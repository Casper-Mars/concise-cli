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

## 进阶使用

### 参数说明

* r: 模板地址。例如github clone的地址
* b: 分支。
* t: 生成的项目的目录名称。可以是 `demo` 或者是绝对路径 `~/project/demo`
* n: 项目名称，用于 `__project_name` 占位符的替换。
* v: 项目的版本，用于 `__project_version` 占位符的替换。
* d: 域名，用于 `__project_domain` 占位符的替换。
* p: 项目的父工程的版本，用于 `__project_parent_version` 占位符的替换。

### 参数占位符使用

例如后端业务项目，可以使用占位符替换实现初始化maven文件(pom.xml)、Makefile、流水线文件等等。  
在模板项目中添加 `concise.yaml` 配置文件，配置有使用了占位符的文件。例如如下：

```yaml
__project_name:
  - pom.xml
  - Makefile
  - k8s/dp.yaml
  - k8s/ingress.yaml
  - k8s/svc.yaml
__project_version:
  - pom.xml
__project_domain:
  - k8s/ingress.yaml
__project_parent_version:
  - pom.xml
```

在使用脚手架创建项目的时候指定参数，则会对配置文件配置的文件进行占位符替换。







