#!/bin/bash
echo "Launching Hue Emulator"
pushd ~/code/Hue-Emulator
git pull
java -jar HueEmulator-v0.8.jar &