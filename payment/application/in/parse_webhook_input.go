package in

import "net/http"

type ParseWebhook struct {
	Payload []byte
	Headers http.Header
}
