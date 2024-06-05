#!/bin/sh
set -x
set -e
echo "Script running as user: $(whoami)"
echo "Compiling..."
echo "Contract: $1"
echo "Version: $2"
echo "EVM: $3"
projectDir=$(pwd)
contractDir="/sources/sol-flattened"
compiledDir="/sources/solc-output"
sudo chown -R $USER:$USER "$projectDir/solc-output"
sudo chown -R $USER:$USER "$projectDir/api"
docker run -v "$projectDir":/sources ethereum/solc:"$2" --allow-paths="$contractDir" -o "$compiledDir" --optimize --optimize-runs 200 --evm-version "$3" --overwrite --bin --abi "$contractDir/$1.sol"


# Run abigen on the host system
abigen --abi "$projectDir/solc-output/$1.abi" --bin="$projectDir/solc-output/$1.bin" --pkg=api --type="$1" --out="$projectDir/api/$1.go"

ls -l "$projectDir/solc-output"
ls -l "$projectDir/api"