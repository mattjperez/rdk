// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NotEqualOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsNotEqualOptions(buf []byte, offset flatbuffers.UOffsetT) *NotEqualOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NotEqualOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNotEqualOptions(buf []byte, offset flatbuffers.UOffsetT) *NotEqualOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NotEqualOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NotEqualOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NotEqualOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func NotEqualOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func NotEqualOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
