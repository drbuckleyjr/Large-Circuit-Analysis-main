package types

import "github.com/dalzilio/rudd"

// Shared types and global variables

type Nd = rudd.Node // Binary Decision Diagram node type alias

type G struct { // Gate structure for S/I and output mapping
	Si, Out4, Out2, Out1, Ns Nd
}

type S []G // State collection type

// SimResult stores the results of a single S/I simulation
// Exported for use in other packages
// Capitalized field names for export
type SimResult struct {
	SI        string // State/Input combination (e.g., "s3i5")
	Outputs   string // Circuit output values
	NextState string // Resulting next state
}

// SIMapping maps S/I strings to their corresponding fault patterns and next states
// Exported for use in other packages
type SIMapping struct {
	FP int // Fault pattern number
	NS Nd  // Next state node
}

// SimResultWithState extends SimResult with additional state information
type SimResultWithState struct {
	SI        string
	Outputs   string
	NextState string
	First     bool
	Ns16h     bool
	Ns8h      bool
	Ns4h      bool
	Ns2h      bool
	Ns1h      bool
}

// Global state variables for circuit simulation and phase management
var (
	// User-selected S/I sequence for testing
	AccumulatedSIs []string

	// Stored simulation results
	AccumulatedFaultFreeSimulations []SimResult
	AccumulatedFaultASimulations    []SimResult

	// Primary fault being tested
	OriginalFaultA string

	// First simulation timeframe flag
	First bool = true

	// Next state holding registers
	Ns16h, Ns8h, Ns4h, Ns2h, Ns1h bool

	// Current fault designations
	FaultA, FaultC string

	// Adaptive fault elimination variables
	CurrentFaultSet   []string // Working fault set that gets progressively filtered
	AllPossibleFaults []string // Original complete fault set (never modified)
)
