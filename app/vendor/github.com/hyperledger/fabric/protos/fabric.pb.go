// Code generated by protoc-gen-go.
// source: fabric.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Transaction_Type int32

const (
	Transaction_UNDEFINED Transaction_Type = 0
	// deploy a chaincode to the network and call `Init` function
	Transaction_CHAINCODE_DEPLOY Transaction_Type = 1
	// call a chaincode `Invoke` function as a transaction
	Transaction_CHAINCODE_INVOKE Transaction_Type = 2
	// call a chaincode `query` function
	Transaction_CHAINCODE_QUERY Transaction_Type = 3
	// terminate a chaincode; not implemented yet
	Transaction_CHAINCODE_TERMINATE Transaction_Type = 4
)

var Transaction_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "CHAINCODE_DEPLOY",
	2: "CHAINCODE_INVOKE",
	3: "CHAINCODE_QUERY",
	4: "CHAINCODE_TERMINATE",
}
var Transaction_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"CHAINCODE_DEPLOY":    1,
	"CHAINCODE_INVOKE":    2,
	"CHAINCODE_QUERY":     3,
	"CHAINCODE_TERMINATE": 4,
}

func (x Transaction_Type) String() string {
	return proto.EnumName(Transaction_Type_name, int32(x))
}
func (Transaction_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 0} }

type PeerEndpoint_Type int32

const (
	PeerEndpoint_UNDEFINED     PeerEndpoint_Type = 0
	PeerEndpoint_VALIDATOR     PeerEndpoint_Type = 1
	PeerEndpoint_NON_VALIDATOR PeerEndpoint_Type = 2
)

var PeerEndpoint_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "VALIDATOR",
	2: "NON_VALIDATOR",
}
var PeerEndpoint_Type_value = map[string]int32{
	"UNDEFINED":     0,
	"VALIDATOR":     1,
	"NON_VALIDATOR": 2,
}

func (x PeerEndpoint_Type) String() string {
	return proto.EnumName(PeerEndpoint_Type_name, int32(x))
}
func (PeerEndpoint_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{8, 0} }

type Message_Type int32

const (
	Message_UNDEFINED               Message_Type = 0
	Message_DISC_HELLO              Message_Type = 1
	Message_DISC_DISCONNECT         Message_Type = 2
	Message_DISC_GET_PEERS          Message_Type = 3
	Message_DISC_PEERS              Message_Type = 4
	Message_DISC_NEWMSG             Message_Type = 5
	Message_CHAIN_TRANSACTION       Message_Type = 6
	Message_SYNC_GET_BLOCKS         Message_Type = 11
	Message_SYNC_BLOCKS             Message_Type = 12
	Message_SYNC_BLOCK_ADDED        Message_Type = 13
	Message_SYNC_STATE_GET_SNAPSHOT Message_Type = 14
	Message_SYNC_STATE_SNAPSHOT     Message_Type = 15
	Message_SYNC_STATE_GET_DELTAS   Message_Type = 16
	Message_SYNC_STATE_DELTAS       Message_Type = 17
	Message_RESPONSE                Message_Type = 20
	Message_CONSENSUS               Message_Type = 21
)

var Message_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "DISC_HELLO",
	2:  "DISC_DISCONNECT",
	3:  "DISC_GET_PEERS",
	4:  "DISC_PEERS",
	5:  "DISC_NEWMSG",
	6:  "CHAIN_TRANSACTION",
	11: "SYNC_GET_BLOCKS",
	12: "SYNC_BLOCKS",
	13: "SYNC_BLOCK_ADDED",
	14: "SYNC_STATE_GET_SNAPSHOT",
	15: "SYNC_STATE_SNAPSHOT",
	16: "SYNC_STATE_GET_DELTAS",
	17: "SYNC_STATE_DELTAS",
	20: "RESPONSE",
	21: "CONSENSUS",
}
var Message_Type_value = map[string]int32{
	"UNDEFINED":               0,
	"DISC_HELLO":              1,
	"DISC_DISCONNECT":         2,
	"DISC_GET_PEERS":          3,
	"DISC_PEERS":              4,
	"DISC_NEWMSG":             5,
	"CHAIN_TRANSACTION":       6,
	"SYNC_GET_BLOCKS":         11,
	"SYNC_BLOCKS":             12,
	"SYNC_BLOCK_ADDED":        13,
	"SYNC_STATE_GET_SNAPSHOT": 14,
	"SYNC_STATE_SNAPSHOT":     15,
	"SYNC_STATE_GET_DELTAS":   16,
	"SYNC_STATE_DELTAS":       17,
	"RESPONSE":                20,
	"CONSENSUS":               21,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (Message_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{12, 0} }

type Response_StatusCode int32

const (
	Response_UNDEFINED Response_StatusCode = 0
	Response_SUCCESS   Response_StatusCode = 200
	Response_FAILURE   Response_StatusCode = 500
)

var Response_StatusCode_name = map[int32]string{
	0:   "UNDEFINED",
	200: "SUCCESS",
	500: "FAILURE",
}
var Response_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"SUCCESS":   200,
	"FAILURE":   500,
}

func (x Response_StatusCode) String() string {
	return proto.EnumName(Response_StatusCode_name, int32(x))
}
func (Response_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{13, 0} }

// Transaction defines a function call to a contract.
// `args` is an array of type string so that the chaincode writer can choose
// whatever format they wish for the arguments for their chaincode.
// For example, they may wish to use JSON, XML, or a custom format.
// TODO: Defined remaining fields.
type Transaction struct {
	Type Transaction_Type `protobuf:"varint,1,opt,name=type,enum=protos.Transaction_Type" json:"type,omitempty"`
	// store ChaincodeID as bytes so its encrypted value can be stored
	ChaincodeID                    []byte                     `protobuf:"bytes,2,opt,name=chaincodeID,proto3" json:"chaincodeID,omitempty"`
	Payload                        []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata                       []byte                     `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Txid                           string                     `protobuf:"bytes,5,opt,name=txid" json:"txid,omitempty"`
	Timestamp                      *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp,omitempty"`
	ConfidentialityLevel           ConfidentialityLevel       `protobuf:"varint,7,opt,name=confidentialityLevel,enum=protos.ConfidentialityLevel" json:"confidentialityLevel,omitempty"`
	ConfidentialityProtocolVersion string                     `protobuf:"bytes,8,opt,name=confidentialityProtocolVersion" json:"confidentialityProtocolVersion,omitempty"`
	Nonce                          []byte                     `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ToValidators                   []byte                     `protobuf:"bytes,10,opt,name=toValidators,proto3" json:"toValidators,omitempty"`
	Cert                           []byte                     `protobuf:"bytes,11,opt,name=cert,proto3" json:"cert,omitempty"`
	Signature                      []byte                     `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Transaction) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// TransactionBlock carries a batch of transactions.
type TransactionBlock struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *TransactionBlock) Reset()                    { *m = TransactionBlock{} }
func (m *TransactionBlock) String() string            { return proto.CompactTextString(m) }
func (*TransactionBlock) ProtoMessage()               {}
func (*TransactionBlock) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *TransactionBlock) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// TransactionResult contains the return value of a transaction. It does
// not track potential state changes that were a result of the transaction.
// txid - The unique identifier of this transaction.
// result - The return value of the transaction.
// errorCode - An error code. 5xx will be logged as a failure in the dashboard.
// error - An error string for logging an issue.
// chaincodeEvent - any event emitted by a transaction
type TransactionResult struct {
	Txid           string          `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
	Result         []byte          `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	ErrorCode      uint32          `protobuf:"varint,3,opt,name=errorCode" json:"errorCode,omitempty"`
	Error          string          `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,5,opt,name=chaincodeEvent" json:"chaincodeEvent,omitempty"`
}

func (m *TransactionResult) Reset()                    { *m = TransactionResult{} }
func (m *TransactionResult) String() string            { return proto.CompactTextString(m) }
func (*TransactionResult) ProtoMessage()               {}
func (*TransactionResult) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *TransactionResult) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

// Block carries The data that describes a block in the blockchain.
// version - Version used to track any protocol changes.
// timestamp - The time at which the block or transaction order
// was proposed. This may not be used by all consensus modules.
// transactions - The ordered list of transactions in the block.
// stateHash - The state hash after running transactions in this block.
// previousBlockHash - The hash of the previous block in the chain.
// consensusMetadata - Consensus modules may optionally store any
// additional metadata in this field.
// nonHashData - Data stored with the block, but not included in the blocks
// hash. This allows this data to be different per peer or discarded without
// impacting the blockchain.
type Block struct {
	Version           uint32                     `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp         *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Transactions      []*Transaction             `protobuf:"bytes,3,rep,name=transactions" json:"transactions,omitempty"`
	StateHash         []byte                     `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	PreviousBlockHash []byte                     `protobuf:"bytes,5,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	ConsensusMetadata []byte                     `protobuf:"bytes,6,opt,name=consensusMetadata,proto3" json:"consensusMetadata,omitempty"`
	NonHashData       *NonHashData               `protobuf:"bytes,7,opt,name=nonHashData" json:"nonHashData,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *Block) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetNonHashData() *NonHashData {
	if m != nil {
		return m.NonHashData
	}
	return nil
}

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
}

func (m *BlockchainInfo) Reset()                    { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string            { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()               {}
func (*BlockchainInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

// NonHashData is data that is recorded on the block, but not included in
// the block hash when verifying the blockchain.
// localLedgerCommitTimestamp - The time at which the block was added
// to the ledger on the local peer.
// chaincodeEvent - is an array ChaincodeEvents, one per transaction in the
// block
type NonHashData struct {
	LocalLedgerCommitTimestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=localLedgerCommitTimestamp" json:"localLedgerCommitTimestamp,omitempty"`
	ChaincodeEvents            []*ChaincodeEvent          `protobuf:"bytes,2,rep,name=chaincodeEvents" json:"chaincodeEvents,omitempty"`
}

func (m *NonHashData) Reset()                    { *m = NonHashData{} }
func (m *NonHashData) String() string            { return proto.CompactTextString(m) }
func (*NonHashData) ProtoMessage()               {}
func (*NonHashData) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *NonHashData) GetLocalLedgerCommitTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.LocalLedgerCommitTimestamp
	}
	return nil
}

func (m *NonHashData) GetChaincodeEvents() []*ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvents
	}
	return nil
}

type PeerAddress struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *PeerAddress) Reset()                    { *m = PeerAddress{} }
func (m *PeerAddress) String() string            { return proto.CompactTextString(m) }
func (*PeerAddress) ProtoMessage()               {}
func (*PeerAddress) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

type PeerID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

type PeerEndpoint struct {
	ID      *PeerID           `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Address string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type    PeerEndpoint_Type `protobuf:"varint,3,opt,name=type,enum=protos.PeerEndpoint_Type" json:"type,omitempty"`
	PkiID   []byte            `protobuf:"bytes,4,opt,name=pkiID,proto3" json:"pkiID,omitempty"`
}

func (m *PeerEndpoint) Reset()                    { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()               {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *PeerEndpoint) GetID() *PeerID {
	if m != nil {
		return m.ID
	}
	return nil
}

type PeersMessage struct {
	Peers []*PeerEndpoint `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *PeersMessage) Reset()                    { *m = PeersMessage{} }
func (m *PeersMessage) String() string            { return proto.CompactTextString(m) }
func (*PeersMessage) ProtoMessage()               {}
func (*PeersMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *PeersMessage) GetPeers() []*PeerEndpoint {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeersAddresses struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *PeersAddresses) Reset()                    { *m = PeersAddresses{} }
func (m *PeersAddresses) String() string            { return proto.CompactTextString(m) }
func (*PeersAddresses) ProtoMessage()               {}
func (*PeersAddresses) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

type HelloMessage struct {
	PeerEndpoint   *PeerEndpoint   `protobuf:"bytes,1,opt,name=peerEndpoint" json:"peerEndpoint,omitempty"`
	BlockchainInfo *BlockchainInfo `protobuf:"bytes,2,opt,name=blockchainInfo" json:"blockchainInfo,omitempty"`
}

func (m *HelloMessage) Reset()                    { *m = HelloMessage{} }
func (m *HelloMessage) String() string            { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()               {}
func (*HelloMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *HelloMessage) GetPeerEndpoint() *PeerEndpoint {
	if m != nil {
		return m.PeerEndpoint
	}
	return nil
}

func (m *HelloMessage) GetBlockchainInfo() *BlockchainInfo {
	if m != nil {
		return m.BlockchainInfo
	}
	return nil
}

type Message struct {
	Type      Message_Type               `protobuf:"varint,1,opt,name=type,enum=protos.Message_Type" json:"type,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *Message) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Response struct {
	Status Response_StatusCode `protobuf:"varint,1,opt,name=status,enum=protos.Response_StatusCode" json:"status,omitempty"`
	Msg    []byte              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

// BlockState is the payload of Message.SYNC_BLOCK_ADDED. When a VP
// commits a new block to the ledger, it will notify its connected NVPs of the
// block and the delta state. The NVP may call the ledger APIs to apply the
// block and the delta state to its ledger if the block's previousBlockHash
// equals to the NVP's current block hash
type BlockState struct {
	Block      *Block `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	StateDelta []byte `protobuf:"bytes,2,opt,name=stateDelta,proto3" json:"stateDelta,omitempty"`
}

func (m *BlockState) Reset()                    { *m = BlockState{} }
func (m *BlockState) String() string            { return proto.CompactTextString(m) }
func (*BlockState) ProtoMessage()               {}
func (*BlockState) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{14} }

func (m *BlockState) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// SyncBlockRange is the payload of Message.SYNC_GET_BLOCKS, where
// start and end indicate the starting and ending blocks inclusively. The order
// in which blocks are returned is defined by the start and end values. For
// example, if start=3 and end=5, the order of blocks will be 3, 4, 5.
// If start=5 and end=3, the order will be 5, 4, 3.
type SyncBlockRange struct {
	CorrelationId uint64 `protobuf:"varint,1,opt,name=correlationId" json:"correlationId,omitempty"`
	Start         uint64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	End           uint64 `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
}

func (m *SyncBlockRange) Reset()                    { *m = SyncBlockRange{} }
func (m *SyncBlockRange) String() string            { return proto.CompactTextString(m) }
func (*SyncBlockRange) ProtoMessage()               {}
func (*SyncBlockRange) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{15} }

// SyncBlocks is the payload of Message.SYNC_BLOCKS, where the range
// indicates the blocks responded to the request SYNC_GET_BLOCKS
type SyncBlocks struct {
	Range  *SyncBlockRange `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Blocks []*Block        `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *SyncBlocks) Reset()                    { *m = SyncBlocks{} }
func (m *SyncBlocks) String() string            { return proto.CompactTextString(m) }
func (*SyncBlocks) ProtoMessage()               {}
func (*SyncBlocks) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{16} }

func (m *SyncBlocks) GetRange() *SyncBlockRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *SyncBlocks) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

// SyncSnapshotRequest Payload for the penchainMessage.SYNC_GET_SNAPSHOT message.
type SyncStateSnapshotRequest struct {
	CorrelationId uint64 `protobuf:"varint,1,opt,name=correlationId" json:"correlationId,omitempty"`
}

func (m *SyncStateSnapshotRequest) Reset()                    { *m = SyncStateSnapshotRequest{} }
func (m *SyncStateSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncStateSnapshotRequest) ProtoMessage()               {}
func (*SyncStateSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{17} }

// SyncStateSnapshot is the payload of Message.SYNC_SNAPSHOT, which is a response
// to penchainMessage.SYNC_GET_SNAPSHOT. It contains the snapshot or a chunk of the
// snapshot on stream, and in which case, the sequence indicate the order
// starting at 0.  The terminating message will have len(delta) == 0.
type SyncStateSnapshot struct {
	Delta       []byte                    `protobuf:"bytes,1,opt,name=delta,proto3" json:"delta,omitempty"`
	Sequence    uint64                    `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	BlockNumber uint64                    `protobuf:"varint,3,opt,name=blockNumber" json:"blockNumber,omitempty"`
	Request     *SyncStateSnapshotRequest `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
}

func (m *SyncStateSnapshot) Reset()                    { *m = SyncStateSnapshot{} }
func (m *SyncStateSnapshot) String() string            { return proto.CompactTextString(m) }
func (*SyncStateSnapshot) ProtoMessage()               {}
func (*SyncStateSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{18} }

func (m *SyncStateSnapshot) GetRequest() *SyncStateSnapshotRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// SyncStateDeltasRequest is the payload of Message.SYNC_GET_STATE.
// blockNumber indicates the block number for the delta which is being
// requested. If no payload is included with SYNC_GET_STATE, it represents
// a request for a snapshot of the current state.
type SyncStateDeltasRequest struct {
	Range *SyncBlockRange `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
}

func (m *SyncStateDeltasRequest) Reset()                    { *m = SyncStateDeltasRequest{} }
func (m *SyncStateDeltasRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncStateDeltasRequest) ProtoMessage()               {}
func (*SyncStateDeltasRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{19} }

func (m *SyncStateDeltasRequest) GetRange() *SyncBlockRange {
	if m != nil {
		return m.Range
	}
	return nil
}

// SyncStateDeltas is the payload of the Message.SYNC_STATE in response to
// the Message.SYNC_GET_STATE message.
type SyncStateDeltas struct {
	Range  *SyncBlockRange `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Deltas [][]byte        `protobuf:"bytes,2,rep,name=deltas,proto3" json:"deltas,omitempty"`
}

func (m *SyncStateDeltas) Reset()                    { *m = SyncStateDeltas{} }
func (m *SyncStateDeltas) String() string            { return proto.CompactTextString(m) }
func (*SyncStateDeltas) ProtoMessage()               {}
func (*SyncStateDeltas) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{20} }

func (m *SyncStateDeltas) GetRange() *SyncBlockRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*TransactionBlock)(nil), "protos.TransactionBlock")
	proto.RegisterType((*TransactionResult)(nil), "protos.TransactionResult")
	proto.RegisterType((*Block)(nil), "protos.Block")
	proto.RegisterType((*BlockchainInfo)(nil), "protos.BlockchainInfo")
	proto.RegisterType((*NonHashData)(nil), "protos.NonHashData")
	proto.RegisterType((*PeerAddress)(nil), "protos.PeerAddress")
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
	proto.RegisterType((*PeersMessage)(nil), "protos.PeersMessage")
	proto.RegisterType((*PeersAddresses)(nil), "protos.PeersAddresses")
	proto.RegisterType((*HelloMessage)(nil), "protos.HelloMessage")
	proto.RegisterType((*Message)(nil), "protos.Message")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*BlockState)(nil), "protos.BlockState")
	proto.RegisterType((*SyncBlockRange)(nil), "protos.SyncBlockRange")
	proto.RegisterType((*SyncBlocks)(nil), "protos.SyncBlocks")
	proto.RegisterType((*SyncStateSnapshotRequest)(nil), "protos.SyncStateSnapshotRequest")
	proto.RegisterType((*SyncStateSnapshot)(nil), "protos.SyncStateSnapshot")
	proto.RegisterType((*SyncStateDeltasRequest)(nil), "protos.SyncStateDeltasRequest")
	proto.RegisterType((*SyncStateDeltas)(nil), "protos.SyncStateDeltas")
	proto.RegisterEnum("protos.Transaction_Type", Transaction_Type_name, Transaction_Type_value)
	proto.RegisterEnum("protos.PeerEndpoint_Type", PeerEndpoint_Type_name, PeerEndpoint_Type_value)
	proto.RegisterEnum("protos.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterEnum("protos.Response_StatusCode", Response_StatusCode_name, Response_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Peer service

type PeerClient interface {
	// Accepts a stream of Message during chat session, while receiving
	// other Message (e.g. from other peers).
	Chat(ctx context.Context, opts ...grpc.CallOption) (Peer_ChatClient, error)
	// Process a transaction from a remote source.
	ProcessTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Response, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Peer_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[0], c.cc, "/protos.Peer/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerChatClient{stream}
	return x, nil
}

type Peer_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type peerChatClient struct {
	grpc.ClientStream
}

func (x *peerChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) ProcessTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Peer/ProcessTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Peer service

type PeerServer interface {
	// Accepts a stream of Message during chat session, while receiving
	// other Message (e.g. from other peers).
	Chat(Peer_ChatServer) error
	// Process a transaction from a remote source.
	ProcessTransaction(context.Context, *Transaction) (*Response, error)
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServer).Chat(&peerChatServer{stream})
}

type Peer_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type peerChatServer struct {
	grpc.ServerStream
}

func (x *peerChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Peer_ProcessTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).ProcessTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Peer/ProcessTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).ProcessTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessTransaction",
			Handler:    _Peer_ProcessTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Peer_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("fabric.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0xf5, 0x67, 0x6b, 0x24, 0xcb, 0xf4, 0xc6, 0x71, 0x18, 0x27, 0x48, 0x0d, 0xb6, 0x05,
	0x8c, 0x20, 0x55, 0x0a, 0x07, 0x41, 0x82, 0x00, 0x2d, 0xa2, 0x88, 0x74, 0x2c, 0x44, 0xa6, 0x94,
	0xa5, 0xec, 0x20, 0x3d, 0xd4, 0xa0, 0xa9, 0xb5, 0x4c, 0x44, 0xe2, 0xaa, 0xdc, 0x95, 0x51, 0x5f,
	0x7b, 0x2a, 0xfa, 0x1a, 0x3d, 0xf6, 0xd8, 0x67, 0xe8, 0xcf, 0x23, 0xf4, 0x2d, 0x7a, 0xe9, 0x03,
	0x14, 0xbb, 0x4b, 0x52, 0xa4, 0xac, 0xa4, 0x49, 0x2f, 0xf6, 0xce, 0x37, 0xdf, 0xec, 0xce, 0xcc,
	0xce, 0xce, 0x50, 0x50, 0x3f, 0xf3, 0x4e, 0xa3, 0xc0, 0x6f, 0x4e, 0x23, 0xca, 0x29, 0xaa, 0xc8,
	0x7f, 0x6c, 0x7b, 0xdd, 0x3f, 0xf7, 0x82, 0xd0, 0xa7, 0x43, 0xa2, 0x14, 0xdb, 0x9b, 0x29, 0x40,
	0x2e, 0x48, 0xc8, 0x63, 0xf4, 0x93, 0x11, 0xa5, 0xa3, 0x31, 0x79, 0x20, 0xa5, 0xd3, 0xd9, 0xd9,
	0x03, 0x1e, 0x4c, 0x08, 0xe3, 0xde, 0x64, 0xaa, 0x08, 0xe6, 0x5f, 0x25, 0xa8, 0x0d, 0x22, 0x2f,
	0x64, 0x9e, 0xcf, 0x03, 0x1a, 0xa2, 0xfb, 0x50, 0xe2, 0x97, 0x53, 0x62, 0x68, 0x3b, 0xda, 0x6e,
	0x63, 0xcf, 0x50, 0x2c, 0xd6, 0xcc, 0x50, 0x9a, 0x83, 0xcb, 0x29, 0xc1, 0x92, 0x85, 0x76, 0xa0,
	0x96, 0x1e, 0xdb, 0xb1, 0x8c, 0xc2, 0x8e, 0xb6, 0x5b, 0xc7, 0x59, 0x08, 0x19, 0xb0, 0x32, 0xf5,
	0x2e, 0xc7, 0xd4, 0x1b, 0x1a, 0x45, 0xa9, 0x4d, 0x44, 0xb4, 0x0d, 0xab, 0x13, 0xc2, 0xbd, 0xa1,
	0xc7, 0x3d, 0xa3, 0x24, 0x55, 0xa9, 0x8c, 0x10, 0x94, 0xf8, 0xf7, 0xc1, 0xd0, 0x28, 0xef, 0x68,
	0xbb, 0x55, 0x2c, 0xd7, 0xe8, 0x09, 0x54, 0x53, 0xe7, 0x8d, 0xca, 0x8e, 0xb6, 0x5b, 0xdb, 0xdb,
	0x6e, 0xaa, 0xf0, 0x9a, 0x49, 0x78, 0xcd, 0x41, 0xc2, 0xc0, 0x73, 0x32, 0xea, 0xc3, 0xa6, 0x4f,
	0xc3, 0xb3, 0x60, 0x48, 0x42, 0x1e, 0x78, 0xe3, 0x80, 0x5f, 0x76, 0xc9, 0x05, 0x19, 0x1b, 0x2b,
	0x32, 0xc6, 0x3b, 0x49, 0x8c, 0xed, 0x25, 0x1c, 0xbc, 0xd4, 0x12, 0xed, 0xc3, 0xdd, 0x05, 0xbc,
	0x2f, 0xf6, 0xf0, 0xe9, 0xf8, 0x98, 0x44, 0x2c, 0xa0, 0xa1, 0xb1, 0x2a, 0x3d, 0xff, 0x0f, 0x16,
	0xda, 0x84, 0x72, 0x48, 0x43, 0x9f, 0x18, 0x55, 0x99, 0x00, 0x25, 0x20, 0x13, 0xea, 0x9c, 0x1e,
	0x7b, 0xe3, 0x60, 0xe8, 0x71, 0x1a, 0x31, 0x03, 0xa4, 0x32, 0x87, 0x89, 0x0c, 0xf9, 0x24, 0xe2,
	0x46, 0x4d, 0xea, 0xe4, 0x1a, 0xdd, 0x81, 0x2a, 0x0b, 0x46, 0xa1, 0xc7, 0x67, 0x11, 0x31, 0xea,
	0x52, 0x31, 0x07, 0x4c, 0x0a, 0x25, 0x71, 0x73, 0x68, 0x0d, 0xaa, 0x47, 0x8e, 0x65, 0xef, 0x77,
	0x1c, 0xdb, 0xd2, 0xaf, 0xa1, 0x4d, 0xd0, 0xdb, 0x07, 0xad, 0x8e, 0xd3, 0xee, 0x59, 0xf6, 0x89,
	0x65, 0xf7, 0xbb, 0xbd, 0x37, 0xba, 0x96, 0x47, 0x3b, 0xce, 0x71, 0xef, 0xa5, 0xad, 0x17, 0xd0,
	0x75, 0x58, 0x9f, 0xa3, 0xaf, 0x8e, 0x6c, 0xfc, 0x46, 0x2f, 0xa2, 0x9b, 0x70, 0x7d, 0x0e, 0x0e,
	0x6c, 0x7c, 0xd8, 0x71, 0x5a, 0x03, 0x5b, 0x2f, 0x99, 0x2f, 0x41, 0xcf, 0x94, 0xcd, 0xf3, 0x31,
	0xf5, 0xdf, 0xa2, 0xc7, 0x50, 0xe7, 0x73, 0x8c, 0x19, 0xda, 0x4e, 0x71, 0xb7, 0xb6, 0x77, 0x7d,
	0x49, 0x99, 0xe1, 0x1c, 0xd1, 0xfc, 0x55, 0x83, 0x8d, 0xac, 0x96, 0xb0, 0xd9, 0x98, 0xa7, 0x75,
	0xa2, 0x65, 0xea, 0x64, 0x0b, 0x2a, 0x91, 0xd4, 0xc6, 0xe5, 0x18, 0x4b, 0x22, 0x3b, 0x24, 0x8a,
	0x68, 0xd4, 0xa6, 0x43, 0x22, 0x6b, 0x71, 0x0d, 0xcf, 0x01, 0x71, 0x13, 0x52, 0x90, 0xa5, 0x58,
	0xc5, 0x4a, 0x40, 0x5f, 0x43, 0x23, 0x2d, 0x66, 0x5b, 0x3c, 0x2b, 0x59, 0x91, 0xb5, 0xbd, 0xad,
	0xb4, 0x66, 0x72, 0x5a, 0xbc, 0xc0, 0x36, 0x7f, 0x2b, 0x40, 0x59, 0x05, 0x6e, 0xc0, 0xca, 0x45,
	0x5c, 0x1a, 0x9a, 0x3c, 0x3b, 0x11, 0xf3, 0x75, 0x5d, 0xf8, 0x98, 0xba, 0x5e, 0x4c, 0x66, 0xf1,
	0x03, 0x93, 0x29, 0x0b, 0x85, 0x7b, 0x9c, 0x1c, 0x78, 0xec, 0x3c, 0x7e, 0x7b, 0x73, 0x00, 0xdd,
	0x87, 0x8d, 0x69, 0x44, 0x2e, 0x02, 0x3a, 0x63, 0xd2, 0x77, 0xc9, 0x2a, 0x4b, 0xd6, 0x55, 0x85,
	0x60, 0xfb, 0x34, 0x64, 0x24, 0x64, 0x33, 0x76, 0x98, 0xbc, 0xe7, 0x8a, 0x62, 0x5f, 0x51, 0xa0,
	0x47, 0x50, 0x0b, 0x69, 0x28, 0x0c, 0x2d, 0xc1, 0x5b, 0x91, 0xe1, 0xa6, 0x1e, 0x3b, 0x73, 0x15,
	0xce, 0xf2, 0xcc, 0x1f, 0x34, 0x68, 0xc8, 0x23, 0x65, 0x7e, 0x3b, 0xe1, 0x19, 0x15, 0xd7, 0x7c,
	0x4e, 0x82, 0xd1, 0x39, 0x97, 0xf9, 0x2c, 0xe1, 0x58, 0x42, 0xf7, 0x40, 0xf7, 0x67, 0x51, 0x44,
	0x42, 0x3e, 0x77, 0x5e, 0x15, 0xc2, 0x15, 0x7c, 0x79, 0xa4, 0xc5, 0x77, 0x44, 0x6a, 0xfe, 0xa2,
	0x41, 0x2d, 0xe3, 0x21, 0xfa, 0x06, 0xb6, 0xc7, 0xd4, 0xf7, 0xc6, 0x5d, 0x32, 0x1c, 0x91, 0xa8,
	0x4d, 0x27, 0x93, 0x80, 0xa7, 0xf7, 0x24, 0xbd, 0x7a, 0xff, 0x4d, 0xbe, 0xc7, 0x1a, 0x3d, 0x83,
	0xf5, 0x7c, 0x29, 0x31, 0xa3, 0x20, 0x6f, 0xf7, 0x5d, 0x95, 0xb7, 0x48, 0x37, 0x1f, 0x41, 0xad,
	0x4f, 0x48, 0xd4, 0x1a, 0x0e, 0x23, 0xc2, 0x64, 0xbf, 0x38, 0xa7, 0x8c, 0x27, 0x2f, 0x45, 0xac,
	0x05, 0x36, 0xa5, 0x91, 0x7a, 0x27, 0x65, 0x2c, 0xd7, 0xe6, 0x1d, 0xa8, 0x08, 0xb3, 0x8e, 0x25,
	0xb4, 0xa1, 0x37, 0x21, 0x89, 0x85, 0x58, 0x9b, 0xbf, 0x6b, 0x50, 0x17, 0x6a, 0x3b, 0x1c, 0x4e,
	0x69, 0x10, 0x72, 0x74, 0x17, 0x0a, 0x1d, 0x2b, 0x8e, 0xb5, 0x91, 0xb8, 0xa6, 0x36, 0xc0, 0x05,
	0xd5, 0xfe, 0x3d, 0xe5, 0x81, 0x3c, 0xa5, 0x8a, 0x13, 0x11, 0x7d, 0x11, 0x0f, 0x9a, 0xa2, 0x6c,
	0xc2, 0xb7, 0xb2, 0xb6, 0xc9, 0xee, 0xd9, 0x49, 0xb3, 0x09, 0xe5, 0xe9, 0xdb, 0xa0, 0x63, 0xc5,
	0xe5, 0xaa, 0x04, 0xf3, 0xf1, 0xf2, 0x9e, 0xb6, 0x06, 0xd5, 0xe3, 0x56, 0xb7, 0x63, 0xb5, 0x06,
	0x3d, 0xac, 0x6b, 0x68, 0x03, 0xd6, 0x9c, 0x9e, 0x73, 0x32, 0x87, 0x0a, 0xe6, 0x53, 0x15, 0x07,
	0x3b, 0x24, 0x8c, 0x79, 0x23, 0x82, 0xee, 0x41, 0x79, 0x2a, 0xe4, 0xb8, 0x21, 0x6d, 0x2e, 0x73,
	0x07, 0x2b, 0x8a, 0xd9, 0x84, 0x86, 0xb4, 0x8d, 0x53, 0x4b, 0xe4, 0x7b, 0xf2, 0x12, 0x41, 0xee,
	0x50, 0xc5, 0x73, 0xc0, 0xfc, 0x51, 0x83, 0xfa, 0x01, 0x19, 0x8f, 0x69, 0x72, 0xd8, 0x13, 0xa8,
	0x4f, 0x33, 0xfb, 0xc6, 0xe9, 0x5b, 0x7e, 0x66, 0x8e, 0x29, 0xfa, 0xd1, 0x69, 0xee, 0x19, 0xc4,
	0x0d, 0x23, 0xad, 0x8a, 0xfc, 0x23, 0xc1, 0x0b, 0x6c, 0xf3, 0xef, 0x22, 0xac, 0x24, 0x5e, 0xec,
	0xe6, 0x26, 0x7d, 0x7a, 0x7a, 0xac, 0xce, 0xe6, 0xfe, 0xff, 0x77, 0xa8, 0x77, 0x4f, 0xff, 0xdc,
	0xac, 0x2a, 0x2d, 0xce, 0xaa, 0x3f, 0x0a, 0xcb, 0x2f, 0xb6, 0x01, 0x60, 0x75, 0xdc, 0xf6, 0xc9,
	0x81, 0xdd, 0xed, 0xf6, 0x74, 0x4d, 0x0c, 0x24, 0x29, 0x8b, 0x3f, 0x3d, 0xc7, 0xb1, 0xdb, 0x03,
	0xbd, 0x80, 0x10, 0x34, 0x24, 0xf8, 0xc2, 0x1e, 0x9c, 0xf4, 0x6d, 0x1b, 0xbb, 0x7a, 0x31, 0x35,
	0x54, 0x72, 0x09, 0xad, 0x43, 0x4d, 0xca, 0x8e, 0xfd, 0xfa, 0xd0, 0x7d, 0xa1, 0x97, 0xd1, 0x0d,
	0xd8, 0x90, 0x53, 0xec, 0x64, 0x80, 0x5b, 0x8e, 0xdb, 0x6a, 0x0f, 0x3a, 0x3d, 0x47, 0xaf, 0x88,
	0x03, 0xdc, 0x37, 0x8e, 0xda, 0xeb, 0x79, 0xb7, 0xd7, 0x7e, 0xe9, 0xea, 0x35, 0x61, 0x2c, 0xc1,
	0x18, 0xa8, 0x8b, 0x69, 0x39, 0x07, 0x4e, 0x5a, 0x96, 0x65, 0x5b, 0xfa, 0x1a, 0xba, 0x0d, 0x37,
	0x25, 0xea, 0x0e, 0x5a, 0x03, 0x5b, 0xee, 0xe0, 0x3a, 0xad, 0xbe, 0x7b, 0xd0, 0x1b, 0xe8, 0x0d,
	0x31, 0x35, 0x33, 0xca, 0x54, 0xb1, 0x8e, 0x6e, 0xc1, 0x8d, 0x05, 0x2b, 0xcb, 0xee, 0x0e, 0x5a,
	0xae, 0xae, 0x0b, 0x1f, 0x33, 0xaa, 0x18, 0xde, 0x40, 0x75, 0x58, 0xc5, 0xb6, 0xdb, 0xef, 0x39,
	0xae, 0xad, 0x6f, 0x8a, 0x8c, 0xb5, 0xc5, 0xd2, 0x71, 0x8f, 0x5c, 0xfd, 0x86, 0xf9, 0x93, 0x06,
	0xab, 0x98, 0xb0, 0xa9, 0xe8, 0xc4, 0xe8, 0x21, 0x54, 0x44, 0x9b, 0x9f, 0xb1, 0xf8, 0xd2, 0x6f,
	0x27, 0x97, 0x9e, 0x30, 0x9a, 0xae, 0x54, 0x8b, 0x89, 0x88, 0x63, 0x2a, 0xd2, 0xa1, 0x38, 0x61,
	0xa3, 0xb8, 0x87, 0x8a, 0xa5, 0xf9, 0x18, 0x60, 0xce, 0x5b, 0xbc, 0xa2, 0x3a, 0xac, 0xb8, 0x47,
	0xed, 0xb6, 0xed, 0xba, 0xfa, 0x9f, 0x9a, 0x90, 0xf6, 0x5b, 0x9d, 0xee, 0x11, 0xb6, 0xf5, 0x7f,
	0x8a, 0xe6, 0x2b, 0x00, 0x59, 0xa0, 0xc2, 0x9a, 0xa0, 0x4f, 0xa1, 0x2c, 0xcb, 0x33, 0xae, 0xff,
	0xb5, 0x5c, 0x0d, 0x63, 0xa5, 0x43, 0x77, 0x01, 0xe4, 0x64, 0xb2, 0xc8, 0x98, 0x7b, 0xb1, 0x13,
	0x19, 0xc4, 0xfc, 0x16, 0x1a, 0xee, 0x65, 0xe8, 0x2b, 0x1b, 0x2f, 0x1c, 0x11, 0xf4, 0x19, 0xac,
	0xf9, 0x34, 0x8a, 0xc8, 0xd8, 0x13, 0xc3, 0xae, 0x33, 0x8c, 0xe7, 0x43, 0x1e, 0x14, 0xfd, 0x84,
	0x71, 0x2f, 0x6e, 0x7e, 0x25, 0xac, 0x04, 0x11, 0x2b, 0x09, 0x55, 0xad, 0x96, 0xb0, 0x58, 0x9a,
	0x1e, 0x40, 0xba, 0x3f, 0x43, 0xf7, 0xa1, 0x1c, 0x89, 0x43, 0x62, 0x97, 0xd3, 0x67, 0x97, 0x77,
	0x01, 0x2b, 0x12, 0xfa, 0x1c, 0x2a, 0x32, 0x88, 0xa4, 0x77, 0x2f, 0x44, 0x18, 0x2b, 0xcd, 0x67,
	0x60, 0x08, 0x7b, 0x99, 0x14, 0x37, 0xf4, 0xa6, 0xec, 0x9c, 0x72, 0x4c, 0xbe, 0x9b, 0x11, 0xc6,
	0x3f, 0x2c, 0x18, 0xf3, 0x67, 0x0d, 0x36, 0xae, 0x6c, 0x21, 0x42, 0x1c, 0xca, 0xac, 0x69, 0xaa,
	0x65, 0x4a, 0x41, 0x7c, 0x76, 0x33, 0xb1, 0xb9, 0xf8, 0xea, 0x54, 0xb1, 0xa7, 0xb2, 0xf8, 0x9c,
	0x97, 0x3e, 0x39, 0xb3, 0xc9, 0x29, 0x89, 0xe2, 0x34, 0x64, 0x21, 0xf4, 0x14, 0x56, 0x22, 0xe5,
	0x9a, 0x7c, 0xb4, 0xb5, 0xbd, 0x9d, 0x6c, 0x0a, 0x96, 0x85, 0x80, 0x13, 0x03, 0x73, 0x1f, 0xb6,
	0x52, 0x92, 0xbc, 0x3c, 0x96, 0x44, 0xf9, 0x51, 0x69, 0x35, 0x5f, 0xc3, 0xfa, 0xc2, 0x3e, 0x1f,
	0x79, 0x2f, 0x5b, 0x50, 0x91, 0xb9, 0x50, 0xf7, 0x52, 0xc7, 0xb1, 0xb4, 0x37, 0x83, 0x92, 0xe8,
	0xbd, 0xa8, 0x09, 0xa5, 0xf6, 0xb9, 0xc7, 0xd1, 0xfa, 0x42, 0x4f, 0xdc, 0x5e, 0x04, 0xcc, 0x6b,
	0xbb, 0xda, 0x97, 0x1a, 0xfa, 0x0a, 0x50, 0x3f, 0xa2, 0x3e, 0x61, 0x2c, 0xfb, 0x4b, 0x6a, 0xd9,
	0x77, 0xd8, 0xb6, 0xbe, 0xf8, 0xe2, 0xcc, 0x6b, 0xa7, 0xea, 0x27, 0xdd, 0xc3, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x71, 0x9c, 0x63, 0xe9, 0x0d, 0x00, 0x00,
}
