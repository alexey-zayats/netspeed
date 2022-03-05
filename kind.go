package netspeed

type Kind int

const (
	// KindUnknown constant represents that no provider defined
	KindUnknown Kind = iota
	// KindOokla constant represent user chosen Ookla's speedtest.net
	KindOokla
	// KindNetflix constant represent user chosen Netfix's fast.com
	KindNetflix
)

const (
	ookla   = "ookla"
	netflix = "netflix"
	unknown = "unknown"
)

// KindFromString makes Kind from given string
func KindFromString(k string) (p Kind) {
	p = KindUnknown
	switch k {
	case ookla:
		p = KindOokla
	case netflix:
		p = KindNetflix
	}
	return
}

// String interface implementation for Kind
func (p Kind) String() (s string) {
	switch p {
	case KindOokla:
		s = ookla
	case KindNetflix:
		s = netflix
	default:
		s = unknown
	}
	return
}
