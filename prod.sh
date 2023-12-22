export PROD_PATH="/opt/musematch"

# pull latest version
git pull

# build
sh build.sh

# copy template files
sudo cp -r public "$PROD_PATH/"
sudo cp .env "$PROD_PATH/.env"
sudo cp musematch "$PROD_PATH/"

# restart server daemon
sudo systemctl restart musematch
