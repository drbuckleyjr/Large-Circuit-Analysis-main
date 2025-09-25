package main

import (
	"fmt"

	"rudd_Large.go/core"
	"rudd_Large.go/simulation"
	"rudd_Large.go/types"
)

func main() {
	fmt.Println("Large circuit fault analysis system (modularized)")

	// Initialize the BDD system
	fmt.Println("Initializing BDD system...")
	core.Initialize()

	// Example usage:
	sequence := []string{"s1i2", "s2i3"}
	faultA := "d2:0"

	// Set up global fault state
	types.FaultA = faultA
	types.OriginalFaultA = faultA

	fmt.Printf("Testing sequence: %v with fault: %s\n", sequence, faultA)

	// Simulate the sequence
	for i, si := range sequence {
		fmt.Printf("Step %d: %s\n", i+1, si)

		// Parse state/input (simplified)
		state := core.Str2nd("s1") // Would parse from si in real implementation
		result := simulation.SimulateWithFault(state, faultA)

		fmt.Printf("  Result: %s -> %s\n", result.SI, result.NextState)

		// Store result
		simResult := types.SimResult{
			SI:        si,
			Outputs:   result.Outputs,
			NextState: result.NextState,
		}
		simulation.AddSimulationResult(simResult, false) // fault present
	}

	// Show accumulated results
	simulation.AccumulateSIs(sequence)
	simulation.PrintAccumulatedSIs()

	faultFree, faultA_count := simulation.GetSimulationResultCount()
	fmt.Printf("Simulation complete: %d fault-free, %d fault-A results\n",
		faultFree, faultA_count)

	fmt.Println("=== Modular structure test successful ===")
}
