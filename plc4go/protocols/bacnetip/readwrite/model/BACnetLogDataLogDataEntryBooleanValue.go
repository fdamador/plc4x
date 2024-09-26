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

// BACnetLogDataLogDataEntryBooleanValue is the corresponding interface of BACnetLogDataLogDataEntryBooleanValue
type BACnetLogDataLogDataEntryBooleanValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetContextTagBoolean
	// IsBACnetLogDataLogDataEntryBooleanValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryBooleanValue()
}

// _BACnetLogDataLogDataEntryBooleanValue is the data-structure of this message
type _BACnetLogDataLogDataEntryBooleanValue struct {
	BACnetLogDataLogDataEntryContract
	BooleanValue BACnetContextTagBoolean
}

var _ BACnetLogDataLogDataEntryBooleanValue = (*_BACnetLogDataLogDataEntryBooleanValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryBooleanValue)(nil)

// NewBACnetLogDataLogDataEntryBooleanValue factory function for _BACnetLogDataLogDataEntryBooleanValue
func NewBACnetLogDataLogDataEntryBooleanValue(peekedTagHeader BACnetTagHeader, booleanValue BACnetContextTagBoolean) *_BACnetLogDataLogDataEntryBooleanValue {
	if booleanValue == nil {
		panic("booleanValue of type BACnetContextTagBoolean for BACnetLogDataLogDataEntryBooleanValue must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryBooleanValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		BooleanValue:                      booleanValue,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetBooleanValue() BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryBooleanValue(structType any) BACnetLogDataLogDataEntryBooleanValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryBooleanValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryBooleanValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryBooleanValue"
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryBooleanValue BACnetLogDataLogDataEntryBooleanValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryBooleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryBooleanValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "booleanValue", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryBooleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryBooleanValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryBooleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryBooleanValue")
		}

		if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryBooleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryBooleanValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) IsBACnetLogDataLogDataEntryBooleanValue() {}

func (m *_BACnetLogDataLogDataEntryBooleanValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) deepCopy() *_BACnetLogDataLogDataEntryBooleanValue {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryBooleanValueCopy := &_BACnetLogDataLogDataEntryBooleanValue{
		m.BACnetLogDataLogDataEntryContract.DeepCopy().(BACnetLogDataLogDataEntryContract),
		m.BooleanValue.DeepCopy().(BACnetContextTagBoolean),
	}
	m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryBooleanValueCopy
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
