# Large-Circuit-Analysis-main

This repository contains the implementation of Automatic Test Pattern Generation (ATPG) algorithms for large-scale circuit analysis.

## Project Structure

```
├── src/                    # Source code
│   ├── circuit_analyzer.py # Main ATPG engine
│   └── circuit_parser.py   # Circuit format parser
├── docs/                   # Documentation
│   └── project_overview.md # Project overview and architecture
├── tests/                  # Unit tests
│   └── test_circuit_parser.py # Parser tests
├── requirements.txt        # Python dependencies
└── README.md              # This file
```

## Features

- **Multi-format support**: Parse Verilog and SPICE circuit descriptions
- **ATPG algorithms**: Framework for implementing various test pattern generation algorithms
- **Fault coverage analysis**: Tools for analyzing test pattern effectiveness
- **Extensible architecture**: Modular design for easy extension and customization

## Quick Start

1. Clone the repository
2. Install dependencies: `pip install -r requirements.txt`
3. Run tests: `python -m pytest tests/`
4. See `docs/project_overview.md` for detailed documentation

## Current Status

This project is in active development. Core infrastructure and parsing capabilities have been implemented, with ATPG algorithm development in progress.
