#!/bin/bash
tail -500 /home/towski/save/df_linux/gamelog.txt | sed '/^x.*/d' | tac > /home/towski/code/dwarfomatic/public/gamelog.txt
