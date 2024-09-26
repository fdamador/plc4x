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

// BACnetPropertyStatesEscalatorMode is the corresponding interface of BACnetPropertyStatesEscalatorMode
type BACnetPropertyStatesEscalatorMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetEscalatorMode returns EscalatorMode (property field)
	GetEscalatorMode() BACnetEscalatorModeTagged
	// IsBACnetPropertyStatesEscalatorMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesEscalatorMode()
}

// _BACnetPropertyStatesEscalatorMode is the data-structure of this message
type _BACnetPropertyStatesEscalatorMode struct {
	BACnetPropertyStatesContract
	EscalatorMode BACnetEscalatorModeTagged
}

var _ BACnetPropertyStatesEscalatorMode = (*_BACnetPropertyStatesEscalatorMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesEscalatorMode)(nil)

// NewBACnetPropertyStatesEscalatorMode factory function for _BACnetPropertyStatesEscalatorMode
func NewBACnetPropertyStatesEscalatorMode(peekedTagHeader BACnetTagHeader, escalatorMode BACnetEscalatorModeTagged) *_BACnetPropertyStatesEscalatorMode {
	if escalatorMode == nil {
		panic("escalatorMode of type BACnetEscalatorModeTagged for BACnetPropertyStatesEscalatorMode must not be nil")
	}
	_result := &_BACnetPropertyStatesEscalatorMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		EscalatorMode:                escalatorMode,
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

func (m *_BACnetPropertyStatesEscalatorMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorMode) GetEscalatorMode() BACnetEscalatorModeTagged {
	return m.EscalatorMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorMode(structType any) BACnetPropertyStatesEscalatorMode {
	if casted, ok := structType.(BACnetPropertyStatesEscalatorMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorMode) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorMode"
}

func (m *_BACnetPropertyStatesEscalatorMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (escalatorMode)
	lengthInBits += m.EscalatorMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEscalatorMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesEscalatorMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesEscalatorMode BACnetPropertyStatesEscalatorMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	escalatorMode, err := ReadSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", ReadComplex[BACnetEscalatorModeTagged](BACnetEscalatorModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalatorMode' field"))
	}
	m.EscalatorMode = escalatorMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesEscalatorMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorMode")
		}

		if err := WriteSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", m.GetEscalatorMode(), WriteComplex[BACnetEscalatorModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'escalatorMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEscalatorMode) IsBACnetPropertyStatesEscalatorMode() {}

func (m *_BACnetPropertyStatesEscalatorMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesEscalatorMode) deepCopy() *_BACnetPropertyStatesEscalatorMode {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesEscalatorModeCopy := &_BACnetPropertyStatesEscalatorMode{
		m.BACnetPropertyStatesContract.DeepCopy().(BACnetPropertyStatesContract),
		m.EscalatorMode.DeepCopy().(BACnetEscalatorModeTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesEscalatorModeCopy
}

func (m *_BACnetPropertyStatesEscalatorMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
