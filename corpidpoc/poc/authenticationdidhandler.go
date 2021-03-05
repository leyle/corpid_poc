package poc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/leyle/corpid_poc/corpidpoc/context"
	"github.com/leyle/go-api-starter/ginhelper"
)

// Authentication DID handlers
func CreateAuthenticationDIDhandler(ctx *context.ApiContext) {
	var form AuthenticationDID
	err := ctx.C.BindJSON(&form)
	ginhelper.StopExec(err)

	dataId := form.Id
	data, err := jsoniter.MarshalToString(form)
	ginhelper.StopExec(err)

	stateReq := &LedgerState{
		Id:         dataId,
		Type:       AuthenticationDIDType,
		DataString: data,
	}

	err = createState(ctx, stateReq)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, dataId)
	return
}

func GetAuthenticationDIDInfoHandler(ctx *context.ApiContext) {
	dataId := ctx.C.Param("id")
	state, err := getStateById(ctx, dataId)
	ginhelper.StopExec(err)

	if state.Type != AuthenticationDIDType {
		err = fmt.Errorf("Invalid data type[%s], expect[%s], maybe called wrong api", state.Type, AuthenticationDIDType)
		ginhelper.StopExec(err)
	}

	var data AuthenticationDID
	err = jsoniter.UnmarshalFromString(state.DataString, &data)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, data)
	return
}

func QueryAuthenticationDIDHandler(ctx *context.ApiContext) {
	states, err := getAllStates(ctx)
	ginhelper.StopExec(err)

	var datas []*AuthenticationDID
	for _, state := range states {
		if state.Type == AuthenticationDIDType {
			var data AuthenticationDID
			err = jsoniter.UnmarshalFromString(state.DataString, &data)
			ginhelper.StopExec(err)

			datas = append(datas, &data)
		}
	}

	ginhelper.ReturnOKJson(ctx.C, datas)
	return
}
