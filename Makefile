build:
	go build -o go-mod-example main.go
run:
	./go-mod-example
clean:
	rm go-mod-example
run-client:
	go run cmd/client/main.go
run-route-guide:
	go run cmd/route-guide/main.go