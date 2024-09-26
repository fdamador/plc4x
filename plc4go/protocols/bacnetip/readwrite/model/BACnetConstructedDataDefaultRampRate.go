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

// BACnetConstructedDataDefaultRampRate is the corresponding interface of BACnetConstructedDataDefaultRampRate
type BACnetConstructedDataDefaultRampRate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDefaultRampRate returns DefaultRampRate (property field)
	GetDefaultRampRate() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataDefaultRampRate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDefaultRampRate()
}

// _BACnetConstructedDataDefaultRampRate is the data-structure of this message
type _BACnetConstructedDataDefaultRampRate struct {
	BACnetConstructedDataContract
	DefaultRampRate BACnetApplicationTagReal
}

var _ BACnetConstructedDataDefaultRampRate = (*_BACnetConstructedDataDefaultRampRate)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDefaultRampRate)(nil)

// NewBACnetConstructedDataDefaultRampRate factory function for _BACnetConstructedDataDefaultRampRate
func NewBACnetConstructedDataDefaultRampRate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, defaultRampRate BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultRampRate {
	if defaultRampRate == nil {
		panic("defaultRampRate of type BACnetApplicationTagReal for BACnetConstructedDataDefaultRampRate must not be nil")
	}
	_result := &_BACnetConstructedDataDefaultRampRate{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DefaultRampRate:               defaultRampRate,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultRampRate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_RAMP_RATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetDefaultRampRate() BACnetApplicationTagReal {
	return m.DefaultRampRate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetDefaultRampRate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultRampRate(structType any) BACnetConstructedDataDefaultRampRate {
	if casted, ok := structType.(BACnetConstructedDataDefaultRampRate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultRampRate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultRampRate) GetTypeName() string {
	return "BACnetConstructedDataDefaultRampRate"
}

func (m *_BACnetConstructedDataDefaultRampRate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (defaultRampRate)
	lengthInBits += m.DefaultRampRate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultRampRate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDefaultRampRate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDefaultRampRate BACnetConstructedDataDefaultRampRate, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultRampRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultRampRate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	defaultRampRate, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "defaultRampRate", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'defaultRampRate' field"))
	}
	m.DefaultRampRate = defaultRampRate

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), defaultRampRate)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultRampRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultRampRate")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDefaultRampRate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDefaultRampRate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultRampRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultRampRate")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "defaultRampRate", m.GetDefaultRampRate(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'defaultRampRate' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultRampRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultRampRate")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultRampRate) IsBACnetConstructedDataDefaultRampRate() {}

func (m *_BACnetConstructedDataDefaultRampRate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDefaultRampRate) deepCopy() *_BACnetConstructedDataDefaultRampRate {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDefaultRampRateCopy := &_BACnetConstructedDataDefaultRampRate{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DefaultRampRate.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDefaultRampRateCopy
}

func (m *_BACnetConstructedDataDefaultRampRate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
