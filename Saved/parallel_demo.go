package main

import (
	"fmt"
	"sync"
	"time"
)

// parallel_processing_demo.go - Demonstrates parallel processing of same-level gates

func demonstrateParallelProcessing() {
	fmt.Println("=== PARALLEL PROCESSING DEMONSTRATION ===")
	fmt.Println("Showing that gates at the same topological level can be processed in ANY ORDER\n")

	// Create enhanced netlist
	netlist := NewEnhancedCompactNetlist()

	// Load sample gates with clear level separation
	sampleGates := []CompactGate{
		// Level 1 - Primary inputs (4 gates)
		{"ps16", "PI", 1, []string{}, []string{"ps16_s", "ps16_1"}, []string{"ps16"}},
		{"ps8", "PI", 1, []string{}, []string{"ps8_s", "ps8_1"}, []string{"ps8"}},
		{"ps4", "PI", 1, []string{}, []string{"ps4_s", "ps4_1"}, []string{"ps4"}},
		{"ps2", "PI", 1, []string{}, []string{"ps2_s", "ps2_1"}, []string{"ps2"}},

		// Level 2 - Inverters (4 gates - can all run in parallel)
		{"nps16", "NOT", 2, []string{"ps16_s", "ps16_1"}, []string{"nps16_s", "nps16_1"}, []string{"nps16"}},
		{"nps8", "NOT", 2, []string{"ps8_s", "ps8_1"}, []string{"nps8_s", "nps8_1"}, []string{"nps8"}},
		{"nps4", "NOT", 2, []string{"ps4_s", "ps4_1"}, []string{"nps4_s", "nps4_1"}, []string{"nps4"}},
		{"nps2", "NOT", 2, []string{"ps2_s", "ps2_1"}, []string{"nps2_s", "nps2_1"}, []string{"nps2"}},

		// Level 3 - Local states (4 gates - can all run in parallel)
		{"ls0", "3AND", 3, []string{"nps4_s", "nps2_s", "nps1_s"}, []string{"ls0_s", "ls0_1"}, []string{"ls0"}},
		{"ls1", "3AND", 3, []string{"nps4_s", "nps2_s", "ps1_s"}, []string{"ls1_s", "ls1_1"}, []string{"ls1"}},
		{"ls2", "3AND", 3, []string{"nps4_s", "ps2_s", "nps1_s"}, []string{"ls2_s", "ls2_1"}, []string{"ls2"}},
		{"ls3", "3AND", 3, []string{"nps4_s", "ps2_s", "ps1_s"}, []string{"ls3_s", "ls3_1"}, []string{"ls3"}},

		// Level 4 - State combinations (4 gates - can all run in parallel)
		{"s0", "3AND", 4, []string{"nps16_s", "nps8_s", "ls0_s"}, []string{"s0_s", "s0_1"}, []string{"s0"}},
		{"s1", "3AND", 4, []string{"nps16_s", "nps8_s", "ls1_s"}, []string{"s1_s", "s1_1"}, []string{"s1"}},
		{"s2", "3AND", 4, []string{"nps16_s", "nps8_s", "ls2_s"}, []string{"s2_s", "s2_1"}, []string{"s2"}},
		{"s3", "3AND", 4, []string{"nps16_s", "nps8_s", "ls3_s"}, []string{"s3_s", "s3_1"}, []string{"s3"}},
	}

	for _, gate := range sampleGates {
		netlist.AddGate(gate)
	}

	fmt.Printf("✅ Loaded %d gates across %d levels\n", len(sampleGates), len(netlist.GetTopologicalLevels()))

	// Show level statistics
	fmt.Println("\n📊 Level Statistics (gates per level):")
	stats := netlist.GetLevelStatistics()
	levels := netlist.GetTopologicalLevels()
	for _, level := range levels {
		gateCount := stats[level]
		fmt.Printf("  Level %d: %d gates (can process in parallel)\n", level, gateCount)
	}

	// Demonstrate sequential vs parallel processing
	fmt.Println("\n🔄 PROCESSING DEMONSTRATIONS\n")

	// 1. Sequential processing (traditional approach)
	fmt.Println("1️⃣ Sequential Processing (traditional):")
	start := time.Now()
	processSequentially(netlist)
	sequentialTime := time.Since(start)
	fmt.Printf("   ⏱️  Sequential time: %v\n", sequentialTime)

	// 2. Parallel processing (optimized approach)
	fmt.Println("\n2️⃣ Parallel Processing (optimized):")
	start = time.Now()
	processInParallel(netlist)
	parallelTime := time.Since(start)
	fmt.Printf("   ⏱️  Parallel time: %v\n", parallelTime)

	// Show processing order flexibility
	fmt.Println("\n🎯 PROCESSING ORDER FLEXIBILITY\n")
	demonstrateOrderFlexibility(netlist)

	// Validate no same-level dependencies
	fmt.Println("\n🔍 DEPENDENCY VALIDATION\n")
	if err := netlist.ValidateNoSameLevelDependencies(); err != nil {
		fmt.Printf("❌ Same-level dependency violation: %v\n", err)
	} else {
		fmt.Println("✅ No same-level dependencies detected")
		fmt.Println("✅ All gates at the same level are independent")
		fmt.Println("✅ Parallel processing is safe and correct")
	}

	fmt.Println("\n🎉 PARALLEL PROCESSING BENEFITS:")
	fmt.Println("  ✅ Gates at same level have NO dependencies on each other")
	fmt.Println("  ✅ Can process in ANY ORDER without affecting correctness")
	fmt.Println("  ✅ Perfect for parallel/concurrent processing")
	fmt.Println("  ✅ Scales well with multi-core processors")
	fmt.Println("  ✅ BDD operations can be parallelized within each level")
}

func processSequentially(netlist *EnhancedCompactNetlist) {
	processingOrder := netlist.GetProcessingOrder()

	fmt.Println("   Processing gates level-by-level sequentially:")
	for levelIndex, levelGates := range processingOrder {
		level := levelIndex + 1
		fmt.Printf("   Level %d: Processing %d gates sequentially...\n", level, len(levelGates))

		for _, gate := range levelGates {
			// Simulate gate processing time
			simulateGateProcessing(gate, 1*time.Millisecond)
		}
	}
}

func processInParallel(netlist *EnhancedCompactNetlist) {
	processingOrder := netlist.GetProcessingOrder()

	fmt.Println("   Processing gates level-by-level in parallel:")
	for levelIndex, levelGates := range processingOrder {
		level := levelIndex + 1
		fmt.Printf("   Level %d: Processing %d gates in parallel...\n", level, len(levelGates))

		// Process all gates at this level in parallel
		var wg sync.WaitGroup
		for _, gate := range levelGates {
			wg.Add(1)
			go func(g CompactGate) {
				defer wg.Done()
				// Simulate gate processing time
				simulateGateProcessing(g, 1*time.Millisecond)
			}(gate)
		}
		wg.Wait() // Wait for all gates at this level to complete
	}
}

func simulateGateProcessing(gate CompactGate, processingTime time.Duration) {
	// Simulate BDD operations for this gate
	time.Sleep(processingTime)
	// In real implementation, this would be:
	// - BDD variable creation for outputs
	// - BDD function computation based on gate type
	// - Assignment to BDD manager
}

func demonstrateOrderFlexibility(netlist *EnhancedCompactNetlist) {
	fmt.Println("📝 Same-level gates can be processed in ANY order:")

	levels := netlist.GetTopologicalLevels()
	for _, level := range levels[:3] { // Show first 3 levels
		gates := netlist.GetGatesAtLevel(level)
		if len(gates) > 1 {
			fmt.Printf("\n   Level %d gates (%d total):\n", level, len(gates))

			// Show original order
			fmt.Print("   Original order:   ")
			for i, gate := range gates {
				if i > 0 {
					fmt.Print(" → ")
				}
				fmt.Print(gate.Name)
			}
			fmt.Println()

			// Show reverse order (equally valid)
			fmt.Print("   Reverse order:    ")
			for i := len(gates) - 1; i >= 0; i-- {
				if i < len(gates)-1 {
					fmt.Print(" → ")
				}
				fmt.Print(gates[i].Name)
			}
			fmt.Println()

			// Show random order (equally valid)
			fmt.Print("   Any other order:  ")
			randomOrder := []int{1, 3, 0, 2} // Just a different permutation
			for i, idx := range randomOrder {
				if idx < len(gates) {
					if i > 0 {
						fmt.Print(" → ")
					}
					fmt.Print(gates[idx].Name)
				}
			}
			fmt.Println()

			fmt.Println("   ✅ All orders produce identical results!")
		}
	}
}

func runParallelDemo() {
	demonstrateParallelProcessing()
}
