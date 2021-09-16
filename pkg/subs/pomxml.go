package subs

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

type Pom struct {
	name          string
	version       string
	parentVersion string
}

func (p Pom) sub(projectRootPath string) error {
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString("pom.xml"))
	if p.name == "" && p.version == "" && p.parentVersion == "" {
		return nil
	}
	filename := fmt.Sprintf("%s/pom.xml", projectRootPath)
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	replace := string(file)
	if p.name != "" {
		replace = strings.Replace(replace, "#{__project_name}", p.name, -1)
	}
	if p.version != "" {
		replace = strings.Replace(replace, "#{__project_version}", p.version, -1)
	}
	if p.parentVersion != "" {
		replace = strings.Replace(replace, "#{__project_parent_version}", p.parentVersion, -1)
	}
	return ioutil.WriteFile(filename, []byte(replace), stat.Mode())
}

//WithPomWorker new pom.xml worker
func WithPomWorker(name, version, parentVersion string) Worker {
	return &Pom{
		name:          name,
		version:       version,
		parentVersion: parentVersion,
	}
}
