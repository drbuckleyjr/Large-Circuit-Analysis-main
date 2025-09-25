package core

import (
	"fmt"

	"github.com/dalzilio/rudd"
	"rudd_Large.go/types"
)

// BDD (Binary Decision Diagram) setup and core operations
var (
	Bdd                                                   *rudd.BDD
	Nd128, Nd64, Nd32, Nd16, Nd8, Nd4, Nd2, Nd1           rudd.Node
	Not                                                   func(rudd.Node) rudd.Node
	And                                                   func(...rudd.Node) rudd.Node
	Or                                                    func(...rudd.Node) rudd.Node
	IsNull                                                func(rudd.Node) bool
	Null                                                  rudd.Node
	Ps16, Ps8, Ps4, Ps2, Ps1, In4, In2, In1               rudd.Node
	Nps16, Nps8, Nps4, Nps2, Nps1, Nin4, Nin2, Nin1       rudd.Node
	S0, S1, S2, S3, S4, S5, S6, S7, S8, S9, S10, S11, S12 rudd.Node
	S13, S14, S15, S16, S17, S18, S19, S20, S21, S22, S23 rudd.Node
	S24, S25, S26, S27, S28, S29, S30, S31                rudd.Node
)

// Initialize initializes the BDD manager and all core variables
func Initialize() {
	var err error
	Bdd, err = rudd.New(8, rudd.Nodesize(10000), rudd.Cachesize(3000))
	if err != nil {
		panic(fmt.Sprintf("Failed to create BDD: %v", err))
	}

	// Create BDD variables
	Nd128 = Bdd.Ithvar(7)
	Nd64 = Bdd.Ithvar(6)
	Nd32 = Bdd.Ithvar(5)
	Nd16 = Bdd.Ithvar(4)
	Nd8 = Bdd.Ithvar(3)
	Nd4 = Bdd.Ithvar(2)
	Nd2 = Bdd.Ithvar(1)
	Nd1 = Bdd.Ithvar(0)

	// Create operation functions
	Not = Bdd.Not
	And = Bdd.And
	Or = Bdd.Or
	Null = Bdd.False()

	IsNull = func(n rudd.Node) bool {
		return Bdd.Equal(n, Null)
	}

	// Initialize primary state and input variables
	Ps16 = Nd128
	Ps8 = Nd64
	Ps4 = Nd32
	Ps2 = Nd16
	Ps1 = Nd8
	In4 = Nd4
	In2 = Nd2
	In1 = Nd1

	// Initialize negated variables
	Nps16 = Not(Ps16)
	Nps8 = Not(Ps8)
	Nps4 = Not(Ps4)
	Nps2 = Not(Ps2)
	Nps1 = Not(Ps1)
	Nin4 = Not(In4)
	Nin2 = Not(In2)
	Nin1 = Not(In1)

	// Initialize state variables
	initializeStates()
}

// initializeStates creates all 32 state variables
func initializeStates() {
	S0 = And(Nps16, Nps8, Nps4, Nps2, Nps1)
	S1 = And(Nps16, Nps8, Nps4, Nps2, Ps1)
	S2 = And(Nps16, Nps8, Nps4, Ps2, Nps1)
	S3 = And(Nps16, Nps8, Nps4, Ps2, Ps1)
	S4 = And(Nps16, Nps8, Ps4, Nps2, Nps1)
	S5 = And(Nps16, Nps8, Ps4, Nps2, Ps1)
	S6 = And(Nps16, Nps8, Ps4, Ps2, Nps1)
	S7 = And(Nps16, Nps8, Ps4, Ps2, Ps1)
	S8 = And(Nps16, Ps8, Nps4, Nps2, Nps1)
	S9 = And(Nps16, Ps8, Nps4, Nps2, Ps1)
	S10 = And(Nps16, Ps8, Nps4, Ps2, Nps1)
	S11 = And(Nps16, Ps8, Nps4, Ps2, Ps1)
	S12 = And(Nps16, Ps8, Ps4, Nps2, Nps1)
	S13 = And(Nps16, Ps8, Ps4, Nps2, Ps1)
	S14 = And(Nps16, Ps8, Ps4, Ps2, Nps1)
	S15 = And(Nps16, Ps8, Ps4, Ps2, Ps1)
	S16 = And(Ps16, Nps8, Nps4, Nps2, Nps1)
	S17 = And(Ps16, Nps8, Nps4, Nps2, Ps1)
	S18 = And(Ps16, Nps8, Nps4, Ps2, Nps1)
	S19 = And(Ps16, Nps8, Nps4, Ps2, Ps1)
	S20 = And(Ps16, Nps8, Ps4, Nps2, Nps1)
	S21 = And(Ps16, Nps8, Ps4, Nps2, Ps1)
	S22 = And(Ps16, Nps8, Ps4, Ps2, Nps1)
	S23 = And(Ps16, Nps8, Ps4, Ps2, Ps1)
	S24 = And(Ps16, Ps8, Nps4, Nps2, Nps1)
	S25 = And(Ps16, Ps8, Nps4, Nps2, Ps1)
	S26 = And(Ps16, Ps8, Nps4, Ps2, Nps1)
	S27 = And(Ps16, Ps8, Nps4, Ps2, Ps1)
	S28 = And(Ps16, Ps8, Ps4, Nps2, Nps1)
	S29 = And(Ps16, Ps8, Ps4, Nps2, Ps1)
	S30 = And(Ps16, Ps8, Ps4, Ps2, Nps1)
	S31 = And(Ps16, Ps8, Ps4, Ps2, Ps1)
}

// GetStateByNumber returns the state variable for a given state number (0-31)
func GetStateByNumber(stateNum int) types.Nd {
	states := []types.Nd{
		S0, S1, S2, S3, S4, S5, S6, S7, S8, S9, S10, S11, S12, S13, S14, S15,
		S16, S17, S18, S19, S20, S21, S22, S23, S24, S25, S26, S27, S28, S29, S30, S31,
	}
	if stateNum >= 0 && stateNum < len(states) {
		return states[stateNum]
	}
	return Null
}

// Str2nd converts a string to a BDD node
func Str2nd(f string) types.Nd {
	switch f {
	case "s0":
		return S0
	case "s1":
		return S1
	case "s2":
		return S2
	case "s3":
		return S3
	case "s4":
		return S4
	case "s5":
		return S5
	case "s6":
		return S6
	case "s7":
		return S7
	case "s8":
		return S8
	case "s9":
		return S9
	case "s10":
		return S10
	case "s11":
		return S11
	case "s12":
		return S12
	case "s13":
		return S13
	case "s14":
		return S14
	case "s15":
		return S15
	case "s16":
		return S16
	case "s17":
		return S17
	case "s18":
		return S18
	case "s19":
		return S19
	case "s20":
		return S20
	case "s21":
		return S21
	case "s22":
		return S22
	case "s23":
		return S23
	case "s24":
		return S24
	case "s25":
		return S25
	case "s26":
		return S26
	case "s27":
		return S27
	case "s28":
		return S28
	case "s29":
		return S29
	case "s30":
		return S30
	case "s31":
		return S31
	}
	return Null
}

// Nd2str converts a BDD node to its string representation
func Nd2str(sy types.Nd) string {
	if IsNull(sy) {
		return "null"
	}

	// Check all states
	stateNames := []string{
		"s0", "s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11", "s12", "s13", "s14", "s15",
		"s16", "s17", "s18", "s19", "s20", "s21", "s22", "s23", "s24", "s25", "s26", "s27", "s28", "s29", "s30", "s31",
	}
	states := []types.Nd{
		S0, S1, S2, S3, S4, S5, S6, S7, S8, S9, S10, S11, S12, S13, S14, S15,
		S16, S17, S18, S19, S20, S21, S22, S23, S24, S25, S26, S27, S28, S29, S30, S31,
	}

	for i, state := range states {
		if Bdd.Equal(sy, state) {
			return stateNames[i]
		}
	}
	return "unknown"
}

// AllSAT returns all satisfying assignments for a BDD node
func AllSAT(f types.Nd, str2nd func(string) types.Nd) []string {
	if IsNull(f) {
		return []string{}
	}

	var result []string
	// Check each state
	for i := 0; i < 32; i++ {
		stateName := fmt.Sprintf("s%d", i)
		stateNode := str2nd(stateName)
		if !IsNull(And(f, stateNode)) {
			result = append(result, stateName)
		}
	}

	// Check each input
	for i := 0; i < 8; i++ {
		inputName := fmt.Sprintf("i%d", i)
		// Note: This is simplified - you may need to implement input nodes
		result = append(result, inputName)
	}

	return result
}

// Circuit logic rules
func PiRule(s1, i1 types.Nd) (types.Nd, types.Nd) {
	return s1, i1
}

func NotRule(s1, i1 types.Nd) (types.Nd, types.Nd) {
	return Not(s1), Not(i1)
}

func Or2Rule(s1, s2, i1, i2 types.Nd) (types.Nd, types.Nd) {
	return Or(s1, s2), Or(i1, i2)
}

func And2Rule(s1, s2, i1, i2 types.Nd) (types.Nd, types.Nd) {
	return And(s1, s2), And(i1, i2)
}

func And3Rule(s1, s2, s3, i1, i2, i3 types.Nd) (types.Nd, types.Nd) {
	return And(s1, s2, s3), And(i1, i2, i3)
}

func Or3Rule(s1, s2, s3, i1, i2, i3 types.Nd) (types.Nd, types.Nd) {
	return Or(s1, s2, s3), Or(i1, i2, i3)
}
