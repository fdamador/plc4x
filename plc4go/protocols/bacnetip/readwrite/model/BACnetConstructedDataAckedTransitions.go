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

// BACnetConstructedDataAckedTransitions is the corresponding interface of BACnetConstructedDataAckedTransitions
type BACnetConstructedDataAckedTransitions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAckedTransitions returns AckedTransitions (property field)
	GetAckedTransitions() BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTransitionBitsTagged
	// IsBACnetConstructedDataAckedTransitions is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAckedTransitions()
}

// _BACnetConstructedDataAckedTransitions is the data-structure of this message
type _BACnetConstructedDataAckedTransitions struct {
	BACnetConstructedDataContract
	AckedTransitions BACnetEventTransitionBitsTagged
}

var _ BACnetConstructedDataAckedTransitions = (*_BACnetConstructedDataAckedTransitions)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAckedTransitions)(nil)

// NewBACnetConstructedDataAckedTransitions factory function for _BACnetConstructedDataAckedTransitions
func NewBACnetConstructedDataAckedTransitions(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ackedTransitions BACnetEventTransitionBitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAckedTransitions {
	if ackedTransitions == nil {
		panic("ackedTransitions of type BACnetEventTransitionBitsTagged for BACnetConstructedDataAckedTransitions must not be nil")
	}
	_result := &_BACnetConstructedDataAckedTransitions{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AckedTransitions:              ackedTransitions,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAckedTransitions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACKED_TRANSITIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetAckedTransitions() BACnetEventTransitionBitsTagged {
	return m.AckedTransitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetActualValue() BACnetEventTransitionBitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventTransitionBitsTagged(m.GetAckedTransitions())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAckedTransitions(structType any) BACnetConstructedDataAckedTransitions {
	if casted, ok := structType.(BACnetConstructedDataAckedTransitions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckedTransitions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAckedTransitions) GetTypeName() string {
	return "BACnetConstructedDataAckedTransitions"
}

func (m *_BACnetConstructedDataAckedTransitions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ackedTransitions)
	lengthInBits += m.AckedTransitions.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAckedTransitions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAckedTransitions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAckedTransitions BACnetConstructedDataAckedTransitions, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckedTransitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAckedTransitions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ackedTransitions, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "ackedTransitions", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackedTransitions' field"))
	}
	m.AckedTransitions = ackedTransitions

	actualValue, err := ReadVirtualField[BACnetEventTransitionBitsTagged](ctx, "actualValue", (*BACnetEventTransitionBitsTagged)(nil), ackedTransitions)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckedTransitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAckedTransitions")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAckedTransitions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAckedTransitions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckedTransitions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAckedTransitions")
		}

		if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "ackedTransitions", m.GetAckedTransitions(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ackedTransitions' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckedTransitions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAckedTransitions")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAckedTransitions) IsBACnetConstructedDataAckedTransitions() {}

func (m *_BACnetConstructedDataAckedTransitions) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAckedTransitions) deepCopy() *_BACnetConstructedDataAckedTransitions {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAckedTransitionsCopy := &_BACnetConstructedDataAckedTransitions{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.AckedTransitions.DeepCopy().(BACnetEventTransitionBitsTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAckedTransitionsCopy
}

func (m *_BACnetConstructedDataAckedTransitions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
