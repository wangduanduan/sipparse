package sipparse

const (
	MInvite    = "INVITE"
	MCancel    = "CANCEL"
	MAck       = "ACK"
	MBye       = "BYE"
	MInfo      = "INFO"
	MOptions   = "OPTIONS"
	MUpdate    = "UPDATE"
	MRegister  = "REGISTER"
	MMessage   = "MESSAGE"
	MSubscribe = "SUBSCRIBE"
	MNotify    = "NOTIFY"
	MPrack     = "PRACK"
	MRefer     = "REFER"
	MPublish   = "PUBLISH"
)

type SIPURI struct {
	user   string
	domain string
}

type SIPCSEQ struct {
	num    int
	method string
}

type SIPMSG struct {
	IsRequest  bool
	CallID     string
	RequestURL SIPURI
	FromURL    SIPURI
	ToRUL      SIPURI
	CSEQ       SIPCSEQ
	raw        string
}
