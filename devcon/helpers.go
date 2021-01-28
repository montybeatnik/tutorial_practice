package devcon

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// NewConfig builds the ClientConfig with the ConnInfo
func NewConfig(connParams ConnInfo) *ConnInfo {
	if connParams.Username == "" {
		connParams.Username = username
	}
	if connParams.Password == "" {
		connParams.Password = password
	}
	connParams.ClientCfg = &ssh.ClientConfig{
		User: connParams.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(connParams.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(time.Second * 5),
	}
	return &connParams
}

// executeOne tries to log into the device and collect the output
func executeOne(c *ConnInfo) (bytes.Buffer, error) {
	b := bytes.Buffer{}
	h := fmt.Sprintf(c.IP + ":" + port)
	client, err := ssh.Dial("tcp", h, c.ClientCfg)
	e := errorChecker(err)
	if err != nil {
		return b, e
	}
	// don't forget to close the session
	defer client.Close()
	// Create session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Wirte stdout to bytes.Buffer
	session.Stdout = &b

	// send the command
	err = session.Run(c.Command)
	e = errorChecker(err)
	if err != nil {
		return b, e
	}
	// a syntax error on the device still fills the bytes.Buffer
	// So we need to check the bytes back from the device for this specific
	// scenario.
	if strings.Contains(b.String(), "syntax error, expecting <command>") ||
		strings.Contains(b.String(), "unknown command") {
		return b, SyntaxErr{}
	}
	return b, nil
}
