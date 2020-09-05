# go-aqi ![tests](https://github.com/MrFlynn/go-aqi/workflows/Tests/badge.svg)
Calculate the air quality index (AQI) from different particulate concentrations.
The method for calculating AQI is based on the Environmental Protection Agency's
(United States) [table method](https://en.wikipedia.org/wiki/Air_quality_index#United_States) 
for calculating AQI. Other countries' AQI calculation methods might come in the
future, but this might require some rethinking of how this library is currently
written.

Currently this library supports the following measurements:

* [PM2.5](https://en.wikipedia.org/wiki/Particulates)
* [PM10](https://en.wikipedia.org/wiki/Particulates)
* Carbon Monoxide
* Sulfur Dioxide
* Nitrogen Dioxide

Ozone is another measurement that I would like to add, but calculating AQI from
ozone is slightly more complicated compared to the other particulate measures.

## Installation

```
go get -u github.com/mrflynn/go-aqi
```

## Example
Here is an example for calculating AQI from multiple particulate measurements.

```go
package main

import (
	"fmt"

	"github.com/mrflynn/go-aqi"
)

func main() {
	results, err := aqi.Calculate(aqi.PM25{20.2}, aqi.CO{4.1}, aqi.NO2{67.6})
	if err != nil {
		fmt.Prinln(err)
		return
	}

	fmt.Printf(
		"The air quality is %s with an AQI of %.3f\n",
		results.Index.Name,
		results.AQI,
	)
}
```

A toy CLI application can be found in the [`examples/`](examples/cli.go) folder.

## Contributing
Contributions are welcome. Create a pull request and I'll take a look at it
whenever I have some time.

Make sure to include unit tests in your pull requests.

## License

[MIT](LICENSE)
