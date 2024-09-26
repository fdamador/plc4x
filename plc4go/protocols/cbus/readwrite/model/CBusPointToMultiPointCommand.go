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

// CBusPointToMultiPointCommand is the corresponding interface of CBusPointToMultiPointCommand
type CBusPointToMultiPointCommand interface {
	CBusPointToMultiPointCommandContract
	CBusPointToMultiPointCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCBusPointToMultiPointCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToMultiPointCommand()
}

// CBusPointToMultiPointCommandContract provides a set of functions which can be overwritten by a sub struct
type CBusPointToMultiPointCommandContract interface {
	utils.Copyable
	// GetPeekedApplication returns PeekedApplication (property field)
	GetPeekedApplication() byte
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// IsCBusPointToMultiPointCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToMultiPointCommand()
}

// CBusPointToMultiPointCommandRequirements provides a set of functions which need to be implemented by a sub struct
type CBusPointToMultiPointCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedApplication returns PeekedApplication (discriminator field)
	GetPeekedApplication() byte
}

// _CBusPointToMultiPointCommand is the data-structure of this message
type _CBusPointToMultiPointCommand struct {
	_SubType          CBusPointToMultiPointCommand
	PeekedApplication byte

	// Arguments.
	CBusOptions CBusOptions
}

var _ CBusPointToMultiPointCommandContract = (*_CBusPointToMultiPointCommand)(nil)

// NewCBusPointToMultiPointCommand factory function for _CBusPointToMultiPointCommand
func NewCBusPointToMultiPointCommand(peekedApplication byte, cBusOptions CBusOptions) *_CBusPointToMultiPointCommand {
	return &_CBusPointToMultiPointCommand{PeekedApplication: peekedApplication, CBusOptions: cBusOptions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToMultiPointCommand) GetPeekedApplication() byte {
	return m.PeekedApplication
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToMultiPointCommand(structType any) CBusPointToMultiPointCommand {
	if casted, ok := structType.(CBusPointToMultiPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToMultiPointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToMultiPointCommand) GetTypeName() string {
	return "CBusPointToMultiPointCommand"
}

func (m *_CBusPointToMultiPointCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_CBusPointToMultiPointCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CBusPointToMultiPointCommandParse[T CBusPointToMultiPointCommand](ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (T, error) {
	return CBusPointToMultiPointCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToMultiPointCommandParseWithBufferProducer[T CBusPointToMultiPointCommand](cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusPointToMultiPointCommandParseWithBuffer[T](ctx, readBuffer, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CBusPointToMultiPointCommandParseWithBuffer[T CBusPointToMultiPointCommand](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusPointToMultiPointCommand{CBusOptions: cBusOptions}).parse(ctx, readBuffer, cBusOptions)
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

func (m *_CBusPointToMultiPointCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (__cBusPointToMultiPointCommand CBusPointToMultiPointCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToMultiPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedApplication, err := ReadPeekField[byte](ctx, "peekedApplication", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedApplication' field"))
	}
	m.PeekedApplication = peekedApplication

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusPointToMultiPointCommand
	switch {
	case peekedApplication == 0xFF: // CBusPointToMultiPointCommandStatus
		if _child, err = new(_CBusPointToMultiPointCommandStatus).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToMultiPointCommandStatus for type-switch of CBusPointToMultiPointCommand")
		}
	case 0 == 0: // CBusPointToMultiPointCommandNormal
		if _child, err = new(_CBusPointToMultiPointCommandNormal).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToMultiPointCommandNormal for type-switch of CBusPointToMultiPointCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedApplication=%v]", peekedApplication)
	}

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToMultiPointCommand")
	}

	return _child, nil
}

func (pm *_CBusPointToMultiPointCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusPointToMultiPointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToMultiPointCommand")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToMultiPointCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusPointToMultiPointCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusPointToMultiPointCommand) IsCBusPointToMultiPointCommand() {}

func (m *_CBusPointToMultiPointCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToMultiPointCommand) deepCopy() *_CBusPointToMultiPointCommand {
	if m == nil {
		return nil
	}
	_CBusPointToMultiPointCommandCopy := &_CBusPointToMultiPointCommand{
		nil, // will be set by child
		m.PeekedApplication,
		m.CBusOptions,
	}
	return _CBusPointToMultiPointCommandCopy
}
