# ISCAS-89 vs Enhanced Compact Format Comparison

## ISCAS-89 Sequential Benchmark Format Analysis

Based on the ISCAS-89 benchmark suite, here's how their netlist format compares to your enhanced compact format:

## **ISCAS-89 Original Format (Typical)**

### Example from s27 benchmark:
```
# s27 - Sequential benchmark with 4 inputs, 1 output, 3 flip-flops
INPUT(G0)
INPUT(G1) 
INPUT(G2)
INPUT(G3)
OUTPUT(G17)

G5 = DFF(G10)
G6 = DFF(G11)
G7 = DFF(G13)

G8 = NOT(G5)
G9 = NOT(G6)
G10 = NAND(G0, G4)
G11 = OR(G1, G9)
G12 = OR(G2, G8)
G13 = NAND(G3, G12)
G14 = NAND(G0, G8)
G15 = NAND(G1, G9)
G16 = NAND(G2, G5)
G17 = OR(G14, G15, G16)
```

### Example from s208 benchmark:
```
INPUT(G1)
INPUT(G2)
INPUT(G3)
...
OUTPUT(G199)
OUTPUT(G200)

G45 = DFF(G10)
G46 = DFF(G11)
...
G50 = AND(G1, G8)
G51 = AND(G45, G47)
G52 = NAND(G2, G46)
...
```

## **Your Enhanced Compact Format**

### Same s27 circuit in your format:
```
# Level 1 - Primary inputs
G0, PI, 1, [], [G0_s, G0_1], [G0]
G1, PI, 1, [], [G1_s, G1_1], [G1] 
G2, PI, 1, [], [G2_s, G2_1], [G2]
G3, PI, 1, [], [G3_s, G3_1], [G3]

# Level 1 - Flip-flop outputs (pseudo-primary inputs)
G5, DFF, 1, [G10_s], [G5_s, G5_1], [G5]
G6, DFF, 1, [G11_s], [G6_s, G6_1], [G6]
G7, DFF, 1, [G13_s], [G7_s, G7_1], [G7]

# Level 2 - Combinational logic
G8, NOT, 2, [G5_s, G5_1], [G8_s, G8_1], [G8]
G9, NOT, 2, [G6_s, G6_1], [G9_s, G9_1], [G9]

# Level 3 - More combinational logic  
G10, NAND, 3, [G0_s, G4_s], [G10_s, G10_1], [G10]
G11, OR, 3, [G1_s, G9_s], [G11_s, G11_1], [G11]
G12, OR, 3, [G2_s, G8_s], [G12_s, G12_1], [G12]
G14, NAND, 3, [G0_s, G8_s], [G14_s, G14_1], [G14]
G15, NAND, 3, [G1_s, G9_s], [G15_s, G15_1], [G15]
G16, NAND, 3, [G2_s, G5_s], [G16_s, G16_1], [G16]

# Level 4 - Next state and outputs
G13, NAND, 4, [G3_s, G12_s], [G13_s, G13_1], [G13]
G17, OR3, 4, [G14_s, G15_s, G16_s], [G17_s, G17_1], [G17]
```

## **Format Comparison Analysis**

| Aspect | ISCAS-89 Original | Your Enhanced Compact |
|--------|------------------|----------------------|
| **Gate Specification** | `G10 = NAND(G0, G4)` | `G10, NAND, 3, [G0_s, G4_s], [G10_s, G10_1], [G10]` |
| **Topological Info** | ❌ None | ✅ Explicit level number |
| **BDD Signal Names** | ❌ Manual generation | ✅ Built-in (_s, _1 suffixes) |
| **Processing Order** | ❌ Must compute | ✅ Level-based processing ready |
| **Parallel Processing** | ❌ Requires analysis | ✅ Same-level gates clearly identified |
| **Input/Output Lists** | ❌ Implicit | ✅ Explicit arrays |
| **Tool Integration** | ❌ Requires parsing | ✅ JSON-compatible structure |

## **Key Similarities (Great News!)**

### 1. **Gate Types Are Compatible**
```
ISCAS-89: NAND, AND, OR, NOT, DFF
Your Format: 2AND, 3AND, 2OR, 3OR, NOT, DFF, PI
```
**Mapping is straightforward!**

### 2. **Signal Names Follow Patterns**
```
ISCAS-89: G0, G1, G10, G17
Your Format: ps16, ls0, s1, out4
```
**Both use meaningful signal names**

### 3. **Sequential Elements Are Identified**
```
ISCAS-89: G5 = DFF(G10)
Your Format: G5, DFF, 1, [G10_s], [G5_s, G5_1], [G5]
```
**Both clearly mark flip-flops**

## **Key Differences (Conversion Needed)**

### 1. **Topological Levels Missing in ISCAS**
- **ISCAS-89**: No level information
- **Your Format**: Explicit level numbers for processing order
- **Solution**: Compute topological levels during conversion

### 2. **BDD Signal Convention Different**
- **ISCAS-89**: Simple names (G0, G10)
- **Your Format**: BDD-ready names (G0_s, G0_1, G10_s, G10_1)
- **Solution**: Auto-generate BDD signal names

### 3. **Gate Types Need Mapping**
- **ISCAS-89**: Generic (AND, OR, NAND)
- **Your Format**: Arity-specific (2AND, 3AND, 2OR, 3OR)
- **Solution**: Analyze input count and map appropriately

## **Conversion Strategy**

### Phase 1: Parse ISCAS-89 Format
```go
type ISCASGate struct {
    Name     string
    Type     string  // "NAND", "AND", "OR", "NOT", "DFF"
    Inputs   []string
    Output   string
}
```

### Phase 2: Compute Topological Levels
```go
func computeTopologicalLevels(gates []ISCASGate) map[string]int {
    // Implement topological sort
    // Assign level numbers
    // Handle sequential feedback properly
}
```

### Phase 3: Convert to Enhanced Compact Format
```go
func convertToCompactFormat(iscasGates []ISCASGate) []CompactGate {
    levels := computeTopologicalLevels(iscasGates)
    
    var compactGates []CompactGate
    for _, gate := range iscasGates {
        compact := CompactGate{
            Name:    gate.Name,
            Type:    mapGateType(gate.Type, len(gate.Inputs)),
            Level:   levels[gate.Name],
            Inputs:  generateBDDInputs(gate.Inputs),
            Outputs: generateBDDOutputs(gate.Output),
            Targets: []string{gate.Name},
        }
        compactGates = append(compactGates, compact)
    }
    return compactGates
}
```

## **Conversion Complexity Assessment**

### ✅ **Easy Conversions**
- **Gate types**: Direct mapping with arity detection
- **Signal names**: Keep original names, generate BDD variants
- **Connectivity**: Direct translation of input/output relationships

### 🟡 **Medium Complexity**
- **Topological levels**: Requires graph analysis
- **Sequential elements**: Need to handle feedback loops properly
- **Primary inputs/outputs**: Extract from INPUT/OUTPUT declarations

### ⚠️ **Challenging Aspects**
- **Sequential feedback**: DFF outputs feeding back to combinational logic
- **Level assignment**: Sequential circuits have feedback cycles
- **State element handling**: Flip-flops break topological ordering

## **Expected Conversion Results**

### ISCAS-89 Benchmark Coverage
```
Small circuits (s27, s208): ~10-20 gates → Easy conversion
Medium circuits (s344, s349): ~50-100 gates → Straightforward
Large circuits (s1196, s1423): 500+ gates → Automated processing essential
```

### Performance Benefits
- **Manual ps2ns() elimination**: No more hand-coding hundreds of gates
- **Parallel processing**: Same-level gates identified automatically
- **Validation capability**: Compare against known ISCAS results
- **Academic credibility**: Standard benchmark compatibility

## **Conclusion**

The ISCAS-89 format is **very compatible** with your enhanced compact format! The main differences are:

1. **Missing topological levels** (computable)
2. **Different BDD naming** (auto-generatable)  
3. **Generic gate types** (mappable with arity analysis)

Your netlist-driven approach could easily support the entire ISCAS-89 benchmark suite, making it a **universal ATPG platform** rather than just a solution for your LARGE circuit.

This would be a **significant research contribution** - demonstrating automated ps2ns() generation for all standard ATPG benchmarks!
