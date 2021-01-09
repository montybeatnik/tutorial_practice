# Devcon | Netacles
The devcon package offers a simple API to talk to network elements.
We are now at version 2.0 where things have been scaled down to the most fundamental elements. At it's essence, devcon, takes in some connection info (ip, command(s), creds), logs into the target device, does what you tell it, and provides a buffer with the output.

# How To Use The Tool
Import the devcon package from the web_server_for_network_devices library.
See an example below.

```golang
import (
    "github.com/montybeatnik/web_server_for_network_devices/devcon"
)
func main() {
    // build the ConnInfo with the constructor
    cfgInfo := NewConfig(connInfo)
    // use the ConnInfo to run the command
    output, err := RunCmd(cfgInfo)
    // check for errors
	if err != nil {
		log.Println(err)
    }
    // print the output
	fmt.Println(output.String())
}
```
# Contributing

```golang
```