package sipparse

var empty = struct{}{}
var acceptMethods map[string]struct{}
var discardMethods map[string]struct{}

type SIP struct {
	Title           string // Method or Status
	IsRequest       bool
	CallID          string
	RequestUsername string
	RequestDomain   string
	ToUsername      string
	ToDomain        string
	FromUsername    string
	FromDomain      string
	CseqNumber      int
	CseqMethod      string
	SrcAddr         string // IP:PORT
	DstAddr         string // IP:PORT
	CreateAt        string
	Protocol        int
	raw             *string // raw sip message
}

func init() {
	am := map[string]struct{}{
		"INVITE":    empty,
		"CANCEL":    empty,
		"ACK":       empty,
		"BYE":       empty,
		"INFO":      empty,
		"OPTIONS":   empty,
		"UPDATE":    empty,
		"REGISTER":  empty,
		"MESSAGE":   empty,
		"SUBSCRIBE": empty,
		"NOTIFY":    empty,
		"PRACK":     empty,
		"REFER":     empty,
		"PUBLISH":   empty,
	}

	// may be read from env
	dm := map[string]struct{}{
		"INFO":      empty,
		"OPTIONS":   empty,
		"REGISTER":  empty,
		"MESSAGE":   empty,
		"SUBSCRIBE": empty,
		"PUBLISH":   empty,
	}

	acceptMethods = am
	discardMethods = dm
}
