package main

import (
	"flag"
	"fmt"
	"github.com/alexey-zayats/netspeed"
	"os"
)

func main() {

	var pkind string

	flag.StringVar(&pkind, "provider", "", "Choose a provider to measure net speed with: ookla, netflix")
	flag.Parse()

	result, err := netspeed.Measure(netspeed.KindFromString(pkind))
	if err != nil {
		fmt.Printf("unable measure network speed with provider %s: %s", pkind, err)
		os.Exit(1)
	}

	fmt.Printf("Provider: %s\n", result.Provider())
	fmt.Printf("Download: %f Mbps\n", result.Download())
	fmt.Printf("Upload: %f Mbps\n", result.Upload())
	fmt.Printf("Time: %s\n", result.Latency().String())
}
