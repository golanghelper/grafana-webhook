package grafana_webhook

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
	"net/http/httptest"
)

// TestHandleWebhook tests Handler function -
func TestHandleWebhook(t *testing.T) {

	t.Run("test_post_request_struct", func(t *testing.T) {
		// read file with test data
		fd, e := ioutil.ReadFile("request-example.json")
		fatalIfError(e, t)

		// prepare request handler
		h := HandleWebhook(func(w http.ResponseWriter, b *Body) {
			if b.Title != "My alert" {
				t.Fatalf("Got: %s, expected: %s", b.Title, "My alert")
			}
		})

		// prepare a request
		r, e := http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(fd))
		fatalIfError(e, t)

		// run habdler
		w := httptest.NewRecorder()
		h(w, r)
	})

}

func fatalIfError(e error, t *testing.T) {
	if e != nil {
		t.Fatal(e)
	}
}