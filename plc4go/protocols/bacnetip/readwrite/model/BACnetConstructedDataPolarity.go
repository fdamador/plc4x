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

// BACnetConstructedDataPolarity is the corresponding interface of BACnetConstructedDataPolarity
type BACnetConstructedDataPolarity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPolarity returns Polarity (property field)
	GetPolarity() BACnetPolarityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPolarityTagged
	// IsBACnetConstructedDataPolarity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPolarity()
}

// _BACnetConstructedDataPolarity is the data-structure of this message
type _BACnetConstructedDataPolarity struct {
	BACnetConstructedDataContract
	Polarity BACnetPolarityTagged
}

var _ BACnetConstructedDataPolarity = (*_BACnetConstructedDataPolarity)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPolarity)(nil)

// NewBACnetConstructedDataPolarity factory function for _BACnetConstructedDataPolarity
func NewBACnetConstructedDataPolarity(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, polarity BACnetPolarityTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPolarity {
	if polarity == nil {
		panic("polarity of type BACnetPolarityTagged for BACnetConstructedDataPolarity must not be nil")
	}
	_result := &_BACnetConstructedDataPolarity{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Polarity:                      polarity,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPolarity) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_POLARITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPolarity) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetPolarity() BACnetPolarityTagged {
	return m.Polarity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPolarity) GetActualValue() BACnetPolarityTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetPolarityTagged(m.GetPolarity())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPolarity(structType any) BACnetConstructedDataPolarity {
	if casted, ok := structType.(BACnetConstructedDataPolarity); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPolarity); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPolarity) GetTypeName() string {
	return "BACnetConstructedDataPolarity"
}

func (m *_BACnetConstructedDataPolarity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (polarity)
	lengthInBits += m.Polarity.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPolarity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPolarity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPolarity BACnetConstructedDataPolarity, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPolarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPolarity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	polarity, err := ReadSimpleField[BACnetPolarityTagged](ctx, "polarity", ReadComplex[BACnetPolarityTagged](BACnetPolarityTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'polarity' field"))
	}
	m.Polarity = polarity

	actualValue, err := ReadVirtualField[BACnetPolarityTagged](ctx, "actualValue", (*BACnetPolarityTagged)(nil), polarity)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPolarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPolarity")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPolarity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPolarity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPolarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPolarity")
		}

		if err := WriteSimpleField[BACnetPolarityTagged](ctx, "polarity", m.GetPolarity(), WriteComplex[BACnetPolarityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'polarity' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPolarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPolarity")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPolarity) IsBACnetConstructedDataPolarity() {}

func (m *_BACnetConstructedDataPolarity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPolarity) deepCopy() *_BACnetConstructedDataPolarity {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPolarityCopy := &_BACnetConstructedDataPolarity{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.Polarity.DeepCopy().(BACnetPolarityTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPolarityCopy
}

func (m *_BACnetConstructedDataPolarity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
