package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Compact netlist format for designer input
// Format: <sig_name, sig_type, level, inputs, outputs, targets>
// Example: b25, 2OR, 5, [s23, a4], [b25_s, b25_1], [b25:1, b25:0]

type CompactGate struct {
	Name         string   // Signal name (e.g., "b25")
	Type         string   // Gate type (e.g., "2OR", "3AND", "NOT", "PI")
	Level        int      // Topological level number
	Inputs       []string // Input signals
	Outputs      []string // Output signals
	FaultTargets []string // Fault target specifications
}

type CompactNetlist struct {
	Name        string
	Description string
	Gates       []CompactGate
}

// Parse compact netlist format from text file
func ParseCompactNetlist(filename string) (*CompactNetlist, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	netlist := &CompactNetlist{
		Name:        "Circuit from compact format",
		Description: "Parsed from designer input",
		Gates:       []CompactGate{},
	}

	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		gate, err := parseCompactGateLine(line, lineNum)
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", lineNum, err)
		}

		netlist.Gates = append(netlist.Gates, gate)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return netlist, nil
}

// Parse a single line of compact format
// Format: sig_name, sig_type, level, [inputs], [outputs], [targets]
func parseCompactGateLine(line string, lineNum int) (CompactGate, error) {
	gate := CompactGate{}

	// Remove angle brackets if present
	line = strings.Trim(line, "<>")

	// Split by comma and clean whitespace
	parts := strings.Split(line, ",")
	if len(parts) != 6 {
		return gate, fmt.Errorf("expected 6 fields, got %d", len(parts))
	}

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	// Parse signal name
	gate.Name = parts[0]
	if gate.Name == "" {
		return gate, fmt.Errorf("signal name cannot be empty")
	}

	// Parse gate type
	gate.Type = parts[1]
	if gate.Type == "" {
		return gate, fmt.Errorf("gate type cannot be empty")
	}

	// Parse level
	level, err := strconv.Atoi(parts[2])
	if err != nil {
		return gate, fmt.Errorf("invalid level number '%s': %v", parts[2], err)
	}
	gate.Level = level

	// Parse inputs array
	gate.Inputs, err = parseStringArray(parts[3])
	if err != nil {
		return gate, fmt.Errorf("invalid inputs format: %v", err)
	}

	// Parse outputs array
	gate.Outputs, err = parseStringArray(parts[4])
	if err != nil {
		return gate, fmt.Errorf("invalid outputs format: %v", err)
	}

	// Parse fault targets array
	gate.FaultTargets, err = parseStringArray(parts[5])
	if err != nil {
		return gate, fmt.Errorf("invalid fault targets format: %v", err)
	}

	return gate, nil
}

// Parse array format [item1, item2, item3] or [item1,item2,item3]
func parseStringArray(arrayStr string) ([]string, error) {
	arrayStr = strings.TrimSpace(arrayStr)

	// Remove brackets
	if !strings.HasPrefix(arrayStr, "[") || !strings.HasSuffix(arrayStr, "]") {
		return nil, fmt.Errorf("array must be enclosed in brackets []")
	}

	content := strings.Trim(arrayStr, "[]")
	if content == "" {
		return []string{}, nil
	}

	// Split by comma and clean
	items := strings.Split(content, ",")
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = strings.TrimSpace(item)
	}

	return result, nil
}

// Convert compact netlist to JSON format
func (cn *CompactNetlist) ToJSON() (*Circuit, error) {
	circuit := &Circuit{
		Name:        cn.Name,
		Description: cn.Description,
		Levels:      []Level{},
	}

	// Group gates by level
	levelMap := make(map[int][]Gate)
	maxLevel := 0

	for _, cgate := range cn.Gates {
		if cgate.Level > maxLevel {
			maxLevel = cgate.Level
		}

		// Convert compact gate to JSON gate format
		jsonGate := Gate{
			Name:         cgate.Name,
			Type:         normalizeGateType(cgate.Type),
			Inputs:       cgate.Inputs,
			Outputs:      cgate.Outputs,
			FaultTargets: cgate.FaultTargets,
		}

		levelMap[cgate.Level] = append(levelMap[cgate.Level], jsonGate)
	}

	// Create level structures
	for level := 1; level <= maxLevel; level++ {
		if gates, exists := levelMap[level]; exists {
			levelStruct := Level{
				Level:       level,
				Description: fmt.Sprintf("Level %d gates", level),
				Gates:       gates,
			}
			circuit.Levels = append(circuit.Levels, levelStruct)
		}
	}

	// Extract primary inputs, outputs, and next states
	circuit.PrimaryInputs, circuit.CircuitInputs, circuit.Outputs, circuit.NextStates = extractCircuitPorts(cn.Gates)

	return circuit, nil
}

// Normalize gate type from compact to JSON format
func normalizeGateType(compactType string) string {
	switch strings.ToUpper(compactType) {
	case "PI":
		return "PI"
	case "NOT", "1NOT":
		return "NOT"
	case "2AND", "AND2":
		return "AND2"
	case "3AND", "AND3":
		return "AND3"
	case "2OR", "OR2":
		return "OR2"
	case "3OR", "OR3":
		return "OR3"
	default:
		return compactType // Return as-is if unknown
	}
}

// Extract circuit ports from gate list
func extractCircuitPorts(gates []CompactGate) ([]string, []string, []string, []string) {
	var primaryInputs, circuitInputs, outputs, nextStates []string

	for _, gate := range gates {
		if gate.Type == "PI" {
			if strings.HasPrefix(gate.Name, "ps") {
				primaryInputs = append(primaryInputs, gate.Name+"_i")
			} else if strings.HasPrefix(gate.Name, "i") {
				circuitInputs = append(circuitInputs, gate.Name)
			}
		}

		for _, output := range gate.Outputs {
			if strings.HasPrefix(output, "out") && strings.HasSuffix(output, "_s") {
				outputs = append(outputs, output)
			} else if strings.HasPrefix(output, "ns") && strings.HasSuffix(output, "_s") {
				nextStates = append(nextStates, output)
			}
		}
	}

	return primaryInputs, circuitInputs, outputs, nextStates
}

// Sort netlist by topological level or signal name
func (cn *CompactNetlist) SortByLevel() {
	sort.Slice(cn.Gates, func(i, j int) bool {
		if cn.Gates[i].Level == cn.Gates[j].Level {
			return cn.Gates[i].Name < cn.Gates[j].Name
		}
		return cn.Gates[i].Level < cn.Gates[j].Level
	})
}

func (cn *CompactNetlist) SortByName() {
	sort.Slice(cn.Gates, func(i, j int) bool {
		return cn.Gates[i].Name < cn.Gates[j].Name
	})
}

// Export compact netlist to text file
func (cn *CompactNetlist) ExportToFile(filename string) error {
	var lines []string

	// Add header
	lines = append(lines, "# Compact Circuit Netlist")
	lines = append(lines, fmt.Sprintf("# Name: %s", cn.Name))
	lines = append(lines, fmt.Sprintf("# Description: %s", cn.Description))
	lines = append(lines, "# Format: <sig_name, sig_type, level, [inputs], [outputs], [targets]>")
	lines = append(lines, "")

	// Add gates
	for _, gate := range cn.Gates {
		line := fmt.Sprintf("%s, %s, %d, [%s], [%s], [%s]",
			gate.Name,
			gate.Type,
			gate.Level,
			strings.Join(gate.Inputs, ", "),
			strings.Join(gate.Outputs, ", "),
			strings.Join(gate.FaultTargets, ", "))
		lines = append(lines, line)
	}

	content := strings.Join(lines, "\n")
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// Convert JSON circuit back to compact format
func JSONToCompact(circuit *Circuit) *CompactNetlist {
	netlist := &CompactNetlist{
		Name:        circuit.Name,
		Description: circuit.Description,
		Gates:       []CompactGate{},
	}

	for _, level := range circuit.Levels {
		for _, gate := range level.Gates {
			compactGate := CompactGate{
				Name:         gate.Name,
				Type:         gate.Type,
				Level:        level.Level,
				Inputs:       gate.Inputs,
				Outputs:      gate.Outputs,
				FaultTargets: gate.FaultTargets,
			}
			netlist.Gates = append(netlist.Gates, compactGate)
		}
	}

	return netlist
}

// Generate compact format from existing JSON
func ConvertJSONToCompact(jsonFile, compactFile string) error {
	// Read JSON file
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	var circuit Circuit
	err = json.Unmarshal(data, &circuit)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Convert to compact format
	compactNetlist := JSONToCompact(&circuit)
	compactNetlist.SortByLevel()

	// Export to compact file
	err = compactNetlist.ExportToFile(compactFile)
	if err != nil {
		return fmt.Errorf("failed to write compact file: %v", err)
	}

	fmt.Printf("Converted %s to %s\n", jsonFile, compactFile)
	fmt.Printf("Compact format has %d gates across %d levels\n",
		len(compactNetlist.Gates), getMaxLevel(compactNetlist.Gates))

	return nil
}

func getMaxLevel(gates []CompactGate) int {
	maxLevel := 0
	for _, gate := range gates {
		if gate.Level > maxLevel {
			maxLevel = gate.Level
		}
	}
	return maxLevel
}

// Enhanced sorting and display options for compact netlists

// Display modes for netlist presentation
type DisplayMode int

const (
	DisplayTopological  DisplayMode = iota // Sorted by topological level (processing order)
	DisplayAlphabetical                    // Sorted by signal name (lookup/reference)
)

// Enhanced compact netlist with sorting and display capabilities
type EnhancedCompactNetlist struct {
	CompactNetlist
	currentSort DisplayMode
	levelIndex  map[int][]int  // Level -> gate indices for fast level access
	nameIndex   map[string]int // Signal name -> gate index for fast lookup
}

// Create enhanced netlist with indexing
func NewEnhancedCompactNetlist() *EnhancedCompactNetlist {
	return &EnhancedCompactNetlist{
		CompactNetlist: CompactNetlist{
			Gates: []CompactGate{},
		},
		currentSort: DisplayTopological,
		levelIndex:  make(map[int][]int),
		nameIndex:   make(map[string]int),
	}
}

// Load and automatically index the netlist
func (ecn *EnhancedCompactNetlist) LoadFromFile(filename string) error {
	// Load using existing parser
	netlist, err := ParseCompactNetlist(filename)
	if err != nil {
		return err
	}

	ecn.CompactNetlist = *netlist
	ecn.rebuildIndices()

	// Default to topological sort since that's how designers create circuits
	ecn.SortByTopologicalLevel()

	return nil
}

// Rebuild internal indices for fast access
func (ecn *EnhancedCompactNetlist) rebuildIndices() {
	ecn.levelIndex = make(map[int][]int)
	ecn.nameIndex = make(map[string]int)

	for i, gate := range ecn.Gates {
		// Build level index
		ecn.levelIndex[gate.Level] = append(ecn.levelIndex[gate.Level], i)

		// Build name index
		ecn.nameIndex[gate.Name] = i
	}
}

// Sort by topological level (primary) and signal name (secondary)
func (ecn *EnhancedCompactNetlist) SortByTopologicalLevel() {
	sort.Slice(ecn.Gates, func(i, j int) bool {
		// Primary sort: topological level
		if ecn.Gates[i].Level != ecn.Gates[j].Level {
			return ecn.Gates[i].Level < ecn.Gates[j].Level
		}
		// Secondary sort: signal name (for gates at same level)
		return ecn.Gates[i].Name < ecn.Gates[j].Name
	})

	ecn.currentSort = DisplayTopological
	ecn.rebuildIndices()

	fmt.Printf("Sorted netlist by topological level (%d gates)\n", len(ecn.Gates))
}

// Sort alphabetically by signal name
func (ecn *EnhancedCompactNetlist) SortBySignalName() {
	sort.Slice(ecn.Gates, func(i, j int) bool {
		return ecn.Gates[i].Name < ecn.Gates[j].Name
	})

	ecn.currentSort = DisplayAlphabetical
	ecn.rebuildIndices()

	fmt.Printf("Sorted netlist alphabetically by signal name (%d gates)\n", len(ecn.Gates))
}

// GetLevelStatistics returns processing statistics for each level
func (ecn *EnhancedCompactNetlist) GetLevelStatistics() map[int]int {
	stats := make(map[int]int)
	for level, indices := range ecn.levelIndex {
		stats[level] = len(indices)
	}
	return stats
}

// ValidateNoSameLevelDependencies ensures no gate depends on another gate at the same level
func (ecn *EnhancedCompactNetlist) ValidateNoSameLevelDependencies() error {
	for _, gate := range ecn.Gates {
		for _, inputSignal := range gate.Inputs {
			// Extract signal name (remove _s, _1 suffixes)
			signalName := strings.TrimSuffix(strings.TrimSuffix(inputSignal, "_s"), "_1")

			if inputGate, found := ecn.FindGateByName(signalName); found {
				if inputGate.Level == gate.Level {
					return fmt.Errorf("same-level dependency violation: gate %s (level %d) depends on %s (also level %d)",
						gate.Name, gate.Level, inputGate.Name, inputGate.Level)
				}
			}
		}
	}
	return nil
}

// AddGate adds a gate to the enhanced netlist and updates indices
func (ecn *EnhancedCompactNetlist) AddGate(gate CompactGate) {
	ecn.Gates = append(ecn.Gates, gate)
	gateIndex := len(ecn.Gates) - 1

	// Update name index for O(1) lookup
	ecn.nameIndex[gate.Name] = gateIndex

	// Update level index for O(1) level access
	if ecn.levelIndex[gate.Level] == nil {
		ecn.levelIndex[gate.Level] = []int{}
	}
	ecn.levelIndex[gate.Level] = append(ecn.levelIndex[gate.Level], gateIndex)
}

// Get gates at a specific topological level (fast access)
func (ecn *EnhancedCompactNetlist) GetGatesAtLevel(level int) []CompactGate {
	indices, exists := ecn.levelIndex[level]
	if !exists {
		return []CompactGate{}
	}

	gates := make([]CompactGate, len(indices))
	for i, idx := range indices {
		gates[i] = ecn.Gates[idx]
	}

	return gates
}

// Find gate by signal name (fast lookup)
func (ecn *EnhancedCompactNetlist) FindGateByName(signalName string) (*CompactGate, bool) {
	idx, exists := ecn.nameIndex[signalName]
	if !exists {
		return nil, false
	}

	return &ecn.Gates[idx], true
}

// Get all topological levels in order
func (ecn *EnhancedCompactNetlist) GetTopologicalLevels() []int {
	levels := make([]int, 0, len(ecn.levelIndex))
	for level := range ecn.levelIndex {
		levels = append(levels, level)
	}
	sort.Ints(levels)
	return levels
}

// Display netlist in current sort order with formatting
func (ecn *EnhancedCompactNetlist) Display(mode DisplayMode) {
	// Switch to requested display mode if different
	if mode != ecn.currentSort {
		switch mode {
		case DisplayTopological:
			ecn.SortByTopologicalLevel()
		case DisplayAlphabetical:
			ecn.SortBySignalName()
		}
	}

	fmt.Printf("\n=== COMPACT NETLIST DISPLAY ===\n")
	fmt.Printf("Circuit: %s\n", ecn.Name)
	fmt.Printf("Description: %s\n", ecn.Description)

	switch mode {
	case DisplayTopological:
		ecn.displayTopological()
	case DisplayAlphabetical:
		ecn.displayAlphabetical()
	}

	fmt.Printf("Total gates: %d\n", len(ecn.Gates))
}

// Display in topological order with level grouping
func (ecn *EnhancedCompactNetlist) displayTopological() {
	fmt.Printf("Sort order: Topological (Level → Signal Name)\n")
	fmt.Printf("Format: <sig_name, sig_type, level, [inputs], [outputs], [targets]>\n\n")

	levels := ecn.GetTopologicalLevels()

	for _, level := range levels {
		gates := ecn.GetGatesAtLevel(level)

		fmt.Printf("# Level %d (%d gates)\n", level, len(gates))

		for _, gate := range gates {
			fmt.Printf("%s, %s, %d, [%s], [%s], [%s]\n",
				gate.Name,
				gate.Type,
				gate.Level,
				strings.Join(gate.Inputs, ", "),
				strings.Join(gate.Outputs, ", "),
				strings.Join(gate.FaultTargets, ", "))
		}
		fmt.Println() // Blank line between levels
	}
}

// Display in alphabetical order with signal grouping
func (ecn *EnhancedCompactNetlist) displayAlphabetical() {
	fmt.Printf("Sort order: Alphabetical (Signal Name)\n")
	fmt.Printf("Format: <sig_name, sig_type, level, [inputs], [outputs], [targets]>\n\n")

	// Group by first letter for easier navigation
	currentPrefix := ""

	for _, gate := range ecn.Gates {
		prefix := strings.ToUpper(string(gate.Name[0]))

		if prefix != currentPrefix {
			if currentPrefix != "" {
				fmt.Println() // Blank line between letter groups
			}
			fmt.Printf("# Signals starting with '%s'\n", prefix)
			currentPrefix = prefix
		}

		fmt.Printf("%s, %s, %d, [%s], [%s], [%s]\n",
			gate.Name,
			gate.Type,
			gate.Level,
			strings.Join(gate.Inputs, ", "),
			strings.Join(gate.Outputs, ", "),
			strings.Join(gate.FaultTargets, ", "))
	}
}

// Export netlist in specified sort order
func (ecn *EnhancedCompactNetlist) ExportSorted(filename string, mode DisplayMode) error {
	// Ensure correct sort order
	if mode != ecn.currentSort {
		switch mode {
		case DisplayTopological:
			ecn.SortByTopologicalLevel()
		case DisplayAlphabetical:
			ecn.SortBySignalName()
		}
	}

	var lines []string

	// Add header with sort information
	lines = append(lines, "# Compact Circuit Netlist")
	lines = append(lines, fmt.Sprintf("# Name: %s", ecn.Name))
	lines = append(lines, fmt.Sprintf("# Description: %s", ecn.Description))

	switch mode {
	case DisplayTopological:
		lines = append(lines, "# Sort Order: Topological (Level → Signal Name)")
		lines = append(lines, "# Optimized for sequential processing")
	case DisplayAlphabetical:
		lines = append(lines, "# Sort Order: Alphabetical (Signal Name)")
		lines = append(lines, "# Optimized for lookup and reference")
	}

	lines = append(lines, "# Format: <sig_name, sig_type, level, [inputs], [outputs], [targets]>")
	lines = append(lines, "")

	// Add gates with level grouping for topological sort
	if mode == DisplayTopological {
		levels := ecn.GetTopologicalLevels()

		for _, level := range levels {
			lines = append(lines, fmt.Sprintf("# Level %d", level))
			gates := ecn.GetGatesAtLevel(level)

			for _, gate := range gates {
				line := fmt.Sprintf("%s, %s, %d, [%s], [%s], [%s]",
					gate.Name,
					gate.Type,
					gate.Level,
					strings.Join(gate.Inputs, ", "),
					strings.Join(gate.Outputs, ", "),
					strings.Join(gate.FaultTargets, ", "))
				lines = append(lines, line)
			}
			lines = append(lines, "") // Blank line between levels
		}
	} else {
		// Alphabetical - add all gates in order
		for _, gate := range ecn.Gates {
			line := fmt.Sprintf("%s, %s, %d, [%s], [%s], [%s]",
				gate.Name,
				gate.Type,
				gate.Level,
				strings.Join(gate.Inputs, ", "),
				strings.Join(gate.Outputs, ", "),
				strings.Join(gate.FaultTargets, ", "))
			lines = append(lines, line)
		}
	}

	content := strings.Join(lines, "\n")
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// Validate topological ordering (ensure no forward references)
func (ecn *EnhancedCompactNetlist) ValidateTopologicalOrder() error {
	fmt.Println("Validating topological order...")

	// Build signal definition levels
	signalLevels := make(map[string]int)

	for _, gate := range ecn.Gates {
		for _, output := range gate.Outputs {
			signalLevels[output] = gate.Level
		}
	}

	// Check each gate's inputs are defined at lower levels
	for _, gate := range ecn.Gates {
		for _, input := range gate.Inputs {
			inputLevel, exists := signalLevels[input]
			if !exists {
				return fmt.Errorf("gate %s (level %d) uses undefined input signal %s",
					gate.Name, gate.Level, input)
			}

			if inputLevel >= gate.Level {
				return fmt.Errorf("gate %s (level %d) uses input %s from level %d (forward reference)",
					gate.Name, gate.Level, input, inputLevel)
			}
		}
	}

	fmt.Printf("✅ Topological order validated: %d gates, %d levels\n",
		len(ecn.Gates), len(ecn.GetTopologicalLevels()))

	return nil
}

// Processing optimization: get gates organized by topological level
// CRITICAL: Gates within each level can be processed in ANY ORDER or IN PARALLEL
// since they have no dependencies on each other (guaranteed by topological sorting)
func (ecn *EnhancedCompactNetlist) GetProcessingOrder() [][]CompactGate {
	levels := ecn.GetTopologicalLevels()
	result := make([][]CompactGate, len(levels))

	for i, level := range levels {
		result[i] = ecn.GetGatesAtLevel(level)
	}

	return result
}

// Demonstration functions
func DemoCompactFormat() {
	fmt.Println("=== COMPACT NETLIST FORMAT DEMO ===")

	// Example compact format entries
	examples := []string{
		"ps16, PI, 1, [], [ps16_s, ps16_1], [ps16]",
		"nps16, NOT, 1, [ps16_s, ps16_1], [nps16_s, nps16_1], [nps16]",
		"ls0, 3AND, 2, [nps4_s, nps2_s, nps1_s, nps4_1, nps2_1, nps1_1], [ls0_s, ls0_1], [ls0]",
		"s0, 3AND, 3, [nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1], [s0_s, s0_1], [s0]",
		"a1, 2AND, 4, [s10_s, i0_s, s10_1, i0_1], [a1_s, a1_1], [a1]",
		"b25, 2OR, 5, [s23_s, a4_s, s23_1, a4_1], [b25_s, b25_1], [b25]",
		"out4, 2OR, 10, [f1_s, f2_s, f1_1, f2_1], [out4_s, out4_1], [out4]",
	}

	fmt.Println("Example compact format entries:")
	for i, example := range examples {
		fmt.Printf("%d: %s\n", i+1, example)
	}

	fmt.Println("\nAdvantages of compact format:")
	fmt.Println("✅ Single line per gate - easy to read and edit")
	fmt.Println("✅ Topological level clearly visible")
	fmt.Println("✅ All gate information in one place")
	fmt.Println("✅ Can be sorted by level or signal name")
	fmt.Println("✅ Much more compact than JSON")
	fmt.Println("✅ Easy for designer manual input")
	fmt.Println("✅ Can be converted to JSON internally for processing")
}
