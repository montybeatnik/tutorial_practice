
# Autochecks
## Overview
Autochecks wraps around the devcon package providing a layer of abstraction around the fundamental aspect of collecting info from a target device.
The API implements a very simple interface that proves incredibly robust. 

```golang
type AutoCheck interface {
	// Run issues the command against the network element
	Run(p Params) (interface{}, error)
}
```

## Usage 
```golang
func main() {
	// build the argument(s)
	a := []string{"ge-0/1/0"}
	// set the parameters
	p := autochecks.Params{
		IP: "10.63.244.76",
		Args: a,
	}
	var autochecks.
}
```

## Contributing
Here are the steps to add autochecks to the package:
1. Run a command against a device and save the output as valid XML. i.e. ('ssh admin@a.b.c.d "show system alarms | display xml" >> xml/sysAlarms.xml')
2. Navigate to the following URL [XML to GO](https://www.onlinetool.io/xmltogo) and paste in the XML to get a go struct.
3. Create a file with a meaningful name. 
4. Add the struct and two methods to it
	1. Mapper()
	2. Run()

## TODO:
Add support for the following formats:
* JSON
* multi-line strings (unstructured)
