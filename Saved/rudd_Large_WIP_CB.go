// rudd_Large_WIP_CB.go
// refactored by GitHub/Copilot 9 JUN 2025 @ 1445 hrs
// This code simulates a laarge circuit with fault propagation
// using RUDD (Reduced Ordered Decision Diagram).

package main

import (
	"fmt"

	"github.com/dalzilio/rudd"
)

type (
	nd = rudd.Node
	g  struct {
		si, out_4, out_2, out_1, ns nd
	}
	S []g
)

func main() {

	// RUDD SETUP ======================================================
	// =================================================================

	bdd, _ := rudd.New(8, rudd.Nodesize(10000), rudd.Cachesize(3000))

	nd128 := bdd.Ithvar(7)
	nd64 := bdd.Ithvar(6)
	nd32 := bdd.Ithvar(5)
	nd16 := bdd.Ithvar(4)
	nd8 := bdd.Ithvar(3)
	nd4 := bdd.Ithvar(2)
	nd2 := bdd.Ithvar(1)
	nd1 := bdd.Ithvar(0)

	// Define the logical operations on BDD nodes
	not := bdd.Not
	and := bdd.And
	or := bdd.Or
	eq := bdd.Equal

	null := bdd.False()
	// True := bdd.True()
	// False := bdd.False()

	// True is type nd
	// False is type nd
	// true is type Bool
	// false is type Bool

	// COMMON SETUP ====================================================
	// =================================================================

	ps16 := nd128
	ps8 := nd64
	ps4 := nd32
	ps2 := nd16
	ps1 := nd8
	in4 := nd4
	in2 := nd2
	in1 := nd1

	nps16 := not(ps16)
	nps8 := not(ps8)
	nps4 := not(ps4)
	nps2 := not(ps2)
	nps1 := not(ps1)
	nin4 := not(in4)
	nin2 := not(in2)
	nin1 := not(in1)

	s0 := and(nps16, and(nps8, and(nps4, and(nps2, nps1))))
	s1 := and(nps16, and(nps8, and(nps4, and(nps2, ps1))))
	s2 := and(nps16, and(nps8, and(nps4, and(ps2, nps1))))
	s3 := and(nps16, and(nps8, and(nps4, and(ps2, ps1))))
	s4 := and(nps16, and(nps8, and(ps4, and(nps2, nps1))))
	s5 := and(nps16, and(nps8, and(ps4, and(nps2, ps1))))
	s6 := and(nps16, and(nps8, and(ps4, and(ps2, nps1))))
	s7 := and(nps16, and(nps8, and(ps4, and(ps2, ps1))))
	s8 := and(nps16, and(ps8, and(nps4, and(nps2, nps1))))
	s9 := and(nps16, and(ps8, and(nps4, and(nps2, ps1))))
	s10 := and(nps16, and(ps8, and(nps4, and(ps2, nps1))))
	s11 := and(nps16, and(ps8, and(nps4, and(ps2, ps1))))
	s12 := and(nps16, and(ps8, and(ps4, and(nps2, nps1))))
	s13 := and(nps16, and(ps8, and(ps4, and(nps2, ps1))))
	s14 := and(nps16, and(ps8, and(ps4, and(ps2, nps1))))
	s15 := and(nps16, and(ps8, and(ps4, and(ps2, ps1))))
	s16 := and(ps16, and(nps8, and(nps4, and(nps2, nps1))))
	s17 := and(ps16, and(nps8, and(nps4, and(nps2, ps1))))
	s18 := and(ps16, and(nps8, and(nps4, and(ps2, nps1))))
	s19 := and(ps16, and(nps8, and(nps4, and(ps2, ps1))))
	s20 := and(ps16, and(nps8, and(ps4, and(nps2, nps1))))
	s21 := and(ps16, and(nps8, and(ps4, and(nps2, ps1))))
	s22 := and(ps16, and(nps8, and(ps4, and(ps2, nps1))))
	s23 := and(ps16, and(nps8, and(ps4, and(ps2, ps1))))
	s24 := and(ps16, and(ps8, and(nps4, and(nps2, nps1))))
	s25 := and(ps16, and(ps8, and(nps4, and(nps2, ps1))))
	s26 := and(ps16, and(ps8, and(nps4, and(ps2, nps1))))
	s27 := and(ps16, and(ps8, and(nps4, and(ps2, ps1))))
	s28 := and(ps16, and(ps8, and(ps4, and(nps2, nps1))))
	s29 := and(ps16, and(ps8, and(ps4, and(nps2, ps1))))
	s30 := and(ps16, and(ps8, and(ps4, and(ps2, nps1))))
	s31 := and(ps16, and(ps8, and(ps4, and(ps2, ps1))))

	i0 := and(nin4, and(nin2, nin1))
	i1 := and(nin4, and(nin2, in1))
	i2 := and(nin4, and(in2, nin1))
	i3 := and(nin4, and(in2, in1))
	i4 := and(in4, and(nin2, nin1))
	i5 := and(in4, and(nin2, in1))
	i6 := and(in4, and(in2, nin1))
	i7 := and(in4, and(in2, in1))

	s0i0 := and(s0, i0)
	s0i1 := and(s0, i1)
	s0i2 := and(s0, i2)
	s0i3 := and(s0, i3)
	s0i4 := and(s0, i4)
	s0i5 := and(s0, i5)
	s0i6 := and(s0, i6)
	s0i7 := and(s0, i7)
	s1i0 := and(s1, i0)
	s1i1 := and(s1, i1)
	s1i2 := and(s1, i2)
	s1i3 := and(s1, i3)
	s1i4 := and(s1, i4)
	s1i5 := and(s1, i5)
	s1i6 := and(s1, i6)
	s1i7 := and(s1, i7)
	s2i0 := and(s2, i0)
	s2i1 := and(s2, i1)
	s2i2 := and(s2, i2)
	s2i3 := and(s2, i3)
	s2i4 := and(s2, i4)
	s2i5 := and(s2, i5)
	s2i6 := and(s2, i6)
	s2i7 := and(s2, i7)
	s3i0 := and(s3, i0)
	s3i1 := and(s3, i1)
	s3i2 := and(s3, i2)
	s3i3 := and(s3, i3)
	s3i4 := and(s3, i4)
	s3i5 := and(s3, i5)
	s3i6 := and(s3, i6)
	s3i7 := and(s3, i7)
	s4i0 := and(s4, i0)
	s4i1 := and(s4, i1)
	s4i2 := and(s4, i2)
	s4i3 := and(s4, i3)
	s4i4 := and(s4, i4)
	s4i5 := and(s4, i5)
	s4i6 := and(s4, i6)
	s4i7 := and(s4, i7)
	s5i0 := and(s5, i0)
	s5i1 := and(s5, i1)
	s5i2 := and(s5, i2)
	s5i3 := and(s5, i3)
	s5i4 := and(s5, i4)
	s5i5 := and(s5, i5)
	s5i6 := and(s5, i6)
	s5i7 := and(s5, i7)
	s6i0 := and(s6, i0)
	s6i1 := and(s6, i1)
	s6i2 := and(s6, i2)
	s6i3 := and(s6, i3)
	s6i4 := and(s6, i4)
	s6i5 := and(s6, i5)
	s6i6 := and(s6, i6)
	s6i7 := and(s6, i7)
	s7i0 := and(s7, i0)
	s7i1 := and(s7, i1)
	s7i2 := and(s7, i2)
	s7i3 := and(s7, i3)
	s7i4 := and(s7, i4)
	s7i5 := and(s7, i5)
	s7i6 := and(s7, i6)
	s7i7 := and(s7, i7)
	s8i0 := and(s8, i0)
	s8i1 := and(s8, i1)
	s8i2 := and(s8, i2)
	s8i3 := and(s8, i3)
	s8i4 := and(s8, i4)
	s8i5 := and(s8, i5)
	s8i6 := and(s8, i6)
	s8i7 := and(s8, i7)
	s9i0 := and(s9, i0)
	s9i1 := and(s9, i1)
	s9i2 := and(s9, i2)
	s9i3 := and(s9, i3)
	s9i4 := and(s9, i4)
	s9i5 := and(s9, i5)
	s9i6 := and(s9, i6)
	s9i7 := and(s9, i7)
	s10i0 := and(s10, i0)
	s10i1 := and(s10, i1)
	s10i2 := and(s10, i2)
	s10i3 := and(s10, i3)
	s10i4 := and(s10, i4)
	s10i5 := and(s10, i5)
	s10i6 := and(s10, i6)
	s10i7 := and(s10, i7)
	s11i0 := and(s11, i0)
	s11i1 := and(s11, i1)
	s11i2 := and(s11, i2)
	s11i3 := and(s11, i3)
	s11i4 := and(s11, i4)
	s11i5 := and(s11, i5)
	s11i6 := and(s11, i6)
	s11i7 := and(s11, i7)
	s12i0 := and(s12, i0)
	s12i1 := and(s12, i1)
	s12i2 := and(s12, i2)
	s12i3 := and(s12, i3)
	s12i4 := and(s12, i4)
	s12i5 := and(s12, i5)
	s12i6 := and(s12, i6)
	s12i7 := and(s12, i7)
	s13i0 := and(s13, i0)
	s13i1 := and(s13, i1)
	s13i2 := and(s13, i2)
	s13i3 := and(s13, i3)
	s13i4 := and(s13, i4)
	s13i5 := and(s13, i5)
	s13i6 := and(s13, i6)
	s13i7 := and(s13, i7)
	s14i0 := and(s14, i0)
	s14i1 := and(s14, i1)
	s14i2 := and(s14, i2)
	s14i3 := and(s14, i3)
	s14i4 := and(s14, i4)
	s14i5 := and(s14, i5)
	s14i6 := and(s14, i6)
	s14i7 := and(s14, i7)
	s15i0 := and(s15, i0)
	s15i1 := and(s15, i1)
	s15i2 := and(s15, i2)
	s15i3 := and(s15, i3)
	s15i4 := and(s15, i4)
	s15i5 := and(s15, i5)
	s15i6 := and(s15, i6)
	s15i7 := and(s15, i7)
	s16i0 := and(s16, i0)
	s16i1 := and(s16, i1)
	s16i2 := and(s16, i2)
	s16i3 := and(s16, i3)
	s16i4 := and(s16, i4)
	s16i5 := and(s16, i5)
	s16i6 := and(s16, i6)
	s16i7 := and(s16, i7)
	s17i0 := and(s17, i0)
	s17i1 := and(s17, i1)
	s17i2 := and(s17, i2)
	s17i3 := and(s17, i3)
	s17i4 := and(s17, i4)
	s17i5 := and(s17, i5)
	s17i6 := and(s17, i6)
	s17i7 := and(s17, i7)
	s18i0 := and(s18, i0)
	s18i1 := and(s18, i1)
	s18i2 := and(s18, i2)
	s18i3 := and(s18, i3)
	s18i4 := and(s18, i4)
	s18i5 := and(s18, i5)
	s18i6 := and(s18, i6)
	s18i7 := and(s18, i7)
	s19i0 := and(s19, i0)
	s19i1 := and(s19, i1)
	s19i2 := and(s19, i2)
	s19i3 := and(s19, i3)
	s19i4 := and(s19, i4)
	s19i5 := and(s19, i5)
	s19i6 := and(s19, i6)
	s19i7 := and(s19, i7)
	s20i0 := and(s20, i0)
	s20i1 := and(s20, i1)
	s20i2 := and(s20, i2)
	s20i3 := and(s20, i3)
	s20i4 := and(s20, i4)
	s20i5 := and(s20, i5)
	s20i6 := and(s20, i6)
	s20i7 := and(s20, i7)
	s21i0 := and(s21, i0)
	s21i1 := and(s21, i1)
	s21i2 := and(s21, i2)
	s21i3 := and(s21, i3)
	s21i4 := and(s21, i4)
	s21i5 := and(s21, i5)
	s21i6 := and(s21, i6)
	s21i7 := and(s21, i7)
	s22i0 := and(s22, i0)
	s22i1 := and(s22, i1)
	s22i2 := and(s22, i2)
	s22i3 := and(s22, i3)
	s22i4 := and(s22, i4)
	s22i5 := and(s22, i5)
	s22i6 := and(s22, i6)
	s22i7 := and(s22, i7)
	s23i0 := and(s23, i0)
	s23i1 := and(s23, i1)
	s23i2 := and(s23, i2)
	s23i3 := and(s23, i3)
	s23i4 := and(s23, i4)
	s23i5 := and(s23, i5)
	s23i6 := and(s23, i6)
	s23i7 := and(s23, i7)
	s24i0 := and(s24, i0)
	s24i1 := and(s24, i1)
	s24i2 := and(s24, i2)
	s24i3 := and(s24, i3)
	s24i4 := and(s24, i4)
	s24i5 := and(s24, i5)
	s24i6 := and(s24, i6)
	s24i7 := and(s24, i7)
	s25i0 := and(s25, i0)
	s25i1 := and(s25, i1)
	s25i2 := and(s25, i2)
	s25i3 := and(s25, i3)
	s25i4 := and(s25, i4)
	s25i5 := and(s25, i5)
	s25i6 := and(s25, i6)
	s25i7 := and(s25, i7)
	s26i0 := and(s26, i0)
	s26i1 := and(s26, i1)
	s26i2 := and(s26, i2)
	s26i3 := and(s26, i3)
	s26i4 := and(s26, i4)
	s26i5 := and(s26, i5)
	s26i6 := and(s26, i6)
	s26i7 := and(s26, i7)
	s27i0 := and(s27, i0)
	s27i1 := and(s27, i1)
	s27i2 := and(s27, i2)
	s27i3 := and(s27, i3)
	s27i4 := and(s27, i4)
	s27i5 := and(s27, i5)
	s27i6 := and(s27, i6)
	s27i7 := and(s27, i7)
	s28i0 := and(s28, i0)
	s28i1 := and(s28, i1)
	s28i2 := and(s28, i2)
	s28i3 := and(s28, i3)
	s28i4 := and(s28, i4)
	s28i5 := and(s28, i5)
	s28i6 := and(s28, i6)
	s28i7 := and(s28, i7)
	s29i0 := and(s29, i0)
	s29i1 := and(s29, i1)
	s29i2 := and(s29, i2)
	s29i3 := and(s29, i3)
	s29i4 := and(s29, i4)
	s29i5 := and(s29, i5)
	s29i6 := and(s29, i6)
	s29i7 := and(s29, i7)
	s30i0 := and(s30, i0)
	s30i1 := and(s30, i1)
	s30i2 := and(s30, i2)
	s30i3 := and(s30, i3)
	s30i4 := and(s30, i4)
	s30i5 := and(s30, i5)
	s30i6 := and(s30, i6)
	s30i7 := and(s30, i7)
	s31i0 := and(s31, i0)
	s31i1 := and(s31, i1)
	s31i2 := and(s31, i2)
	s31i3 := and(s31, i3)
	s31i4 := and(s31, i4)
	s31i5 := and(s31, i5)
	s31i6 := and(s31, i6)
	s31i7 := and(s31, i7)

	/* NOTES ==========================================================

	    try these faults: e42:0 d2:0 c6:0 e18:1 e5:1 c5:1 s6:0

	    because they have longer test sequences.

	================================================================= */

	// COMMON LOGIC GATE RULES =========================================
	// =================================================================

	// These functions create fault-propagation (name_s) and 1-set
	// path-enabling (name_1) rules for a primary input, the 'NOT'
	// gate, and for 2- and 3-input AND,and OR gates.

	piRule := func(s1, i1 nd) (nd, nd) {
		// computes the propagation function
		o_s := s1
		// computes the 1-set
		o_1 := i1
		return o_s, o_1
	}

	notRule := func(s1, i1 nd) (nd, nd) {
		// computes the propagation function
		o_s := s1
		// computes the 1-set
		o_1 := not(i1)
		return o_s, o_1
	}

	and2Rule := func(s1, s2, i1, i2 nd) (nd, nd) {
		// computes the propagation function
		o_s := or(and(not(i1), s1, not(i2), s2),
			and(i2, s1, not(s2)),
			and(i1, s2, not(s1)),
			and(i1, i2, or(s1, s2)))
		// computes the 1-set
		o_1 := and(i1, i2)
		return o_s, o_1
	}

	or2Rule := func(s1, s2, i1, i2 nd) (nd, nd) {
		// computes the propagation function
		o_s := or(and(i1, s1, i2, s2),
			and(not(i2), s1, not(s2)),
			and(not(i1), s2, not(s1)),
			and(not(i1), not(i2), or(s1, s2)))
		// computes the 1-set
		o_1 := or(i1, i2)
		return o_s, o_1
	}

	and3Rule := func(s1, s2, s3, i1, i2, i3 nd) (nd, nd) {
		s, i := and2Rule(s1, s2, i1, i2)
		o_s, o_1 := and2Rule(s, s3, i, i3)
		return o_s, o_1
	}

	or3Rule := func(s1, s2, s3, i1, i2, i3 nd) (nd, nd) {
		s, i := or2Rule(s1, s2, i1, i2)
		o_s, o_1 := or2Rule(s, s3, i, i3)
		return o_s, o_1

	} // end of rules ==================================================

	// allSAT ==========================================================
	// =================================================================

	// allSat receives a BCD object and parses it into constituent S/Is,
	// returning a list of zero or more S/Is, each of type string.

	allSAT := func(f nd, str2nd func(string) nd) []string {
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

				// Check if the conjunction of `f` and the current state-input is not null
				if and(f, ndValue) != null {
					g = append(g, si)
				}
			}
		}

		return g // Return the list of S/I's in the nd function
	}

	// ps2ns ===========================================================
	// =================================================================

	ps2ns := func(ps16_i, ps8_i, ps4_i, ps2_i, ps1_i nd, fault_A string) (nd, nd, nd, nd, nd, nd, nd, nd) {

		// Receives S/Is via present-state lines ps16_s, ps8_s. ect.
		// Uses fault to activate local_fault; propagates local_fault
		// to output and next_state lines, out4_s, out2_s, out1_s,
		// ns16_s, ns8_s, etc., as S/Is.

		// Helper function to apply faults to signals
		applyFault := func(signal nd, faultSignal nd, fault string, faultName string) nd {
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

		//-----------------------------------

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
		s31_s = applyFault(s31_s, s31_1, fault_A, "s31")

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
		s16_s = applyFault(s16_s, s16_1, fault_A, "s16")

		s17_s, s17_1 := and3Rule(ps16_s, nps8_s, ls1_s, ps16_1, nps8_1, ls1_1)
		s17_s = applyFault(s17_s, s17_1, fault_A, "s17")

		s18_s, s18_1 := and3Rule(ps16_s, nps8_s, ls2_s, ps16_1, nps8_1, ls2_1)
		s18_s = applyFault(s18_s, s18_1, fault_A, "s18")

		s19_s, s19_1 := and3Rule(ps16_s, nps8_s, ls3_s, ps16_1, nps8_1, ls3_1)
		s19_s = applyFault(s19_s, s19_1, fault_A, "s19")

		s20_s, s20_1 := and3Rule(ps16_s, nps8_s, ls4_s, ps16_1, nps8_1, ls4_1)
		s20_s = applyFault(s20_s, s20_1, fault_A, "s20")

		s21_s, s21_1 := and3Rule(ps16_s, nps8_s, ls5_s, ps16_1, nps8_1, ls5_1)
		s21_s = applyFault(s21_s, s21_1, fault_A, "s21")

		s22_s, s22_1 := and3Rule(ps16_s, nps8_s, ls6_s, ps16_1, nps8_1, ls6_1)
		s22_s = applyFault(s22_s, s22_1, fault_A, "s22")

		s23_s, s23_1 := and3Rule(ps16_s, nps8_s, ls7_s, ps16_1, nps8_1, ls7_1)
		s23_s = applyFault(s23_s, s23_1, fault_A, "s23")

		s24_s, s24_1 := and3Rule(ps16_s, ps8_s, ls0_s, ps16_1, ps8_1, ls0_1)
		s24_s = applyFault(s24_s, s24_1, fault_A, "s24")

		s25_s, s25_1 := and3Rule(ps16_s, ps8_s, ls1_s, ps16_1, ps8_1, ls1_1)
		s25_s = applyFault(s25_s, s25_1, fault_A, "s25")

		s26_s, s26_1 := and3Rule(ps16_s, ps8_s, ls2_s, ps16_1, ps8_1, ls2_1)
		s26_s = applyFault(s26_s, s26_1, fault_A, "s26")

		s27_s, s27_1 := and3Rule(ps16_s, ps8_s, ls3_s, ps16_1, ps8_1, ls3_1)
		s27_s = applyFault(s27_s, s27_1, fault_A, "s27")

		s28_s, s28_1 := and3Rule(ps16_s, ps8_s, ls4_s, ps16_1, ps8_1, ls4_1)
		s28_s = applyFault(s28_s, s28_1, fault_A, "s28")

		s29_s, s29_1 := and3Rule(ps16_s, ps8_s, ls5_s, ps16_1, ps8_1, ls5_1)
		s29_s = applyFault(s29_s, s29_1, fault_A, "s29")

		s30_s, s30_1 := and3Rule(ps16_s, ps8_s, ls6_s, ps16_1, ps8_1, ls6_1)
		s30_s = applyFault(s30_s, s30_1, fault_A, "s30")

		b2_s, b2_1 := or3Rule(i5_s, i3_s, i2_s, i5_1, i3_1, i2_1)
		b2_s = applyFault(b2_s, b2_1, fault_A, "b2")

		b7_s, b7_1 := or2Rule(i5_s, i1_s, i5_1, i1_1)
		b7_s = applyFault(b7_s, b7_1, fault_A, "b7")

		c5_s, c5_1 := or2Rule(i7_s, i5_s, i7_1, i5_1)
		c5_s = applyFault(c5_s, c5_1, fault_A, "c5")

		c13_s, c13_1 := or2Rule(i3_s, i2_s, i3_1, i2_1)
		c13_s = applyFault(c13_s, c13_1, fault_A, "c13")

		e5_s, e5_1 := or2Rule(i5_s, i4_s, i5_1, i4_1)
		e5_s = applyFault(e5_s, e5_1, fault_A, "e5")

		e10_s, e10_1 := or3Rule(i6_s, i2_s, i0_s, i6_1, i2_1, i0_1)
		e10_s = applyFault(e10_s, e10_1, fault_A, "e10")

		e12_s, e12_1 := or2Rule(i6_s, i3_s, i6_1, i3_1)
		e12_s = applyFault(e12_s, e12_1, fault_A, "e12")

		e18_s, e18_1 := or2Rule(i6_s, i2_s, i6_1, i2_1)
		e18_s = applyFault(e18_s, e18_1, fault_A, "e18")

		e24_s, e24_1 := or2Rule(i7_s, i1_s, i7_1, i1_1)
		e24_s = applyFault(e24_s, e24_1, fault_A, "e24")

		// level 4

		a1_s, a1_1 := and2Rule(s10_s, i0_s, s10_1, i0_1)
		a1_s = applyFault(a1_s, a1_1, fault_A, "a1")

		a2_s, a2_1 := and2Rule(s15_s, i5_s, s15_1, i5_1)
		a2_s = applyFault(a2_s, a2_1, fault_A, "a2")

		a3_s, a3_1 := and2Rule(s18_s, ni6_s, s18_1, ni6_1)
		a3_s = applyFault(a3_s, a3_1, fault_A, "a3")

		a4_s, a4_1 := or3Rule(s20_s, s21_s, s22_s, s20_1, s21_1, s22_1)
		a4_s = applyFault(a4_s, a4_1, fault_A, "a4")

		a5_s, a5_1 := and2Rule(s24_s, ni7_s, s24_1, ni7_1)
		a5_s = applyFault(a5_s, a5_1, fault_A, "a5")

		a6_s, a6_1 := or2Rule(s25_s, s26_s, s25_1, s26_1)
		a6_s = applyFault(a6_s, a6_1, fault_A, "a6")

		a7_s, a7_1 := or3Rule(s27_s, s28_s, s29_s, s27_1, s28_1, s29_1)
		a7_s = applyFault(a7_s, a7_1, fault_A, "a7")

		a9_s, a9_1 := and3Rule(s31_s, ni5_s, ni2_s, s31_1, ni5_1, ni2_1)
		a9_s = applyFault(a9_s, a9_1, fault_A, "a9")

		b1_s, b1_1 := and2Rule(s3_s, i2_s, s3_1, i2_1)
		b1_s = applyFault(b1_s, b1_1, fault_A, "b1")

		b3_s, b3_1 := and2Rule(b2_s, s7_s, b2_1, s7_1)
		b3_s = applyFault(b3_s, b3_1, fault_A, "b3")

		b4_s, b4_1 := and2Rule(s10_s, ni6_s, s10_1, ni6_1)
		b4_s = applyFault(b4_s, b4_1, fault_A, "b4")

		b5_s, b5_1 := or3Rule(s12_s, s13_s, s14_s, s12_1, s13_1, s14_1)
		b5_s = applyFault(b5_s, b5_1, fault_A, "b5")

		b6_s, b6_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		b6_s = applyFault(b6_s, b6_1, fault_A, "b6")

		b8_s, b8_1 := and2Rule(s23_s, b7_s, s23_1, b7_1)
		b8_s = applyFault(b8_s, b8_1, fault_A, "b8")

		b9_s, b9_1 := or2Rule(s25_s, s26_s, s25_1, s26_1)
		b9_s = applyFault(b9_s, b9_1, fault_A, "b9")

		b10_s, b10_1 := or3Rule(s27_s, s28_s, s29_s, s27_1, s28_1, s29_1)
		b10_s = applyFault(b10_s, b10_1, fault_A, "b10")

		c2_s, c2_1 := and3Rule(s7_s, ni5_s, ni3_s, s7_1, ni5_1, ni3_1)
		c2_s = applyFault(c2_s, c2_1, fault_A, "c2")

		c3_s, c3_1 := and2Rule(s11_s, i7_s, s11_1, i7_1)
		c3_s = applyFault(c3_s, c3_1, fault_A, "c3")

		c4_s, c4_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		c4_s = applyFault(c4_s, c4_1, fault_A, "c4")

		c6_s, c6_1 := and2Rule(s23_s, ni5_s, s23_1, ni5_1)
		c6_s = applyFault(c6_s, c6_1, fault_A, "c6")

		c8_s, c8_1 := and2Rule(s19_s, c5_s, s19_1, c5_1)
		c8_s = applyFault(c8_s, c8_1, fault_A, "c8")

		c14_s, c14_1 := and2Rule(c13_s, s3_s, c13_1, s3_1)
		c14_s = applyFault(c14_s, c14_1, fault_A, "c14")

		c15_s, c15_1 := and2Rule(s27_s, i7_s, s27_1, i7_1)
		c15_s = applyFault(c15_s, c15_1, fault_A, "c15")

		d1_s, d1_1 := and2Rule(s1_s, i2_s, s1_1, i2_1)
		d1_s = applyFault(d1_s, d1_1, fault_A, "d1")

		d2_s, d2_1 := and3Rule(s3_s, ni3_s, ni2_s, s3_1, ni3_1, ni2_1)
		d2_s = applyFault(d2_s, d2_1, fault_A, "d2")

		d3_s, d3_1 := and2Rule(s5_s, i0_s, s5_1, i0_1)
		d3_s = applyFault(d3_s, d3_1, fault_A, "d3")

		d5_s, d5_1 := and2Rule(s9_s, i2_s, s9_1, i2_1)
		d5_s = applyFault(d5_s, d5_1, fault_A, "d5")

		d6_s, d6_1 := and2Rule(s11_s, ni7_s, s11_1, ni7_1)
		d6_s = applyFault(d6_s, d6_1, fault_A, "d6")

		d7_s, d7_1 := and2Rule(s13_s, i0_s, s13_1, i0_1)
		d7_s = applyFault(d7_s, d7_1, fault_A, "d7")

		d9_s, d9_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		d9_s = applyFault(d9_s, d9_1, fault_A, "d9")

		d10_s, d10_1 := and2Rule(s17_s, i2_s, s17_1, i2_1)
		d10_s = applyFault(d10_s, d10_1, fault_A, "d10")

		d11_s, d11_1 := and2Rule(s19_s, ni7_s, s19_1, ni7_1)
		d11_s = applyFault(d11_s, d11_1, fault_A, "d11")

		d12_s, d12_1 := and2Rule(s21_s, i0_s, s21_1, i0_1)
		d12_s = applyFault(d12_s, d12_1, fault_A, "d12")

		d13_s, d13_1 := and3Rule(s23_s, ni5_s, ni1_s, s23_1, ni5_1, ni1_1)
		d13_s = applyFault(d13_s, d13_1, fault_A, "d13")

		d14_s, d14_1 := and2Rule(s25_s, i2_s, s25_1, i2_1)
		d14_s = applyFault(d14_s, d14_1, fault_A, "d14")

		d15_s, d15_1 := and2Rule(s29_s, i0_s, s29_1, i0_1)
		d15_s = applyFault(d15_s, d15_1, fault_A, "d15")

		d27_s, d27_1 := and2Rule(s27_s, ni7_s, s27_1, ni7_1)
		d27_s = applyFault(d27_s, d27_1, fault_A, "d27")

		e1_s, e1_1 := and2Rule(s0_s, i1_s, s0_1, i1_1)
		e1_s = applyFault(e1_s, e1_1, fault_A, "e1")

		e2_s, e2_1 := and2Rule(s1_s, ni2_s, s1_1, ni2_1)
		e2_s = applyFault(e2_s, e2_1, fault_A, "e2")

		e3_s, e3_1 := and2Rule(s2_s, i2_s, s2_1, i2_1)
		e3_s = applyFault(e3_s, e3_1, fault_A, "e3")

		e4_s, e4_1 := and3Rule(s3_s, ni3_s, ni2_s, s3_1, ni3_1, ni2_1)
		e4_s = applyFault(e4_s, e4_1, fault_A, "e4")

		e6_s, e6_1 := and2Rule(s5_s, ni0_s, s5_1, ni0_1)
		e6_s = applyFault(e6_s, e6_1, fault_A, "e6")

		e7_s, e7_1 := and2Rule(s6_s, i7_s, s6_1, i7_1)
		e7_s = applyFault(e7_s, e7_1, fault_A, "e7")

		e8_s, e8_1 := and2Rule(s8_s, i1_s, s8_1, i1_1)
		e8_s = applyFault(e8_s, e8_1, fault_A, "e8")

		e9_s, e9_1 := and2Rule(s9_s, ni2_s, s9_1, ni2_1)
		e9_s = applyFault(e9_s, e9_1, fault_A, "e9")

		e11_s, e11_1 := and2Rule(s11_s, ni7_s, s11_1, ni7_1)
		e11_s = applyFault(e11_s, e11_1, fault_A, "e11")

		e13_s, e13_1 := and2Rule(s13_s, ni0_s, s13_1, ni0_1)
		e13_s = applyFault(e13_s, e13_1, fault_A, "e13")

		e14_s, e14_1 := and2Rule(s14_s, i7_s, s14_1, i7_1)
		e14_s = applyFault(e14_s, e14_1, fault_A, "e14")

		e15_s, e15_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		e15_s = applyFault(e15_s, e15_1, fault_A, "e15")

		e16_s, e16_1 := and2Rule(s16_s, i1_s, s16_1, i1_1)
		e16_s = applyFault(e16_s, e16_1, fault_A, "e16")

		e17_s, e17_1 := and2Rule(s17_s, ni2_s, s17_1, ni2_1)
		e17_s = applyFault(e17_s, e17_1, fault_A, "e17")

		e19_s, e19_1 := and2Rule(s19_s, ni7_s, s19_1, ni7_1)
		e19_s = applyFault(e19_s, e19_1, fault_A, "e19")

		e20_s, e20_1 := and2Rule(s20_s, e12_s, s20_1, e12_1)
		e20_s = applyFault(e20_s, e20_1, fault_A, "e20")

		e21_s, e21_1 := and2Rule(s21_s, ni0_s, s21_1, ni0_1)
		e21_s = applyFault(e21_s, e21_1, fault_A, "e21")

		e22_s, e22_1 := and2Rule(s22_s, i7_s, s22_1, i7_1)
		e22_s = applyFault(e22_s, e22_1, fault_A, "e22")

		e23_s, e23_1 := and2Rule(s23_s, ni5_s, s23_1, ni5_1)
		e23_s = applyFault(e23_s, e23_1, fault_A, "e23")

		e25_s, e25_1 := and2Rule(s25_s, ni2_s, s25_1, ni2_1)
		e25_s = applyFault(e25_s, e25_1, fault_A, "e25")

		e26_s, e26_1 := and2Rule(s26_s, i2_s, s26_1, i2_1)
		e26_s = applyFault(e26_s, e26_1, fault_A, "e26")

		e27_s, e27_1 := and2Rule(s27_s, ni7_s, s27_1, ni7_1)
		e27_s = applyFault(e27_s, e27_1, fault_A, "e27")

		e28_s, e28_1 := and2Rule(s28_s, e12_s, s28_1, e12_1)
		e28_s = applyFault(e28_s, e28_1, fault_A, "e28")

		e29_s, e29_1 := and2Rule(s29_s, ni0_s, s29_1, ni0_1)
		e29_s = applyFault(e29_s, e29_1, fault_A, "e29")

		e30_s, e30_1 := and2Rule(s30_s, i7_s, s30_1, i7_1)
		e30_s = applyFault(e30_s, e30_1, fault_A, "e30")

		e31_s, e31_1 := and2Rule(s4_s, e5_s, s4_1, e5_1)
		e31_s = applyFault(e31_s, e31_1, fault_A, "e31")

		e32_s, e32_1 := and2Rule(s10_s, e10_s, s10_1, e10_1)
		e32_s = applyFault(e32_s, e32_1, fault_A, "e32")

		e33_s, e33_1 := and2Rule(s12_s, e12_s, s12_1, e12_1)
		e33_s = applyFault(e33_s, e33_1, fault_A, "e33")

		e34_s, e34_1 := and2Rule(s18_s, e18_s, s18_1, e18_1)
		e34_s = applyFault(e34_s, e34_1, fault_A, "e34")

		e35_s, e35_1 := and2Rule(s24_s, e24_s, s24_1, e24_1)
		e35_s = applyFault(e35_s, e35_1, fault_A, "e35")

		f1_s, f1_1 := and2Rule(s12_s, i5_s, s12_1, i5_1)
		f1_s = applyFault(f1_s, f1_1, fault_A, "f1")

		f2_s, f2_1 := and2Rule(s27_s, i4_s, s27_1, i4_1)
		f2_s = applyFault(f2_s, f2_1, fault_A, "f2")

		f3_s, f3_1 := and2Rule(s15_s, i0_s, s15_1, i0_1)
		f3_s = applyFault(f3_s, f3_1, fault_A, "f3")

		f4_s, f4_1 := and2Rule(s27_s, i2_s, s27_1, i2_1)
		f4_s = applyFault(f4_s, f4_1, fault_A, "f4")

		f5_s, f5_1 := and2Rule(s0_s, i7_s, s0_1, i7_1)
		f5_s = applyFault(f5_s, f5_1, fault_A, "f5")

		f6_s, f6_1 := and2Rule(s27_s, i1_s, s27_1, i1_1)
		f6_s = applyFault(f6_s, f6_1, fault_A, "f6")

		// level 5

		a8_s, a8_1 := or2Rule(a7_s, s30_s, a7_1, s30_1)
		a8_s = applyFault(a8_s, a8_1, fault_A, "a8")

		a10_s, a10_1 := or3Rule(a1_s, a2_s, s16_s, a1_1, a2_1, s16_1)
		a10_s = applyFault(a10_s, a10_1, fault_A, "a10")

		b11_s, b11_1 := or2Rule(b10_s, s30_s, b10_1, s30_1)
		b11_s = applyFault(b11_s, b11_1, fault_A, "b11")

		b12_s, b12_1 := or3Rule(b1_s, b3_s, s8_s, b1_1, b3_1, s8_1)
		b12_s = applyFault(b12_s, b12_1, fault_A, "b12")

		b13_s, b13_1 := or3Rule(s9_s, b4_s, s11_s, s9_1, b4_1, s11_1)
		b13_s = applyFault(b13_s, b13_1, fault_A, "b13")

		c1_s, c1_1 := or3Rule(c14_s, s4_s, s5_s, c14_1, s4_1, s5_1)
		c1_s = applyFault(c1_s, c1_1, fault_A, "c1")

		c7_s, c7_1 := and2Rule(c2_s, ni2_s, c2_1, ni2_1)
		c7_s = applyFault(c7_s, c7_1, fault_A, "c7")

		c16_s, c16_1 := or3Rule(c15_s, s28_s, s29_s, c15_1, s28_1, s29_1)
		c16_s = applyFault(c16_s, c16_1, fault_A, "c16")

		d17_s, d17_1 := or3Rule(d1_s, d2_s, d3_s, d1_1, d2_1, d3_1)
		d17_s = applyFault(d17_s, d17_1, fault_A, "d17")

		d19_s, d19_1 := or3Rule(s10_s, d6_s, d7_s, s10_1, d6_1, d7_1)
		d19_s = applyFault(d19_s, d19_1, fault_A, "d19")

		d20_s, d20_1 := or3Rule(s14_s, d9_s, d10_s, s14_1, d9_1, d10_1)
		d20_s = applyFault(d20_s, d20_1, fault_A, "d20")

		d21_s, d21_1 := or3Rule(s18_s, d11_s, d12_s, s18_1, d11_1, d12_1)
		d21_s = applyFault(d21_s, d21_1, fault_A, "d21")

		d22_s, d22_1 := or3Rule(s22_s, d13_s, d14_s, s22_1, d13_1, d14_1)
		d22_s = applyFault(d22_s, d22_1, fault_A, "d22")

		d23_s, d23_1 := or3Rule(s26_s, d15_s, s30_s, s26_1, d15_1, s30_1)
		d23_s = applyFault(d23_s, d23_1, fault_A, "d23")

		e36_s, e36_1 := or3Rule(e1_s, e2_s, e3_s, e1_1, e2_1, e3_1)
		e36_s = applyFault(e36_s, e36_1, fault_A, "e36")

		e37_s, e37_1 := or3Rule(e4_s, e31_s, e6_s, e4_1, e31_1, e6_1)
		e37_s = applyFault(e37_s, e37_1, fault_A, "e37")

		e39_s, e39_1 := or2Rule(e9_s, e32_s, e9_1, e32_1)
		e39_s = applyFault(e39_s, e39_1, fault_A, "e39")

		e40_s, e40_1 := or3Rule(e11_s, e33_s, e13_s, e11_1, e33_1, e13_1)
		e40_s = applyFault(e40_s, e40_1, fault_A, "e40")

		e41_s, e41_1 := or3Rule(e14_s, e15_s, e16_s, e14_1, e15_1, e16_1)
		e41_s = applyFault(e41_s, e41_1, fault_A, "e41")

		e42_s, e42_1 := or3Rule(e17_s, e34_s, e19_s, e17_1, e34_1, e19_1)
		e42_s = applyFault(e42_s, e42_1, fault_A, "e42")

		e43_s, e43_1 := or3Rule(e20_s, e30_s, a9_s, e20_1, e30_1, a9_1)
		e43_s = applyFault(e43_s, e43_1, fault_A, "e43")

		e44_s, e44_1 := or3Rule(e21_s, e22_s, e23_s, e21_1, e22_1, e23_1)
		e44_s = applyFault(e44_s, e44_1, fault_A, "e44")

		e45_s, e45_1 := or3Rule(e35_s, e25_s, e26_s, e35_1, e25_1, e26_1)
		e45_s = applyFault(e45_s, e45_1, fault_A, "e45")

		e46_s, e46_1 := or3Rule(e27_s, e28_s, e29_s, e27_1, e28_1, e29_1)
		e46_s = applyFault(e46_s, e46_1, fault_A, "e46")

		out4_s, out4_1 := or2Rule(f1_s, f2_s, f1_1, f2_1)
		out4_s = applyFault(out4_s, out4_1, fault_A, "out4")

		out2_s, out2_1 := or2Rule(f3_s, f4_s, f3_1, f4_1)
		out2_s = applyFault(out2_s, out2_1, fault_A, "out2")

		out1_s, out1_1 := or2Rule(f5_s, f6_s, f5_1, f6_1)
		out1_s = applyFault(out1_s, out1_1, fault_A, "out1")

		// level 6

		a11_s, a11_1 := or3Rule(a10_s, s17_s, a3_s, a10_1, s17_1, a3_1)
		a11_s = applyFault(a11_s, a11_1, fault_A, "a11")

		b14_s, b14_1 := or3Rule(b12_s, b13_s, b5_s, b12_1, b13_1, b5_1)
		b14_s = applyFault(b14_s, b14_1, fault_A, "b14")

		c9_s, c9_1 := or3Rule(c1_s, s6_s, c7_s, c1_1, s6_1, c7_1)
		c9_s = applyFault(c9_s, c9_1, fault_A, "c9")

		c17_s, c17_1 := or2Rule(c16_s, s30_s, c16_1, s30_1)
		c17_s = applyFault(c17_s, c17_1, fault_A, "c17")

		d18_s, d18_1 := or3Rule(s6_s, c7_s, d5_s, s6_1, c7_1, d5_1)
		d18_s = applyFault(d18_s, d18_1, fault_A, "d18")

		d28_s, d28_1 := or2Rule(d23_s, d27_s, d23_1, d27_1)
		d28_s = applyFault(d28_s, d28_1, fault_A, "d28")

		e38_s, e38_1 := or3Rule(e7_s, c7_s, e8_s, e7_1, c7_1, e8_1)
		e38_s = applyFault(e38_s, e38_1, fault_A, "e38")

		e47_s, e47_1 := or3Rule(e36_s, e37_s, e44_s, e36_1, e37_1, e44_1)
		e47_s = applyFault(e47_s, e47_1, fault_A, "e47")

		e49_s, e49_1 := or3Rule(e39_s, e40_s, e41_s, e39_1, e40_1, e41_1)
		e49_s = applyFault(e49_s, e49_1, fault_A, "e49")

		// level 7

		a12_s, a12_1 := or3Rule(a11_s, s19_s, a4_s, a11_1, s19_1, a4_1)
		a12_s = applyFault(a12_s, a12_1, fault_A, "a12")

		b15_s, b15_1 := or3Rule(b14_s, b6_s, b8_s, b14_1, b6_1, b8_1)
		b15_s = applyFault(b15_s, b15_1, fault_A, "b15")

		c10_s, c10_1 := or3Rule(c9_s, c3_s, b5_s, c9_1, c3_1, b5_1)
		c10_s = applyFault(c10_s, c10_1, fault_A, "c10")

		d24_s, d24_1 := or3Rule(d17_s, d18_s, d19_s, d17_1, d18_1, d19_1)
		d24_s = applyFault(d24_s, d24_1, fault_A, "d24")

		e48_s, e48_1 := or3Rule(e45_s, e46_s, e38_s, e45_1, e46_1, e38_1)
		e48_s = applyFault(e48_s, e48_1, fault_A, "e48")

		// level 8

		a13_s, a13_1 := or3Rule(a12_s, s23_s, a5_s, a12_1, s23_1, a5_1)
		a13_s = applyFault(a13_s, a13_1, fault_A, "a13")

		b16_s, b16_1 := or3Rule(b15_s, s24_s, b9_s, b15_1, s24_1, b9_1)
		b16_s = applyFault(b16_s, b16_1, fault_A, "b16")

		c11_s, c11_1 := or3Rule(c10_s, c4_s, c8_s, c10_1, c4_1, c8_1)
		c11_s = applyFault(c11_s, c11_1, fault_A, "c11")

		d25_s, d25_1 := or3Rule(d24_s, d20_s, d21_s, d24_1, d20_1, d21_1)
		d25_s = applyFault(d25_s, d25_1, fault_A, "d25")

		e50_s, e50_1 := or3Rule(e47_s, e48_s, e49_s, e47_1, e48_1, e49_1)
		e50_s = applyFault(e50_s, e50_1, fault_A, "e50")

		// level 9

		ns1_s, ns1_1 := or3Rule(e50_s, e42_s, e43_s, e50_1, e42_1, e43_1)
		ns1_s = applyFault(ns1_s, ns1_1, fault_A, "ns1")

		ns8_s, ns8_1 := or3Rule(b16_s, b11_s, a9_s, b16_1, b11_1, a9_1)
		ns8_s = applyFault(ns8_s, ns8_1, fault_A, "ns8")

		a14_s, a14_1 := or3Rule(a13_s, a6_s, a8_s, a13_1, a6_1, a8_1)
		a14_s = applyFault(a14_s, a14_1, fault_A, "a14")

		c12_s, c12_1 := or3Rule(c11_s, a4_s, c6_s, c11_1, a4_1, c6_1)
		c12_s = applyFault(c12_s, c12_1, fault_A, "c12")

		d26_s, d26_1 := or3Rule(d25_s, d22_s, d28_s, d25_1, d22_1, d28_1)
		d26_s = applyFault(d26_s, d26_1, fault_A, "d26")

		// level 10

		ns2_s, ns2_1 := or3Rule(d26_s, s2_s, a9_s, d26_1, s2_1, a9_1)
		ns2_s = applyFault(ns2_s, ns2_1, fault_A, "ns2")

		ns4_s, ns4_1 := or3Rule(c12_s, c17_s, a9_s, c12_1, c17_1, a9_1)
		ns4_s = applyFault(ns4_s, ns4_1, fault_A, "ns4")

		ns16_s, ns16_1 := or2Rule(a14_s, a9_s, a14_1, a9_1)
		ns16_s = applyFault(ns16_s, ns16_1, fault_A, "ns16")

		return out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s

	} // End of ps2ns() ================================================

	// ns2fp ===========================================================
	// =================================================================

	ns2fp := func(out4_s, out2_s, out1_s,
		ns16_s, ns8_s, ns4_s, ns2_s, ns1_s nd) []S {

		// create fault pattern collections
		var fp1, fp2, fp3, fp4, fp5, fp6, fp7, fp8,
			fp9, fp10, fp11, fp12, fp13, fp14, fp15, fp16,
			fp17, fp18, fp19, fp20, fp21, fp22, fp23,
			fp24, fp25, fp26, fp27, fp28, fp29, fp30, fp31 nd

		// fault patterns for ns2fp
		fp1 = and(not(ns16_s), not(ns8_s), not(ns4_s), not(ns2_s), ns1_s)
		fp2 = and(not(ns16_s), not(ns8_s), not(ns4_s), ns2_s, not(ns1_s))
		fp3 = and(not(ns16_s), not(ns8_s), not(ns4_s), ns2_s, ns1_s)
		fp4 = and(not(ns16_s), not(ns8_s), ns4_s, not(ns2_s), not(ns1_s))
		fp5 = and(not(ns16_s), not(ns8_s), ns4_s, not(ns2_s), ns1_s)
		fp6 = and(not(ns16_s), not(ns8_s), ns4_s, ns2_s, not(ns1_s))
		fp7 = and(not(ns16_s), not(ns8_s), ns4_s, ns2_s, ns1_s)
		fp8 = and(not(ns16_s), ns8_s, not(ns4_s), not(ns2_s), not(ns1_s))
		fp9 = and(not(ns16_s), ns8_s, not(ns4_s), not(ns2_s), ns1_s)
		fp10 = and(not(ns16_s), ns8_s, not(ns4_s), ns2_s, not(ns1_s))
		fp11 = and(not(ns16_s), ns8_s, not(ns4_s), ns2_s, ns1_s)
		fp12 = and(not(ns16_s), ns8_s, ns4_s, not(ns2_s), not(ns1_s))
		fp13 = and(not(ns16_s), ns8_s, ns4_s, not(ns2_s), ns1_s)
		fp14 = and(not(ns16_s), ns8_s, ns4_s, ns2_s, not(ns1_s))
		fp15 = and(not(ns16_s), ns8_s, ns4_s, ns2_s, ns1_s)
		fp16 = and(ns16_s, not(ns8_s), not(ns4_s), not(ns2_s), not(ns1_s))
		fp17 = and(ns16_s, not(ns8_s), not(ns4_s), not(ns2_s), ns1_s)
		fp18 = and(ns16_s, not(ns8_s), not(ns4_s), ns2_s, not(ns1_s))
		fp19 = and(ns16_s, not(ns8_s), not(ns4_s), ns2_s, ns1_s)
		fp20 = and(ns16_s, not(ns8_s), ns4_s, not(ns2_s), not(ns1_s))
		fp21 = and(ns16_s, not(ns8_s), ns4_s, not(ns2_s), ns1_s)
		fp22 = and(ns16_s, not(ns8_s), ns4_s, ns2_s, not(ns1_s))
		fp23 = and(ns16_s, not(ns8_s), ns4_s, ns2_s, ns1_s)
		fp24 = and(ns16_s, ns8_s, not(ns4_s), not(ns2_s), not(ns1_s))
		fp25 = and(ns16_s, ns8_s, not(ns4_s), not(ns2_s), ns1_s)
		fp26 = and(ns16_s, ns8_s, not(ns4_s), ns2_s, not(ns1_s))
		fp27 = and(ns16_s, ns8_s, not(ns4_s), ns2_s, ns1_s)
		fp28 = and(ns16_s, ns8_s, ns4_s, not(ns2_s), not(ns1_s))
		fp29 = and(ns16_s, ns8_s, ns4_s, not(ns2_s), ns1_s)
		fp30 = and(ns16_s, ns8_s, ns4_s, ns2_s, not(ns1_s))
		fp31 = and(ns16_s, ns8_s, ns4_s, ns2_s, ns1_s)

		var collectedSIs S

		decompMap := func(fpx, out4_s, out2_s, out1_s nd, cSI S) S {

			//----------------------------------------------------------
			if and(fpx, s0i0) != null {
				cSI = append(cSI, g{s0i0, and(s0i0,
					out4_s), and(s0i0, out2_s), and(s0i0, out1_s), s0})
			}
			if and(fpx, s0i1) != null {
				cSI = append(cSI, g{s0i1, and(s0i1,
					out4_s), and(s0i1, out2_s), and(s0i1, out1_s), s1})
			}
			if and(fpx, s0i2) != null {
				cSI = append(cSI, g{s0i2, and(s0i2,
					out4_s), and(s0i2, out2_s), and(s0i2, out1_s), s0})
			}
			if and(fpx, s0i3) != null {
				cSI = append(cSI, g{s0i3, and(s0i3,
					out4_s), and(s0i3, out2_s), and(s0i3, out1_s), s0})
			}
			if and(fpx, s0i4) != null {
				cSI = append(cSI, g{s0i4, and(s0i4,
					out4_s), and(s0i4, out2_s), and(s0i4, out1_s), s0})
			}
			if and(fpx, s0i5) != null {
				cSI = append(cSI, g{s0i5, and(s0i5,
					out4_s), and(s0i5, out2_s), and(s0i5, out1_s), s0})
			}
			if and(fpx, s0i6) != null {
				cSI = append(cSI, g{s0i6, and(s0i6,
					out4_s), and(s0i6, out2_s), and(s0i6, out1_s), s0})
			}
			if and(fpx, s0i7) != null {
				cSI = append(cSI, g{s0i7, and(s0i7,
					out4_s), and(s0i7, out2_s), and(s0i7, out1_s), s0})
			}
			//----------------------------------------------------------
			if and(fpx, s1i0) != null {
				cSI = append(cSI, g{s1i0, and(s1i0,
					out4_s), and(s1i0, out2_s), and(s1i0, out1_s), s1})
			}
			if and(fpx, s1i1) != null {
				cSI = append(cSI, g{s1i1, and(s1i1,
					out4_s), and(s1i1, out2_s), and(s1i1, out1_s), s1})
			}
			if and(fpx, s1i2) != null {
				cSI = append(cSI, g{s1i2, and(s1i2,
					out4_s), and(s1i2, out2_s), and(s1i2, out1_s), s2})
			}
			if and(fpx, s1i3) != null {
				cSI = append(cSI, g{s1i3, and(s1i3,
					out4_s), and(s1i3, out2_s), and(s1i3, out1_s), s1})
			}
			if and(fpx, s1i4) != null {
				cSI = append(cSI, g{s1i4, and(s1i4,
					out4_s), and(s1i4, out2_s), and(s1i4, out1_s), s1})
			}
			if and(fpx, s1i5) != null {
				cSI = append(cSI, g{s1i5, and(s1i5,
					out4_s), and(s1i5, out2_s), and(s1i5, out1_s), s1})
			}
			if and(fpx, s1i6) != null {
				cSI = append(cSI, g{s1i6, and(s1i6,
					out4_s), and(s1i6, out2_s), and(s1i6, out1_s), s1})
			}
			if and(fpx, s1i7) != null {
				cSI = append(cSI, g{s1i7, and(s1i7,
					out4_s), and(s1i7, out2_s), and(s1i7, out1_s), s1})
			}
			//----------------------------------------------------------
			if and(fpx, s2i0) != null {
				cSI = append(cSI, g{s2i0, and(s2i0,
					out4_s), and(s2i0, out2_s), and(s2i0, out1_s), s2})
			}
			if and(fpx, s2i1) != null {
				cSI = append(cSI, g{s2i1, and(s2i1,
					out4_s), and(s2i1, out2_s), and(s2i1, out1_s), s2})
			}
			if and(fpx, s2i2) != null {
				cSI = append(cSI, g{s2i2, and(s2i2,
					out4_s), and(s2i2, out2_s), and(s2i2, out1_s), s3})
			}
			if and(fpx, s2i3) != null {
				cSI = append(cSI, g{s2i3, and(s2i3,
					out4_s), and(s2i3, out2_s), and(s2i3, out1_s), s2})
			}
			if and(fpx, s2i4) != null {
				cSI = append(cSI, g{s2i4, and(s2i4,
					out4_s), and(s2i4, out2_s), and(s2i4, out1_s), s2})
			}
			if and(fpx, s2i5) != null {
				cSI = append(cSI, g{s2i5, and(s2i5,
					out4_s), and(s2i5, out2_s), and(s2i5, out1_s), s2})
			}
			if and(fpx, s2i6) != null {
				cSI = append(cSI, g{s2i6, and(s2i6,
					out4_s), and(s2i6, out2_s), and(s2i6, out1_s), s2})
			}
			if and(fpx, s2i7) != null {
				cSI = append(cSI, g{s2i7, and(s2i7,
					out4_s), and(s2i7, out2_s), and(s2i7, out1_s), s2})
			}
			//----------------------------------------------------------
			if and(fpx, s3i0) != null {
				cSI = append(cSI, g{s3i0, and(s3i0,
					out4_s), and(s3i0, out2_s), and(s3i0, out1_s), s3})
			}
			if and(fpx, s3i1) != null {
				cSI = append(cSI, g{s3i1, and(s3i1,
					out4_s), and(s3i1, out2_s), and(s3i1, out1_s), s3})
			}
			if and(fpx, s3i2) != null {
				cSI = append(cSI, g{s3i2, and(s3i2,
					out4_s), and(s3i2, out2_s), and(s3i2, out1_s), s12})
			}
			if and(fpx, s3i3) != null {
				cSI = append(cSI, g{s3i3, and(s3i3,
					out4_s), and(s3i3, out2_s), and(s3i3, out1_s), s4})
			}
			if and(fpx, s3i4) != null {
				cSI = append(cSI, g{s3i4, and(s3i4,
					out4_s), and(s3i4, out2_s), and(s3i4, out1_s), s3})
			}
			if and(fpx, s3i5) != null {
				cSI = append(cSI, g{s3i5, and(s3i5,
					out4_s), and(s3i5, out2_s), and(s3i5, out1_s), s3})
			}
			if and(fpx, s3i6) != null {
				cSI = append(cSI, g{s3i6, and(s3i6,
					out4_s), and(s3i6, out2_s), and(s3i6, out1_s), s3})
			}
			if and(fpx, s3i7) != null {
				cSI = append(cSI, g{s3i7, and(s3i7,
					out4_s), and(s3i7, out2_s), and(s3i7, out1_s), s3})
			}
			//----------------------------------------------------------
			if and(fpx, s4i0) != null {
				cSI = append(cSI, g{s4i0, and(s4i0,
					out4_s), and(s4i0, out2_s), and(s4i0, out1_s), s4})
			}
			if and(fpx, s4i1) != null {
				cSI = append(cSI, g{s4i1, and(s4i1,
					out4_s), and(s4i1, out2_s), and(s4i1, out1_s), s4})
			}
			if and(fpx, s4i2) != null {
				cSI = append(cSI, g{s4i2, and(s4i2,
					out4_s), and(s4i2, out2_s), and(s4i2, out1_s), s4})
			}
			if and(fpx, s4i3) != null {
				cSI = append(cSI, g{s4i3, and(s4i3,
					out4_s), and(s4i3, out2_s), and(s4i3, out1_s), s4})
			}
			if and(fpx, s4i4) != null {
				cSI = append(cSI, g{s4i4, and(s4i4,
					out4_s), and(s4i4, out2_s), and(s4i4, out1_s), s5})
			}
			if and(fpx, s4i5) != null {
				cSI = append(cSI, g{s4i5, and(s4i5,
					out4_s), and(s4i5, out2_s), and(s4i5, out1_s), s5})
			}
			if and(fpx, s4i6) != null {
				cSI = append(cSI, g{s4i6, and(s4i6,
					out4_s), and(s4i6, out2_s), and(s4i6, out1_s), s4})
			}
			if and(fpx, s4i7) != null {
				cSI = append(cSI, g{s4i7, and(s4i7,
					out4_s), and(s4i7, out2_s), and(s4i7, out1_s), s4})
			}
			//----------------------------------------------------------
			if and(fpx, s5i0) != null {
				cSI = append(cSI, g{s5i0, and(s5i0,
					out4_s), and(s5i0, out2_s), and(s5i0, out1_s), s6})
			}
			if and(fpx, s5i1) != null {
				cSI = append(cSI, g{s5i1, and(s5i1,
					out4_s), and(s5i1, out2_s), and(s5i1, out1_s), s5})
			}
			if and(fpx, s5i2) != null {
				cSI = append(cSI, g{s5i2, and(s5i2,
					out4_s), and(s5i2, out2_s), and(s5i2, out1_s), s5})
			}
			if and(fpx, s5i3) != null {
				cSI = append(cSI, g{s5i3, and(s5i3,
					out4_s), and(s5i3, out2_s), and(s5i3, out1_s), s5})
			}
			if and(fpx, s5i4) != null {
				cSI = append(cSI, g{s5i4, and(s5i4,
					out4_s), and(s5i4, out2_s), and(s5i4, out1_s), s5})
			}
			if and(fpx, s5i5) != null {
				cSI = append(cSI, g{s5i5, and(s5i5,
					out4_s), and(s5i5, out2_s), and(s5i5, out1_s), s5})
			}
			if and(fpx, s5i6) != null {
				cSI = append(cSI, g{s5i6, and(s5i6,
					out4_s), and(s5i6, out2_s), and(s5i6, out1_s), s5})
			}
			if and(fpx, s5i7) != null {
				cSI = append(cSI, g{s5i7, and(s5i7,
					out4_s), and(s5i7, out2_s), and(s5i7, out1_s), s5})
			}
			//----------------------------------------------------------
			if and(fpx, s6i0) != null {
				cSI = append(cSI, g{s6i0, and(s6i0,
					out4_s), and(s6i0, out2_s), and(s6i0, out1_s), s6})
			}
			if and(fpx, s6i1) != null {
				cSI = append(cSI, g{s6i1, and(s6i1,
					out4_s), and(s6i1, out2_s), and(s6i1, out1_s), s6})
			}
			if and(fpx, s6i2) != null {
				cSI = append(cSI, g{s6i2, and(s6i2,
					out4_s), and(s6i2, out2_s), and(s6i2, out1_s), s6})
			}
			if and(fpx, s6i3) != null {
				cSI = append(cSI, g{s6i3, and(s6i3,
					out4_s), and(s6i3, out2_s), and(s6i3, out1_s), s6})
			}
			if and(fpx, s6i4) != null {
				cSI = append(cSI, g{s6i4, and(s6i4,
					out4_s), and(s6i4, out2_s), and(s6i4, out1_s), s6})
			}
			if and(fpx, s6i5) != null {
				cSI = append(cSI, g{s6i5, and(s6i5,
					out4_s), and(s6i5, out2_s), and(s6i5, out1_s), s6})
			}
			if and(fpx, s6i6) != null {
				cSI = append(cSI, g{s6i6, and(s6i6,
					out4_s), and(s6i6, out2_s), and(s6i6, out1_s), s6})
			}
			if and(fpx, s6i7) != null {
				cSI = append(cSI, g{s6i7, and(s6i7,
					out4_s), and(s6i7, out2_s), and(s6i7, out1_s), s7})
			}
			//----------------------------------------------------------
			if and(fpx, s7i0) != null {
				cSI = append(cSI, g{s7i0, and(s7i0,
					out4_s), and(s7i0, out2_s), and(s7i0, out1_s), s7})
			}
			if and(fpx, s7i1) != null {
				cSI = append(cSI, g{s7i1, and(s7i1,
					out4_s), and(s7i1, out2_s), and(s7i1, out1_s), s7})
			}
			if and(fpx, s7i2) != null {
				cSI = append(cSI, g{s7i2, and(s7i2,
					out4_s), and(s7i2, out2_s), and(s7i2, out1_s), s8})
			}
			if and(fpx, s7i3) != null {
				cSI = append(cSI, g{s7i3, and(s7i3,
					out4_s), and(s7i3, out2_s), and(s7i3, out1_s), s8})
			}
			if and(fpx, s7i4) != null {
				cSI = append(cSI, g{s7i4, and(s7i4,
					out4_s), and(s7i4, out2_s), and(s7i4, out1_s), s7})
			}
			if and(fpx, s7i5) != null {
				cSI = append(cSI, g{s7i5, and(s7i5,
					out4_s), and(s7i5, out2_s), and(s7i5, out1_s), s8})
			}
			if and(fpx, s7i6) != null {
				cSI = append(cSI, g{s7i6, and(s7i6,
					out4_s), and(s7i6, out2_s), and(s7i6, out1_s), s7})
			}
			if and(fpx, s7i7) != null {
				cSI = append(cSI, g{s7i7, and(s7i7,
					out4_s), and(s7i7, out2_s), and(s7i7, out1_s), s7})
			}
			//----------------------------------------------------------
			if and(fpx, s8i0) != null {
				cSI = append(cSI, g{s8i0, and(s8i0,
					out4_s), and(s8i0, out2_s), and(s8i0, out1_s), s8})
			}
			if and(fpx, s8i1) != null {
				cSI = append(cSI, g{s8i1, and(s8i1,
					out4_s), and(s8i1, out2_s), and(s8i1, out1_s), s9})
			}
			if and(fpx, s8i2) != null {
				cSI = append(cSI, g{s8i2, and(s8i2,
					out4_s), and(s8i2, out2_s), and(s8i2, out1_s), s8})
			}
			if and(fpx, s8i3) != null {
				cSI = append(cSI, g{s8i3, and(s8i3,
					out4_s), and(s8i3, out2_s), and(s8i3, out1_s), s8})
			}
			if and(fpx, s8i4) != null {
				cSI = append(cSI, g{s8i4, and(s8i4,
					out4_s), and(s8i4, out2_s), and(s8i4, out1_s), s8})
			}
			if and(fpx, s8i5) != null {
				cSI = append(cSI, g{s8i5, and(s8i5,
					out4_s), and(s8i5, out2_s), and(s8i5, out1_s), s8})
			}
			if and(fpx, s8i6) != null {
				cSI = append(cSI, g{s8i6, and(s8i6,
					out4_s), and(s8i6, out2_s), and(s8i6, out1_s), s8})
			}
			if and(fpx, s8i7) != null {
				cSI = append(cSI, g{s8i7, and(s8i7,
					out4_s), and(s8i7, out2_s), and(s8i7, out1_s), s8})
			}
			//----------------------------------------------------------
			if and(fpx, s9i0) != null {
				cSI = append(cSI, g{s9i0, and(s9i0,
					out4_s), and(s9i0, out2_s), and(s9i0, out1_s), s9})
			}
			if and(fpx, s9i1) != null {
				cSI = append(cSI, g{s9i1, and(s9i1,
					out4_s), and(s9i1, out2_s), and(s9i1, out1_s), s9})
			}
			if and(fpx, s9i2) != null {
				cSI = append(cSI, g{s9i2, and(s9i2,
					out4_s), and(s9i2, out2_s), and(s9i2, out1_s), s10})
			}
			if and(fpx, s9i3) != null {
				cSI = append(cSI, g{s9i3, and(s9i3,
					out4_s), and(s9i3, out2_s), and(s9i3, out1_s), s9})
			}
			if and(fpx, s9i4) != null {
				cSI = append(cSI, g{s9i4, and(s9i4,
					out4_s), and(s9i4, out2_s), and(s9i4, out1_s), s9})
			}
			if and(fpx, s9i5) != null {
				cSI = append(cSI, g{s9i5, and(s9i5,
					out4_s), and(s9i5, out2_s), and(s9i5, out1_s), s9})
			}
			if and(fpx, s9i6) != null {
				cSI = append(cSI, g{s9i6, and(s9i6,
					out4_s), and(s9i6, out2_s), and(s9i6, out1_s), s9})
			}
			if and(fpx, s9i7) != null {
				cSI = append(cSI, g{s9i7, and(s9i7,
					out4_s), and(s9i7, out2_s), and(s9i7, out1_s), s9})
			}
			//----------------------------------------------------------
			if and(fpx, s10i0) != null {
				cSI = append(cSI, g{s10i0, and(s10i0,
					out4_s), and(s10i0, out2_s), and(s10i0, out1_s), s27})
			}
			if and(fpx, s10i1) != null {
				cSI = append(cSI, g{s10i1, and(s10i1,
					out4_s), and(s10i1, out2_s), and(s10i1, out1_s), s10})
			}
			if and(fpx, s10i2) != null {
				cSI = append(cSI, g{s10i2, and(s10i2,
					out4_s), and(s10i2, out2_s), and(s10i2, out1_s), s11})
			}
			if and(fpx, s10i3) != null {
				cSI = append(cSI, g{s10i3, and(s10i3,
					out4_s), and(s10i3, out2_s), and(s10i3, out1_s), s10})
			}
			if and(fpx, s10i4) != null {
				cSI = append(cSI, g{s10i4, and(s10i4,
					out4_s), and(s10i4, out2_s), and(s10i4, out1_s), s10})
			}
			if and(fpx, s10i5) != null {
				cSI = append(cSI, g{s10i5, and(s10i5,
					out4_s), and(s10i5, out2_s), and(s10i5, out1_s), s10})
			}
			if and(fpx, s10i6) != null {
				cSI = append(cSI, g{s10i6, and(s10i6,
					out4_s), and(s10i6, out2_s), and(s10i6, out1_s), s3})
			}
			if and(fpx, s10i7) != null {
				cSI = append(cSI, g{s10i7, and(s10i7,
					out4_s), and(s10i7, out2_s), and(s10i7, out1_s), s10})
			}
			//----------------------------------------------------------
			if and(fpx, s11i0) != null {
				cSI = append(cSI, g{s11i0, and(s11i0,
					out4_s), and(s11i0, out2_s), and(s11i0, out1_s), s11})
			}
			if and(fpx, s11i1) != null {
				cSI = append(cSI, g{s11i1, and(s11i1,
					out4_s), and(s11i1, out2_s), and(s11i1, out1_s), s11})
			}
			if and(fpx, s11i2) != null {
				cSI = append(cSI, g{s11i2, and(s11i2,
					out4_s), and(s11i2, out2_s), and(s11i2, out1_s), s11})
			}
			if and(fpx, s11i3) != null {
				cSI = append(cSI, g{s11i3, and(s11i3,
					out4_s), and(s11i3, out2_s), and(s11i3, out1_s), s11})
			}
			if and(fpx, s11i4) != null {
				cSI = append(cSI, g{s11i4, and(s11i4,
					out4_s), and(s11i4, out2_s), and(s11i4, out1_s), s11})
			}
			if and(fpx, s11i5) != null {
				cSI = append(cSI, g{s11i5, and(s11i5,
					out4_s), and(s11i5, out2_s), and(s11i5, out1_s), s11})
			}
			if and(fpx, s11i6) != null {
				cSI = append(cSI, g{s11i6, and(s11i6,
					out4_s), and(s11i6, out2_s), and(s11i6, out1_s), s11})
			}
			if and(fpx, s11i7) != null {
				cSI = append(cSI, g{s11i7, and(s11i7,
					out4_s), and(s11i7, out2_s), and(s11i7, out1_s), s12})
			}
			//----------------------------------------------------------
			if and(fpx, s12i0) != null {
				cSI = append(cSI, g{s12i0, and(s12i0,
					out4_s), and(s12i0, out2_s), and(s12i0, out1_s), s12})
			}
			if and(fpx, s12i1) != null {
				cSI = append(cSI, g{s12i1, and(s12i1,
					out4_s), and(s12i1, out2_s), and(s12i1, out1_s), s12})
			}
			if and(fpx, s12i2) != null {
				cSI = append(cSI, g{s12i2, and(s12i2,
					out4_s), and(s12i2, out2_s), and(s12i2, out1_s), s12})
			}
			if and(fpx, s12i3) != null {
				cSI = append(cSI, g{s12i3, and(s12i3,
					out4_s), and(s12i3, out2_s), and(s12i3, out1_s), s13})
			}
			if and(fpx, s12i4) != null {
				cSI = append(cSI, g{s12i4, and(s12i4,
					out4_s), and(s12i4, out2_s), and(s12i4, out1_s), s12})
			}
			if and(fpx, s12i5) != null {
				cSI = append(cSI, g{s12i5, and(s12i5,
					out4_s), and(s12i5, out2_s), and(s12i5, out1_s), s12})
			}
			if and(fpx, s12i6) != null {
				cSI = append(cSI, g{s12i6, and(s12i6,
					out4_s), and(s12i6, out2_s), and(s12i6, out1_s), s13})
			}
			if and(fpx, s12i7) != null {
				cSI = append(cSI, g{s12i7, and(s12i7,
					out4_s), and(s12i7, out2_s), and(s12i7, out1_s), s12})
			}
			//----------------------------------------------------------
			if and(fpx, s13i0) != null {
				cSI = append(cSI, g{s13i0, and(s13i0,
					out4_s), and(s13i0, out2_s), and(s13i0, out1_s), s14})
			}
			if and(fpx, s13i1) != null {
				cSI = append(cSI, g{s13i1, and(s13i1,
					out4_s), and(s13i1, out2_s), and(s13i1, out1_s), s13})
			}
			if and(fpx, s13i2) != null {
				cSI = append(cSI, g{s13i2, and(s13i2,
					out4_s), and(s13i2, out2_s), and(s13i2, out1_s), s13})
			}
			if and(fpx, s13i3) != null {
				cSI = append(cSI, g{s13i3, and(s13i3,
					out4_s), and(s13i3, out2_s), and(s13i3, out1_s), s13})
			}
			if and(fpx, s13i4) != null {
				cSI = append(cSI, g{s13i4, and(s13i4,
					out4_s), and(s13i4, out2_s), and(s13i4, out1_s), s13})
			}
			if and(fpx, s13i5) != null {
				cSI = append(cSI, g{s13i5, and(s13i5,
					out4_s), and(s13i5, out2_s), and(s13i5, out1_s), s13})
			}
			if and(fpx, s13i6) != null {
				cSI = append(cSI, g{s13i6, and(s13i6,
					out4_s), and(s13i6, out2_s), and(s13i6, out1_s), s13})
			}
			if and(fpx, s13i7) != null {
				cSI = append(cSI, g{s13i7, and(s13i7,
					out4_s), and(s13i7, out2_s), and(s13i7, out1_s), s13})
			}
			//----------------------------------------------------------
			if and(fpx, s14i0) != null {
				cSI = append(cSI, g{s14i0, and(s14i0,
					out4_s), and(s14i0, out2_s), and(s14i0, out1_s), s14})
			}
			if and(fpx, s14i1) != null {
				cSI = append(cSI, g{s14i1, and(s14i1,
					out4_s), and(s14i1, out2_s), and(s14i1, out1_s), s14})
			}
			if and(fpx, s14i2) != null {
				cSI = append(cSI, g{s14i2, and(s14i2,
					out4_s), and(s14i2, out2_s), and(s14i2, out1_s), s14})
			}
			if and(fpx, s14i3) != null {
				cSI = append(cSI, g{s14i3, and(s14i3,
					out4_s), and(s14i3, out2_s), and(s14i3, out1_s), s14})
			}
			if and(fpx, s14i4) != null {
				cSI = append(cSI, g{s14i4, and(s14i4,
					out4_s), and(s14i4, out2_s), and(s14i4, out1_s), s14})
			}
			if and(fpx, s14i5) != null {
				cSI = append(cSI, g{s14i5, and(s14i5,
					out4_s), and(s14i5, out2_s), and(s14i5, out1_s), s14})
			}
			if and(fpx, s14i6) != null {
				cSI = append(cSI, g{s14i6, and(s14i6,
					out4_s), and(s14i6, out2_s), and(s14i6, out1_s), s14})
			}
			if and(fpx, s14i7) != null {
				cSI = append(cSI, g{s14i7, and(s14i7,
					out4_s), and(s14i7, out2_s), and(s14i7, out1_s), s15})
			}
			//----------------------------------------------------------
			if and(fpx, s15i0) != null {
				cSI = append(cSI, g{s15i0, and(s15i0,
					out4_s), and(s15i0, out2_s), and(s15i0, out1_s), s15})
			}
			if and(fpx, s15i1) != null {
				cSI = append(cSI, g{s15i1, and(s15i1,
					out4_s), and(s15i1, out2_s), and(s15i1, out1_s), s15})
			}
			if and(fpx, s15i2) != null {
				cSI = append(cSI, g{s15i2, and(s15i2,
					out4_s), and(s15i2, out2_s), and(s15i2, out1_s), s15})
			}
			if and(fpx, s15i3) != null {
				cSI = append(cSI, g{s15i3, and(s15i3,
					out4_s), and(s15i3, out2_s), and(s15i3, out1_s), s15})
			}
			if and(fpx, s15i4) != null {
				cSI = append(cSI, g{s15i4, and(s15i4,
					out4_s), and(s15i4, out2_s), and(s15i4, out1_s), s15})
			}
			if and(fpx, s15i5) != null {
				cSI = append(cSI, g{s15i5, and(s15i5,
					out4_s), and(s15i5, out2_s), and(s15i5, out1_s), s16})
			}
			if and(fpx, s15i6) != null {
				cSI = append(cSI, g{s15i6, and(s15i6,
					out4_s), and(s15i6, out2_s), and(s15i6, out1_s), s15})
			}
			if and(fpx, s15i7) != null {
				cSI = append(cSI, g{s15i7, and(s15i7,
					out4_s), and(s15i7, out2_s), and(s15i7, out1_s), s15})
			}
			//----------------------------------------------------------
			if and(fpx, s16i0) != null {
				cSI = append(cSI, g{s16i0, and(s16i0,
					out4_s), and(s16i0, out2_s), and(s16i0, out1_s), s16})
			}
			if and(fpx, s16i1) != null {
				cSI = append(cSI, g{s16i1, and(s16i1,
					out4_s), and(s16i1, out2_s), and(s16i1, out1_s), s17})
			}
			if and(fpx, s16i2) != null {
				cSI = append(cSI, g{s16i2, and(s16i2,
					out4_s), and(s16i2, out2_s), and(s16i2, out1_s), s16})
			}
			if and(fpx, s16i3) != null {
				cSI = append(cSI, g{s16i3, and(s16i3,
					out4_s), and(s16i3, out2_s), and(s16i3, out1_s), s16})
			}
			if and(fpx, s16i4) != null {
				cSI = append(cSI, g{s16i4, and(s16i4,
					out4_s), and(s16i4, out2_s), and(s16i4, out1_s), s16})
			}
			if and(fpx, s16i5) != null {
				cSI = append(cSI, g{s16i5, and(s16i5,
					out4_s), and(s16i5, out2_s), and(s16i5, out1_s), s16})
			}
			if and(fpx, s16i6) != null {
				cSI = append(cSI, g{s16i6, and(s16i6,
					out4_s), and(s16i6, out2_s), and(s16i6, out1_s), s16})
			}
			if and(fpx, s16i7) != null {
				cSI = append(cSI, g{s16i7, and(s16i7,
					out4_s), and(s16i7, out2_s), and(s16i7, out1_s), s16})
			}
			//----------------------------------------------------------
			if and(fpx, s17i0) != null {
				cSI = append(cSI, g{s17i0, and(s17i0,
					out4_s), and(s17i0, out2_s), and(s17i0, out1_s), s17})
			}
			if and(fpx, s17i1) != null {
				cSI = append(cSI, g{s17i1, and(s17i1,
					out4_s), and(s17i1, out2_s), and(s17i1, out1_s), s17})
			}
			if and(fpx, s17i2) != null {
				cSI = append(cSI, g{s17i2, and(s17i2,
					out4_s), and(s17i2, out2_s), and(s17i2, out1_s), s18})
			}
			if and(fpx, s17i3) != null {
				cSI = append(cSI, g{s17i3, and(s17i3,
					out4_s), and(s17i3, out2_s), and(s17i3, out1_s), s17})
			}
			if and(fpx, s17i4) != null {
				cSI = append(cSI, g{s17i4, and(s17i4,
					out4_s), and(s17i4, out2_s), and(s17i4, out1_s), s17})
			}
			if and(fpx, s17i5) != null {
				cSI = append(cSI, g{s17i5, and(s17i5,
					out4_s), and(s17i5, out2_s), and(s17i5, out1_s), s17})
			}
			if and(fpx, s17i6) != null {
				cSI = append(cSI, g{s17i6, and(s17i6,
					out4_s), and(s17i6, out2_s), and(s17i6, out1_s), s17})
			}
			if and(fpx, s17i7) != null {
				cSI = append(cSI, g{s17i7, and(s17i7,
					out4_s), and(s17i7, out2_s), and(s17i7, out1_s), s17})
			}
			//----------------------------------------------------------
			if and(fpx, s18i0) != null {
				cSI = append(cSI, g{s18i0, and(s18i0,
					out4_s), and(s18i0, out2_s), and(s18i0, out1_s), s18})
			}
			if and(fpx, s18i1) != null {
				cSI = append(cSI, g{s18i1, and(s18i1,
					out4_s), and(s18i1, out2_s), and(s18i1, out1_s), s18})
			}
			if and(fpx, s18i2) != null {
				cSI = append(cSI, g{s18i2, and(s18i2,
					out4_s), and(s18i2, out2_s), and(s18i2, out1_s), s19})
			}
			if and(fpx, s18i3) != null {
				cSI = append(cSI, g{s18i3, and(s18i3,
					out4_s), and(s18i3, out2_s), and(s18i3, out1_s), s18})
			}
			if and(fpx, s18i4) != null {
				cSI = append(cSI, g{s18i4, and(s18i4,
					out4_s), and(s18i4, out2_s), and(s18i4, out1_s), s18})
			}
			if and(fpx, s18i5) != null {
				cSI = append(cSI, g{s18i5, and(s18i5,
					out4_s), and(s18i5, out2_s), and(s18i5, out1_s), s18})
			}
			if and(fpx, s18i6) != null {
				cSI = append(cSI, g{s18i6, and(s18i6,
					out4_s), and(s18i6, out2_s), and(s18i6, out1_s), s7})
			}
			if and(fpx, s18i7) != null {
				cSI = append(cSI, g{s18i7, and(s18i7,
					out4_s), and(s18i7, out2_s), and(s18i7, out1_s), s18})
			}
			//----------------------------------------------------------
			if and(fpx, s19i0) != null {
				cSI = append(cSI, g{s19i0, and(s19i0,
					out4_s), and(s19i0, out2_s), and(s19i0, out1_s), s19})
			}
			if and(fpx, s19i1) != null {
				cSI = append(cSI, g{s19i1, and(s19i1,
					out4_s), and(s19i1, out2_s), and(s19i1, out1_s), s19})
			}
			if and(fpx, s19i2) != null {
				cSI = append(cSI, g{s19i2, and(s19i2,
					out4_s), and(s19i2, out2_s), and(s19i2, out1_s), s19})
			}
			if and(fpx, s19i3) != null {
				cSI = append(cSI, g{s19i3, and(s19i3,
					out4_s), and(s19i3, out2_s), and(s19i3, out1_s), s19})
			}
			if and(fpx, s19i4) != null {
				cSI = append(cSI, g{s19i4, and(s19i4,
					out4_s), and(s19i4, out2_s), and(s19i4, out1_s), s19})
			}
			if and(fpx, s19i5) != null {
				cSI = append(cSI, g{s19i5, and(s19i5,
					out4_s), and(s19i5, out2_s), and(s19i5, out1_s), s23})
			}
			if and(fpx, s19i6) != null {
				cSI = append(cSI, g{s19i6, and(s19i6,
					out4_s), and(s19i6, out2_s), and(s19i6, out1_s), s19})
			}
			if and(fpx, s19i7) != null {
				cSI = append(cSI, g{s19i7, and(s19i7,
					out4_s), and(s19i7, out2_s), and(s19i7, out1_s), s20})
			}
			//----------------------------------------------------------
			if and(fpx, s20i0) != null {
				cSI = append(cSI, g{s20i0, and(s20i0,
					out4_s), and(s20i0, out2_s), and(s20i0, out1_s), s20})
			}
			if and(fpx, s20i1) != null {
				cSI = append(cSI, g{s20i1, and(s20i1,
					out4_s), and(s20i1, out2_s), and(s20i1, out1_s), s20})
			}
			if and(fpx, s20i2) != null {
				cSI = append(cSI, g{s20i2, and(s20i2,
					out4_s), and(s20i2, out2_s), and(s20i2, out1_s), s20})
			}
			if and(fpx, s20i3) != null {
				cSI = append(cSI, g{s20i3, and(s20i3,
					out4_s), and(s20i3, out2_s), and(s20i3, out1_s), s21})
			}
			if and(fpx, s20i4) != null {
				cSI = append(cSI, g{s20i4, and(s20i4,
					out4_s), and(s20i4, out2_s), and(s20i4, out1_s), s20})
			}
			if and(fpx, s20i5) != null {
				cSI = append(cSI, g{s20i5, and(s20i5,
					out4_s), and(s20i5, out2_s), and(s20i5, out1_s), s20})
			}
			if and(fpx, s20i6) != null {
				cSI = append(cSI, g{s20i6, and(s20i6,
					out4_s), and(s20i6, out2_s), and(s20i6, out1_s), s21})
			}
			if and(fpx, s20i7) != null {
				cSI = append(cSI, g{s20i7, and(s20i7,
					out4_s), and(s20i7, out2_s), and(s20i7, out1_s), s20})
			}
			//----------------------------------------------------------
			if and(fpx, s21i0) != null {
				cSI = append(cSI, g{s21i0, and(s21i0,
					out4_s), and(s21i0, out2_s), and(s21i0, out1_s), s22})
			}
			if and(fpx, s21i1) != null {
				cSI = append(cSI, g{s21i1, and(s21i1,
					out4_s), and(s21i1, out2_s), and(s21i1, out1_s), s21})
			}
			if and(fpx, s21i2) != null {
				cSI = append(cSI, g{s21i2, and(s21i2,
					out4_s), and(s21i2, out2_s), and(s21i2, out1_s), s21})
			}
			if and(fpx, s21i3) != null {
				cSI = append(cSI, g{s21i3, and(s21i3,
					out4_s), and(s21i3, out2_s), and(s21i3, out1_s), s21})
			}
			if and(fpx, s21i4) != null {
				cSI = append(cSI, g{s21i4, and(s21i4,
					out4_s), and(s21i4, out2_s), and(s21i4, out1_s), s21})
			}
			if and(fpx, s21i5) != null {
				cSI = append(cSI, g{s21i5, and(s21i5,
					out4_s), and(s21i5, out2_s), and(s21i5, out1_s), s21})
			}
			if and(fpx, s21i6) != null {
				cSI = append(cSI, g{s21i6, and(s21i6,
					out4_s), and(s21i6, out2_s), and(s21i6, out1_s), s21})
			}
			if and(fpx, s21i7) != null {
				cSI = append(cSI, g{s21i7, and(s21i7,
					out4_s), and(s21i7, out2_s), and(s21i7, out1_s), s21})
			}
			//----------------------------------------------------------
			if and(fpx, s22i0) != null {
				cSI = append(cSI, g{s22i0, and(s22i0,
					out4_s), and(s22i0, out2_s), and(s22i0, out1_s), s22})
			}
			if and(fpx, s22i1) != null {
				cSI = append(cSI, g{s22i1, and(s22i1,
					out4_s), and(s22i1, out2_s), and(s22i1, out1_s), s22})
			}
			if and(fpx, s22i2) != null {
				cSI = append(cSI, g{s22i2, and(s22i2,
					out4_s), and(s22i2, out2_s), and(s22i2, out1_s), s22})
			}
			if and(fpx, s22i3) != null {
				cSI = append(cSI, g{s22i3, and(s22i3,
					out4_s), and(s22i3, out2_s), and(s22i3, out1_s), s22})
			}
			if and(fpx, s22i4) != null {
				cSI = append(cSI, g{s22i4, and(s22i4,
					out4_s), and(s22i4, out2_s), and(s22i4, out1_s), s22})
			}
			if and(fpx, s22i5) != null {
				cSI = append(cSI, g{s22i5, and(s22i5,
					out4_s), and(s22i5, out2_s), and(s22i5, out1_s), s22})
			}
			if and(fpx, s22i6) != null {
				cSI = append(cSI, g{s22i6, and(s22i6,
					out4_s), and(s22i6, out2_s), and(s22i6, out1_s), s22})
			}
			if and(fpx, s22i7) != null {
				cSI = append(cSI, g{s22i7, and(s22i7,
					out4_s), and(s22i7, out2_s), and(s22i7, out1_s), s23})
			}
			//----------------------------------------------------------
			if and(fpx, s23i0) != null {
				cSI = append(cSI, g{s23i0, and(s23i0,
					out4_s), and(s23i0, out2_s), and(s23i0, out1_s), s23})
			}
			if and(fpx, s23i1) != null {
				cSI = append(cSI, g{s23i1, and(s23i1,
					out4_s), and(s23i1, out2_s), and(s23i1, out1_s), s29})
			}
			if and(fpx, s23i2) != null {
				cSI = append(cSI, g{s23i2, and(s23i2,
					out4_s), and(s23i2, out2_s), and(s23i2, out1_s), s23})
			}
			if and(fpx, s23i3) != null {
				cSI = append(cSI, g{s23i3, and(s23i3,
					out4_s), and(s23i3, out2_s), and(s23i3, out1_s), s23})
			}
			if and(fpx, s23i4) != null {
				cSI = append(cSI, g{s23i4, and(s23i4,
					out4_s), and(s23i4, out2_s), and(s23i4, out1_s), s23})
			}
			if and(fpx, s23i5) != null {
				cSI = append(cSI, g{s23i5, and(s23i5,
					out4_s), and(s23i5, out2_s), and(s23i5, out1_s), s24})
			}
			if and(fpx, s23i6) != null {
				cSI = append(cSI, g{s23i6, and(s23i6,
					out4_s), and(s23i6, out2_s), and(s23i6, out1_s), s23})
			}
			if and(fpx, s23i7) != null {
				cSI = append(cSI, g{s23i7, and(s23i7,
					out4_s), and(s23i7, out2_s), and(s23i7, out1_s), s23})
			}
			//----------------------------------------------------------
			if and(fpx, s24i0) != null {
				cSI = append(cSI, g{s24i0, and(s24i0,
					out4_s), and(s24i0, out2_s), and(s24i0, out1_s), s24})
			}
			if and(fpx, s24i1) != null {
				cSI = append(cSI, g{s24i1, and(s24i1,
					out4_s), and(s24i1, out2_s), and(s24i1, out1_s), s25})
			}
			if and(fpx, s24i2) != null {
				cSI = append(cSI, g{s24i2, and(s24i2,
					out4_s), and(s24i2, out2_s), and(s24i2, out1_s), s24})
			}
			if and(fpx, s24i3) != null {
				cSI = append(cSI, g{s24i3, and(s24i3,
					out4_s), and(s24i3, out2_s), and(s24i3, out1_s), s24})
			}
			if and(fpx, s24i4) != null {
				cSI = append(cSI, g{s24i4, and(s24i4,
					out4_s), and(s24i4, out2_s), and(s24i4, out1_s), s24})
			}
			if and(fpx, s24i5) != null {
				cSI = append(cSI, g{s24i5, and(s24i5,
					out4_s), and(s24i5, out2_s), and(s24i5, out1_s), s24})
			}
			if and(fpx, s24i6) != null {
				cSI = append(cSI, g{s24i6, and(s24i6,
					out4_s), and(s24i6, out2_s), and(s24i6, out1_s), s24})
			}
			if and(fpx, s24i7) != null {
				cSI = append(cSI, g{s24i7, and(s24i7,
					out4_s), and(s24i7, out2_s), and(s24i7, out1_s), s9})
			}
			//----------------------------------------------------------
			if and(fpx, s25i0) != null {
				cSI = append(cSI, g{s25i0, and(s25i0,
					out4_s), and(s25i0, out2_s), and(s25i0, out1_s), s25})
			}
			if and(fpx, s25i1) != null {
				cSI = append(cSI, g{s25i1, and(s25i1,
					out4_s), and(s25i1, out2_s), and(s25i1, out1_s), s25})
			}
			if and(fpx, s25i2) != null {
				cSI = append(cSI, g{s25i2, and(s25i2,
					out4_s), and(s25i2, out2_s), and(s25i2, out1_s), s26})
			}
			if and(fpx, s25i3) != null {
				cSI = append(cSI, g{s25i3, and(s25i3,
					out4_s), and(s25i3, out2_s), and(s25i3, out1_s), s25})
			}
			if and(fpx, s25i4) != null {
				cSI = append(cSI, g{s25i4, and(s25i4,
					out4_s), and(s25i4, out2_s), and(s25i4, out1_s), s25})
			}
			if and(fpx, s25i5) != null {
				cSI = append(cSI, g{s25i5, and(s25i5,
					out4_s), and(s25i5, out2_s), and(s25i5, out1_s), s25})
			}
			if and(fpx, s25i6) != null {
				cSI = append(cSI, g{s25i6, and(s25i6,
					out4_s), and(s25i6, out2_s), and(s25i6, out1_s), s25})
			}
			if and(fpx, s25i7) != null {
				cSI = append(cSI, g{s25i7, and(s25i7,
					out4_s), and(s25i7, out2_s), and(s25i7, out1_s), s25})
			}
			//----------------------------------------------------------
			if and(fpx, s26i0) != null {
				cSI = append(cSI, g{s26i0, and(s26i0,
					out4_s), and(s26i0, out2_s), and(s26i0, out1_s), s26})
			}
			if and(fpx, s26i1) != null {
				cSI = append(cSI, g{s26i1, and(s26i1,
					out4_s), and(s26i1, out2_s), and(s26i1, out1_s), s26})
			}
			if and(fpx, s26i2) != null {
				cSI = append(cSI, g{s26i2, and(s26i2,
					out4_s), and(s26i2, out2_s), and(s26i2, out1_s), s27})
			}
			if and(fpx, s26i3) != null {
				cSI = append(cSI, g{s26i3, and(s26i3,
					out4_s), and(s26i3, out2_s), and(s26i3, out1_s), s26})
			}
			if and(fpx, s26i4) != null {
				cSI = append(cSI, g{s26i4, and(s26i4,
					out4_s), and(s26i4, out2_s), and(s26i4, out1_s), s26})
			}
			if and(fpx, s26i5) != null {
				cSI = append(cSI, g{s26i5, and(s26i5,
					out4_s), and(s26i5, out2_s), and(s26i5, out1_s), s26})
			}
			if and(fpx, s26i6) != null {
				cSI = append(cSI, g{s26i6, and(s26i6,
					out4_s), and(s26i6, out2_s), and(s26i6, out1_s), s26})
			}
			if and(fpx, s26i7) != null {
				cSI = append(cSI, g{s26i7, and(s26i7,
					out4_s), and(s26i7, out2_s), and(s26i7, out1_s), s26})
			}
			//----------------------------------------------------------
			if and(fpx, s27i0) != null {
				cSI = append(cSI, g{s27i0, and(s27i0,
					out4_s), and(s27i0, out2_s), and(s27i0, out1_s), s27})
			}
			if and(fpx, s27i1) != null {
				cSI = append(cSI, g{s27i1, and(s27i1,
					out4_s), and(s27i1, out2_s), and(s27i1, out1_s), s27})
			}
			if and(fpx, s27i2) != null {
				cSI = append(cSI, g{s27i2, and(s27i2,
					out4_s), and(s27i2, out2_s), and(s27i2, out1_s), s27})
			}
			if and(fpx, s27i3) != null {
				cSI = append(cSI, g{s27i3, and(s27i3,
					out4_s), and(s27i3, out2_s), and(s27i3, out1_s), s27})
			}
			if and(fpx, s27i4) != null {
				cSI = append(cSI, g{s27i4, and(s27i4,
					out4_s), and(s27i4, out2_s), and(s27i4, out1_s), s27})
			}
			if and(fpx, s27i5) != null {
				cSI = append(cSI, g{s27i5, and(s27i5,
					out4_s), and(s27i5, out2_s), and(s27i5, out1_s), s27})
			}
			if and(fpx, s27i6) != null {
				cSI = append(cSI, g{s27i6, and(s27i6,
					out4_s), and(s27i6, out2_s), and(s27i6, out1_s), s27})
			}
			if and(fpx, s27i7) != null {
				cSI = append(cSI, g{s27i7, and(s27i7,
					out4_s), and(s27i7, out2_s), and(s27i7, out1_s), s28})
			}
			//----------------------------------------------------------
			if and(fpx, s28i0) != null {
				cSI = append(cSI, g{s28i0, and(s28i0,
					out4_s), and(s28i0, out2_s), and(s28i0, out1_s), s28})
			}
			if and(fpx, s28i1) != null {
				cSI = append(cSI, g{s28i1, and(s28i1,
					out4_s), and(s28i1, out2_s), and(s28i1, out1_s), s28})
			}
			if and(fpx, s28i2) != null {
				cSI = append(cSI, g{s28i2, and(s28i2,
					out4_s), and(s28i2, out2_s), and(s28i2, out1_s), s28})
			}
			if and(fpx, s28i3) != null {
				cSI = append(cSI, g{s28i3, and(s28i3,
					out4_s), and(s28i3, out2_s), and(s28i3, out1_s), s29})
			}
			if and(fpx, s28i4) != null {
				cSI = append(cSI, g{s28i4, and(s28i4,
					out4_s), and(s28i4, out2_s), and(s28i4, out1_s), s28})
			}
			if and(fpx, s28i5) != null {
				cSI = append(cSI, g{s28i5, and(s28i5,
					out4_s), and(s28i5, out2_s), and(s28i5, out1_s), s28})
			}
			if and(fpx, s28i6) != null {
				cSI = append(cSI, g{s28i6, and(s28i6,
					out4_s), and(s28i6, out2_s), and(s28i6, out1_s), s29})
			}
			if and(fpx, s28i7) != null {
				cSI = append(cSI, g{s28i7, and(s28i7,
					out4_s), and(s28i7, out2_s), and(s28i7, out1_s), s28})
			}
			//----------------------------------------------------------
			if and(fpx, s29i0) != null {
				cSI = append(cSI, g{s29i0, and(s29i0,
					out4_s), and(s29i0, out2_s), and(s29i0, out1_s), s30})
			}
			if and(fpx, s29i1) != null {
				cSI = append(cSI, g{s29i1, and(s29i1,
					out4_s), and(s29i1, out2_s), and(s29i1, out1_s), s29})
			}
			if and(fpx, s29i2) != null {
				cSI = append(cSI, g{s29i2, and(s29i2,
					out4_s), and(s29i2, out2_s), and(s29i2, out1_s), s29})
			}
			if and(fpx, s29i3) != null {
				cSI = append(cSI, g{s29i3, and(s29i3,
					out4_s), and(s29i3, out2_s), and(s29i3, out1_s), s29})
			}
			if and(fpx, s29i4) != null {
				cSI = append(cSI, g{s29i4, and(s29i4,
					out4_s), and(s29i4, out2_s), and(s29i4, out1_s), s29})
			}
			if and(fpx, s29i5) != null {
				cSI = append(cSI, g{s29i5, and(s29i5,
					out4_s), and(s29i5, out2_s), and(s29i5, out1_s), s29})
			}
			if and(fpx, s29i6) != null {
				cSI = append(cSI, g{s29i6, and(s29i6,
					out4_s), and(s29i6, out2_s), and(s29i6, out1_s), s29})
			}
			if and(fpx, s29i7) != null {
				cSI = append(cSI, g{s29i7, and(s29i7,
					out4_s), and(s29i7, out2_s), and(s29i7, out1_s), s29})
			}
			//----------------------------------------------------------
			if and(fpx, s30i0) != null {
				cSI = append(cSI, g{s30i0, and(s30i0,
					out4_s), and(s30i0, out2_s), and(s30i0, out1_s), s30})
			}
			if and(fpx, s30i1) != null {
				cSI = append(cSI, g{s30i1, and(s30i1,
					out4_s), and(s30i1, out2_s), and(s30i1, out1_s), s30})
			}
			if and(fpx, s30i2) != null {
				cSI = append(cSI, g{s30i2, and(s30i2,
					out4_s), and(s30i2, out2_s), and(s30i2, out1_s), s30})
			}
			if and(fpx, s30i3) != null {
				cSI = append(cSI, g{s30i3, and(s30i3,
					out4_s), and(s30i3, out2_s), and(s30i3, out1_s), s30})
			}
			if and(fpx, s30i4) != null {
				cSI = append(cSI, g{s30i4, and(s30i4,
					out4_s), and(s30i4, out2_s), and(s30i4, out1_s), s30})
			}
			if and(fpx, s30i5) != null {
				cSI = append(cSI, g{s30i5, and(s30i5,
					out4_s), and(s30i5, out2_s), and(s30i5, out1_s), s30})
			}
			if and(fpx, s30i6) != null {
				cSI = append(cSI, g{s30i6, and(s30i6,
					out4_s), and(s30i6, out2_s), and(s30i6, out1_s), s30})
			}
			if and(fpx, s30i7) != null {
				cSI = append(cSI, g{s30i7, and(s30i7,
					out4_s), and(s30i7, out2_s), and(s30i7, out1_s), s31})
			}
			//----------------------------------------------------------
			if and(fpx, s31i0) != null {
				cSI = append(cSI, g{s31i0, and(s31i0,
					out4_s), and(s31i0, out2_s), and(s31i0, out1_s), s31})
			}
			if and(fpx, s31i1) != null {
				cSI = append(cSI, g{s31i1, and(s31i1,
					out4_s), and(s31i1, out2_s), and(s31i1, out1_s), s31})
			}
			if and(fpx, s31i2) != null {
				cSI = append(cSI, g{s31i2, and(s31i2,
					out4_s), and(s31i2, out2_s), and(s31i2, out1_s), s0})
			}
			if and(fpx, s31i3) != null {
				cSI = append(cSI, g{s31i3, and(s31i3,
					out4_s), and(s31i3, out2_s), and(s31i3, out1_s), s31})
			}
			if and(fpx, s31i4) != null {
				cSI = append(cSI, g{s31i4, and(s31i4,
					out4_s), and(s31i4, out2_s), and(s31i4, out1_s), s31})
			}
			if and(fpx, s31i5) != null {
				cSI = append(cSI, g{s31i5, and(s31i5,
					out4_s), and(s31i5, out2_s), and(s31i5, out1_s), s0})
			}
			if and(fpx, s31i6) != null {
				cSI = append(cSI, g{s31i6, and(s31i6,
					out4_s), and(s31i6, out2_s), and(s31i6, out1_s), s31})
			}
			if and(fpx, s31i7) != null {
				cSI = append(cSI, g{s31i7, and(s31i7,
					out4_s), and(s31i7, out2_s), and(s31i7, out1_s), s31})
			}
			//----------------------------------------------------------

			return cSI

		} // end of decompMap()

		var dMfp1, dMfp2, dMfp3, dMfp4, dMfp5, dMfp6, dMfp7,
			dMfp8, dMfp9, dMfp10, dMfp11, dMfp12, dMfp13, dMfp14,
			dMfp15, dMfp16, dMfp17, dMfp18, dMfp19, dMfp20, dMfp21,
			dMfp22, dMfp23, dMfp24, dMfp25, dMfp26, dMfp27, dMfp28,
			dMfp29, dMfp30, dMfp31 S

		dMfp1 = decompMap(fp1, out4_s, out2_s, out1_s, collectedSIs)
		dMfp2 = decompMap(fp2, out4_s, out2_s, out1_s, collectedSIs)
		dMfp3 = decompMap(fp3, out4_s, out2_s, out1_s, collectedSIs)
		dMfp4 = decompMap(fp4, out4_s, out2_s, out1_s, collectedSIs)
		dMfp5 = decompMap(fp5, out4_s, out2_s, out1_s, collectedSIs)
		dMfp6 = decompMap(fp6, out4_s, out2_s, out1_s, collectedSIs)
		dMfp7 = decompMap(fp7, out4_s, out2_s, out1_s, collectedSIs)
		dMfp8 = decompMap(fp8, out4_s, out2_s, out1_s, collectedSIs)
		dMfp9 = decompMap(fp9, out4_s, out2_s, out1_s, collectedSIs)
		dMfp10 = decompMap(fp10, out4_s, out2_s, out1_s, collectedSIs)
		dMfp11 = decompMap(fp11, out4_s, out2_s, out1_s, collectedSIs)
		dMfp12 = decompMap(fp12, out4_s, out2_s, out1_s, collectedSIs)
		dMfp13 = decompMap(fp13, out4_s, out2_s, out1_s, collectedSIs)
		dMfp14 = decompMap(fp14, out4_s, out2_s, out1_s, collectedSIs)
		dMfp15 = decompMap(fp15, out4_s, out2_s, out1_s, collectedSIs)
		dMfp16 = decompMap(fp16, out4_s, out2_s, out1_s, collectedSIs)
		dMfp17 = decompMap(fp17, out4_s, out2_s, out1_s, collectedSIs)
		dMfp18 = decompMap(fp18, out4_s, out2_s, out1_s, collectedSIs)
		dMfp19 = decompMap(fp19, out4_s, out2_s, out1_s, collectedSIs)
		dMfp20 = decompMap(fp20, out4_s, out2_s, out1_s, collectedSIs)
		dMfp21 = decompMap(fp21, out4_s, out2_s, out1_s, collectedSIs)
		dMfp22 = decompMap(fp22, out4_s, out2_s, out1_s, collectedSIs)
		dMfp23 = decompMap(fp23, out4_s, out2_s, out1_s, collectedSIs)
		dMfp24 = decompMap(fp24, out4_s, out2_s, out1_s, collectedSIs)
		dMfp25 = decompMap(fp25, out4_s, out2_s, out1_s, collectedSIs)
		dMfp26 = decompMap(fp26, out4_s, out2_s, out1_s, collectedSIs)
		dMfp27 = decompMap(fp27, out4_s, out2_s, out1_s, collectedSIs)
		dMfp28 = decompMap(fp28, out4_s, out2_s, out1_s, collectedSIs)
		dMfp29 = decompMap(fp29, out4_s, out2_s, out1_s, collectedSIs)
		dMfp30 = decompMap(fp30, out4_s, out2_s, out1_s, collectedSIs)
		dMfp31 = decompMap(fp31, out4_s, out2_s, out1_s, collectedSIs)

		// dMfps is a slice of S which is a slice of g{si, out4, out2, out1, ns} structs
        var dMfps = []S{dMfp1, dMfp2, dMfp3, dMfp4, dMfp5, dMfp6, dMfp7,
			dMfp8, dMfp9, dMfp10, dMfp11, dMfp12, dMfp13, dMfp14,
			dMfp15, dMfp16, dMfp17, dMfp18, dMfp19, dMfp20, dMfp21,
			dMfp22, dMfp23, dMfp24, dMfp25, dMfp26, dMfp27, dMfp28,
			dMfp29, dMfp30, dMfp31}
		
			return dMfps

	} // end of ns2fp

	// END of core =====================================================
	// =================================================================
	str2nd := func(f string) nd {
		mapping := map[string]nd{
			"null": null, "s0": s0, "s1": s1, "s2": s2, "s3": s3, "s4": s4,
			"s5": s5, "s6": s6, "s7": s7, "s8": s8, "s9": s9, "s10": s10,
			"s11": s11, "s12": s12, "s13": s13, "s14": s14, "s15": s15,
			"s16": s16, "s17": s17, "s18": s18, "s19": s19, "s20": s20,
			"s21": s21, "s22": s22, "s23": s23, "s24": s24, "s25": s25,
			"s26": s26, "s27": s27, "s28": s28, "s29": s29, "s30": s30,
			"s31": s31,
		}
		return mapping[f]
	}


	nd2str := func(sy nd) string {
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

	// setUP ===========================================================
	// =================================================================

	// setUP() converts a user selected S/I, usi, to a tuple
	// (ps16, ps8, ps4, ps2, ps1, in3, in2, in1) for use by one_set_BOOL()

	setUP := func(usi string, first, ns16b, ns8b, ns4b, ns2b, ns1b bool) (bool, bool, bool, bool, bool, bool, bool, bool) {

		var ps16i, ps8i, ps4i, ps2i, ps1i, in4i, in2i, in1i bool

		// Initialize all values to false
		for s := 1; s <= 31; s++ {
			for i := 0; i <= 7; i++ {
				if usi == fmt.Sprintf("s%di%d", s, i) {
					ps16i = s&16 != 0
					ps8i = s&8 != 0
					ps4i = s&4 != 0
					ps2i = s&2 != 0
					ps1i = s&1 != 0
					in4i = i&4 != 0
					in2i = i&2 != 0
					in1i = i&1 != 0
				}
			}
		}

		// Return the computed values
		return ps16i, ps8i, ps4i, ps2i, ps1i, in4i, in2i, in1i
	}

	// END of setUP ====================================================
	// =================================================================

	// one_set_BOOL ====================================================
	// =================================================================

	// uses Boolean values true, false to provide
	// a Boolean simulation of a time-frame

	one_set_BOOL := func(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i bool, fault_C string) (bool, bool, bool, bool, bool, bool, bool, bool, string) {
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

		// Apply faults to inputs
		ps1b := applyFault(ps1a, "ps1")
		ps0b := applyFault(ps0a, "ps0")
		in1b := applyFault(in1i, "in1")
		in0b := applyFault(in0i, "in0")

		// Negated values with faults
		nps1b := applyFault(!ps1b, "nps1")
		nps0b := applyFault(!ps0b, "nps0")
		nin1b := applyFault(!in1b, "nin1")
		nin0b := applyFault(!in0b, "nin0")

		// Intermediate signals with faults
		s0b := applyFault(nps1b && nps0b, "s0")
		s1b := applyFault(nps1b && ps0b, "s1")
		s2b := applyFault(ps1b && nps0b, "s2")
		i0b := applyFault(nin1b && nin0b, "i0")
		i1b := applyFault(nin1b && in0b, "i1")
		i2b := applyFault(in1b && nin0b, "i2")
		i3b := applyFault(in1b && in0b, "i3")
		ni3b := applyFault(!i3b, "ni3")

		// More intermediate signals with faults
		s0i1b := applyFault(s0b && i1b, "s0i1")
		s1i3b := applyFault(s1b && i3b, "s1i3")
		s2i0b := applyFault(s2b && i0b, "s2i0")
		s2i1b := applyFault(s2b && i1b, "s2i1")
		outb := applyFault(s1b && i2b, "out")
		s1ni3b := applyFault(s1b && ni3b, "s1ni3")

		// Final outputs with faults
		ns1b := applyFault(s1i3b || s2i1b, "ns2")
		ns0b := applyFault(s0i1b || s1ni3b || s2i0b, "ns1")

		// Determine state
		var state string
		switch {
		case !ns1b && !ns0b:
			state = "s0"
		case !ns1b && ns0b:
			state = "s1"
		case ns1b && !ns0b:
			state = "s2"
		case ns1b && ns0b:
			state = "s3"
		}

		return out4b, out2b, out1b, ns16b, ns8b, ns4b, ns2b, ns1b, state
	}

	// END of one_set_BOOL =============================================
	// =================================================================

	// peek ============================================================
	// =================================================================

	peek := func(fp int, s nd, fault_A, fault_C string) {
		// Helper function to convert `fp` to keys
		convertKeys := func(fp int, sx nd) (nd, nd, nd, nd, nd) {
			var keys [5]nd
			for i := 0; i < 5; i++ {
				if fp&(1<<i) != 0 {
					keys[4-i] = sx
				} else {
					keys[4-i] = null
				}
			}
			return keys[0], keys[1], keys[2], keys[3], keys[4]
		}

		// Convert `fp` to keys
		pt16, pt8, pt4, pt2, pt1 := convertKeys(fp, s)

		// ps2ns =======================================================
		out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s :=
			ps2ns(pt16, pt8, pt4, pt2, pt1, fault_A)
		// =============================================================

		// ns2fp =======================================================
		dMfps := ns2fp(out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s)
		// =============================================================
        
		
		// Determine if at least one key is non-null
		lns := or(pt1, or(pt2, or(pt4, or(pt8, pt16))))

		// Print circuit information
		fmt.Println("-------------------------------------------------")
		fmt.Printf("circuit: Large, A: %s, B: %s, C: %s, fp: %d, s: %s\n",
			fault_A, fault_C, fp, nd2str(s))
		fmt.Println("-------------------------------------------------")

		// Helper function to display `dMfpx`
		dispdM := func(dMfpx S, lns nd, fpn int) {
			for _, dM := range dMfpx {
				if !eq(and(lns, dM.si), null) || eq(lns, null) {
					fmt.Printf("%s  %s %s %s  fp = %d  ns = %s\n",
						allSAT(dM.si, str2nd), allSAT(dM.out_4, str2nd), allSAT(dM.out_2, str2nd),
						allSAT(dM.out_1, str2nd), fpn, nd2str(dM.ns))
				}
			}
		}
        
        // Display all `dMfps`
		for i, dMfpx := range dMfps {
			dispdM(dMfpx, lns, i+1)
		}
	}

	// END of peek =====================================================
	// =================================================================

	// foo  ============================================================
	// =================================================================

	// foo() organizes use of peek(), reports results, and accomodates
	// any created-and-unused code

	foo := func() {
		
		// Outer loop: Wait for user input
		for {
			// Get user input for faults
			first, ns16h, ns8h, ns4h, ns2h, ns1h := true, false, false, false, false, false
			var fault_A, fault_C string
			fmt.Print("Enter a fault_A name:    ")
			fmt.Scanln(&fault_A)
			fmt.Print("Enter a fault_C name:    ")
			fmt.Scanln(&fault_C)

			// Inner loop: Process fault patterns
			for {
				// Get fault pattern value
				var ufp int
				fmt.Print("Enter a fault_pattern value, fp:    ")
				fmt.Scanln(&ufp)
				if ufp == 999 {
					break // Exit to outer loop
				}
				fp := ufp

				// Get next state value
				var uns string
				fmt.Print("Enter a next_state value, s:    ")
				fmt.Scanln(&uns)
				s := str2nd(uns)

				// Call `peek` function
				peek(fp, s, fault_A, fault_C)
				fmt.Println("========================================================3")

				// Get selected S/I
				var usi string
				fmt.Print("                                   Enter selected S/I: ")
				fmt.Scanln(&usi)
				fmt.Println()

				// Call `setUP` and `one_set_BOOL`
				ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i := setUP(usi, first, ns16h, ns8h, ns4h, ns2h, ns1h)
				out4b, out2b, out1b, ns16b, ns8b, ns4b, ns2b, ns1b, state := one_set_BOOL(ps16a, ps8a, ps4a, ps2a, ps1a, in4i, in2i, in1i, fault_C)

				// Update state variables
				ns16h = ns16b
				ns8h = ns8b
				ns4h = ns4b
				ns2h = ns2b
				ns1h = ns1b
				first = false

				// Print results
				fmt.Printf("                                  outb: %v\n"    next_state: %s    fault_C: %s\n", outb, state, fault_C)
				fmt.Printf("                                  next_state: %s")
				fmt.Printf("                                  fault_C: %s\n")
				fmt.Println("========================================================4")
			}
		}
	}

	// END of foo ======================================================
	// =================================================================

	foo()

	// =================================================================

} // end of main()

// END END END =========================================================
