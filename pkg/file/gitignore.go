package file

import "os"

var gitIgnore = `.project
.settings
target
.classpath
bin
log/
/.apt_generated/
/.apt_generated_tests/
.springBeans
.factorypath
.idea
.sts4-cache
*.iml
temp
`

//BuildGitIgnore 构建gitignore文件
func BuildGitIgnore(path string) error {
	target, err := os.OpenFile(path+"/.gitignore", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.Write([]byte(gitIgnore))
	return err
}
