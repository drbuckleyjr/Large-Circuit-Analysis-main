#!/bin/bash
# Code Lock-Down Script for Backup.go
# Creates timestamped backups and maintains a clean working version

echo "=== Code Lock-Down Script ==="
echo "Creating backup with timestamp..."

# Create timestamped backup
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
cp Backup.go "Backup_LOCKED_${TIMESTAMP}.go"

echo "Backup created: Backup_LOCKED_${TIMESTAMP}.go"
echo "File size: $(wc -l Backup.go | awk '{print $1}') lines"

# Show current status
echo ""
echo "Current backup files:"
ls -la Backup_*.go | tail -5

echo ""
echo "=== LOCK-DOWN COMPLETE ==="
echo "Your code is now safely backed up!"
echo ""
echo "To restore from backup:"
echo "cp Backup_LOCKED_${TIMESTAMP}.go Backup.go"
