version=0.1.9

# 进行测试
.PHONY: test
test:
	make clean
	mkdir out
	go build -o out/cli cmd/concise/concisectl.go
	cd out \
	&& \
	./cli -m test-cli -p 1.0.1

# 清理测试产生的文件
.PHONY: clean
clean:
	rm -rf out

# 构建可执行程序
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "-X main.cliVersion=${version}" -o concise-cli-linux-amd64 cmd/concise/concisectl.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags "-X main.cliVersion=${version}" -o concise-cli-windows-amd64.exe cmd/concise/concisectl.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -ldflags "-X main.cliVersion=${version}" -o concise-cli-darwin-amd64 cmd/concise/concisectl.go
