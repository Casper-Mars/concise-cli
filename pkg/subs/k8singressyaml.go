package subs

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

type Ingress struct {
	name   string
	domain string
}

func (i Ingress) sub(projectRootPath string) error {
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString("k8s/ingress.yaml"))
	if i.name == "" && i.domain == "" {
		return nil
	}
	filename := fmt.Sprintf("%s/k8s/ingress.yaml", projectRootPath)
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	replace := string(file)
	if i.name != "" {
		replace = strings.Replace(replace, "#{__project_name}", i.name, -1)
	}
	if i.domain != "" {
		replace = strings.Replace(replace, "#{__project_domain}", i.domain, -1)
	}
	return ioutil.WriteFile(filename, []byte(replace), stat.Mode())
}

//NewIngressWorker new k8s/ingress.yaml worker
func WithIngressWorker(name, domain string) Worker {
	return &Ingress{
		name:   name,
		domain: domain,
	}
}
