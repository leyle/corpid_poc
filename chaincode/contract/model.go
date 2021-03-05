package contract

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var ErrNoIdData = errors.New("no data for this id")

type Contract struct {
	contractapi.Contract
}

type DIDLedgerIn struct {
	Id        string          `json:"id"`
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
	CreatedAt int64           `json:"createdAt"`
	UpdatedAt int64           `json:"updatedAt"`
}

type DIDLedgerOut struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
