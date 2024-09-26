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

// S7PayloadUserDataItemCyclicServicesPush is the corresponding interface of S7PayloadUserDataItemCyclicServicesPush
type S7PayloadUserDataItemCyclicServicesPush interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetItemsCount returns ItemsCount (property field)
	GetItemsCount() uint16
	// GetItems returns Items (property field)
	GetItems() []AssociatedValueType
	// IsS7PayloadUserDataItemCyclicServicesPush is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesPush()
}

// _S7PayloadUserDataItemCyclicServicesPush is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesPush struct {
	S7PayloadUserDataItemContract
	ItemsCount uint16
	Items      []AssociatedValueType
}

var _ S7PayloadUserDataItemCyclicServicesPush = (*_S7PayloadUserDataItemCyclicServicesPush)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesPush)(nil)

// NewS7PayloadUserDataItemCyclicServicesPush factory function for _S7PayloadUserDataItemCyclicServicesPush
func NewS7PayloadUserDataItemCyclicServicesPush(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, itemsCount uint16, items []AssociatedValueType) *_S7PayloadUserDataItemCyclicServicesPush {
	_result := &_S7PayloadUserDataItemCyclicServicesPush{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		ItemsCount:                    itemsCount,
		Items:                         items,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetItemsCount() uint16 {
	return m.ItemsCount
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetItems() []AssociatedValueType {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesPush(structType any) S7PayloadUserDataItemCyclicServicesPush {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesPush); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesPush); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesPush"
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (itemsCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesPush S7PayloadUserDataItemCyclicServicesPush, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesPush"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesPush")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemsCount, err := ReadSimpleField(ctx, "itemsCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsCount' field"))
	}
	m.ItemsCount = itemsCount

	items, err := ReadCountArrayField[AssociatedValueType](ctx, "items", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer), uint64(itemsCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesPush"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesPush")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesPush"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesPush")
		}

		if err := WriteSimpleField[uint16](ctx, "itemsCount", m.GetItemsCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesPush"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesPush")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) IsS7PayloadUserDataItemCyclicServicesPush() {}

func (m *_S7PayloadUserDataItemCyclicServicesPush) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) deepCopy() *_S7PayloadUserDataItemCyclicServicesPush {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCyclicServicesPushCopy := &_S7PayloadUserDataItemCyclicServicesPush{
		m.S7PayloadUserDataItemContract.DeepCopy().(S7PayloadUserDataItemContract),
		m.ItemsCount,
		utils.DeepCopySlice[AssociatedValueType, AssociatedValueType](m.Items),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCyclicServicesPushCopy
}

func (m *_S7PayloadUserDataItemCyclicServicesPush) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
