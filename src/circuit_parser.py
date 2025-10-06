#!/usr/bin/env python3
"""
Circuit Parser for processing various circuit description formats
Part of the Large Circuit Analysis ATPG project
"""

import re
from typing import Dict, List, Tuple

class CircuitParser:
    """Parser for circuit description files (Verilog, SPICE, etc.)"""
    
    def __init__(self):
        self.nodes = {}
        self.components = []
        self.nets = []
    
    def parse_verilog(self, content: str) -> Dict:
        """Parse Verilog circuit description."""
        circuit_info = {
            'modules': [],
            'inputs': [],
            'outputs': [],
            'wires': [],
            'gates': []
        }
        
        lines = content.strip().split('\n')
        current_module = None
        
        for line in lines:
            line = line.strip()
            if not line or line.startswith('//'):
                continue
                
            # Module declaration with port list
            if line.startswith('module'):
                module_match = re.match(r'module\s+(\w+)', line)
                if module_match:
                    current_module = module_match.group(1)
                    circuit_info['modules'].append(current_module)
                
                # Parse ports from module declaration
                if '(' in line:
                    port_section = line.split('(')[1].split(')')[0] if ')' in line else line.split('(')[1]
                    
                    # Split into tokens and process sequentially
                    tokens = [t.strip() for t in re.split(r'[,\s]+', port_section) if t.strip()]
                    i = 0
                    while i < len(tokens):
                        if tokens[i] == 'input':
                            i += 1
                            while i < len(tokens) and tokens[i] not in ['input', 'output']:
                                circuit_info['inputs'].append(tokens[i])
                                i += 1
                        elif tokens[i] == 'output':
                            i += 1
                            while i < len(tokens) and tokens[i] not in ['input', 'output']:
                                circuit_info['outputs'].append(tokens[i])
                                i += 1
                        else:
                            i += 1
            
            # Standalone Input/Output declarations
            elif line.startswith('input'):
                inputs = re.findall(r'\w+', line[5:])
                circuit_info['inputs'].extend(inputs)
            
            elif line.startswith('output'):
                outputs = re.findall(r'\w+', line[6:])
                circuit_info['outputs'].extend(outputs)
            
            elif line.startswith('wire'):
                wires = re.findall(r'\w+', line[4:])
                circuit_info['wires'].extend(wires)
        
        return circuit_info
    
    def parse_spice(self, content: str) -> Dict:
        """Parse SPICE netlist format."""
        circuit_info = {
            'title': '',
            'components': [],
            'nodes': set()
        }
        
        lines = content.strip().split('\n')
        
        for i, line in enumerate(lines):
            line = line.strip()
            if not line or line.startswith('*'):
                continue
            
            if i == 0:
                circuit_info['title'] = line
                continue
            
            # Parse component lines
            if line.startswith(('R', 'C', 'L', 'M', 'Q')):
                parts = line.split()
                if len(parts) >= 3:
                    component = {
                        'name': parts[0],
                        'nodes': parts[1:-1],
                        'value': parts[-1]
                    }
                    circuit_info['components'].append(component)
                    circuit_info['nodes'].update(parts[1:-1])
        
        circuit_info['nodes'] = list(circuit_info['nodes'])
        return circuit_info
    
    def validate_circuit(self, circuit_info: Dict) -> Tuple[bool, List[str]]:
        """Validate parsed circuit for common issues."""
        errors = []
        
        # Check for disconnected nodes
        # Check for short circuits
        # Validate component connections
        
        is_valid = len(errors) == 0
        return is_valid, errors

if __name__ == "__main__":
    parser = CircuitParser()
    print("Circuit Parser module loaded")