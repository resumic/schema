package theme

import (
	"io"
	"os"
	"path"

	"github.com/gobuffalo/packd"
	packr "github.com/gobuffalo/packr/v2"
)

func getDefaultThemesPath(cachePath string) string {
	return path.Join(cachePath, "default_themes")
}

func extractDefaultThemes(root string) error {
	box := packr.New("themes", "../theme/defaults")
	err := os.RemoveAll(root)
	if err != nil {
		return err
	}
	return box.Walk(func(name string, file packd.File) error {
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
