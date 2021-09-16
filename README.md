# concise项目的脚手架工具

---

- [concise项目的脚手架工具](#concise项目的脚手架工具)
    - [安装](##安装)
        - [linux/mac](###linux/mac)
        - [windows](###windows)
    - [快速入门](##快速入门)

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
concise-cli version
```

如果出现命令找不到，则安装失败，可能是环境没刷新。否则就是成功。

### windows

* 下载最新的release

* 把下载好的二进制包重命名为 `concise-cli.exe` 并移动到用户自定义的任意一个目录下

* 把该目录加入到环境变量PATH中

* 打开cmd验证安装

```shell
concise-cli version
```

## 快速入门

* 查看帮助

```shell
concise-cli help
```

* 初始化项目

```shell
concise-cli -r http://xxx.xxxx.xx/xxx/xxx.git
```

> 命令执行成功后，会在执行的目录中看到多出一个demo目录，然后打开idea/eclipse打开(导入)即可。
