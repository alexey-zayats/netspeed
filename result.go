package netspeed

import (
	"time"
)

// Result interface provide measurement results
type Result interface {
	Provider() Kind
	Download() float64
	Upload() float64
	Latency() time.Duration
}

// result struct hods real measurement results
type result struct {
	provider Kind
	download float64
	upload   float64
	latency  time.Duration
}

// Latency returns result latency
func (r result) Latency() time.Duration {
	return r.latency
}

// Provider returns provider kind
func (r result) Provider() Kind {
	return r.provider
}

// Download returns download speed
func (r result) Download() float64 {
	return r.download
}

// Upload returns upload speed
func (r result) Upload() float64 {
	return r.upload
}
