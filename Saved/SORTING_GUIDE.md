# Netlist Sorting and Search Guide

## Overview
The enhanced compact netlist system supports **two primary display modes** optimized for different use cases in the design workflow.

## Two Essential Sorting Modes

### 1. Topological Sorting (Processing Order)
**Use Case**: Circuit processing, ps2ns generation, validation
**Sort Key**: Level number (primary) → Signal name (secondary)
**Optimization**: Sequential processing efficiency

```
# Level 1 - Primary inputs
ps16, PI, 1, [], [ps16_s, ps16_1], [ps16]
ps8, PI, 1, [], [ps8_s, ps8_1], [ps8]
nps16, NOT, 1, [ps16_s, ps16_1], [nps16_s, nps16_1], [nps16]

# Level 2 - Local states  
ls0, 3AND, 2, [nps4_s, nps2_s, nps1_s], [ls0_s, ls0_1], [ls0]
ls1, 3AND, 2, [nps4_s, nps2_s, ps1_s], [ls1_s, ls1_1], [ls1]

# Level 3 - State combinations
s0, 3AND, 3, [nps16_s, nps8_s, ls0_s], [s0_s, s0_1], [s0]
s1, 3AND, 3, [nps16_s, nps8_s, ls1_s], [s1_s, s1_1], [s1]
```

**Advantages**:
- ✅ **Processing ready**: Gates can be executed level by level
- ✅ **Dependency safe**: No forward references possible
- ✅ **Natural creation order**: Matches how designers think
- ✅ **Validation friendly**: Easy to verify topological correctness

### 2. Alphabetical Sorting (Lookup/Reference)
**Use Case**: Debugging, cross-referencing, gate lookup, documentation
**Sort Key**: Signal name alphabetically
**Optimization**: Fast signal location and reference

```
# Signals starting with 'A'
a1, 2AND, 4, [s10_s, i0_s], [a1_s, a1_1], [a1]
a2, 2AND, 4, [s15_s, i5_s], [a2_s, a2_1], [a2]

# Signals starting with 'L'  
ls0, 3AND, 2, [nps4_s, nps2_s, nps1_s], [ls0_s, ls0_1], [ls0]
ls1, 3AND, 2, [nps4_s, nps2_s, ps1_s], [ls1_s, ls1_1], [ls1]

# Signals starting with 'P'
ps16, PI, 1, [], [ps16_s, ps16_1], [ps16]
ps8, PI, 1, [], [ps8_s, ps8_1], [ps8]
```

**Advantages**:
- ✅ **Fast lookup**: Find any signal quickly
- ✅ **Cross-reference**: Easy to trace signal usage
- ✅ **Debugging**: Locate problematic gates by name
- ✅ **Documentation**: Better for reference materials

## Usage Patterns

### Designer Workflow
```go
// 1. Create netlist in natural topological order
netlist := NewEnhancedCompactNetlist()
netlist.LoadFromFile("designer_input.compact")

// 2. Validate topological correctness
err := netlist.ValidateTopologicalOrder()
if err != nil {
    log.Fatal("Topological error:", err)
}

// 3. Process in topological order
processingOrder := netlist.GetProcessingOrder()
for _, levelGates := range processingOrder {
    // Process all gates at this level (can be parallel)
    for _, gate := range levelGates {
        processGate(gate)
    }
}
```

### Debugging Workflow
```go
// Switch to alphabetical view for debugging
netlist.Display(DisplayAlphabetical)

// Fast signal lookup during debugging
gate, found := netlist.FindGateByName("problematic_signal")
if found {
    fmt.Printf("Found %s at level %d\n", gate.Name, gate.Level)
    // Examine inputs, outputs, etc.
}

// Export alphabetical reference
netlist.ExportSorted("circuit_reference.compact", DisplayAlphabetical)
```

## Fast Access Operations

### Level-Based Processing (O(1) level access)
```go
// Get all gates at a specific level instantly
level5Gates := netlist.GetGatesAtLevel(5)
fmt.Printf("Level 5 has %d gates\n", len(level5Gates))

// Process levels in order
levels := netlist.GetTopologicalLevels() // [1, 2, 3, ..., 10]
for _, level := range levels {
    gates := netlist.GetGatesAtLevel(level)
    // Process this level
}
```

### Signal Lookup (O(1) name access)
```go
// Find specific gates instantly
if gate, found := netlist.FindGateByName("out4"); found {
    fmt.Printf("Output gate at level %d\n", gate.Level)
}

// Trace signal dependencies
if gate, found := netlist.FindGateByName("a1"); found {
    fmt.Printf("Gate %s uses inputs: %v\n", gate.Name, gate.Inputs)
}
```

## File Organization

### Topological Export (for processing)
```
circuit_processing.compact
├── # Level 1 - Primary inputs
├── # Level 2 - Local states  
├── # Level 3 - State combinations
├── ...
└── # Level 10 - Final outputs
```

### Alphabetical Export (for reference)
```
circuit_reference.compact
├── # Signals starting with 'A'
├── # Signals starting with 'B'
├── # Signals starting with 'C'
├── ...
└── # Signals starting with 'Z'
```

## Performance Characteristics

| Operation | Topological Sort | Alphabetical Sort |
|-----------|-----------------|-------------------|
| **Sequential Processing** | ✅ Optimal | ❌ Requires reordering |
| **Signal Lookup** | ❌ Linear search | ✅ O(1) with index |
| **Level Access** | ✅ O(1) with index | ❌ Requires filtering |
| **Validation** | ✅ Natural order | ❌ Must check dependencies |
| **Debugging** | ❌ Hard to find signals | ✅ Easy signal location |

## Best Practices

### For Circuit Creation
1. **Start topological**: Create gates in natural dependency order
2. **Validate early**: Check topological order frequently during development
3. **Use level comments**: Group gates by level with clear headers
4. **Consistent naming**: Follow signal naming conventions

### For Circuit Analysis
1. **Switch to alphabetical**: Use alphabetical view for signal tracing
2. **Fast lookup**: Use FindGateByName() for debugging specific signals
3. **Level analysis**: Use GetGatesAtLevel() to understand circuit hierarchy
4. **Export both formats**: Keep both sorted versions for different uses

### For Processing Systems
1. **Load topological**: Always process in topological order
2. **Index both ways**: Build both level and name indices
3. **Validate dependencies**: Ensure no forward references
4. **Optimize access**: Use O(1) lookup operations

## Tool Commands

```go
// Basic operations
netlist.SortByTopologicalLevel()  // Natural processing order
netlist.SortBySignalName()        // Lookup/reference order

// Display modes
netlist.Display(DisplayTopological)  // Show in processing order
netlist.Display(DisplayAlphabetical) // Show in lookup order

// Fast access
gates := netlist.GetGatesAtLevel(5)           // All level 5 gates
gate, found := netlist.FindGateByName("a1")  // Find specific gate
levels := netlist.GetTopologicalLevels()     // [1,2,3,...,10]

// Validation and export
netlist.ValidateTopologicalOrder()  // Check dependencies
netlist.ExportSorted(file, mode)    // Save in optimal format
```

This dual-sorting system provides the **flexibility designers need** while maintaining the **processing efficiency** required for automated netlist-driven ps2ns generation!
