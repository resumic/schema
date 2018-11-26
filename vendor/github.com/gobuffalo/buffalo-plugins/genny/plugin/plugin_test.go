package plugin

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools/gomods"
	"github.com/stretchr/testify/require"
)

func Test_Generator(t *testing.T) {
	r := require.New(t)
	opts := &Options{
		PluginPkg: "github.com/foo/buffalo-bar",
		Year:      1999,
		Author:    "Homer Simpson",
		ShortName: "bar",
	}
	run := genny.DryRunner(context.Background())
	run.Root = os.TempDir()

	gg, err := New(opts)
	r.NoError(err)
	run.WithGroup(gg)

	r.NoError(run.Run())
	res := run.Results()

	var cmds []string
	if !gomods.On() {
		cmds = []string{"git init", "go get github.com/alecthomas/gometalinter", "gometalinter --install"}
	} else {
		cmds = []string{"git init", "go mod init github.com/foo/buffalo-bar", "go get github.com/alecthomas/gometalinter", "gometalinter --install", "go mod tidy"}
	}

	r.Len(res.Commands, len(cmds))
	for i, x := range cmds {
		r.Equal(x, strings.TrimSpace(strings.Join(res.Commands[i].Args, " ")))
	}

	files := []string{
		".gitignore",
		".gometalinter.json",
		".goreleaser.yml.plush",
		".travis.yml",
		"LICENSE",
		"Makefile",
		"README.md",
		"bar/listen.go",
		"bar/version.go",
		"cmd/available.go",
		"cmd/bar.go",
		"cmd/root.go",
		"cmd/version.go",
		"main.go",
		"readme.md",
	}

	for i, f := range res.Files {
		r.Equal(files[i], f.Name())
	}
	r.Len(res.Files, len(files))

	f := res.Files[6]
	r.Equal("README.md", f.Name())
	r.Contains(f.String(), opts.PluginPkg)

	f = res.Files[12]
	r.Equal("cmd/version.go", f.Name())
	r.Contains(f.String(), opts.PluginPkg+"/"+opts.ShortName)
	r.Contains(f.String(), opts.ShortName+".Version")

	f = res.Files[13]
	r.Equal("main.go", f.Name())
	r.Contains(f.String(), "github.com/foo/buffalo-bar/cmd")

}
