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

// BACnetConstructedDataLastCredentialRemovedTime is the corresponding interface of BACnetConstructedDataLastCredentialRemovedTime
type BACnetConstructedDataLastCredentialRemovedTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastCredentialRemovedTime returns LastCredentialRemovedTime (property field)
	GetLastCredentialRemovedTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataLastCredentialRemovedTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastCredentialRemovedTime()
}

// _BACnetConstructedDataLastCredentialRemovedTime is the data-structure of this message
type _BACnetConstructedDataLastCredentialRemovedTime struct {
	BACnetConstructedDataContract
	LastCredentialRemovedTime BACnetDateTime
}

var _ BACnetConstructedDataLastCredentialRemovedTime = (*_BACnetConstructedDataLastCredentialRemovedTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastCredentialRemovedTime)(nil)

// NewBACnetConstructedDataLastCredentialRemovedTime factory function for _BACnetConstructedDataLastCredentialRemovedTime
func NewBACnetConstructedDataLastCredentialRemovedTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastCredentialRemovedTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCredentialRemovedTime {
	if lastCredentialRemovedTime == nil {
		panic("lastCredentialRemovedTime of type BACnetDateTime for BACnetConstructedDataLastCredentialRemovedTime must not be nil")
	}
	_result := &_BACnetConstructedDataLastCredentialRemovedTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastCredentialRemovedTime:     lastCredentialRemovedTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLastCredentialRemovedTime() BACnetDateTime {
	return m.LastCredentialRemovedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetLastCredentialRemovedTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialRemovedTime(structType any) BACnetConstructedDataLastCredentialRemovedTime {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialRemovedTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemovedTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemovedTime"
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastCredentialRemovedTime)
	lengthInBits += m.LastCredentialRemovedTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastCredentialRemovedTime BACnetConstructedDataLastCredentialRemovedTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemovedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialRemovedTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastCredentialRemovedTime, err := ReadSimpleField[BACnetDateTime](ctx, "lastCredentialRemovedTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCredentialRemovedTime' field"))
	}
	m.LastCredentialRemovedTime = lastCredentialRemovedTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), lastCredentialRemovedTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemovedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialRemovedTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemovedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialRemovedTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "lastCredentialRemovedTime", m.GetLastCredentialRemovedTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCredentialRemovedTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemovedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialRemovedTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) IsBACnetConstructedDataLastCredentialRemovedTime() {
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) deepCopy() *_BACnetConstructedDataLastCredentialRemovedTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastCredentialRemovedTimeCopy := &_BACnetConstructedDataLastCredentialRemovedTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LastCredentialRemovedTime.DeepCopy().(BACnetDateTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastCredentialRemovedTimeCopy
}

func (m *_BACnetConstructedDataLastCredentialRemovedTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
