package devcon

var (
	ip       = "10.28.12.251"
	commands = []string{"show version", "show system uptime", "exit"}
)

func NewConnInfo() ConnInfo {
	return ConnInfo{
		Username: username,
		Password: password,
		IP:       ip,
		Command:  "show version",
		Commands: commands,
		Conf:     "",
	}
}
