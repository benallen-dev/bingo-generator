#!/bin/bash

echo "Collating all rounds..."
for dir in ./output/*/; do
	echo "$dir"
    ./scripts/collate.sh "$(basename "$dir")"
done


# echo "Collating round 1..."
# ./scripts/collate.sh round-1
# echo "Collating round 1 winners..."
# ./scripts/collate.sh round-1-winners

# echo "Collating round 2..."
# ./scripts/collate.sh round-2
# echo "Collating round 2 winners..."
# ./scripts/collate.sh round-2-winners

# echo "Collating round 3..."
# ./scripts/collate.sh round-3
# echo "Collating round 3 winners..."
# ./scripts/collate.sh round3-winners

# echo "Collating round 4..."
# ./scripts/collate.sh round4
# echo "Collating round 4 winners..."
# ./scripts/collate.sh round4-winners

echo "Done!"
