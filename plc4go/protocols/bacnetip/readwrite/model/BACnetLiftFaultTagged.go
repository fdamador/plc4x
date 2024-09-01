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

// BACnetLiftFaultTagged is the corresponding interface of BACnetLiftFaultTagged
type BACnetLiftFaultTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftFault
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetLiftFaultTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftFaultTagged.
// This is useful for switch cases.
type BACnetLiftFaultTaggedExactly interface {
	BACnetLiftFaultTagged
	isBACnetLiftFaultTagged() bool
}

// _BACnetLiftFaultTagged is the data-structure of this message
type _BACnetLiftFaultTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLiftFault
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftFaultTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftFaultTagged) GetValue() BACnetLiftFault {
	return m.Value
}

func (m *_BACnetLiftFaultTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetLiftFaultTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLiftFault_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftFaultTagged factory function for _BACnetLiftFaultTagged
func NewBACnetLiftFaultTagged(header BACnetTagHeader, value BACnetLiftFault, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLiftFaultTagged {
	return &_BACnetLiftFaultTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftFaultTagged(structType any) BACnetLiftFaultTagged {
	if casted, ok := structType.(BACnetLiftFaultTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftFaultTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftFaultTagged) GetTypeName() string {
	return "BACnetLiftFaultTagged"
}

func (m *_BACnetLiftFaultTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetLiftFaultTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftFaultTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLiftFaultTagged, error) {
	return BACnetLiftFaultTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLiftFaultTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftFaultTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftFaultTagged, error) {
		return BACnetLiftFaultTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLiftFaultTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftFaultTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftFaultTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftFaultTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetLiftFault](ctx, "value", readBuffer, EnsureType[BACnetLiftFault](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLiftFault_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLiftFault_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftFaultTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftFaultTagged")
	}

	// Create the instance
	return &_BACnetLiftFaultTagged{
		TagNumber:        tagNumber,
		TagClass:         tagClass,
		Header:           header,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetLiftFaultTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftFaultTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftFaultTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftFaultTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLiftFault](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetLiftFaultTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftFaultTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftFaultTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLiftFaultTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLiftFaultTagged) isBACnetLiftFaultTagged() bool {
	return true
}

func (m *_BACnetLiftFaultTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
