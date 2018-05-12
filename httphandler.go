package grafana_webhook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HandleWebhook returns a http handler function
// 'h' HandlerFunc parameter is called after request successfully unmarshaled to the Body pointer
func HandleWebhook(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Grafana request body
		var b *Body

		// parse POST/ PUT values to the Grafana Body model
		reqData, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqData, &b)

		// run Grafana HandlerFunc
		if h != nil {
			h(w, b)
		}
	}
}

type HandlerFunc func(w http.ResponseWriter, b *Body)
