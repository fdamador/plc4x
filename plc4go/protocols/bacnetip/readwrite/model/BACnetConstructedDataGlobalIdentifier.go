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

// BACnetConstructedDataGlobalIdentifier is the corresponding interface of BACnetConstructedDataGlobalIdentifier
type BACnetConstructedDataGlobalIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetGlobalIdentifier returns GlobalIdentifier (property field)
	GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataGlobalIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGlobalIdentifier()
}

// _BACnetConstructedDataGlobalIdentifier is the data-structure of this message
type _BACnetConstructedDataGlobalIdentifier struct {
	BACnetConstructedDataContract
	GlobalIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataGlobalIdentifier = (*_BACnetConstructedDataGlobalIdentifier)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGlobalIdentifier)(nil)

// NewBACnetConstructedDataGlobalIdentifier factory function for _BACnetConstructedDataGlobalIdentifier
func NewBACnetConstructedDataGlobalIdentifier(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, globalIdentifier BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGlobalIdentifier {
	if globalIdentifier == nil {
		panic("globalIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataGlobalIdentifier must not be nil")
	}
	_result := &_BACnetConstructedDataGlobalIdentifier{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		GlobalIdentifier:              globalIdentifier,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GLOBAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.GlobalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetGlobalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGlobalIdentifier(structType any) BACnetConstructedDataGlobalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataGlobalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataGlobalIdentifier"
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (globalIdentifier)
	lengthInBits += m.GlobalIdentifier.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGlobalIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGlobalIdentifier BACnetConstructedDataGlobalIdentifier, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGlobalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	globalIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "globalIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'globalIdentifier' field"))
	}
	m.GlobalIdentifier = globalIdentifier

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), globalIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGlobalIdentifier")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGlobalIdentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "globalIdentifier", m.GetGlobalIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'globalIdentifier' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGlobalIdentifier")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGlobalIdentifier) IsBACnetConstructedDataGlobalIdentifier() {}

func (m *_BACnetConstructedDataGlobalIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGlobalIdentifier) deepCopy() *_BACnetConstructedDataGlobalIdentifier {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGlobalIdentifierCopy := &_BACnetConstructedDataGlobalIdentifier{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.GlobalIdentifier.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGlobalIdentifierCopy
}

func (m *_BACnetConstructedDataGlobalIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
