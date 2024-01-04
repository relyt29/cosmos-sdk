package cometbft

import (
	coreappmgr "cosmossdk.io/server/v2/core/appmanager"
	"cosmossdk.io/server/v2/core/event"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/proto"
)

// TODO
func parseQueryResponse(_ proto.Message) (*abci.ResponseQuery, error) {
	return &abci.ResponseQuery{}, nil
}

func parseFinalizeBlockResponse(in *coreappmgr.BlockResponse, err error) (*abci.ResponseFinalizeBlock, error) {
	allEvents := append(in.BeginBlockEvents, in.EndBlockEvents...)

	resp := &abci.ResponseFinalizeBlock{
		Events:                intoABCIEvents(allEvents),
		TxResults:             intoABCITxResults(in.TxResults),
		ValidatorUpdates:      intoABCIValidatorUpdates(in.ValidatorUpdates),
		AppHash:               in.Apphash,
		ConsensusParamUpdates: nil, // TODO: figure out consensus params here, maybe parse the tx responses or events?
	}
	return resp, nil
}

func intoABCIValidatorUpdates(updates []coreappmgr.ValidatorUpdate) []abci.ValidatorUpdate {
	valsetUpdates := make([]abci.ValidatorUpdate, len(updates))

	for i := range updates {
		valsetUpdates[i] = abci.ValidatorUpdate{
			PubKey: cmtprotocrypto.PublicKey{
				Sum: &cmtprotocrypto.PublicKey_Ed25519{ // TODO: check if this is ok
					Ed25519: updates[i].PubKey,
				},
			},
			Power: updates[i].Power,
		}
	}

	return valsetUpdates
}

func intoABCITxResults(results []coreappmgr.TxResult) []*abci.ExecTxResult {
	res := make([]*abci.ExecTxResult, len(results))
	for i := range results {
		if results[i].Error == nil {
			res[i] = sdkerrors.ResponseExecTxResultWithEvents(
				results[i].Error,
				0, // TODO: gas wanted?
				results[i].GasUsed,
				intoABCIEvents(results[i].Events),
				false,
			)
			continue
		}

		// TODO: handle properly once the we decide on the type of TxResult.Resp
	}

	return res
}

func intoABCIEvents(events []event.Event) []abci.Event {
	abciEvents := make([]abci.Event, len(events))
	for i := range events {
		abciEvents[i] = abci.Event{
			Type:       events[i].Type,
			Attributes: intoABCIAttributes(events[i].Attributes),
		}
	}
	return abciEvents
}

func intoABCIAttributes(attributes []event.Attribute) []abci.EventAttribute {
	abciAttributes := make([]abci.EventAttribute, len(attributes))
	for i := range attributes {
		abciAttributes[i] = abci.EventAttribute{
			Key:   "",
			Value: "",
			Index: false, // TODO: who holds this config?
		}
	}
	return abciAttributes
}
