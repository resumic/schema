package theme

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/gobuffalo/packd"
	packr "github.com/gobuffalo/packr/v2"
)

func extractDefaultTheme(themesName, themesDir string, box *packr.Box) error {
	err := os.RemoveAll(themesDir)
	if err != nil {
		return err
	}

	return box.WalkPrefix(themesName, func(name string, file packd.File) error {
		name = strings.TrimPrefix(name, themesName)
		p := path.Join(themesDir, name)
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

func getDefaultThemesDir(themesName, cacheDir string) (string, error) {
	box := packr.New("themes", "../theme/defaults")
	if !box.Has(themesName) {
		return "", &ThemeNotFoundError{name: themesName}
	}

	themesDir := path.Join(cacheDir, "default_themes", themesName)
	err := extractDefaultTheme(themesName, themesDir, box)
	if err != nil {
		return "", err
	}
	return themesDir, nil
}
