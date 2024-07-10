build:
	go build -o runners-app main.go

docker-build:
	docker build -t runners-app:latest .  

test:
	go test -v

run: 
	go run main.go

docker-run:
	docker run --name runners-app -p 8080:8080 --mount type=bind,src=./runners.toml,dst=/app/runners.toml runners-app:latest

run-sample:
	cp sample-config-runners.toml runners.toml
	go run main.go

coverage:
	go test -v -coverpkg=./.. ./... -coverprofile=coverage.out

view-coverage:
	go tool cover -html=coverage.out

clean:
	rm runners-app &2>1
	rm coverage.out
