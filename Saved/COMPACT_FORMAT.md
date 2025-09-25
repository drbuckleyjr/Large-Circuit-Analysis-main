# Compact Netlist Format for Circuit Design

## Overview
The compact netlist format provides a **designer-friendly** way to specify circuit netlists without the verbosity of JSON. Each gate is defined on a single line with all necessary information.

## Format Specification

### Line Format
```
<sig_name, sig_type, level, [inputs], [outputs], [targets]>
```

### Field Definitions

| Field | Description | Example |
|-------|-------------|---------|
| `sig_name` | Signal/gate name | `b25`, `out4`, `ns16` |
| `sig_type` | Gate type | `PI`, `NOT`, `2AND`, `3AND`, `2OR`, `3OR` |
| `level` | Topological level (1-N) | `1`, `5`, `10` |
| `[inputs]` | Input signal names | `[s23_s, a4_s, s23_1, a4_1]` |
| `[outputs]` | Output signal names | `[b25_s, b25_1]` |
| `[targets]` | Fault injection targets | `[b25:1, b25:0]` or `[b25]` |

### Gate Types

| Type | Description | Inputs | Outputs |
|------|-------------|--------|---------|
| `PI` | Primary Input | None | `[name_s, name_1]` |
| `NOT` | Inverter | `[in_s, in_1]` | `[out_s, out_1]` |
| `2AND` | 2-input AND | `[in1_s, in2_s, in1_1, in2_1]` | `[out_s, out_1]` |
| `3AND` | 3-input AND | `[in1_s, in2_s, in3_s, in1_1, in2_1, in3_1]` | `[out_s, out_1]` |
| `2OR` | 2-input OR | `[in1_s, in2_s, in1_1, in2_1]` | `[out_s, out_1]` |
| `3OR` | 3-input OR | `[in1_s, in2_s, in3_s, in1_1, in2_1, in3_1]` | `[out_s, out_1]` |

## Examples

### Primary Input
```
ps16, PI, 1, [], [ps16_s, ps16_1], [ps16]
```
- Creates primary input `ps16` at level 1
- Outputs fault propagation signal `ps16_s` and path enabling signal `ps16_1`
- Fault target is `ps16`

### Inverter
```
nps16, NOT, 1, [ps16_s, ps16_1], [nps16_s, nps16_1], [nps16]
```
- Creates inverter `nps16` at level 1
- Takes `ps16_s` and `ps16_1` as inputs
- Produces `nps16_s` and `nps16_1` as outputs

### 3-input AND Gate
```
ls0, 3AND, 2, [nps4_s, nps2_s, nps1_s, nps4_1, nps2_1, nps1_1], [ls0_s, ls0_1], [ls0]
```
- Creates 3-AND gate `ls0` at level 2
- ANDs three fault signals and three path enabling signals
- Output is `ls0_s` (fault) and `ls0_1` (path enable)

### 2-input OR Gate  
```
out4, 2OR, 10, [f1_s, f2_s, f1_1, f2_1], [out4_s, out4_1], [out4]
```
- Creates primary output `out4` at final level 10
- ORs two intermediate signals
- This is a circuit output

## Topological Levels

Gates are organized by **topological level** indicating processing order:

- **Level 1**: Primary inputs and basic inversions
- **Level 2**: Local state generation from primary inputs
- **Level 3**: State combinations and additional logic
- **Levels 4-9**: Progressive logic reduction and combinations
- **Level 10**: Final outputs and next-state computation

**Processing Rule**: All gates at level N can be processed in any order because:
- Their inputs come from levels < N
- Their outputs are used at levels > N

## Signal Naming Convention

### Standard Signals
- **Fault propagation**: `name_s` (e.g., `ps16_s`, `out4_s`)
- **Path enabling**: `name_1` (e.g., `ps16_1`, `out4_1`)

### Signal Categories
- **Primary inputs**: `ps16`, `ps8`, `ps4`, `ps2`, `ps1` (state)
- **Circuit inputs**: `i7`, `i6`, `i5`, etc. (mapped to `in4`, `in2`, `in1`)
- **Outputs**: `out4_s`, `out2_s`, `out1_s`
- **Next states**: `ns16_s`, `ns8_s`, `ns4_s`, `ns2_s`, `ns1_s`

## Comments and Formatting

```
# This is a comment line
// This is also a comment

# Empty lines are ignored

# Brackets can be omitted (but recommended for clarity):
ps16, PI, 1, [], [ps16_s, ps16_1], [ps16]

# Whitespace is flexible:
ps16,PI,1,[],[ps16_s,ps16_1],[ps16]
```

## Advantages Over JSON

| Feature | Compact Format | JSON Format |
|---------|---------------|-------------|
| **Readability** | ✅ Single line per gate | ❌ Multiple lines, nested structure |
| **Editability** | ✅ Easy to modify individual gates | ❌ Must navigate complex structure |
| **Compactness** | ✅ ~80% smaller file size | ❌ Very verbose |
| **Designer Input** | ✅ Natural for manual creation | ❌ Too complex for hand editing |
| **Sorting** | ✅ Easy to sort by level or name | ❌ Requires special tools |
| **Processing** | ✅ Converts to JSON internally | ✅ Native JSON processing |

## Conversion Workflow

```
Designer Input (Compact) → JSON (Internal) → ps2ns Generation
                        ↓
                   Validation & Analysis
```

1. **Designer creates** compact netlist file
2. **System converts** to JSON for internal processing  
3. **Validation framework** tests correctness
4. **ps2ns generator** creates executable code

## Tool Support

The system provides:
- **Parser**: `ParseCompactNetlist()` - Load compact format
- **Converter**: `ToJSON()` - Convert to internal JSON format
- **Exporter**: `ExportToFile()` - Save compact format
- **Sorter**: `SortByLevel()`, `SortByName()` - Organization tools
- **Round-trip**: Full conversion cycle validation

## Best Practices

### For Designers
1. **Start with level 1** - Define all primary inputs first
2. **Maintain level order** - Process dependencies correctly
3. **Use consistent naming** - Follow `name_s`/`name_1` convention
4. **Group by level** - Keep same-level gates together
5. **Add comments** - Document complex logic sections

### For Tools
1. **Validate dependencies** - Ensure inputs exist before use
2. **Check level consistency** - No forward references
3. **Verify completeness** - All signals properly defined
4. **Support sorting** - Allow reorganization by level/name

This compact format makes circuit specification **designer-friendly** while maintaining full compatibility with the automated processing system!
