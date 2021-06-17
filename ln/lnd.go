package ln

import (
	"context"
	"encoding/hex"
	"errors"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/lightningnetwork/lnd/lnrpc"
)

// LNDclient is an implementation of the wall.LNClient and pay.LNClient interface
// for the lnd Lightning Network node implementation.
type LNDclient struct {
	lndClient lnrpc.LightningClient
	ctx       context.Context
	conn      *grpc.ClientConn
}

// GenerateInvoice generates an invoice with the given price and memo.
func (c LNDclient) GenerateInvoice(amount int64, memo string) (Invoice, error) {
	result := Invoice{}

	// Create the request and send it
	invoice := lnrpc.Invoice{
		Memo:  memo,
		Value: amount,
	}
	stdOutLogger.Println("Creating invoice for a new API request")
	res, err := c.lndClient.AddInvoice(c.ctx, &invoice)
	if err != nil {
		return result, err
	}

	result.ImplDepID = hex.EncodeToString(res.RHash)
	result.PaymentHash = result.ImplDepID
	result.PaymentRequest = res.PaymentRequest
	return result, nil
}

// CheckInvoice takes an invoice ID (LN node implementation specific) and checks if the corresponding invoice was settled.
// An error is returned if no corresponding invoice was found.
// False is returned if the invoice isn't settled.
func (c LNDclient) CheckInvoice(id string) (bool, error) {
	// In the case of lnd, the ID is the hex encoded preimage hash.
	plainHash, err := hex.DecodeString(id)
	if err != nil {
		return false, err
	}

	stdOutLogger.Printf("Checking invoice for hash %v\n", id)

	// Get the invoice for that hash
	paymentHash := lnrpc.PaymentHash{
		RHash: plainHash,
		// Hex encoded, must be exactly 32 byte
		RHashStr: id,
	}
	invoice, err := c.lndClient.LookupInvoice(c.ctx, &paymentHash)
	if err != nil {
		return false, err
	}

	// Check if invoice was settled
	if !invoice.GetSettled() {
		return false, nil
	}
	return true, nil
}

// Pay pays the invoice and returns the preimage (hex encoded) on success, or an error on failure.
func (c LNDclient) Pay(invoice string) (string, error) {
	// Decode payment request (a.k.a. invoice).
	// TODO: Decoded values are only used for logging, so maybe make this optional to make fewer RPC calls
	payReqString := lnrpc.PayReqString{
		PayReq: invoice,
	}
	decodedPayReq, err := c.lndClient.DecodePayReq(c.ctx, &payReqString)
	if err != nil {
		return "", err
	}

	// Send payment
	sendReq := lnrpc.SendRequest{
		PaymentRequest: invoice,
	}
	stdOutLogger.Printf("Sending payment with %v Satoshis to %v (memo: \"%v\")",
		decodedPayReq.NumSatoshis, decodedPayReq.Dest