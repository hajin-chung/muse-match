# Muse Match

golang (fiber, slqx) + sqlite3 + html templates

# TODO

- [x] setup ec2
- [x] setup daemon server service
- [x] migrate r2 to s3
- [x] deploy github actions
- [x] auth
- [x] edit profile
- [x] user art list
- [x] create new art
- [x] exhibit page
- [x] admin page
- [x] edit art
- [x] art url qr code
- [x] proper error handling
  - message to slack
  - logging errors with environment
  - refine error page (500 status page)
- [x] handle 404 not found
- [x] copy data from astro server
- [ ] production ready db
  - automatical backup script
- [x] health check
- [x] tailwindcss build step

# production setting

0. install requirements
  
  - go
  - sqlite3
  - tailwindcss cli

1. create database file

```
mkdir $PROD_PATH/db
sqlite3 $PROD_PATH/db/musematch.db < db/init.sql
```

2. run build script

```
./build.sh
```
