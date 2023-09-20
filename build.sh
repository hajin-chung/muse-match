#!/bin/bash

# TODO: change this path to outside git directory
PROD_PATH="./prod"

# copy template files
cp -r views "$PROD_PATH/"

# build server binary
go build -o "$PROD_PATH/server"

# reload server daemon
sudo systemctl reload musematch
