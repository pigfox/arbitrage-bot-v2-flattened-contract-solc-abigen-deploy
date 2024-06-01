#!/bin/sh
set -x
set -e
echo "Compiling..."
echo "Contract: $1"
echo "Version: $2"
projectDir=$(pwd)
contractDir="/sources/sol-flattened"
compiledDir="/sources/solc-output"
docker run -v "$projectDir":/sources ethereum/solc:"$2" --allow-paths="$contractDir" -o "$compiledDir" --optimize --overwrite --bin --abi "$contractDir/$1.sol"
# Run abigen on the host system
abigen --abi "$projectDir/solc-output/$1.abi" --bin="$projectDir/solc-output/$1.bin" --pkg=api --type="$1" --out="$projectDir/api/$1.go"

ls -l "$projectDir/solc-output"
ls -l "$projectDir/api"
echo "Done Compiling..."