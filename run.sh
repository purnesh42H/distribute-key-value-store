#!/bin/bash

for i in 0 1 2 3; do
    ./server 510 $i 4 & # value is passed as command-line parameter
done
