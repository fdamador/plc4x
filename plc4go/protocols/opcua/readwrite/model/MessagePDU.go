/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MessagePDU is the corresponding interface of MessagePDU
type MessagePDU interface {
	MessagePDUContract
	MessagePDURequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsMessagePDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMessagePDU()
	// CreateBuilder creates a MessagePDUBuilder
	CreateMessagePDUBuilder() MessagePDUBuilder
}

// MessagePDUContract provides a set of functions which can be overwritten by a sub struct
type MessagePDUContract interface {
	// GetChunk returns Chunk (property field)
	GetChunk() ChunkType
	// GetBinary() returns a parser argument
	GetBinary() bool
	// IsMessagePDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMessagePDU()
	// CreateBuilder creates a MessagePDUBuilder
	CreateMessagePDUBuilder() MessagePDUBuilder
}

// MessagePDURequirements provides a set of functions which need to be implemented by a sub struct
type MessagePDURequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() string
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// _MessagePDU is the data-structure of this message
type _MessagePDU struct {
	_SubType interface {
		MessagePDUContract
		MessagePDURequirements
	}
	Chunk ChunkType

	// Arguments.
	Binary bool
}

var _ MessagePDUContract = (*_MessagePDU)(nil)

// NewMessagePDU factory function for _MessagePDU
func NewMessagePDU(chunk ChunkType, binary bool) *_MessagePDU {
	return &_MessagePDU{Chunk: chunk, Binary: binary}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MessagePDUBuilder is a builder for MessagePDU
type MessagePDUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(chunk ChunkType) MessagePDUBuilder
	// WithChunk adds Chunk (property field)
	WithChunk(ChunkType) MessagePDUBuilder
	// AsOpcuaHelloRequest converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaHelloRequest() interface {
		OpcuaHelloRequestBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaAcknowledgeResponse converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaAcknowledgeResponse() interface {
		OpcuaAcknowledgeResponseBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaOpenRequest converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaOpenRequest() interface {
		OpcuaOpenRequestBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaOpenResponse converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaOpenResponse() interface {
		OpcuaOpenResponseBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaCloseRequest converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaCloseRequest() interface {
		OpcuaCloseRequestBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaMessageRequest converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaMessageRequest() interface {
		OpcuaMessageRequestBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaMessageResponse converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaMessageResponse() interface {
		OpcuaMessageResponseBuilder
		Done() MessagePDUBuilder
	}
	// AsOpcuaMessageError converts this build to a subType of MessagePDU. It is always possible to return to current builder using Done()
	AsOpcuaMessageError() interface {
		OpcuaMessageErrorBuilder
		Done() MessagePDUBuilder
	}
	// Build builds the MessagePDU or returns an error if something is wrong
	PartialBuild() (MessagePDUContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() MessagePDUContract
	// Build builds the MessagePDU or returns an error if something is wrong
	Build() (MessagePDU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MessagePDU
}

// NewMessagePDUBuilder() creates a MessagePDUBuilder
func NewMessagePDUBuilder() MessagePDUBuilder {
	return &_MessagePDUBuilder{_MessagePDU: new(_MessagePDU)}
}

type _MessagePDUChildBuilder interface {
	utils.Copyable
	setParent(MessagePDUContract)
	buildForMessagePDU() (MessagePDU, error)
}

type _MessagePDUBuilder struct {
	*_MessagePDU

	childBuilder _MessagePDUChildBuilder

	err *utils.MultiError
}

var _ (MessagePDUBuilder) = (*_MessagePDUBuilder)(nil)

func (b *_MessagePDUBuilder) WithMandatoryFields(chunk ChunkType) MessagePDUBuilder {
	return b.WithChunk(chunk)
}

func (b *_MessagePDUBuilder) WithChunk(chunk ChunkType) MessagePDUBuilder {
	b.Chunk = chunk
	return b
}

func (b *_MessagePDUBuilder) PartialBuild() (MessagePDUContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MessagePDU.deepCopy(), nil
}

func (b *_MessagePDUBuilder) PartialMustBuild() MessagePDUContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MessagePDUBuilder) AsOpcuaHelloRequest() interface {
	OpcuaHelloRequestBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaHelloRequestBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaHelloRequestBuilder().(*_OpcuaHelloRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaAcknowledgeResponse() interface {
	OpcuaAcknowledgeResponseBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaAcknowledgeResponseBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaAcknowledgeResponseBuilder().(*_OpcuaAcknowledgeResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaOpenRequest() interface {
	OpcuaOpenRequestBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaOpenRequestBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaOpenRequestBuilder().(*_OpcuaOpenRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaOpenResponse() interface {
	OpcuaOpenResponseBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaOpenResponseBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaOpenResponseBuilder().(*_OpcuaOpenResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaCloseRequest() interface {
	OpcuaCloseRequestBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaCloseRequestBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaCloseRequestBuilder().(*_OpcuaCloseRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaMessageRequest() interface {
	OpcuaMessageRequestBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaMessageRequestBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaMessageRequestBuilder().(*_OpcuaMessageRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaMessageResponse() interface {
	OpcuaMessageResponseBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaMessageResponseBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaMessageResponseBuilder().(*_OpcuaMessageResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) AsOpcuaMessageError() interface {
	OpcuaMessageErrorBuilder
	Done() MessagePDUBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		OpcuaMessageErrorBuilder
		Done() MessagePDUBuilder
	}); ok {
		return cb
	}
	cb := NewOpcuaMessageErrorBuilder().(*_OpcuaMessageErrorBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MessagePDUBuilder) Build() (MessagePDU, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForMessagePDU()
}

func (b *_MessagePDUBuilder) MustBuild() MessagePDU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MessagePDUBuilder) DeepCopy() any {
	_copy := b.CreateMessagePDUBuilder().(*_MessagePDUBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_MessagePDUChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMessagePDUBuilder creates a MessagePDUBuilder
func (b *_MessagePDU) CreateMessagePDUBuilder() MessagePDUBuilder {
	if b == nil {
		return NewMessagePDUBuilder()
	}
	return &_MessagePDUBuilder{_MessagePDU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MessagePDU) GetChunk() ChunkType {
	return m.Chunk
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMessagePDU(structType any) MessagePDU {
	if casted, ok := structType.(MessagePDU); ok {
		return casted
	}
	if casted, ok := structType.(*MessagePDU); ok {
		return *casted
	}
	return nil
}

func (m *_MessagePDU) GetTypeName() string {
	return "MessagePDU"
}

func (m *_MessagePDU) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 24

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (totalLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MessagePDU) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_MessagePDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func MessagePDUParse[T MessagePDU](ctx context.Context, theBytes []byte, response bool, binary bool) (T, error) {
	return MessagePDUParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response, binary)
}

func MessagePDUParseWithBufferProducer[T MessagePDU](response bool, binary bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MessagePDUParseWithBuffer[T](ctx, readBuffer, response, binary)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func MessagePDUParseWithBuffer[T MessagePDU](ctx context.Context, readBuffer utils.ReadBuffer, response bool, binary bool) (T, error) {
	v, err := (&_MessagePDU{Binary: binary}).parse(ctx, readBuffer, response, binary)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_MessagePDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool, binary bool) (__messagePDU MessagePDU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MessagePDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MessagePDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[string](ctx, "messageType", ReadString(readBuffer, uint32(24)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	chunk, err := ReadEnumField[ChunkType](ctx, "chunk", "ChunkType", ReadEnum(ChunkTypeByValue, ReadString(readBuffer, uint32(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chunk' field"))
	}
	m.Chunk = chunk

	totalLength, err := ReadImplicitField[uint32](ctx, "totalLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalLength' field"))
	}
	_ = totalLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MessagePDU
	switch {
	case messageType == "HEL" && response == bool(false): // OpcuaHelloRequest
		if _child, err = new(_OpcuaHelloRequest).parse(ctx, readBuffer, m, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaHelloRequest for type-switch of MessagePDU")
		}
	case messageType == "ACK" && response == bool(true): // OpcuaAcknowledgeResponse
		if _child, err = new(_OpcuaAcknowledgeResponse).parse(ctx, readBuffer, m, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaAcknowledgeResponse for type-switch of MessagePDU")
		}
	case messageType == "OPN" && response == bool(false): // OpcuaOpenRequest
		if _child, err = new(_OpcuaOpenRequest).parse(ctx, readBuffer, m, totalLength, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaOpenRequest for type-switch of MessagePDU")
		}
	case messageType == "OPN" && response == bool(true): // OpcuaOpenResponse
		if _child, err = new(_OpcuaOpenResponse).parse(ctx, readBuffer, m, totalLength, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaOpenResponse for type-switch of MessagePDU")
		}
	case messageType == "CLO" && response == bool(false): // OpcuaCloseRequest
		if _child, err = new(_OpcuaCloseRequest).parse(ctx, readBuffer, m, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaCloseRequest for type-switch of MessagePDU")
		}
	case messageType == "MSG" && response == bool(false): // OpcuaMessageRequest
		if _child, err = new(_OpcuaMessageRequest).parse(ctx, readBuffer, m, totalLength, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageRequest for type-switch of MessagePDU")
		}
	case messageType == "MSG" && response == bool(true): // OpcuaMessageResponse
		if _child, err = new(_OpcuaMessageResponse).parse(ctx, readBuffer, m, totalLength, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageResponse for type-switch of MessagePDU")
		}
	case messageType == "ERR" && response == bool(true): // OpcuaMessageError
		if _child, err = new(_OpcuaMessageError).parse(ctx, readBuffer, m, response, binary); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageError for type-switch of MessagePDU")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v, response=%v]", messageType, response)
	}

	if closeErr := readBuffer.CloseContext("MessagePDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MessagePDU")
	}

	return _child, nil
}

func (pm *_MessagePDU) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MessagePDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MessagePDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MessagePDU")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteString(writeBuffer, 24)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}

	if err := WriteSimpleEnumField[ChunkType](ctx, "chunk", "ChunkType", m.GetChunk(), WriteEnum[ChunkType, string](ChunkType.GetValue, ChunkType.PLC4XEnumName, WriteString(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'chunk' field")
	}
	totalLength := uint32(uint32(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "totalLength", totalLength, WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MessagePDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MessagePDU")
	}
	return nil
}

////
// Arguments Getter

func (m *_MessagePDU) GetBinary() bool {
	return m.Binary
}

//
////

func (m *_MessagePDU) IsMessagePDU() {}

func (m *_MessagePDU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MessagePDU) deepCopy() *_MessagePDU {
	if m == nil {
		return nil
	}
	_MessagePDUCopy := &_MessagePDU{
		nil, // will be set by child
		m.Chunk,
		m.Binary,
	}
	return _MessagePDUCopy
}
