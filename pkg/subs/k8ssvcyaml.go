package subs

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

type Svc struct {
	name string
}

func (s Svc) sub(projectRootPath string) error {
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString("k8s/svc.yaml"))
	if s.name == "" {
		return nil
	}
	filename := fmt.Sprintf("%s/k8s/svc.yaml", projectRootPath)
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	replace := strings.Replace(string(file), "#{__project_name}", s.name, -1)
	return ioutil.WriteFile(filename, []byte(replace), stat.Mode())
}

//NewSvcWorker new k8s/svc.yaml worker
func WithSvcWorker(name string) Worker {
	return &Svc{
		name: name,
	}
}
