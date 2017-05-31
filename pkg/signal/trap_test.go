package signal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDumpStacks(t *testing.T) {
	directory, errorDir := ioutil.TempDir("", "test")
	assert.NoError(t, errorDir)

	defer os.RemoveAll(directory) // clean up

	_, error := DumpStacks(directory)
	path := filepath.Join(directory, fmt.Sprintf(stacksLogNameTemplate, strings.Replace(time.Now().Format(time.RFC3339), ":", "", -1)))

	readFile, _ := ioutil.ReadFile(path)
	fileData := string(readFile)

	assert.NotEqual(t, fileData, "")

	assert.NoError(t, error)

	path, errorPath := DumpStacks("")
	assert.NoError(t, errorPath)
	file := os.Stderr
	assert.EqualValues(t, file.Name(), path)
}
