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

func TestDestroyAction(t *testing.T) {
	r := require.New(t)
	r.NoError(testhelpers.EnsureBuffaloCMD(t))

	testhelpers.RunWithinTempFolder(t, func(tt *testing.T) {
		rr := require.New(tt)

		out, err := testhelpers.RunBuffaloCMD(t, []string{"new", "testapp"})
		tt.Log(out)
		rr.NoError(err)

		os.Chdir("testapp")

		out, err = testhelpers.RunBuffaloCMD(t, []string{"generate", "action", "ouch", "show"})
		tt.Log(out)
		rr.NoError(err)

		out, err = testhelpers.RunBuffaloCMD(t, []string{"d", "action", "ouch", "-y"})
		tt.Log(out)
		rr.NoError(err)

		r.NoFileExists(filepath.Join("actions", "ouch.go"))
		r.NoFileExists(filepath.Join("actions", "ouch_test.go"))
	})
}
