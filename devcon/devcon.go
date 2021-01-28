package devcon

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

const (
	// define conn params
	username = "chernautomation"
	password = "T8600cr01"
	port     = "22"
)

// ConnInfo holds the elements to setup an SSH client
type ConnInfo struct {
	Username      string
	Password      string
	IP            string
	Command       string
	Commands      []string
	Conf          string
	CommitComment string
	ClientCfg     *ssh.ClientConfig
}

// RunCmd takes in an pointer to ConnInfo and returns a bytes.Buffer
// and a nil error if successful.
// Should an error occur,
// it returns a nil bytes.Buffer and a meaningful error.
func RunCmd(c *ConnInfo) (bytes.Buffer, error) {
	b, err := executeOne(c)
	if err != nil {
		return b, err
	}
	return b, nil
}
