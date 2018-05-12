package grafana_webhook

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

// HandleWebhook returns a http handler function
// 'h' HandlerFunc parameter is called after request successfully unmarshaled to the Body pointer
func HandleWebhook(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Grafana request body
		var b *Body

		// parse POST/ PUT values to the Grafana Body model
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			// request limit: @TODO - make it configurable, the limit depends on the matches amount...
			r.Body = http.MaxBytesReader(w, r.Body, 8192)
			reqData, e := ioutil.ReadAll(r.Body)
			if e != nil {
				// read body action has failed
				b = BodyOnReadAllSizeLimitErr()
			} else {
				json.Unmarshal(reqData, &b)
			}
			// run Grafana HandlerFunc
			if h != nil {
				h(w, b)
			}
		}
	}
}

type HandlerFunc func(w http.ResponseWriter, b *Body)
