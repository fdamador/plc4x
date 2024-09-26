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

// BACnetPropertyStatesLiftCarMode is the corresponding interface of BACnetPropertyStatesLiftCarMode
type BACnetPropertyStatesLiftCarMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLiftCarMode returns LiftCarMode (property field)
	GetLiftCarMode() BACnetLiftCarModeTagged
	// IsBACnetPropertyStatesLiftCarMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLiftCarMode()
}

// _BACnetPropertyStatesLiftCarMode is the data-structure of this message
type _BACnetPropertyStatesLiftCarMode struct {
	BACnetPropertyStatesContract
	LiftCarMode BACnetLiftCarModeTagged
}

var _ BACnetPropertyStatesLiftCarMode = (*_BACnetPropertyStatesLiftCarMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLiftCarMode)(nil)

// NewBACnetPropertyStatesLiftCarMode factory function for _BACnetPropertyStatesLiftCarMode
func NewBACnetPropertyStatesLiftCarMode(peekedTagHeader BACnetTagHeader, liftCarMode BACnetLiftCarModeTagged) *_BACnetPropertyStatesLiftCarMode {
	if liftCarMode == nil {
		panic("liftCarMode of type BACnetLiftCarModeTagged for BACnetPropertyStatesLiftCarMode must not be nil")
	}
	_result := &_BACnetPropertyStatesLiftCarMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LiftCarMode:                  liftCarMode,
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

func (m *_BACnetPropertyStatesLiftCarMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarMode) GetLiftCarMode() BACnetLiftCarModeTagged {
	return m.LiftCarMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarMode(structType any) BACnetPropertyStatesLiftCarMode {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarMode) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarMode"
}

func (m *_BACnetPropertyStatesLiftCarMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (liftCarMode)
	lengthInBits += m.LiftCarMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftCarMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLiftCarMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLiftCarMode BACnetPropertyStatesLiftCarMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	liftCarMode, err := ReadSimpleField[BACnetLiftCarModeTagged](ctx, "liftCarMode", ReadComplex[BACnetLiftCarModeTagged](BACnetLiftCarModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'liftCarMode' field"))
	}
	m.LiftCarMode = liftCarMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLiftCarMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftCarMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarMode")
		}

		if err := WriteSimpleField[BACnetLiftCarModeTagged](ctx, "liftCarMode", m.GetLiftCarMode(), WriteComplex[BACnetLiftCarModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'liftCarMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftCarMode) IsBACnetPropertyStatesLiftCarMode() {}

func (m *_BACnetPropertyStatesLiftCarMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLiftCarMode) deepCopy() *_BACnetPropertyStatesLiftCarMode {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLiftCarModeCopy := &_BACnetPropertyStatesLiftCarMode{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.LiftCarMode.DeepCopy().(BACnetLiftCarModeTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLiftCarModeCopy
}

func (m *_BACnetPropertyStatesLiftCarMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
