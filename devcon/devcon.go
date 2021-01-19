package devcon

import (
	"bytes"
	"fmt"

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

// RunManyCmds takes in a pointer to ConnInfo and if nothing goes wrong
// returns a bytes.Buffer and a nil error.
// Should an error occur,
// it returns a nil bytes.Buffer and a meaningful error.
func RunManyCmds(c *ConnInfo) (bytes.Buffer, error) {
	b, err := executeMany(c)
	if err != nil {
		return b, err
	}
	return b, nil
}

// DifferSet takes in a pointer to ConnInfo
// and returns a buffer and an error if nothing goes wrong.
// You need only supply the IP and the config because it will be rolled back.
func DifferSet(c *ConnInfo) (bytes.Buffer, error) {

	// JunOS CLI arguments to provide a diff
	c.Commands = []string{
		"configure exclusive",
		c.Conf,
		"commit check",
		"show | compare",
		"rollback",
		"exit",
		"exit",
	}
	b, err := executeMany(c)
	if err != nil {
		return b, err
	}
	return b, nil
}

// PusherSet Pushes config to the device
// It drops into the device's config exclusive mode,
// does a commit check and a 'show | compare' as well as applies a commit message.
// You have to supply the config in set format and a meaningful commit message.
// The func formats everything for you.
func PusherSet(c *ConnInfo) (bytes.Buffer, error) {
	cc := fmt.Sprintf("commit and-quit comment %v", c.CommitComment)
	c.Commands = []string{
		"configure exclusive",
		c.Conf,
		"commit check",
		"show | compare",
		cc,
		"exit",
		"exit",
	}
	b, err := executeMany(c)
	if err != nil {
		return b, err
	}
	return b, nil
}

// PusherNID Pushes config to a NID
func PusherNID(c *ConnInfo) (bytes.Buffer, error) {
	c.Commands = []string{
		c.Conf,
		"exit",
	}
	b, err := executeMany(c)
	if err != nil {
		return b, err
	}
	return b, nil
}
