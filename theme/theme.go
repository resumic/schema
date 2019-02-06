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

func UpdateCache(cacheDir string) error {
	defaultThemesDir := getDefaultThemesDir(cacheDir)
	err := extractDefaultThemes(defaultThemesDir)
	if err != nil {
		return err
	}
	officialThemesDir := getOfficialThemesDir(cacheDir)
	return pullOfficialThemes(officialThemesDir)
}

func GetThemesDir(themesName, cacheDir string) (string, error) {
	defaultThemesDir := getDefaultThemesDir(cacheDir)
	themesDir := path.Join(defaultThemesDir, themesName)
	if _, err := os.Stat(themesDir); !os.IsNotExist(err) {
		return themesDir, err
	}

	officialThemesDir := getOfficialThemesDir(cacheDir)
	themesDir = path.Join(officialThemesDir, themesName)
	if _, err := os.Stat(themesDir); !os.IsNotExist(err) {
		return themesDir, err
	}
	return "", &ThemeNotFoundError{name: themesName}
}
