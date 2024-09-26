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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetNumericValue returns NumericValue (property field)
	GetNumericValue() BACnetContextTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric()
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric struct {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	NumericValue BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)(nil)
var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numericValue BACnetContextTagUnsignedInteger, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if numericValue == nil {
		panic("numericValue of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
		NumericValue: numericValue,
	}
	_result.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = _result
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

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract {
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetNumericValue() BACnetContextTagUnsignedInteger {
	return m.NumericValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).getLengthInBits(ctx))

	// Simple field (numericValue)
	lengthInBits += m.NumericValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, tagNumber uint8) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric, err error) {
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numericValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "numericValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numericValue' field"))
	}
	m.NumericValue = numericValue

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "numericValue", m.GetNumericValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'numericValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.DeepCopy().(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract),
		m.NumericValue.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
