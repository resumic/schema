package theme

import (
	"os"
	"path"
)

type ThemeNotFoundError struct {
	name string
}

func (e *ThemeNotFoundError) Error() string {
	return "theme: " + e.name + " could not be found."
}

func UpdateCache(cachePath string) error {
	defaultPath := getDefaultThemesPath(cachePath)
	return extractDefaultThemes(defaultPath)
}

func GetThemesPath(themesName, cachePath string) (string, error) {
	defaultPath := getDefaultThemesPath(cachePath)
	themesPath := path.Join(defaultPath, themesName)
	if _, err := os.Stat(themesPath); os.IsNotExist(err) {
		return "", &ThemeNotFoundError{name: themesName}
	}
	return themesPath, nil
}
