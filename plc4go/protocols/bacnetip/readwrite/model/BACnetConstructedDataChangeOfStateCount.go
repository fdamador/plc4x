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

// BACnetConstructedDataChangeOfStateCount is the corresponding interface of BACnetConstructedDataChangeOfStateCount
type BACnetConstructedDataChangeOfStateCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetChangeIfStateCount returns ChangeIfStateCount (property field)
	GetChangeIfStateCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataChangeOfStateCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataChangeOfStateCount()
}

// _BACnetConstructedDataChangeOfStateCount is the data-structure of this message
type _BACnetConstructedDataChangeOfStateCount struct {
	BACnetConstructedDataContract
	ChangeIfStateCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataChangeOfStateCount = (*_BACnetConstructedDataChangeOfStateCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataChangeOfStateCount)(nil)

// NewBACnetConstructedDataChangeOfStateCount factory function for _BACnetConstructedDataChangeOfStateCount
func NewBACnetConstructedDataChangeOfStateCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, changeIfStateCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChangeOfStateCount {
	if changeIfStateCount == nil {
		panic("changeIfStateCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataChangeOfStateCount must not be nil")
	}
	_result := &_BACnetConstructedDataChangeOfStateCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ChangeIfStateCount:            changeIfStateCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGE_OF_STATE_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetChangeIfStateCount() BACnetApplicationTagUnsignedInteger {
	return m.ChangeIfStateCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetChangeIfStateCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChangeOfStateCount(structType any) BACnetConstructedDataChangeOfStateCount {
	if casted, ok := structType.(BACnetConstructedDataChangeOfStateCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangeOfStateCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetTypeName() string {
	return "BACnetConstructedDataChangeOfStateCount"
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (changeIfStateCount)
	lengthInBits += m.ChangeIfStateCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChangeOfStateCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataChangeOfStateCount BACnetConstructedDataChangeOfStateCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangeOfStateCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangeOfStateCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	changeIfStateCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "changeIfStateCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'changeIfStateCount' field"))
	}
	m.ChangeIfStateCount = changeIfStateCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), changeIfStateCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangeOfStateCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangeOfStateCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangeOfStateCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangeOfStateCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "changeIfStateCount", m.GetChangeIfStateCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'changeIfStateCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangeOfStateCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangeOfStateCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChangeOfStateCount) IsBACnetConstructedDataChangeOfStateCount() {}

func (m *_BACnetConstructedDataChangeOfStateCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataChangeOfStateCount) deepCopy() *_BACnetConstructedDataChangeOfStateCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataChangeOfStateCountCopy := &_BACnetConstructedDataChangeOfStateCount{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.ChangeIfStateCount.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataChangeOfStateCountCopy
}

func (m *_BACnetConstructedDataChangeOfStateCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
