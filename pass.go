package gopass

import (
	"errors"
	"os"
)

var ErrInterrupted = errors.New("Interrupted")

type outputMode int
const (
	hidden outputMode = iota
	masked outputMode = iota
	visible outputMode = iota
)

// getPasswd returns the input read from terminal.
// If masked is true, typing will be matched by asterisks on the screen.
// Otherwise, typing will echo nothing.
func getPasswd(om outputMode) ([]byte, error) {
	var pass, bs, mask []byte
	if om == masked {
		bs = []byte("\b \b")
		mask = []byte("*")
	}

	var err error
	for {
		if v := getch(); v == 127 || v == 8 {
			if l := len(pass); l > 0 {
				pass = pass[:l-1]
				if om == masked {
					os.Stdout.Write(bs)
				}
			}
		} else if v == 13 || v == 10 || v == 4 {
			break
		} else if v == 3 {
			err = ErrInterrupted
			break
		} else if v != 0 {
			pass = append(pass, v)
			var towrite []byte
			if om == masked {
				towrite = mask
			} else if om == visible {
				towrite = []byte{ v }
			}
			os.Stdout.Write(towrite)
		}
	}
	println()
	if err != nil {
		return nil, err
	}
	return pass, nil
}

// GetPasswd returns the password read from the terminal without echoing input.
// The returned byte array does not include end-of-line characters.
func GetPasswd() ([]byte, error) {
	return getPasswd(hidden)
}

// GetPasswdMasked returns the password read from the terminal, echoing asterisks.
// The returned byte array does not include end-of-line characters.
func GetPasswdMasked() ([]byte, error) {
	return getPasswd(masked)
}

// GetPrompt is as above, but without any masking or hiding
func GetPrompt() ([]byte, error) {
	return getPasswd(visible)
}
