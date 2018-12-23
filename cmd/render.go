package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/render"
	"github.com/spf13/cobra"
)

func renderRun(cmd *cobra.Command, args []string) error {
	resumePath := args[0]
	htmlPath := args[1]
	themeName, err := cmd.Flags().GetString("theme")
	if err != nil {
		panic(err)
	}
	resume, err := ioutil.ReadFile(resumePath)
	if err != nil {
		return fmt.Errorf("Couldn't read the json resume from %s: %s", resumePath, err)
	}
	html, err := render.RenderHTML(resume, themeName)
	if err != nil {
		return fmt.Errorf("Couldn't render the html resume: %s", err)
	}
	err = ioutil.WriteFile(htmlPath, html, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the html resume to %s: %s", htmlPath, err)
	}
	fmt.Printf("%s rendered successfully to %s using %s as theme\n", resumePath, htmlPath, themeName)
	return nil
}

var renderCmd = &cobra.Command{
	Use:   "render JSON_PATH HTML_PATH",
	Short: "Render json resume to html",
	Args:  cobra.ExactArgs(2),
	RunE:  renderRun,
}

func init() {
	renderCmd.Flags().StringP("theme", "t", "test-theme", "Theme to use")
	rootCmd.AddCommand(renderCmd)
}
