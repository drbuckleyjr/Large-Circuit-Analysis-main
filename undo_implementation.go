// undo_implementation.go
// Standalone implementation of S/I undo functionality with state restoration

package main

import (
	"fmt"
)

// undoLastSI removes the most recent S/I from accumulatedSIs and restores the search state
// that existed before that S/I was selected. This requires replaying the remaining S/I sequence.
func undoLastSI() (bool, int, nd) {
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I to undo.")
		return false, 0, s0
	}

	// Store the removed S/I for user feedback
	removed := accumulatedSIs[len(accumulatedSIs)-1]
	accumulatedSIs = accumulatedSIs[:len(accumulatedSIs)-1]

	fmt.Printf("Removed last S/I: %s\n", removed)
	fmt.Printf("Accumulated S/I sequence now has %d entries\n", len(accumulatedSIs))

	// Reset to initial state
	fp := 0
	ns := s0

	// Replay the remaining S/I sequence to restore the correct state
	if len(accumulatedSIs) == 0 {
		fmt.Println("Reset to initial state: fp=0, ns=s0")
		return true, fp, ns
	}

	fmt.Printf("Replaying %d S/I entries to restore state:\n", len(accumulatedSIs))
	for i, si := range accumulatedSIs {
		// Look up the S/I in siMap to get its state transition
		if mapping, exists := siMap[si]; exists {
			fp = mapping.fp
			ns = mapping.ns
			fmt.Printf("  %d. %s -> fp=%d, ns=%s\n", i+1, si, fp, nd2str(ns))
		} else {
			fmt.Printf("  %d. %s -> ERROR: not found in siMap\n", i+1, si)
			// Continue with current state if S/I not found
		}
	}

	fmt.Printf("State restoration complete: fp=%d, ns=%s\n", fp, nd2str(ns))
	fmt.Println("You can now continue selecting S/I from this restored state.")

	return true, fp, ns
}

// Helper function to validate undo operation works correctly
func validateUndoOperation() {
	fmt.Println("\n=== Undo Operation Validation ===")
	fmt.Printf("Current S/I sequence: %v\n", accumulatedSIs)

	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I sequence to validate undo against.")
		return
	}

	// Show what would be removed
	toRemove := accumulatedSIs[len(accumulatedSIs)-1]
	remaining := accumulatedSIs[:len(accumulatedSIs)-1]

	fmt.Printf("Would remove: %s\n", toRemove)
	fmt.Printf("Would remain: %v\n", remaining)

	// Calculate what the restored state would be
	testFp := 0
	testNs := s0

	for _, si := range remaining {
		if mapping, exists := siMap[si]; exists {
			testFp = mapping.fp
			testNs = mapping.ns
		}
	}

	fmt.Printf("Restored state would be: fp=%d, ns=%s\n", testFp, nd2str(testNs))
}
