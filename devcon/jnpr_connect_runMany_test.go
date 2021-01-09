package devcon

import (
	"fmt"
	"testing"
)

func TestRunManyCmds(t *testing.T) {
	cInfo := NewConnInfo()
	c := NewConfig(cInfo)
	b, err := RunManyCmds(c)
	if err != nil {
		t.Errorf("RunCmd failed with %v", err)
	}
	if len(b.String()) == 0 {
		t.Error("Bytes Buffer is empty!")
	}
}

func TestRunManyCmdsAuthProblem(t *testing.T) {
	authErr := AuthErr{}
	cInfo := NewConnInfo()
	cInfo.Username = "fail"
	c := NewConfig(cInfo)
	_, err := RunManyCmds(c)
	if err != authErr {
		t.Errorf("expected authProblem, got %v", err)
	}
}

func TestRunManyCmdsSyntaxErr(t *testing.T) {
	syntaxErr := SyntaxErr{}
	cInfo := NewConnInfo()
	cInfo.Commands = []string{"show version", "show blah", "exit"}
	c := NewConfig(cInfo)
	_, err := RunManyCmds(c)
	if err != syntaxErr {
		t.Errorf("expected syntaxProblem, got %v", err)
	}
}

func TestRunManyCmdsTimeOutErr(t *testing.T) {
	timeOutErr := IOTimeOutErr{}
	cInfo := NewConnInfo()
	cInfo.IP = "1.1.1.1"
	c := NewConfig(cInfo)
	_, err := RunCmd(c)
	if err != timeOutErr {
		t.Errorf("expected TimeOutErr, got %v", err)
	}
}

// Configuration Tests
func TestDifferSet(t *testing.T) {
	cInfo := NewConnInfo()
	cInfo.Conf = "set interfaces ge-0/0/5.105 description TestLTES1__TEST"
	c := NewConfig(cInfo)
	b, err := DifferSet(c)
	if err != nil {
		t.Errorf("RunCmd failed with %v", err)
	}
	if len(b.String()) == 0 {
		t.Error("Bytes Buffer is empty!")
	}
}

func rollback(c *ConnInfo) error {
	c.Conf = "rollback 1"
	c.CommitComment = "ROLLING_BACK_TEST"
	
	b, err := PusherSet(c)
	if err != nil {
		return fmt.Errorf("RunCmd failed with %v", err)
	}
	if len(b.String()) == 0 {
		return fmt.Errorf("Bytes Buffer is empty!")
	}
	return nil
}

func TestPusherSet(t *testing.T) {
	cInfo := NewConnInfo()
	cInfo.Conf = "set interfaces ge-0/0/5.105 description TestLTES1__TEST"
	cInfo.CommitComment = "TESTING_PUSHER"
	c := NewConfig(cInfo)
	b, err := PusherSet(c)
	if err != nil {
		t.Errorf("RunCmd failed with %v", err)
	}
	if len(b.String()) == 0 {
		t.Error("Bytes Buffer is empty!")
	}
	if err := rollback(c); err != nil {
		t.Errorf("Rollback failed: %v", err)
	}
}