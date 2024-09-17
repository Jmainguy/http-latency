package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tcnksm/go-httpstat"
)

func latency(url string, log *logrus.Logger) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(io.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	log.WithFields(logrus.Fields{
		"DNSLookup":        result.DNSLookup / time.Millisecond,
		"TCPConnection":    result.TCPConnection / time.Millisecond,
		"TLSHandshake":     result.TLSHandshake / time.Millisecond,
		"ServerProcessing": result.ServerProcessing / time.Millisecond,
		"NameLookup":       result.NameLookup / time.Millisecond,
		"Connect":          result.Connect / time.Millisecond,
		"Pretransfer":      result.Connect / time.Millisecond,
		"StartTransfer":    result.StartTransfer / time.Millisecond,
		"URL":              url,
	}).Info("Results in Milliseconds")
}

func main() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	url := os.Getenv("URL")
	for {
		latency(url, log)
		time.Sleep(1 * time.Second)
	}
}
