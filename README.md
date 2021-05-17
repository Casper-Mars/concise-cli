# concise项目的脚手架工具

<p align="center">
  <a href="https://github.com/Casper-Mars/concise-cli/actions/workflows/go.yml"><img src="https://github.com/Casper-Mars/concise-cli/actions/workflows/go.yml/badge.svg" alt="Github Actions status"></a>
</p>

---

- [concise项目的脚手架工具](#concise项目的脚手架工具)
    - [安装](##安装)
        - [linux/mac](###linux/mac)
        - [windows](###windows)
    - [快速入门](##快速入门)
    - [参与贡献](##参与贡献)

---

## 安装

### linux/mac

* 下载最新的release

* 把下载好的二进制包重命名为 `concise-cli` 并移动到 `/usr/local/bin` 下

```shell
mv concise-cli /usr/local/bin/
```

* 查询版本号来验证安装结果

```shell
concise-cli -v
```

如果出现命令找不到，则安装失败，可能是环境没刷新。否则就是成功。

### windows

## 快速入门

* 查看帮助

```shell
~: concise-cli -help
```

* 初始化项目

```shell
~: concise-cli-${os}-{architecture}-${version} -m dao -p 1.0.1
```

> 命令执行成功后，会在执行的目录中看到多出一个dao目录，然后打开idea/eclipse打开(导入)即可。

## 参与贡献

项目比较简单，只要会使用make就能轻松上手。具体的make命令参考Makefile的注释。