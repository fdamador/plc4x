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

// DeviceConfigurationRequestDataBlock is the corresponding interface of DeviceConfigurationRequestDataBlock
type DeviceConfigurationRequestDataBlock interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// IsDeviceConfigurationRequestDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationRequestDataBlock()
}

// _DeviceConfigurationRequestDataBlock is the data-structure of this message
type _DeviceConfigurationRequestDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ DeviceConfigurationRequestDataBlock = (*_DeviceConfigurationRequestDataBlock)(nil)

// NewDeviceConfigurationRequestDataBlock factory function for _DeviceConfigurationRequestDataBlock
func NewDeviceConfigurationRequestDataBlock(communicationChannelId uint8, sequenceCounter uint8) *_DeviceConfigurationRequestDataBlock {
	return &_DeviceConfigurationRequestDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequestDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DeviceConfigurationRequestDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequestDataBlock(structType any) DeviceConfigurationRequestDataBlock {
	if casted, ok := structType.(DeviceConfigurationRequestDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequestDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequestDataBlock) GetTypeName() string {
	return "DeviceConfigurationRequestDataBlock"
}

func (m *_DeviceConfigurationRequestDataBlock) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DeviceConfigurationRequestDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceConfigurationRequestDataBlockParse(ctx context.Context, theBytes []byte) (DeviceConfigurationRequestDataBlock, error) {
	return DeviceConfigurationRequestDataBlockParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceConfigurationRequestDataBlockParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationRequestDataBlock, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationRequestDataBlock, error) {
		return DeviceConfigurationRequestDataBlockParseWithBuffer(ctx, readBuffer)
	}
}

func DeviceConfigurationRequestDataBlockParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationRequestDataBlock, error) {
	v, err := (&_DeviceConfigurationRequestDataBlock{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DeviceConfigurationRequestDataBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequestDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	sequenceCounter, err := ReadSimpleField(ctx, "sequenceCounter", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceCounter' field"))
	}
	m.SequenceCounter = sequenceCounter

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequestDataBlock")
	}

	return m, nil
}

func (m *_DeviceConfigurationRequestDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationRequestDataBlock) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DeviceConfigurationRequestDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequestDataBlock")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sequenceCounter", m.GetSequenceCounter(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sequenceCounter' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if popErr := writeBuffer.PopContext("DeviceConfigurationRequestDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequestDataBlock")
	}
	return nil
}

func (m *_DeviceConfigurationRequestDataBlock) IsDeviceConfigurationRequestDataBlock() {}

func (m *_DeviceConfigurationRequestDataBlock) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeviceConfigurationRequestDataBlock) deepCopy() *_DeviceConfigurationRequestDataBlock {
	if m == nil {
		return nil
	}
	_DeviceConfigurationRequestDataBlockCopy := &_DeviceConfigurationRequestDataBlock{
		m.CommunicationChannelId,
		m.SequenceCounter,
		m.reservedField0,
	}
	return _DeviceConfigurationRequestDataBlockCopy
}

func (m *_DeviceConfigurationRequestDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
