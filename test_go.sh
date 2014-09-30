#!/bin/bash
./load_schema.sh dwarfomatic_test
~/save/go/bin/go test $@
