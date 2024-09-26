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

// OpcuaProtocolLimits is the corresponding interface of OpcuaProtocolLimits
type OpcuaProtocolLimits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetReceiveBufferSize returns ReceiveBufferSize (property field)
	GetReceiveBufferSize() uint32
	// GetSendBufferSize returns SendBufferSize (property field)
	GetSendBufferSize() uint32
	// GetMaxMessageSize returns MaxMessageSize (property field)
	GetMaxMessageSize() uint32
	// GetMaxChunkCount returns MaxChunkCount (property field)
	GetMaxChunkCount() uint32
	// IsOpcuaProtocolLimits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaProtocolLimits()
}

// _OpcuaProtocolLimits is the data-structure of this message
type _OpcuaProtocolLimits struct {
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

var _ OpcuaProtocolLimits = (*_OpcuaProtocolLimits)(nil)

// NewOpcuaProtocolLimits factory function for _OpcuaProtocolLimits
func NewOpcuaProtocolLimits(receiveBufferSize uint32, sendBufferSize uint32, maxMessageSize uint32, maxChunkCount uint32) *_OpcuaProtocolLimits {
	return &_OpcuaProtocolLimits{ReceiveBufferSize: receiveBufferSize, SendBufferSize: sendBufferSize, MaxMessageSize: maxMessageSize, MaxChunkCount: maxChunkCount}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaProtocolLimits) GetReceiveBufferSize() uint32 {
	return m.ReceiveBufferSize
}

func (m *_OpcuaProtocolLimits) GetSendBufferSize() uint32 {
	return m.SendBufferSize
}

func (m *_OpcuaProtocolLimits) GetMaxMessageSize() uint32 {
	return m.MaxMessageSize
}

func (m *_OpcuaProtocolLimits) GetMaxChunkCount() uint32 {
	return m.MaxChunkCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaProtocolLimits(structType any) OpcuaProtocolLimits {
	if casted, ok := structType.(OpcuaProtocolLimits); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaProtocolLimits); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaProtocolLimits) GetTypeName() string {
	return "OpcuaProtocolLimits"
}

func (m *_OpcuaProtocolLimits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (receiveBufferSize)
	lengthInBits += 32

	// Simple field (sendBufferSize)
	lengthInBits += 32

	// Simple field (maxMessageSize)
	lengthInBits += 32

	// Simple field (maxChunkCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_OpcuaProtocolLimits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaProtocolLimitsParse(ctx context.Context, theBytes []byte) (OpcuaProtocolLimits, error) {
	return OpcuaProtocolLimitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaProtocolLimitsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaProtocolLimits, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaProtocolLimits, error) {
		return OpcuaProtocolLimitsParseWithBuffer(ctx, readBuffer)
	}
}

func OpcuaProtocolLimitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaProtocolLimits, error) {
	v, err := (&_OpcuaProtocolLimits{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_OpcuaProtocolLimits) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__opcuaProtocolLimits OpcuaProtocolLimits, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaProtocolLimits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaProtocolLimits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	receiveBufferSize, err := ReadSimpleField(ctx, "receiveBufferSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiveBufferSize' field"))
	}
	m.ReceiveBufferSize = receiveBufferSize

	sendBufferSize, err := ReadSimpleField(ctx, "sendBufferSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sendBufferSize' field"))
	}
	m.SendBufferSize = sendBufferSize

	maxMessageSize, err := ReadSimpleField(ctx, "maxMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxMessageSize' field"))
	}
	m.MaxMessageSize = maxMessageSize

	maxChunkCount, err := ReadSimpleField(ctx, "maxChunkCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxChunkCount' field"))
	}
	m.MaxChunkCount = maxChunkCount

	if closeErr := readBuffer.CloseContext("OpcuaProtocolLimits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaProtocolLimits")
	}

	return m, nil
}

func (m *_OpcuaProtocolLimits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaProtocolLimits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("OpcuaProtocolLimits"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for OpcuaProtocolLimits")
	}

	if err := WriteSimpleField[uint32](ctx, "receiveBufferSize", m.GetReceiveBufferSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'receiveBufferSize' field")
	}

	if err := WriteSimpleField[uint32](ctx, "sendBufferSize", m.GetSendBufferSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'sendBufferSize' field")
	}

	if err := WriteSimpleField[uint32](ctx, "maxMessageSize", m.GetMaxMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'maxMessageSize' field")
	}

	if err := WriteSimpleField[uint32](ctx, "maxChunkCount", m.GetMaxChunkCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'maxChunkCount' field")
	}

	if popErr := writeBuffer.PopContext("OpcuaProtocolLimits"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for OpcuaProtocolLimits")
	}
	return nil
}

func (m *_OpcuaProtocolLimits) IsOpcuaProtocolLimits() {}

func (m *_OpcuaProtocolLimits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaProtocolLimits) deepCopy() *_OpcuaProtocolLimits {
	if m == nil {
		return nil
	}
	_OpcuaProtocolLimitsCopy := &_OpcuaProtocolLimits{
		m.ReceiveBufferSize,
		m.SendBufferSize,
		m.MaxMessageSize,
		m.MaxChunkCount,
	}
	return _OpcuaProtocolLimitsCopy
}

func (m *_OpcuaProtocolLimits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
