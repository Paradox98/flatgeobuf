// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package FlatGeobuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Column struct {
	_tab flatbuffers.Table
}

func GetRootAsColumn(buf []byte, offset flatbuffers.UOffsetT) *Column {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Column{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Column) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Column) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Column) Type() ColumnType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Column) MutateType(n ColumnType) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func ColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ColumnAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ColumnAddType(builder *flatbuffers.Builder, type_ byte) {
	builder.PrependByteSlot(1, type_, 0)
}
func ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
