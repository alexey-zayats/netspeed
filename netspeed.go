package netspeed

import (
	"fmt"
)

type NetSpeed interface {
	measure() (Result, error)
}

var (
	// ErrUnknownProvider error about no known provider is given
	ErrUnknownProvider = fmt.Errorf("unknown provider given")
)

// Measure is the main entrypoint for the library
// Input: provider.Kind for the provider which must be used to make a measurements
// Output: Result interface or error
func Measure(kind Kind) (Result, error) {

	var ns NetSpeed

	switch kind {
	case KindOokla:
		ns = &ooklaProvider{}
	case KindNetflix:
		ns = &netflixProvider{}
	default:
		return nil, ErrUnknownProvider
	}

	return ns.measure()
}
