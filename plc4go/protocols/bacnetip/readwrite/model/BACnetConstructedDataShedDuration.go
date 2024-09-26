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

// BACnetConstructedDataShedDuration is the corresponding interface of BACnetConstructedDataShedDuration
type BACnetConstructedDataShedDuration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetShedDuration returns ShedDuration (property field)
	GetShedDuration() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataShedDuration is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataShedDuration()
}

// _BACnetConstructedDataShedDuration is the data-structure of this message
type _BACnetConstructedDataShedDuration struct {
	BACnetConstructedDataContract
	ShedDuration BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataShedDuration = (*_BACnetConstructedDataShedDuration)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataShedDuration)(nil)

// NewBACnetConstructedDataShedDuration factory function for _BACnetConstructedDataShedDuration
func NewBACnetConstructedDataShedDuration(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, shedDuration BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataShedDuration {
	if shedDuration == nil {
		panic("shedDuration of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataShedDuration must not be nil")
	}
	_result := &_BACnetConstructedDataShedDuration{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ShedDuration:                  shedDuration,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataShedDuration) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SHED_DURATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataShedDuration) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetShedDuration() BACnetApplicationTagUnsignedInteger {
	return m.ShedDuration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetShedDuration())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataShedDuration(structType any) BACnetConstructedDataShedDuration {
	if casted, ok := structType.(BACnetConstructedDataShedDuration); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataShedDuration); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataShedDuration) GetTypeName() string {
	return "BACnetConstructedDataShedDuration"
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (shedDuration)
	lengthInBits += m.ShedDuration.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataShedDuration) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataShedDuration BACnetConstructedDataShedDuration, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataShedDuration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataShedDuration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	shedDuration, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "shedDuration", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'shedDuration' field"))
	}
	m.ShedDuration = shedDuration

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), shedDuration)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataShedDuration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataShedDuration")
	}

	return m, nil
}

func (m *_BACnetConstructedDataShedDuration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataShedDuration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataShedDuration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataShedDuration")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "shedDuration", m.GetShedDuration(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'shedDuration' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataShedDuration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataShedDuration")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataShedDuration) IsBACnetConstructedDataShedDuration() {}

func (m *_BACnetConstructedDataShedDuration) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataShedDuration) deepCopy() *_BACnetConstructedDataShedDuration {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataShedDurationCopy := &_BACnetConstructedDataShedDuration{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.ShedDuration.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataShedDurationCopy
}

func (m *_BACnetConstructedDataShedDuration) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
