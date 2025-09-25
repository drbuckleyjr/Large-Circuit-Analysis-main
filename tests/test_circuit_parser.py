#!/usr/bin/env python3
"""
Unit tests for Circuit Parser module
"""

import unittest
import sys
import os

# Add src directory to path for imports
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'src'))

from circuit_parser import CircuitParser

class TestCircuitParser(unittest.TestCase):
    """Test cases for CircuitParser class"""
    
    def setUp(self):
        """Set up test fixtures before each test method."""
        self.parser = CircuitParser()
    
    def test_verilog_module_parsing(self):
        """Test parsing of basic Verilog module declaration."""
        verilog_content = """
        module test_circuit(input a, b, output y);
            wire w1;
            and gate1(w1, a, b);
            or gate2(y, w1, a);
        endmodule
        """
        
        result = self.parser.parse_verilog(verilog_content)
        
        self.assertIn('test_circuit', result['modules'])
        self.assertIn('a', result['inputs'])
        self.assertIn('b', result['inputs'])
        self.assertIn('y', result['outputs'])
        self.assertIn('w1', result['wires'])
    
    def test_spice_parsing(self):
        """Test parsing of basic SPICE netlist."""
        spice_content = """
        Test Circuit
        R1 1 2 1k
        R2 2 0 2k
        C1 1 0 1u
        .end
        """
        
        result = self.parser.parse_spice(spice_content)
        
        self.assertEqual(result['title'], 'Test Circuit')
        self.assertEqual(len(result['components']), 3)
        self.assertIn('1', result['nodes'])
        self.assertIn('2', result['nodes'])
        self.assertIn('0', result['nodes'])
    
    def test_empty_content(self):
        """Test parser behavior with empty content."""
        result_verilog = self.parser.parse_verilog("")
        result_spice = self.parser.parse_spice("")
        
        self.assertEqual(len(result_verilog['modules']), 0)
        self.assertEqual(len(result_spice['components']), 0)
    
    def test_circuit_validation(self):
        """Test circuit validation functionality."""
        circuit_info = {
            'modules': ['test'],
            'inputs': ['a', 'b'],
            'outputs': ['y'],
            'wires': ['w1']
        }
        
        is_valid, errors = self.parser.validate_circuit(circuit_info)
        
        # Basic validation should pass for well-formed circuit info
        self.assertIsInstance(is_valid, bool)
        self.assertIsInstance(errors, list)

if __name__ == '__main__':
    unittest.main()