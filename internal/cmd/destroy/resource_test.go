//go:build integration
// +build integration

package destroy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/cli/internal/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestDestroyResource(t *testing.T) {
	r := require.New(t)
	r.NoError(testhelpers.EnsureBuffaloCMD(t))

	testhelpers.RunWithinTempFolder(t, func(tt *testing.T) {
		rr := require.New(tt)

		out, err := testhelpers.RunBuffaloCMD(t, []string{"new", "testapp"})
		tt.Log(out)
		rr.NoError(err)

		os.Chdir("testapp")

		out, err = testhelpers.RunBuffaloCMD(t, []string{"generate", "resource", "ouch"})
		tt.Log(out)
		rr.NoError(err)

		out, err = testhelpers.RunBuffaloCMD(t, []string{"d", "resource", "ouch", "-y"})
		tt.Log(out)
		rr.NoError(err)

		r.NoFileExists(filepath.Join("models", "ouch.go"))
		r.NoFileExists(filepath.Join("models", "ouch_test.go"))
		r.NoFileExists(filepath.Join("actions", "ouches.go"))
		r.NoFileExists(filepath.Join("actions", "ouches_test.go"))
		r.NoFileExists(filepath.Join("locales", "ouches.en-us.yaml"))

		fcontent, err := os.ReadFile(filepath.Join("actions", "app.go"))
		rr.NoError(err)

		r.NotContains(string(fcontent), `app.Resource(\"/ouches\", OuchesResource{})`)

	})

}
