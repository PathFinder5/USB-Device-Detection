build_server:
	go build server.go helper.go routes.go 

server_run:
	sudo ./server

client_run:
	go run client.go