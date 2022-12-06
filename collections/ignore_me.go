package collections

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

func ProtoValueEncoder[T any, PT interface {
	*T
	codec.ProtoMarshaler
}](cdc codec.BinaryCodec) ValueEncoder[T] {
	return protoValueEncoder[T, PT]{cdc: cdc}
}

type protoValueEncoder[T any, PT interface {
	*T
	codec.ProtoMarshaler
}] struct {
	cdc codec.BinaryCodec
	_   T
}

func (p protoValueEncoder[T, PT]) Encode(value T) ([]byte, error) {
	return p.cdc.Marshal(PT(&value))
}

func (p protoValueEncoder[T, PT]) Decode(b []byte) (T, error) {
	value := new(T)
	err := p.cdc.Unmarshal(b, PT(value))
	if err != nil {
		var v T
		return v, err
	}
	return *value, nil
}

func (p protoValueEncoder[T, PT]) Stringify(value T) string {
	return PT(&value).String()
}

func (p protoValueEncoder[T, PT]) ValueType() string {
	return proto.MessageName(PT(new(T)))
}
