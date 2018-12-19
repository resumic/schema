package render

import (
	"io"
	"os"
	"path"

	"github.com/gobuffalo/packd"
	packr "github.com/gobuffalo/packr/v2"
)

var themesList = map[string]bool{
	"test-theme": true,
}

type InvalidThemeError struct {
	name string
}

func (e *InvalidThemeError) Error() string {
	return "render: " + e.name + " is not a valid theme."
}

func extractTheme(root, theme string) error {
	if !themesList[theme] {
		return &InvalidThemeError{name: theme}
	}

	box := packr.New("themes", "../render/themes")
	return box.WalkPrefix(theme+"/", func(name string, file packd.File) error {
		p := path.Join(root, name)
		os.MkdirAll(path.Dir(p), 0777)
		dst, err := os.Create(p)
		defer dst.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, file)
		return err
	})
}
