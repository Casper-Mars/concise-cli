version=0.4.0


# 进行测试
.PHONY: test-kit
test:
	make clean
	mkdir out
	go build -o out/cli cmd/concise/main.go
	cd out \
	&& \
	./cli kit -m test-cli -p 1.0.1

# 清理测试产生的文件
.PHONY: clean
clean:
	rm -rf concise-*

# 构建可执行程序
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-darwin-amd64 main.go
