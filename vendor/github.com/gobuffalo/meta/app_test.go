package meta

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/require"
)

func Test_ModulesPackageName(t *testing.T) {
	r := require.New(t)
	tmp := os.TempDir()
	envy.Set(envy.GO111MODULE, "on")

	r.NoError(os.Chdir(tmp))

	tcases := []struct {
		Content     string
		PackageName string
	}{
		{"module github.com/wawandco/zekito", "github.com/wawandco/zekito"},
		{"module zekito", "zekito"},
		{"module gopkg.in/some/friday.v2", "gopkg.in/some/friday.v2"},
		{"", "zekito"},
	}

	for _, tcase := range tcases {
		envy.Set("GOPATH", tmp)

		t.Run(tcase.Content, func(st *testing.T) {
			r := require.New(st)

			r.NoError(ioutil.WriteFile("go.mod", []byte(tcase.Content), 0644))

			a := New(filepath.Join(tmp, "zekito"))
			r.Equal(tcase.PackageName, a.PackagePkg)
		})
	}
}

func Test_App_Encoding(t *testing.T) {
	r := require.New(t)

	a := New(".")
	bb := &bytes.Buffer{}
	r.NoError(a.Encode(bb))

	b := App{}

	r.NoError((&b).Decode(bb))

	r.Equal(a.String(), b.String())
}

func Test_App_IsZero(t *testing.T) {
	r := require.New(t)

	app := App{}
	r.True(app.IsZero())
	app = New(".")
	r.False(app.IsZero())
}

func Test_App_PackageRoot(t *testing.T) {
	r := require.New(t)

	app := App{}

	app.PackageRoot("foo.com/bar")
	r.Equal("foo.com/bar/actions", app.ActionsPkg)
	r.Equal("foo.com/bar/models", app.ModelsPkg)
	r.Equal("foo.com/bar/grifts", app.GriftsPkg)
}
