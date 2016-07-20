// Code generated by protoc-gen-gogo.
// source: keyserverlocal.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ReplicaState contains the persistent internal state of a single replica.
// Additional on-disk state is descried in server/table.go.
type ReplicaState struct {
	// cached values derived purely from the state of the log
	NextIndexLog                    uint64         `protobuf:"varint,1,opt,name=next_index_log,json=nextIndexLog,proto3" json:"next_index_log,omitempty"`
	NextIndexVerifier               uint64         `protobuf:"varint,2,opt,name=next_index_verifier,json=nextIndexVerifier,proto3" json:"next_index_verifier,omitempty"`
	PreviousSummaryHash             []byte         `protobuf:"bytes,3,opt,name=previous_summary_hash,json=previousSummaryHash,proto3" json:"previous_summary_hash,omitempty"`
	LastEpochDelimiter              EpochDelimiter `protobuf:"bytes,4,opt,name=last_epoch_delimiter,json=lastEpochDelimiter" json:"last_epoch_delimiter"`
	ThisReplicaNeedsToSignLastEpoch bool           `protobuf:"varint,5,opt,name=this_replica_needs_to_sign_last_epoch,json=thisReplicaNeedsToSignLastEpoch,proto3" json:"this_replica_needs_to_sign_last_epoch,omitempty"`
	PendingUpdates                  bool           `protobuf:"varint,6,opt,name=pending_updates,json=pendingUpdates,proto3" json:"pending_updates,omitempty"`
	// local variables
	LatestTreeSnapshot         uint64 `protobuf:"varint,7,opt,name=latest_tree_snapshot,json=latestTreeSnapshot,proto3" json:"latest_tree_snapshot,omitempty"`
	LastEpochNeedsRatification bool   `protobuf:"varint,8,opt,name=last_epoch_needs_ratification,json=lastEpochNeedsRatification,proto3" json:"last_epoch_needs_ratification,omitempty"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorKeyserverlocal, []int{0} }

func (m *ReplicaState) GetLastEpochDelimiter() EpochDelimiter {
	if m != nil {
		return m.LastEpochDelimiter
	}
	return EpochDelimiter{}
}

func init() {
	proto1.RegisterType((*ReplicaState)(nil), "proto.ReplicaState")
}
func (this *ReplicaState) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ReplicaState)
	if !ok {
		that2, ok := that.(ReplicaState)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ReplicaState")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ReplicaState but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ReplicaState but is not nil && this == nil")
	}
	if this.NextIndexLog != that1.NextIndexLog {
		return fmt.Errorf("NextIndexLog this(%v) Not Equal that(%v)", this.NextIndexLog, that1.NextIndexLog)
	}
	if this.NextIndexVerifier != that1.NextIndexVerifier {
		return fmt.Errorf("NextIndexVerifier this(%v) Not Equal that(%v)", this.NextIndexVerifier, that1.NextIndexVerifier)
	}
	if !bytes.Equal(this.PreviousSummaryHash, that1.PreviousSummaryHash) {
		return fmt.Errorf("PreviousSummaryHash this(%v) Not Equal that(%v)", this.PreviousSummaryHash, that1.PreviousSummaryHash)
	}
	if !this.LastEpochDelimiter.Equal(&that1.LastEpochDelimiter) {
		return fmt.Errorf("LastEpochDelimiter this(%v) Not Equal that(%v)", this.LastEpochDelimiter, that1.LastEpochDelimiter)
	}
	if this.ThisReplicaNeedsToSignLastEpoch != that1.ThisReplicaNeedsToSignLastEpoch {
		return fmt.Errorf("ThisReplicaNeedsToSignLastEpoch this(%v) Not Equal that(%v)", this.ThisReplicaNeedsToSignLastEpoch, that1.ThisReplicaNeedsToSignLastEpoch)
	}
	if this.PendingUpdates != that1.PendingUpdates {
		return fmt.Errorf("PendingUpdates this(%v) Not Equal that(%v)", this.PendingUpdates, that1.PendingUpdates)
	}
	if this.LatestTreeSnapshot != that1.LatestTreeSnapshot {
		return fmt.Errorf("LatestTreeSnapshot this(%v) Not Equal that(%v)", this.LatestTreeSnapshot, that1.LatestTreeSnapshot)
	}
	if this.LastEpochNeedsRatification != that1.LastEpochNeedsRatification {
		return fmt.Errorf("LastEpochNeedsRatification this(%v) Not Equal that(%v)", this.LastEpochNeedsRatification, that1.LastEpochNeedsRatification)
	}
	return nil
}
func (this *ReplicaState) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ReplicaState)
	if !ok {
		that2, ok := that.(ReplicaState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.NextIndexLog != that1.NextIndexLog {
		return false
	}
	if this.NextIndexVerifier != that1.NextIndexVerifier {
		return false
	}
	if !bytes.Equal(this.PreviousSummaryHash, that1.PreviousSummaryHash) {
		return false
	}
	if !this.LastEpochDelimiter.Equal(&that1.LastEpochDelimiter) {
		return false
	}
	if this.ThisReplicaNeedsToSignLastEpoch != that1.ThisReplicaNeedsToSignLastEpoch {
		return false
	}
	if this.PendingUpdates != that1.PendingUpdates {
		return false
	}
	if this.LatestTreeSnapshot != that1.LatestTreeSnapshot {
		return false
	}
	if this.LastEpochNeedsRatification != that1.LastEpochNeedsRatification {
		return false
	}
	return true
}
func (this *ReplicaState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&proto.ReplicaState{")
	s = append(s, "NextIndexLog: "+fmt.Sprintf("%#v", this.NextIndexLog)+",\n")
	s = append(s, "NextIndexVerifier: "+fmt.Sprintf("%#v", this.NextIndexVerifier)+",\n")
	s = append(s, "PreviousSummaryHash: "+fmt.Sprintf("%#v", this.PreviousSummaryHash)+",\n")
	s = append(s, "LastEpochDelimiter: "+strings.Replace(this.LastEpochDelimiter.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "ThisReplicaNeedsToSignLastEpoch: "+fmt.Sprintf("%#v", this.ThisReplicaNeedsToSignLastEpoch)+",\n")
	s = append(s, "PendingUpdates: "+fmt.Sprintf("%#v", this.PendingUpdates)+",\n")
	s = append(s, "LatestTreeSnapshot: "+fmt.Sprintf("%#v", this.LatestTreeSnapshot)+",\n")
	s = append(s, "LastEpochNeedsRatification: "+fmt.Sprintf("%#v", this.LastEpochNeedsRatification)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringKeyserverlocal(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringKeyserverlocal(e map[int32]github_com_maditya_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *ReplicaState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReplicaState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NextIndexLog != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintKeyserverlocal(data, i, uint64(m.NextIndexLog))
	}
	if m.NextIndexVerifier != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintKeyserverlocal(data, i, uint64(m.NextIndexVerifier))
	}
	if len(m.PreviousSummaryHash) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintKeyserverlocal(data, i, uint64(len(m.PreviousSummaryHash)))
		i += copy(data[i:], m.PreviousSummaryHash)
	}
	data[i] = 0x22
	i++
	i = encodeVarintKeyserverlocal(data, i, uint64(m.LastEpochDelimiter.Size()))
	n1, err := m.LastEpochDelimiter.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.ThisReplicaNeedsToSignLastEpoch {
		data[i] = 0x28
		i++
		if m.ThisReplicaNeedsToSignLastEpoch {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.PendingUpdates {
		data[i] = 0x30
		i++
		if m.PendingUpdates {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.LatestTreeSnapshot != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintKeyserverlocal(data, i, uint64(m.LatestTreeSnapshot))
	}
	if m.LastEpochNeedsRatification {
		data[i] = 0x40
		i++
		if m.LastEpochNeedsRatification {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Keyserverlocal(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Keyserverlocal(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintKeyserverlocal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedReplicaState(r randyKeyserverlocal, easy bool) *ReplicaState {
	this := &ReplicaState{}
	this.NextIndexLog = uint64(uint64(r.Uint32()))
	this.NextIndexVerifier = uint64(uint64(r.Uint32()))
	v1 := r.Intn(100)
	this.PreviousSummaryHash = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.PreviousSummaryHash[i] = byte(r.Intn(256))
	}
	v2 := NewPopulatedEpochDelimiter(r, easy)
	this.LastEpochDelimiter = *v2
	this.ThisReplicaNeedsToSignLastEpoch = bool(bool(r.Intn(2) == 0))
	this.PendingUpdates = bool(bool(r.Intn(2) == 0))
	this.LatestTreeSnapshot = uint64(uint64(r.Uint32()))
	this.LastEpochNeedsRatification = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyKeyserverlocal interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneKeyserverlocal(r randyKeyserverlocal) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringKeyserverlocal(r randyKeyserverlocal) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneKeyserverlocal(r)
	}
	return string(tmps)
}
func randUnrecognizedKeyserverlocal(r randyKeyserverlocal, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldKeyserverlocal(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldKeyserverlocal(data []byte, r randyKeyserverlocal, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateKeyserverlocal(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateKeyserverlocal(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateKeyserverlocal(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateKeyserverlocal(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateKeyserverlocal(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateKeyserverlocal(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateKeyserverlocal(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.NextIndexLog != 0 {
		n += 1 + sovKeyserverlocal(uint64(m.NextIndexLog))
	}
	if m.NextIndexVerifier != 0 {
		n += 1 + sovKeyserverlocal(uint64(m.NextIndexVerifier))
	}
	l = len(m.PreviousSummaryHash)
	if l > 0 {
		n += 1 + l + sovKeyserverlocal(uint64(l))
	}
	l = m.LastEpochDelimiter.Size()
	n += 1 + l + sovKeyserverlocal(uint64(l))
	if m.ThisReplicaNeedsToSignLastEpoch {
		n += 2
	}
	if m.PendingUpdates {
		n += 2
	}
	if m.LatestTreeSnapshot != 0 {
		n += 1 + sovKeyserverlocal(uint64(m.LatestTreeSnapshot))
	}
	if m.LastEpochNeedsRatification {
		n += 2
	}
	return n
}

func sovKeyserverlocal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKeyserverlocal(x uint64) (n int) {
	return sovKeyserverlocal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReplicaState) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaState{`,
		`NextIndexLog:` + fmt.Sprintf("%v", this.NextIndexLog) + `,`,
		`NextIndexVerifier:` + fmt.Sprintf("%v", this.NextIndexVerifier) + `,`,
		`PreviousSummaryHash:` + fmt.Sprintf("%v", this.PreviousSummaryHash) + `,`,
		`LastEpochDelimiter:` + strings.Replace(strings.Replace(this.LastEpochDelimiter.String(), "EpochDelimiter", "EpochDelimiter", 1), `&`, ``, 1) + `,`,
		`ThisReplicaNeedsToSignLastEpoch:` + fmt.Sprintf("%v", this.ThisReplicaNeedsToSignLastEpoch) + `,`,
		`PendingUpdates:` + fmt.Sprintf("%v", this.PendingUpdates) + `,`,
		`LatestTreeSnapshot:` + fmt.Sprintf("%v", this.LatestTreeSnapshot) + `,`,
		`LastEpochNeedsRatification:` + fmt.Sprintf("%v", this.LastEpochNeedsRatification) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringKeyserverlocal(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReplicaState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyserverlocal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndexLog", wireType)
			}
			m.NextIndexLog = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndexLog |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndexVerifier", wireType)
			}
			m.NextIndexVerifier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndexVerifier |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSummaryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKeyserverlocal
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousSummaryHash = append(m.PreviousSummaryHash[:0], data[iNdEx:postIndex]...)
			if m.PreviousSummaryHash == nil {
				m.PreviousSummaryHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpochDelimiter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKeyserverlocal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastEpochDelimiter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThisReplicaNeedsToSignLastEpoch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ThisReplicaNeedsToSignLastEpoch = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingUpdates", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PendingUpdates = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestTreeSnapshot", wireType)
			}
			m.LatestTreeSnapshot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LatestTreeSnapshot |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpochNeedsRatification", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LastEpochNeedsRatification = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipKeyserverlocal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeyserverlocal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKeyserverlocal(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyserverlocal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKeyserverlocal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKeyserverlocal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKeyserverlocal
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipKeyserverlocal(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthKeyserverlocal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyserverlocal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("keyserverlocal.proto", fileDescriptorKeyserverlocal) }

var fileDescriptorKeyserverlocal = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x6f, 0xd4, 0x30,
	0x14, 0xc6, 0x2f, 0xf4, 0x5a, 0x2a, 0x73, 0x2a, 0xaa, 0xdb, 0x4a, 0xd1, 0x49, 0xa4, 0x08, 0x81,
	0x60, 0xca, 0x55, 0x65, 0x61, 0xa5, 0x02, 0x09, 0xa4, 0xd2, 0x21, 0x29, 0xac, 0x96, 0x2f, 0x79,
	0x97, 0x58, 0x24, 0x76, 0x64, 0x3b, 0x55, 0x6f, 0xe3, 0xcf, 0xe1, 0x4f, 0x60, 0x64, 0xec, 0xd8,
	0x05, 0x89, 0x09, 0xb5, 0x9d, 0x18, 0x6f, 0x64, 0xe4, 0xc5, 0xce, 0x1d, 0xc7, 0xf0, 0x29, 0xce,
	0xf7, 0xfb, 0x9e, 0xed, 0xf7, 0x4c, 0xf6, 0x3f, 0xc3, 0xdc, 0x80, 0xbe, 0x00, 0x5d, 0xa9, 0x8c,
	0x57, 0x71, 0xa3, 0x95, 0x55, 0x74, 0xd3, 0x7d, 0xc6, 0x47, 0x85, 0xb0, 0x65, 0x3b, 0x8d, 0x33,
	0x55, 0x4f, 0x6a, 0x9e, 0x0b, 0x3b, 0xe7, 0x13, 0x47, 0xa6, 0xed, 0x6c, 0x52, 0xa8, 0x42, 0xb9,
	0x1f, 0xb7, 0xf2, 0x85, 0xe3, 0x5d, 0x0d, 0x4d, 0x25, 0x32, 0x6e, 0x85, 0x92, 0xde, 0x7a, 0xf2,
	0x63, 0x83, 0x8c, 0x12, 0xef, 0xa6, 0x96, 0x5b, 0xa0, 0x4f, 0xc9, 0x8e, 0x84, 0x4b, 0xcb, 0x84,
	0xcc, 0xe1, 0x92, 0x55, 0xaa, 0x08, 0x83, 0xc7, 0xc1, 0x8b, 0x61, 0x32, 0xea, 0xdc, 0xf7, 0x9d,
	0x79, 0xaa, 0x0a, 0x1a, 0x93, 0xbd, 0xb5, 0x14, 0xde, 0x4f, 0xcc, 0x04, 0xe8, 0xf0, 0x9e, 0x8b,
	0xee, 0xae, 0xa2, 0x9f, 0x7a, 0x40, 0x8f, 0xc9, 0x41, 0xa3, 0xe1, 0x42, 0xa8, 0xd6, 0x30, 0xd3,
	0xd6, 0x35, 0xd7, 0x73, 0x56, 0x72, 0x53, 0x86, 0x1b, 0x58, 0x31, 0x4a, 0xf6, 0x96, 0x30, 0xf5,
	0xec, 0x1d, 0x22, 0xfa, 0x81, 0xec, 0x57, 0xdc, 0x58, 0x06, 0x8d, 0xca, 0x4a, 0x96, 0x43, 0x25,
	0x6a, 0x61, 0xf1, 0x90, 0x21, 0x96, 0x3c, 0x38, 0x3e, 0xf0, 0x0d, 0xc4, 0x6f, 0x3b, 0xfa, 0x66,
	0x09, 0x4f, 0x86, 0x57, 0xbf, 0x0e, 0x07, 0x09, 0xed, 0x0a, 0xff, 0x27, 0xf4, 0x8c, 0x3c, 0xb3,
	0xa5, 0x30, 0xac, 0x9f, 0x01, 0x93, 0x00, 0xb9, 0x61, 0x56, 0x31, 0x23, 0x0a, 0xc9, 0xfe, 0x9d,
	0x14, 0x6e, 0xe2, 0xfe, 0xdb, 0xc9, 0x61, 0x17, 0xee, 0x27, 0x73, 0xd6, 0x45, 0xcf, 0x55, 0x8a,
	0xc1, 0xd3, 0xe5, 0xc6, 0xf4, 0x39, 0x79, 0xd8, 0x80, 0xcc, 0x85, 0x2c, 0x58, 0xdb, 0xe4, 0x38,
	0x3a, 0x13, 0x6e, 0xb9, 0xca, 0x9d, 0xde, 0xfe, 0xe8, 0x5d, 0x7a, 0xd4, 0xf5, 0x81, 0x0b, 0xcb,
	0xac, 0x06, 0x60, 0x46, 0xf2, 0xc6, 0x94, 0xca, 0x86, 0xf7, 0xdd, 0xb0, 0xa8, 0x67, 0xe7, 0x88,
	0xd2, 0x9e, 0xd0, 0xd7, 0xe4, 0xd1, 0x5a, 0xe7, 0xfe, 0xa2, 0x1a, 0x5f, 0x6d, 0xd6, 0xbf, 0x5d,
	0xb8, 0xed, 0x0e, 0x1a, 0xaf, 0xba, 0x74, 0x17, 0x4c, 0xd6, 0x12, 0x27, 0xaf, 0xae, 0x6f, 0xa3,
	0xc1, 0x4f, 0xd4, 0xcd, 0x6d, 0x14, 0x2c, 0x50, 0x7f, 0x50, 0x5f, 0xee, 0xa2, 0xe0, 0x2b, 0xea,
	0x1b, 0xea, 0x3b, 0xea, 0x0a, 0x75, 0x8d, 0xba, 0x41, 0xfd, 0xbe, 0x8b, 0x06, 0x0b, 0xfc, 0x4e,
	0xb7, 0xdc, 0x5c, 0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x94, 0xcf, 0xd3, 0x7c, 0x02,
	0x00, 0x00,
}
