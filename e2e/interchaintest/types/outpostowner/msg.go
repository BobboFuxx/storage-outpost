/* Code generated by github.com/srdtrk/go-codegen, DO NOT EDIT. */
package outpostowner

type InstantiateMsg struct {
	Admin *string `json:"admin,omitempty"`
	StorageOutpostCodeId int `json:"storage_outpost_code_id"`
}

type ExecuteMsg struct {
	CreateIcaContract *ExecuteMsg_CreateIcaContract `json:"create_ica_contract,omitempty"`
}

type QueryMsg struct {
	// GetContractState returns the contact's state.
	GetContractState *QueryMsg_GetContractState `json:"get_contract_state,omitempty"`
	// GetIcaState returns the ICA state for the given ICA ID.
	GetIcaContractState *QueryMsg_GetIcaContractState `json:"get_ica_contract_state,omitempty"`
	// GetIcaCount returns the number of ICAs.
	GetIcaCount *QueryMsg_GetIcaCount `json:"get_ica_count,omitempty"`
}

type ExecuteMsg_CreateIcaContract struct {
	ChannelOpenInitOptions ChannelOpenInitOptions `json:"channel_open_init_options"`
	Salt *string `json:"salt,omitempty"`
}

// Encoding is the encoding of the transactions sent to the ICA host.
type TxEncoding string

const (
	// Protobuf is the protobuf serialization of the CosmosSDK's Any.
	TxEncoding_Proto3 TxEncoding = "proto3"
	// Proto3Json is the json serialization of the CosmosSDK's Any.
	TxEncoding_Proto3Json TxEncoding = "proto3json"
)

type QueryMsg_GetIcaCount struct{}

/*
A human readable address.

In Cosmos, this is typically bech32 encoded. But for multi-chain smart contracts no assumptions should be made other than being UTF-8 encoded and of reasonable length.

This type represents a validated address. It can be created in the following ways 1. Use `Addr::unchecked(input)` 2. Use `let checked: Addr = deps.api.addr_validate(input)?` 3. Use `let checked: Addr = deps.api.addr_humanize(canonical_addr)?` 4. Deserialize from JSON. This must only be done from JSON that was validated before such as a contract's state. `Addr` must not be used in messages sent by the user because this would result in unvalidated instances.

This type is immutable. If you really need to mutate it (Really? Are you sure?), create a mutable copy using `let mut mutable = Addr::to_string()` and operate on that `String` instance.
*/
type Addr string

// ContractState is the state of the IBC application.
type ContractState struct {
	// The admin of this contract.
	Admin Addr `json:"admin"`
	// The code ID of the storage-outpost contract.
	StorageOutpostCodeId int `json:"storage_outpost_code_id"`
}

// IcaState is the state of the ICA.
type IcaState struct {
	ChannelState ChannelState `json:"channel_state"`
	IcaAddr string `json:"ica_addr"`
	IcaId int `json:"ica_id"`
	TxEncoding TxEncoding `json:"tx_encoding"`
}

type QueryMsg_GetIcaContractState struct {
	IcaId int `json:"ica_id"`
}

// IbcChannel defines all information on a channel. This is generally used in the hand-shake process, but can be queried directly.
type IbcChannel struct {
	// The connection upon which this channel was created. If this is a multi-hop channel, we only expose the first hop.
	ConnectionId string `json:"connection_id"`
	CounterpartyEndpoint IbcEndpoint `json:"counterparty_endpoint"`
	Endpoint IbcEndpoint `json:"endpoint"`
	Order IbcOrder `json:"order"`
	// Note: in ibcv3 this may be "", in the IbcOpenChannel handshake messages
	Version string `json:"version"`
}

// IbcOrder defines if a channel is ORDERED or UNORDERED Values come from https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/ibc/core/channel/v1/channel.proto#L69-L80 Naming comes from the protobuf files and go translations.
type IbcOrder string

const (
	IbcOrder_OrderUnordered IbcOrder = "ORDER_UNORDERED"
	IbcOrder_OrderOrdered   IbcOrder = "ORDER_ORDERED"
)

// IcaContractState is the state of the storage-outpost contract.
type IcaContractState struct {
	ContractAddr Addr `json:"contract_addr"`
	IcaState *IcaState `json:"ica_state,omitempty"`
}

// The message used to provide the MsgChannelOpenInit with the required data.
type ChannelOpenInitOptions struct {
	// The counterparty connection id on the counterparty chain.
	CounterpartyConnectionId string `json:"counterparty_connection_id"`
	// The counterparty port id. If not specified, [crate::ibc::types::keys::HOST_PORT_ID] is used. Currently, this contract only supports the host port.
	CounterpartyPortId *string `json:"counterparty_port_id,omitempty"`
	// TxEncoding is the encoding used for the ICA txs. If not specified, [TxEncoding::Protobuf] is used.
	TxEncoding *TxEncoding `json:"tx_encoding,omitempty"`
	// The connection id on this chain.
	ConnectionId string `json:"connection_id"`
}

type QueryMsg_GetContractState struct{}

// ContractChannelState is the state of the IBC application's channel. This application only supports one channel.
type ChannelState struct {
	// The status of the channel.
	ChannelStatus ChannelStatus `json:"channel_status"`
	// The IBC channel, as defined by cosmwasm.
	Channel IbcChannel `json:"channel"`
}

// ChannelState is the state of the IBC channel.
type ChannelStatus string

const (
	ChannelStatus_StateUninitializedUnspecified ChannelStatus = "STATE_UNINITIALIZED_UNSPECIFIED"
	ChannelStatus_StateInit                     ChannelStatus = "STATE_INIT"
	ChannelStatus_StateTryopen                  ChannelStatus = "STATE_TRYOPEN"
	ChannelStatus_StateOpen                     ChannelStatus = "STATE_OPEN"
	ChannelStatus_StateClosed                   ChannelStatus = "STATE_CLOSED"
)

type IbcEndpoint struct {
	PortId string `json:"port_id"`
	ChannelId string `json:"channel_id"`
}
