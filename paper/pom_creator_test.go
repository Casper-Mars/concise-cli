package paper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePom(t *testing.T) {

	moduleName := "test-module"
	parentVersion := "1.0.0"

	err := CreatePom(".", moduleName, parentVersion)
	require.NoError(t, err)

}
