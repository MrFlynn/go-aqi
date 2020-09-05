package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrflynn/go-aqi"
)

// This is an example cli utility that converts PM2.5 measurements to the corresponding
// AQI value.

var (
	pm25 float64
	help bool
)

func init() {
	flag.Float64Var(&pm25, "v", 0.0, "PM2.5 value in micrograms per meter cubed")
	flag.BoolVar(&help, "h", false, "show this screen")
	flag.Parse()
}

func main() {
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if pm25 == 0.0 {
		fmt.Println("missing argument -v")
		flag.PrintDefaults()
		os.Exit(1)
	}

	result, err := aqi.Calculate(aqi.PM25{pm25})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("AQI is %.3f\n", result.AQI)
}