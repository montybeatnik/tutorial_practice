package autochecks

type Params struct {
	IP   string
	Args []string
}

// AutoCheck is the API or the method signature to represent what needs to happen
// for an autocheck implementation
type AutoCheck interface {
	// Run issues the command against the network element
	Run(p Params) (interface{}, error)
}
