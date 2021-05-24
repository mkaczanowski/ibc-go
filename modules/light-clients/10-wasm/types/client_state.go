package types

import (
	"encoding/json"
	"fmt"
	"github.com/CosmWasm/wasmvm/api"
	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

func (c *ClientState) Initialize(context sdk.Context, marshaler codec.BinaryCodec, store sdk.KVStore, state exported.ConsensusState) error {
	const InitializeState = "initializestate"
	payload := make(map[string]map[string]interface{})
	payload[InitializeState] = make(map[string]interface{})
	inner := payload[InitializeState]
	inner["me"] = c
	inner["consensus_state"] = state

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	// Under the hood there are two calls to wasm contract for initialization as by design
	// cosmwasm does not allow init call to return any value.

	_, err = initContract(c.CodeId, context, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToInit, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToCall, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if !output.Result.IsValid {
		return fmt.Errorf("%s error ocurred while initializing client state", output.Result.ErrorMsg)
	}
	output.resetImmutables(c)

	*c = *output.Me
	return nil
}

func (c *ClientState) CheckHeaderAndUpdateState(context sdk.Context, marshaler codec.BinaryCodec, store sdk.KVStore, header exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	consensusState, err := GetConsensusState(store, marshaler, c.LatestHeight)
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(err, "could not get trusted consensus state from clientStore for Header at Height: %s", header.GetHeight())
	}

	const CheckHeaderAndUpdateState = "checkheaderandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[CheckHeaderAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckHeaderAndUpdateState]
	inner["me"] = c
	inner["header"] = header
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToCall, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if !output.Result.IsValid {
		return nil, nil, fmt.Errorf("%s error ocurred while updating client state", output.Result.ErrorMsg)
	}
	output.resetImmutables(c)
	return output.NewClientState, output.NewConsensusState, nil
}

func (c *ClientState) CheckMisbehaviourAndUpdateState(context sdk.Context, marshaler codec.BinaryCodec, store sdk.KVStore, misbehaviour exported.Misbehaviour) (exported.ClientState, error) {
	wasmMisbehaviour, ok := misbehaviour.(*Misbehaviour)
	if !ok {
		return nil, sdkerrors.Wrapf(
			clienttypes.ErrInvalidMisbehaviour,
			"invalid misbehaviour type %T, expected %T", wasmMisbehaviour, &Misbehaviour{},
		)
	}

	// Get consensus bytes from clientStore
	consensusState1, err := GetConsensusState(store, marshaler, wasmMisbehaviour.Header1.Height)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "could not get trusted consensus state from clientStore for Header1 at Height: %s", wasmMisbehaviour.Header1)
	}

	// Get consensus bytes from clientStore
	consensusState2, err := GetConsensusState(store, marshaler, wasmMisbehaviour.Header2.Height)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "could not get trusted consensus state from clientStore for Header2 at Height: %s", wasmMisbehaviour.Header2)
	}

	const CheckMisbehaviourAndUpdateState = "checkmisbehaviourandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[CheckMisbehaviourAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckMisbehaviourAndUpdateState]
	inner["me"] = c
	inner["misbehaviour"] = wasmMisbehaviour
	inner["consensus_state1"] = consensusState1
	inner["consensus_state2"] = consensusState2

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	out, err := callContract(c.CodeId, context, store, encodedData)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToCall, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if !output.Result.IsValid {
		return nil, fmt.Errorf("%s error ocurred while updating client state", output.Result.ErrorMsg)
	}
	output.resetImmutables(c)
	return output.NewClientState, nil
}

func (c *ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore,
	substituteClientStore sdk.KVStore, substituteClient exported.ClientState,
	initialHeight exported.Height,
) (exported.ClientState, error) {
	var (
		SubjectPrefix    = []byte("subject/")
		SubstitutePrefix = []byte("substitute/")
	)

	consensusState, err := GetConsensusState(subjectClientStore, cdc, c.LatestHeight)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			err, "unexpected error: could not get consensus state from clientstore at height: %d", c.GetLatestHeight(),
		)
	}

	store := NewWrappedStore(subjectClientStore, subjectClientStore, SubjectPrefix, SubstitutePrefix)

	const CheckSubstituteAndUpdateState = "checksubstituteandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[CheckSubstituteAndUpdateState] = make(map[string]interface{})
	inner := payload[CheckSubstituteAndUpdateState]
	inner["me"] = c
	inner["subject_consensus_state"] = consensusState
	inner["substitute_client_state"] = substituteClient
	inner["initial_height"] = initialHeight

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	out, err := callContract(c.CodeId, ctx, store, encodedData)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToCall, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		return nil, sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if !output.Result.IsValid {
		return nil, fmt.Errorf("%s error ocurred while updating client state", output.Result.ErrorMsg)
	}

	output.resetImmutables(c)
	return output.NewClientState, nil
}

func (c *ClientState) VerifyUpgradeAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, store sdk.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient, proofUpgradeConsState []byte) (exported.ClientState, exported.ConsensusState, error) {
	wasmUpgradeConsState, ok := newConsState.(*ConsensusState)
	if !ok {
		return nil, nil, sdkerrors.Wrapf(clienttypes.ErrInvalidConsensus, "upgraded consensus state must be Tendermint consensus state. expected %T, got: %T",
			&ConsensusState{}, wasmUpgradeConsState)
	}

	// last height of current counterparty chain must be client's latest height
	lastHeight := c.LatestHeight
	lastHeightConsensusState, err := GetConsensusState(store, cdc, lastHeight)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "could not retrieve consensus state for lastHeight")
	}

	const VerifyUpgradeAndUpdateState = "verifyupgradeandupdatestate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyUpgradeAndUpdateState] = make(map[string]interface{})
	inner := payload[VerifyUpgradeAndUpdateState]
	inner["me"] = c
	inner["new_client_state"] = newClient
	inner["new_consensus_state"] = newConsState
	inner["client_upgrade_proof"] = proofUpgradeClient
	inner["consensus_state_upgrade_proof"] = proofUpgradeConsState
	inner["last_height_consensus_state"] = lastHeightConsensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	out, err := callContract(c.CodeId, ctx, store, encodedData)
	if err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToCall, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		return nil, nil, sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if !output.Result.IsValid {
		return nil, nil, fmt.Errorf("%s error ocurred while updating client state", output.Result.ErrorMsg)
	}
	output.resetImmutables(c)
	return output.NewClientState, output.NewConsensusState, nil
}

func (c *ClientState) ZeroCustomFields() exported.ClientState {
	const ZeroCustomFields = "zerocustomfields"
	payload := make(map[string]map[string]interface{})
	payload[ZeroCustomFields] = make(map[string]interface{})
	inner := payload[ZeroCustomFields]
	inner["me"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}

	gasMeter := sdk.NewGasMeter(0)
	out, err := callContractWithEnvAndMeter(c.CodeId, nil, &FailKVStore{}, api.MockEnv(), gasMeter, encodedData)
	if err != nil {
		// TODO: Handle error
	}
	output := clientStateCallResponse{}
	if err := json.Unmarshal(out.Data, &output); err != nil {
		// TODO: Handle error
	}
	output.resetImmutables(c)
	return output.Me
}

/**
Following functions only queries the state so should be part of query call
*/

func (c *ClientState) ClientType() string {
	return c.Type
}

func (c *ClientState) ExportMetadata(store sdk.KVStore) []exported.GenesisMetadata {
	const ExportMetadataQuery = "exportmetadata"
	payload := make(map[string]map[string]interface{})
	payload[ExportMetadataQuery] = make(map[string]interface{})
	inner := payload[ExportMetadataQuery]
	inner["me"] = c

	encodedData, err := json.Marshal(payload)
	if err != nil {
		// TODO: Handle error
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		// TODO: Handle error
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		// TODO: Handle error
	}

	genesisMetadata := make([]exported.GenesisMetadata, len(output.GenesisMetadata))
	for i, metadata := range output.GenesisMetadata {
		genesisMetadata[i] = metadata
	}
	return genesisMetadata
}

func (c *ClientState) GetLatestHeight() exported.Height {
	return c.LatestHeight
}

func (c *ClientState) Status(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec) exported.Status {
	consensusState, err := GetConsensusState(store, cdc, c.LatestHeight)
	if err != nil {
		return exported.Unknown
	}

	const Status = "status"
	payload := make(map[string]map[string]interface{})
	payload[Status] = make(map[string]interface{})
	inner := payload[Status]
	inner["me"] = c
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return exported.Unknown
	}

	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return exported.Unknown
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return exported.Unknown
	}

	return output.Status
}

func (c *ClientState) Validate() error {
	if c.Data == nil || len(c.Data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}

	if c.CodeId == nil || len(c.CodeId) == 0 {
		return fmt.Errorf("codeid cannot be empty")
	}

	return nil
}

func (c *ClientState) GetProofSpecs() []*ics23.ProofSpec {
	return c.ProofSpecs
}

func (c *ClientState) VerifyClientState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, counterpartyClientIdentifier string, proof []byte, clientState exported.ClientState) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}
	const VerifyClientStateQuery = "verifyclientstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyClientStateQuery] = make(map[string]interface{})
	inner := payload[VerifyClientStateQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["counterparty_client_identifier"] = counterpartyClientIdentifier
	inner["proof"] = proof
	inner["counterparty_client_state"] = clientState
	inner["consensus_state"] = consensusState
	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while validating client state", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyClientConsensusState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, counterpartyClientIdentifier string, consensusHeight exported.Height, prefix exported.Prefix, proof []byte, consensusState exported.ConsensusState) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyClientConsensusStateQuery = "verifyclientconsensusstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyClientConsensusStateQuery] = make(map[string]interface{})
	inner := payload[VerifyClientConsensusStateQuery]
	inner["me"] = c
	inner["height"] = height
	inner["consensus_height"] = consensusHeight
	inner["commitment_prefix"] = prefix
	inner["counterparty_client_identifier"] = counterpartyClientIdentifier
	inner["proof"] = proof
	inner["counterparty_consensus_state"] = consensusState
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying consensus state", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyConnectionState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, connectionID string, connectionEnd exported.ConnectionI) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyConnectionStateQuery = "verifyconnectionstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyConnectionStateQuery] = make(map[string]interface{})
	inner := payload[VerifyConnectionStateQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["connection_id"] = connectionID
	inner["connection_end"] = connectionEnd
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying connection state", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyChannelState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, portID, channelID string, channel exported.ChannelI) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyChannelStateQuery = "verifychannelstate"
	payload := make(map[string]map[string]interface{})
	payload[VerifyChannelStateQuery] = make(map[string]interface{})
	inner := payload[VerifyChannelStateQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["channel"] = channel
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying channel state", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyPacketCommitment(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64, commitmentBytes []byte) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyPacketCommitmentQuery = "verifypacketcommitment"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketCommitmentQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketCommitmentQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence
	inner["commitment_bytes"] = commitmentBytes
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying packet commitment", output.Result.ErrorMsg)
	}

}

func (c *ClientState) VerifyPacketAcknowledgement(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64, acknowledgement []byte) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyPacketAcknowledgementQuery = "verifypacketacknowledgement"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketAcknowledgementQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketAcknowledgementQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence
	inner["acknowledgement"] = acknowledgement
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying packet acknowledgement", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyPacketReceiptAbsence(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, sequence uint64) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyPacketReceiptAbsenceQuery = "verifypacketreceiptabsence"
	payload := make(map[string]map[string]interface{})
	payload[VerifyPacketReceiptAbsenceQuery] = make(map[string]interface{})
	inner := payload[VerifyPacketReceiptAbsenceQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["sequence"] = sequence
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verifying packet receipt absense", output.Result.ErrorMsg)
	}
}

func (c *ClientState) VerifyNextSequenceRecv(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, currentTimestamp uint64, delayPeriod uint64, prefix exported.Prefix, proof []byte, portID, channelID string, nextSequenceRecv uint64) error {
	consensusState, err := GetConsensusState(store, cdc, height)
	if err != nil {
		return err
	}

	const VerifyNextSequenceRecvQuery = "verifynextsequencerecv"
	payload := make(map[string]map[string]interface{})
	payload[VerifyNextSequenceRecvQuery] = make(map[string]interface{})
	inner := payload[VerifyNextSequenceRecvQuery]
	inner["me"] = c
	inner["height"] = height
	inner["commitment_prefix"] = prefix
	inner["proof"] = proof
	inner["port_id"] = portID
	inner["channel_id"] = channelID
	inner["current_timestamp"] = currentTimestamp
	inner["delay_period"] = delayPeriod
	inner["next_sequence_recv"] = nextSequenceRecv
	inner["consensus_state"] = consensusState

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToMarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}
	response, err := queryContractWithStore(c.CodeId, store, encodedData)
	if err != nil {
		return sdkerrors.Wrapf(ErrUnableToQuery, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	output := queryResponse{}
	if err := json.Unmarshal(response, &output); err != nil {
		return sdkerrors.Wrapf(ErrUnableToUnmarshalPayload, fmt.Sprintf("underlying error: %s", err.Error()))
	}

	if output.Result.IsValid {
		return nil
	} else {
		return fmt.Errorf("%s error while verify next sequence", output.Result.ErrorMsg)
	}
}
