package theme

import (
	"os"
	"path"

	git "gopkg.in/src-d/go-git.v4"
)

const officialThemesRepo = "https://github.com/ArmanMazdaee/test-resumic-themes.git"

func pullOfficialRepo(repoDir string) error {
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		_, err := git.PlainClone(repoDir, false, &git.CloneOptions{
			URL: officialThemesRepo,
		})
		return err
	}

	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		return err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	return worktree.Pull(&git.PullOptions{})
}

func getOfficialThemesDir(themesName, cacheDir string) (string, error) {
	repoDir := path.Join(cacheDir, "official_themes")
	if err := pullOfficialRepo(repoDir); err != nil {
		if err != git.NoErrAlreadyUpToDate {
			return "", err
		}
	}

	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		return "", err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	submodule, err := worktree.Submodule(themesName)
	if err != nil {
		if err == git.ErrSubmoduleNotFound {
			return "", &ThemeNotFoundError{name: themesName}
		}
		return "", err
	}

	if err = submodule.Init(); err != nil {
		if err != git.ErrSubmoduleAlreadyInitialized {
			return "", err
		}
	}

	if err = submodule.Update(&git.SubmoduleUpdateOptions{}); err != nil {
		if err != git.NoErrAlreadyUpToDate {
			return "", err
		}
	}

	themesDir := path.Join(repoDir, submodule.Config().Path)
	return themesDir, nil
}
