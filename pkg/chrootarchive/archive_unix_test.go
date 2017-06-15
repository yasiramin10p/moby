package chrootarchive

import (
	"testing"
	"encoding/json"
	"os"

	"github.com/docker/docker/pkg/archive"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"io/ioutil"
	"github.com/docker/docker/pkg/system"
)

func TestNewDecoder(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "docker-TestNewEncoder")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	destination := filepath.Join(tempDir, "dest")

	err = system.MkdirAll(destination, 0700)
	require.NoError(t, err)

	//read the options from the pipe "ExtraFiles"
	var options *archive.TarOptions
	err = json.NewDecoder(os.NewFile(3, destination)).Decode(&options)
	//assert.NoError(t,err)

	if _, err := flush(os.Stdin); err != nil {
		fatal(err)
	}

}
