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

// ApduDataExtPropertyValueWrite is the corresponding interface of ApduDataExtPropertyValueWrite
type ApduDataExtPropertyValueWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint16
	// GetData returns Data (property field)
	GetData() []byte
	// IsApduDataExtPropertyValueWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtPropertyValueWrite()
}

// _ApduDataExtPropertyValueWrite is the data-structure of this message
type _ApduDataExtPropertyValueWrite struct {
	ApduDataExtContract
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
	Data        []byte
}

var _ ApduDataExtPropertyValueWrite = (*_ApduDataExtPropertyValueWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtPropertyValueWrite)(nil)

// NewApduDataExtPropertyValueWrite factory function for _ApduDataExtPropertyValueWrite
func NewApduDataExtPropertyValueWrite(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []byte, length uint8) *_ApduDataExtPropertyValueWrite {
	_result := &_ApduDataExtPropertyValueWrite{
		ApduDataExtContract: NewApduDataExt(length),
		ObjectIndex:         objectIndex,
		PropertyId:          propertyId,
		Count:               count,
		Index:               index,
		Data:                data,
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyValueWrite) GetExtApciType() uint8 {
	return 0x17
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyValueWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyValueWrite) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyValueWrite) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyValueWrite) GetCount() uint8 {
	return m.Count
}

func (m *_ApduDataExtPropertyValueWrite) GetIndex() uint16 {
	return m.Index
}

func (m *_ApduDataExtPropertyValueWrite) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyValueWrite(structType any) ApduDataExtPropertyValueWrite {
	if casted, ok := structType.(ApduDataExtPropertyValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtPropertyValueWrite"
}

func (m *_ApduDataExtPropertyValueWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataExtPropertyValueWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtPropertyValueWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtPropertyValueWrite ApduDataExtPropertyValueWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIndex, err := ReadSimpleField(ctx, "objectIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIndex' field"))
	}
	m.ObjectIndex = objectIndex

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	count, err := ReadSimpleField(ctx, "count", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}
	m.Index = index

	data, err := readBuffer.ReadByteArray("data", int(int32(length)-int32(int32(5))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyValueWrite")
	}

	return m, nil
}

func (m *_ApduDataExtPropertyValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyValueWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyValueWrite")
		}

		if err := WriteSimpleField[uint8](ctx, "objectIndex", m.GetObjectIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIndex' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "count", m.GetCount(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if err := WriteSimpleField[uint16](ctx, "index", m.GetIndex(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'index' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyValueWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyValueWrite) IsApduDataExtPropertyValueWrite() {}

func (m *_ApduDataExtPropertyValueWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtPropertyValueWrite) deepCopy() *_ApduDataExtPropertyValueWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtPropertyValueWriteCopy := &_ApduDataExtPropertyValueWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
		m.ObjectIndex,
		m.PropertyId,
		m.Count,
		m.Index,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtPropertyValueWriteCopy
}

func (m *_ApduDataExtPropertyValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
