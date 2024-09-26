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

// FirmataCommandSysex is the corresponding interface of FirmataCommandSysex
type FirmataCommandSysex interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataCommand
	// GetCommand returns Command (property field)
	GetCommand() SysexCommand
	// IsFirmataCommandSysex is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataCommandSysex()
}

// _FirmataCommandSysex is the data-structure of this message
type _FirmataCommandSysex struct {
	FirmataCommandContract
	Command SysexCommand
	// Reserved Fields
	reservedField0 *uint8
}

var _ FirmataCommandSysex = (*_FirmataCommandSysex)(nil)
var _ FirmataCommandRequirements = (*_FirmataCommandSysex)(nil)

// NewFirmataCommandSysex factory function for _FirmataCommandSysex
func NewFirmataCommandSysex(command SysexCommand, response bool) *_FirmataCommandSysex {
	if command == nil {
		panic("command of type SysexCommand for FirmataCommandSysex must not be nil")
	}
	_result := &_FirmataCommandSysex{
		FirmataCommandContract: NewFirmataCommand(response),
		Command:                command,
	}
	_result.FirmataCommandContract.(*_FirmataCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSysex) GetCommandCode() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSysex) GetParent() FirmataCommandContract {
	return m.FirmataCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandSysex) GetCommand() SysexCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataCommandSysex(structType any) FirmataCommandSysex {
	if casted, ok := structType.(FirmataCommandSysex); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSysex); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSysex) GetTypeName() string {
	return "FirmataCommandSysex"
}

func (m *_FirmataCommandSysex) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataCommandContract.(*_FirmataCommand).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_FirmataCommandSysex) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataCommandSysex) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataCommand, response bool) (__firmataCommandSysex FirmataCommandSysex, err error) {
	m.FirmataCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSysex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSysex")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[SysexCommand](ctx, "command", ReadComplex[SysexCommand](SysexCommandParseWithBufferProducer[SysexCommand]((bool)(response)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0xF7))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("FirmataCommandSysex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSysex")
	}

	return m, nil
}

func (m *_FirmataCommandSysex) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSysex) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSysex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSysex")
		}

		if err := WriteSimpleField[SysexCommand](ctx, "command", m.GetCommand(), WriteComplex[SysexCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0xF7), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSysex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSysex")
		}
		return nil
	}
	return m.FirmataCommandContract.(*_FirmataCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataCommandSysex) IsFirmataCommandSysex() {}

func (m *_FirmataCommandSysex) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataCommandSysex) deepCopy() *_FirmataCommandSysex {
	if m == nil {
		return nil
	}
	_FirmataCommandSysexCopy := &_FirmataCommandSysex{
		m.FirmataCommandContract.(*_FirmataCommand).deepCopy(),
		m.Command.DeepCopy().(SysexCommand),
		m.reservedField0,
	}
	m.FirmataCommandContract.(*_FirmataCommand)._SubType = m
	return _FirmataCommandSysexCopy
}

func (m *_FirmataCommandSysex) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
