// +build !windows

package readpass

import (
	"fmt"
	"os"
)
import "github.com/gokyle/readpass/readpass"

func unixReadPassword(prompt string) (password string, err error) {
	binPass, err := readpass.ReadPass(prompt)
	fmt.Fprintf(os.Stderr, "\n")
	if err == nil {
		password = string(binPass)
	}
	return
}

func unixReadPasswordBytes(prompt string) (password []byte, err error) {
	password, err = readpass.ReadPass(prompt)
	fmt.Fprintf(os.Stderr, "\n")
	if err != nil {
		password = nil
	}
	return
}

func init() {
	PasswordPrompt = unixReadPassword
	PasswordPromptBytes = unixReadPasswordBytes
}
