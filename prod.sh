export PROD_PATH="/opt/musematch"

# pull latest version
git pull

# build
sh build.sh

# copy template files
cp -r public "$PROD_PATH/"
cp .env "$PROD_PATH/.env"
cp musematch "#PROD_PATH/"

# restart server daemon
sudo systemctl restart musematch
