package devcon

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
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

// assignStdInAndOut takes in a pointer to an ssh.Session and returns a Reader, 
// Writer, and an error
func assignStdInAndOut(sess *ssh.Session) (io.Reader, io.WriteCloser, error) {
	// Store the session output to an io.Reader
	sshOut, err := sess.StdoutPipe()
	if err != nil {
		var sIn io.WriteCloser
		return sshOut, sIn, errors.New(fmt.Sprintf("Failed to get stdOut: %v", err))
	}
	// StdinPipe for commands
	stdIn, err := sess.StdinPipe()
	if err != nil {
		return sshOut, stdIn, errors.New(fmt.Sprintf("Failed to get stdIn: %v", err))
	}
	return sshOut, stdIn, nil
}

// parse takes in an io.Reader and iterates through each line
// if "error:" appears in the line, it will log the line.
// if "configuration check succeeds" appears in the line, it will log "commit check successful"
// TODO: look for the following:
//  - need to account for session locks
//  - commit check error
func parse(ip string, r io.Reader) (bytes.Buffer, error) {
	var b bytes.Buffer
	// create a reader through which we'll iterate line ('\n') by line
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return b, err
			}
		}
		// If error is in the string, return an empty buffer and the error
		if strings.Contains(line, "error:") {
			return b, fmt.Errorf("Error: %v", line)
		}
		// let the user know the config check was sussessful
		if strings.Contains(line, "configuration check succeeds") {
			b.WriteString("commit check successful")
			continue
		}
		// return a syntax problem if the text shows up in the line
		if strings.Contains(line, "syntax error, expecting") {
			b.WriteString("syntax problem")
			return b, SyntaxErr{}
		}
		// format the IP info with the buffered line
		l := fmt.Sprintf("[%v] %v", ip, line)
		b.WriteString(l)
	}
	return b, nil
}

// executeMany sets up an interactive session with the target device
func executeMany(c *ConnInfo) (bytes.Buffer, error) {
	var b bytes.Buffer
	h := fmt.Sprintf(c.IP + ":" + port)
	client, err := ssh.Dial("tcp", h, c.ClientCfg)
	e := errorChecker(err)
	if err != nil {
		return b, e
	}
	// close the client
	defer client.Close()
	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		return b, err
	}
	// close the session
	defer sess.Close()

	sshOut, stdIn, err := assignStdInAndOut(sess)
	if err != nil {
		return b, err
	}
	// Start remote shell
	err = sess.Shell()
	if err != nil {
		return b, errors.New(fmt.Sprintf("Failed connect to shell: %v", err))
	}

	// send the commands
	for _, cmd := range c.Commands {
		_, err = fmt.Fprintf(stdIn, "%s\n", cmd)
		if err != nil {
			return b, errors.New(fmt.Sprintf("Failed to get cmd output: %v", err))
		}
	}
	buf, err := parse(c.IP, sshOut)
	if err != nil {
		return buf, err
	}
	err = sess.Wait()
	if err != nil {
		return buf, errors.New(fmt.Sprintf("Failed to exit: %v", err))
	}
	return buf, nil
}