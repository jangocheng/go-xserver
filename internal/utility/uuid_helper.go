package utility

import (
	"bytes"
	"encoding/binary"

	"github.com/fananchong/go-xserver/common"
	"github.com/fananchong/go-xserver/internal/protocol"
	uuid "github.com/satori/go.uuid"
)

// NewNID : 生成一个NID
func NewNID() common.NodeID {
	return UUID2NodeID(uuid.NewV2(uuid.DomainPerson))
}

// UUID2NodeID : uuid.UUID 转化为 common.NodeID
func UUID2NodeID(uid uuid.UUID) common.NodeID {
	nid := common.NodeID{}
	copy(nid[:], uid[:])
	return nid
}

// NodeID2UUID : common.NodeID 转化为 uuid.UUID
func NodeID2UUID(nid common.NodeID) uuid.UUID {
	uid := uuid.UUID{}
	copy(uid[:], nid[:])
	return uid
}

// NodeID2ServerID : common.NodeID 转化为 protocol.SERVER_ID
func NodeID2ServerID(nid common.NodeID) *protocol.SERVER_ID {
	sid := &protocol.SERVER_ID{}
	sid.Low = binary.LittleEndian.Uint64(nid[:8])
	sid.High = binary.LittleEndian.Uint64(nid[8:])
	return sid
}

// ServerID2NodeID : protocol.SERVER_ID 转化为 common.NodeID
func ServerID2NodeID(sid *protocol.SERVER_ID) common.NodeID {
	nid := common.NodeID{}
	binary.LittleEndian.PutUint64(nid[:8], sid.GetLow())
	binary.LittleEndian.PutUint64(nid[8:], sid.GetHigh())
	return nid
}

// ServerID2UUID : protocol.SERVER_ID 转化为 uuid.UUID
func ServerID2UUID(sid *protocol.SERVER_ID) uuid.UUID {
	uid := uuid.UUID{}
	binary.LittleEndian.PutUint64(uid[:8], sid.GetLow())
	binary.LittleEndian.PutUint64(uid[8:], sid.GetHigh())
	return uid
}

// EqualSID : 2 个 SID 是否相等
func EqualSID(sid1 *protocol.SERVER_ID, sid2 *protocol.SERVER_ID) bool {
	return sid1.GetLow() == sid2.GetLow() && sid1.GetHigh() == sid2.GetHigh()
}

// EqualNID : 2 个 NID 是否相等
func EqualNID(nid1 common.NodeID, nid2 common.NodeID) bool {
	return bytes.Equal(nid1[:], nid2[:])
}
