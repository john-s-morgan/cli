//go:build integration
// +build integration

package destroy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/cli/internal/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestDestroyMailer(t *testing.T) {
	r := require.New(t)
	r.NoError(testhelpers.EnsureBuffaloCMD(t))

	testhelpers.RunWithinTempFolder(t, func(tt *testing.T) {
		rr := require.New(tt)

		out, err := testhelpers.RunBuffaloCMD(t, []string{"new", "testapp"})
		tt.Log(out)
		rr.NoError(err)

		os.Chdir("testapp")

		out, err = testhelpers.RunBuffaloCMD(t, []string{"generate", "mailer", "ouch"})
		tt.Log(out)
		rr.NoError(err)

		out, err = testhelpers.RunBuffaloCMD(t, []string{"d", "mailer", "ouch", "-y"})
		tt.Log(out)
		rr.NoError(err)

		r.NoFileExists(filepath.Join("mailers", "ouch.go"))
		r.NoFileExists(filepath.Join("templates", "mail", "ouch.plush.html"))
	})

}
