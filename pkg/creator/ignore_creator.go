package creator

import "os"

var ignoreFileTemplate = ".project\n.settings\ntarget\n.classpath\nbin\nlog/\n.idea/\nlogs/\n*.iml\n*.uml\nLOG_HOME_IS_UNDEFINED/\n*/.sts4-cache\n*/.springBeans\n"

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
