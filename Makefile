version=0.1.0
.PHONY: test
test:
	make clean
	mkdir out
	go build -o out/cli main.go
	cd out \
	&& \
	./cli -m test-cli -p 1.0.1

.PHONY: clean
clean:
	rm -rf out

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o concise-cli-linux-amd64-${version} main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o concise-cli-windows-amd64-${version}.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o concise-cli-darwin-amd64-${version} main.go
