package poc

import (
	"encoding/json"
	"github.com/leyle/corpid_poc/corpidpoc/context"
)

func createState(ctx *context.ApiContext, form *LedgerState) error {
	contractRet := GetContract(ctx, ctx.Cfg.Chaincode.Channel, ctx.Cfg.Chaincode.Chaincode)
	if contractRet.Err != nil {
		return contractRet.Err
	}
	defer contractRet.Close()
	contract := contractRet.Contract

	args, err := json.Marshal(form)
	if err != nil {
		ctx.Logger().Error().Err(err).Msg("marshal create state form failed")
		return err
	}

	ret, err := contract.SubmitTransaction(MethodCreate, string(args))
	if err != nil {
		ctx.Logger().Error().Err(err).Msg("submit create state failed")
		return err
	}
	ctx.Logger().Info().Str("id", form.Id).Str("resp", string(ret)).Msg("create state success")

	return nil
}

func getStateById(ctx *context.ApiContext, id string) (*LedgerState, error) {
	contractRet := GetContract(ctx, ctx.Cfg.Chaincode.Channel, ctx.Cfg.Chaincode.Chaincode)
	if contractRet.Err != nil {
		return nil, contractRet.Err
	}
	defer contractRet.Close()
	contract := contractRet.Contract

	ret, err := contract.EvaluateTransaction(MethodGetById, id)
	if err != nil {
		ctx.Logger().Error().Err(err).Str("id", id).Msg("get state by id failed")
		return nil, err
	}

	var state LedgerState
	err = json.Unmarshal(ret, &state)
	if err != nil {
		ctx.Logger().Error().Err(err).Str("id", id).Msg("get state by id, unmarshal failed")
		return nil, err
	}

	return &state, nil
}

func getAllStates(ctx *context.ApiContext) ([]*LedgerState, error) {
	contractRet := GetContract(ctx, ctx.Cfg.Chaincode.Channel, ctx.Cfg.Chaincode.Chaincode)
	if contractRet.Err != nil {
		return nil, contractRet.Err
	}
	defer contractRet.Close()
	contract := contractRet.Contract

	rets, err := contract.EvaluateTransaction(MethodGetAll)
	if err != nil {
		ctx.Logger().Error().Err(err).Msg("get all states failed")
		return nil, err
	}

	var states []*LedgerState
	err = json.Unmarshal(rets, &states)
	if err != nil {
		ctx.Logger().Error().Err(err).Msg("get all states, unmarshal failed")
		return nil, err
	}

	return states, nil
}
