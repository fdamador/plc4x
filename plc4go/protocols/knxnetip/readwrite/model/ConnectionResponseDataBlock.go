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

// ConnectionResponseDataBlock is the corresponding interface of ConnectionResponseDataBlock
type ConnectionResponseDataBlock interface {
	ConnectionResponseDataBlockContract
	ConnectionResponseDataBlockRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsConnectionResponseDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionResponseDataBlock()
}

// ConnectionResponseDataBlockContract provides a set of functions which can be overwritten by a sub struct
type ConnectionResponseDataBlockContract interface {
	utils.Copyable
	// IsConnectionResponseDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionResponseDataBlock()
}

// ConnectionResponseDataBlockRequirements provides a set of functions which need to be implemented by a sub struct
type ConnectionResponseDataBlockRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetConnectionType returns ConnectionType (discriminator field)
	GetConnectionType() uint8
}

// _ConnectionResponseDataBlock is the data-structure of this message
type _ConnectionResponseDataBlock struct {
	_SubType ConnectionResponseDataBlock
}

var _ ConnectionResponseDataBlockContract = (*_ConnectionResponseDataBlock)(nil)

// NewConnectionResponseDataBlock factory function for _ConnectionResponseDataBlock
func NewConnectionResponseDataBlock() *_ConnectionResponseDataBlock {
	return &_ConnectionResponseDataBlock{}
}

// Deprecated: use the interface for direct cast
func CastConnectionResponseDataBlock(structType any) ConnectionResponseDataBlock {
	if casted, ok := structType.(ConnectionResponseDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponseDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponseDataBlock) GetTypeName() string {
	return "ConnectionResponseDataBlock"
}

func (m *_ConnectionResponseDataBlock) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionResponseDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ConnectionResponseDataBlockParse[T ConnectionResponseDataBlock](ctx context.Context, theBytes []byte) (T, error) {
	return ConnectionResponseDataBlockParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConnectionResponseDataBlockParseWithBufferProducer[T ConnectionResponseDataBlock]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ConnectionResponseDataBlockParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ConnectionResponseDataBlockParseWithBuffer[T ConnectionResponseDataBlock](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ConnectionResponseDataBlock{}).parse(ctx, readBuffer)
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

func (m *_ConnectionResponseDataBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__connectionResponseDataBlock ConnectionResponseDataBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponseDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	connectionType, err := ReadDiscriminatorField[uint8](ctx, "connectionType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ConnectionResponseDataBlock
	switch {
	case connectionType == 0x03: // ConnectionResponseDataBlockDeviceManagement
		if _child, err = new(_ConnectionResponseDataBlockDeviceManagement).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionResponseDataBlockDeviceManagement for type-switch of ConnectionResponseDataBlock")
		}
	case connectionType == 0x04: // ConnectionResponseDataBlockTunnelConnection
		if _child, err = new(_ConnectionResponseDataBlockTunnelConnection).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionResponseDataBlockTunnelConnection for type-switch of ConnectionResponseDataBlock")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [connectionType=%v]", connectionType)
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponseDataBlock")
	}

	return _child, nil
}

func (pm *_ConnectionResponseDataBlock) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ConnectionResponseDataBlock, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ConnectionResponseDataBlock")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteDiscriminatorField(ctx, "connectionType", m.GetConnectionType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ConnectionResponseDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ConnectionResponseDataBlock")
	}
	return nil
}

func (m *_ConnectionResponseDataBlock) IsConnectionResponseDataBlock() {}

func (m *_ConnectionResponseDataBlock) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionResponseDataBlock) deepCopy() *_ConnectionResponseDataBlock {
	if m == nil {
		return nil
	}
	_ConnectionResponseDataBlockCopy := &_ConnectionResponseDataBlock{
		nil, // will be set by child
	}
	return _ConnectionResponseDataBlockCopy
}
