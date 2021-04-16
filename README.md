# concise-cli

concise项目的脚手架工具

## 快速入门

* 查看帮助

```shell
~: concise-cli-${os}-{architecture}-${version} -help

Usage of ./concise-cli-darwin-amd64-0.1.0:
  -m string
    	指定模块名称，用于pom文件的artifactId
  -p string
    	指定父工程的版
```

* 初始化项目

```shell
~: concise-cli-${os}-{architecture}-${version} -m dao -p 1.0.1
```

> 命令执行成功后，会在执行的目录中看到多出一个dao目录，然后打开idea/eclipse打开(导入)即可。

# 参与贡献

项目比较简单，只要会使用make就能轻松上手。具体的make命令参考Makefile的注释。