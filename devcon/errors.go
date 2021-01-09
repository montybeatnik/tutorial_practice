package devcon

import "strings"

// AuthErr handles credential issues
type AuthErr struct{}

func (h AuthErr) Error() string {
	return "Auth Failure"
}

// TimeOutErr handles network connectivity problems
type TimeOutErr struct{}

func (h TimeOutErr) Error() string {
	return "Timeout Failure"
}

// IOTimeOutErr handles network connectivity problems
type IOTimeOutErr struct{}

func (h IOTimeOutErr) Error() string {
	return "Timeout Failure"
}

// NoRoute handles routing problems
type NoRoute struct{}

func (h NoRoute) Error() string {
	return "No Route to Host"
}

// SyntaxErr handles syntax problems
type SyntaxErr struct{}

func (h SyntaxErr) Error() string {
	return "Syntax Problem"
}

// errorChecker checks for various network socket problems
func errorChecker(err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "timed out") {
			return TimeOutErr{}
		}
		if strings.Contains(err.Error(), "i/o timeout") {
			return IOTimeOutErr{}
		}
		if strings.Contains(err.Error(), "no route to host") {
			return NoRoute{}
		}
		if strings.Contains(err.Error(), "unable to authenticate") {
			return AuthErr{}
		}
		if strings.Contains(err.Error(), "syntax error, expecting <command>") {
			return SyntaxErr{}
		}
		return err
	}
	return nil
}