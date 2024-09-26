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

// BACnetPropertyStatesEscalatorOperationDirection is the corresponding interface of BACnetPropertyStatesEscalatorOperationDirection
type BACnetPropertyStatesEscalatorOperationDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetEscalatorOperationDirection returns EscalatorOperationDirection (property field)
	GetEscalatorOperationDirection() BACnetEscalatorOperationDirectionTagged
	// IsBACnetPropertyStatesEscalatorOperationDirection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesEscalatorOperationDirection()
}

// _BACnetPropertyStatesEscalatorOperationDirection is the data-structure of this message
type _BACnetPropertyStatesEscalatorOperationDirection struct {
	BACnetPropertyStatesContract
	EscalatorOperationDirection BACnetEscalatorOperationDirectionTagged
}

var _ BACnetPropertyStatesEscalatorOperationDirection = (*_BACnetPropertyStatesEscalatorOperationDirection)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesEscalatorOperationDirection)(nil)

// NewBACnetPropertyStatesEscalatorOperationDirection factory function for _BACnetPropertyStatesEscalatorOperationDirection
func NewBACnetPropertyStatesEscalatorOperationDirection(peekedTagHeader BACnetTagHeader, escalatorOperationDirection BACnetEscalatorOperationDirectionTagged) *_BACnetPropertyStatesEscalatorOperationDirection {
	if escalatorOperationDirection == nil {
		panic("escalatorOperationDirection of type BACnetEscalatorOperationDirectionTagged for BACnetPropertyStatesEscalatorOperationDirection must not be nil")
	}
	_result := &_BACnetPropertyStatesEscalatorOperationDirection{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		EscalatorOperationDirection:  escalatorOperationDirection,
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

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetEscalatorOperationDirection() BACnetEscalatorOperationDirectionTagged {
	return m.EscalatorOperationDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorOperationDirection(structType any) BACnetPropertyStatesEscalatorOperationDirection {
	if casted, ok := structType.(BACnetPropertyStatesEscalatorOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorOperationDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorOperationDirection"
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (escalatorOperationDirection)
	lengthInBits += m.EscalatorOperationDirection.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesEscalatorOperationDirection BACnetPropertyStatesEscalatorOperationDirection, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	escalatorOperationDirection, err := ReadSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "escalatorOperationDirection", ReadComplex[BACnetEscalatorOperationDirectionTagged](BACnetEscalatorOperationDirectionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalatorOperationDirection' field"))
	}
	m.EscalatorOperationDirection = escalatorOperationDirection

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorOperationDirection")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorOperationDirection")
		}

		if err := WriteSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "escalatorOperationDirection", m.GetEscalatorOperationDirection(), WriteComplex[BACnetEscalatorOperationDirectionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'escalatorOperationDirection' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorOperationDirection")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) IsBACnetPropertyStatesEscalatorOperationDirection() {
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) deepCopy() *_BACnetPropertyStatesEscalatorOperationDirection {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesEscalatorOperationDirectionCopy := &_BACnetPropertyStatesEscalatorOperationDirection{
		m.BACnetPropertyStatesContract.DeepCopy().(BACnetPropertyStatesContract),
		m.EscalatorOperationDirection.DeepCopy().(BACnetEscalatorOperationDirectionTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesEscalatorOperationDirectionCopy
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
