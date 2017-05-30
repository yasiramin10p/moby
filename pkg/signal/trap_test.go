package signal

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDumpStacks(t *testing.T) {
	_, error := DumpStacks("InvalidValues")
	assert.EqualValues(t, error.Error(), "failed to open file to write the goroutine stacks: open InvalidValues/"+fmt.Sprintf(stacksLogNameTemplate, strings.Replace(time.Now().Format(time.RFC3339), ":", "", -1))+": no such file or directory")

	path, error := DumpStacks("")
	assert.NoError(t, error)
	var file *os.File
	file = os.Stderr
	assert.EqualValues(t, file.Name(), path)
}
