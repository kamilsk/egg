package deps_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/egg/internal/cmd/deps"
)

func Test(t *testing.T) {
	buf := new(bytes.Buffer)
	root := New()
	root.SetOut(buf)

	t.Run("check", func(t *testing.T) {
		root.SetArgs([]string{"check"})
		err := root.Execute()
		assert.NoError(t, err)
	})
}
