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
