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

		
        // Completed examples
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

        // Instructions for Copilot
		// Add applyFault() statements for each signal below in the format:
        // signal_s = applyFault(signal_s, signal_1, fault_A, "signal")

		// Bare signal definitions

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
		 