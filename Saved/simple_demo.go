package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Simple demo runner for enhanced compact netlist system

// CompactGate represents a gate in compact format
type CompactGate struct {
	Name    string
	Type    string
	Level   int
	Inputs  []string
	Outputs []string
}

// EnhancedCompactNetlist with dual sorting and indexing
type EnhancedCompactNetlist struct {
	Gates      []CompactGate
	levelIndex map[int][]int  // level -> []gateIndex
	nameIndex  map[string]int // name -> gateIndex
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

	// Update name index
	ecn.nameIndex[gate.Name] = gateIndex

	// Update level index
	if _, exists := ecn.levelIndex[gate.Level]; !exists {
		ecn.levelIndex[gate.Level] = []int{}
	}
	ecn.levelIndex[gate.Level] = append(ecn.levelIndex[gate.Level], gateIndex)
}

// SortByTopologicalLevel sorts gates by level then name
func (ecn *EnhancedCompactNetlist) SortByTopologicalLevel() {
	sort.Slice(ecn.Gates, func(i, j int) bool {
		if ecn.Gates[i].Level != ecn.Gates[j].Level {
			return ecn.Gates[i].Level < ecn.Gates[j].Level
		}
		return ecn.Gates[i].Name < ecn.Gates[j].Name
	})
	ecn.rebuildIndices()
}

// SortBySignalName sorts gates alphabetically by name
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

// Display prints the netlist in the current sort order
func (ecn *EnhancedCompactNetlist) Display(mode string) {
	fmt.Printf("\n--- %s ---\n", mode)
	for i, gate := range ecn.Gates {
		if i >= 10 { // Limit output for demo
			fmt.Printf("... (%d more gates)\n", len(ecn.Gates)-i)
			break
		}
		fmt.Printf("%s, %s, %d, %v, %v\n",
			gate.Name, gate.Type, gate.Level, gate.Inputs, gate.Outputs)
	}
}

func main() {
	fmt.Println("=== ENHANCED COMPACT NETLIST DEMONSTRATION ===\n")

	// Create netlist
	netlist := NewEnhancedCompactNetlist()

	// Load sample gates
	sampleGates := []CompactGate{
		{"ps16", "PI", 1, []string{}, []string{"ps16_s", "ps16_1"}},
		{"ps8", "PI", 1, []string{}, []string{"ps8_s", "ps8_1"}},
		{"ps4", "PI", 1, []string{}, []string{"ps4_s", "ps4_1"}},
		{"ps2", "PI", 1, []string{}, []string{"ps2_s", "ps2_1"}},

		{"nps16", "NOT", 2, []string{"ps16_s", "ps16_1"}, []string{"nps16_s", "nps16_1"}},
		{"nps8", "NOT", 2, []string{"ps8_s", "ps8_1"}, []string{"nps8_s", "nps8_1"}},
		{"nps4", "NOT", 2, []string{"ps4_s", "ps4_1"}, []string{"nps4_s", "nps4_1"}},

		{"ls0", "3AND", 3, []string{"nps4_s", "nps2_s", "nps1_s"}, []string{"ls0_s", "ls0_1"}},
		{"ls1", "3AND", 3, []string{"nps4_s", "nps2_s", "ps1_s"}, []string{"ls1_s", "ls1_1"}},
		{"ls2", "3AND", 3, []string{"nps4_s", "ps2_s", "nps1_s"}, []string{"ls2_s", "ls2_1"}},

		{"s0", "3AND", 4, []string{"nps16_s", "nps8_s", "ls0_s"}, []string{"s0_s", "s0_1"}},
		{"s1", "3AND", 4, []string{"nps16_s", "nps8_s", "ls1_s"}, []string{"s1_s", "s1_1"}},
		{"s2", "3AND", 4, []string{"nps16_s", "nps8_s", "ls2_s"}, []string{"s2_s", "s2_1"}},

		{"a1", "2AND", 5, []string{"s1_s", "i0_s"}, []string{"a1_s", "a1_1"}},
		{"a2", "2AND", 5, []string{"s2_s", "i1_s"}, []string{"a2_s", "a2_1"}},
		{"a3", "2AND", 5, []string{"s3_s", "i2_s"}, []string{"a3_s", "a3_1"}},

		{"out1", "2OR", 6, []string{"a1_s", "a4_s"}, []string{"out1_s", "out1_1"}},
		{"out2", "2OR", 6, []string{"a2_s", "a5_s"}, []string{"out2_s", "out2_1"}},
		{"out3", "2OR", 6, []string{"a3_s", "a6_s"}, []string{"out3_s", "out3_1"}},
	}

	for _, gate := range sampleGates {
		netlist.AddGate(gate)
	}

	fmt.Printf("✅ Loaded %d gates across %d levels\n\n",
		len(netlist.Gates), len(netlist.GetTopologicalLevels()))

	// Demonstrate dual sorting
	fmt.Println("📋 DUAL SORTING DEMONSTRATION")

	// Topological sorting
	netlist.SortByTopologicalLevel()
	netlist.Display("Topological Order (Processing Mode)")

	// Alphabetical sorting
	netlist.SortBySignalName()
	netlist.Display("Alphabetical Order (Lookup Mode)")

	// Fast access demonstration
	fmt.Println("\n⚡ FAST ACCESS OPERATIONS")

	// Switch back to topological for level access
	netlist.SortByTopologicalLevel()

	// Level-based access
	fmt.Println("\n🎯 Level-based access:")
	levels := netlist.GetTopologicalLevels()
	for _, level := range levels {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("  Level %d: %d gates", level, len(gates))
		if len(gates) <= 3 {
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

	// Signal lookup
	fmt.Println("\n🔍 Signal lookup:")
	searchSignals := []string{"s0", "a1", "out1", "ps16", "nonexistent"}
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

	// Validation
	fmt.Println("\n🔍 TOPOLOGICAL VALIDATION")
	if err := netlist.ValidateTopologicalOrder(); err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Topological order is valid")
	}

	// Performance summary
	fmt.Println("\n📊 PERFORMANCE CHARACTERISTICS")
	fmt.Println("  • O(1) signal lookup with name index")
	fmt.Println("  • O(1) level access with level index")
	fmt.Println("  • Dual sorting for optimal workflow support")
	fmt.Println("  • Topological validation ensures processing safety")

	// Workflow summary
	fmt.Println("\n🎯 WORKFLOW INTEGRATION")
	fmt.Println("Designer workflow:")
	fmt.Println("  1. Create gates in topological order (natural)")
	fmt.Println("  2. Switch to alphabetical for debugging/lookup")
	fmt.Println("  3. Export in appropriate format for target use")
	fmt.Println("\nProcessing workflow:")
	fmt.Println("  1. Load compact format (any order)")
	fmt.Println("  2. Sort topologically for processing")
	fmt.Println("  3. Process level-by-level for ps2ns generation")

	fmt.Println("\n🎉 ENHANCED COMPACT NETLIST READY!")
	fmt.Println("Next: Integration with validation framework and automated ps2ns generation")
}
