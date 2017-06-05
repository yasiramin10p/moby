package signal

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSignal(t *testing.T) {
	_, error := ParseSignal("SIG")
	assert.EqualValues(t, error.Error(), "Invalid signal: SIG")

	elements := map[string]string{
		"KILL": syscall.SIGKILL.String(),
		"ABRT": syscall.SIGABRT.String(),
		"TERM": syscall.SIGTERM.String(),
		"BUS":  syscall.SIGBUS.String(),
		"CLD":  syscall.SIGCLD.String(),
		"CONT": syscall.SIGCONT.String(),
		"FPE":  syscall.SIGFPE.String(),
		"HUP":  syscall.SIGHUP.String(),
		"ILL":  syscall.SIGILL.String(),
		"IO":   syscall.SIGIO.String(),
	}

	for arrayList := range elements {
		responseSignal, error := ParseSignal(arrayList)
		value, _ := elements[arrayList]
		assert.NoError(t, error)
		assert.EqualValues(t, value, responseSignal.String())
	}
}

func TestCatchAll(t *testing.T) {
	sigs := make(chan os.Signal, 2)
	CatchAll(sigs)
}

func TestStopCatch(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	StopCatch(sigs)
}
func TestValidSignalForPlatform(t *testing.T) {
	var sig uint64
	var isValidSignal = false
	sig = 0
	isValidSignal = ValidSignalForPlatform(syscall.Signal(sig))
	assert.EqualValues(t, false, isValidSignal)

	sig = 45
	isValidSignal = ValidSignalForPlatform(syscall.Signal(sig))
	assert.EqualValues(t, true, isValidSignal)

}
