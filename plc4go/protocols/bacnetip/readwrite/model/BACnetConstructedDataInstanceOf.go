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

// BACnetConstructedDataInstanceOf is the corresponding interface of BACnetConstructedDataInstanceOf
type BACnetConstructedDataInstanceOf interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInstanceOf returns InstanceOf (property field)
	GetInstanceOf() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataInstanceOf is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInstanceOf()
}

// _BACnetConstructedDataInstanceOf is the data-structure of this message
type _BACnetConstructedDataInstanceOf struct {
	BACnetConstructedDataContract
	InstanceOf BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataInstanceOf = (*_BACnetConstructedDataInstanceOf)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInstanceOf)(nil)

// NewBACnetConstructedDataInstanceOf factory function for _BACnetConstructedDataInstanceOf
func NewBACnetConstructedDataInstanceOf(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, instanceOf BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInstanceOf {
	if instanceOf == nil {
		panic("instanceOf of type BACnetApplicationTagCharacterString for BACnetConstructedDataInstanceOf must not be nil")
	}
	_result := &_BACnetConstructedDataInstanceOf{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InstanceOf:                    instanceOf,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInstanceOf) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANCE_OF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetInstanceOf() BACnetApplicationTagCharacterString {
	return m.InstanceOf
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetInstanceOf())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInstanceOf(structType any) BACnetConstructedDataInstanceOf {
	if casted, ok := structType.(BACnetConstructedDataInstanceOf); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstanceOf); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInstanceOf) GetTypeName() string {
	return "BACnetConstructedDataInstanceOf"
}

func (m *_BACnetConstructedDataInstanceOf) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (instanceOf)
	lengthInBits += m.InstanceOf.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInstanceOf) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInstanceOf) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInstanceOf BACnetConstructedDataInstanceOf, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstanceOf"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstanceOf")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	instanceOf, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "instanceOf", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceOf' field"))
	}
	m.InstanceOf = instanceOf

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), instanceOf)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstanceOf"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstanceOf")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInstanceOf) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInstanceOf) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstanceOf"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstanceOf")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "instanceOf", m.GetInstanceOf(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceOf' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstanceOf"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstanceOf")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInstanceOf) IsBACnetConstructedDataInstanceOf() {}

func (m *_BACnetConstructedDataInstanceOf) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataInstanceOf) deepCopy() *_BACnetConstructedDataInstanceOf {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataInstanceOfCopy := &_BACnetConstructedDataInstanceOf{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.InstanceOf.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataInstanceOfCopy
}

func (m *_BACnetConstructedDataInstanceOf) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
