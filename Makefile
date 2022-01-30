run:
	echo "starting server on localhost:9090"
	go run main/main.go

register:
	curl -X POST \
          http://localhost:9090/register \
		  -H 'email: ${Email}' \
		  -H 'password-hash: ${Password}' \
		  -H 'first-name: ${First-Name}' \
		  -H 'last-name: ${Last-Name}'


login:
	curl -X POST \
          http://localhost:9090/login \
          -H 'email: ${Email}' \
          -H 'password-hash: ${Password}'

check-token:
	curl -X GET \
          http://localhost:9090/home \
          -H 'token: ${Token}'
