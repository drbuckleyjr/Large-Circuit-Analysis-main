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
		nps8_s = applyFault(nps8_s, nps8_1, fault_A, "nps8")1
		 
		nps16_s, nps16_1 := notRule(ps16_s, ps16_1)
		nps16_s = applyFault(nps16_s, nps16_1, fault_A, "nps16")
		 
		nin1_s, nin1_1 := notRule(in1_s, in1_1)
		nin1_s = applyFault(nin1_s, nin1_1, fault_A, "nin1")
		 
		nin2_s, nin2_1 := notRule(in2_s, in2_1)
		nin2_s = applyFault(nin2_s, nin2_1, fault_A, "nin2")
		 
		nin4_s, nin4_1 := notRule(in4_s, in4_1)
		nin4_s = applyFault(nin4_s, nin4_1, fault_A, "nin4")
		 
		i7_s, i7_1 := and3Rule(in4_s, in2_s, in1_s, in4_1, in2_1, in1_1)
		I7_s = applyFault(i7_s, i7_1, fault_A, "i7")
		 

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


		// Note: The applyFault() statements for the signals below
		
		// Add applyFault() statements for the remaining signals below.
        // Each signal should follow this format:
        // signal_s = applyFault(signal_s, signal_1, fault_A, "<signal>") 
		
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

		// Add applyFault() statement for s15_s, s15_1
		// Note: s15_s, s15_1 is defined below, so we can apply the fault here. 
		s15_s, s15_1 := and3Rule(nps16_s, ps8_s, ls7_s, nps16_1, ps8_1, ls7_1)
		s15_s = applyFault(s15_s, s15_1, fault_A, "s15")
		// Copilot: ADD applyFault() statement for all signals below.
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
		 
		s19_s, s19_1 := and3Rule(ps16_s, nps8_s, ls3_s, ps16_1, nps8_1, ls3_1)
		 
		s20_s, s20_1 := and3Rule(ps16_s, nps8_s, ls4_s, ps16_1, nps8_1, ls4_1)
		 
		s21_s, s21_1 := and3Rule(ps16_s, nps8_s, ls5_s, ps16_1, nps8_1, ls5_1)
		 
		s22_s, s22_1 := and3Rule(ps16_s, nps8_s, ls6_s, ps16_1, nps8_1, ls6_1)
		 
		s23_s, s23_1 := and3Rule(ps16_s, nps8_s, ls7_s, ps16_1, nps8_1, ls7_1)
		 
		s24_s, s24_1 := and3Rule(ps16_s, ps8_s, ls0_s, ps16_1, ps8_1, ls0_1)
		 
		s25_s, s25_1 := and3Rule(ps16_s, ps8_s, ls1_s, ps16_1, ps8_1, ls1_1)
		 
		s26_s, s26_1 := and3Rule(ps16_s, ps8_s, ls2_s, ps16_1, ps8_1, ls2_1)
		
		s27_s, s27_1 := and3Rule(ps16_s, ps8_s, ls3_s, ps16_1, ps8_1, ls3_1)

		s28_s, s28_1 := and3Rule(ps16_s, ps8_s, ls4_s, ps16_1, ps8_1, ls4_1)
		 
		s29_s, s29_1 := and3Rule(ps16_s, ps8_s, ls5_s, ps16_1, ps8_1, ls5_1)
		 
		s30_s, s30_1 := and3Rule(ps16_s, ps8_s, ls6_s, ps16_1, ps8_1, ls6_1)
		 
		b2_s, b2_1 := or3Rule(i5_s, i3_s, i2_s, i5_1, i3_1, i2_1)
		 
		b7_s, b7_1 := or2Rule(i5_s, i1_s, i5_1, i1_1)
		 
		c5_s, c5_1 := or2Rule(i7_s, i5_s, i7_1, i5_1)
		 
		c13_s, c13_1 := or2Rule(i3_s, i2_s, i3_1, i2_1)
		 
		e5_s, e5_1 := or2Rule(i5_s, i4_s, i5_1, i4_1)
		 
		e10_s, e10_1 := or3Rule(i6_s, i2_s, i0_s, i6_1, i2_1, i0_1)
		 
		e12_s, e12_1 := or2Rule(i6_s, i3_s, i6_1, i3_1)
		 
		e18_s, e18_1 := or2Rule(i6_s, i2_s, i6_1, i2_1)
		 
		e24_s, e24_1 := or2Rule(i7_s, i1_s, i7_1, i1_1)
		 

		// level 4

		 
		a1_s, a1_1 := and2Rule(s10_s, i0_s, s10_1, i0_1)
		 
		a2_s, a2_1 := and2Rule(s15_s, i5_s, s15_1, i5_1)
		 
		a3_s, a3_1 := and2Rule(s18_s, ni6_s, s18_1, ni6_1)
		 
		a4_s, a4_1 := or3Rule(s20_s, s21_s, s22_s, s20_1, s21_1, s22_1)
		 
		a5_s, a5_1 := and2Rule(s24_s, ni7_s, s24_1, ni7_1)
		 
		a6_s, a6_1 := or2Rule(s25_s, s26_s, s25_1, s26_1)
		 
		a7_s, a7_1 := or3Rule(s27_s, s28_s, s29_s, s27_1, s28_1, s29_1)
		 
		a9_s, a9_1 := and3Rule(s31_s, ni5_s, ni2_s, s31_1, ni5_1, ni2_1)
		 
		b1_s, b1_1 := and2Rule(s3_s, i2_s, s3_1, i2_1)
		 
		b3_s, b3_1 := and2Rule(b2_s, s7_s, b2_1, s7_1)
		 
		b4_s, b4_1 := and2Rule(s10_s, ni6_s, s10_1, ni6_1)
		 
		b5_s, b5_1 := or3Rule(s12_s, s13_s, s14_s, s12_1, s13_1, s14_1)
		 
		b6_s, b6_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		 
		b8_s, b8_1 := and2Rule(s23_s, b7_s, s23_1, b7_1)
		 
		b9_s, b9_1 := or2Rule(s25_s, s26_s, s25_1, s26_1)
		 
		b10_s, b10_1 := or3Rule(s27_s, s28_s, s29_s, s27_1, s28_1, s29_1)
		 
		c2_s, c2_1 := and3Rule(s7_s, ni5_s, ni3_s, s7_1, ni5_1, ni3_1)
		 
		c3_s, c3_1 := and2Rule(s11_s, i7_s, s11_1, i7_1)
		 
		c4_s, c4_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		 
		c6_s, c6_1 := and2Rule(s23_s, ni5_s, s23_1, ni5_1)
		 
		c8_s, c8_1 := and2Rule(s19_s, c5_s, s19_1, c5_1)
		 
		c14_s, c14_1 := and2Rule(c13_s, s3_s, c13_1, s3_1)
		 
		c15_s, c15_1 := and2Rule(s27_s, i7_s, s27_1, i7_1)
		 
		d1_s, d1_1 := and2Rule(s1_s, i2_s, s1_1, i2_1)
		 
		d2_s, d2_1 := and3Rule(s3_s, ni3_s, ni2_s, s3_1, ni3_1, ni2_1)
		 
		d3_s, d3_1 := and2Rule(s5_s, i0_s, s5_1, i0_1)
		 
		d5_s, d5_1 := and2Rule(s9_s, i2_s, s9_1, i2_1)
		 
		d6_s, d6_1 := and2Rule(s11_s, ni7_s, s11_1, ni7_1)
		 
		d7_s, d7_1 := and2Rule(s13_s, i0_s, s13_1, i0_1)
		 
		d9_s, d9_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		 
		d10_s, d10_1 := and2Rule(s17_s, i2_s, s17_1, i2_1)
		 
		d11_s, d11_1 := and2Rule(s19_s, ni7_s, s19_1, ni7_1)
		 
		d12_s, d12_1 := and2Rule(s21_s, i0_s, s21_1, i0_1)
		 
		d13_s, d13_1 := and3Rule(s23_s, ni5_s, ni1_s, s23_1, ni5_1, ni1_1)
		 
		d14_s, d14_1 := and2Rule(s25_s, i2_s, s25_1, i2_1)
		 
		d15_s, d15_1 := and2Rule(s29_s, i0_s, s29_1, i0_1)
		 
		d27_s, d27_1 := and2Rule(s27_s, ni7_s, s27_1, ni7_1)
		 
		e1_s, e1_1 := and2Rule(s0_s, i1_s, s0_1, i1_1)
		 
		e2_s, e2_1 := and2Rule(s1_s, ni2_s, s1_1, ni2_1)
		 
		e3_s, e3_1 := and2Rule(s2_s, i2_s, s2_1, i2_1)
		 
		e4_s, e4_1 := and3Rule(s3_s, ni3_s, ni2_s, s3_1, ni3_1, ni2_1)
		 
		e6_s, e6_1 := and2Rule(s5_s, ni0_s, s5_1, ni0_1)
		 
		e7_s, e7_1 := and2Rule(s6_s, i7_s, s6_1, i7_1)
		 
		e8_s, e8_1 := and2Rule(s8_s, i1_s, s8_1, i1_1)
		 
		e9_s, e9_1 := and2Rule(s9_s, ni2_s, s9_1, ni2_1)
		 
		e11_s, e11_1 := and2Rule(s11_s, ni7_s, s11_1, ni7_1)
		 
		e13_s, e13_1 := and2Rule(s13_s, ni0_s, s13_1, ni0_1)
		 
		e14_s, e14_1 := and2Rule(s14_s, i7_s, s14_1, i7_1)
		 
		e15_s, e15_1 := and2Rule(s15_s, ni5_s, s15_1, ni5_1)
		 
		e16_s, e16_1 := and2Rule(s16_s, i1_s, s16_1, i1_1)
		 
		e17_s, e17_1 := and2Rule(s17_s, ni2_s, s17_1, ni2_1)
		 
		e19_s, e19_1 := and2Rule(s19_s, ni7_s, s19_1, ni7_1)
		 
		e20_s, e20_1 := and2Rule(s20_s, e12_s, s20_1, e12_1)
		 
		e21_s, e21_1 := and2Rule(s21_s, ni0_s, s21_1, ni0_1)
		 
		e22_s, e22_1 := and2Rule(s22_s, i7_s, s22_1, i7_1)
		 
		e23_s, e23_1 := and2Rule(s23_s, ni5_s, s23_1, ni5_1)
		 
		e25_s, e25_1 := and2Rule(s25_s, ni2_s, s25_1, ni2_1)
		 
		e26_s, e26_1 := and2Rule(s26_s, i2_s, s26_1, i2_1)
		 
		e27_s, e27_1 := and2Rule(s27_s, ni7_s, s27_1, ni7_1)
		 
		e28_s, e28_1 := and2Rule(s28_s, e12_s, s28_1, e12_1)
		 
		e29_s, e29_1 := and2Rule(s29_s, ni0_s, s29_1, ni0_1)
		 
		e30_s, e30_1 := and2Rule(s30_s, i7_s, s30_1, i7_1)
		 
		e31_s, e31_1 := and2Rule(s4_s, e5_s, s4_1, e5_1)
		 
		e32_s, e32_1 := and2Rule(s10_s, e10_s, s10_1, e10_1)
		 
		e33_s, e33_1 := and2Rule(s12_s, e12_s, s12_1, e12_1)
		 
		e34_s, e34_1 := and2Rule(s18_s, e18_s, s18_1, e18_1)
		 
		e35_s, e35_1 := and2Rule(s24_s, e24_s, s24_1, e24_1)
		 
		f1_s, f1_1 := and2Rule(s12_s, i5_s, s12_1, i5_1)
		 
		f2_s, f2_1 := and2Rule(s27_s, i4_s, s27_1, i4_1)
		 
		f3_s, f3_1 := and2Rule(s15_s, i0_s, s15_1, i0_1)
		 
		f4_s, f4_1 := and2Rule(s27_s, i2_s, s27_1, i2_1)
		 
		f5_s, f5_1 := and2Rule(s0_s, i7_s, s0_1, i7_1)
		 
		f6_s, f6_1 := and2Rule(s27_s, i1_s, s27_1, i1_1)
		 

		// level 5

		 
		a8_s, a8_1 := or2Rule(a7_s, s30_s, a7_1, s30_1)
		 
		a10_s, a10_1 := or3Rule(a1_s, a2_s, s16_s, a1_1, a2_1, s16_1)
		 
		b11_s, b11_1 := or2Rule(b10_s, s30_s, b10_1, s30_1)
		 
		b12_s, b12_1 := or3Rule(b1_s, b3_s, s8_s, b1_1, b3_1, s8_1)
		 
		b13_s, b13_1 := or3Rule(s9_s, b4_s, s11_s, s9_1, b4_1, s11_1)
		 
		c1_s, c1_1 := or3Rule(c14_s, s4_s, s5_s, c14_1, s4_1, s5_1)
		 
		c7_s, c7_1 := and2Rule(c2_s, ni2_s, c2_1, ni2_1)
		 
		c16_s, c16_1 := or3Rule(c15_s, s28_s, s29_s, c15_1, s28_1, s29_1)
		 
		d17_s, d17_1 := or3Rule(d1_s, d2_s, d3_s, d1_1, d2_1, d3_1)
		 
		d19_s, d19_1 := or3Rule(s10_s, d6_s, d7_s, s10_1, d6_1, d7_1)
		 
		d20_s, d20_1 := or3Rule(s14_s, d9_s, d10_s, s14_1, d9_1, d10_1)
		 
		d21_s, d21_1 := or3Rule(s18_s, d11_s, d12_s, s18_1, d11_1, d12_1)
		 
		d22_s, d22_1 := or3Rule(s22_s, d13_s, d14_s, s22_1, d13_1, d14_1)
		 
		d23_s, d23_1 := or3Rule(s26_s, d15_s, s30_s, s26_1, d15_1, s30_1)
		 
		e36_s, e36_1 := or3Rule(e1_s, e2_s, e3_s, e1_1, e2_1, e3_1)
		 
		e37_s, e37_1 := or3Rule(e4_s, e31_s, e6_s, e4_1, e31_1, e6_1)
		 
		e39_s, e39_1 := or2Rule(e9_s, e32_s, e9_1, e32_1)
		 
		e40_s, e40_1 := or3Rule(e11_s, e33_s, e13_s, e11_1, e33_1, e13_1)
		 
		e41_s, e41_1 := or3Rule(e14_s, e15_s, e16_s, e14_1, e15_1, e16_1)
		 
		e42_s, e42_1 := or3Rule(e17_s, e34_s, e19_s, e17_1, e34_1, e19_1)
		 
		e43_s, e43_1 := or3Rule(e20_s, e30_s, a9_s, e20_1, e30_1, a9_1)
		 
		e44_s, e44_1 := or3Rule(e21_s, e22_s, e23_s, e21_1, e22_1, e23_1)
		 
		e45_s, e45_1 := or3Rule(e35_s, e25_s, e26_s, e35_1, e25_1, e26_1)
		 
		e46_s, e46_1 := or3Rule(e27_s, e28_s, e29_s, e27_1, e28_1, e29_1)
		 
		out4_s, out4_1 := or2Rule(f1_s, f2_s, f1_1, f2_1)
		 
		out2_s, out2_1 := or2Rule(f3_s, f4_s, f3_1, f4_1)
		 
		out1_s, out1_1 := or2Rule(f5_s, f6_s, f5_1, f6_1)
		 

		// level 6

		 
		a11_s, a11_1 := or3Rule(a10_s, s17_s, a3_s, a10_1, s17_1, a3_1)
		 
		b14_s, b14_1 := or3Rule(b12_s, b13_s, b5_s, b12_1, b13_1, b5_1)
		 
		c9_s, c9_1 := or3Rule(c1_s, s6_s, c7_s, c1_1, s6_1, c7_1)
		 
		c17_s, c17_1 := or2Rule(c16_s, s30_s, c16_1, s30_1)
		 
		d18_s, d18_1 := or3Rule(s6_s, c7_s, d5_s, s6_1, c7_1, d5_1)
		 
		d28_s, d28_1 := or2Rule(d23_s, d27_s, d23_1, d27_1)
		 
		e38_s, e38_1 := or3Rule(e7_s, c7_s, e8_s, e7_1, c7_1, e8_1)
		 
		e47_s, e47_1 := or3Rule(e36_s, e37_s, e44_s, e36_1, e37_1, e44_1)
		 
		e49_s, e49_1 := or3Rule(e39_s, e40_s, e41_s, e39_1, e40_1, e41_1)
		 

		// level 7

		 
		a12_s, a12_1 := or3Rule(a11_s, s19_s, a4_s, a11_1, s19_1, a4_1)
		 
		b15_s, b15_1 := or3Rule(b14_s, b6_s, b8_s, b14_1, b6_1, b8_1)
		 
		c10_s, c10_1 := or3Rule(c9_s, c3_s, b5_s, c9_1, c3_1, b5_1)
		 
		d24_s, d24_1 := or3Rule(d17_s, d18_s, d19_s, d17_1, d18_1, d19_1)
		 
		e48_s, e48_1 := or3Rule(e45_s, e46_s, e38_s, e45_1, e46_1, e38_1)
		 

		// level 8

		 
		a13_s, a13_1 := or3Rule(a12_s, s23_s, a5_s, a12_1, s23_1, a5_1)
		 
		b16_s, b16_1 := or3Rule(b15_s, s24_s, b9_s, b15_1, s24_1, b9_1)
		 
		c11_s, c11_1 := or3Rule(c10_s, c4_s, c8_s, c10_1, c4_1, c8_1)
		 
		d25_s, d25_1 := or3Rule(d24_s, d20_s, d21_s, d24_1, d20_1, d21_1)
		 
		e50_s, e50_1 := or3Rule(e47_s, e48_s, e49_s, e47_1, e48_1, e49_1)
		 

		// level 9

		 
		ns1_s, ns1_1 := or3Rule(e50_s, e42_s, e43_s, e50_1, e42_1, e43_1)
		 
		ns8_s, ns8_1 := or3Rule(b16_s, b11_s, a9_s, b16_1, b11_1, a9_1)
		 
		a14_s, a14_1 := or3Rule(a13_s, a6_s, a8_s, a13_1, a6_1, a8_1)
		 
		c12_s, c12_1 := or3Rule(c11_s, a4_s, c6_s, c11_1, a4_1, c6_1)
		 
		d26_s, d26_1 := or3Rule(d25_s, d22_s, d28_s, d25_1, d22_1, d28_1)
		 

		// level 10

		 
		ns2_s, ns2_1 := or3Rule(d26_s, s2_s, a9_s, d26_1, s2_1, a9_1)
		 
		ns4_s, ns4_1 := or3Rule(c12_s, c17_s, a9_s, c12_1, c17_1, a9_1)
		 
		ns16_s, ns16_1 := or2Rule(a14_s, a9_s, a14_1, a9_1)
		 

		return out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s

	} // End of ps2ns() ================================================