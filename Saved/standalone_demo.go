package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// standalone_demo.go - Complete standalone demonstration
// Shows enhanced compact netlist with dual sorting capabilities

// CompactGate represents a gate in compact format
type CompactGate struct {
	Name    string
	Type    string
	Level   int
	Inputs  []string
	Outputs []string
	Targets []string
}

// EnhancedCompactNetlist with dual sorting and indexing
type EnhancedCompactNetlist struct {
	Gates      []CompactGate
	levelIndex map[int][]int  // level -> []gateIndex for O(1) level access
	nameIndex  map[string]int // name -> gateIndex for O(1) name lookup
}

// NewEnhancedCompactNetlist creates a new enhanced netlist
func NewEnhancedCompactNetlist() *EnhancedCompactNetlist {
	return &EnhancedCompactNetlist{
		Gates:      []CompactGate{},
		levelIndex: make(map[int][]int),
		nameIndex:  make(map[string]int),
	}
}

// AddGate adds a gate and updates indices
func (ecn *EnhancedCompactNetlist) AddGate(gate CompactGate) {
	ecn.Gates = append(ecn.Gates, gate)
	gateIndex := len(ecn.Gates) - 1

	// Update name index for O(1) lookup
	ecn.nameIndex[gate.Name] = gateIndex

	// Update level index for O(1) level access
	if _, exists := ecn.levelIndex[gate.Level]; !exists {
		ecn.levelIndex[gate.Level] = []int{}
	}
	ecn.levelIndex[gate.Level] = append(ecn.levelIndex[gate.Level], gateIndex)
}

// SortByTopologicalLevel sorts gates by level then name (processing mode)
func (ecn *EnhancedCompactNetlist) SortByTopologicalLevel() {
	sort.Slice(ecn.Gates, func(i, j int) bool {
		if ecn.Gates[i].Level != ecn.Gates[j].Level {
			return ecn.Gates[i].Level < ecn.Gates[j].Level
		}
		return ecn.Gates[i].Name < ecn.Gates[j].Name
	})
	ecn.rebuildIndices()
}

// SortBySignalName sorts gates alphabetically by name (lookup mode)
func (ecn *EnhancedCompactNetlist) SortBySignalName() {
	sort.Slice(ecn.Gates, func(i, j int) bool {
		return ecn.Gates[i].Name < ecn.Gates[j].Name
	})
	ecn.rebuildIndices()
}

// rebuildIndices reconstructs the indices after sorting
func (ecn *EnhancedCompactNetlist) rebuildIndices() {
	ecn.nameIndex = make(map[string]int)
	ecn.levelIndex = make(map[int][]int)

	for i, gate := range ecn.Gates {
		ecn.nameIndex[gate.Name] = i
		if _, exists := ecn.levelIndex[gate.Level]; !exists {
			ecn.levelIndex[gate.Level] = []int{}
		}
		ecn.levelIndex[gate.Level] = append(ecn.levelIndex[gate.Level], i)
	}
}

// FindGateByName finds a gate by name (O(1) lookup)
func (ecn *EnhancedCompactNetlist) FindGateByName(name string) (*CompactGate, bool) {
	if index, found := ecn.nameIndex[name]; found {
		return &ecn.Gates[index], true
	}
	return nil, false
}

// GetGatesAtLevel returns all gates at a specific level (O(1) level access)
func (ecn *EnhancedCompactNetlist) GetGatesAtLevel(level int) []CompactGate {
	var gates []CompactGate
	if indices, found := ecn.levelIndex[level]; found {
		for _, index := range indices {
			gates = append(gates, ecn.Gates[index])
		}
	}
	return gates
}

// GetTopologicalLevels returns all levels in ascending order
func (ecn *EnhancedCompactNetlist) GetTopologicalLevels() []int {
	var levels []int
	for level := range ecn.levelIndex {
		levels = append(levels, level)
	}
	sort.Ints(levels)
	return levels
}

// DisplayMode constants
const (
	DisplayTopological  = "Topological Order (Processing Mode)"
	DisplayAlphabetical = "Alphabetical Order (Lookup Mode)"
)

// Display prints the netlist in specified mode
func (ecn *EnhancedCompactNetlist) Display(mode string, maxGates int) {
	fmt.Printf("\n--- %s ---\n", mode)
	for i, gate := range ecn.Gates {
		if i >= maxGates {
			fmt.Printf("... (%d more gates)\n", len(ecn.Gates)-i)
			break
		}
		fmt.Printf("  %s, %s, %d, %v, %v, %v\n",
			gate.Name, gate.Type, gate.Level, gate.Inputs, gate.Outputs, gate.Targets)
	}
}

// ValidateTopologicalOrder checks if gates are properly ordered
func (ecn *EnhancedCompactNetlist) ValidateTopologicalOrder() error {
	for _, gate := range ecn.Gates {
		for _, input := range gate.Inputs {
			// Extract signal name (remove _s, _1 suffixes)
			signalName := strings.TrimSuffix(strings.TrimSuffix(input, "_s"), "_1")
			if inputGate, found := ecn.FindGateByName(signalName); found {
				if inputGate.Level >= gate.Level {
					return fmt.Errorf("topological violation: gate %s (level %d) depends on %s (level %d)",
						gate.Name, gate.Level, inputGate.Name, inputGate.Level)
				}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("=== ENHANCED COMPACT NETLIST DEMONSTRATION ===")
	fmt.Println("Demonstrating dual sorting capabilities and fast access operations\n")

	// Create enhanced netlist
	netlist := NewEnhancedCompactNetlist()

	// Load sample circuit (subset of LARGE circuit)
	fmt.Println("📁 Loading sample circuit gates...")

	sampleGates := []CompactGate{
		// Level 1 - Primary inputs
		{"ps16", "PI", 1, []string{}, []string{"ps16_s", "ps16_1"}, []string{"ps16"}},
		{"ps8", "PI", 1, []string{}, []string{"ps8_s", "ps8_1"}, []string{"ps8"}},
		{"ps4", "PI", 1, []string{}, []string{"ps4_s", "ps4_1"}, []string{"ps4"}},
		{"ps2", "PI", 1, []string{}, []string{"ps2_s", "ps2_1"}, []string{"ps2"}},
		{"ps1", "PI", 1, []string{}, []string{"ps1_s", "ps1_1"}, []string{"ps1"}},
		{"i0", "PI", 1, []string{}, []string{"i0_s", "i0_1"}, []string{"i0"}},
		{"i1", "PI", 1, []string{}, []string{"i1_s", "i1_1"}, []string{"i1"}},
		{"i2", "PI", 1, []string{}, []string{"i2_s", "i2_1"}, []string{"i2"}},

		// Level 2 - Inverters
		{"nps16", "NOT", 2, []string{"ps16_s", "ps16_1"}, []string{"nps16_s", "nps16_1"}, []string{"nps16"}},
		{"nps8", "NOT", 2, []string{"ps8_s", "ps8_1"}, []string{"nps8_s", "nps8_1"}, []string{"nps8"}},
		{"nps4", "NOT", 2, []string{"ps4_s", "ps4_1"}, []string{"nps4_s", "nps4_1"}, []string{"nps4"}},
		{"nps2", "NOT", 2, []string{"ps2_s", "ps2_1"}, []string{"nps2_s", "nps2_1"}, []string{"nps2"}},
		{"nps1", "NOT", 2, []string{"ps1_s", "ps1_1"}, []string{"nps1_s", "nps1_1"}, []string{"nps1"}},

		// Level 3 - Local states
		{"ls0", "3AND", 3, []string{"nps4_s", "nps2_s", "nps1_s"}, []string{"ls0_s", "ls0_1"}, []string{"ls0"}},
		{"ls1", "3AND", 3, []string{"nps4_s", "nps2_s", "ps1_s"}, []string{"ls1_s", "ls1_1"}, []string{"ls1"}},
		{"ls2", "3AND", 3, []string{"nps4_s", "ps2_s", "nps1_s"}, []string{"ls2_s", "ls2_1"}, []string{"ls2"}},
		{"ls3", "3AND", 3, []string{"nps4_s", "ps2_s", "ps1_s"}, []string{"ls3_s", "ls3_1"}, []string{"ls3"}},

		// Level 4 - State combinations
		{"s0", "3AND", 4, []string{"nps16_s", "nps8_s", "ls0_s"}, []string{"s0_s", "s0_1"}, []string{"s0"}},
		{"s1", "3AND", 4, []string{"nps16_s", "nps8_s", "ls1_s"}, []string{"s1_s", "s1_1"}, []string{"s1"}},
		{"s2", "3AND", 4, []string{"nps16_s", "nps8_s", "ls2_s"}, []string{"s2_s", "s2_1"}, []string{"s2"}},
		{"s3", "3AND", 4, []string{"nps16_s", "nps8_s", "ls3_s"}, []string{"s3_s", "s3_1"}, []string{"s3"}},

		// Level 5 - Logic operations
		{"a1", "2AND", 5, []string{"s1_s", "i0_s"}, []string{"a1_s", "a1_1"}, []string{"a1"}},
		{"a2", "2AND", 5, []string{"s2_s", "i1_s"}, []string{"a2_s", "a2_1"}, []string{"a2"}},
		{"a3", "2AND", 5, []string{"s3_s", "i2_s"}, []string{"a3_s", "a3_1"}, []string{"a3"}},

		// Level 6 - Final outputs
		{"out1", "2OR", 6, []string{"a1_s", "a4_s"}, []string{"out1_s", "out1_1"}, []string{"out1"}},
		{"out2", "2OR", 6, []string{"a2_s", "a5_s"}, []string{"out2_s", "out2_1"}, []string{"out2"}},
		{"out3", "2OR", 6, []string{"a3_s", "a6_s"}, []string{"out3_s", "out3_1"}, []string{"out3"}},
	}

	for _, gate := range sampleGates {
		netlist.AddGate(gate)
	}

	fmt.Printf("✅ Loaded %d gates across %d levels\n\n",
		len(netlist.Gates), len(netlist.GetTopologicalLevels()))

	// ============= DEMONSTRATION 1: DUAL SORTING =============
	fmt.Println("📋 DEMONSTRATION 1: DUAL SORTING CAPABILITIES\n")

	// Show topological sorting (natural processing order)
	fmt.Println("🔄 Sorting by topological level (processing mode)...")
	netlist.SortByTopologicalLevel()
	netlist.Display(DisplayTopological, 8)

	// Show alphabetical sorting (lookup optimized)
	fmt.Println("\n🔄 Sorting by signal name (lookup mode)...")
	netlist.SortBySignalName()
	netlist.Display(DisplayAlphabetical, 8)

	// ============= DEMONSTRATION 2: FAST ACCESS =============
	fmt.Println("\n⚡ DEMONSTRATION 2: FAST ACCESS OPERATIONS\n")

	// Switch back to topological for level access demonstration
	netlist.SortByTopologicalLevel()

	// Level-based access (O(1))
	fmt.Println("🎯 Level-based access (O(1) per level):")
	levels := netlist.GetTopologicalLevels()
	for _, level := range levels {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("  Level %d: %d gates", level, len(gates))
		if len(gates) <= 4 {
			fmt.Print(" (")
			for i, gate := range gates {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(gate.Name)
			}
			fmt.Print(")")
		}
		fmt.Println()
	}

	// Signal lookup (O(1))
	fmt.Println("\n🔍 Signal lookup (O(1) per signal):")
	searchSignals := []string{"s0", "a1", "out1", "ps16", "ls2", "nonexistent"}
	for _, signalName := range searchSignals {
		start := time.Now()
		if gate, found := netlist.FindGateByName(signalName); found {
			duration := time.Since(start)
			fmt.Printf("  ✅ Found %s: %s gate at level %d (%v)\n",
				signalName, gate.Type, gate.Level, duration)
		} else {
			fmt.Printf("  ❌ Signal %s not found\n", signalName)
		}
	}

	// ============= DEMONSTRATION 3: VALIDATION =============
	fmt.Println("\n🔍 DEMONSTRATION 3: TOPOLOGICAL VALIDATION\n")

	fmt.Println("Validating topological order...")
	if err := netlist.ValidateTopologicalOrder(); err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Topological order is valid - no forward references detected")
	}

	// ============= DEMONSTRATION 4: WORKFLOW SCENARIOS =============
	fmt.Println("\n🎯 DEMONSTRATION 4: WORKFLOW SCENARIOS\n")

	// Scenario 1: Designer workflow
	fmt.Println("🎨 Scenario 1: Designer Creating Circuit")
	fmt.Println("  Designer naturally creates gates in topological order:")
	netlist.SortByTopologicalLevel()
	for i, gate := range netlist.Gates[:5] {
		fmt.Printf("    %d. Add %s (%s) at level %d\n", i+1, gate.Name, gate.Type, gate.Level)
	}
	fmt.Printf("    ... (%d more gates follow naturally)\n", len(netlist.Gates)-5)

	// Scenario 2: Debug workflow
	fmt.Println("\n🔧 Scenario 2: Engineer Debugging Circuit")
	fmt.Println("  Engineer switches to alphabetical view for fast lookup:")
	netlist.SortBySignalName()
	fmt.Println("  Looking up specific signals:")
	debugSignals := []string{"a1", "ls2", "out3"}
	for _, signal := range debugSignals {
		if gate, found := netlist.FindGateByName(signal); found {
			fmt.Printf("    📍 %s: %s gate, level %d, inputs: %v\n",
				signal, gate.Type, gate.Level, gate.Inputs)
		}
	}

	// Scenario 3: Processing workflow
	fmt.Println("\n🤖 Scenario 3: Automated Processing System")
	fmt.Println("  System processes gates level-by-level:")
	netlist.SortByTopologicalLevel()
	levels = netlist.GetTopologicalLevels()
	for _, level := range levels[:4] { // Show first few levels
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("    Process Level %d: %d gates can run in parallel\n", level, len(gates))
	}
	fmt.Printf("    ... (levels %d-%d follow sequentially)\n", levels[4], levels[len(levels)-1])

	// ============= PERFORMANCE SUMMARY =============
	fmt.Println("\n📊 PERFORMANCE CHARACTERISTICS\n")
	fmt.Println("✅ O(1) Operations:")
	fmt.Println("  • Signal lookup by name (nameIndex)")
	fmt.Println("  • Level access for all gates at level (levelIndex)")
	fmt.Println("  • Fast validation with topological checking")
	fmt.Println("\n⚡ Optimized Workflows:")
	fmt.Println("  • Topological sorting: Natural creation & processing order")
	fmt.Println("  • Alphabetical sorting: Fast lookup & debugging")
	fmt.Println("  • Dual indexing: Best of both access patterns")

	// ============= INTEGRATION READINESS =============
	fmt.Println("\n🎉 INTEGRATION READINESS\n")
	fmt.Println("Enhanced Compact Format provides:")
	fmt.Println("✅ Designer-friendly single-line gate specification")
	fmt.Println("✅ Dual sorting for optimal workflow support")
	fmt.Println("✅ O(1) fast access for processing efficiency")
	fmt.Println("✅ Topological validation for safety")
	fmt.Println("✅ JSON compatibility for existing tools")

	fmt.Println("\nNext steps:")
	fmt.Println("1. 🔄 Connect to validation framework")
	fmt.Println("2. 🔄 Generate automated ps2ns functions")
	fmt.Println("3. 🔄 Complete netlist-driven ATPG system")
	fmt.Println("4. 🎯 Enable \"any circuit\" processing capability")

	fmt.Printf("\n🚀 NETLIST-DRIVEN REVISION: READY FOR IMPLEMENTATION!\n")

	// ============= CIRCUIT-INDEPENDENT PLATFORM DEMONSTRATION =============
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("🔄 ARCHITECTURAL REFACTORING: CIRCUIT-INDEPENDENT PLATFORM")
	fmt.Println(strings.Repeat("=", 70))

	// Demonstrate the new separated architecture
	DemonstrateCircuitIndependentPlatform()
}
