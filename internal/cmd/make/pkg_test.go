package make_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/egg/internal/cmd/make"
)

func Test(t *testing.T) {
	buf := new(bytes.Buffer)
	root := New()
	root.SetOut(buf)

	t.Run("build", func(t *testing.T) {
		root.SetArgs([]string{"build"})
		err := root.Execute()
		assert.NoError(t, err)
	})
}
