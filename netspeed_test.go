package netspeed

import (
	"errors"
	"testing"
)

func TestMeasure_Unknown(t *testing.T) {

	_, err := Measure(KindUnknown)

	if errors.Is(err, ErrUnknownProvider) {
		t.Log(err.Error())
	}
}

func TestMeasure_Provider(t *testing.T) {

	var tests = []struct {
		name     string
		provider Kind
		want     Kind
	}{
		{"ookla", KindOokla, KindOokla},
		{"netflix", KindNetflix, KindNetflix},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := Measure(tt.provider)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			if r.Provider() != tt.provider {
				t.Errorf("got %s, want %s", r.Provider(), tt.provider)
			}
		})
	}
}

func BenchmarkOokla_Measure(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Measure(KindOokla)
	}
}

func BenchmarkNetflix_Measure(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Measure(KindNetflix)
	}
}
