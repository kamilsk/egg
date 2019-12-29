package vanity_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/egg/internal/cmd/vanity"
)

func Test(t *testing.T) {
	buf := new(bytes.Buffer)
	root := New()
	root.SetOut(buf)

	t.Run("generate", func(t *testing.T) {
		root.SetArgs([]string{"generate"})
		err := root.Execute()
		assert.NoError(t, err)
	})
}
