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

// BACnetLogRecordLogDatumAnyValue is the corresponding interface of BACnetLogRecordLogDatumAnyValue
type BACnetLogRecordLogDatumAnyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetAnyValue returns AnyValue (property field)
	GetAnyValue() BACnetConstructedData
	// IsBACnetLogRecordLogDatumAnyValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumAnyValue()
}

// _BACnetLogRecordLogDatumAnyValue is the data-structure of this message
type _BACnetLogRecordLogDatumAnyValue struct {
	BACnetLogRecordLogDatumContract
	AnyValue BACnetConstructedData
}

var _ BACnetLogRecordLogDatumAnyValue = (*_BACnetLogRecordLogDatumAnyValue)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumAnyValue)(nil)

// NewBACnetLogRecordLogDatumAnyValue factory function for _BACnetLogRecordLogDatumAnyValue
func NewBACnetLogRecordLogDatumAnyValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, anyValue BACnetConstructedData, tagNumber uint8) *_BACnetLogRecordLogDatumAnyValue {
	_result := &_BACnetLogRecordLogDatumAnyValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		AnyValue:                        anyValue,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
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

func (m *_BACnetLogRecordLogDatumAnyValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumAnyValue) GetAnyValue() BACnetConstructedData {
	return m.AnyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumAnyValue(structType any) BACnetLogRecordLogDatumAnyValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumAnyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumAnyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumAnyValue"
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Optional Field (anyValue)
	if m.AnyValue != nil {
		lengthInBits += m.AnyValue.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumAnyValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumAnyValue BACnetLogRecordLogDatumAnyValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumAnyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumAnyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var anyValue BACnetConstructedData
	_anyValue, err := ReadOptionalField[BACnetConstructedData](ctx, "anyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(10)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'anyValue' field"))
	}
	if _anyValue != nil {
		anyValue = *_anyValue
		m.AnyValue = anyValue
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumAnyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumAnyValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumAnyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumAnyValue")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "anyValue", GetRef(m.GetAnyValue()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'anyValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumAnyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumAnyValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumAnyValue) IsBACnetLogRecordLogDatumAnyValue() {}

func (m *_BACnetLogRecordLogDatumAnyValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumAnyValue) deepCopy() *_BACnetLogRecordLogDatumAnyValue {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumAnyValueCopy := &_BACnetLogRecordLogDatumAnyValue{
		m.BACnetLogRecordLogDatumContract.DeepCopy().(BACnetLogRecordLogDatumContract),
		m.AnyValue.DeepCopy().(BACnetConstructedData),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumAnyValueCopy
}

func (m *_BACnetLogRecordLogDatumAnyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
