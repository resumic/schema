package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/resumic/schema/render"
	"github.com/resumic/schema/theme"
	"github.com/spf13/cobra"
)

// Added in case we change default filenames.
const (
	defaultHTML   = "resume.html"
	defaultResume = "resume.json"
)

func renderRun(cmd *cobra.Command, args []string) error {
	var (
		resumePath string
		htmlPath   string
	)
	switch argLen := len(args); argLen {
	// No args: use defaults for input file and output file.
	case 0:
		resumePath = defaultResume
		htmlPath = defaultHTML
	// 1 argument: check to see what the file type is. If JSON, use as input file.
	// If HTML, use as output file. This doesn't permit using files with no extension.
	case 1:
		if strings.Contains(args[0], ".json") {
			resumePath = args[0]
			htmlPath = defaultHTML
		} else if strings.Contains(args[0], ".html") {
			resumePath = defaultResume
			htmlPath = args[0]
		} else {
			return fmt.Errorf("Please provide argument(s) with a valid .json or .html extension")
		}
	// 2 args: use args for input and output files.
	case 2:
		resumePath = args[0]
		htmlPath = args[1]
	}
	// resumePath := args[0]
	// htmlPath := args[1]
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
	Use:   "render JSON_PATH (defaults to resume.json) HTML_PATH (defaults to resume.html)",
	Short: "Render json resume to html",
	Args:  cobra.MaximumNArgs(2),
	RunE:  renderRun,
}

func init() {
	renderCmd.Flags().StringP("theme", "t", "test-theme", "Theme to use")
	renderCmd.Flags().StringP("themesDir", "d", "", "Filesystem path to themes directory")
	rootCmd.AddCommand(renderCmd)
}
