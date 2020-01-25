package tools

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/izumin5210/gex"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCommand_GoModWithDep(t *testing.T) {
	var (
		issue    = "testdata/issue-52/tools"
		expected = "tools.golden"
		obtained = "tools.go"
	)
	golden, err := os.Stat(filepath.Join(issue, expected))
	require.NoError(t, err)

	buf, cfg := new(bytes.Buffer), &(*gex.Default)
	cfg.FS = afero.NewMemMapFs()
	cfg.ErrWriter = buf
	cfg.OutWriter = buf
	cfg.RootDir = issue
	cfg.WorkingDir = issue
	cmd := NewAddCommand(cfg)
	cmd.SetOut(buf)

	assert.NoError(t, cmd.RunE(cmd, []string{"golang.org/x/vgo"}))

	manifest, err := cfg.FS.Stat(filepath.Join(issue, obtained))
	assert.NoError(t, err)
	assert.Equal(t, golden.Size(), manifest.Size())
}
