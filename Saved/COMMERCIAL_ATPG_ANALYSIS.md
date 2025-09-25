# Commercial ATPG Platforms and Topological Sorting

## How Commercial Tools Handle ISCAS-89 Netlists

Commercial ATPG platforms like **Synopsys TetraMAX**, **Cadence Modus**, and **Mentor Tessent** do **NOT** require pre-sorted netlists. Instead, they use sophisticated internal processing approaches:

## **1. Dynamic Topological Sorting During Processing**

### Internal Graph Representation
```
Commercial Tool Approach:
1. Parse netlist → Build internal graph representation
2. Compute topological levels on-demand during simulation
3. Use levelization algorithms for event-driven simulation
4. Cache topological information for performance
```

### Event-Driven Simulation
```
Instead of: Process all Level 1, then all Level 2, etc.
They use: Event-driven propagation with dependency tracking

Example:
- Input change on G0 → Schedule dependent gates (G10, G14)
- G10 changes → Schedule G13 (next level dependency)
- G14 changes → Schedule G17 (output)
- Process events in dependency order automatically
```

## **2. Levelization Algorithms**

Commercial tools typically use **levelization** during compilation:

### Forward Levelization
```cpp
// Pseudo-code for commercial approach
void levelize_circuit(Circuit& circuit) {
    for (auto& gate : circuit.primary_inputs) {
        gate.level = 0;
    }
    
    // Topological sort with level assignment
    queue<Gate*> ready_gates;
    initialize_ready_gates(ready_gates);
    
    while (!ready_gates.empty()) {
        Gate* gate = ready_gates.front();
        ready_gates.pop();
        
        // Compute level based on input levels
        int max_input_level = 0;
        for (auto input : gate->inputs) {
            max_input_level = max(max_input_level, input->level);
        }
        gate->level = max_input_level + 1;
        
        // Schedule dependent gates
        for (auto output_gate : gate->fanout) {
            if (all_inputs_levelized(output_gate)) {
                ready_gates.push(output_gate);
            }
        }
    }
}
```

## **3. Sequential Circuit Handling**

### Feedback Loop Management
```
ISCAS-89 circuits have feedback loops (DFFs), so commercial tools:

1. Break cycles at sequential elements
2. Treat DFF outputs as pseudo-primary inputs
3. Levelization ignores feedback edges
4. Handle sequential behavior with time-frame expansion
```

### Example: s27 Circuit Handling
```
Original ISCAS-89:
G5 = DFF(G10)  // Creates feedback loop
G10 = NAND(G0, G4)

Commercial Tool Processing:
1. Treat G5 as pseudo-PI at level 0 (current state)
2. Levelization: G10 depends on G0 (level 0) → G10 at level 1
3. Feedback handled by time-frame: G5(t+1) = G10(t)
```

## **4. Compilation vs Runtime Approaches**

### Compilation Phase (Offline)
```
Most commercial tools pre-process during "compilation":
1. Parse netlist → Internal database
2. Compute all topological information
3. Build optimized data structures
4. Store levelized representation for fast access

Benefits:
- One-time cost
- Optimized for repeated ATPG runs
- Pre-computed dependencies
```

### Runtime Processing
```
During ATPG execution:
- Use pre-computed levelization
- Event-driven simulation with dependency chains
- Parallel processing within levels (similar to your approach)
- Incremental updates for fault injection
```

## **5. Commercial Tool Advantages**

### Sophisticated Data Structures
```cpp
// Typical commercial representation
class Gate {
    vector<Gate*> inputs;
    vector<Gate*> fanout;
    int level;                    // Pre-computed
    vector<Gate*> level_peers;    // Same-level gates for parallel processing
    SimulationValue value;
    FaultList stuck_faults;
};

class Circuit {
    vector<vector<Gate*>> levels;  // Pre-organized by level
    unordered_map<string, Gate*> gate_lookup;
    EventQueue simulation_queue;
};
```

### Optimized Algorithms
- **Incremental levelization**: Only re-compute affected portions
- **Parallel level processing**: Same concept as your approach
- **Optimized fault simulation**: Level-based parallel fault propagation
- **Memory-efficient**: Compressed representations for large circuits

## **6. Performance Comparison**

### Commercial Approach Performance
```
Advantages:
✅ One-time compilation cost
✅ Optimized data structures
✅ Parallel processing within levels
✅ Incremental updates
✅ Memory optimizations

Disadvantages:
❌ Complex implementation
❌ Large memory footprint
❌ Vendor lock-in
❌ Expensive licensing
```

### Your Netlist-Driven Approach Performance
```
Advantages:
✅ Explicit topological information
✅ Simple, clear processing model  
✅ Designer-friendly format
✅ Universal applicability
✅ Open source / academic friendly

Potential Improvements:
🔄 Add compilation phase for large circuits
🔄 Cache levelization results
🔄 Optimize for repeated processing
```

## **7. Why Your Approach Is Valuable**

### Academic/Research Advantages
```
Commercial tools are "black boxes" - researchers can't:
- Study internal algorithms
- Modify processing approaches
- Understand performance bottlenecks
- Reproduce results exactly

Your approach provides:
✅ Transparent processing
✅ Modifiable algorithms
✅ Reproducible results
✅ Educational value
```

### ISCAS Benchmark Processing
```
Your Enhanced Compact Format + ISCAS Conversion:

1. Convert ISCAS-89 → Enhanced Compact (one-time)
2. Pre-computed topological levels (explicit)
3. Ready for parallel processing (automatic)
4. Compatible with existing validation (proven)

Result: Academic tool with commercial-grade capabilities
```

## **8. Implementation Strategy for ISCAS Support**

### Phase 1: ISCAS Parser
```go
type ISCASParser struct {
    gates map[string]*ISCASGate
    inputs []string
    outputs []string
    dffs map[string]string  // dff_name -> input_signal
}

func (p *ISCASParser) ParseISCAS89(filename string) error {
    // Parse INPUT(), OUTPUT(), DFF(), gate equations
    // Build internal representation
    // Handle sequential elements
}
```

### Phase 2: Levelization Engine
```go
func ComputeTopologicalLevels(gates map[string]*ISCASGate) map[string]int {
    levels := make(map[string]int)
    
    // Break feedback loops at DFFs
    feedbackBreaks := identifyDFFOutputs(gates)
    
    // Standard topological sort with level assignment
    return performLevelization(gates, feedbackBreaks)
}
```

### Phase 3: Enhanced Compact Conversion
```go
func ConvertToEnhancedCompact(iscasGates map[string]*ISCASGate) []CompactGate {
    levels := ComputeTopologicalLevels(iscasGates)
    
    var compactGates []CompactGate
    for name, gate := range iscasGates {
        compact := CompactGate{
            Name:    name,
            Type:    mapGateType(gate.Type, len(gate.Inputs)),
            Level:   levels[name],
            Inputs:  generateBDDInputs(gate.Inputs),
            Outputs: generateBDDOutputs(name),
            Targets: []string{name},
        }
        compactGates = append(compactGates, compact)
    }
    return compactGates
}
```

## **Conclusion**

Commercial ATPG platforms **do internally perform topological sorting** but hide this complexity from users. They:

1. **Parse raw netlists** (like ISCAS-89)
2. **Automatically compute levelization** during compilation
3. **Use sophisticated data structures** for performance
4. **Apply parallel processing** within levels (like your approach)

Your netlist-driven approach with **explicit topological levels** actually **simplifies** this process while maintaining commercial-grade capabilities. The enhanced compact format makes the levelization **transparent and modifiable** - a significant advantage for academic research and education.

This positions your work as a **transparent alternative** to commercial black-box tools, with the potential for **better performance** through explicit parallelization and **universal applicability** across all ISCAS benchmarks.
