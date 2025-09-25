# Large Circuit Analysis - ATPG Project

## Overview

This project implements Automatic Test Pattern Generation (ATPG) algorithms for large-scale circuit analysis. The system is designed to handle complex digital circuits and generate comprehensive test patterns for fault detection and coverage analysis.

## Project Goals

- **Large-scale circuit support**: Handle circuits with thousands of gates
- **Multiple format support**: Parse Verilog, SPICE, and other circuit description formats
- **Advanced ATPG algorithms**: Implement D-algorithm, PODEM, and other modern ATPG techniques
- **Fault coverage analysis**: Comprehensive reporting of test pattern effectiveness
- **Performance optimization**: Efficient algorithms for large circuit analysis

## Architecture

### Core Components

1. **Circuit Parser** (`src/circuit_parser.py`)
   - Supports multiple circuit description formats
   - Validates circuit connectivity and structure
   - Converts to internal representation

2. **Circuit Analyzer** (`src/circuit_analyzer.py`)
   - Main ATPG engine
   - Test pattern generation
   - Fault coverage analysis

3. **Test Patterns** (Planned)
   - Pattern optimization
   - Compression techniques
   - Export to various test formats

## Current Status

- ✅ Basic project structure established
- ✅ Circuit parser foundation implemented
- ✅ Core analyzer framework created
- 🔄 ATPG algorithm implementation in progress
- ⏳ Test suite development planned
- ⏳ Performance benchmarking planned

## Usage

```python
from src.circuit_analyzer import CircuitAnalyzer

# Initialize analyzer
analyzer = CircuitAnalyzer()

# Load circuit
analyzer.load_circuit("path/to/circuit.v")

# Generate test patterns
patterns = analyzer.generate_test_patterns("circuit.v")

# Analyze coverage
coverage = analyzer.analyze_fault_coverage(patterns)
```

## Dependencies

- Python 3.7+
- Standard library modules (re, typing)

## Development Roadmap

1. **Phase 1**: Core infrastructure and parsing
2. **Phase 2**: Basic ATPG algorithm implementation
3. **Phase 3**: Advanced algorithms and optimization
4. **Phase 4**: Large-scale testing and benchmarking
5. **Phase 5**: Performance optimization and deployment

## Contributing

This project is part of ongoing research into large-scale circuit analysis and ATPG methodologies.