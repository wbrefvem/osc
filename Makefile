client: 
	go build -o osc ./client

server: 
	go build -o osc-server ./server

docker:
	docker build .

clean:
	rm -rf ./osc
	rm -rf ./osc-server