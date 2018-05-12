# grafana-webhook

grafana-webhook provides an easy way to write go http handlers for webhook channels

## Usage example

Handle Grafana request and send a message by a service:

```go

...

// sending messages server
handler := http.DefaultServeMux
handler.HandleFunc("/sms", HandleWebhook(func(w http.ResponseWriter, b *Body) {

    msg := fmt.Sprintf("Grafana status: %s\n%s", b.Title, b.Message)
    sendMessage(msg)

}, 0))

go http.ListenAndServe(addr, handler)
log.Println(fmt.Sprintf("API is listening on: %s", addr))

...

```

Above listener can be used to fill the url input in [Grafana Webhook notification channel](http://docs.grafana.org/alerting/notifications/#webhook).