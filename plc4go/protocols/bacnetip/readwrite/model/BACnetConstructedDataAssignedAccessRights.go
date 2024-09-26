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

// BACnetConstructedDataAssignedAccessRights is the corresponding interface of BACnetConstructedDataAssignedAccessRights
type BACnetConstructedDataAssignedAccessRights interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAssignedAccessRights returns AssignedAccessRights (property field)
	GetAssignedAccessRights() []BACnetAssignedAccessRights
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAssignedAccessRights is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAssignedAccessRights()
}

// _BACnetConstructedDataAssignedAccessRights is the data-structure of this message
type _BACnetConstructedDataAssignedAccessRights struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	AssignedAccessRights []BACnetAssignedAccessRights
}

var _ BACnetConstructedDataAssignedAccessRights = (*_BACnetConstructedDataAssignedAccessRights)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAssignedAccessRights)(nil)

// NewBACnetConstructedDataAssignedAccessRights factory function for _BACnetConstructedDataAssignedAccessRights
func NewBACnetConstructedDataAssignedAccessRights(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, assignedAccessRights []BACnetAssignedAccessRights, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAssignedAccessRights {
	_result := &_BACnetConstructedDataAssignedAccessRights{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AssignedAccessRights:          assignedAccessRights,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAssignedAccessRights) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAssignedAccessRights) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ASSIGNED_ACCESS_RIGHTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAssignedAccessRights) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAssignedAccessRights) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAssignedAccessRights) GetAssignedAccessRights() []BACnetAssignedAccessRights {
	return m.AssignedAccessRights
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAssignedAccessRights) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAssignedAccessRights(structType any) BACnetConstructedDataAssignedAccessRights {
	if casted, ok := structType.(BACnetConstructedDataAssignedAccessRights); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAssignedAccessRights); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAssignedAccessRights) GetTypeName() string {
	return "BACnetConstructedDataAssignedAccessRights"
}

func (m *_BACnetConstructedDataAssignedAccessRights) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AssignedAccessRights) > 0 {
		for _, element := range m.AssignedAccessRights {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAssignedAccessRights) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAssignedAccessRights) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAssignedAccessRights BACnetConstructedDataAssignedAccessRights, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAssignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAssignedAccessRights")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	assignedAccessRights, err := ReadTerminatedArrayField[BACnetAssignedAccessRights](ctx, "assignedAccessRights", ReadComplex[BACnetAssignedAccessRights](BACnetAssignedAccessRightsParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedAccessRights' field"))
	}
	m.AssignedAccessRights = assignedAccessRights

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAssignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAssignedAccessRights")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAssignedAccessRights) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAssignedAccessRights) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAssignedAccessRights"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAssignedAccessRights")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "assignedAccessRights", m.GetAssignedAccessRights(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'assignedAccessRights' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAssignedAccessRights"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAssignedAccessRights")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAssignedAccessRights) IsBACnetConstructedDataAssignedAccessRights() {}

func (m *_BACnetConstructedDataAssignedAccessRights) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAssignedAccessRights) deepCopy() *_BACnetConstructedDataAssignedAccessRights {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAssignedAccessRightsCopy := &_BACnetConstructedDataAssignedAccessRights{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetAssignedAccessRights, BACnetAssignedAccessRights](m.AssignedAccessRights),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAssignedAccessRightsCopy
}

func (m *_BACnetConstructedDataAssignedAccessRights) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
