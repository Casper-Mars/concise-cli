package subs

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

type Deployment struct {
	name string
}

func (d Deployment) sub(projectRootPath string) error {
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString("k8s/dp.yaml"))
	if d.name == "" {
		return nil
	}
	filename := fmt.Sprintf("%s/k8s/dp.yaml", projectRootPath)
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	replace := strings.Replace(string(file), "#{__project_name}", d.name, -1)
	return ioutil.WriteFile(filename, []byte(replace), stat.Mode())
}

//NewDpWorker new k8s/dp.yaml worker
func WithDpWorker(name string) Worker {
	return &Deployment{
		name: name,
	}
}
