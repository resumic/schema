package theme

type ThemeNotFoundError struct {
	name string
}

func (e *ThemeNotFoundError) Error() string {
	return "theme: " + e.name + " could not be found."
}

func isThemeNotFoundError(err error) bool {
	_, ok := err.(*ThemeNotFoundError)
	return ok
}

func GetThemesDir(themesName, cacheDir string) (string, error) {
	themesDir, err := getDefaultThemesDir(themesName, cacheDir)
	if err == nil || !isThemeNotFoundError(err) {
		return themesDir, err
	}

	return getOfficialThemesDir(themesName, cacheDir)
}
