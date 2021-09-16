package subs

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

type Makefile struct {
	name string
}

func (m Makefile) sub(projectRootPath string) error {
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString("Makefile"))
	if m.name == "" {
		return nil
	}
	filename := fmt.Sprintf("%s/Makefile", projectRootPath)
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	replace := strings.Replace(string(file), "#{__project_name}", m.name, -1)
	return ioutil.WriteFile(filename, []byte(replace), stat.Mode())
}

//WithMakefileWorker new Makefile worker
func WithMakefileWorker(name string) Worker {
	return &Makefile{
		name: name,
	}
}
