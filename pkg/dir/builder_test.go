package dir

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuild(t *testing.T) {

	testYaml := `
name: _root
child:
  - name: _hack
  - name: _src
    child:
      - name: _main
`
	err := Build([]byte(testYaml), ".")
	assert.NoError(t, err)
	assert.True(t, isExist("./_root"))
	assert.True(t, isExist("./_root/_hack"))
	assert.True(t, isExist("./_root/_src/_main"))
	os.RemoveAll("./_root")
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
