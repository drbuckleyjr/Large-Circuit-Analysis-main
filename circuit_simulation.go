package main

// circuit_simulation.go - Circuit simulation functions
// This file contains all simulation-related functions extracted from the monolithic Backup.go

import (
	"fmt"
)

// =================================================================
// SIMULATION FUNCTIONS - Extracted from Backup.go
// =================================================================

// accumulateFaultFreeSimulations: Accumulate fault-free simulation results for each S/I
func accumulateFaultFreeSimulations() {
	accumulatedFaultFreeSimulations = []SimResult{}
	for _, si := range accumulatedSIs {
		result := simulateStateInputWithFault(si, "")
		accumulatedFaultFreeSimulations = append(accumulatedFaultFreeSimulations, result)
	}
}

// accumulateFaultASimulations: Accumulate fault_A simulation results for each S/I
func accumulateFaultASimulations(fault_A string) {
	accumulatedFaultASimulations = []SimResult{}
	for _, si := range accumulatedSIs {
		result := simulateStateInputWithFault(si, fault_A)
		accumulatedFaultASimulations = append(accumulatedFaultASimulations, result)
	}
}

// xx3: For user-input fault_Cs, simulate each S/I and categorize results
func xx3(userFaultCs []string) (identicalToFaultFree, identicalToFaultA, different [][]SimResult) {
	identicalToFaultFree = [][]SimResult{}
	identicalToFaultA = [][]SimResult{}
	different = [][]SimResult{}
	for _, fault_C := range userFaultCs {
		simResults := []SimResult{}
		for _, si := range accumulatedSIs {
			result := simulateStateInputWithFault(si, fault_C)
			simResults = append(simResults, result)
		}
		if simResultsEqual(simResults, accumulatedFaultFreeSimulations) {
			identicalToFaultFree = append(identicalToFaultFree, simResults)
		} else if simResultsEqual(simResults, accumulatedFaultASimulations) {
			identicalToFaultA = append(identicalToFaultA, simResults)
		} else {
			different = append(different, simResults)
		}
	}
	return
}

// xx4: Same as xx3, but with all possible fault_Cs
func xx4(getAllPossibleFaults func() []string) (identicalToFaultFree, identicalToFaultA, different [][]SimResult) {
	return xx3(getAllPossibleFaults())
}

// simResultsEqual compares two slices of SimResult for string equality
func simResultsEqual(a, b []SimResult) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Function to check if fault exists in circuit
func isValidFault(fault string) bool {
	allFaults := getAllPossibleFaults()
	for _, f := range allFaults {
		if f == fault {
			return true
		}
	}
	return false
}

// Simulation phase choice xx1: fault-free simulation
func simulationPhaseXX1() {
	fmt.Println("\n=== SIMULATION PHASE XX1: Fault-Free Circuit ===")
	fmt.Printf("fault_A: %s, fault_C: \"\"\n", originalFaultA)

	// Show current S/I sequence
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated.")
		return
	}

	fmt.Printf("fault_A: %s | Current S/I sequence (%d entries):\n", originalFaultA, len(accumulatedSIs))
	for i, si := range accumulatedSIs {
		fmt.Printf("  %d: %s\n", i+1, si)
	}

	fmt.Println("\nPress ENTER to step through each S/I simulation:")
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	for i, si := range accumulatedSIs {
		fmt.Printf("\nStep %d: %s\n", i+1, si)
		fmt.Scanln() // Wait for ENTER

		result := simulateStateInputWithFault(si, "")
		fmt.Printf("  Outputs: %s\n", result.outputs)
		fmt.Printf("  Next State: %s\n", result.nextState)
	}

	fmt.Println("\nSimulation ended. Returning to switch.")
}

// Simulation phase choice xx2: fault_A simulation
func simulationPhaseXX2() {
	fmt.Println("\n=== SIMULATION PHASE XX2: Faulty Circuit (fault_A) ===")
	fmt.Printf("fault_A: %s, fault_C: %s\n", originalFaultA, originalFaultA)

	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated.")
		return
	}

	// Show current S/I sequence
	fmt.Printf("fault_A: %s | Current S/I sequence (%d entries):\n", originalFaultA, len(accumulatedSIs))
	for i, si := range accumulatedSIs {
		fmt.Printf("  %d: %s\n", i+1, si)
	}

	fmt.Println("\nPress ENTER to step through each S/I simulation:")
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	for i, si := range accumulatedSIs {
		fmt.Printf("\nStep %d: %s\n", i+1, si)
		fmt.Scanln() // Wait for ENTER

		result := simulateStateInputWithFault(si, originalFaultA)
		fmt.Printf("  Outputs: %s\n", result.outputs)
		fmt.Printf("  Next State: %s\n", result.nextState)
	}

	fmt.Println("\nSimulation ended. Returning to switch.")
}

// Helper function for paginated display
func displayPaginated(lines []string, title string) {
	const linesPerPage = 20 // Adjust based on typical screen size
	fmt.Printf("\n%s\n", title)
	fmt.Println("Press SPACE for next page, or any other key to continue...")

	for i := 0; i < len(lines); i += linesPerPage {
		end := i + linesPerPage
		if end > len(lines) {
			end = len(lines)
		}

		// Display current page
		for j := i; j < end; j++ {
			fmt.Println(lines[j])
		}

		// If more pages remain, wait for user input
		if end < len(lines) {
			fmt.Printf("\nPage %d/%d - Press SPACE for next page, any other key to continue: ",
				(i/linesPerPage)+1, (len(lines)+linesPerPage-1)/linesPerPage)
			var input string
			fmt.Scanln(&input)
			if input != " " && input != "" {
				break
			}
		}
	}
}

// Minimal simulation function for xx4 workflow
func simulateAndCategorizeFaults(faultList []string, phaseName string) {
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated.")
		return
	}

	fmt.Printf("Testing %d faults with accumulated S/I sequence\n", len(faultList))
	fmt.Printf("S/I sequence (%d entries): %v\n", len(accumulatedSIs), accumulatedSIs)

	// Simple categorization for xx4
	sameFaultFree := 0
	sameFaultA := 0
	different := 0

	// Run fault-free simulation once
	var faultFreeResults []SimResult
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	for _, si := range accumulatedSIs {
		result := simulateStateInputWithFault(si, "fault_free")
		faultFreeResults = append(faultFreeResults, result)
	}

	// Run fault_A simulation once
	var faultAResults []SimResult
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	for _, si := range accumulatedSIs {
		result := simulateStateInputWithFault(si, originalFaultA)
		faultAResults = append(faultAResults, result)
	}

	// Categorize each fault
	for _, fault := range faultList {
		var faultResults []SimResult
		first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
		for _, si := range accumulatedSIs {
			faultResults = append(faultResults, simulateStateInputWithFault(si, fault))
		}

		// Compare outputs only
		matchesFaultFree := true
		matchesFaultA := true
		for i := range faultResults {
			if faultResults[i].outputs != faultFreeResults[i].outputs {
				matchesFaultFree = false
			}
			if faultResults[i].outputs != faultAResults[i].outputs {
				matchesFaultA = false
			}
		}

		if matchesFaultFree {
			sameFaultFree++
		} else if matchesFaultA {
			sameFaultA++
		} else {
			different++
		}
	}

	// Display summary results
	fmt.Printf("\n=== FAULT CATEGORIZATION RESULTS ===\n")
	fmt.Printf("Fault-free behavior: %d faults\n", sameFaultFree)
	fmt.Printf("Same as fault_A (%s): %d faults\n", originalFaultA, sameFaultA)
	fmt.Printf("Different behavior: %d faults\n", different)
	fmt.Printf("======================================\n")
}

// Simulation phase choice xx4: all possible faults
func simulationPhaseXX4() {
	fmt.Println("\n=== SIMULATION PHASE XX4: All Possible Faults ===")
	fmt.Printf("fault_A: %s\n", originalFaultA)

	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated.")
		return
	}

	// Show current S/I sequence
	fmt.Printf("fault_A: %s | Current S/I sequence (%d entries):\n", originalFaultA, len(accumulatedSIs))
	for i, si := range accumulatedSIs {
		fmt.Printf("  %d: %s\n", i+1, si)
	}

	allFaults := getAllPossibleFaults()

	// Use shared simulation function
	simulateAndCategorizeFaults(allFaults, "XX4")
}

// extractIPart extracts the input part from a user-entered S/I string
// Example: "s12i5" returns 5, "s0i7" returns 7
func extractIPart(si string) int {
	var stateNum, inputNum int
	n, err := fmt.Sscanf(si, "s%di%d", &stateNum, &inputNum)
	if n != 2 || err != nil {
		fmt.Printf("Error parsing S/I '%s': %v\n", si, err)
		return -1 // Return -1 to indicate parsing error
	}
	fmt.Printf("Extracted I-part from '%s': I=%d (S=%d)\n", si, inputNum, stateNum)
	return inputNum
}

// extractSIParts extracts both state and input parts from a user-entered S/I string
// Example: "s12i5" returns (12, 5), "s0i7" returns (0, 7)
func extractSIParts(si string) (int, int) {
	var stateNum, inputNum int
	n, err := fmt.Sscanf(si, "s%di%d", &stateNum, &inputNum)
	if n != 2 || err != nil {
		fmt.Printf("Error parsing S/I '%s': %v\n", si, err)
		return -1, -1 // Return -1, -1 to indicate parsing error
	}
	fmt.Printf("Extracted S/I parts from '%s': S=%d, I=%d\n", si, stateNum, inputNum)
	return stateNum, inputNum
}

// convertInputToNodes converts an input number (0-7) to corresponding rudd.Node values
// for in4_i, in2_i, in1_i based on 3-bit binary encoding
// If bit is 1: use the node directly (in4, in2, in1)
// If bit is 0: use the negation (not(in4), not(in2), not(in1))
func convertInputToNodes(inputNum int) (nd, nd, nd) {
	// Convert input number to 3-bit binary
	in4Bit := (inputNum & 0x04) != 0 // bit 2 (value 4)
	in2Bit := (inputNum & 0x02) != 0 // bit 1 (value 2)
	in1Bit := (inputNum & 0x01) != 0 // bit 0 (value 1)

	// Convert to rudd.Node values
	var in4_i, in2_i, in1_i nd
	if in4Bit {
		in4_i = in4 // bit is 1, use node directly
	} else {
		in4_i = not(in4) // bit is 0, use negation
	}

	if in2Bit {
		in2_i = in2 // bit is 1, use node directly
	} else {
		in2_i = not(in2) // bit is 0, use negation
	}

	if in1Bit {
		in1_i = in1 // bit is 1, use node directly
	} else {
		in1_i = not(in1) // bit is 0, use negation
	}

	return in4_i, in2_i, in1_i
}

// setUP performs setup calculations for simulation
func setUP(usi string, first, ns16b, ns8b, ns4b, ns2b,
	ns1b bool) (bool, bool, bool, bool, bool, bool, bool, bool) {
	// usi is like "s12i5"
	var ps16i, ps8i, ps4i, ps2i, ps1i, in4i, in2i, in1i bool

	// Parse state and input numbers
	var s, i int
	n, err := fmt.Sscanf(usi, "s%di%d", &s, &i)
	if n != 2 || err != nil {
		// fallback: all false if parsing fails
		return false, false, false, false, false, false, false, false
	}

	// State bits (5 bits: s0..s31)
	ps16i = (s & 0x10) != 0
	ps8i = (s & 0x08) != 0
	ps4i = (s & 0x04) != 0
	ps2i = (s & 0x02) != 0
	ps1i = (s & 0x01) != 0

	// Input bits (3 bits: i0..i7)
	in4i = (i & 0x04) != 0
	in2i = (i & 0x02) != 0
	in1i = (i & 0x01) != 0

	// Now handle the "first" logic as before
	var ps16a, ps8a, ps4a, ps2a, ps1a bool
	ps16a = (!first && ns16b) || (first && ps16i)
	ps8a = (!first && ns8b) || (first && ps8i)
	ps4a = (!first && ns4b) || (first && ps4i)
	ps2a = (!first && ns2b) || (first && ps2i)
	ps1a = (!first && ns1b) || (first && ps1i)

	return ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i
}

// simulateSingleTimeframe performs Boolean simulation of one circuit timeframe
// Takes present state and inputs, applies faults, returns outputs and next state
func simulateSingleTimeframe(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i,
	in1i bool, fault_C string) (bool, bool, bool, bool, bool, bool, bool, bool, string) {

	// Helper function to apply faults
	applyFault := func(value bool, faultKey string) bool {
		if fault_C == faultKey+":0" {
			return false
		}
		if fault_C == faultKey+":1" {
			return true
		}
		return value
	}

	// Receives S/Is via present-state lines ps16_i, ps8_i. ect.
	// Uses fault_C to activate local_fault; propagates local_fault
	// to output and next_state lines, out4, out2, out1,
	// ns16, ns8, etc., as S/Is.

	// ----------- Simulation NETLIST ------------
	// -- level 0 --
	ps16b := applyFault(ps16a, "ps16")
	ps8b := applyFault(ps8a, "ps8")
	ps4b := applyFault(ps4a, "ps4")
	ps2b := applyFault(ps2a, "ps2")
	ps1b := applyFault(ps1a, "ps1")
	in4b := applyFault(in4i, "in4")
	in2b := applyFault(in2i, "in2")
	in1b := applyFault(in1i, "in1")
	// -- level 1 --
	nps1b := applyFault(!ps1b, "nps1")
	nps2b := applyFault(!ps2b, "nps2")
	nps4b := applyFault(!ps4b, "nps4")
	nps8b := applyFault(!ps8b, "nps8")
	nps16b := applyFault(!ps16b, "nps16")
	nin1b := applyFault(!in1b, "nin1")
	nin2b := applyFault(!in2b, "nin2")
	nin4b := applyFault(!in4b, "nin4")
	i7b := applyFault(in4b && in2b && in1b, "i7")
	// -- level 2 --
	i0b := applyFault(nin4b && nin2b && nin1b, "i0")
	i1b := applyFault(nin4b && nin2b && in1b, "i1")
	i2b := applyFault(nin4b && in2b && nin1b, "i2")
	i3b := applyFault(nin4b && in2b && in1b, "i3")
	i4b := applyFault(in4b && nin2b && nin1b, "i4")
	i5b := applyFault(in4b && nin2b && in1b, "i5")
	i6b := applyFault(in4b && in2b && nin1b, "i6")
	ls0b := applyFault(nps4b && nps2b && nps1b, "ls0")
	ls1b := applyFault(nps4b && nps2b && ps1b, "ls1")
	ls2b := applyFault(nps4b && ps2b && nps1b, "ls2")
	ls3b := applyFault(nps4b && ps2b && ps1b, "ls3")
	ls4b := applyFault(ps4b && nps2b && nps1b, "ls4")
	ls5b := applyFault(ps4b && nps2b && ps1b, "ls5")
	ls6b := applyFault(ps4b && ps2b && nps1b, "ls6")
	ls7b := applyFault(ps4b && ps2b && ps1b, "ls7")
	ni7b := applyFault(!i7b, "ni7")
	s31b := applyFault(ps16b && ps8b && ls7b, "s31")
	// level 3
	ni0b := applyFault(!i0b, "ni0")
	ni1b := applyFault(!i1b, "ni1")
	ni2b := applyFault(!i2b, "ni2")
	ni3b := applyFault(!i3b, "ni3")
	ni5b := applyFault(!i5b, "ni5")
	ni6b := applyFault(!i6b, "ni6")
	s0b := applyFault(nps16b && nps8b && ls0b, "s0")
	s1b := applyFault(nps16b && nps8b && ls1b, "s1")
	s2b := applyFault(nps16b && nps8b && ls2b, "s2")
	s3b := applyFault(nps16b && nps8b && ls3b, "s3")
	s4b := applyFault(nps16b && nps8b && ls4b, "s4")
	s5b := applyFault(nps16b && nps8b && ls5b, "s5")
	s6b := applyFault(nps16b && nps8b && ls6b, "s6")
	s7b := applyFault(nps16b && nps8b && ls7b, "s7")
	s8b := applyFault(nps16b && ps8b && ls0b, "s8")
	s9b := applyFault(nps16b && ps8b && ls1b, "s9")
	s10b := applyFault(nps16b && ps8b && ls2b, "s10")
	s11b := applyFault(nps16b && ps8b && ls3b, "s11")
	s12b := applyFault(nps16b && ps8b && ls4b, "s12")
	s13b := applyFault(nps16b && ps8b && ls5b, "s13")
	s14b := applyFault(nps16b && ps8b && ls6b, "s14")
	s15b := applyFault(nps16b && ps8b && ls7b, "s15")
	s16b := applyFault(ps16b && nps8b && ls0b, "s16")
	s17b := applyFault(ps16b && nps8b && ls1b, "s17")
	s18b := applyFault(ps16b && nps8b && ls2b, "s18")
	s19b := applyFault(ps16b && nps8b && ls3b, "s19")
	s20b := applyFault(ps16b && nps8b && ls4b, "s20")
	s21b := applyFault(ps16b && nps8b && ls5b, "s21")
	s22b := applyFault(ps16b && nps8b && ls6b, "s22")
	s23b := applyFault(ps16b && nps8b && ls7b, "s23")
	s24b := applyFault(ps16b && ps8b && ls0b, "s24")
	s25b := applyFault(ps16b && ps8b && ls1b, "s25")
	s26b := applyFault(ps16b && ps8b && ls2b, "s26")
	s27b := applyFault(ps16b && ps8b && ls3b, "s27")
	s28b := applyFault(ps16b && ps8b && ls4b, "s28")
	s29b := applyFault(ps16b && ps8b && ls5b, "s29")
	s30b := applyFault(ps16b && ps8b && ls6b, "s30")
	b2b := applyFault(i5b || i3b || i2b, "b2")
	b7b := applyFault(i5b || i1b, "b7")
	c5b := applyFault(i7b || i5b, "c5")
	c13b := applyFault(i3b || i2b, "c13")
	e5b := applyFault(i5b || i4b, "e5")
	e10b := applyFault(i6b || i2b || i0b, "e10")
	e12b := applyFault(i6b || i3b, "e12")
	e18b := applyFault(i6b || i2b, "e18")
	e24b := applyFault(i7b || i1b, "e24")
	// -- level 4 --
	a1b := applyFault(s10b && i0b, "a1")
	a2b := applyFault(s15b && i5b, "a2")
	a3b := applyFault(s18b && ni6b, "a3")
	a4b := applyFault(s20b || s21b || s22b, "a4")
	a5b := applyFault(s24b && ni7b, "a5")
	a6b := applyFault(s25b || s26b, "a6")
	a7b := applyFault(s27b || s28b || s29b, "a7")
	a9b := applyFault(s31b && ni5b && ni2b, "a9")
	b1b := applyFault(s3b && i2b, "b1")
	b3b := applyFault(b2b && s7b, "b3")
	b4b := applyFault(s10b && ni6b, "b4")
	b5b := applyFault(s12b || s13b || s14b, "b5")
	b6b := applyFault(s15b && ni5b, "b6")
	b8b := applyFault(s23b && b7b, "b8")
	b9b := applyFault(s25b || s26b, "b9")
	b10b := applyFault(s27b || s28b || s29b, "b10")
	c2b := applyFault(s7b && ni5b && ni3b, "c2")
	c3b := applyFault(s11b && i7b, "c3")
	c4b := applyFault(s15b && ni5b, "c4")
	c6b := applyFault(s23b && ni5b, "c6")
	c8b := applyFault(s19b && c5b, "c8")
	c14b := applyFault(c13b && s3b, "c14")
	c15b := applyFault(s27b && i7b, "c15")
	d1b := applyFault(s1b && i2b, "d1")
	d2b := applyFault(s3b && ni3b && ni2b, "d2")
	d3b := applyFault(s5b && i0b, "d3")
	d5b := applyFault(s9b && i2b, "d5")
	d6b := applyFault(s11b && ni7b, "d6")
	d7b := applyFault(s13b && i0b, "d7")
	d9b := applyFault(s15b && ni5b, "d9")
	d10b := applyFault(s17b && i2b, "d10")
	d11b := applyFault(s19b && ni7b, "d11")
	d12b := applyFault(s21b && i0b, "d12")
	d13b := applyFault(s23b && ni5b && ni1b, "d13")
	d14b := applyFault(s25b && i2b, "d14")
	d15b := applyFault(s29b && i0b, "d15")
	d27b := applyFault(s27b && ni7b, "d27")
	e1b := applyFault(s0b && i1b, "e1")
	e2b := applyFault(s1b && ni2b, "e2")
	e3b := applyFault(s2b && i2b, "e3")
	e4b := applyFault(s3b && ni3b && ni2b, "e4")
	e6b := applyFault(s5b && ni0b, "e6")
	e7b := applyFault(s6b && i7b, "e7")
	e8b := applyFault(s8b && i1b, "e8")
	e9b := applyFault(s9b && ni2b, "e9")
	e11b := applyFault(s11b && ni7b, "e11")
	e13b := applyFault(s13b && ni0b, "e13")
	e14b := applyFault(s14b && i7b, "e14")
	e15b := applyFault(s15b && ni5b, "e15")
	e16b := applyFault(s16b && i1b, "e16")
	e17b := applyFault(s17b && ni2b, "e17")
	e19b := applyFault(s19b && ni7b, "e19")
	e20b := applyFault(s20b && e12b, "e20")
	e21b := applyFault(s21b && ni0b, "e21")
	e22b := applyFault(s22b && i7b, "e22")
	e23b := applyFault(s23b && ni5b, "e23")
	e25b := applyFault(s25b && ni2b, "e25")
	e26b := applyFault(s26b && i2b, "e26")
	e27b := applyFault(s27b && ni7b, "e27")
	e28b := applyFault(s28b && e12b, "e28")
	e29b := applyFault(s29b && ni0b, "e29")
	e30b := applyFault(s30b && i7b, "e30")
	e31b := applyFault(s4b && e5b, "e31")
	e32b := applyFault(s10b && e10b, "e32")
	e33b := applyFault(s12b && e12b, "e33")
	e34b := applyFault(s18b && e18b, "e34")
	e35b := applyFault(s24b && e24b, "e35")
	f1b := applyFault(s12b && i5b, "f1")
	f2b := applyFault(s27b && i4b, "f2")
	f3b := applyFault(s15b && i0b, "f3")
	f4b := applyFault(s27b && i2b, "f4")
	f5b := applyFault(s0b && i7b, "f5")
	f6b := applyFault(s27b && i1b, "f6")

	// level 5

	a8b := applyFault(a7b || s30b, "a8")
	a10b := applyFault(a1b || a2b || s16b, "a10")
	b11b := applyFault(b10b || s30b, "b11")
	b12b := applyFault(b1b || b3b || s8b, "b12")
	b13b := applyFault(s9b || b4b || s11b, "b13")
	c1b := applyFault(c14b || s4b || s5b, "c1")
	c7b := applyFault(c2b && ni2b, "c7")
	c16b := applyFault(c15b || s28b || s29b, "c16")
	d17b := applyFault(d1b || d2b || d3b, "d17")
	d19b := applyFault(s10b || d6b || d7b, "d19")
	d20b := applyFault(s14b || d9b || d10b, "d20")
	d21b := applyFault(s18b || d11b || d12b, "d21")
	d22b := applyFault(s22b || d13b || d14b, "d22")
	d23b := applyFault(s26b || d15b || s30b, "d23")
	e36b := applyFault(e1b || e2b || e3b, "e36")
	e37b := applyFault(e4b || e31b || e6b, "e37")
	e39b := applyFault(e9b || e32b, "e39")
	e40b := applyFault(e11b || e33b || e13b, "e40")
	e41b := applyFault(e14b || e15b || e16b, "e41")
	e42b := applyFault(e17b || e34b || e19b, "e42")
	e43b := applyFault(e20b || e30b || a9b, "e43")
	e44b := applyFault(e21b || e22b || e23b, "e44")
	e45b := applyFault(e35b || e25b || e26b, "e45")
	e46b := applyFault(e27b || e28b || e29b, "e46")
	out4b := applyFault(f1b || f2b, "out4")
	out2b := applyFault(f3b || f4b, "out2")
	out1b := applyFault(f5b || f6b, "out1")
	// -- level 6 --
	a11b := applyFault(a10b || s17b || a3b, "a11")
	b14b := applyFault(b12b || b13b || b5b, "b14")
	c9b := applyFault(c1b || s6b || c7b, "c9")
	c17b := applyFault(c16b || s30b, "c17")
	d18b := applyFault(s6b || c7b || d5b, "d18")
	d28b := applyFault(d23b || d27b, "d28")
	e38b := applyFault(e7b || c7b || e8b, "e38")
	e47b := applyFault(e36b || e37b || e44b, "e47")
	e49b := applyFault(e39b || e40b || e41b, "e49")
	// -- level 7 --
	a12b := applyFault(a11b || s19b || a4b, "a12")
	b15b := applyFault(b14b || b6b || b8b, "b15")
	c10b := applyFault(c9b || c3b || b5b, "c10")
	d24b := applyFault(d17b || d18b || d19b, "d24")
	e48b := applyFault(e45b || e46b || e38b, "e48")
	// -- level 8 --
	a13b := applyFault(a12b || s23b || a5b, "a13")
	b16b := applyFault(b15b || s24b || b9b, "b16")
	c11b := applyFault(c10b || c4b || c8b, "c11")
	d25b := applyFault(d24b || d20b || d21b, "d25")
	e50b := applyFault(e47b || e48b || e49b, "e50")
	// -- level 9 --
	ns1b := applyFault(e50b || e42b || e43b, "ns1")
	ns8b := applyFault(b16b || b11b || a9b, "ns8")
	a14b := applyFault(a13b || a6b || a8b, "a14")
	c12b := applyFault(c11b || a4b || c6b, "c12")
	d26b := applyFault(d25b || d22b || d28b, "d26")
	// -- level 10 --
	ns2b := applyFault(d26b || s2b || a9b, "ns2")
	ns4b := applyFault(c12b || c17b || a9b, "ns4")
	ns16b := applyFault(a14b || a9b, "ns16")
	// -- END OF CIRCUIT NET LIST --

	// Determine state
	var state string
	switch {
	case !ns16b && !ns8b && !ns4b && !ns2b && !ns1b:
		state = "s0"
	case !ns16b && !ns8b && !ns4b && !ns2b && ns1b:
		state = "s1"
	case !ns16b && !ns8b && !ns4b && ns2b && !ns1b:
		state = "s2"
	case !ns16b && !ns8b && !ns4b && ns2b && ns1b:
		state = "s3"
	case !ns16b && !ns8b && ns4b && !ns2b && !ns1b:
		state = "s4"
	case !ns16b && !ns8b && ns4b && !ns2b && ns1b:
		state = "s5"
	case !ns16b && !ns8b && ns4b && ns2b && !ns1b:
		state = "s6"
	case !ns16b && !ns8b && ns4b && ns2b && ns1b:
		state = "s7"
	case !ns16b && ns8b && !ns4b && !ns2b && !ns1b:
		state = "s8"
	case !ns16b && ns8b && !ns4b && !ns2b && ns1b:
		state = "s9"
	case !ns16b && ns8b && !ns4b && ns2b && !ns1b:
		state = "s10"
	case !ns16b && ns8b && !ns4b && ns2b && ns1b:
		state = "s11"
	case !ns16b && ns8b && ns4b && !ns2b && !ns1b:
		state = "s12"
	case !ns16b && ns8b && ns4b && !ns2b && ns1b:
		state = "s13"
	case !ns16b && ns8b && ns4b && ns2b && !ns1b:
		state = "s14"
	case !ns16b && ns8b && ns4b && ns2b && ns1b:
		state = "s15"
	case ns16b && !ns8b && !ns4b && !ns2b && !ns1b:
		state = "s16"
	case ns16b && !ns8b && !ns4b && !ns2b && ns1b:
		state = "s17"
	case ns16b && !ns8b && !ns4b && ns2b && !ns1b:
		state = "s18"
	case ns16b && !ns8b && !ns4b && ns2b && ns1b:
		state = "s19"
	case ns16b && !ns8b && ns4b && !ns2b && !ns1b:
		state = "s20"
	case ns16b && !ns8b && ns4b && !ns2b && ns1b:
		state = "s21"
	case ns16b && !ns8b && ns4b && ns2b && !ns1b:
		state = "s22"
	case ns16b && !ns8b && ns4b && ns2b && ns1b:
		state = "s23"
	case ns16b && ns8b && !ns4b && !ns2b && !ns1b:
		state = "s24"
	case ns16b && ns8b && !ns4b && !ns2b && ns1b:
		state = "s25"
	case ns16b && ns8b && !ns4b && ns2b && !ns1b:
		state = "s26"
	case ns16b && ns8b && !ns4b && ns2b && ns1b:
		state = "s27"
	case ns16b && ns8b && ns4b && !ns2b && !ns1b:
		state = "s28"
	case ns16b && ns8b && ns4b && !ns2b && ns1b:
		state = "s29"
	case ns16b && ns8b && ns4b && ns2b && !ns1b:
		state = "s30"
	case ns16b && ns8b && ns4b && ns2b && ns1b:
		state = "s31"
	}

	return out4b, out2b, out1b, ns16b, ns8b, ns4b, ns2b, ns1b, state

}

// simulateStateInputWithFault simulates a single S/I combination with optional fault injection
func simulateStateInputWithFault(si string, fault_C string) SimResult {

	// These are global variables: first, ns16h, ns8h, ns4h, ns2h, ns1h

	// Parse S/I string to extract state and input numbers
	var stateNum, inputNum int
	fmt.Sscanf(si, "s%di%d", &stateNum, &inputNum)

	// Perform simulation using setUP and simulateSingleTimeframe
	// setUP := func(usi string, first, ns16b, ns8b, ns4b, ns2b, ns1b bool)
	ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i := setUP(si, first, ns16h, ns8h, ns4h, ns2h, ns1h)
	// simulateSingleTimeframe := func(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i bool,fault_C string)
	// Simulate one timeframe with Boolean logic and fault injection
	out4b, out2b, out1b, ns16b, ns8b, ns4b, ns2b, ns1b, state := simulateSingleTimeframe(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i, fault_C)

	// Establish next state based on current state
	// These are global variables and are being updated globally
	ns16h = ns16b
	ns8h = ns8b
	ns4h = ns4b
	ns2h = ns2b
	ns1h = ns1b
	first = false

	// Format outputs and next state
	// fmt.Println("  fault_A:", originalFaultA, "fault_C:", fault_C) // Commented out to reduce output during bulk simulation
	outputs := fmt.Sprintf("out4:%t out2:%t out1:%t", out4b, out2b, out1b)
	nextState := state // Use the state return value instead of individual bits

	return SimResult{
		si:        si,
		outputs:   outputs,
		nextState: nextState,
	}
}

// SimResultWithState represents simulation results with state information
type SimResultWithState struct {
	si                       string
	outputs                  string
	nextState                string
	ns16, ns8, ns4, ns2, ns1 bool
}

// simulateStateInputWithFaultLarge simulates a single S/I combination with explicit state parameters (for fault interference analysis)
func simulateStateInputWithFaultLarge(si string, fault_C string, first_param, ns16h_param, ns8h_param, ns4h_param, ns2h_param, ns1h_param bool) SimResultWithState {

	// Parse S/I string to extract state and input numbers
	var stateNum, inputNum int
	fmt.Sscanf(si, "s%di%d", &stateNum, &inputNum)

	// Perform simulation using setUP and simulateSingleTimeframe with passed parameters
	ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i := setUP(si, first_param, ns16h_param, ns8h_param, ns4h_param, ns2h_param, ns1h_param)
	// Simulate one timeframe with Boolean logic and fault injection
	out4b, out2b, out1b, ns16b, ns8b, ns4b, ns2b, ns1b, state := simulateSingleTimeframe(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i, fault_C)

	// Format outputs and next state
	outputs := fmt.Sprintf("out4:%t out2:%t out1:%t", out4b, out2b, out1b)
	nextState := state

	return SimResultWithState{
		si:        si,
		outputs:   outputs,
		nextState: nextState,
		ns16:      ns16b,
		ns8:       ns8b,
		ns4:       ns4b,
		ns2:       ns2b,
		ns1:       ns1b,
	}
}
//
// All core simulation functions (setUP, simulateSingleTimeframe, 
// simulateStateInputWithFault, etc.) remain in Backup.go to avoid
// circular dependencies and redeclaration errors.
//
// The main simulation logic continues to reside in Backup.go to maintain
// the monolithic structure until a full refactoring is completed.