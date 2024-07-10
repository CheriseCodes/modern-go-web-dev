build:
	go build -o runners-app main.go

test:
	go test -v

run: 
	go run main.go

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
