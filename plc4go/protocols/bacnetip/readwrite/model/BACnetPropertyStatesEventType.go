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

// BACnetPropertyStatesEventType is the corresponding interface of BACnetPropertyStatesEventType
type BACnetPropertyStatesEventType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// IsBACnetPropertyStatesEventType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesEventType()
}

// _BACnetPropertyStatesEventType is the data-structure of this message
type _BACnetPropertyStatesEventType struct {
	BACnetPropertyStatesContract
	EventType BACnetEventTypeTagged
}

var _ BACnetPropertyStatesEventType = (*_BACnetPropertyStatesEventType)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesEventType)(nil)

// NewBACnetPropertyStatesEventType factory function for _BACnetPropertyStatesEventType
func NewBACnetPropertyStatesEventType(peekedTagHeader BACnetTagHeader, eventType BACnetEventTypeTagged) *_BACnetPropertyStatesEventType {
	if eventType == nil {
		panic("eventType of type BACnetEventTypeTagged for BACnetPropertyStatesEventType must not be nil")
	}
	_result := &_BACnetPropertyStatesEventType{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		EventType:                    eventType,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
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

func (m *_BACnetPropertyStatesEventType) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEventType) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEventType(structType any) BACnetPropertyStatesEventType {
	if casted, ok := structType.(BACnetPropertyStatesEventType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEventType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEventType) GetTypeName() string {
	return "BACnetPropertyStatesEventType"
}

func (m *_BACnetPropertyStatesEventType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEventType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesEventType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesEventType BACnetPropertyStatesEventType, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEventType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}
	m.EventType = eventType

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEventType")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesEventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEventType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEventType")
		}

		if err := WriteSimpleField[BACnetEventTypeTagged](ctx, "eventType", m.GetEventType(), WriteComplex[BACnetEventTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEventType")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEventType) IsBACnetPropertyStatesEventType() {}

func (m *_BACnetPropertyStatesEventType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesEventType) deepCopy() *_BACnetPropertyStatesEventType {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesEventTypeCopy := &_BACnetPropertyStatesEventType{
		m.BACnetPropertyStatesContract.DeepCopy().(BACnetPropertyStatesContract),
		m.EventType.DeepCopy().(BACnetEventTypeTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesEventTypeCopy
}

func (m *_BACnetPropertyStatesEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
