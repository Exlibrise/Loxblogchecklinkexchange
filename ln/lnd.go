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

// CheckInvoice takes an invoice ID (LN node implementation specific) and checks if th