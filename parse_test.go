package sipparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSeq(t *testing.T) {
	successCases := []struct {
		in                 string
		expectedCseqNumber string
		expectedCseqMethod string
	}{
		{
			in:                 "CSeq: 123 INVITE\r\n",
			expectedCseqNumber: "123",
			expectedCseqMethod: "INVITE",
		},
		{
			in:                 "CSeq: 123 INVITE\r\n",
			expectedCseqNumber: "123",
			expectedCseqMethod: "INVITE",
		},
	}

	for _, item := range successCases {
		sip := SIP{
			raw: &item.in,
		}
		sip.ParseCseq()
		assert.Equal(t, item.expectedCseqMethod, sip.CSeqMethod)
		assert.Equal(t, item.expectedCseqNumber, sip.CSeqNumber)
	}
}

func TestPraseFirstLine(t *testing.T) {
	successCases := []struct {
		in            string
		header        string
		expectedValue string
	}{
		{
			in:            "CSeq: 123 INVITE\r\n",
			header:        "CSeq",
			expectedValue: "123 INVITE",
		},
		{
			in:            "Call-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "Call-ID",
			expectedValue: "1234",
		},
		{
			in:            "Call-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "Call-ID\r\n",
			expectedValue: "",
		},
		{
			in:            "User-Agent: wdd\r\nCall-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "",
			expectedValue: "",
		},
		{
			in:            "User-Agent: wdd\r\nCall-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "User-Agent",
			expectedValue: "wdd",
		},
		{
			in:            "",
			header:        "User-Agent",
			expectedValue: "",
		},
	}

	for _, item := range successCases {
		sip := SIP{
			raw: &item.in,
		}
		assert.Equal(t, item.expectedValue, sip.GetHeaderValue(item.header))
	}
}

func TestGetHeaderValue(t *testing.T) {
	successCases := []struct {
		in            string
		header        string
		expectedValue string
	}{
		{
			in:            "CSeq: 123 INVITE\r\n",
			header:        "CSeq",
			expectedValue: "123 INVITE",
		},
		{
			in:            "Call-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "Call-ID",
			expectedValue: "1234",
		},
		{
			in:            "Call-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "Call-ID\r\n",
			expectedValue: "",
		},
		{
			in:            "User-Agent: wdd\r\nCall-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "",
			expectedValue: "",
		},
		{
			in:            "User-Agent: wdd\r\nCall-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "User-Agent",
			expectedValue: "wdd",
		},
		{
			in:            "User-Agent: wdd\r\nCall-ID: 1234\r\nCSeq: 123 INVITE\r\n",
			header:        "Call-ID",
			expectedValue: "1234",
		},
		{
			in:            "",
			header:        "User-Agent",
			expectedValue: "",
		},
	}

	for _, item := range successCases {
		sip := SIP{
			raw: &item.in,
		}
		assert.Equal(t, item.expectedValue, sip.GetHeaderValue(item.header))
	}
}

func TestParseSIPURL(t *testing.T) {
	successCases := []struct {
		msg    string
		user   string
		domain string
	}{
		{"<sip:1002@192.168.159.12>;tag=feffa1b1ce68471d8e0a97eb9a1dcb32", "1002", "192.168.159.12"},
		{"<sip:1002@192.168.159.12>", "1002", "192.168.159.12"},
		{"\"800004\" <sip:800004@001.com>;tag=v2Gl2kKLlTNNC.Nydij-ri02clU52sTZ", "800004", "001.com"},
		{"sip:800004@001.com", "800004", "001.com"},
		{"Carol <sip:carol@chicago.com>", "carol", "chicago.com"},
		{"sip:+12125551212@phone2net.com;tag=887s", "+12125551212", "phone2net.com"},
	}

	for _, c := range successCases {
		user, domain := ParseSIPURL(c.msg)
		assert.Equal(t, c.user, user)
		assert.Equal(t, c.domain, domain)
	}
}
