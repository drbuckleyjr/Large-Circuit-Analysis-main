package main

import (
	"fmt"
)

// Demo the compact format conversion
func RunCompactDemo() {
	fmt.Println("=== COMPACT NETLIST FORMAT CONVERSION ===")

	// Convert existing JSON to compact format
	err := ConvertJSONToCompact("netlist_large.json", "netlist_large.compact")
	if err != nil {
		fmt.Printf("Conversion error: %v\n", err)
		return
	}

	// Also convert the small circuit
	err = ConvertJSONToCompact("netlist_small.json", "netlist_small.compact")
	if err != nil {
		fmt.Printf("Small circuit conversion error: %v\n", err)
		// Continue even if small circuit fails
	}

	// Demonstrate the format
	DemoCompactFormat()

	fmt.Println("\n=== TESTING ROUND-TRIP CONVERSION ===")

	// Test round-trip: Compact -> JSON -> Compact
	fmt.Println("1. Loading compact format...")
	compactNetlist, err := ParseCompactNetlist("netlist_large.compact")
	if err != nil {
		fmt.Printf("Error parsing compact format: %v\n", err)
		return
	}

	fmt.Printf("   Loaded %d gates\n", len(compactNetlist.Gates))

	fmt.Println("2. Converting to JSON format...")
	jsonCircuit, err := compactNetlist.ToJSON()
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}

	fmt.Printf("   Created circuit with %d levels\n", len(jsonCircuit.Levels))

	fmt.Println("3. Converting back to compact format...")
	compactNetlist2 := JSONToCompact(jsonCircuit)
	compactNetlist2.SortByLevel()

	fmt.Printf("   Round-trip complete: %d gates\n", len(compactNetlist2.Gates))

	fmt.Println("\n✅ Compact format system ready for designer input!")
}
