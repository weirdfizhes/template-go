# {Project Name} - Golang

This is {Project Name} Application for End-User App.

Tech Stack:
1. App
    - Go (minimum go1.19)
2. Core Package
    - Echo
    - SQLX
3. Database
    - PostgreSQL (v14.7)
4. Runner
    - SocketMaster
    - Nodemon

## SQLC Migrate and Generate Models
### Generate migration file
```
$ make gen.pg.migration name=<migration_name>
```
### Migrate up
```
$ make pg.migrate.up
```
### Migrate down
```
$ make pg.migrate.down
```
## Application Manual
### Install dependencies

    $ make deps

### Make env file

    $ make env

### Run development server

    $ make dev

### Build the goApp executable file

    $ make build.go

### Start goApp

    $ make start

### Start goApp binary

    $ make start.binary

### Start goApp with SocketMaster

    $ make start.sm

# REST API Documentation

REST API Documentation for {Project Name} Application is stored in Postman App, or in `docs/postman` folder, and you can download the JSON file here.

<!-- Change link to your postman link -->
[Download here!](https://postman.com)

<br>


![PLABS.ID LOGO](https://www.plabs.id/LogoPlabs.svg "https://www.plabs.id")

&copy; 2023 PLABS.ID. All Rights Reserved
