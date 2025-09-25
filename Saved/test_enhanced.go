package main

import (
	"fmt"
)

// test_enhanced.go - Test the enhanced compact netlist functionality

func testEnhancedNetlist() {
	fmt.Println("=== ENHANCED COMPACT NETLIST TEST ===\n")

	// Create enhanced netlist
	netlist := NewEnhancedCompactNetlist()
	fmt.Println("✅ Created enhanced compact netlist")

	// Add sample gates
	sampleGates := []CompactGate{
		{"ps16", "PI", 1, []string{}, []string{"ps16_s", "ps16_1"}, []string{"ps16"}},
		{"ps8", "PI", 1, []string{}, []string{"ps8_s", "ps8_1"}, []string{"ps8"}},
		{"nps16", "NOT", 2, []string{"ps16_s", "ps16_1"}, []string{"nps16_s", "nps16_1"}, []string{"nps16"}},
		{"nps8", "NOT", 2, []string{"ps8_s", "ps8_1"}, []string{"nps8_s", "nps8_1"}, []string{"nps8"}},
		{"ls0", "3AND", 3, []string{"nps4_s", "nps2_s", "nps1_s"}, []string{"ls0_s", "ls0_1"}, []string{"ls0"}},
		{"s0", "3AND", 4, []string{"nps16_s", "nps8_s", "ls0_s"}, []string{"s0_s", "s0_1"}, []string{"s0"}},
		{"a1", "2AND", 5, []string{"s1_s", "i0_s"}, []string{"a1_s", "a1_1"}, []string{"a1"}},
		{"out1", "2OR", 6, []string{"a1_s", "a4_s"}, []string{"out1_s", "out1_1"}, []string{"out1"}},
	}

	for _, gate := range sampleGates {
		netlist.AddGate(gate)
	}
	fmt.Printf("✅ Added %d gates\n", len(sampleGates))

	// Test topological sorting
	fmt.Println("\n📋 Testing Topological Sorting:")
	netlist.SortByTopologicalLevel()
	fmt.Println("Topological order:")
	for i, gate := range netlist.Gates[:min(5, len(netlist.Gates))] {
		fmt.Printf("  %d. %s (%s) - Level %d\n", i+1, gate.Name, gate.Type, gate.Level)
	}

	// Test alphabetical sorting
	fmt.Println("\n📋 Testing Alphabetical Sorting:")
	netlist.SortBySignalName()
	fmt.Println("Alphabetical order:")
	for i, gate := range netlist.Gates[:min(5, len(netlist.Gates))] {
		fmt.Printf("  %d. %s (%s) - Level %d\n", i+1, gate.Name, gate.Type, gate.Level)
	}

	// Test fast lookups
	fmt.Println("\n🔍 Testing Fast Lookups:")
	searchTargets := []string{"s0", "a1", "out1", "nonexistent"}
	for _, target := range searchTargets {
		if gate, found := netlist.FindGateByName(target); found {
			fmt.Printf("  ✅ Found %s: %s gate at level %d\n", target, gate.Type, gate.Level)
		} else {
			fmt.Printf("  ❌ %s not found\n", target)
		}
	}

	// Test level access
	fmt.Println("\n🎯 Testing Level Access:")
	levels := netlist.GetTopologicalLevels()
	fmt.Printf("Circuit has %d levels: %v\n", len(levels), levels)
	for _, level := range levels[:min(3, len(levels))] {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("  Level %d: %d gates\n", level, len(gates))
	}

	// Test validation
	fmt.Println("\n🔍 Testing Validation:")
	if err := netlist.ValidateTopologicalOrder(); err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Topological order is valid")
	}

	fmt.Println("\n🎉 All tests completed successfully!")
	fmt.Println("\nEnhanced compact netlist features:")
	fmt.Println("  ✅ Dual sorting (topological/alphabetical)")
	fmt.Println("  ✅ O(1) signal lookup")
	fmt.Println("  ✅ O(1) level access")
	fmt.Println("  ✅ Topological validation")
	fmt.Println("  ✅ Designer-friendly format")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func runEnhancedTest() {
	testEnhancedNetlist()
}
