package main

import (
	"fmt"
	"strings"
)

// Demonstrate the dual sorting and search capabilities
func DemoSortingAndSearch() {
	fmt.Println("=== DUAL SORTING AND SEARCH DEMO ===")

	// Create enhanced netlist
	netlist := NewEnhancedCompactNetlist()
	netlist.Name = "LARGE Circuit - Sorting Demo"
	netlist.Description = "Demonstrates topological vs alphabetical sorting"

	// Add sample gates to demonstrate sorting
	sampleGates := []CompactGate{
		{Name: "ps16", Type: "PI", Level: 1, Inputs: []string{}, Outputs: []string{"ps16_s", "ps16_1"}, FaultTargets: []string{"ps16"}},
		{Name: "a1", Type: "2AND", Level: 4, Inputs: []string{"s10_s", "i0_s"}, Outputs: []string{"a1_s", "a1_1"}, FaultTargets: []string{"a1"}},
		{Name: "nps16", Type: "NOT", Level: 1, Inputs: []string{"ps16_s", "ps16_1"}, Outputs: []string{"nps16_s", "nps16_1"}, FaultTargets: []string{"nps16"}},
		{Name: "out4", Type: "2OR", Level: 10, Inputs: []string{"f1_s", "f2_s"}, Outputs: []string{"out4_s", "out4_1"}, FaultTargets: []string{"out4"}},
		{Name: "b25", Type: "2OR", Level: 5, Inputs: []string{"s23_s", "a4_s"}, Outputs: []string{"b25_s", "b25_1"}, FaultTargets: []string{"b25"}},
		{Name: "ls0", Type: "3AND", Level: 2, Inputs: []string{"nps4_s", "nps2_s", "nps1_s"}, Outputs: []string{"ls0_s", "ls0_1"}, FaultTargets: []string{"ls0"}},
		{Name: "i7", Type: "PI", Level: 1, Inputs: []string{}, Outputs: []string{"i7_s", "i7_1"}, FaultTargets: []string{"i7"}},
		{Name: "s0", Type: "3AND", Level: 3, Inputs: []string{"nps16_s", "nps8_s", "ls0_s"}, Outputs: []string{"s0_s", "s0_1"}, FaultTargets: []string{"s0"}},
	}

	netlist.Gates = sampleGates
	netlist.rebuildIndices()

	fmt.Printf("Loaded %d sample gates for demonstration\n\n", len(sampleGates))

	// Demonstrate topological sorting (processing order)
	fmt.Println("=== TOPOLOGICAL SORTING (Processing Order) ===")
	netlist.Display(DisplayTopological)

	// Demonstrate alphabetical sorting (lookup/reference)
	fmt.Println("\n=== ALPHABETICAL SORTING (Lookup/Reference) ===")
	netlist.Display(DisplayAlphabetical)

	// Demonstrate fast level access
	fmt.Println("\n=== FAST LEVEL ACCESS ===")
	fmt.Println("Processing gates level by level:")

	levels := netlist.GetTopologicalLevels()
	for _, level := range levels {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("Level %d: %d gates - ", level, len(gates))

		gateNames := make([]string, len(gates))
		for i, gate := range gates {
			gateNames[i] = gate.Name
		}
		fmt.Printf("[%s]\n", strings.Join(gateNames, ", "))
	}

	// Demonstrate fast signal lookup
	fmt.Println("\n=== FAST SIGNAL LOOKUP ===")
	testSignals := []string{"a1", "out4", "ps16", "unknown_signal"}

	for _, signalName := range testSignals {
		gate, found := netlist.FindGateByName(signalName)
		if found {
			fmt.Printf("Signal '%s': %s gate at level %d\n", signalName, gate.Type, gate.Level)
		} else {
			fmt.Printf("Signal '%s': NOT FOUND\n", signalName)
		}
	}

	// Demonstrate topological validation
	fmt.Println("\n=== TOPOLOGICAL VALIDATION ===")
	err := netlist.ValidateTopologicalOrder()
	if err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Topological order is valid")
	}

	// Demonstrate processing order optimization
	fmt.Println("\n=== PROCESSING ORDER OPTIMIZATION ===")
	processingOrder := netlist.GetProcessingOrder()

	fmt.Println("Optimized processing sequence:")
	for levelIdx, levelGates := range processingOrder {
		fmt.Printf("Step %d (Level %d): Process %d gates in parallel\n",
			levelIdx+1, levels[levelIdx], len(levelGates))

		for _, gate := range levelGates {
			fmt.Printf("  - Execute %s gate '%s'\n", gate.Type, gate.Name)
		}
	}

	// Show file export options
	fmt.Println("\n=== EXPORT OPTIONS ===")
	fmt.Println("Available export formats:")
	fmt.Println("1. Topological sort → Optimized for sequential processing")
	fmt.Println("2. Alphabetical sort → Optimized for lookup and reference")

	// Export examples (would create actual files)
	fmt.Println("\nExport commands:")
	fmt.Println("netlist.ExportSorted(\"circuit_topological.compact\", DisplayTopological)")
	fmt.Println("netlist.ExportSorted(\"circuit_alphabetical.compact\", DisplayAlphabetical)")
}

// Demonstrate search efficiency comparison
func DemoSearchEfficiency() {
	fmt.Println("\n=== SEARCH EFFICIENCY DEMONSTRATION ===")

	// Create large netlist for timing comparison
	netlist := NewEnhancedCompactNetlist()

	// Generate many gates to show search efficiency
	for level := 1; level <= 10; level++ {
		for i := 0; i < 100; i++ {
			gate := CompactGate{
				Name:         fmt.Sprintf("gate_%d_%d", level, i),
				Type:         "2AND",
				Level:        level,
				Inputs:       []string{fmt.Sprintf("in1_%d_%d", level, i), fmt.Sprintf("in2_%d_%d", level, i)},
				Outputs:      []string{fmt.Sprintf("out_%d_%d_s", level, i), fmt.Sprintf("out_%d_%d_1", level, i)},
				FaultTargets: []string{fmt.Sprintf("gate_%d_%d", level, i)},
			}
			netlist.Gates = append(netlist.Gates, gate)
		}
	}

	netlist.rebuildIndices()
	fmt.Printf("Created large netlist: %d gates across %d levels\n", len(netlist.Gates), 10)

	// Demonstrate fast level access
	fmt.Println("\n⚡ Fast Level Access (using index):")
	for level := 1; level <= 3; level++ {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("Level %d: Found %d gates instantly\n", level, len(gates))
	}

	// Demonstrate fast signal lookup
	fmt.Println("\n🔍 Fast Signal Lookup (using index):")
	testSignals := []string{"gate_5_42", "gate_1_99", "gate_10_1"}

	for _, signalName := range testSignals {
		gate, found := netlist.FindGateByName(signalName)
		if found {
			fmt.Printf("Signal '%s': Found at level %d (instant lookup)\n", signalName, gate.Level)
		} else {
			fmt.Printf("Signal '%s': Not found\n", signalName)
		}
	}

	fmt.Println("\n✅ Index-based search provides O(1) lookup performance")
	fmt.Println("✅ Level-based processing enables efficient topological execution")
}

// Main demonstration function
func RunSortingDemo() {
	DemoSortingAndSearch()
	DemoSearchEfficiency()

	fmt.Println("\n=== SUMMARY ===")
	fmt.Println("✅ Topological sorting: Optimized for sequential processing")
	fmt.Println("✅ Alphabetical sorting: Optimized for lookup and reference")
	fmt.Println("✅ Fast level access: O(1) retrieval of gates at any level")
	fmt.Println("✅ Fast signal lookup: O(1) finding of any gate by name")
	fmt.Println("✅ Validation: Ensures no forward references in topological order")
	fmt.Println("✅ Export flexibility: Save in optimal format for intended use")

	fmt.Println("\n🎯 Perfect for designer workflow:")
	fmt.Println("   • Create netlist in natural topological order")
	fmt.Println("   • Switch to alphabetical view for debugging/reference")
	fmt.Println("   • Process efficiently using topological order")
	fmt.Println("   • Fast lookup during development and validation")
}
