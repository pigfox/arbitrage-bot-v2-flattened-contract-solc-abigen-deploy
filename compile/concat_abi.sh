#!/bin/sh
output_dir="/home/peter/Documents/crypto-projects/arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/solc-output"
output_file="Pigfox_combined.abi"

echo "[" > "${output_dir}/${output_file}"
first=1
for file in "${output_dir}"/*.abi; do
    if [ "$first" -eq 1 ]; then
        cat "$file" >> "${output_dir}/${output_file}"
        first=0
    else
        echo "," >> "${output_dir}/${output_file}"
        cat "$file" >> "${output_dir}/${output_file}"
    fi
done
echo "]" >> "${output_dir}/${output_file}"
mv Pigfox_combined.abi Pigfox.abi
