module github.com/joshinjohnson/authentication-manager

go 1.16

replace (
	github.com/joshinjohnson/authentication-engine v0.1.1 => ../authentication-engine
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/joshinjohnson/authentication-engine v0.1.1
)
