package poc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/leyle/corpid_poc/corpidpoc/context"
	"github.com/leyle/go-api-starter/ginhelper"
)

// Authentication DID handlers
func CreateCredentialDIDhandler(ctx *context.ApiContext) {
	var form CredentialDID
	err := ctx.C.BindJSON(&form)
	ginhelper.StopExec(err)

	dataId := form.Id
	data, err := jsoniter.MarshalToString(form)
	ginhelper.StopExec(err)

	stateReq := &LedgerState{
		Id:         dataId,
		Type:       CredentialDIDType,
		DataString: data,
	}

	err = createState(ctx, stateReq)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, dataId)
	return
}

func GetCredentialDIDInfoHandler(ctx *context.ApiContext) {
	dataId := ctx.C.Param("id")
	state, err := getStateById(ctx, dataId)
	ginhelper.StopExec(err)

	if state.Type != CredentialDIDType {
		err = fmt.Errorf("Invalid data type[%s], expect[%s], maybe called wrong api", state.Type, CredentialDIDType)
		ginhelper.StopExec(err)
	}

	var data CredentialDID
	err = jsoniter.UnmarshalFromString(state.DataString, &data)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, data)
	return
}

func QueryCredentialDIDHandler(ctx *context.ApiContext) {
	states, err := getAllStates(ctx)
	ginhelper.StopExec(err)

	var datas []*CredentialDID
	for _, state := range states {
		if state.Type == CredentialDIDType {
			var data CredentialDID
			err = jsoniter.UnmarshalFromString(state.DataString, &data)
			ginhelper.StopExec(err)

			datas = append(datas, &data)
		}
	}

	ginhelper.ReturnOKJson(ctx.C, datas)
	return
}
