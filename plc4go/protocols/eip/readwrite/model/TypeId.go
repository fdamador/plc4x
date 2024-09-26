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

// TypeId is the corresponding interface of TypeId
type TypeId interface {
	TypeIdContract
	TypeIdRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsTypeId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTypeId()
}

// TypeIdContract provides a set of functions which can be overwritten by a sub struct
type TypeIdContract interface {
	utils.Copyable
	// IsTypeId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTypeId()
}

// TypeIdRequirements provides a set of functions which need to be implemented by a sub struct
type TypeIdRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetId returns Id (discriminator field)
	GetId() uint16
}

// _TypeId is the data-structure of this message
type _TypeId struct {
	_SubType TypeId
}

var _ TypeIdContract = (*_TypeId)(nil)

// NewTypeId factory function for _TypeId
func NewTypeId() *_TypeId {
	return &_TypeId{}
}

// Deprecated: use the interface for direct cast
func CastTypeId(structType any) TypeId {
	if casted, ok := structType.(TypeId); ok {
		return casted
	}
	if casted, ok := structType.(*TypeId); ok {
		return *casted
	}
	return nil
}

func (m *_TypeId) GetTypeName() string {
	return "TypeId"
}

func (m *_TypeId) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (id)
	lengthInBits += 16

	return lengthInBits
}

func (m *_TypeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func TypeIdParse[T TypeId](ctx context.Context, theBytes []byte) (T, error) {
	return TypeIdParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func TypeIdParseWithBufferProducer[T TypeId]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := TypeIdParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func TypeIdParseWithBuffer[T TypeId](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_TypeId{}).parse(ctx, readBuffer)
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

func (m *_TypeId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__typeId TypeId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TypeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	id, err := ReadDiscriminatorField[uint16](ctx, "id", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child TypeId
	switch {
	case id == 0x0000: // NullAddressItem
		if _child, err = new(_NullAddressItem).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NullAddressItem for type-switch of TypeId")
		}
	case id == 0x0100: // ServicesResponse
		if _child, err = new(_ServicesResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ServicesResponse for type-switch of TypeId")
		}
	case id == 0x00A1: // ConnectedAddressItem
		if _child, err = new(_ConnectedAddressItem).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectedAddressItem for type-switch of TypeId")
		}
	case id == 0x00B1: // ConnectedDataItem
		if _child, err = new(_ConnectedDataItem).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectedDataItem for type-switch of TypeId")
		}
	case id == 0x00B2: // UnConnectedDataItem
		if _child, err = new(_UnConnectedDataItem).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type UnConnectedDataItem for type-switch of TypeId")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [id=%v]", id)
	}

	if closeErr := readBuffer.CloseContext("TypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TypeId")
	}

	return _child, nil
}

func (pm *_TypeId) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TypeId, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TypeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TypeId")
	}

	if err := WriteDiscriminatorField(ctx, "id", m.GetId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'id' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("TypeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TypeId")
	}
	return nil
}

func (m *_TypeId) IsTypeId() {}

func (m *_TypeId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TypeId) deepCopy() *_TypeId {
	if m == nil {
		return nil
	}
	_TypeIdCopy := &_TypeId{
		nil, // will be set by child
	}
	return _TypeIdCopy
}
