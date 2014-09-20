#!/bin/bash
STONESENSE=$(xwininfo -name "Stonesense" | head -2 | tail -1 | cut -f4 -d' ')
DWARFFORTRESS=$(xwininfo -name "Dwarf Fortress" | head -2 | tail -1 | cut -f4 -d' ')
echo "Starting with Stonesense window $STONESENSE and $DWARFFORTRESS"
LD_LIBRARY_PATH=/home/towski/save/df_linux/hack/ ./foo $STONESENSE $DWARFFORTRESS
