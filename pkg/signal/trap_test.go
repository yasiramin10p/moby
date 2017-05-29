package signal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDumpStacks(t *testing.T) {
	path, error := DumpStacks("")
	var file *os.File
	file = os.Stderr
	assert.NoError(t, error)
	assert.EqualValues(t,file.Name(),path)
}
func TestTrap(t *testing.T) {
	Trap()
}