package subs

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

var ErrFileNotExist = errors.New("file not exist")

type Worker struct {
	file        string
	placeholder map[string]string
}

func (w Worker) Substitute(ctx context.Context) error {
	if len(w.placeholder) == 0 {
		return nil
	}
	fmt.Printf("[%s] %s\n", color.GreenString("Substitute"), color.WhiteString(w.file))
	stat, err := os.Stat(w.file)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, w.file)
	}
	data, err := ioutil.ReadFile(w.file)
	if err != nil {
		return err
	}
	replace := string(data)
	for holder, value := range w.placeholder {
		replace = strings.Replace(replace, fmt.Sprintf("#{%s}", holder), value, -1)
	}
	return ioutil.WriteFile(w.file, []byte(replace), stat.Mode())
}
