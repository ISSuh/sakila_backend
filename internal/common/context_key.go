package common

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	ReqeustContextKey  = contextKey("request")
	ResponseContextKey = contextKey("response")
)
