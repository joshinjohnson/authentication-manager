run:
	echo "starting server on localhost:9090"
	go run main/main.go

register:
	curl --data "Email=value1&Password-Hash=value2&First-Name=asd&Last-Name=asd" localhost:9090/register

login:
	curl --data "Email=value1&Password-Hash=value2" localhost:9090/login
