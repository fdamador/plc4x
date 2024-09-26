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

// BACnetPriorityValueDate is the corresponding interface of BACnetPriorityValueDate
type BACnetPriorityValueDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// IsBACnetPriorityValueDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueDate()
}

// _BACnetPriorityValueDate is the data-structure of this message
type _BACnetPriorityValueDate struct {
	BACnetPriorityValueContract
	DateValue BACnetApplicationTagDate
}

var _ BACnetPriorityValueDate = (*_BACnetPriorityValueDate)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueDate)(nil)

// NewBACnetPriorityValueDate factory function for _BACnetPriorityValueDate
func NewBACnetPriorityValueDate(peekedTagHeader BACnetTagHeader, dateValue BACnetApplicationTagDate, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueDate {
	if dateValue == nil {
		panic("dateValue of type BACnetApplicationTagDate for BACnetPriorityValueDate must not be nil")
	}
	_result := &_BACnetPriorityValueDate{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		DateValue:                   dateValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
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

func (m *_BACnetPriorityValueDate) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueDate(structType any) BACnetPriorityValueDate {
	if casted, ok := structType.(BACnetPriorityValueDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueDate) GetTypeName() string {
	return "BACnetPriorityValueDate"
}

func (m *_BACnetPriorityValueDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueDate BACnetPriorityValueDate, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueDate")
	}

	return m, nil
}

func (m *_BACnetPriorityValueDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueDate")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueDate")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueDate) IsBACnetPriorityValueDate() {}

func (m *_BACnetPriorityValueDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueDate) deepCopy() *_BACnetPriorityValueDate {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueDateCopy := &_BACnetPriorityValueDate{
		m.BACnetPriorityValueContract.DeepCopy().(BACnetPriorityValueContract),
		m.DateValue.DeepCopy().(BACnetApplicationTagDate),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueDateCopy
}

func (m *_BACnetPriorityValueDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
