package client

type Resolver string

func (s Resolver) String() string {
	return string(s)
}

const (
	DNSResolver    = Resolver("DNS")
	ConsulResolver = Resolver("CONSUL")
)

const (
	XRequestID = "RequestID"
	sessionKey = "session_key"
)
