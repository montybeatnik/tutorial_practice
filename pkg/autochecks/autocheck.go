package autochecks

import "bytes"

type Params struct {
	IP   string
	Args []string
}

// Check is the API or the method signature to represent what needs to happen
// for an autocheck implementation
type Check interface {
	// Run issues the command against the network element
	Run(p Params) (bytes.Buffer, error)
}
