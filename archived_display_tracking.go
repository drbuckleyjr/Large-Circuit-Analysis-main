// ARCHIVED DISPLAY AND TRACKING FUNCTIONS
// Moved from Backup.go to reduce file size
// These functions implement debugging and tracking features
//
// NOTE: This file is for reference only and cannot be compiled standalone
// as it depends on variables and functions from the main application

package archived

import (
	"fmt"
)

// trackFaultObjects displays fault object movement during search phase
// Shows primary inputs, present-state lines, output lines, and next-state lines using allSAT
func trackFaultObjects(fp int, presentState nd, nextState nd, siSequence []string) {
	fmt.Printf("\n=== FAULT OBJECT TRACKING ===\n")
	fmt.Printf("Present timeframe -> Next timeframe\n")
	fmt.Printf("fp=%d, Present State: %s -> Next State: %s\n", fp, nd2str(presentState), nd2str(nextState))

	// Safety check for null states
	if isNull(presentState) || isNull(nextState) {
		fmt.Printf("Warning: null state detected - skipping fault object tracking\n")
		fmt.Printf("=============================\n\n")
		return
	}

	// Show present-state line inputs (where fault objects arrive)
	fmt.Printf("Present-state lines (fault objects arriving):\n")
	for i := 1; i <= 32; i++ {
		// Check if this line has any S/Is
		stateName := fmt.Sprintf("s%d", i-1)
		lineVar := str2nd(stateName)
		if !isNull(lineVar) {
			combined := and(presentState, lineVar)
			if !isNull(combined) && !isNull(and(presentState, lineVar)) {
				lineObjects := allSAT(and(presentState, lineVar), str2nd)
				if len(lineObjects) > 0 {
					fmt.Printf("  Line %d: %v\n", i, lineObjects)
				}
			}
		}
	}

	// Show next-state line outputs (where fault objects go to)
	fmt.Printf("Next-state lines (fault objects departing):\n")
	for i := 1; i <= 32; i++ {
		stateName := fmt.Sprintf("s%d", i-1)
		lineVar := str2nd(stateName)
		if !isNull(lineVar) {
			combined := and(nextState, lineVar)
			if !isNull(combined) && !isNull(and(nextState, lineVar)) {
				lineObjects := allSAT(and(nextState, lineVar), str2nd)
				if len(lineObjects) > 0 {
					fmt.Printf("  Line %d: %v\n", i, lineObjects)
				}
			}
		}
	}

	// Show S/I sequence being processed
	if len(siSequence) > 0 {
		fmt.Printf("S/I sequence processed: %v\n", siSequence)
	}

	fmt.Printf("=============================\n\n")
}

// debugFaultSets shows the current state of fault elimination
func debugFaultSets() {
	fmt.Printf("\n=== FAULT SET DEBUGGING ===\n")
	fmt.Printf("Original fault set size: %d\n", len(allPossibleFaults))
	fmt.Printf("Current fault set size: %d\n", len(currentFaultSet))
	fmt.Printf("Accumulated S/I sequence: %v\n", accumulatedSIs)

	if len(accumulatedSIs) > 0 {
		fmt.Printf("--- Current vs Standard Methods ---\n")
		fmt.Printf("\n--- Testing Adaptive Method (currentFaultSet) ---\n")
		countIdentical := countSameAsFaultAGlobal(accumulatedSIs)
		fmt.Printf("Adaptive method: %d faults same as fault_A\n", countIdentical)

		fmt.Printf("\n--- Testing Standard Method (allPossibleFaults) ---\n")
		_, identicalToFaultA_standard, _ := categorizeFaultsByName(allPossibleFaults)
		fmt.Printf("Standard method: %d faults same as fault_A\n", len(identicalToFaultA_standard))

		fmt.Printf("\nThis explains the adaptive vs standard difference!\n")
	}
	fmt.Printf("===============================\n\n")
}

// updateFaultSetAfterSelection updates currentFaultSet to contain only faults
// that are identical to fault_A for the current accumulated S/I sequence
func updateFaultSetAfterSelection() {
	if len(currentFaultSet) == 0 || len(accumulatedSIs) == 0 {
		return
	}

	// Use our new function to get the actual fault names
	_, identicalToFaultA, _ := categorizeFaultsByName(currentFaultSet)

	// Update currentFaultSet to only contain the identical faults
	oldCount := len(currentFaultSet)
	currentFaultSet = identicalToFaultA

	fmt.Printf("Updated fault set: %d -> %d faults remaining\n", oldCount, len(currentFaultSet))
}
