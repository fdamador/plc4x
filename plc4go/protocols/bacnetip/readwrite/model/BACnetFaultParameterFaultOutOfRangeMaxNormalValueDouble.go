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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble()
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doubleValue BACnetApplicationTagDouble, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		DoubleValue: doubleValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = _result
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleCopy := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).deepCopy(),
		m.DoubleValue.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
