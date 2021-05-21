package creator

import "os"

var ignoreFileTemplate = `
.project
.settings
target
.classpath
bin
log/
.idea/
logs/
*.iml
*.uml
LOG_HOME_IS_UNDEFINED/
*/.sts4-cache
*/.springBeans
`

func CreateIgnoreFile(basePath string) error {
	ignore, err := os.Create(basePath + "/.gitignore")
	if err != nil {
		return err
	}
	_, err = ignore.WriteString(ignoreFileTemplate)
	if err != nil {
		return err
	}
	return ignore.Close()
}
