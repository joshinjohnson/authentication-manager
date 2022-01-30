# authentication-manager

## Installation
`go get -u github.com/joshinjohnson/authentication-manager`

## Run
- to start server, use `make run`
- to register an user, use `make register Email=<EMAIL> Password=<PASSWORD> First-Name=<FIRSTNAME> Last-Name=<LASTNAME>`
- to login, use `make login Email=<EMAIL> Password=<PASSWORD>`
- to check token, use `make check-token Token=<TOKEN>`
