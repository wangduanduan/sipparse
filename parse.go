package sipparse

import (
	"strings"
)

var empty = struct{}{}
var acceptMethods map[string]struct{}
var discardMethods map[string]struct{}

const CRLF = "\r\n"

const (
	ParseOk = iota
	ECanNotFindHeader
	EBadHeaderValue
)

const EmptyStr = ""
const (
	HeaderCallID = "Call-ID"
	HeaderFrom   = "From"
	HeaderTo     = "To"
	HeaderUA     = "User-Agent"
	HeaderCSeq   = "CSeq"
)

type SIP struct {
	Title           string // Method or Status
	IsRequest       bool
	CallID          string
	RequestURL      string
	RequestUsername string
	RequestDomain   string
	ToUsername      string
	ToDomain        string
	FromUsername    string
	FromDomain      string
	CSeqNumber      string
	CSeqMethod      string
	UserAgent       string
	SrcAddr         string // IP:PORT
	DstAddr         string // IP:PORT
	CreateAt        string
	Protocol        int
	UID             string  // correlative id for AB call leg
	FSCallID        string  // freeswitch CallID
	raw             *string // raw sip message
}

// 	Request 	: INVITE bob@example.com SIP/2.0
// 	Response 	: SIP/2.0 200 OK
// 	Response	: SIP/2.0 501 Not Implemented
func (p *SIP) ParseFirstLine() {
	if *p.raw == EmptyStr {
		return
	}

	firstLineIndex := strings.Index(*p.raw, CRLF)
	if firstLineIndex == -1 {
		return
	}
	firstLine := (*p.raw)[:firstLineIndex]
	firstLineMeta := strings.SplitN(firstLine, " ", 3)

	if len(firstLineMeta) != 3 {
		return
	}
	if strings.HasPrefix(firstLineMeta[0], "SIP") {
		p.IsRequest = false
		p.Title = firstLineMeta[1]
		return
	}
	p.IsRequest = true
	p.Title = firstLineMeta[0]
	p.RequestURL = firstLineMeta[1]
}

func (p *SIP) ParseFrom() {}
func (p *SIP) ParseRequestURL() {
	if p.RequestURL == "" {
		return
	}

}
func (p *SIP) ParseTo()        {}
func (p *SIP) ParseUserAgent() {}

func ParseSIPURL(s string) (string, string) {
	if s == "" {
		return "", ""
	}

	newURL := s

	if strings.Contains(s, "<") {
		start := strings.Index(s, "<")
		end := strings.Index(s, ">")
		if start > end {
			return "", ""
		}
		newURL = s[start:end]
	}

	a := strings.Index(newURL, ":")
	b := strings.Index(newURL, "@")
	c := strings.Index(newURL, ";")

	if a == -1 || b == -1 || a > b {
		return "", ""
	}

	if c == -1 {
		c = len(newURL)
	}

	user := newURL[a+1 : b]
	domain := newURL[b+1 : c]
	return user, domain
}

func (p *SIP) ParseCseq() {
	cseqValue := p.GetHeaderValue(HeaderCSeq)
	if cseqValue == EmptyStr {
		return
	}
	cs := strings.SplitN(cseqValue, " ", 2)
	if len(cs) != 2 {
		return
	}
	p.CSeqNumber = cs[0]
	p.CSeqMethod = cs[1]
}

func (p *SIP) GetHeaderValue(header string) (v string) {
	if *p.raw == EmptyStr || header == EmptyStr {
		return EmptyStr
	}

	if strings.Contains(header, CRLF) || strings.Contains(header, " ") {
		return EmptyStr
	}

	startIndex := strings.Index(*p.raw, header+":")

	if startIndex == -1 {
		return EmptyStr
	}

	newStr := (*p.raw)[startIndex:]

	endIndex := strings.Index(newStr, CRLF)

	if endIndex == -1 {
		return EmptyStr
	}

	if len(header)+1 > endIndex {
		return EmptyStr
	}

	return strings.TrimSpace(newStr[len(header)+1 : endIndex])
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
