version=1.0.5

# 清理测试产生的文件
.PHONY: clean
clean:
	rm -rf concise-*

.PHONY: lint
lint:
	golangci-lint run --timeout=5m --config=".golangci.yml"



# 构建可执行程序
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -ldflags "-X 'github.com/Casper-Mars/concise-cli/cmd.CliVersion=${version}'" -o concise-cli-darwin-amd64 main.go
