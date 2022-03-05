package netspeed

import "testing"

func TestKind_String(t *testing.T) {
	tests := []struct {
		name string
		kind Kind
		want string
	}{
		{"Unknown", KindUnknown, unknown},
		{"Netflix", KindNetflix, netflix},
		{"OOkla", KindOokla, ookla},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := tt.kind.String()
			if given != tt.want {
				t.Errorf("want %s, given %s", tt.want, given)
			}
		})
	}
}

func TestKindFromString(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want Kind
	}{
		{"Unknown", unknown, KindUnknown},
		{"Netflix", netflix, KindNetflix},
		{"OOkla", ookla, KindOokla},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := KindFromString(tt.kind)
			if given != tt.want {
				t.Errorf("want %s, given %s", tt.want, given)
			}
		})
	}
}
