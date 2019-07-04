package ipset

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var httpClient = new(http.Client)

//NotifyIPSetAPI notifies the ipset-api that new IP Addresses should be included in the blacklist
func NotifyIPSetAPI() error {
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		return errors.Wrap(err, "Failed to create a new HTTP Request")
	}

	req.Header.Set("Content-Type", "application/json")
	r, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "Failed accomplish the HTTP Request")
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(errors.Wrap(err, "Failed to read the HTTP response body"))
		return nil
	}

	log.WithFields(log.Fields{
		"status":  r.Status,
		"headers": r.Header,
		"body":    string(b),
	}).Info("Received response from the `ipset-api`")

	return nil
}
