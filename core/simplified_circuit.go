package core

import (
	"fmt"

	"github.com/dalzilio/rudd"
	"rudd_Large.go/types"
)

// Simplified circuit operations for initial setup

// SimplifiedActivatePropagateFaultA provides a basic version for testing
func SimplifiedActivatePropagateFaultA(ps16_i, ps8_i, ps4_i, ps2_i, ps1_i types.Nd,
	fault_A string) (types.Nd, types.Nd, types.Nd, types.Nd, types.Nd, types.Nd, types.Nd, types.Nd) {

	fmt.Printf("=== SIMPLIFIED activatePropagateFaultA ===\n")
	fmt.Printf("Input state: ps16=%s, ps8=%s, ps4=%s, ps2=%s, ps1=%s\n",
		Nd2str(ps16_i), Nd2str(ps8_i), Nd2str(ps4_i), Nd2str(ps2_i), Nd2str(ps1_i))
	fmt.Printf("fault_A: %s\n", fault_A)

	// For now, return simplified outputs
	out4_s := Null
	out2_s := Null
	out1_s := Null
	ns16_s := ps16_i
	ns8_s := ps8_i
	ns4_s := ps4_i
	ns2_s := ps2_i
	ns1_s := ps1_i

	fmt.Printf("Output state: out4=%s, out2=%s, out1=%s\n",
		Nd2str(out4_s), Nd2str(out2_s), Nd2str(out1_s))
	fmt.Printf("Next state: ns16=%s, ns8=%s, ns4=%s, ns2=%s, ns1=%s\n",
		Nd2str(ns16_s), Nd2str(ns8_s), Nd2str(ns4_s), Nd2str(ns2_s), Nd2str(ns1_s))

	return out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s
}

// Helper functions for circuit analysis
func GetBDD() *rudd.BDD {
	return Bdd
}

func CreateStateVector(s16, s8, s4, s2, s1 bool) types.Nd {
	var result types.Nd = Bdd.True()

	if s16 {
		result = And(result, Ps16)
	} else {
		result = And(result, Nps16)
	}

	if s8 {
		result = And(result, Ps8)
	} else {
		result = And(result, Nps8)
	}

	if s4 {
		result = And(result, Ps4)
	} else {
		result = And(result, Nps4)
	}

	if s2 {
		result = And(result, Ps2)
	} else {
		result = And(result, Nps2)
	}

	if s1 {
		result = And(result, Ps1)
	} else {
		result = And(result, Nps1)
	}

	return result
}
