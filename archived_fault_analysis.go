// ARCHIVED FAULT ANALYSIS FUNCTIONS
// Moved from Backup.go to reduce file size
// These functions implement fault counting and categorization features
//
// NOTE: This file is for reference only and cannot be compiled standalone
// as it depends on variables and functions from the main application

package archived

import (
	"fmt"
)

// countSameAsFaultAGlobal returns the number of faults in currentFaultSet that produce
// results identical to fault_A when tested with the given S/I sequence
func countSameAsFaultAGlobal(testSequence []string) int {
	if len(currentFaultSet) == 0 || len(testSequence) == 0 {
		return 0
	}

	// Temporarily save current accumulated S/I sequence
	savedAccumulatedSIs := make([]string, len(accumulatedSIs))
	copy(savedAccumulatedSIs, accumulatedSIs)

	// Set accumulatedSIs to our test sequence
	accumulatedSIs = make([]string, len(testSequence))
	copy(accumulatedSIs, testSequence)

	// Run the fault-free and fault_A simulations for this test sequence
	accumulateFaultFreeSimulations()
	accumulateFaultASimulations(originalFaultA)

	// Use our categorization function to count identical faults
	_, identicalToFaultA, _ := categorizeFaultsByName(currentFaultSet)

	// Restore original accumulated S/I sequence
	accumulatedSIs = savedAccumulatedSIs

	return len(identicalToFaultA)
}

// showAdaptiveFaultSet identifies and displays the faults that behave the same as fault_A
// for the current accumulated S/I sequence (this is what the adaptive method should contain)
func showAdaptiveFaultSet() {
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated yet.")
		return
	}

	fmt.Printf("\n=== ADAPTIVE FAULT SET ANALYSIS ===\n")
	fmt.Printf("Fault_A: %s\n", fault_A)
	fmt.Printf("Accumulated S/I sequence (%d entries): %v\n", len(accumulatedSIs), accumulatedSIs)

	// First accumulate the fault_A simulations for comparison
	accumulateFaultASimulations(fault_A)

	// Test all possible faults to find which ones behave the same as fault_A
	_, identicalToFaultA, _ := categorizeFaultsByName(allPossibleFaults)

	fmt.Printf("\nFaults behaving the same as fault_A: %d faults\n", len(identicalToFaultA))
	if len(identicalToFaultA) > 0 {
		fmt.Printf("Fault list: ")
		for i, fault := range identicalToFaultA {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", fault)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nThis should be the content of the adaptive fault set!\n")
	fmt.Printf("=====================================\n\n")
}

// categorizeFaultsByName - like xx3 but returns fault names instead of simulation results
func categorizeFaultsByName(faultList []string) (identicalToFaultFree, identicalToFaultA, different []string) {
	identicalToFaultFree = []string{}
	identicalToFaultA = []string{}
	different = []string{}

	for _, fault_C := range faultList {
		simResults := []SimResult{}
		for _, si := range accumulatedSIs {
			result := simulateStateInputWithFault(si, fault_C)
			simResults = append(simResults, result)
		}
		if simResultsEqual(simResults, accumulatedFaultFreeSimulations) {
			identicalToFaultFree = append(identicalToFaultFree, fault_C)
		} else if simResultsEqual(simResults, accumulatedFaultASimulations) {
			identicalToFaultA = append(identicalToFaultA, fault_C)
		} else {
			different = append(different, fault_C)
		}
	}
	return
}

// simulateAndCategorizeFaults runs simulation and categorization for analysis
func simulateAndCategorizeFaults(faultList []string, phaseName string) {
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence accumulated.")
		return
	}

	if len(faultList) == 0 {
		fmt.Println("No valid faults provided.")
		return
	}

	fmt.Printf("Testing %d faults with accumulated S/I sequence\n", len(faultList))
	fmt.Printf("S/I sequence (%d entries): %v\n", len(accumulatedSIs), accumulatedSIs)
	fmt.Printf("Primary fault_A: %s\n\n", originalFaultA)

	// Simulate fault-free and fault_A responses for comparison
	var faultFreeResults []SimResult
	var faultAResults []SimResult

	fmt.Println("--- Running fault-free simulation ---")
	// Use centralized accumulation functions instead of inline duplication
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	accumulateFaultFreeSimulations()
	faultFreeResults = make([]SimResult, len(accumulatedFaultFreeSimulations))
	copy(faultFreeResults, accumulatedFaultFreeSimulations)

	fmt.Printf("--- Running fault_A (%s) simulation ---\n", originalFaultA)
	first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
	accumulateFaultASimulations(originalFaultA)
	faultAResults = make([]SimResult, len(accumulatedFaultASimulations))
	copy(faultAResults, accumulatedFaultASimulations)

	fmt.Println("\n--- Categorizing all faults ---")

	// Categorize each fault using three-category system
	sameFaultFree := []string{}
	sameFaultA := []string{}
	different := []string{}

	for _, fault := range faultList {
		var faultResults []SimResult
		// Reset globals before each fault simulation
		first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
		for _, si := range accumulatedSIs {
			faultResults = append(faultResults, simulateStateInputWithFault(si, fault))
		}

		// Compare with fault-free and fault_A (outputs only - black box analysis)
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
			sameFaultFree = append(sameFaultFree, fault)
		} else if matchesFaultA {
			sameFaultA = append(sameFaultA, fault)
		} else {
			different = append(different, fault)
		}
	}

	// Display three-category fault categorization results
	fmt.Printf("\n=== FAULT CATEGORIZATION RESULTS ===\n")
	fmt.Printf("(1) Fault-free behavior (%d faults): %v\n", len(sameFaultFree), sameFaultFree)
	fmt.Printf("(2) Same as fault_A (%s) (%d faults): %v\n", originalFaultA, len(sameFaultA), sameFaultA)
	fmt.Printf("(3) Different behavior (%d faults): %v\n", len(different), different)

	// Add fault-masking effectiveness analysis
	fmt.Printf("Fault-masking effectiveness:\n")
	fmt.Printf("  - Faults eliminated by test sequence: %d\n", len(different))
	fmt.Printf("  - Faults still masked as fault-free: %d\n", len(sameFaultFree))
	fmt.Printf("  - Faults indistinguishable from fault_A: %d\n", len(sameFaultA))
	fmt.Printf("======================================\n\n")
}

// countSameAsFaultA - local function for counting fault interference in S/I selection
// This was embedded in displayAvailableTransitions but moved here for archival
func countSameAsFaultA_archived(candidateSI string, accumulatedSIs []string, fault_A string) int {
	// Create test sequence: accumulated + candidate S/I
	testSequence := make([]string, len(accumulatedSIs))
	copy(testSequence, accumulatedSIs)
	testSequence = append(testSequence, candidateSI)

	// Simulate with fault_A to get reference behavior
	first_ref, ns16h_ref, ns8h_ref, ns4h_ref, ns2h_ref, ns1h_ref := true, false, false, false, false, false
	var faultAResults []SimResult

	for _, si := range testSequence {
		result := simulateStateInputWithFaultLarge(si, fault_A, first_ref, ns16h_ref, ns8h_ref, ns4h_ref, ns2h_ref, ns1h_ref)
		first_ref = false
		ns16h_ref, ns8h_ref, ns4h_ref, ns2h_ref, ns1h_ref = result.ns16, result.ns8, result.ns4, result.ns2, result.ns1
		faultAResults = append(faultAResults, SimResult{
			si: si, outputs: result.outputs, nextState: result.nextState,
		})
	}

	// Count how many other faults produce identical behavior
	sameAsFaultACount := 0
	allPossibleFaults := getAllPossibleFaults()

	for _, testFault := range allPossibleFaults {
		if testFault == fault_A {
			continue // Skip fault_A itself
		}

		// Simulate with test fault
		first_test, ns16h_test, ns8h_test, ns4h_test, ns2h_test, ns1h_test := true, false, false, false, false, false
		identicalBehavior := true

		for i, si := range testSequence {
			result := simulateStateInputWithFaultLarge(si, testFault, first_test, ns16h_test, ns8h_test, ns4h_test, ns2h_test, ns1h_test)
			first_test = false
			ns16h_test, ns8h_test, ns4h_test, ns2h_test, ns1h_test = result.ns16, result.ns8, result.ns4, result.ns2, result.ns1

			// Compare with fault_A result
			if result.outputs != faultAResults[i].outputs || result.nextState != faultAResults[i].nextState {
				identicalBehavior = false
				break
			}
		}

		if identicalBehavior {
			sameAsFaultACount++
		}
	}

	return sameAsFaultACount
}
