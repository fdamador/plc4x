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

// BACnetConstructedDataAttemptedSamples is the corresponding interface of BACnetConstructedDataAttemptedSamples
type BACnetConstructedDataAttemptedSamples interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAttemptedSamples returns AttemptedSamples (property field)
	GetAttemptedSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAttemptedSamples is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAttemptedSamples()
}

// _BACnetConstructedDataAttemptedSamples is the data-structure of this message
type _BACnetConstructedDataAttemptedSamples struct {
	BACnetConstructedDataContract
	AttemptedSamples BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAttemptedSamples = (*_BACnetConstructedDataAttemptedSamples)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAttemptedSamples)(nil)

// NewBACnetConstructedDataAttemptedSamples factory function for _BACnetConstructedDataAttemptedSamples
func NewBACnetConstructedDataAttemptedSamples(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, attemptedSamples BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAttemptedSamples {
	if attemptedSamples == nil {
		panic("attemptedSamples of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAttemptedSamples must not be nil")
	}
	_result := &_BACnetConstructedDataAttemptedSamples{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AttemptedSamples:              attemptedSamples,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAttemptedSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ATTEMPTED_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetAttemptedSamples() BACnetApplicationTagUnsignedInteger {
	return m.AttemptedSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetAttemptedSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAttemptedSamples(structType any) BACnetConstructedDataAttemptedSamples {
	if casted, ok := structType.(BACnetConstructedDataAttemptedSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAttemptedSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAttemptedSamples) GetTypeName() string {
	return "BACnetConstructedDataAttemptedSamples"
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (attemptedSamples)
	lengthInBits += m.AttemptedSamples.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAttemptedSamples) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAttemptedSamples BACnetConstructedDataAttemptedSamples, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAttemptedSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAttemptedSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	attemptedSamples, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "attemptedSamples", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attemptedSamples' field"))
	}
	m.AttemptedSamples = attemptedSamples

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), attemptedSamples)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAttemptedSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAttemptedSamples")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAttemptedSamples) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAttemptedSamples) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAttemptedSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAttemptedSamples")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "attemptedSamples", m.GetAttemptedSamples(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'attemptedSamples' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAttemptedSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAttemptedSamples")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAttemptedSamples) IsBACnetConstructedDataAttemptedSamples() {}

func (m *_BACnetConstructedDataAttemptedSamples) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAttemptedSamples) deepCopy() *_BACnetConstructedDataAttemptedSamples {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAttemptedSamplesCopy := &_BACnetConstructedDataAttemptedSamples{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.AttemptedSamples.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAttemptedSamplesCopy
}

func (m *_BACnetConstructedDataAttemptedSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
