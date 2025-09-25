package simulation

import (
	"fmt"
	"sort"
	"strings"

	"rudd_Large.go/core"
	"rudd_Large.go/types"
)

// ExtractSIs extracts all S/I pairs from a BDD node using allSAT
func ExtractSIs(f types.Nd) []types.SIMapping {
	if core.IsNull(f) {
		return []types.SIMapping{}
	}

	// Get all satisfying assignments
	assignments := core.AllSAT(f, core.Str2nd)

	var siMappings []types.SIMapping

	// Parse assignments to extract S/I pairs
	for _, assignment := range assignments {
		if strings.HasPrefix(assignment, "s") {
			// Extract state number
			mapping := types.SIMapping{
				FP: 0, // Default fault pattern
				NS: f, // Store the node
			}
			siMappings = append(siMappings, mapping)
		}
	}

	return siMappings
}

// ProcessSIs processes S/I collections and converts them to our internal format
func ProcessSIs(sis []types.SIMapping) []types.G {
	var result []types.G

	for _, si := range sis {
		g := types.G{
			Si:   si.NS,
			Out4: core.Null,
			Out2: core.Null,
			Out1: core.Null,
			Ns:   si.NS,
		}
		result = append(result, g)
	}

	return result
}

// SimulateWithFault runs circuit simulation with a specific fault
func SimulateWithFault(state types.Nd, fault string) types.SimResultWithState {
	// Extract state components (simplified for now)
	stateStr := core.Nd2str(state)

	// Run the circuit simulation
	out4, out2, out1, ns16, ns8, ns4, ns2, ns1 := core.SimplifiedActivatePropagateFaultA(
		state, state, state, state, state, fault)

	// Create next state from outputs
	nextState := core.And(ns16, ns8, ns4, ns2, ns1)

	result := types.SimResultWithState{
		SI: stateStr + "i0", // Simplified
		Outputs: fmt.Sprintf("out4=%s,out2=%s,out1=%s",
			core.Nd2str(out4), core.Nd2str(out2), core.Nd2str(out1)),
		NextState: core.Nd2str(nextState),
		First:     types.First,
		Ns16h:     types.Ns16h,
		Ns8h:      types.Ns8h,
		Ns4h:      types.Ns4h,
		Ns2h:      types.Ns2h,
		Ns1h:      types.Ns1h,
	}

	return result
}

// AccumulateSIs adds new S/I pairs to the global accumulator
func AccumulateSIs(newSIs []string) {
	types.AccumulatedSIs = append(types.AccumulatedSIs, newSIs...)

	// Remove duplicates (simplified approach)
	seen := make(map[string]bool)
	var unique []string

	for _, si := range types.AccumulatedSIs {
		if !seen[si] {
			seen[si] = true
			unique = append(unique, si)
		}
	}

	types.AccumulatedSIs = unique

	fmt.Printf("Accumulated %d unique S/I pairs\n", len(types.AccumulatedSIs))
}

// GetAccumulatedSICount returns the current count of accumulated S/I pairs
func GetAccumulatedSICount() int {
	return len(types.AccumulatedSIs)
}

// ClearAccumulatedSIs resets the accumulated S/I collection
func ClearAccumulatedSIs() {
	types.AccumulatedSIs = []string{}
	fmt.Println("Cleared accumulated S/I pairs")
}

// SortAccumulatedSIs sorts the accumulated S/I pairs for consistent ordering
func SortAccumulatedSIs() {
	sort.Strings(types.AccumulatedSIs)
}

// PrintAccumulatedSIs displays all accumulated S/I pairs
func PrintAccumulatedSIs() {
	fmt.Printf("=== Accumulated S/I Pairs (%d total) ===\n", len(types.AccumulatedSIs))
	for i, si := range types.AccumulatedSIs {
		fmt.Printf("%d: %s\n", i+1, si)
	}
	fmt.Println("=====================================")
}

// ValidateSI checks if an S/I pair is valid for the circuit
func ValidateSI(siPair string) bool {
	// Check format: should be like "s0i0", "s31i7", etc.
	if len(siPair) < 4 || !strings.HasPrefix(siPair, "s") {
		return false
	}

	// Find the 'i' separator
	iPos := strings.Index(siPair, "i")
	if iPos == -1 {
		return false
	}

	// Extract and validate state part
	state := siPair[1:iPos]
	for i := 0; i < 32; i++ {
		if state == fmt.Sprintf("%d", i) {
			return true
		}
	}

	return false
}

// FilterValidSIs removes invalid S/I pairs from a collection
func FilterValidSIs(sis []string) []string {
	var valid []string

	for _, si := range sis {
		if ValidateSI(si) {
			valid = append(valid, si)
		}
	}

	return valid
}

// AddSimulationResult stores a simulation result
func AddSimulationResult(result types.SimResult, faultFree bool) {
	if faultFree {
		types.AccumulatedFaultFreeSimulations = append(types.AccumulatedFaultFreeSimulations, result)
	} else {
		types.AccumulatedFaultASimulations = append(types.AccumulatedFaultASimulations, result)
	}
}

// GetSimulationResultCount returns counts of stored simulation results
func GetSimulationResultCount() (faultFree, faultA int) {
	return len(types.AccumulatedFaultFreeSimulations), len(types.AccumulatedFaultASimulations)
}

// ClearSimulationResults resets all stored simulation results
func ClearSimulationResults() {
	types.AccumulatedFaultFreeSimulations = []types.SimResult{}
	types.AccumulatedFaultASimulations = []types.SimResult{}
}
