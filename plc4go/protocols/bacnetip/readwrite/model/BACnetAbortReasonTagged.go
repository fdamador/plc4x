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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAbortReasonTagged is the corresponding interface of BACnetAbortReasonTagged
type BACnetAbortReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetValue returns Value (property field)
	GetValue() BACnetAbortReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAbortReasonTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAbortReasonTagged()
}

// _BACnetAbortReasonTagged is the data-structure of this message
type _BACnetAbortReasonTagged struct {
	Value            BACnetAbortReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

var _ BACnetAbortReasonTagged = (*_BACnetAbortReasonTagged)(nil)

// NewBACnetAbortReasonTagged factory function for _BACnetAbortReasonTagged
func NewBACnetAbortReasonTagged(value BACnetAbortReason, proprietaryValue uint32, actualLength uint32) *_BACnetAbortReasonTagged {
	return &_BACnetAbortReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAbortReasonTagged) GetValue() BACnetAbortReason {
	return m.Value
}

func (m *_BACnetAbortReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAbortReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAbortReasonTagged(structType any) BACnetAbortReasonTagged {
	if casted, ok := structType.(BACnetAbortReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAbortReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAbortReasonTagged) GetTypeName() string {
	return "BACnetAbortReasonTagged"
}

func (m *_BACnetAbortReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAbortReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAbortReasonTaggedParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetAbortReasonTagged, error) {
	return BACnetAbortReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetAbortReasonTaggedParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAbortReasonTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAbortReasonTagged, error) {
		return BACnetAbortReasonTaggedParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetAbortReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetAbortReasonTagged, error) {
	v, err := (&_BACnetAbortReasonTagged{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAbortReasonTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetAbortReasonTagged BACnetAbortReasonTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAbortReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAbortReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadManualField[BACnetAbortReason](ctx, "value", readBuffer, EnsureType[BACnetAbortReason](ReadEnumGeneric(ctx, readBuffer, actualLength, BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, actualLength, isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAbortReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAbortReasonTagged")
	}

	return m, nil
}

func (m *_BACnetAbortReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAbortReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAbortReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAbortReasonTagged")
	}

	if err := WriteManualField[BACnetAbortReason](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAbortReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAbortReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAbortReasonTagged) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetAbortReasonTagged) IsBACnetAbortReasonTagged() {}

func (m *_BACnetAbortReasonTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAbortReasonTagged) deepCopy() *_BACnetAbortReasonTagged {
	if m == nil {
		return nil
	}
	_BACnetAbortReasonTaggedCopy := &_BACnetAbortReasonTagged{
		m.Value,
		m.ProprietaryValue,
		m.ActualLength,
	}
	return _BACnetAbortReasonTaggedCopy
}

func (m *_BACnetAbortReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
