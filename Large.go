// Large.go
// 22 Aug 2025 @ 1820 hrs
// Search/Simulation are separated into distinct phases
// Based on rudd_Large_070925.go with phase separation and enhanced UI
// Feature: suggests S/I selections tending to reduce size of fault_A fault-class

package main

import (
	"bufio"
	"fmt"
	"os"
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

// accumulate_simulatedSIs aggregates or processes simulated SI values as needed.
func accumulate_simulatedSIs(simulatedSIs []float64) float64 {
	var total float64
	for _, si := range simulatedSIs {
		total += si
	}
	return total
}

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

// Helper: Count members of fault_A fault-class for a Driving S/I sequence
func countFaultAClassForSI(drivingSI []string, faultA string) int {
    // Simulate fault_A for reference
    refResults := make([]SimResult, len(drivingSI))
    for i, si := range drivingSI {
        refResults[i] = simulateStateInputWithFaultLargeStateless(si, faultA)
    }

    count := 0
    for _, fault := range allPossibleFaults {
        results := make([]SimResult, len(drivingSI))
        for i, si := range drivingSI {
            results[i] = simulateStateInputWithFaultLargeStateless(si, fault)
        }
        if simResultsEqual(results, refResults) {
            count++
        }
    }
    return count
}

// Stateless simulation for a fault and S/I (no global state)
func simulateStateInputWithFaultLargeStateless(si string, fault string) SimResult {
    // Wrapper for stateless simulation logic
    return simulateStateInputWithFaultLarge(si, fault, true, false, false, false, false, false)
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

// Simulation phase choice xx3: user-provided fault list
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

// simulateAndCategorizeFaults performs fault simulation and categorization for a given fault list
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

	if len(sameFaultA) > 0 {
		fmt.Printf("⚠️  OPTIMIZATION OPPORTUNITY: Consider extending S/I sequence to reduce category (2)\n")
	}

	// For XX3, offer detailed results with pagination
	if phaseName == "XX3" {
		fmt.Print("\nShow detailed simulation results? (y/n): ")
		var showDetails string
		fmt.Scanln(&showDetails)

		if strings.ToLower(showDetails) == "y" {
			showDetailedResultsPaginated(accumulatedSIs, faultFreeResults, faultAResults, different)
		}
		fmt.Println("\nSimulation ended. Returning to switch.")
		return
	}

	fmt.Println("\nSimulation ended. Returning to switch.")
}

// Helper function to show detailed results with pagination
func showDetailedResultsPaginated(accumulatedSIs []string, faultFreeResults []SimResult, faultAResults []SimResult, different []string) {
	fmt.Printf("\nDetailed Results for %d S/I inputs:\n", len(accumulatedSIs))

	// Prepare fault-free baseline lines
	var faultFreeLines []string
	for i, si := range accumulatedSIs {
		result := faultFreeResults[i]
		faultFreeLines = append(faultFreeLines, fmt.Sprintf("  %s -> Out:%s NS:%s", si, result.outputs, result.nextState))
	}
	displayPaginated(faultFreeLines, "Fault-Free Reference:")

	// Prepare fault_A baseline lines
	var faultALines []string
	for i, si := range accumulatedSIs {
		result := faultAResults[i]
		faultALines = append(faultALines, fmt.Sprintf("  %s -> Out:%s NS:%s", si, result.outputs, result.nextState))
	}
	displayPaginated(faultALines, fmt.Sprintf("Fault_A (%s) Reference:", originalFaultA))

	// Show different faults in detail with pagination
	if len(different) > 0 {
		var differentLines []string
		for _, fault := range different {
			differentLines = append(differentLines, fmt.Sprintf("\nFault: %s", fault))
			// Re-simulate to show results
			first, ns16h, ns8h, ns4h, ns2h, ns1h = true, false, false, false, false, false
			for _, si := range accumulatedSIs {
				result := simulateStateInputWithFault(si, fault)
				differentLines = append(differentLines, fmt.Sprintf("  %s -> Out:%s NS:%s", si, result.outputs, result.nextState))
			}
		}
		displayPaginated(differentLines, "Faults with Different Responses:")
	}
}

func simulationPhaseXX3() {
	fmt.Println("\n=== SIMULATION PHASE XX3: User-Provided Fault List ===")
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

	// Get fault list from user
	fmt.Println("\nEnter fault list (format: signal:0 or signal:1, separated by spaces):")
	fmt.Print("Faults: ")

	// Use bufio to read the entire line including spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	faultListStr := scanner.Text()

	faultList := strings.Fields(faultListStr)
	fmt.Println("Received faults:", faultList)

	// Validate faults
	var validFaults []string
	for _, fault := range faultList {
		if isValidFault(fault) {
			validFaults = append(validFaults, fault)
		} else {
			fmt.Printf("Warning: Invalid fault %s ignored.\n", fault)
		}
	}

	// Use shared simulation function with XX3-specific behavior
	simulateAndCategorizeFaults(validFaults, "XX3")
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

// =================================================================
// Beginning of SEARCH functions
// =================================================================

// activatePropagateFaultA simulates the circuit from present state to next state
// Receives S/I inputs via present-state lines and propagates through circuit logic
// with optional fault injection to produce outputs and next-state values
func activatePropagateFaultA(ps16_i, ps8_i, ps4_i, ps2_i, ps1_i nd,
	fault_A string) (nd, nd, nd, nd, nd, nd, nd, nd) {

	// Receives S/I inputs via present-state lines ps16_i, ps8_i, etc.
	// Applies fault_A to activate local faults and propagates through circuit
	// to output and next-state lines: out4_s, out2_s, out1_s, ns16_s, ns8_s, etc.

	// Helper function to apply faults to signals
	applyFault := func(signal nd, faultSignal nd,
		fault string, faultName string) nd {
		if fault == faultName+":0" {
			return faultSignal
		}
		if fault == faultName+":1" {
			return not(faultSignal)
		}
		return signal
	}

	// A circuit-level-sorted sequence of gates
	// Such that each signal has already been defined when used

	// ---- Fault Propagation NETLIST for LARGE Circuit ----

	// level 1

	ps16_s, ps16_1 := piRule(ps16_i, ps16)
	ps16_s = applyFault(ps16_s, ps16_1, fault_A, "ps16")

	ps8_s, ps8_1 := piRule(ps8_i, ps8)
	ps8_s = applyFault(ps8_s, ps8_1, fault_A, "ps8")

	ps4_s, ps4_1 := piRule(ps4_i, ps4)
	ps4_s = applyFault(ps4_s, ps4_1, fault_A, "ps4")

	ps2_s, ps2_1 := piRule(ps2_i, ps2)
	ps2_s = applyFault(ps2_s, ps2_1, fault_A, "ps2")

	ps1_s, ps1_1 := piRule(ps1_i, ps1)
	ps1_s = applyFault(ps1_s, ps1_1, fault_A, "ps1")

	in4_s, in4_1 := piRule(null, in4)
	in4_s = applyFault(in4_s, in4_1, fault_A, "in4")

	in2_s, in2_1 := piRule(null, in2)
	in2_s = applyFault(in2_s, in2_1, fault_A, "in2")

	in1_s, in1_1 := piRule(null, in1)
	in1_s = applyFault(in1_s, in1_1, fault_A, "in1")

	nps1_s, nps1_1 := notRule(ps1_s, ps1_1)
	nps1_s = applyFault(nps1_s, nps1_1, fault_A, "nps1")

	nps2_s, nps2_1 := notRule(ps2_s, ps2_1)
	nps2_s = applyFault(nps2_s, nps2_1, fault_A, "nps2")

	nps4_s, nps4_1 := notRule(ps4_s, ps4_1)
	nps4_s = applyFault(nps4_s, nps4_1, fault_A, "nps4")

	nps8_s, nps8_1 := notRule(ps8_s, ps8_1)
	nps8_s = applyFault(nps8_s, nps8_1, fault_A, "nps8")

	nps16_s, nps16_1 := notRule(ps16_s, ps16_1)
	nps16_s = applyFault(nps16_s, nps16_1, fault_A, "nps16")

	nin1_s, nin1_1 := notRule(in1_s, in1_1)
	nin1_s = applyFault(nin1_s, nin1_1, fault_A, "nin1")

	nin2_s, nin2_1 := notRule(in2_s, in2_1)
	nin2_s = applyFault(nin2_s, nin2_1, fault_A, "nin2")

	nin4_s, nin4_1 := notRule(in4_s, in4_1)
	nin4_s = applyFault(nin4_s, nin4_1, fault_A, "nin4")

	i7_s, i7_1 := and3Rule(in4_s, in2_s, in1_s, in4_1, in2_1, in1_1)
	i7_s = applyFault(i7_s, i7_1, fault_A, "i7")

	// level 2

	i0_s, i0_1 := and3Rule(nin4_s, nin2_s, nin1_s, nin4_1, nin2_1, nin1_1)
	i0_s = applyFault(i0_s, i0_1, fault_A, "i0")

	i1_s, i1_1 := and3Rule(nin4_s, nin2_s, in1_s, nin4_1, nin2_1, in1_1)
	i1_s = applyFault(i1_s, i1_1, fault_A, "i1")

	i2_s, i2_1 := and3Rule(nin4_s, in2_s, nin1_s, nin4_1, in2_1, nin1_1)
	i2_s = applyFault(i2_s, i2_1, fault_A, "i2")

	i3_s, i3_1 := and3Rule(nin4_s, in2_s, in1_s, nin4_1, in2_1, in1_1)
	i3_s = applyFault(i3_s, i3_1, fault_A, "i3")

	i4_s, i4_1 := and3Rule(in4_s, nin2_s, nin1_s, in4_1, nin2_1, nin1_1)
	i4_s = applyFault(i4_s, i4_1, fault_A, "i4")

	i5_s, i5_1 := and3Rule(in4_s, nin2_s, in1_s, in4_1, nin2_1, in1_1)
	i5_s = applyFault(i5_s, i5_1, fault_A, "i5")

	i6_s, i6_1 := and3Rule(in4_s, in2_s, nin1_s, in4_1, in2_1, nin1_1)
	i6_s = applyFault(i6_s, i6_1, fault_A, "i6")

	ls0_s, ls0_1 := and3Rule(nps4_s, nps2_s, nps1_s, nps4_1, nps2_1, nps1_1)
	ls0_s = applyFault(ls0_s, ls0_1, fault_A, "ls0")

	ls1_s, ls1_1 := and3Rule(nps4_s, nps2_s, ps1_s, nps4_1, nps2_1, ps1_1)
	ls1_s = applyFault(ls1_s, ls1_1, fault_A, "ls1")

	ls2_s, ls2_1 := and3Rule(nps4_s, ps2_s, nps1_s, nps4_1, ps2_1, nps1_1)
	ls2_s = applyFault(ls2_s, ls2_1, fault_A, "ls2")

	ls3_s, ls3_1 := and3Rule(nps4_s, ps2_s, ps1_s, nps4_1, ps2_1, ps1_1)
	ls3_s = applyFault(ls3_s, ls3_1, fault_A, "ls3")

	ls4_s, ls4_1 := and3Rule(ps4_s, nps2_s, nps1_s, ps4_1, nps2_1, nps1_1)
	ls4_s = applyFault(ls4_s, ls4_1, fault_A, "ls4")

	ls5_s, ls5_1 := and3Rule(ps4_s, nps2_s, ps1_s, ps4_1, nps2_1, ps1_1)
	ls5_s = applyFault(ls5_s, ls5_1, fault_A, "ls5")

	ls6_s, ls6_1 := and3Rule(ps4_s, ps2_s, nps1_s, ps4_1, ps2_1, nps1_1)
	ls6_s = applyFault(ls6_s, ls6_1, fault_A, "ls6")

	ls7_s, ls7_1 := and3Rule(ps4_s, ps2_s, ps1_s, ps4_1, ps2_1, ps1_1)
	ls7_s = applyFault(ls7_s, ls7_1, fault_A, "ls7")

	ni7_s, ni7_1 := notRule(i7_s, i7_1)
	ni7_s = applyFault(ni7_s, ni7_1, fault_A, "ni7")

	s31_s, s31_1 := and3Rule(ps16_s, ps8_s, ls7_s, ps16_1, ps8_1, ls7_1)
	s31_s = applyFault(s31_s, s31_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s, s16_1, fault_A, "0")

	// level 3

	ni0_s, ni0_1 := notRule(i0_s, i0_1)
	ni0_s = applyFault(ni0_s, ni0_1, fault_A, "ni0")

	ni1_s, ni1_1 := notRule(i1_s, i1_1)
	ni1_s = applyFault(ni1_s, ni1_1, fault_A, "ni1")

	ni2_s, ni2_1 := notRule(i2_s, i2_1)
	ni2_s = applyFault(ni2_s, ni2_1, fault_A, "ni2")

	ni3_s, ni3_1 := notRule(i3_s, i3_1)
	ni3_s = applyFault(ni3_s, ni3_1, fault_A, "ni3")

	ni5_s, ni5_1 := notRule(i5_s, i5_1)
	ni5_s = applyFault(ni5_s, ni5_1, fault_A, "ni5")

	ni6_s, ni6_1 := notRule(i6_s, i6_1)
	ni6_s = applyFault(ni6_s, ni6_1, fault_A, "ni6")

	s0_s, s0_1 := and3Rule(nps16_s, nps8_s, ls0_s, nps16_1, nps8_1, ls0_1)
	s0_s = applyFault(s0_s, s0_1, fault_A, "s0")

	s1_s, s1_1 := and3Rule(nps16_s, nps8_s, ls1_s, nps16_1, nps8_1, ls1_1)
	s1_s = applyFault(s1_s, s1_1, fault_A, "s1")

	s2_s, s2_1 := and3Rule(nps16_s, nps8_s, ls2_s, nps16_1, nps8_1, ls2_1)
	s2_s = applyFault(s2_s, s2_1, fault_A, "s2")

	s3_s, s3_1 := and3Rule(nps16_s, nps8_s, ls3_s, nps16_1, nps8_1, ls3_1)
	s3_s = applyFault(s3_s, s3_1, fault_A, "s3")

	s4_s, s4_1 := and3Rule(nps16_s, nps8_s, ls4_s, nps16_1, nps8_1, ls4_1)
	s4_s = applyFault(s4_s, s4_1, fault_A, "s4")

	s5_s, s5_1 := and3Rule(nps16_s, nps8_s, ls5_s, nps16_1, nps8_1, ls5_1)
	s5_s = applyFault(s5_s, s5_1, fault_A, "s5")

	s6_s, s6_1 := and3Rule(nps16_s, nps8_s, ls6_s, nps16_1, nps8_1, ls6_1)
	s6_s = applyFault(s6_s, s6_1, fault_A, "s6")

	s7_s, s7_1 := and3Rule(nps16_s, nps8_s, ls7_s, nps16_1, nps8_1, ls7_1)
	s7_s = applyFault(s7_s, s7_1, fault_A, "s7")

	s8_s, s8_1 := and3Rule(nps16_s, ps8_s, ls0_s, nps16_1, ps8_1, ls0_1)
	s8_s = applyFault(s8_s, s8_1, fault_A, "s8")

	s9_s, s9_1 := and3Rule(nps16_s, ps8_s, ls1_s, nps16_1, ps8_1, ls1_1)
	s9_s = applyFault(s9_s, s9_1, fault_A, "s9")

	s10_s, s10_1 := and3Rule(nps16_s, ps8_s, ls2_s, nps16_1, ps8_1, ls2_1)
	s10_s = applyFault(s10_s, s10_1, fault_A, "s10")

	s11_s, s11_1 := and3Rule(nps16_s, ps8_s, ls3_s, nps16_1, ps8_1, ls3_1)
	s11_s = applyFault(s11_s, s11_1, fault_A, "s11")

	s12_s, s12_1 := and3Rule(nps16_s, ps8_s, ls4_s, nps16_1, ps8_1, ls4_1)
	s12_s = applyFault(s12_s, s12_1, fault_A, "s12")

	s13_s, s13_1 := and3Rule(nps16_s, ps8_s, ls5_s, nps16_1, ps8_1, ls5_1)
	s13_s = applyFault(s13_s, s13_1, fault_A, "s13")

	s14_s, s14_1 := and3Rule(nps16_s, ps8_s, ls6_s, nps16_1, ps8_1, ls6_1)
	s14_s = applyFault(s14_s, s14_1, fault_A, "s14")

	s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
	s15_s = applyFault(s15_s, s15_1, fault_A, "s15")

	s16_s, s16_1 := and3Rule(ps16_s, nps8_s, ls0_s, ps16_1, nps8_1, ls0_1)
	s16_s = applyFault(s16_s