#!/bin/bash

set -e

echo "Checking 150M generated single threaded with ent..."
frandom | head -c 150M | ent

echo "Checking 150M generated multi threaded with ent..."
frandom -t 16 | head -c 150M | ent

echo "Checking single threaded generated with dieharder"
frandom | dieharder -g 200 -d 0
frandom | dieharder -g 200 -d 7
frandom | dieharder -g 200 -d 102

echo "Checking multi threaded generated with dieharder"
frandom -t 16 | dieharder -g 200 -d 0
frandom -t 16 | dieharder -g 200 -d 7
frandom -t 16 | dieharder -g 200 -d 102