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

// BACnetCalendarEntryDate is the corresponding interface of BACnetCalendarEntryDate
type BACnetCalendarEntryDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetCalendarEntry
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetContextTagDate
	// IsBACnetCalendarEntryDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntryDate()
}

// _BACnetCalendarEntryDate is the data-structure of this message
type _BACnetCalendarEntryDate struct {
	BACnetCalendarEntryContract
	DateValue BACnetContextTagDate
}

var _ BACnetCalendarEntryDate = (*_BACnetCalendarEntryDate)(nil)
var _ BACnetCalendarEntryRequirements = (*_BACnetCalendarEntryDate)(nil)

// NewBACnetCalendarEntryDate factory function for _BACnetCalendarEntryDate
func NewBACnetCalendarEntryDate(peekedTagHeader BACnetTagHeader, dateValue BACnetContextTagDate) *_BACnetCalendarEntryDate {
	if dateValue == nil {
		panic("dateValue of type BACnetContextTagDate for BACnetCalendarEntryDate must not be nil")
	}
	_result := &_BACnetCalendarEntryDate{
		BACnetCalendarEntryContract: NewBACnetCalendarEntry(peekedTagHeader),
		DateValue:                   dateValue,
	}
	_result.BACnetCalendarEntryContract.(*_BACnetCalendarEntry)._SubType = _result
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

func (m *_BACnetCalendarEntryDate) GetParent() BACnetCalendarEntryContract {
	return m.BACnetCalendarEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryDate) GetDateValue() BACnetContextTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryDate(structType any) BACnetCalendarEntryDate {
	if casted, ok := structType.(BACnetCalendarEntryDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryDate) GetTypeName() string {
	return "BACnetCalendarEntryDate"
}

func (m *_BACnetCalendarEntryDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).getLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetCalendarEntryDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetCalendarEntry) (__bACnetCalendarEntryDate BACnetCalendarEntryDate, err error) {
	m.BACnetCalendarEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetContextTagDate](ctx, "dateValue", ReadComplex[BACnetContextTagDate](BACnetContextTagParseWithBufferProducer[BACnetContextTagDate]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_DATE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryDate")
	}

	return m, nil
}

func (m *_BACnetCalendarEntryDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryDate")
		}

		if err := WriteSimpleField[BACnetContextTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetContextTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryDate")
		}
		return nil
	}
	return m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryDate) IsBACnetCalendarEntryDate() {}

func (m *_BACnetCalendarEntryDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCalendarEntryDate) deepCopy() *_BACnetCalendarEntryDate {
	if m == nil {
		return nil
	}
	_BACnetCalendarEntryDateCopy := &_BACnetCalendarEntryDate{
		m.BACnetCalendarEntryContract.DeepCopy().(BACnetCalendarEntryContract),
		m.DateValue.DeepCopy().(BACnetContextTagDate),
	}
	m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry)._SubType = m
	return _BACnetCalendarEntryDateCopy
}

func (m *_BACnetCalendarEntryDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
