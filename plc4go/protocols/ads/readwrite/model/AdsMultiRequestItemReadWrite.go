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

// AdsMultiRequestItemReadWrite is the corresponding interface of AdsMultiRequestItemReadWrite
type AdsMultiRequestItemReadWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemReadLength returns ItemReadLength (property field)
	GetItemReadLength() uint32
	// GetItemWriteLength returns ItemWriteLength (property field)
	GetItemWriteLength() uint32
	// IsAdsMultiRequestItemReadWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsMultiRequestItemReadWrite()
}

// _AdsMultiRequestItemReadWrite is the data-structure of this message
type _AdsMultiRequestItemReadWrite struct {
	AdsMultiRequestItemContract
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
	ItemWriteLength uint32
}

var _ AdsMultiRequestItemReadWrite = (*_AdsMultiRequestItemReadWrite)(nil)
var _ AdsMultiRequestItemRequirements = (*_AdsMultiRequestItemReadWrite)(nil)

// NewAdsMultiRequestItemReadWrite factory function for _AdsMultiRequestItemReadWrite
func NewAdsMultiRequestItemReadWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32, itemWriteLength uint32) *_AdsMultiRequestItemReadWrite {
	_result := &_AdsMultiRequestItemReadWrite{
		AdsMultiRequestItemContract: NewAdsMultiRequestItem(),
		ItemIndexGroup:              itemIndexGroup,
		ItemIndexOffset:             itemIndexOffset,
		ItemReadLength:              itemReadLength,
		ItemWriteLength:             itemWriteLength,
	}
	_result.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetIndexGroup() uint32 {
	return uint32(61570)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetParent() AdsMultiRequestItemContract {
	return m.AdsMultiRequestItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemReadWrite) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemReadWrite) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemReadWrite) GetItemReadLength() uint32 {
	return m.ItemReadLength
}

func (m *_AdsMultiRequestItemReadWrite) GetItemWriteLength() uint32 {
	return m.ItemWriteLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemReadWrite(structType any) AdsMultiRequestItemReadWrite {
	if casted, ok := structType.(AdsMultiRequestItemReadWrite); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemReadWrite); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemReadWrite) GetTypeName() string {
	return "AdsMultiRequestItemReadWrite"
}

func (m *_AdsMultiRequestItemReadWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).getLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemReadWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsMultiRequestItemReadWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsMultiRequestItem, indexGroup uint32) (__adsMultiRequestItemReadWrite AdsMultiRequestItemReadWrite, err error) {
	m.AdsMultiRequestItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemReadWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemReadWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndexGroup, err := ReadSimpleField(ctx, "itemIndexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexGroup' field"))
	}
	m.ItemIndexGroup = itemIndexGroup

	itemIndexOffset, err := ReadSimpleField(ctx, "itemIndexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexOffset' field"))
	}
	m.ItemIndexOffset = itemIndexOffset

	itemReadLength, err := ReadSimpleField(ctx, "itemReadLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemReadLength' field"))
	}
	m.ItemReadLength = itemReadLength

	itemWriteLength, err := ReadSimpleField(ctx, "itemWriteLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemWriteLength' field"))
	}
	m.ItemWriteLength = itemWriteLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemReadWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemReadWrite")
	}

	return m, nil
}

func (m *_AdsMultiRequestItemReadWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemReadWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemReadWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemReadWrite")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexGroup", m.GetItemIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexOffset", m.GetItemIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemReadLength", m.GetItemReadLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemReadLength' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemWriteLength", m.GetItemWriteLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemReadWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemReadWrite")
		}
		return nil
	}
	return m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemReadWrite) IsAdsMultiRequestItemReadWrite() {}

func (m *_AdsMultiRequestItemReadWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsMultiRequestItemReadWrite) deepCopy() *_AdsMultiRequestItemReadWrite {
	if m == nil {
		return nil
	}
	_AdsMultiRequestItemReadWriteCopy := &_AdsMultiRequestItemReadWrite{
		m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).deepCopy(),
		m.ItemIndexGroup,
		m.ItemIndexOffset,
		m.ItemReadLength,
		m.ItemWriteLength,
	}
	m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem)._SubType = m
	return _AdsMultiRequestItemReadWriteCopy
}

func (m *_AdsMultiRequestItemReadWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
