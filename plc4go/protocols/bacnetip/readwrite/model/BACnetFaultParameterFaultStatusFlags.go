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

// BACnetFaultParameterFaultStatusFlags is the corresponding interface of BACnetFaultParameterFaultStatusFlags
type BACnetFaultParameterFaultStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetStatusFlagsReference returns StatusFlagsReference (property field)
	GetStatusFlagsReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultStatusFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultStatusFlags()
}

// _BACnetFaultParameterFaultStatusFlags is the data-structure of this message
type _BACnetFaultParameterFaultStatusFlags struct {
	BACnetFaultParameterContract
	OpeningTag           BACnetOpeningTag
	StatusFlagsReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag           BACnetClosingTag
}

var _ BACnetFaultParameterFaultStatusFlags = (*_BACnetFaultParameterFaultStatusFlags)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultStatusFlags)(nil)

// NewBACnetFaultParameterFaultStatusFlags factory function for _BACnetFaultParameterFaultStatusFlags
func NewBACnetFaultParameterFaultStatusFlags(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, statusFlagsReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultStatusFlags {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultStatusFlags must not be nil")
	}
	if statusFlagsReference == nil {
		panic("statusFlagsReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetFaultParameterFaultStatusFlags must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultStatusFlags must not be nil")
	}
	_result := &_BACnetFaultParameterFaultStatusFlags{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		StatusFlagsReference:         statusFlagsReference,
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

func (m *_BACnetFaultParameterFaultStatusFlags) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultStatusFlags) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetStatusFlagsReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.StatusFlagsReference
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultStatusFlags(structType any) BACnetFaultParameterFaultStatusFlags {
	if casted, ok := structType.(BACnetFaultParameterFaultStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetTypeName() string {
	return "BACnetFaultParameterFaultStatusFlags"
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (statusFlagsReference)
	lengthInBits += m.StatusFlagsReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultStatusFlags BACnetFaultParameterFaultStatusFlags, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(5))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	statusFlagsReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "statusFlagsReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlagsReference' field"))
	}
	m.StatusFlagsReference = statusFlagsReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(5))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultStatusFlags")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultStatusFlags")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "statusFlagsReference", m.GetStatusFlagsReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlagsReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultStatusFlags")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultStatusFlags) IsBACnetFaultParameterFaultStatusFlags() {}

func (m *_BACnetFaultParameterFaultStatusFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultStatusFlags) deepCopy() *_BACnetFaultParameterFaultStatusFlags {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultStatusFlagsCopy := &_BACnetFaultParameterFaultStatusFlags{
		m.BACnetFaultParameterContract.DeepCopy().(BACnetFaultParameterContract),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.StatusFlagsReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultStatusFlagsCopy
}

func (m *_BACnetFaultParameterFaultStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
