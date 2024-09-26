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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageCommand is the corresponding interface of FirmataMessageCommand
type FirmataMessageCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataMessage
	// GetCommand returns Command (property field)
	GetCommand() FirmataCommand
	// IsFirmataMessageCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataMessageCommand()
}

// _FirmataMessageCommand is the data-structure of this message
type _FirmataMessageCommand struct {
	FirmataMessageContract
	Command FirmataCommand
}

var _ FirmataMessageCommand = (*_FirmataMessageCommand)(nil)
var _ FirmataMessageRequirements = (*_FirmataMessageCommand)(nil)

// NewFirmataMessageCommand factory function for _FirmataMessageCommand
func NewFirmataMessageCommand(command FirmataCommand, response bool) *_FirmataMessageCommand {
	if command == nil {
		panic("command of type FirmataCommand for FirmataMessageCommand must not be nil")
	}
	_result := &_FirmataMessageCommand{
		FirmataMessageContract: NewFirmataMessage(response),
		Command:                command,
	}
	_result.FirmataMessageContract.(*_FirmataMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageCommand) GetMessageType() uint8 {
	return 0xF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageCommand) GetParent() FirmataMessageContract {
	return m.FirmataMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageCommand) GetCommand() FirmataCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataMessageCommand(structType any) FirmataMessageCommand {
	if casted, ok := structType.(FirmataMessageCommand); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageCommand); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageCommand) GetTypeName() string {
	return "FirmataMessageCommand"
}

func (m *_FirmataMessageCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataMessageContract.(*_FirmataMessage).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_FirmataMessageCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataMessageCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataMessage, response bool) (__firmataMessageCommand FirmataMessageCommand, err error) {
	m.FirmataMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[FirmataCommand](ctx, "command", ReadComplex[FirmataCommand](FirmataCommandParseWithBufferProducer[FirmataCommand]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("FirmataMessageCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageCommand")
	}

	return m, nil
}

func (m *_FirmataMessageCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageCommand")
		}

		if err := WriteSimpleField[FirmataCommand](ctx, "command", m.GetCommand(), WriteComplex[FirmataCommand](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageCommand")
		}
		return nil
	}
	return m.FirmataMessageContract.(*_FirmataMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageCommand) IsFirmataMessageCommand() {}

func (m *_FirmataMessageCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataMessageCommand) deepCopy() *_FirmataMessageCommand {
	if m == nil {
		return nil
	}
	_FirmataMessageCommandCopy := &_FirmataMessageCommand{
		m.FirmataMessageContract.DeepCopy().(FirmataMessageContract),
		m.Command.DeepCopy().(FirmataCommand),
	}
	m.FirmataMessageContract.(*_FirmataMessage)._SubType = m
	return _FirmataMessageCommandCopy
}

func (m *_FirmataMessageCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
