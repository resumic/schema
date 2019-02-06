package theme

import (
	"os"
	"path"

	git "gopkg.in/src-d/go-git.v4"
)

const officialThemesRepo = "https://github.com/ArmanMazdaee/test-resumic-themes.git"

func getOfficialThemesDir(cacheDir string) string {
	return path.Join(cacheDir, "official_themes")
}

func pullOfficialThemes(root string) error {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		_, err := git.PlainClone(root, false, &git.CloneOptions{
			URL: officialThemesRepo,
		})
		return err
	}

	repo, err := git.PlainOpen(root)
	if err != nil {
		return err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	return worktree.Pull(&git.PullOptions{})
}
