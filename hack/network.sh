#!/bin/bash

# List all ips on the local network
if [ "$(uname)" == "Darwin" ]; then
    arp -a | grep -oh "\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*\)"    
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    arp -a | grep -oh "\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*\)"
fi
