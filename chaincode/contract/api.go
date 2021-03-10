package contract

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

// create new record
// args is a json string from CreateForm struct
type CreateFrom struct {
	// id comes from restful api input
	Id string `json:"id"`

	// data is a json string
	Data string `json:"data"`

	// type is data type, e.g CredentialDataSchema / AuthenticationDID / CredentialDID
	// type is used by restful api for unmarshal data
	Type string `json:"type"`
}

func (c *Contract) Create(ctx contractapi.TransactionContextInterface, args string) error {
	var form CreateFrom
	err := json.Unmarshal([]byte(args), &form)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// check if data has already existed
	/*
		exists, err := c.IsExist(ctx, form.Id)
		if err != nil {
			return err
		}
		if exists {
			err = fmt.Errorf("data id[%s] has existed, create new data failed", form.Id)
			fmt.Println(err.Error())
			return err
		}
	*/

	// save data
	state := &DIDLedgerIn{
		Id:        form.Id,
		Type:      form.Type,
		Data:      []byte(form.Data),
		CreatedAt: time.Now().Unix(),
	}
	state.UpdatedAt = state.CreatedAt

	stateJson, err := json.Marshal(state)
	if err != nil {
		fmt.Println("create new data, marshal state failed, ", err.Error())
		return err
	}

	err = ctx.GetStub().PutState(form.Id, stateJson)
	if err != nil {
		fmt.Println("create new data, save to ledger failed, ", err.Error())
		return err
	}

	return nil
}

// get by id
func (c *Contract) GetById(ctx contractapi.TransactionContextInterface, id string) (*DIDLedgerOut, error) {
	stateByte, err := ctx.GetStub().GetState(id)
	if err != nil {
		fmt.Println("get state by id failed, ", id, err.Error())
		return nil, err
	}
	if stateByte == nil {
		fmt.Println("get state by id, no data with that id", id)
		return nil, ErrNoIdData
	}

	// unmarshal
	var state DIDLedgerIn
	err = json.Unmarshal(stateByte, &state)
	if err != nil {
		fmt.Println("get state by id, unmashal bytes failed", id, err.Error())
		return nil, err
	}

	// convert business data to string
	stateOut := copyToOut(&state)

	return stateOut, nil
}

func (c *Contract) GetAll(ctx contractapi.TransactionContextInterface) ([]*DIDLedgerOut, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		fmt.Println("get all state failed, ", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()

	var states []*DIDLedgerOut
	for resultsIterator.HasNext() {
		resp, err := resultsIterator.Next()
		if err != nil {
			fmt.Println("get all state, iterate next failed, ", err.Error())
			return nil, err
		}

		var state DIDLedgerIn
		err = json.Unmarshal(resp.Value, &state)
		if err != nil {
			fmt.Println("get all state, unmarshal state failed", err.Error())
			return nil, err
		}
		stateOut := copyToOut(&state)
		states = append(states, stateOut)
	}

	return states, nil
}

func (c *Contract) IsExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	state, err := ctx.GetStub().GetState(id)
	if err != nil {
		fmt.Println("check if data existed failed, id", id, err.Error())
		return false, err
	}

	if state == nil {
		return false, nil
	}
	return true, nil
}

func copyToOut(in *DIDLedgerIn) *DIDLedgerOut {
	if in == nil {
		return nil
	}

	out := &DIDLedgerOut{
		Id:        in.Id,
		Type:      in.Type,
		Data:      string(in.Data),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return out
}
