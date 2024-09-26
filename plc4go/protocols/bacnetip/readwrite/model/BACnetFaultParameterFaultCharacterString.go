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

// BACnetFaultParameterFaultCharacterString is the corresponding interface of BACnetFaultParameterFaultCharacterString
type BACnetFaultParameterFaultCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultCharacterStringListOfFaultValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultCharacterString()
}

// _BACnetFaultParameterFaultCharacterString is the data-structure of this message
type _BACnetFaultParameterFaultCharacterString struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultCharacterString = (*_BACnetFaultParameterFaultCharacterString)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultCharacterString)(nil)

// NewBACnetFaultParameterFaultCharacterString factory function for _BACnetFaultParameterFaultCharacterString
func NewBACnetFaultParameterFaultCharacterString(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultCharacterString {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	if listOfFaultValues == nil {
		panic("listOfFaultValues of type BACnetFaultParameterFaultCharacterStringListOfFaultValues for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	_result := &_BACnetFaultParameterFaultCharacterString{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfFaultValues:            listOfFaultValues,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultCharacterString) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultCharacterString) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultCharacterString) GetListOfFaultValues() BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultCharacterString) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultCharacterString(structType any) BACnetFaultParameterFaultCharacterString {
	if casted, ok := structType.(BACnetFaultParameterFaultCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultCharacterString) GetTypeName() string {
	return "BACnetFaultParameterFaultCharacterString"
}

func (m *_BACnetFaultParameterFaultCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultCharacterString BACnetFaultParameterFaultCharacterString, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultCharacterStringListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultCharacterStringListOfFaultValues](BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultCharacterString")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultCharacterString")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultCharacterStringListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultCharacterStringListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultCharacterString")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultCharacterString) IsBACnetFaultParameterFaultCharacterString() {}

func (m *_BACnetFaultParameterFaultCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultCharacterString) deepCopy() *_BACnetFaultParameterFaultCharacterString {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultCharacterStringCopy := &_BACnetFaultParameterFaultCharacterString{
		m.BACnetFaultParameterContract.DeepCopy().(BACnetFaultParameterContract),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ListOfFaultValues.DeepCopy().(BACnetFaultParameterFaultCharacterStringListOfFaultValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultCharacterStringCopy
}

func (m *_BACnetFaultParameterFaultCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
