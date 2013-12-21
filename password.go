package readpass

// This file contains utillity functions for decrypting password protecting keys
// and password protecting keys.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The PasswordPrompt function is the function that is called to prompt the user for
// a password.
var PasswordPrompt func(prompt string) (password string, err error) = DefaultPasswordPrompt

// DefaultPasswordPrompt is a simple (but echoing) password entry function
// that takes a prompt and reads the password.
func DefaultPasswordPrompt(prompt string) (password string, err error) {
	fmt.Printf(prompt)
	rd := bufio.NewReader(os.Stdin)
	line, err := rd.ReadString('\n')
	if err != nil {
		return
	}
	password = strings.TrimSpace(line)
	return
}
