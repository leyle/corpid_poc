package poc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/leyle/corpid_poc/corpidpoc/context"
	"github.com/leyle/go-api-starter/ginhelper"
)

// create credential data schema
func CreateCredentialDataHandler(ctx *context.ApiContext) {
	var form CredentialDataSchema
	err := ctx.C.BindJSON(&form)
	ginhelper.StopExec(err)

	// convert input data into json string
	dataId := form.Id
	data, err := jsoniter.MarshalToString(form)
	ginhelper.StopExec(err)

	stateReq := &LedgerState{
		Id:         dataId,
		Type:       CredentialDataType,
		DataString: data,
	}

	err = createState(ctx, stateReq)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, dataId)
	return
}

func GetCredentialDataInfoHandler(ctx *context.ApiContext) {
	dataId := ctx.C.Param("id")

	state, err := getStateById(ctx, dataId)
	ginhelper.StopExec(err)

	// convert data string to go struct
	if state.Type != CredentialDataType {
		err = fmt.Errorf("Invalid data type[%s], expect[%s], maybe call wrong api", state.Type, CredentialDataType)
		ginhelper.StopExec(err)
	}

	var data CredentialDataSchema
	err = jsoniter.UnmarshalFromString(state.DataString, &data)
	ginhelper.StopExec(err)

	ginhelper.ReturnOKJson(ctx.C, data)
	return
}

func QueryCredentialDataHandler(ctx *context.ApiContext) {
	states, err := getAllStates(ctx)
	ginhelper.StopExec(err)

	var datas []*CredentialDataSchema
	for _, state := range states {
		if state.Type == CredentialDataType {
			var data CredentialDataSchema
			err = jsoniter.UnmarshalFromString(state.DataString, &data)
			ginhelper.StopExec(err)

			datas = append(datas, &data)
		}
	}

	ginhelper.ReturnOKJson(ctx.C, datas)
	return
}
