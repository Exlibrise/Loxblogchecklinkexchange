package ln

import (
	"context"
	"encoding/hex"
	"errors"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/ligh