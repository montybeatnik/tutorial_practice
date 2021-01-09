package devcon

import "testing"

var (
	accIP       = "10.46.52.122"
)

func TestRunAccCmd(t *testing.T) {
	cInfo := NewConnInfo()
	c := NewConfig(cInfo)
	c.IP = accIP
	c.Commands = []string{"filter show l2", "exit"}
	b, err := RunManyCmds(c)
	if err != nil {
		t.Errorf("RunCmd failed with %v", err)
	}
	if len(b.String()) == 0 {
		t.Error("Bytes Buffer is empty!")
	}
}

// func TestRunCmdAuthProblem(t *testing.T) {
// 	authErr := AuthErr{}
// 	cInfo := newConnInfo()
// 	cInfo.Username = "fail"
// 	c := NewConfig(cInfo)
// 	_, err := RunCmd(c)
// 	if err != authErr {
// 		t.Errorf("expected authProblem, got %v", err)
// 	}
// }

// func TestRunCmdSyntaxErr(t *testing.T) {
// 	syntaxErr := SyntaxErr{}
// 	cInfo := newConnInfo()
// 	cInfo.Command = "blah"
// 	c := NewConfig(cInfo)
// 	_, err := RunCmd(c)
// 	if err != syntaxErr {
// 		t.Errorf("expected syntaxProblem, got %v", err)
// 	}
// }

// func TestRunCmdTimeOutErr(t *testing.T) {
// 	timeOutErr := IOTimeOutErr{}
// 	cInfo := newConnInfo()
// 	cInfo.IP = "1.1.1.1"
// 	c := NewConfig(cInfo)
// 	_, err := RunCmd(c)
// 	if err != timeOutErr {
// 		t.Errorf("expected TimeOutErr, got %v", err)
// 	}
// }
