package chrootarchive

import (
	"testing"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"io/ioutil"
	"github.com/docker/docker/pkg/system"
	"github.com/stretchr/testify/assert"
)

func TestChroot(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "docker-TestChroot1")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	destination := filepath.Join(tempDir, "dest")

	err = system.MkdirAll(destination, 0700)
	require.NoError(t, err)

	err = chroot("")
        assert.Equal(t,"Error after fallback to chroot: no such file or directory",err.Error())

	err = chroot(destination)

	assert.NoError(t,err)
}
