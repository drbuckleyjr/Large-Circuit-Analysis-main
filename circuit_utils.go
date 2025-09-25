package main

// circuit_utils.go - Utility functions for circuit analysis
// This file contains utility functions used across search and simulation modules

import (
	"fmt"
)

// User input utility function  
func getUserInput() string {
	var input string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&input)
	return input
}

// Display circuit statistics
func displayCircuitStats() {
	fmt.Println("Circuit: 32 states, 8 inputs, 3 outputs")
	fmt.Printf("All possible faults: %d\n", len(allPossibleFaults))
	fmt.Printf("Current fault set: %d faults\n", len(currentFaultSet))
	fmt.Printf("Accumulated S/I: %d entries\n", len(accumulatedSIs))
}

// Reset circuit state for new analysis
func resetCircuitState() {
	first = true
	ns16h = false
	ns8h = false
	ns4h = false
	ns2h = false
	ns1h = false
	fmt.Println("Circuit state reset to initial conditions")
}

// =================================================================
// BDD UTILITY FUNCTIONS - Used by both search and simulation
// =================================================================

// ---- type changing function S/I string -> S/I nd ----
func str2nd(f string) nd {
	mapping := map[string]nd{
		"null": null,
		// s0 with all inputs
		"s0i0": s0i0, "s0i1": s0i1, "s0i2": s0i2, "s0i3": s0i3, "s0i4": s0i4,
		"s0i5": s0i5, "s0i6": s0i6, "s0i7": s0i7,
		// s1 with all inputs
		"s1i0": s1i0, "s1i1": s1i1, "s1i2": s1i2, "s1i3": s1i3, "s1i4": s1i4,
		"s1i5": s1i5, "s1i6": s1i6, "s1i7": s1i7,
		// s2 with all inputs
		"s2i0": s2i0, "s2i1": s2i1, "s2i2": s2i2, "s2i3": s2i3, "s2i4": s2i4,
		"s2i5": s2i5, "s2i6": s2i6, "s2i7": s2i7,
		// s3 with all inputs
		"s3i0": s3i0, "s3i1": s3i1, "s3i2": s3i2, "s3i3": s3i3, "s3i4": s3i4,
		"s3i5": s3i5, "s3i6": s3i6, "s3i7": s3i7,
		// s4 with all inputs
		"s4i0": s4i0, "s4i1": s4i1, "s4i2": s4i2, "s4i3": s4i3, "s4i4": s4i4,
		"s4i5": s4i5, "s4i6": s4i6, "s4i7": s4i7,
		// s5 with all inputs
		"s5i0": s5i0, "s5i1": s5i1, "s5i2": s5i2, "s5i3": s5i3, "s5i4": s5i4,
		"s5i5": s5i5, "s5i6": s5i6, "s5i7": s5i7,
		// s6 with all inputs
		"s6i0": s6i0, "s6i1": s6i1, "s6i2": s6i2, "s6i3": s6i3, "s6i4": s6i4,
		"s6i5": s6i5, "s6i6": s6i6, "s6i7": s6i7,
		// s7 with all inputs
		"s7i0": s7i0, "s7i1": s7i1, "s7i2": s7i2, "s7i3": s7i3, "s7i4": s7i4,
		"s7i5": s7i5, "s7i6": s7i6, "s7i7": s7i7,
		// s8 with all inputs
		"s8i0": s8i0, "s8i1": s8i1, "s8i2": s8i2, "s8i3": s8i3, "s8i4": s8i4,
		"s8i5": s8i5, "s8i6": s8i6, "s8i7": s8i7,
		// s9 with all inputs
		"s9i0": s9i0, "s9i1": s9i1, "s9i2": s9i2, "s9i3": s9i3, "s9i4": s9i4,
		"s9i5": s9i5, "s9i6": s9i6, "s9i7": s9i7,
		// s10 with all inputs
		"s10i0": s10i0, "s10i1": s10i1, "s10i2": s10i2, "s10i3": s10i3,
		"s10i4": s10i4, "s10i5": s10i5, "s10i6": s10i6, "s10i7": s10i7,
		// s11 with all inputs
		"s11i0": s11i0, "s11i1": s11i1, "s11i2": s11i2, "s11i3": s11i3,
		"s11i4": s11i4, "s11i5": s11i5, "s11i6": s11i6, "s11i7": s11i7,
		// s12 with all inputs
		"s12i0": s12i0, "s12i1": s12i1, "s12i2": s12i2, "s12i3": s12i3,
		"s12i4": s12i4, "s12i5": s12i5, "s12i6": s12i6, "s12i7": s12i7,
		// s13 with all inputs
		"s13i0": s13i0, "s13i1": s13i1, "s13i2": s13i2, "s13i3": s13i3,
		"s13i4": s13i4, "s13i5": s13i5, "s13i6": s13i6, "s13i7": s13i7,
		// s14 with all inputs
		"s14i0": s14i0, "s14i1": s14i1, "s14i2": s14i2, "s14i3": s14i3,
		"s14i4": s14i4, "s14i5": s14i5, "s14i6": s14i6, "s14i7": s14i7,
		// s15 with all inputs
		"s15i0": s15i0, "s15i1": s15i1, "s15i2": s15i2, "s15i3": s15i3,
		"s15i4": s15i4, "s15i5": s15i5, "s15i6": s15i6, "s15i7": s15i7,
		// s16 with all inputs
		"s16i0": s16i0, "s16i1": s16i1, "s16i2": s16i2, "s16i3": s16i3,
		"s16i4": s16i4, "s16i5": s16i5, "s16i6": s16i6, "s16i7": s16i7,
		// s17 with all inputs
		"s17i0": s17i0, "s17i1": s17i1, "s17i2": s17i2, "s17i3": s17i3,
		"s17i4": s17i4, "s17i5": s17i5, "s17i6": s17i6, "s17i7": s17i7,
		// s18 with all inputs
		"s18i0": s18i0, "s18i1": s18i1, "s18i2": s18i2, "s18i3": s18i3,
		"s18i4": s18i4, "s18i5": s18i5, "s18i6": s18i6, "s18i7": s18i7,
		// s19 with all inputs
		"s19i0": s19i0, "s19i1": s19i1, "s19i2": s19i2, "s19i3": s19i3,
		"s19i4": s19i4, "s19i5": s19i5, "s19i6": s19i6, "s19i7": s19i7,
		// s20 with all inputs
		"s20i0": s20i0, "s20i1": s20i1, "s20i2": s20i2, "s20i3": s20i3,
		"s20i4": s20i4, "s20i5": s20i5, "s20i6": s20i6, "s20i7": s20i7,
		// s21 with all inputs
		"s21i0": s21i0, "s21i1": s21i1, "s21i2": s21i2, "s21i3": s21i3,
		"s21i4": s21i4, "s21i5": s21i5, "s21i6": s21i6, "s21i7": s21i7,
		// s22 with all inputs
		"s22i0": s22i0, "s22i1": s22i1, "s22i2": s22i2, "s22i3": s22i3,
		"s22i4": s22i4, "s22i5": s22i5, "s22i6": s22i6, "s22i7": s22i7,
		// s23 with all inputs
		"s23i0": s23i0, "s23i1": s23i1, "s23i2": s23i2, "s23i3": s23i3,
		"s23i4": s23i4, "s23i5": s23i5, "s23i6": s23i6, "s23i7": s23i7,
		// s24 with all inputs
		"s24i0": s24i0, "s24i1": s24i1, "s24i2": s24i2, "s24i3": s24i3,
		"s24i4": s24i4, "s24i5": s24i5, "s24i6": s24i6, "s24i7": s24i7,
		// s25 with all inputs
		"s25i0": s25i0, "s25i1": s25i1, "s25i2": s25i2, "s25i3": s25i3,
		"s25i4": s25i4, "s25i5": s25i5, "s25i6": s25i6, "s25i7": s25i7,
		// s26 with all inputs
		"s26i0": s26i0, "s26i1": s26i1, "s26i2": s26i2, "s26i3": s26i3,
		"s26i4": s26i4, "s26i5": s26i5, "s26i6": s26i6, "s26i7": s26i7,
		// s27 with all inputs
		"s27i0": s27i0, "s27i1": s27i1, "s27i2": s27i2, "s27i3": s27i3,
		"s27i4": s27i4, "s27i5": s27i5, "s27i6": s27i6, "s27i7": s27i7,
		// s28 with all inputs
		"s28i0": s28i0, "s28i1": s28i1, "s28i2": s28i2, "s28i3": s28i3,
		"s28i4": s28i4, "s28i5": s28i5, "s28i6": s28i6, "s28i7": s28i7,
		// s29 with all inputs
		"s29i0": s29i0, "s29i1": s29i1, "s29i2": s29i2, "s29i3": s29i3,
		"s29i4": s29i4, "s29i5": s29i5, "s29i6": s29i6, "s29i7": s29i7,
		// s30 with all inputs
		"s30i0": s30i0, "s30i1": s30i1, "s30i2": s30i2, "s30i3": s30i3,
		"s30i4": s30i4, "s30i5": s30i5, "s30i6": s30i6, "s30i7": s30i7,
		// s31 with all inputs
		"s31i0": s31i0, "s31i1": s31i1, "s31i2": s31i2, "s31i3": s31i3,
		"s31i4": s31i4, "s31i5": s31i5, "s31i6": s31i6, "s31i7": s31i7,
	}
	return mapping[f]
}

// ---- type changing function s nd -> s string ----
func nd2str(sy nd) string {
	mapping := map[nd]string{
		null: "null", s0: "s0", s1: "s1", s2: "s2", s3: "s3", s4: "s4",
		s5: "s5", s6: "s6", s7: "s7", s8: "s8", s9: "s9", s10: "s10",
		s11: "s11", s12: "s12", s13: "s13", s14: "s14", s15: "s15",
		s16: "s16", s17: "s17", s18: "s18", s19: "s19", s20: "s20",
		s21: "s21", s22: "s22", s23: "s23", s24: "s24", s25: "s25",
		s26: "s26", s27: "s27", s28: "s28", s29: "s29", s30: "s30",
		s31: "s31",
	}
	return mapping[sy]
}

// allSat receives a BCD object and parses it into constituent S/Is,
// returning a list of zero or more S/Is, each of type string.
func allSAT(f nd, str2nd func(string) nd) []string {
	// Define the states and inputs
	states := 32 // Number of states (s0 to s31)
	inputs := 8  // Number of inputs (i0 to i7)

	// Initialize the result slice
	g := []string{}

	// Iterate through all states and inputs
	for s := 0; s < states; s++ {
		for i := 0; i < inputs; i++ {
			// Construct the state-input string (e.g., "s0i0", "s1i1", etc.)
			si := fmt.Sprintf("s%di%d", s, i)

			// Convert the string to the nd type using str2nd
			ndValue := str2nd(si)

			// Debug: Check for nil values before calling and()
			if f == nil {
				fmt.Printf("ERROR: f is nil at si=%s\n", si)
				continue
			}
			if ndValue == nil {
				fmt.Printf("ERROR: ndValue is nil at si=%s\n", si)
				continue
			}

			// Check if the conjunction of `f` and the current state-input is not null
			if and(f, ndValue) != null {
				g = append(g, si)
			}
		}
	}

	return g // Return the list of S/I's in the nd function
}

// REAL LOGIC GATE RULE FUNCTIONS =================================
// 2- and 3-input AND and OR gates for BDD circuit simulation
// =================================================================

func piRule(s1, i1 nd) (nd, nd) {
	// computes the propagation function
	o_s := s1
	// computes the 1-set
	o_1 := i1
	return o_s, o_1
}

func notRule(s1, i1 nd) (nd, nd) {
	// computes the propagation function
	o_s := s1
	// computes the 1-set
	o_1 := not(i1)
	return o_s, o_1
}

func and2Rule(s1, s2, i1, i2 nd) (nd, nd) {
	// computes the propagation function
	o_s := or(and(not(i1), s1, not(i2), s2),
		and(i2, s1, not(s2)),
		and(i1, s2, not(s1)),
		and(i1, i2, or(s1, s2)))
	// computes the 1-set
	o_1 := and(i1, i2)
	return o_s, o_1
}

func or2Rule(s1, s2, i1, i2 nd) (nd, nd) {
	// computes the propagation function
	o_s := or(and(i1, s1, i2, s2),
		and(not(i2), s1, not(s2)),
		and(not(i1), s2, not(s1)),
		and(not(i1), not(i2), or(s1, s2)))
	// computes the 1-set
	o_1 := or(i1, i2)
	return o_s, o_1
}

func and3Rule(s1, s2, s3, i1, i2, i3 nd) (nd, nd) {
	s, i := and2Rule(s1, s2, i1, i2)
	o_s, o_1 := and2Rule(s, s3, i, i3)
	return o_s, o_1
}

func or3Rule(s1, s2, s3, i1, i2, i3 nd) (nd, nd) {
	s, i := or2Rule(s1, s2, i1, i2)
	o_s, o_1 := or2Rule(s, s3, i, i3)
	return o_s, o_1
}
