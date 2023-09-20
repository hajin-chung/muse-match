#!/bin/bash

PROD_PATH = "/opt/musematch"

# copy template files
cp -r views "$PROD_PATH/"

# build server binary
go build -o "$PROD_PATH/server"

# reload server daemon
sudo systemctl reload musematch
