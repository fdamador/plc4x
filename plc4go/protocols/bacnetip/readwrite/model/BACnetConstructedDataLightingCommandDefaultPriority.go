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

// BACnetConstructedDataLightingCommandDefaultPriority is the corresponding interface of BACnetConstructedDataLightingCommandDefaultPriority
type BACnetConstructedDataLightingCommandDefaultPriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLightingCommandDefaultPriority returns LightingCommandDefaultPriority (property field)
	GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataLightingCommandDefaultPriority is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLightingCommandDefaultPriority()
}

// _BACnetConstructedDataLightingCommandDefaultPriority is the data-structure of this message
type _BACnetConstructedDataLightingCommandDefaultPriority struct {
	BACnetConstructedDataContract
	LightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataLightingCommandDefaultPriority = (*_BACnetConstructedDataLightingCommandDefaultPriority)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLightingCommandDefaultPriority)(nil)

// NewBACnetConstructedDataLightingCommandDefaultPriority factory function for _BACnetConstructedDataLightingCommandDefaultPriority
func NewBACnetConstructedDataLightingCommandDefaultPriority(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lightingCommandDefaultPriority BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommandDefaultPriority {
	if lightingCommandDefaultPriority == nil {
		panic("lightingCommandDefaultPriority of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataLightingCommandDefaultPriority must not be nil")
	}
	_result := &_BACnetConstructedDataLightingCommandDefaultPriority{
		BACnetConstructedDataContract:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LightingCommandDefaultPriority: lightingCommandDefaultPriority,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND_DEFAULT_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLightingCommandDefaultPriority() BACnetApplicationTagUnsignedInteger {
	return m.LightingCommandDefaultPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLightingCommandDefaultPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommandDefaultPriority(structType any) BACnetConstructedDataLightingCommandDefaultPriority {
	if casted, ok := structType.(BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommandDefaultPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetTypeName() string {
	return "BACnetConstructedDataLightingCommandDefaultPriority"
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lightingCommandDefaultPriority)
	lengthInBits += m.LightingCommandDefaultPriority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLightingCommandDefaultPriority BACnetConstructedDataLightingCommandDefaultPriority, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommandDefaultPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommandDefaultPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightingCommandDefaultPriority, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lightingCommandDefaultPriority", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingCommandDefaultPriority' field"))
	}
	m.LightingCommandDefaultPriority = lightingCommandDefaultPriority

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lightingCommandDefaultPriority)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommandDefaultPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommandDefaultPriority")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommandDefaultPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommandDefaultPriority")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lightingCommandDefaultPriority", m.GetLightingCommandDefaultPriority(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightingCommandDefaultPriority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommandDefaultPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommandDefaultPriority")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) IsBACnetConstructedDataLightingCommandDefaultPriority() {
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) deepCopy() *_BACnetConstructedDataLightingCommandDefaultPriority {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLightingCommandDefaultPriorityCopy := &_BACnetConstructedDataLightingCommandDefaultPriority{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.LightingCommandDefaultPriority.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLightingCommandDefaultPriorityCopy
}

func (m *_BACnetConstructedDataLightingCommandDefaultPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
