#!/bin/bash

# Update package list
sudo apt-get update

# Install Go if not already installed
if ! command -v go &> /dev/null; then
    sudo apt-get install -y golang
fi

# Set GOPATH and update PATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Add GOPATH and PATH to .bashrc for persistence
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

# Download and build abigen
mkdir -p $HOME/go-ethereum
cd $HOME/go-ethereum
git clone https://github.com/ethereum/go-ethereum.git .
cd cmd/abigen
go build

# Check if abigen was built successfully
if [ -f "./abigen" ]; then
    echo "abigen built successfully"
    sudo cp ./abigen /usr/local/bin/
    echo "abigen copied to /usr/local/bin"
else
    echo "abigen build failed"
fi

# Verify installation
which abigen
abigen --version
