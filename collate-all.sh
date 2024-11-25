#!/bin/bash

echo "Collating all rounds..."

echo "Collating round 1..."
./scripts/collate.sh round1
echo "Collating round 1 winners..."
./scripts/collate.sh round1-winners

echo "Collating round 2..."
./scripts/collate.sh round2
echo "Collating round 2 winners..."
./scripts/collate.sh round2-winners

echo "Collating round 3..."
./scripts/collate.sh round3
echo "Collating round 3 winners..."
./scripts/collate.sh round3-winners

echo "Collating round 4..."
./scripts/collate.sh round4
echo "Collating round 4 winners..."
./scripts/collate.sh round4-winners

echo "Done!"
