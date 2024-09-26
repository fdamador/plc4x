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

// BACnetPropertyStatesDoorSecuredStatus is the corresponding interface of BACnetPropertyStatesDoorSecuredStatus
type BACnetPropertyStatesDoorSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetDoorSecuredStatus returns DoorSecuredStatus (property field)
	GetDoorSecuredStatus() BACnetDoorSecuredStatusTagged
	// IsBACnetPropertyStatesDoorSecuredStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesDoorSecuredStatus()
}

// _BACnetPropertyStatesDoorSecuredStatus is the data-structure of this message
type _BACnetPropertyStatesDoorSecuredStatus struct {
	BACnetPropertyStatesContract
	DoorSecuredStatus BACnetDoorSecuredStatusTagged
}

var _ BACnetPropertyStatesDoorSecuredStatus = (*_BACnetPropertyStatesDoorSecuredStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesDoorSecuredStatus)(nil)

// NewBACnetPropertyStatesDoorSecuredStatus factory function for _BACnetPropertyStatesDoorSecuredStatus
func NewBACnetPropertyStatesDoorSecuredStatus(peekedTagHeader BACnetTagHeader, doorSecuredStatus BACnetDoorSecuredStatusTagged) *_BACnetPropertyStatesDoorSecuredStatus {
	if doorSecuredStatus == nil {
		panic("doorSecuredStatus of type BACnetDoorSecuredStatusTagged for BACnetPropertyStatesDoorSecuredStatus must not be nil")
	}
	_result := &_BACnetPropertyStatesDoorSecuredStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		DoorSecuredStatus:            doorSecuredStatus,
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

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetDoorSecuredStatus() BACnetDoorSecuredStatusTagged {
	return m.DoorSecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorSecuredStatus(structType any) BACnetPropertyStatesDoorSecuredStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorSecuredStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorSecuredStatus"
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (doorSecuredStatus)
	lengthInBits += m.DoorSecuredStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesDoorSecuredStatus BACnetPropertyStatesDoorSecuredStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorSecuredStatus, err := ReadSimpleField[BACnetDoorSecuredStatusTagged](ctx, "doorSecuredStatus", ReadComplex[BACnetDoorSecuredStatusTagged](BACnetDoorSecuredStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorSecuredStatus' field"))
	}
	m.DoorSecuredStatus = doorSecuredStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorSecuredStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorSecuredStatus")
		}

		if err := WriteSimpleField[BACnetDoorSecuredStatusTagged](ctx, "doorSecuredStatus", m.GetDoorSecuredStatus(), WriteComplex[BACnetDoorSecuredStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorSecuredStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorSecuredStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) IsBACnetPropertyStatesDoorSecuredStatus() {}

func (m *_BACnetPropertyStatesDoorSecuredStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) deepCopy() *_BACnetPropertyStatesDoorSecuredStatus {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesDoorSecuredStatusCopy := &_BACnetPropertyStatesDoorSecuredStatus{
		m.BACnetPropertyStatesContract.DeepCopy().(BACnetPropertyStatesContract),
		m.DoorSecuredStatus.DeepCopy().(BACnetDoorSecuredStatusTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesDoorSecuredStatusCopy
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
