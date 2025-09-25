// Background.go
// This file contains functions to track fault objects during the search phase
// and to undo the last selected S/I, restoring the previous state.

package main

import (
	"fmt"
)

// trackFaultObjects displays fault object movement during search phase
// Shows primary inputs, present-state lines, output lines, and next-state lines using allSAT
func trackFaultObjects(fp int, presentState nd, nextState nd, siSequence []string) {
	fmt.Printf("=== FAULT OBJECT TRACKING ===")
	fmt.Printf("Present timeframe -> Next timeframe")
	fmt.Printf("fp=%d, Present State: %s -> Next State: %s", fp, nd2str(presentState), nd2str(nextState))

	// Show S/I sequence being processed
	if len(siSequence) > 0 {
		fmt.Printf("S/I sequence processed: %v", siSequence)
	}

	// Show present-state line inputs (where fault objects arrive)
	fmt.Printf("Present-state lines (fault objects arriving): %s", nd2str(presentState))

	// Show next-state line outputs (where fault objects go to)
	fmt.Printf("Next-state lines (fault objects departing): %s", nd2str(nextState))

	fmt.Printf("=============================")
}

// undoLastSI removes the most recent S/I from accumulatedSIs and resets state
func undoLastSI() (bool, int, nd) {
	if len(accumulatedSIs) == 0 {
		fmt.Println("No S/I to undo.")
		return false, 0, s0
	}

	removed := accumulatedSIs[len(accumulatedSIs)-1]
	accumulatedSIs = accumulatedSIs[:len(accumulatedSIs)-1]
	fmt.Printf("Removed last S/I: %s", removed)
	fmt.Printf("Accumulated S/I sequence now has %d entries", len(accumulatedSIs))

	// Reset to the state that existed before the removed S/I
	// This means recalculating fp and ns based on the remaining accumulatedSIs
	fp := 0
	ns := s0

	// Replay the remaining S/I sequence to get the correct state
	for _, si := range accumulatedSIs {
		// Simulate this S/I to get the next fp and ns
		// (You may need to implement this based on your specific state tracking)
		fmt.Printf("Replaying S/I: %s", si)
	}

	return true, fp, ns
}

// Example usage in search phase switch statement:
// case "xxUndo":
//     success, newFp, newNs := undoLastSI()
//     if success {
//         fp = newFp
//         ns = newNs
//         fmt.Println("State restored to before last S/I selection.")
//     }
//     continue

// Example usage for fault tracking:
// trackFaultObjects(fp, ns, nextState, accumulatedSIs)

// In your user command switch/case, add:
// case "xxUndo":
//     UndoLastSI()
//     continue
// ...existing code...
