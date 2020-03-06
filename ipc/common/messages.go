package common

import (
	"fmt"
)

// MessageKind is
type MessageKind uint32

const (
	FirstKind MessageKind = iota
	Stop
	ContractDeployRequest
	ContractCallRequest
	ContractResponse
	BlockchainAccountExistsRequest
	BlockchainAccountExistsResponse
	BlockchainNewAddressRequest
	BlockchainNewAddressResponse
	BlockchainGetBalanceRequest
	BlockchainGetBalanceResponse
	BlockchainGetNonceRequest
	BlockchainGetNonceResponse
	BlockchainGetStorageDataRequest
	BlockchainGetStorageDataResponse
	BlockchainIsCodeEmptyRequest
	BlockchainIsCodeEmptyResponse
	BlockchainGetCodeRequest
	BlockchainGetCodeResponse
	BlockchainGetBlockhashRequest
	BlockchainGetBlockhashResponse
	BlockchainLastNonceRequest
	BlockchainLastNonceResponse
	BlockchainLastRoundRequest
	BlockchainLastRoundResponse
	BlockchainLastTimeStampRequest
	BlockchainLastTimeStampResponse
	BlockchainLastRandomSeedRequest
	BlockchainLastRandomSeedResponse
	BlockchainLastEpochRequest
	BlockchainLastEpochResponse
	BlockchainGetStateRootHashRequest
	BlockchainGetStateRootHashResponse
	BlockchainCurrentNonceRequest
	BlockchainCurrentNonceResponse
	BlockchainCurrentRoundRequest
	BlockchainCurrentRoundResponse
	BlockchainCurrentTimeStampRequest
	BlockchainCurrentTimeStampResponse
	BlockchainCurrentRandomSeedRequest
	BlockchainCurrentRandomSeedResponse
	BlockchainCurrentEpochRequest
	BlockchainCurrentEpochResponse
	DiagnoseWaitRequest
	DiagnoseWaitResponse
	LastKind
)

var messageKindNameByID = map[MessageKind]string{}

func init() {
	messageKindNameByID[FirstKind] = "FirstKind"
	messageKindNameByID[Stop] = "Stop"
	messageKindNameByID[ContractDeployRequest] = "ContractDeployRequest"
	messageKindNameByID[ContractCallRequest] = "ContractCallRequest"
	messageKindNameByID[ContractResponse] = "ContractResponse"
	messageKindNameByID[BlockchainAccountExistsRequest] = "BlockchainAccountExistsRequest"
	messageKindNameByID[BlockchainAccountExistsResponse] = "BlockchainAccountExistsResponse"
	messageKindNameByID[BlockchainNewAddressRequest] = "BlockchainNewAddressRequest"
	messageKindNameByID[BlockchainNewAddressResponse] = "BlockchainNewAddressResponse"
	messageKindNameByID[BlockchainGetBalanceRequest] = "BlockchainGetBalanceRequest"
	messageKindNameByID[BlockchainGetBalanceResponse] = "BlockchainGetBalanceResponse"
	messageKindNameByID[BlockchainGetNonceRequest] = "BlockchainGetNonceRequest"
	messageKindNameByID[BlockchainGetNonceResponse] = "BlockchainGetNonceResponse"
	messageKindNameByID[BlockchainGetStorageDataRequest] = "BlockchainGetStorageDataRequest"
	messageKindNameByID[BlockchainGetStorageDataResponse] = "BlockchainGetStorageDataResponse"
	messageKindNameByID[BlockchainIsCodeEmptyRequest] = "BlockchainIsCodeEmptyRequest"
	messageKindNameByID[BlockchainIsCodeEmptyResponse] = "BlockchainIsCodeEmptyResponse"
	messageKindNameByID[BlockchainGetCodeRequest] = "BlockchainGetCodeRequest"
	messageKindNameByID[BlockchainGetCodeResponse] = "BlockchainGetCodeResponse"
	messageKindNameByID[BlockchainGetBlockhashRequest] = "BlockchainGetBlockhashRequest"
	messageKindNameByID[BlockchainGetBlockhashResponse] = "BlockchainGetBlockhashResponse"
	messageKindNameByID[BlockchainLastNonceRequest] = "BlockchainLastNonceRequest"
	messageKindNameByID[BlockchainLastNonceResponse] = "BlockchainLastNonceResponse"
	messageKindNameByID[BlockchainLastRoundRequest] = "BlockchainLastRoundRequest"
	messageKindNameByID[BlockchainLastRoundResponse] = "BlockchainLastRoundResponse"
	messageKindNameByID[BlockchainLastTimeStampRequest] = "BlockchainLastTimeStampRequest"
	messageKindNameByID[BlockchainLastTimeStampResponse] = "BlockchainLastTimeStampResponse"
	messageKindNameByID[BlockchainLastRandomSeedRequest] = "BlockchainLastRandomSeedRequest"
	messageKindNameByID[BlockchainLastRandomSeedResponse] = "BlockchainLastRandomSeedResponse"
	messageKindNameByID[BlockchainLastEpochRequest] = "BlockchainLastEpochRequest"
	messageKindNameByID[BlockchainLastEpochResponse] = "BlockchainLastEpochResponse"
	messageKindNameByID[BlockchainGetStateRootHashRequest] = "BlockchainGetStateRootHashRequest"
	messageKindNameByID[BlockchainGetStateRootHashResponse] = "BlockchainGetStateRootHashResponse"
	messageKindNameByID[BlockchainCurrentNonceRequest] = "BlockchainCurrentNonceRequest"
	messageKindNameByID[BlockchainCurrentNonceResponse] = "BlockchainCurrentNonceResponse"
	messageKindNameByID[BlockchainCurrentRoundRequest] = "BlockchainCurrentRoundRequest"
	messageKindNameByID[BlockchainCurrentRoundResponse] = "BlockchainCurrentRoundResponse"
	messageKindNameByID[BlockchainCurrentTimeStampRequest] = "BlockchainCurrentTimeStampRequest"
	messageKindNameByID[BlockchainCurrentTimeStampResponse] = "BlockchainCurrentTimeStampResponse"
	messageKindNameByID[BlockchainCurrentRandomSeedRequest] = "BlockchainCurrentRandomSeedRequest"
	messageKindNameByID[BlockchainCurrentRandomSeedResponse] = "BlockchainCurrentRandomSeedResponse"
	messageKindNameByID[BlockchainCurrentEpochRequest] = "BlockchainCurrentEpochRequest"
	messageKindNameByID[BlockchainCurrentEpochResponse] = "BlockchainCurrentEpochResponse"
	messageKindNameByID[DiagnoseWaitRequest] = "DiagnoseWaitRequest"
	messageKindNameByID[DiagnoseWaitResponse] = "DiagnoseWaitResponse"
	messageKindNameByID[LastKind] = "LastKind"
}

// MessageHandler is
type MessageHandler interface {
	GetNonce() uint32
	SetNonce(nonce uint32)
	GetKind() MessageKind
	SetKind(kind MessageKind)
	GetError() error
	SetError(err error)
}

// Message is
type Message struct {
	DialogueNonce uint32
	Kind          MessageKind
	ErrorMessage  string
}

// GetNonce gets
func (message *Message) GetNonce() uint32 {
	return message.DialogueNonce
}

// SetNonce sets
func (message *Message) SetNonce(nonce uint32) {
	message.DialogueNonce = nonce
}

// GetKind gets
func (message *Message) GetKind() MessageKind {
	return message.Kind
}

// SetKind sets
func (message *Message) SetKind(kind MessageKind) {
	message.Kind = kind
}

// GetError gets
func (message *Message) GetError() error {
	if message.ErrorMessage == "" {
		return nil
	}

	return fmt.Errorf(message.ErrorMessage)
}

// SetError sets
func (message *Message) SetError(err error) {
	if err != nil {
		message.ErrorMessage = err.Error()
	}
}

func (message *Message) String() string {
	kindName, _ := messageKindNameByID[message.Kind]
	return fmt.Sprintf("[kind=%s nonce=%d err=%s]", kindName, message.DialogueNonce, message.ErrorMessage)
}

// MessageStop is
type MessageStop struct {
	Message
}

// NewMessageStop creates a message
func NewMessageStop() *MessageStop {
	message := &MessageStop{}
	message.Kind = Stop
	return message
}

// MessageReplier is a callback
type MessageReplier func(MessageHandler) MessageHandler

func noopReplier(message MessageHandler) MessageHandler {
	panic("NO-OP replier called")
}

// CreateReplySlots creates
func CreateReplySlots() []MessageReplier {
	slots := make([]MessageReplier, LastKind)
	for i := 0; i < len(slots); i++ {
		slots[i] = noopReplier
	}

	return slots
}

// IsHookCall returns
func IsHookCall(message MessageHandler) bool {
	kind := message.GetKind()
	return kind >= BlockchainAccountExistsRequest && kind <= BlockchainCurrentEpochResponse
}

// IsStopRequest returns
func IsStopRequest(message MessageHandler) bool {
	return message.GetKind() == Stop
}

// IsContractResponse returns
func IsContractResponse(message MessageHandler) bool {
	return message.GetKind() == ContractResponse
}

// IsDiagnose returns
func IsDiagnose(message MessageHandler) bool {
	kind := message.GetKind()
	return kind >= DiagnoseWaitRequest && kind <= DiagnoseWaitResponse
}
