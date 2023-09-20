#!/bin/bash

export PROD_PATH="/opt/musematch"

# pull latest version
git pull

# copy template files
cp -r views "$PROD_PATH/"

# build server binary
go build -o "$PROD_PATH/server"

# restart server daemon
sudo systemctl restart musematch
