package main

import (
	"fmt"

	"rudd_Large.go/core"
	"rudd_Large.go/simulation"
	"rudd_Large.go/types"
)

func main() {
	fmt.Println("=== Modular Large Circuit Test ===")

	// Initialize BDD system
	fmt.Println("Initializing BDD system...")
	core.Initialize()

	// Test basic state creation
	fmt.Println("Testing state creation...")
	state0 := core.GetStateByNumber(0)
	fmt.Printf("State 0: %s\n", core.Nd2str(state0))

	state31 := core.GetStateByNumber(31)
	fmt.Printf("State 31: %s\n", core.Nd2str(state31))

	// Test simulation
	fmt.Println("Testing circuit simulation...")
	result := simulation.SimulateWithFault(state0, "no_fault")
	fmt.Printf("Simulation result: SI=%s, Outputs=%s, NextState=%s\n",
		result.SI, result.Outputs, result.NextState)

	// Test S/I accumulation
	fmt.Println("Testing S/I accumulation...")
	testSIs := []string{"s0i0", "s1i1", "s2i2"}
	simulation.AccumulateSIs(testSIs)
	simulation.PrintAccumulatedSIs()

	// Test validation
	fmt.Println("Testing S/I validation...")
	valid := simulation.ValidateSI("s5i3")
	fmt.Printf("s5i3 is valid: %t\n", valid)

	invalid := simulation.ValidateSI("invalid")
	fmt.Printf("'invalid' is valid: %t\n", invalid)

	// Show current global state
	fmt.Printf("Current fault A: %s\n", types.FaultA)
	fmt.Printf("Current fault C: %s\n", types.FaultC)
	fmt.Printf("First simulation flag: %t\n", types.First)

	faultFree, faultA := simulation.GetSimulationResultCount()
	fmt.Printf("Simulation results - Fault-free: %d, Fault-A: %d\n", faultFree, faultA)

	fmt.Println("=== Test Complete ===")
}
