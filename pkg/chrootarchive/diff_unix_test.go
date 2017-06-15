package chrootarchive

import (
	"testing"
	"github.com/docker/docker/pkg/system"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"

	"github.com/docker/docker/pkg/archive"
	"fmt"
	"encoding/json"
)

func TestUmask(t *testing.T) {
	oldmask, err := system.Umask(0)
	defer system.Umask(oldmask)
	assert.NoError(t,err)
}

func TestUnPackLayer(t *testing.T) {
	var options *archive.TarOptions

	tmpDir, err := ioutil.TempDir("/", "temp-docker-extract")
	assert.NoError(t,err)


	os.Setenv("TMPDIR", tmpDir)
	size, err := archive.UnpackLayer("/", os.Stdin, options)
	os.RemoveAll(tmpDir)
	assert.NoError(t,err)

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(applyLayerResponse{size}); err != nil {
		fatal(fmt.Errorf("unable to encode layerSize JSON: %s", err))
	}
}