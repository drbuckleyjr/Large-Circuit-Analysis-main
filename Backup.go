// Backup.go
// 31 Aug 2025 @ 0740 hrs
// Search/Simulation are separated into distinct phases
// Based on rudd_Large_070925.go with phase separation and enhanced UI
// Feature: suggests S/I selections tending to reduce size of fault_A fault-class

package main

import (
	"fmt"
	"strings"

	"github.com/dalzilio/rudd"
)

type (
	nd = rudd.Node // Binary Decision Diagram node type alias
	g  struct {    // Gate structure for S/I and output mapping
		si, out_4, out_2, out_1, ns nd
	}
	S []g // State collection type

	// SimResult stores the results of a single S/I simulation
	SimResult struct {
		si        string // State/Input combination (e.g., "s3i5")
		outputs   string // Circuit output values
		nextState string // Resulting next state
	}

	// SIMapping maps S/I strings to their corresponding fault patterns and next states
	SIMapping struct {
		fp int // Fault pattern number
		ns nd  // Next state node
	}
)

// Global variables for circuit simulation and phase management
var accumulatedSIs []string                     // User-selected S/I sequence for testing
var accumulatedFaultFreeSimulations []SimResult // Stored fault-free simulation results
var accumulatedFaultASimulations []SimResult    // Stored fault_A simulation results
var originalFaultA string                       // Primary fault being tested
var first bool = true                           // First simulation timeframe flag
var ns16h, ns8h, ns4h, ns2h, ns1h bool          // Next state holding registers
var fault_A, fault_C string                     // Current fault designations

// Adaptive fault elimination variables
var currentFaultSet []string   // Working fault set that gets progressively filtered
var allPossibleFaults []string // Original complete fault set (never modified)

// BDD (Binary Decision Diagram) setup and core operations
// ================================================================
var bdd *rudd.BDD                                         // Main BDD manager instance
var nd128, nd64, nd32, nd16, nd8, nd4, nd2, nd1 rudd.Node // BDD variable nodes
var not func(rudd.Node) rudd.Node                         // BDD negation operation
var and func(...rudd.Node) rudd.Node                      // BDD conjunction (variadic)
var or func(...rudd.Node) rudd.Node                       // BDD disjunction (variadic)
var isNull func(rudd.Node) bool                           // Check if node is NULL/invalid
var null rudd.Node

func init() {
	bdd, _ = rudd.New(8, rudd.Nodesize(10000), rudd.Cachesize(3000))
	nd128 = bdd.Ithvar(7)
	nd64 = bdd.Ithvar(6)
	nd32 = bdd.Ithvar(5)
	nd16 = bdd.Ithvar(4)
	nd8 = bdd.Ithvar(3)
	nd4 = bdd.Ithvar(2)
	nd2 = bdd.Ithvar(1)
	nd1 = bdd.Ithvar(0)
	not = bdd.Not
	and = bdd.And
	or = bdd.Or

	// Helper function to check if node is NULL/invalid
	isNull = func(n rudd.Node) bool {
		return bdd.Equal(n, null)
	}
	null = bdd.False()

	// Initialize primary state and input variables
	ps16 = nd128
	ps8 = nd64
	ps4 = nd32
	ps2 = nd16
	ps1 = nd8
	in4 = nd4
	in2 = nd2
	in1 = nd1

	// Initialize negated variables
	nps16 = not(ps16)
	nps8 = not(ps8)
	nps4 = not(ps4)
	nps2 = not(ps2)
	nps1 = not(ps1)
	nin4 = not(in4)
	nin2 = not(in2)
	nin1 = not(in1)

	// Initialize state variables
	s0 = and(nps16, and(nps8, and(nps4, and(nps2, nps1))))
	s1 = and(nps16, and(nps8, and(nps4, and(nps2, ps1))))
	s2 = and(nps16, and(nps8, and(nps4, and(ps2, nps1))))
	s3 = and(nps16, and(nps8, and(nps4, and(ps2, ps1))))
	s4 = and(nps16, and(nps8, and(ps4, and(nps2, nps1))))
	s5 = and(nps16, and(nps8, and(ps4, and(nps2, ps1))))
	s6 = and(nps16, and(nps8, and(ps4, and(ps2, nps1))))
	s7 = and(nps16, and(nps8, and(ps4, and(ps2, ps1))))
	s8 = and(nps16, and(ps8, and(nps4, and(nps2, nps1))))
	s9 = and(nps16, and(ps8, and(nps4, and(nps2, ps1))))
	s10 = and(nps16, and(ps8, and(nps4, and(ps2, nps1))))
	s11 = and(nps16, and(ps8, and(nps4, and(ps2, ps1))))
	s12 = and(nps16, and(ps8, and(ps4, and(nps2, nps1))))
	s13 = and(nps16, and(ps8, and(ps4, and(nps2, ps1))))
	s14 = and(nps16, and(ps8, and(ps4, and(ps2, nps1))))
	s15 = and(nps16, and(ps8, and(ps4, and(ps2, ps1))))
	s16 = and(ps16, and(nps8, and(nps4, and(nps2, nps1))))
	s17 = and(ps16, and(nps8, and(nps4, and(nps2, ps1))))
	s18 = and(ps16, and(nps8, and(nps4, and(ps2, nps1))))
	s19 = and(ps16, and(nps8, and(nps4, and(ps2, ps1))))
	s20 = and(ps16, and(nps8, and(ps4, and(nps2, nps1))))
	s21 = and(ps16, and(nps8, and(ps4, and(nps2, ps1))))
	s22 = and(ps16, and(nps8, and(ps4, and(ps2, nps1))))
	s23 = and(ps16, and(nps8, and(ps4, and(ps2, ps1))))
	s24 = and(ps16, and(ps8, and(nps4, and(nps2, nps1))))
	s25 = and(ps16, and(ps8, and(nps4, and(nps2, ps1))))
	s26 = and(ps16, and(ps8, and(nps4, and(ps2, nps1))))
	s27 = and(ps16, and(ps8, and(nps4, and(ps2, ps1))))
	s28 = and(ps16, and(ps8, and(ps4, and(nps2, nps1))))
	s29 = and(ps16, and(ps8, and(ps4, and(nps2, ps1))))
	s30 = and(ps16, and(ps8, and(ps4, and(ps2, nps1))))
	s31 = and(ps16, and(ps8, and(ps4, and(ps2, ps1))))

	// Initialize input variables
	i0 = and(nin4, and(nin2, nin1))
	i1 = and(nin4, and(nin2, in1))
	i2 = and(nin4, and(in2, nin1))
	i3 = and(nin4, and(in2, in1))
	i4 = and(in4, and(nin2, nin1))
	i5 = and(in4, and(nin2, in1))
	i6 = and(in4, and(in2, nin1))
	i7 = and(in4, and(in2, in1))

	// Initialize state-input combinations
	s0i0 = and(s0, i0)
	s0i1 = and(s0, i1)
	s0i2 = and(s0, i2)
	s0i3 = and(s0, i3)
	s0i4 = and(s0, i4)
	s0i5 = and(s0, i5)
	s0i6 = and(s0, i6)
	s0i7 = and(s0, i7)
	s1i0 = and(s1, i0)
	s1i1 = and(s1, i1)
	s1i2 = and(s1, i2)
	s1i3 = and(s1, i3)
	s1i4 = and(s1, i4)
	s1i5 = and(s1, i5)
	s1i6 = and(s1, i6)
	s1i7 = and(s1, i7)
	s2i0 = and(s2, i0)
	s2i1 = and(s2, i1)
	s2i2 = and(s2, i2)
	s2i3 = and(s2, i3)
	s2i4 = and(s2, i4)
	s2i5 = and(s2, i5)
	s2i6 = and(s2, i6)
	s2i7 = and(s2, i7)
	s3i0 = and(s3, i0)
	s3i1 = and(s3, i1)
	s3i2 = and(s3, i2)
	s3i3 = and(s3, i3)
	s3i4 = and(s3, i4)
	s3i5 = and(s3, i5)
	s3i6 = and(s3, i6)
	s3i7 = and(s3, i7)
	s4i0 = and(s4, i0)
	s4i1 = and(s4, i1)
	s4i2 = and(s4, i2)
	s4i3 = and(s4, i3)
	s4i4 = and(s4, i4)
	s4i5 = and(s4, i5)
	s4i6 = and(s4, i6)
	s4i7 = and(s4, i7)
	s5i0 = and(s5, i0)
	s5i1 = and(s5, i1)
	s5i2 = and(s5, i2)
	s5i3 = and(s5, i3)
	s5i4 = and(s5, i4)
	s5i5 = and(s5, i5)
	s5i6 = and(s5, i6)
	s5i7 = and(s5, i7)
	s6i0 = and(s6, i0)
	s6i1 = and(s6, i1)
	s6i2 = and(s6, i2)
	s6i3 = and(s6, i3)
	s6i4 = and(s6, i4)
	s6i5 = and(s6, i5)
	s6i6 = and(s6, i6)
	s6i7 = and(s6, i7)
	s7i0 = and(s7, i0)
	s7i1 = and(s7, i1)
	s7i2 = and(s7, i2)
	s7i3 = and(s7, i3)
	s7i4 = and(s7, i4)
	s7i5 = and(s7, i5)
	s7i6 = and(s7, i6)
	s7i7 = and(s7, i7)
	s8i0 = and(s8, i0)
	s8i1 = and(s8, i1)
	s8i2 = and(s8, i2)
	s8i3 = and(s8, i3)
	s8i4 = and(s8, i4)
	s8i5 = and(s8, i5)
	s8i6 = and(s8, i6)
	s8i7 = and(s8, i7)
	s9i0 = and(s9, i0)
	s9i1 = and(s9, i1)
	s9i2 = and(s9, i2)
	s9i3 = and(s9, i3)
	s9i4 = and(s9, i4)
	s9i5 = and(s9, i5)
	s9i6 = and(s9, i6)
	s9i7 = and(s9, i7)
	s10i0 = and(s10, i0)
	s10i1 = and(s10, i1)
	s10i2 = and(s10, i2)
	s10i3 = and(s10, i3)
	s10i4 = and(s10, i4)
	s10i5 = and(s10, i5)
	s10i6 = and(s10, i6)
	s10i7 = and(s10, i7)
	s11i0 = and(s11, i0)
	s11i1 = and(s11, i1)
	s11i2 = and(s11, i2)
	s11i3 = and(s11, i3)
	s11i4 = and(s11, i4)
	s11i5 = and(s11, i5)
	s11i6 = and(s11, i6)
	s11i7 = and(s11, i7)
	s12i0 = and(s12, i0)
	s12i1 = and(s12, i1)
	s12i2 = and(s12, i2)
	s12i3 = and(s12, i3)
	s12i4 = and(s12, i4)
	s12i5 = and(s12, i5)
	s12i6 = and(s12, i6)
	s12i7 = and(s12, i7)
	s13i0 = and(s13, i0)
	s13i1 = and(s13, i1)
	s13i2 = and(s13, i2)
	s13i3 = and(s13, i3)
	s13i4 = and(s13, i4)
	s13i5 = and(s13, i5)
	s13i6 = and(s13, i6)
	s13i7 = and(s13, i7)
	s14i0 = and(s14, i0)
	s14i1 = and(s14, i1)
	s14i2 = and(s14, i2)
	s14i3 = and(s14, i3)
	s14i4 = and(s14, i4)
	s14i5 = and(s14, i5)
	s14i6 = and(s14, i6)
	s14i7 = and(s14, i7)
	s15i0 = and(s15, i0)
	s15i1 = and(s15, i1)
	s15i2 = and(s15, i2)
	s15i3 = and(s15, i3)
	s15i4 = and(s15, i4)
	s15i5 = and(s15, i5)
	s15i6 = and(s15, i6)
	s15i7 = and(s15, i7)
	s16i0 = and(s16, i0)
	s16i1 = and(s16, i1)
	s16i2 = and(s16, i2)
	s16i3 = and(s16, i3)
	s16i4 = and(s16, i4)
	s16i5 = and(s16, i5)
	s16i6 = and(s16, i6)
	s16i7 = and(s16, i7)
	s17i0 = and(s17, i0)
	s17i1 = and(s17, i1)
	s17i2 = and(s17, i2)
	s17i3 = and(s17, i3)
	s17i4 = and(s17, i4)
	s17i5 = and(s17, i5)
	s17i6 = and(s17, i6)
	s17i7 = and(s17, i7)
	s18i0 = and(s18, i0)
	s18i1 = and(s18, i1)
	s18i2 = and(s18, i2)
	s18i3 = and(s18, i3)
	s18i4 = and(s18, i4)
	s18i5 = and(s18, i5)
	s18i6 = and(s18, i6)
	s18i7 = and(s18, i7)
	s19i0 = and(s19, i0)
	s19i1 = and(s19, i1)
	s19i2 = and(s19, i2)
	s19i3 = and(s19, i3)
	s19i4 = and(s19, i4)
	s19i5 = and(s19, i5)
	s19i6 = and(s19, i6)
	s19i7 = and(s19, i7)
	s20i0 = and(s20, i0)
	s20i1 = and(s20, i1)
	s20i2 = and(s20, i2)
	s20i3 = and(s20, i3)
	s20i4 = and(s20, i4)
	s20i5 = and(s20, i5)
	s20i6 = and(s20, i6)
	s20i7 = and(s20, i7)
	s21i0 = and(s21, i0)
	s21i1 = and(s21, i1)
	s21i2 = and(s21, i2)
	s21i3 = and(s21, i3)
	s21i4 = and(s21, i4)
	s21i5 = and(s21, i5)
	s21i6 = and(s21, i6)
	s21i7 = and(s21, i7)
	s22i0 = and(s22, i0)
	s22i1 = and(s22, i1)
	s22i2 = and(s22, i2)
	s22i3 = and(s22, i3)
	s22i4 = and(s22, i4)
	s22i5 = and(s22, i5)
	s22i6 = and(s22, i6)
	s22i7 = and(s22, i7)
	s23i0 = and(s23, i0)
	s23i1 = and(s23, i1)
	s23i2 = and(s23, i2)
	s23i3 = and(s23, i3)
	s23i4 = and(s23, i4)
	s23i5 = and(s23, i5)
	s23i6 = and(s23, i6)
	s23i7 = and(s23, i7)
	s24i0 = and(s24, i0)
	s24i1 = and(s24, i1)
	s24i2 = and(s24, i2)
	s24i3 = and(s24, i3)
	s24i4 = and(s24, i4)
	s24i5 = and(s24, i5)
	s24i6 = and(s24, i6)
	s24i7 = and(s24, i7)
	s25i0 = and(s25, i0)
	s25i1 = and(s25, i1)
	s25i2 = and(s25, i2)
	s25i3 = and(s25, i3)
	s25i4 = and(s25, i4)
	s25i5 = and(s25, i5)
	s25i6 = and(s25, i6)
	s25i7 = and(s25, i7)
	s26i0 = and(s26, i0)
	s26i1 = and(s26, i1)
	s26i2 = and(s26, i2)
	s26i3 = and(s26, i3)
	s26i4 = and(s26, i4)
	s26i5 = and(s26, i5)
	s26i6 = and(s26, i6)
	s26i7 = and(s26, i7)
	s27i0 = and(s27, i0)
	s27i1 = and(s27, i1)
	s27i2 = and(s27, i2)
	s27i3 = and(s27, i3)
	s27i4 = and(s27, i4)
	s27i5 = and(s27, i5)
	s27i6 = and(s27, i6)
	s27i7 = and(s27, i7)
	s28i0 = and(s28, i0)
	s28i1 = and(s28, i1)
	s28i2 = and(s28, i2)
	s28i3 = and(s28, i3)
	s28i4 = and(s28, i4)
	s28i5 = and(s28, i5)
	s28i6 = and(s28, i6)
	s28i7 = and(s28, i7)
	s29i0 = and(s29, i0)
	s29i1 = and(s29, i1)
	s29i2 = and(s29, i2)
	s29i3 = and(s29, i3)
	s29i4 = and(s29, i4)
	s29i5 = and(s29, i5)
	s29i6 = and(s29, i6)
	s29i7 = and(s29, i7)
	s30i0 = and(s30, i0)
	s30i1 = and(s30, i1)
	s30i2 = and(s30, i2)
	s30i3 = and(s30, i3)
	s30i4 = and(s30, i4)
	s30i5 = and(s30, i5)
	s30i6 = and(s30, i6)
	s30i7 = and(s30, i7)
	s31i0 = and(s31, i0)
	s31i1 = and(s31, i1)
	s31i2 = and(s31, i2)
	s31i3 = and(s31, i3)
	s31i4 = and(s31, i4)
	s31i5 = and(s31, i5)
	s31i6 = and(s31, i6)
	s31i7 = and(s31, i7)
}

// Declare all variables without initialization
var ps16, ps8, ps4, ps2, ps1 rudd.Node
var in4, in2, in1 rudd.Node
var nps16, nps8, nps4, nps2, nps1 rudd.Node
var nin4, nin2, nin1 rudd.Node
var s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15 rudd.Node
var s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31 rudd.Node
var i0, i1, i2, i3, i4, i5, i6, i7 rudd.Node
var s0i0, s0i1, s0i2, s0i3, s0i4, s0i5, s0i6, s0i7 rudd.Node
var s1i0, s1i1, s1i2, s1i3, s1i4, s1i5, s1i6, s1i7 rudd.Node
var s2i0, s2i1, s2i2, s2i3, s2i4, s2i5, s2i6, s2i7 rudd.Node
var s3i0, s3i1, s3i2, s3i3, s3i4, s3i5, s3i6, s3i7 rudd.Node
var s4i0, s4i1, s4i2, s4i3, s4i4, s4i5, s4i6, s4i7 rudd.Node
var s5i0, s5i1, s5i2, s5i3, s5i4, s5i5, s5i6, s5i7 rudd.Node
var s6i0, s6i1, s6i2, s6i3, s6i4, s6i5, s6i6, s6i7 rudd.Node
var s7i0, s7i1, s7i2, s7i3, s7i4, s7i5, s7i6, s7i7 rudd.Node
var s8i0, s8i1, s8i2, s8i3, s8i4, s8i5, s8i6, s8i7 rudd.Node
var s9i0, s9i1, s9i2, s9i3, s9i4, s9i5, s9i6, s9i7 rudd.Node
var s10i0, s10i1, s10i2, s10i3, s10i4, s10i5, s10i6, s10i7 rudd.Node
var s11i0, s11i1, s11i2, s11i3, s11i4, s11i5, s11i6, s11i7 rudd.Node
var s12i0, s12i1, s12i2, s12i3, s12i4, s12i5, s12i6, s12i7 rudd.Node
var s13i0, s13i1, s13i2, s13i3, s13i4, s13i5, s13i6, s13i7 rudd.Node
var s14i0, s14i1, s14i2, s14i3, s14i4, s14i5, s14i6, s14i7 rudd.Node
var s15i0, s15i1, s15i2, s15i3, s15i4, s15i5, s15i6, s15i7 rudd.Node
var s16i0, s16i1, s16i2, s16i3, s16i4, s16i5, s16i6, s16i7 rudd.Node
var s17i0, s17i1, s17i2, s17i3, s17i4, s17i5, s17i6, s17i7 rudd.Node
var s18i0, s18i1, s18i2, s18i3, s18i4, s18i5, s18i6, s18i7 rudd.Node
var s19i0, s19i1, s19i2, s19i3, s19i4, s19i5, s19i6, s19i7 rudd.Node
var s20i0, s20i1, s20i2, s20i3, s20i4, s20i5, s20i6, s20i7 rudd.Node
var s21i0, s21i1, s21i2, s21i3, s21i4, s21i5, s21i6, s21i7 rudd.Node
var s22i0, s22i1, s22i2, s22i3, s22i4, s22i5, s22i6, s22i7 rudd.Node
var s23i0, s23i1, s23i2, s23i3, s23i4, s23i5, s23i6, s23i7 rudd.Node
var s24i0, s24i1, s24i2, s24i3, s24i4, s24i5, s24i6, s24i7 rudd.Node
var s25i0, s25i1, s25i2, s25i3, s25i4, s25i5, s25i6, s25i7 rudd.Node
var s26i0, s26i1, s26i2, s26i3, s26i4, s26i5, s26i6, s26i7 rudd.Node
var s27i0, s27i1, s27i2, s27i3, s27i4, s27i5, s27i6, s27i7 rudd.Node
var s28i0, s28i1, s28i2, s28i3, s28i4, s28i5, s28i6, s28i7 rudd.Node
var s29i0, s29i1, s29i2, s29i3, s29i4, s29i5, s29i6, s29i7 rudd.Node
var s30i0, s30i1, s30i2, s30i3, s30i4, s30i5, s30i6, s30i7 rudd.Node
var s31i0, s31i1, s31i2, s31i3, s31i4, s31i5, s31i6, s31i7 rudd.Node

// [REMOVED] accumulate_simulatedSIs - was unused

// Function to reset accumulated S/I sequence
func resetAccumulatedSIs() {
	accumulatedSIs = []string{}
	fmt.Println("S/I sequence reset.")
}

/*
func addToAccumulatedSIsVerbose(si string) {
	accumulatedSIs = append(accumulatedSIs, si)
	fmt.Printf("Added %s to sequence. Total entries: %d\n", si, len(accumulatedSIs))
}
*/
// addToAccumulatedSIs appends a selected S/I value to the accumulatedSIs slice.
func addToAccumulatedSIs(selectedSI string) {
	accumulatedSIs = append(accumulatedSIs, selectedSI)
}

// clearAccumulations resets all accumulation slices (for option 999)
func clearAccumulations() {
	accumulatedSIs = []string{}
	accumulatedFaultFreeSimulations = []SimResult{}
	accumulatedFaultASimulations = []SimResult{}
}

// initializeFaultSets sets up the adaptive fault elimination system
func initializeFaultSets() {
	allPossibleFaults = getAllPossibleFaults()
	currentFaultSet = make([]string, len(allPossibleFaults))
	copy(currentFaultSet, allPossibleFaults)
	fmt.Printf("Initialized fault sets: %d total faults\n", len(allPossibleFaults))
}

// [MOVED TO archived_fault_analysis.go]
// countSameAsFaultAGlobal - fault counting function moved to archive

// debugFaultSets shows the current state of fault elimination
func debugFaultSets() {
	fmt.Printf("\n=== FAULT SET DEBUGGING ===\n")
	fmt.Printf("Original fault set size: %d\n", len(allPossibleFaults))
	fmt.Printf("Current fault set size: %d\n", len(currentFaultSet))
	fmt.Printf("Accumulated S/I sequence: %v\n", accumulatedSIs)

	fmt.Printf("\nRemaining faults in currentFaultSet:\n")
	for i, fault := range currentFaultSet {
		if i < 10 { // Show first 10
			fmt.Printf("  %d: %s\n", i+1, fault)
		}
	}
	if len(currentFaultSet) > 10 {
		fmt.Printf("  ... and %d more\n", len(currentFaultSet)-10)
	}

	// Test both methods and compare
	if len(accumulatedSIs) > 0 {
		// Test adaptive method (using currentFaultSet)
		fmt.Printf("\n--- Testing Adaptive Method (currentFaultSet) ---\n")
		_, identicalToFaultA_adaptive, _ := categorizeFaultsByName(currentFaultSet)
		fmt.Printf("Adaptive method: %d faults same as fault_A\n", len(identicalToFaultA_adaptive))

		// Test standard method (using allPossibleFaults)
		fmt.Printf("\n--- Testing Standard Method (allPossibleFaults) ---\n")
		_, identicalToFaultA_standard, _ := categorizeFaultsByName(allPossibleFaults)
		fmt.Printf("Standard method: %d faults same as fault_A\n", len(identicalToFaultA_standard))

		fmt.Printf("\nThis explains the 4 vs 155 difference!\n")
	}
	fmt.Printf("===============================\n\n")
}

// [MOVED TO archived_fault_analysis.go]
// showAdaptiveFaultSet - adaptive fault set analysis function moved to archive

// Minimal showAdaptiveFaultSet for essential workflow
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
	fmt.Printf("=====================================\n\n")
}

// trackFaultObjects displays fault object movement during search phase
// Shows primary inputs, present-state lines, output lines, and next-state lines using allSAT
// [MOVED TO archived_display_tracking.go]
// trackFaultObjects - fault object tracking function moved to archive

// undoLastSI removes the most recent S/I from accumulatedSIs and resets state
func undoLastSI(siMap map[string]SIMapping) (bool, int, nd) {
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
} // updateFaultSetAfterSelection updates currentFaultSet to contain only faults
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

// [MOVED TO archived_fault_analysis.go]
// categorizeFaultsByName - fault categorization function moved to archive

// Minimal categorizeFaultsByName for essential workflow functions
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

// =================================================================
// SIMULATION FUNCTIONS moved to circuit_simulation.go
// Functions moved: accumulateFaultFreeSimulations, accumulateFaultASimulations, 
// xx3, xx4, simResultsEqual, isValidFault, simulationPhaseXX1, simulationPhaseXX2,
// displayPaginated, simulateAndCategorizeFaults, simulationPhaseXX4, 
// extractIPart, extractSIParts, convertInputToNodes, setUP, 
// simulateSingleTimeframe, simulateStateInputWithFault, simulateStateInputWithFaultLarge
// =================================================================

func getAllPossibleFaults() []string {
	// Return list of all signal names with :0 and :1 faults
	// --- REQUIRES ALL FAULTS FROM activatePropagateFaultA() ---   <<< ======================
	signals := []string{
		//--- level 0 ---
		"ps16", "ps8", "ps4", "ps2", "ps1", "in4", "in2", "in1",
		//--- level 1 ---
		"nps16", "nps8", "nps4", "nps2", "nps1", "nin4", "nin2", "nin1", "i7",
		//--- level 2 ---
		"i0", "i1", "i2", "i3", "i4", "i5", "i6", "ls0", "ls1", "ls2", "ls3",
		"ls4", "ls5", "ls6", "ls7", "ni7", "s31",
		//--- level 3 ---
		"ni0", "ni1", "ni2", "ni3", "ni5", "ni6", "s0", "s1", "s2", "s3",
		"s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11", "s12", "s13",
		"s14", "s15", "s16", "s17", "s18", "s19", "s20", "s21", "s22", "s23",
		"s24", "s25", "s26", "s27", "s28", "s29", "s30", "b2", "b7", "c5",
		"c13", "e5", "e10", "e12", "e18", "e24",
		//--- level 4 ---
		"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a9", "b1", "b3", "b4",
		"b5", "b6", "b8", "b9", "b10", "c2", "c3", "c4", "c6", "c8", "c14",
		"c15", "d1", "d2", "d3", "d5", "d6", "d7", "d9", "d10", "d11", "d12",
		"d13", "d14", "d15", "d27", "e1", "e2", "e3", "e4", "e6", "e7", "e8",
		"e9", "e11", "e13", "e14", "e15", "e16", "e17", "e19", "e20", "e21",
		"e22", "e23", "e25", "e26", "e27", "e28", "e29", "e30", "e31", "e32",
		"e33", "e34", "e35", "f1", "f2", "f3", "f4", "f5", "f6",
		//--- level 5 ---
		"a8", "a10", "b11", "b12", "b13", "c1", "c7", "c16", "d17", "d19",
		"d20", "d21", "d22", "d23", "e36", "e37", "e39", "e40", "e41", "e42",
		"e43", "e44", "e45", "e46", "out4", "out2", "out1",
		//--- level 6 ---
		"a11", "b14", "c9", "c17", "d18", "d28", "e38", "e47", "e49",
		//--- level 7 ---
		"a12", "b15", "c10", "d24", "e48",
		//--- level 8 ---
		"a13", "b16", "c11", "d25", "e50",
		//--- level 9 ---
		"ns1", "ns8", "a14", "c12", "d26",
		//--- level 10 ---
		"ns2", "ns4", "ns16",
	}

	var faults []string
	for _, signal := range signals {
		faults = append(faults, signal+":0")
		faults = append(faults, signal+":1")
	}
	return faults
}

// End of SIMULATION functions ========================================

func main() {

	fmt.Println("==== Starting main loop ====")
	// Modified main loop with separated phases

	// Initialize everything
	resetAccumulatedSIs() // also reset accumulated results
	clearAccumulations()  // reset accumulated results
	initializeFaultSets() // initialize adaptive fault elimination
	fault_A = ""
	fault_C = ""
	originalFaultA = ""
	// reset simResults
	var breakHelper bool = false
	var skipTransitionDisplay bool = false // Flag to skip automatic transition display

	// Main loop for user interaction
	for {
		if breakHelper {
			fmt.Println("At start of loop, breakHelper is true, resetting it.")
			breakHelper = false
			continue // Continue to the next iteration of the loop
		}
		// Get fault_A and fault_C from user
		fmt.Print("\nEnter fault_A (leave empty to exit): ")
		fmt.Scanln(&fault_A)
		// ---------------------------------------------------------------------------
		if fault_A == "" { // I want to retain the sequence and go to 999, xx1, . . .
			// User has completed search and doesn't want to display another transition view

			break
		}
		// -------- I want to go to line 3305 or equivalent -------------------------
		originalFaultA = fault_A

		fmt.Print("Enter fault_C (leave empty for search phase): ")
		fmt.Scanln(&fault_C)

		if fault_C == "" {
			// SEARCH PHASE
			fmt.Println("\n=== SEARCH PHASE ===")
			fmt.Printf("Searching for test sequence for fault_A: %s\n", fault_A)
			// --- Activate fault_A ---
			var fp int = 0                      // Fault position
			var ns nd = s0                      // next-state
			siMap := make(map[string]SIMapping) // Reset S/I mapping

			for {
				if breakHelper {
					break
				}

				// Only show transition display if we're not skipping it
				if !skipTransitionDisplay {
					fmt.Println("fp =", fp, ", ns =", nd2str(ns))
					// --- Display available transitions --- Get S/I mapping for current fault_A and fault_C
					siMap = displayAvailableTransitions(fp, ns, fault_A, fault_C, allSAT, nd2str, siMap)
					// --------------------------------------------------------------

					if siMap == nil {
						breakHelper = true
						fmt.Println("No S/I mapping found. Exiting search phase.")
						break
					}
				}

				// Reset skipPeek flag after first iteration
				skipTransitionDisplay = false

				// Display current accumulated sequence
				if len(accumulatedSIs) > 0 {
					fmt.Printf("\nCurrent S/I sequence (%d entries):\n",
						len(accumulatedSIs))
					for i, si := range accumulatedSIs {
						fmt.Printf("  %d: %s\n", i+1, si)
					}
				}

				// Get S/I selection from user

				fmt.Println("Select an S/I from the displayed sequence")
				fmt.Print("\nEnter selected S/I (example: s12i5) or control (999/xx0/xx1/xx2/xx4/debug/xxAdaptive/xxUndo): ")
				var input string
				fmt.Scanln(&input)

				switch input {
				case "999":
					// Reset and return to search phase
					resetAccumulatedSIs()
					fp = 0
					ns = s0
					fault_A = ""
					fault_C = ""
					fmt.Println("Resetting to search phase.")
					breakHelper = true

				case "xx0":
					// Return to search mode (extend S/I sequence)
					fmt.Println("\n=== #4.5 ===")
					fmt.Println("Returning to search mode to extend S/I sequence.")

					// Show available transitions display for most recent fp, ns, and show fault_A
					fmt.Printf("Current: fp=%d, ns=%s, fault_A=%s\n", fp, nd2str(ns), originalFaultA)
					displayAvailableTransitions(fp, ns, originalFaultA, fault_C, allSAT, nd2str, siMap)

					// Don't set breakHelper - we want to continue in search mode

				case "xx1":
					// Simulation phase: fault-free
					simulationPhaseXX1()
					skipTransitionDisplay = true // Skip automatic transition display after simulation

				case "xx2":
					// Simulation phase: fault_A
					simulationPhaseXX2()
					skipTransitionDisplay = true // Skip automatic transition display after simulation

				case "xx4":
					// Simulation phase: all faults
					simulationPhaseXX4()
					skipTransitionDisplay = true // Skip automatic transition display after simulation

				case "debug":
					// Debug fault sets
					debugFaultSets()
					continue

				case "xxAdaptive":
					// Show faults remaining in adaptive fault set
					showAdaptiveFaultSet()
					skipTransitionDisplay = true // Skip automatic transition display after simulation

				case "xxUndo":
					// Undo the last S/I selection
					success, newFp, newNs := undoLastSI(siMap)
					if success {
						fp = newFp
						ns = newNs
						fmt.Println("State restored to before last S/I selection.")
					}
					continue

				default:
					// Reset and return to search phase || assumes input is S/I
					// addToAccumulatedSI(input)
					// fmt.Println("Resetting to search phase.")

					// Parse S/I input
					if strings.HasPrefix(input, "s") && strings.Contains(input, "i") {
						// Validate S/I format
						var testS, testI int
						if n, _ := fmt.Sscanf(input, "s%di%d", &testS, &testI); n == 2 {
							if testS >= 0 && testS <= 31 && testI >= 0 && testI <= 7 {
								addToAccumulatedSIs(input)

								// Update fault set after user selection (adaptive filtering)
								updateFaultSetAfterSelection()

								// Look up the correct fp and ns from siMap
								if mapping, exists := siMap[input]; exists {
									fp = mapping.fp
									ns = mapping.ns
									fmt.Printf("Retrieved from siMap: %s -> fp=%d, ns=%s, fault_A=%s\n", input, fp, nd2str(ns), originalFaultA)

									// Track fault objects for this timeframe (temporarily disabled due to crash)
									// trackFaultObjects(fp, ns, ns, accumulatedSIs)
									// Now continue to the top of the for loop
									continue
								} else {
									fmt.Printf("Warning: %s not found in siMap\n", input)
								}
							} else {
								fmt.Println("Invalid S/I range. State: 0-31, Input: 0-7")
							}
						} else {
							fmt.Println("Invalid S/I format. Example: s12i5")
						}
					} else {
						fmt.Println("Invalid input. Example S/I format (s12i5) or control codes (999/xx0/xx1/xx2/xx4)")
					}

				} // end of switch input
			}
			breakHelper = true // Set breakHelper to true to continue the loop
			break
		} else {
			// Direct simulation mode (original behavior)
			fmt.Printf("Direct simulation mode: fault_A=%s, fault_C=%s\n", fault_A, fault_C)
			// Implement direct simulation if needed
		}
	}

} // end of main()

// END END END =========================================================
