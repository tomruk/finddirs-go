//go:build unix

package finddirs

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestUnixAppDirs(t *testing.T) {
	config := &AppConfig{
		Subdir:       "foo/bar",
		UseRoaming:   false,
		ConfigSubdir: "config",
		StateSubdir:  "state",
		CacheSubdir:  "cache",
	}
	d, err := RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, d.ConfigDir, filepath.Join(home, ".config/foo/bar"))
	require.Equal(t, d.StateDir, filepath.Join(home, ".local/state/foo/bar"))
	require.Equal(t, d.CacheDir, filepath.Join(home, ".cache/foo/bar"))
}

func TestUnixUserDirs(t *testing.T) {
	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, d.Desktop, filepath.Join(home, "Desktop"))
	require.Equal(t, d.Downloads, filepath.Join(home, "Downloads"))
	require.Equal(t, d.Documents, filepath.Join(home, "Documents"))
	require.Equal(t, d.Pictures, filepath.Join(home, "Pictures"))
	require.Equal(t, d.Videos, filepath.Join(home, "Videos"))
	require.Equal(t, d.Music, filepath.Join(home, "Music"))
	require.Equal(t, d.Templates, filepath.Join(home, "Templates"))
	require.Equal(t, d.PublicShare, filepath.Join(home, "Public"))

	require.Equal(t,
		[]string{
			filepath.Join(home, ".local/share/fonts"),
			filepath.Join(home, ".fonts"),
			"/usr/share/fonts",
			"/usr/local/share/fonts",
		},
		d.Fonts,
	)
}