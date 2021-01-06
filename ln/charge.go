
package ln

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ChargeClient is an implementation of the wall.LNclient interface for "Lightning Charge"
// running on top of the c-lightning Lightning Network node implementation.
type ChargeClient struct {
	client   *http.Client
	baseURL  string
	apiToken string
}

// GenerateInvoice generates an invoice with the given price and memo.
func (c ChargeClient) GenerateInvoice(amount int64, memo string) (Invoice, error) {
	result := Invoice{}

	data := make(url.Values)
	// Possible values as documented in https://github.com/ElementsProject/lightning-charge/blob/master/README.md:
	// msatoshi, currency, amount, description, expiry, metadata and webhook
	// But with *either* msatoshi *or* currency + amount
	mSatoshi := strconv.FormatInt(1000*amount, 10)
	data.Add("msatoshi", mSatoshi)
	data.Add("description", memo)

	// Send request
	req, err := http.NewRequest("POST", c.baseURL+"/invoice", strings.NewReader(data.Encode()))
	if err != nil {
		return result, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api-token", c.apiToken) // This might seem strange, but it's how Lightning Charge expects it
	stdOutLogger.Println("Creating invoice for a new API request")
	res, err := c.client.Do(req)
	if err != nil {
		return result, err
	}

	// Read and deserialize response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	err = res.Body.Close()
	if err != nil {
		return result, err
	}
	invoice, err := deserializeInvoice(body)
	if err != nil {
		return result, err
	}

	result.ImplDepID = invoice.ID
	result.PaymentHash = invoice.Rhash
	result.PaymentRequest = invoice.Payreq
	return result, nil
}

// CheckInvoice takes an invoice ID (LN node implementation specific) and checks if the corresponding invoice was settled.
// An error is returned if the invoice info couldn't be fetched from Lightning Charge or deserialized etc.
// False is returned if the invoice isn't settled.
func (c ChargeClient) CheckInvoice(id string) (bool, error) {
	stdOutLogger.Printf("Checking invoice %v\n", id)

	// Fetch invoice
	req, err := http.NewRequest("GET", c.baseURL+"/invoice/"+id, nil)
	if err != nil {
		return false, err
	}
	req.SetBasicAuth("api-token", c.apiToken) // This might seem strange, but it's how Lightning Charge expects it
	res, err := c.client.Do(req)
	if err != nil {
		return false, err
	}

	invoiceJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	invoice, err := deserializeInvoice(invoiceJSON)
	if err != nil {
		return false, err
	}

	if invoice.Status == "unpaid" {
		return false, nil