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

// BACnetIPModeTagged is the corresponding interface of BACnetIPModeTagged
type BACnetIPModeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetIPMode
	// IsBACnetIPModeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetIPModeTagged()
}

// _BACnetIPModeTagged is the data-structure of this message
type _BACnetIPModeTagged struct {
	Header BACnetTagHeader
	Value  BACnetIPMode

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetIPModeTagged = (*_BACnetIPModeTagged)(nil)

// NewBACnetIPModeTagged factory function for _BACnetIPModeTagged
func NewBACnetIPModeTagged(header BACnetTagHeader, value BACnetIPMode, tagNumber uint8, tagClass TagClass) *_BACnetIPModeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetIPModeTagged must not be nil")
	}
	return &_BACnetIPModeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetIPModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetIPModeTagged) GetValue() BACnetIPMode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetIPModeTagged(structType any) BACnetIPModeTagged {
	if casted, ok := structType.(BACnetIPModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetIPModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetIPModeTagged) GetTypeName() string {
	return "BACnetIPModeTagged"
}

func (m *_BACnetIPModeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetIPModeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetIPModeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetIPModeTagged, error) {
	return BACnetIPModeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetIPModeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetIPModeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetIPModeTagged, error) {
		return BACnetIPModeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetIPModeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetIPModeTagged, error) {
	v, err := (&_BACnetIPModeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetIPModeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetIPModeTagged BACnetIPModeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetIPModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetIPModeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetIPMode](ctx, "value", readBuffer, EnsureType[BACnetIPMode](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetIPMode_NORMAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetIPModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetIPModeTagged")
	}

	return m, nil
}

func (m *_BACnetIPModeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetIPModeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetIPModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetIPModeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetIPMode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetIPModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetIPModeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetIPModeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetIPModeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetIPModeTagged) IsBACnetIPModeTagged() {}

func (m *_BACnetIPModeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetIPModeTagged) deepCopy() *_BACnetIPModeTagged {
	if m == nil {
		return nil
	}
	_BACnetIPModeTaggedCopy := &_BACnetIPModeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetIPModeTaggedCopy
}

func (m *_BACnetIPModeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
