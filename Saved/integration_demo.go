package main

import (
	"fmt"
	"log"
	"time"
)

// integration_demo.go - Shows complete netlist-driven workflow
// Demonstrates the full pipeline from compact format to BDD processing

// JSONGate represents the JSON format structure for integration
type JSONGate struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Level   int      `json:"level"`
	Inputs  []string `json:"inputs"`
	Outputs []string `json:"outputs"`
}

// convertCompactToJSON converts enhanced compact format to JSON format
func convertCompactToJSON(netlist *EnhancedCompactNetlist) []JSONGate {
	var jsonGates []JSONGate
	for _, gate := range netlist.Gates {
		jsonGate := JSONGate{
			Name:    gate.Name,
			Type:    gate.Type,
			Level:   gate.Level,
			Inputs:  gate.Inputs,
			Outputs: gate.Outputs,
		}
		jsonGates = append(jsonGates, jsonGate)
	}
	return jsonGates
}

func demonstrateIntegratedWorkflow() {
	fmt.Println("=== INTEGRATED NETLIST-DRIVEN WORKFLOW ===\n")

	// Step 1: Load compact format (designer input)
	fmt.Println("📁 STEP 1: Loading Enhanced Compact Netlist")
	netlist := NewEnhancedCompactNetlist()

	// Simulate loading from file (would be actual file in practice)
	loadSampleCompactNetlist(netlist)

	fmt.Printf("✅ Loaded %d gates across %d levels\n",
		len(netlist.Gates), len(netlist.GetTopologicalLevels()))

	// Step 2: Validate topological order
	fmt.Println("\n🔍 STEP 2: Validating Topological Order")
	if err := netlist.ValidateTopologicalOrder(); err != nil {
		log.Fatal("Topological validation failed:", err)
	}
	fmt.Println("✅ Topological order is valid")

	// Step 3: Display dual sorting capabilities
	fmt.Println("\n📋 STEP 3: Demonstrating Dual Sorting")

	fmt.Println("\n--- Topological Order (Processing Mode) ---")
	netlist.SortByTopologicalLevel()
	showFirstFewGates(netlist, 5)

	fmt.Println("\n--- Alphabetical Order (Lookup Mode) ---")
	netlist.SortBySignalName()
	showFirstFewGates(netlist, 5)

	// Step 4: Fast access operations
	fmt.Println("\n⚡ STEP 4: Fast Access Operations")
	demonstrateFastAccess(netlist)

	// Step 5: Convert to JSON format
	fmt.Println("\n🔄 STEP 5: Converting to JSON Format")
	jsonGates := convertCompactToJSON(netlist)
	fmt.Printf("✅ Converted to JSON format: %d gates\n", len(jsonGates))

	// Step 6: Integration with BDD processing
	fmt.Println("\n🧠 STEP 6: BDD Processing Integration")
	demonstrateBDDIntegration(netlist, jsonGates)

	// Step 7: Performance comparison
	fmt.Println("\n📊 STEP 7: Performance Analysis")
	demonstratePerformance(netlist)
}

func loadSampleCompactNetlist(netlist *EnhancedCompactNetlist) {
	// Load a subset of the LARGE circuit for demonstration
	sampleGates := []CompactGate{
		{"ps16", "PI", 1, []string{}, []string{"ps16_s", "ps16_1"}, []string{"ps16"}},
		{"ps8", "PI", 1, []string{}, []string{"ps8_s", "ps8_1"}, []string{"ps8"}},
		{"ps4", "PI", 1, []string{}, []string{"ps4_s", "ps4_1"}, []string{"ps4"}},

		{"nps16", "NOT", 2, []string{"ps16_s", "ps16_1"}, []string{"nps16_s", "nps16_1"}, []string{"nps16"}},
		{"nps8", "NOT", 2, []string{"ps8_s", "ps8_1"}, []string{"nps8_s", "nps8_1"}, []string{"nps8"}},

		{"ls0", "3AND", 3, []string{"nps4_s", "nps2_s", "nps1_s"}, []string{"ls0_s", "ls0_1"}, []string{"ls0"}},
		{"ls1", "3AND", 3, []string{"nps4_s", "nps2_s", "ps1_s"}, []string{"ls1_s", "ls1_1"}, []string{"ls1"}},

		{"s0", "3AND", 4, []string{"nps16_s", "nps8_s", "ls0_s"}, []string{"s0_s", "s0_1"}, []string{"s0"}},
		{"s1", "3AND", 4, []string{"nps16_s", "nps8_s", "ls1_s"}, []string{"s1_s", "s1_1"}, []string{"s1"}},

		{"a1", "2AND", 5, []string{"s1_s", "i0_s"}, []string{"a1_s", "a1_1"}, []string{"a1"}},
		{"a2", "2AND", 5, []string{"s2_s", "i1_s"}, []string{"a2_s", "a2_1"}, []string{"a2"}},

		{"out1", "2OR", 6, []string{"a1_s", "a3_s"}, []string{"out1_s", "out1_1"}, []string{"out1"}},
		{"out2", "2OR", 6, []string{"a2_s", "a4_s"}, []string{"out2_s", "out2_1"}, []string{"out2"}},
	}

	for _, gate := range sampleGates {
		netlist.Gates = append(netlist.Gates, gate)
		// Update indices
		netlist.nameIndex[gate.Name] = len(netlist.Gates) - 1
		if _, exists := netlist.levelIndex[gate.Level]; !exists {
			netlist.levelIndex[gate.Level] = []int{}
		}
		netlist.levelIndex[gate.Level] = append(netlist.levelIndex[gate.Level], len(netlist.Gates)-1)
	}
}

func showFirstFewGates(netlist *EnhancedCompactNetlist, count int) {
	displayed := 0
	for _, gate := range netlist.Gates {
		if displayed >= count {
			break
		}
		fmt.Printf("  %s, %s, %d, %v, %v\n",
			gate.Name, gate.Type, gate.Level, gate.Inputs, gate.Outputs)
		displayed++
	}
	if len(netlist.Gates) > count {
		fmt.Printf("  ... (%d more gates)\n", len(netlist.Gates)-count)
	}
}

func demonstrateFastAccess(netlist *EnhancedCompactNetlist) {
	// Level-based access
	fmt.Println("🎯 Level-based access:")
	level3Gates := netlist.GetGatesAtLevel(3)
	fmt.Printf("  Level 3 has %d gates: ", len(level3Gates))
	for i, gate := range level3Gates {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(gate.Name)
	}
	fmt.Println()

	// Signal lookup
	fmt.Println("\n🔍 Signal lookup:")
	searchSignals := []string{"s0", "a1", "out1", "nonexistent"}
	for _, signalName := range searchSignals {
		if gate, found := netlist.FindGateByName(signalName); found {
			fmt.Printf("  ✅ Found %s: %s gate at level %d\n",
				signalName, gate.Type, gate.Level)
		} else {
			fmt.Printf("  ❌ Signal %s not found\n", signalName)
		}
	}

	// Level enumeration
	fmt.Println("\n📊 Circuit structure:")
	levels := netlist.GetTopologicalLevels()
	for _, level := range levels {
		gateCount := len(netlist.GetGatesAtLevel(level))
		fmt.Printf("  Level %d: %d gates\n", level, gateCount)
	}
}

func demonstrateBDDIntegration(netlist *EnhancedCompactNetlist, jsonGates []JSONGate) {
	fmt.Println("🔗 Integration points:")

	// Show how compact format enables automated processing
	fmt.Println("  1. Compact → JSON conversion: ✅ Complete")
	fmt.Println("  2. Topological ordering: ✅ Validated")
	fmt.Println("  3. BDD signal naming: ✅ Ready")

	// Count different gate types for BDD processing
	gateTypeCounts := make(map[string]int)
	for _, gate := range netlist.Gates {
		gateTypeCounts[gate.Type]++
	}

	fmt.Println("\n📈 Gate type distribution for BDD processing:")
	for gateType, count := range gateTypeCounts {
		fmt.Printf("  %s: %d gates\n", gateType, count)
	}

	// Show signal naming convention compatibility
	fmt.Println("\n🏷️  RUDD BDD signal naming compatibility:")
	for _, gate := range netlist.Gates[:3] { // Show first few
		fmt.Printf("  %s → BDD signals: %s, %s\n",
			gate.Name, gate.Outputs[0], gate.Outputs[1])
	}
}

func demonstratePerformance(netlist *EnhancedCompactNetlist) {
	fmt.Println("⏱️  Performance comparison:")

	// Measure topological sorting
	start := time.Now()
	netlist.SortByTopologicalLevel()
	topologicalTime := time.Since(start)

	// Measure alphabetical sorting
	start = time.Now()
	netlist.SortBySignalName()
	alphabeticalTime := time.Since(start)

	// Measure fast lookups
	searchTargets := []string{"s0", "a1", "out1", "ps16", "ls0"}
	start = time.Now()
	for _, target := range searchTargets {
		netlist.FindGateByName(target)
	}
	lookupTime := time.Since(start)

	// Measure level access
	start = time.Now()
	for level := 1; level <= 6; level++ {
		netlist.GetGatesAtLevel(level)
	}
	levelAccessTime := time.Since(start)

	fmt.Printf("  Topological sort: %v\n", topologicalTime)
	fmt.Printf("  Alphabetical sort: %v\n", alphabeticalTime)
	fmt.Printf("  5 signal lookups: %v (avg: %v per lookup)\n",
		lookupTime, lookupTime/time.Duration(len(searchTargets)))
	fmt.Printf("  6 level accesses: %v (avg: %v per level)\n",
		levelAccessTime, levelAccessTime/6)

	fmt.Println("\n💡 Performance insights:")
	fmt.Println("  • O(1) signal lookup enables fast debugging")
	fmt.Println("  • O(1) level access enables efficient processing")
	fmt.Println("  • Dual indexing provides best of both access patterns")
}

func demonstrateWorkflowScenarios() {
	fmt.Println("\n=== WORKFLOW SCENARIOS ===\n")

	netlist := NewEnhancedCompactNetlist()
	loadSampleCompactNetlist(netlist)

	// Scenario 1: Designer creating circuit
	fmt.Println("🎨 SCENARIO 1: Designer Creating Circuit")
	fmt.Println("Designer workflow:")
	fmt.Println("1. Creates gates in natural topological order")
	fmt.Println("2. Validates topological correctness as they go")
	fmt.Println("3. Exports in topological format for processing")

	netlist.SortByTopologicalLevel()
	fmt.Println("\nTopological view (natural creation order):")
	showFirstFewGates(netlist, 3)

	// Scenario 2: Engineer debugging circuit
	fmt.Println("\n🔧 SCENARIO 2: Engineer Debugging Circuit")
	fmt.Println("Debug workflow:")
	fmt.Println("1. Switches to alphabetical view for fast lookup")
	fmt.Println("2. Finds problematic signals quickly")
	fmt.Println("3. Traces dependencies and outputs")

	netlist.SortBySignalName()
	fmt.Println("\nAlphabetical view (lookup optimized):")
	showFirstFewGates(netlist, 3)

	if gate, found := netlist.FindGateByName("a1"); found {
		fmt.Printf("\nDebug example - Found gate 'a1':")
		fmt.Printf("\n  Type: %s, Level: %d", gate.Type, gate.Level)
		fmt.Printf("\n  Inputs: %v", gate.Inputs)
		fmt.Printf("\n  Outputs: %v", gate.Outputs)
	}

	// Scenario 3: Automated processing system
	fmt.Println("\n\n🤖 SCENARIO 3: Automated Processing System")
	fmt.Println("Processing workflow:")
	fmt.Println("1. Loads compact format from designer")
	fmt.Println("2. Validates topological order")
	fmt.Println("3. Processes level-by-level for ps2ns generation")

	levels := netlist.GetTopologicalLevels()
	fmt.Printf("\nProcessing order (%d levels):\n", len(levels))
	for _, level := range levels {
		gates := netlist.GetGatesAtLevel(level)
		fmt.Printf("  Level %d: %d gates ready for parallel processing\n",
			level, len(gates))
	}
}

func runIntegrationDemo() {
	demonstrateIntegratedWorkflow()
	demonstrateWorkflowScenarios()

	fmt.Println("\n🎉 INTEGRATION COMPLETE!")
	fmt.Println("\nNext steps:")
	fmt.Println("1. ✅ Enhanced compact format with dual sorting")
	fmt.Println("2. ✅ Fast lookup and level access capabilities")
	fmt.Println("3. ✅ Integration with JSON format")
	fmt.Println("4. 🔄 Connect to validation framework")
	fmt.Println("5. 🔄 Generate automated ps2ns functions")
	fmt.Println("6. 🔄 Complete netlist-driven ATPG system")
}
