package signal

import(
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSignal(t *testing.T) {
	signal, error := ParseSignal("TERM")
	assert.NoError(t, error)
	assert.EqualValues(t,syscall.SIGTERM.String(),signal.String())
}

func TestValidSignalForPlatform (t *testing.T) {
	var sig uint64
	var isValidSignal = false
	sig = 0
	isValidSignal = ValidSignalForPlatform(syscall.Signal(sig))
	assert.EqualValues(t,false,isValidSignal)

	sig = 45
	isValidSignal = ValidSignalForPlatform(syscall.Signal(sig))
	assert.EqualValues(t,true,isValidSignal)

}
