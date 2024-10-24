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

// S7ParameterUserDataItem is the corresponding interface of S7ParameterUserDataItem
type S7ParameterUserDataItem interface {
	S7ParameterUserDataItemContract
	S7ParameterUserDataItemRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsS7ParameterUserDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterUserDataItem()
	// CreateBuilder creates a S7ParameterUserDataItemBuilder
	CreateS7ParameterUserDataItemBuilder() S7ParameterUserDataItemBuilder
}

// S7ParameterUserDataItemContract provides a set of functions which can be overwritten by a sub struct
type S7ParameterUserDataItemContract interface {
	// IsS7ParameterUserDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterUserDataItem()
	// CreateBuilder creates a S7ParameterUserDataItemBuilder
	CreateS7ParameterUserDataItemBuilder() S7ParameterUserDataItemBuilder
}

// S7ParameterUserDataItemRequirements provides a set of functions which need to be implemented by a sub struct
type S7ParameterUserDataItemRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
}

// _S7ParameterUserDataItem is the data-structure of this message
type _S7ParameterUserDataItem struct {
	_SubType interface {
		S7ParameterUserDataItemContract
		S7ParameterUserDataItemRequirements
	}
}

var _ S7ParameterUserDataItemContract = (*_S7ParameterUserDataItem)(nil)

// NewS7ParameterUserDataItem factory function for _S7ParameterUserDataItem
func NewS7ParameterUserDataItem() *_S7ParameterUserDataItem {
	return &_S7ParameterUserDataItem{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterUserDataItemBuilder is a builder for S7ParameterUserDataItem
type S7ParameterUserDataItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7ParameterUserDataItemBuilder
	// AsS7ParameterUserDataItemCPUFunctions converts this build to a subType of S7ParameterUserDataItem. It is always possible to return to current builder using Done()
	AsS7ParameterUserDataItemCPUFunctions() interface {
		S7ParameterUserDataItemCPUFunctionsBuilder
		Done() S7ParameterUserDataItemBuilder
	}
	// Build builds the S7ParameterUserDataItem or returns an error if something is wrong
	PartialBuild() (S7ParameterUserDataItemContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() S7ParameterUserDataItemContract
	// Build builds the S7ParameterUserDataItem or returns an error if something is wrong
	Build() (S7ParameterUserDataItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterUserDataItem
}

// NewS7ParameterUserDataItemBuilder() creates a S7ParameterUserDataItemBuilder
func NewS7ParameterUserDataItemBuilder() S7ParameterUserDataItemBuilder {
	return &_S7ParameterUserDataItemBuilder{_S7ParameterUserDataItem: new(_S7ParameterUserDataItem)}
}

type _S7ParameterUserDataItemChildBuilder interface {
	utils.Copyable
	setParent(S7ParameterUserDataItemContract)
	buildForS7ParameterUserDataItem() (S7ParameterUserDataItem, error)
}

type _S7ParameterUserDataItemBuilder struct {
	*_S7ParameterUserDataItem

	childBuilder _S7ParameterUserDataItemChildBuilder

	err *utils.MultiError
}

var _ (S7ParameterUserDataItemBuilder) = (*_S7ParameterUserDataItemBuilder)(nil)

func (b *_S7ParameterUserDataItemBuilder) WithMandatoryFields() S7ParameterUserDataItemBuilder {
	return b
}

func (b *_S7ParameterUserDataItemBuilder) PartialBuild() (S7ParameterUserDataItemContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterUserDataItem.deepCopy(), nil
}

func (b *_S7ParameterUserDataItemBuilder) PartialMustBuild() S7ParameterUserDataItemContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterUserDataItemBuilder) AsS7ParameterUserDataItemCPUFunctions() interface {
	S7ParameterUserDataItemCPUFunctionsBuilder
	Done() S7ParameterUserDataItemBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterUserDataItemCPUFunctionsBuilder
		Done() S7ParameterUserDataItemBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterUserDataItemCPUFunctionsBuilder().(*_S7ParameterUserDataItemCPUFunctionsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterUserDataItemBuilder) Build() (S7ParameterUserDataItem, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForS7ParameterUserDataItem()
}

func (b *_S7ParameterUserDataItemBuilder) MustBuild() S7ParameterUserDataItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterUserDataItemBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterUserDataItemBuilder().(*_S7ParameterUserDataItemBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_S7ParameterUserDataItemChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterUserDataItemBuilder creates a S7ParameterUserDataItemBuilder
func (b *_S7ParameterUserDataItem) CreateS7ParameterUserDataItemBuilder() S7ParameterUserDataItemBuilder {
	if b == nil {
		return NewS7ParameterUserDataItemBuilder()
	}
	return &_S7ParameterUserDataItemBuilder{_S7ParameterUserDataItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterUserDataItem(structType any) S7ParameterUserDataItem {
	if casted, ok := structType.(S7ParameterUserDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserDataItem) GetTypeName() string {
	return "S7ParameterUserDataItem"
}

func (m *_S7ParameterUserDataItem) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterUserDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7ParameterUserDataItemParse[T S7ParameterUserDataItem](ctx context.Context, theBytes []byte) (T, error) {
	return S7ParameterUserDataItemParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7ParameterUserDataItemParseWithBufferProducer[T S7ParameterUserDataItem]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7ParameterUserDataItemParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func S7ParameterUserDataItemParseWithBuffer[T S7ParameterUserDataItem](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_S7ParameterUserDataItem{}).parse(ctx, readBuffer)
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

func (m *_S7ParameterUserDataItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__s7ParameterUserDataItem S7ParameterUserDataItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemType, err := ReadDiscriminatorField[uint8](ctx, "itemType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7ParameterUserDataItem
	switch {
	case itemType == 0x12: // S7ParameterUserDataItemCPUFunctions
		if _child, err = new(_S7ParameterUserDataItemCPUFunctions).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterUserDataItemCPUFunctions for type-switch of S7ParameterUserDataItem")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [itemType=%v]", itemType)
	}

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserDataItem")
	}

	return _child, nil
}

func (pm *_S7ParameterUserDataItem) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7ParameterUserDataItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7ParameterUserDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7ParameterUserDataItem")
	}

	if err := WriteDiscriminatorField(ctx, "itemType", m.GetItemType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7ParameterUserDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7ParameterUserDataItem")
	}
	return nil
}

func (m *_S7ParameterUserDataItem) IsS7ParameterUserDataItem() {}

func (m *_S7ParameterUserDataItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterUserDataItem) deepCopy() *_S7ParameterUserDataItem {
	if m == nil {
		return nil
	}
	_S7ParameterUserDataItemCopy := &_S7ParameterUserDataItem{
		nil, // will be set by child
	}
	return _S7ParameterUserDataItemCopy
}
