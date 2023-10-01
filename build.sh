#!/bin/bash

export PROD_PATH="/opt/musematch"

# pull latest version
git pull

# copy template files
cp -r views "$PROD_PATH/"
cp -r public "$PROD_PATH/"
cp .env "$PROD_PATH/.env"

# build tailwindcss styles
sudo ~/tailwindcss -i "$PROD_PATH/views/input.css" -o "$PROD_PATH/public/output.css" --minify

# build server binary
go build -o "$PROD_PATH/server"

# restart server daemon
sudo systemctl restart musematch
