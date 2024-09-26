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

// Apdu is the corresponding interface of Apdu
type Apdu interface {
	ApduContract
	ApduRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsApdu is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApdu()
}

// ApduContract provides a set of functions which can be overwritten by a sub struct
type ApduContract interface {
	// GetNumbered returns Numbered (property field)
	GetNumbered() bool
	// GetCounter returns Counter (property field)
	GetCounter() uint8
	// GetDataLength() returns a parser argument
	GetDataLength() uint8
	// IsApdu is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApdu()
}

// ApduRequirements provides a set of functions which need to be implemented by a sub struct
type ApduRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetControl returns Control (discriminator field)
	GetControl() uint8
}

// _Apdu is the data-structure of this message
type _Apdu struct {
	_SubType Apdu
	Numbered bool
	Counter  uint8

	// Arguments.
	DataLength uint8
}

var _ ApduContract = (*_Apdu)(nil)

// NewApdu factory function for _Apdu
func NewApdu(numbered bool, counter uint8, dataLength uint8) *_Apdu {
	return &_Apdu{Numbered: numbered, Counter: counter, DataLength: dataLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Apdu) GetNumbered() bool {
	return m.Numbered
}

func (m *_Apdu) GetCounter() uint8 {
	return m.Counter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApdu(structType any) Apdu {
	if casted, ok := structType.(Apdu); ok {
		return casted
	}
	if casted, ok := structType.(*Apdu); ok {
		return *casted
	}
	return nil
}

func (m *_Apdu) GetTypeName() string {
	return "Apdu"
}

func (m *_Apdu) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (control)
	lengthInBits += 1

	// Simple field (numbered)
	lengthInBits += 1

	// Simple field (counter)
	lengthInBits += 4

	return lengthInBits
}

func (m *_Apdu) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ApduParse[T Apdu](ctx context.Context, theBytes []byte, dataLength uint8) (T, error) {
	return ApduParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduParseWithBufferProducer[T Apdu](dataLength uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ApduParseWithBuffer[T](ctx, readBuffer, dataLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ApduParseWithBuffer[T Apdu](ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (T, error) {
	v, err := (&_Apdu{DataLength: dataLength}).parse(ctx, readBuffer, dataLength)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_Apdu) parse(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (__apdu Apdu, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Apdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Apdu")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	control, err := ReadDiscriminatorField[uint8](ctx, "control", ReadUnsignedByte(readBuffer, uint8(1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'control' field"))
	}

	numbered, err := ReadSimpleField(ctx, "numbered", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numbered' field"))
	}
	m.Numbered = numbered

	counter, err := ReadSimpleField(ctx, "counter", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'counter' field"))
	}
	m.Counter = counter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child Apdu
	switch {
	case control == uint8(1): // ApduControlContainer
		if _child, err = new(_ApduControlContainer).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduControlContainer for type-switch of Apdu")
		}
	case control == uint8(0): // ApduDataContainer
		if _child, err = new(_ApduDataContainer).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataContainer for type-switch of Apdu")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [control=%v]", control)
	}

	if closeErr := readBuffer.CloseContext("Apdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Apdu")
	}

	return _child, nil
}

func (pm *_Apdu) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Apdu, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Apdu"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Apdu")
	}

	if err := WriteDiscriminatorField(ctx, "control", m.GetControl(), WriteUnsignedByte(writeBuffer, 1)); err != nil {
		return errors.Wrap(err, "Error serializing 'control' field")
	}

	if err := WriteSimpleField[bool](ctx, "numbered", m.GetNumbered(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'numbered' field")
	}

	if err := WriteSimpleField[uint8](ctx, "counter", m.GetCounter(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'counter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Apdu"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Apdu")
	}
	return nil
}

////
// Arguments Getter

func (m *_Apdu) GetDataLength() uint8 {
	return m.DataLength
}

//
////

func (m *_Apdu) IsApdu() {}

func (m *_Apdu) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Apdu) deepCopy() *_Apdu {
	if m == nil {
		return nil
	}
	_ApduCopy := &_Apdu{
		nil, // will be set by child
		m.Numbered,
		m.Counter,
		m.DataLength,
	}
	return _ApduCopy
}
