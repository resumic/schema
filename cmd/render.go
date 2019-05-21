package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/render"
	"github.com/resumic/schema/theme"
	"github.com/spf13/cobra"
)

func renderRun(cmd *cobra.Command, args []string) error {
	resumePath := args[0]
	htmlPath := args[1]
	cacheDir, err := cmd.Flags().GetString("cacheDir")
	if err != nil {
		panic(err)
	}
	themesName, err := cmd.Flags().GetString("theme")
	if err != nil {
		panic(err)
	}
	themesDir, err := cmd.Flags().GetString("themesDir")
	if err != nil {
		panic(err)
	}

	if themesDir == "" {
		themesDir, err = theme.GetThemesDir(themesName, cacheDir)
		if err != nil {
			return err
		}
	}

	resume, err := ioutil.ReadFile(resumePath)
	if err != nil {
		return fmt.Errorf("Couldn't read the json resume from %s: %s", resumePath, err)
	}
	html, err := render.RenderHTML(resume, themesDir)
	if err != nil {
		return fmt.Errorf("Couldn't render the html resume: %s", err)
	}
	err = ioutil.WriteFile(htmlPath, html, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the html resume to %s: %s", htmlPath, err)
	}
	fmt.Printf("%s rendered successfully to %s using %s as theme\n", resumePath, htmlPath, themesDir)
	return nil
}

var renderCmd = &cobra.Command{
	Use:   "render JSON_PATH HTML_PATH",
	Short: "Render json resume to html",
	Args:  cobra.ExactArgs(2),
	RunE:  renderRun,
}

func init() {
	renderCmd.Flags().StringP("theme", "t", "orbit", "Theme to use")
	renderCmd.Flags().StringP("themesDir", "d", "", "Filesystem path to themes directory")
	rootCmd.AddCommand(renderCmd)
}
