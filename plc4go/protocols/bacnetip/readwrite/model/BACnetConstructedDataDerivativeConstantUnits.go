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

// BACnetConstructedDataDerivativeConstantUnits is the corresponding interface of BACnetConstructedDataDerivativeConstantUnits
type BACnetConstructedDataDerivativeConstantUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
	// IsBACnetConstructedDataDerivativeConstantUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDerivativeConstantUnits()
}

// _BACnetConstructedDataDerivativeConstantUnits is the data-structure of this message
type _BACnetConstructedDataDerivativeConstantUnits struct {
	BACnetConstructedDataContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetConstructedDataDerivativeConstantUnits = (*_BACnetConstructedDataDerivativeConstantUnits)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDerivativeConstantUnits)(nil)

// NewBACnetConstructedDataDerivativeConstantUnits factory function for _BACnetConstructedDataDerivativeConstantUnits
func NewBACnetConstructedDataDerivativeConstantUnits(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, units BACnetEngineeringUnitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDerivativeConstantUnits {
	if units == nil {
		panic("units of type BACnetEngineeringUnitsTagged for BACnetConstructedDataDerivativeConstantUnits must not be nil")
	}
	_result := &_BACnetConstructedDataDerivativeConstantUnits{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Units:                         units,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DERIVATIVE_CONSTANT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDerivativeConstantUnits(structType any) BACnetConstructedDataDerivativeConstantUnits {
	if casted, ok := structType.(BACnetConstructedDataDerivativeConstantUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDerivativeConstantUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetTypeName() string {
	return "BACnetConstructedDataDerivativeConstantUnits"
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDerivativeConstantUnits BACnetConstructedDataDerivativeConstantUnits, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDerivativeConstantUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDerivativeConstantUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	actualValue, err := ReadVirtualField[BACnetEngineeringUnitsTagged](ctx, "actualValue", (*BACnetEngineeringUnitsTagged)(nil), units)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDerivativeConstantUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDerivativeConstantUnits")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDerivativeConstantUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDerivativeConstantUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDerivativeConstantUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDerivativeConstantUnits")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) IsBACnetConstructedDataDerivativeConstantUnits() {
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) deepCopy() *_BACnetConstructedDataDerivativeConstantUnits {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDerivativeConstantUnitsCopy := &_BACnetConstructedDataDerivativeConstantUnits{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.Units.DeepCopy().(BACnetEngineeringUnitsTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDerivativeConstantUnitsCopy
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
