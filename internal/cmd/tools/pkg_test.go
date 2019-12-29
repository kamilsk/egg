package tools_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/egg/internal/cmd/tools"
)

func Test(t *testing.T) {
	buf := new(bytes.Buffer)
	root := New()
	root.SetOut(buf)

	t.Run("help", func(t *testing.T) {
		root.SetArgs([]string{"--help"})
		err := root.Execute()
		assert.NoError(t, err)
	})
}
