#!/bin/bash

echo "Collating all rounds..."
for dir in ./output/*/; do
	echo Generating PDF for "$(basename "$dir")"
    ./scripts/collate.sh "$(basename "$dir")"
done
