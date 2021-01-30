// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/loki/loki.proto

package loki

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PushRequest struct {
	Streams []*Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams"`
}

func (m *PushRequest) Reset()      { *m = PushRequest{} }
func (*PushRequest) ProtoMessage() {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20a29f97a8c073b4, []int{0}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

type PushResponse struct {
}

func (m *PushResponse) Reset()      { *m = PushResponse{} }
func (*PushResponse) ProtoMessage() {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20a29f97a8c073b4, []int{1}
}
func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return m.Size()
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

type Stream struct {
	Labels  string  `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels"`
	Entries []Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
}

func (m *Stream) Reset()      { *m = Stream{} }
func (*Stream) ProtoMessage() {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_20a29f97a8c073b4, []int{2}
}
func (m *Stream) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(m, src)
}
func (m *Stream) XXX_Size() int {
	return m.Size()
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Stream) GetEntries() []Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"ts"`
	Line      string    `protobuf:"bytes,2,opt,name=line,proto3" json:"line"`
}

func (m *Entry) Reset()      { *m = Entry{} }
func (*Entry) ProtoMessage() {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_20a29f97a8c073b4, []int{3}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Entry) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "loki.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "loki.PushResponse")
	proto.RegisterType((*Stream)(nil), "loki.Stream")
	proto.RegisterType((*Entry)(nil), "loki.Entry")
}

func init() { proto.RegisterFile("proto/loki/loki.proto", fileDescriptor_20a29f97a8c073b4) }

var fileDescriptor_20a29f97a8c073b4 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x67, 0x10, 0x41, 0xa6, 0x04, 0x93, 0x26, 0x26, 0x84, 0x98, 0x57, 0xd3, 0x15, 0x1b,
	0xdb, 0x04, 0x8c, 0x07, 0x68, 0x62, 0xe2, 0xd2, 0x8c, 0x5e, 0x80, 0xea, 0x58, 0x1a, 0xdb, 0x0e,
	0x76, 0xa6, 0x0b, 0x77, 0x1e, 0x81, 0x63, 0x78, 0x14, 0x96, 0x2c, 0x59, 0x55, 0x19, 0x36, 0x86,
	0x15, 0x47, 0x30, 0x9d, 0x69, 0x65, 0xf3, 0xfa, 0xbf, 0x97, 0xaf, 0xff, 0xff, 0xe6, 0x91, 0x8b,
	0x45, 0xce, 0x25, 0xf7, 0x13, 0xfe, 0x16, 0xeb, 0xe2, 0xe9, 0xde, 0x6e, 0x57, 0x7a, 0xe4, 0x44,
	0x9c, 0x47, 0x09, 0xf3, 0xf5, 0x2c, 0x2c, 0x5e, 0x7d, 0x19, 0xa7, 0x4c, 0xc8, 0x59, 0xba, 0x30,
	0xd8, 0xe8, 0x3a, 0x8a, 0xe5, 0xbc, 0x08, 0xbd, 0x67, 0x9e, 0xfa, 0x11, 0x8f, 0xf8, 0x91, 0xac,
	0x3a, 0x63, 0x5d, 0x29, 0x83, 0xbb, 0x01, 0xb1, 0x1e, 0x0a, 0x31, 0xa7, 0xec, 0xbd, 0x60, 0x42,
	0xda, 0x53, 0xd2, 0x15, 0x32, 0x67, 0xb3, 0x54, 0x0c, 0xf1, 0xd5, 0xc9, 0xd8, 0x9a, 0xf4, 0x3d,
	0xbd, 0xc2, 0xa3, 0x1e, 0x06, 0xd6, 0xbe, 0x74, 0x1a, 0x80, 0x36, 0xc2, 0x1d, 0x90, 0xbe, 0xf1,
	0x10, 0x0b, 0x9e, 0x09, 0xe6, 0xbe, 0x90, 0x8e, 0xe1, 0x6d, 0x97, 0x74, 0x92, 0x59, 0xc8, 0x92,
	0xca, 0x0d, 0x8f, 0x7b, 0x01, 0xd9, 0x97, 0x4e, 0x3d, 0xa1, 0xf5, 0xd7, 0xbe, 0x25, 0x5d, 0x96,
	0xc9, 0x3c, 0x66, 0x62, 0xd8, 0xd2, 0x91, 0x96, 0x89, 0xbc, 0xcb, 0x64, 0xfe, 0x11, 0x9c, 0xaf,
	0x4a, 0x07, 0x55, 0xa9, 0x35, 0x43, 0x1b, 0xe1, 0x72, 0x72, 0xaa, 0x11, 0xfb, 0x9e, 0xf4, 0xfe,
	0x8f, 0xa0, 0x73, 0xac, 0xc9, 0xc8, 0x33, 0x67, 0xf2, 0x9a, 0xc7, 0x7b, 0x4f, 0x0d, 0x11, 0x0c,
	0x6a, 0xc7, 0x96, 0x14, 0xcb, 0x6f, 0x07, 0xd3, 0xe3, 0xcf, 0xf6, 0x25, 0x69, 0x27, 0x71, 0xc6,
	0x86, 0x2d, 0xbd, 0xec, 0xd9, 0xbe, 0x74, 0x74, 0x4f, 0x75, 0x0d, 0x6e, 0xd6, 0x5b, 0x40, 0x9b,
	0x2d, 0xa0, 0xc3, 0x16, 0xf0, 0xa7, 0x02, 0xfc, 0xa5, 0x00, 0xaf, 0x14, 0xe0, 0xb5, 0x02, 0xfc,
	0xa3, 0x00, 0xff, 0x2a, 0x40, 0x07, 0x05, 0x78, 0xb9, 0x03, 0xb4, 0xde, 0x01, 0xda, 0xec, 0x00,
	0x85, 0x1d, 0xbd, 0xc2, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa4, 0x24, 0x23, 0xd6, 0x01,
	0x00, 0x00,
}

func (this *PushRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushRequest)
	if !ok {
		that2, ok := that.(PushRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Streams) != len(that1.Streams) {
		return false
	}
	for i := range this.Streams {
		if !this.Streams[i].Equal(that1.Streams[i]) {
			return false
		}
	}
	return true
}
func (this *PushResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushResponse)
	if !ok {
		that2, ok := that.(PushResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *Stream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Stream)
	if !ok {
		that2, ok := that.(Stream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Labels != that1.Labels {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(&that1.Entries[i]) {
			return false
		}
	}
	return true
}
func (this *Entry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Entry)
	if !ok {
		that2, ok := that.(Entry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if this.Line != that1.Line {
		return false
	}
	return true
}
func (this *PushRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&loki.PushRequest{")
	if this.Streams != nil {
		s = append(s, "Streams: "+fmt.Sprintf("%#v", this.Streams)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&loki.PushResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Stream) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&loki.Stream{")
	s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	if this.Entries != nil {
		vs := make([]Entry, len(this.Entries))
		for i := range vs {
			vs[i] = this.Entries[i]
		}
		s = append(s, "Entries: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Entry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&loki.Entry{")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Line: "+fmt.Sprintf("%#v", this.Line)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLoki(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for iNdEx := len(m.Streams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Streams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLoki(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Stream) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stream) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stream) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLoki(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		i -= len(m.Labels)
		copy(dAtA[i:], m.Labels)
		i = encodeVarintLoki(dAtA, i, uint64(len(m.Labels)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Line) > 0 {
		i -= len(m.Line)
		copy(dAtA[i:], m.Line)
		i = encodeVarintLoki(dAtA, i, uint64(len(m.Line)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLoki(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLoki(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoki(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PushRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for _, e := range m.Streams {
			l = e.Size()
			n += 1 + l + sovLoki(uint64(l))
		}
	}
	return n
}

func (m *PushResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Stream) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Labels)
	if l > 0 {
		n += 1 + l + sovLoki(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovLoki(uint64(l))
		}
	}
	return n
}

func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLoki(uint64(l))
	l = len(m.Line)
	if l > 0 {
		n += 1 + l + sovLoki(uint64(l))
	}
	return n
}

func sovLoki(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoki(x uint64) (n int) {
	return sovLoki(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PushRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForStreams := "[]*Stream{"
	for _, f := range this.Streams {
		repeatedStringForStreams += strings.Replace(f.String(), "Stream", "Stream", 1) + ","
	}
	repeatedStringForStreams += "}"
	s := strings.Join([]string{`&PushRequest{`,
		`Streams:` + repeatedStringForStreams + `,`,
		`}`,
	}, "")
	return s
}
func (this *PushResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PushResponse{`,
		`}`,
	}, "")
	return s
}
func (this *Stream) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEntries := "[]Entry{"
	for _, f := range this.Entries {
		repeatedStringForEntries += strings.Replace(strings.Replace(f.String(), "Entry", "Entry", 1), `&`, ``, 1) + ","
	}
	repeatedStringForEntries += "}"
	s := strings.Join([]string{`&Stream{`,
		`Labels:` + fmt.Sprintf("%v", this.Labels) + `,`,
		`Entries:` + repeatedStringForEntries + `,`,
		`}`,
	}, "")
	return s
}
func (this *Entry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Entry{`,
		`Timestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "timestamp.Timestamp", 1), `&`, ``, 1) + `,`,
		`Line:` + fmt.Sprintf("%v", this.Line) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLoki(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoki
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLoki
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoki
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams, &Stream{})
			if err := m.Streams[len(m.Streams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoki(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoki
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLoki
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
func (m *PushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoki
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLoki(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoki
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLoki
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
func (m *Stream) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoki
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stream: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stream: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoki
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoki
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLoki
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoki
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoki(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoki
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLoki
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
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoki
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLoki
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoki
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoki
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoki
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Line = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoki(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoki
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLoki
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
func skipLoki(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoki
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLoki
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLoki
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoki
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoki
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoki        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoki          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoki = fmt.Errorf("proto: unexpected end of group")
)