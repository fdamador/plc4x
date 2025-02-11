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

// BACnetValueSource is the corresponding interface of BACnetValueSource
type BACnetValueSource interface {
	BACnetValueSourceContract
	BACnetValueSourceRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetValueSource is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSource()
	// CreateBuilder creates a BACnetValueSourceBuilder
	CreateBACnetValueSourceBuilder() BACnetValueSourceBuilder
}

// BACnetValueSourceContract provides a set of functions which can be overwritten by a sub struct
type BACnetValueSourceContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetValueSource is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSource()
	// CreateBuilder creates a BACnetValueSourceBuilder
	CreateBACnetValueSourceBuilder() BACnetValueSourceBuilder
}

// BACnetValueSourceRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetValueSourceRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetValueSource is the data-structure of this message
type _BACnetValueSource struct {
	_SubType interface {
		BACnetValueSourceContract
		BACnetValueSourceRequirements
	}
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetValueSourceContract = (*_BACnetValueSource)(nil)

// NewBACnetValueSource factory function for _BACnetValueSource
func NewBACnetValueSource(peekedTagHeader BACnetTagHeader) *_BACnetValueSource {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetValueSource must not be nil")
	}
	return &_BACnetValueSource{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetValueSourceBuilder is a builder for BACnetValueSource
type BACnetValueSourceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetValueSourceBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetValueSourceBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetValueSourceBuilder
	// AsBACnetValueSourceNone converts this build to a subType of BACnetValueSource. It is always possible to return to current builder using Done()
	AsBACnetValueSourceNone() interface {
		BACnetValueSourceNoneBuilder
		Done() BACnetValueSourceBuilder
	}
	// AsBACnetValueSourceObject converts this build to a subType of BACnetValueSource. It is always possible to return to current builder using Done()
	AsBACnetValueSourceObject() interface {
		BACnetValueSourceObjectBuilder
		Done() BACnetValueSourceBuilder
	}
	// AsBACnetValueSourceAddress converts this build to a subType of BACnetValueSource. It is always possible to return to current builder using Done()
	AsBACnetValueSourceAddress() interface {
		BACnetValueSourceAddressBuilder
		Done() BACnetValueSourceBuilder
	}
	// Build builds the BACnetValueSource or returns an error if something is wrong
	PartialBuild() (BACnetValueSourceContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetValueSourceContract
	// Build builds the BACnetValueSource or returns an error if something is wrong
	Build() (BACnetValueSource, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetValueSource
}

// NewBACnetValueSourceBuilder() creates a BACnetValueSourceBuilder
func NewBACnetValueSourceBuilder() BACnetValueSourceBuilder {
	return &_BACnetValueSourceBuilder{_BACnetValueSource: new(_BACnetValueSource)}
}

type _BACnetValueSourceChildBuilder interface {
	utils.Copyable
	setParent(BACnetValueSourceContract)
	buildForBACnetValueSource() (BACnetValueSource, error)
}

type _BACnetValueSourceBuilder struct {
	*_BACnetValueSource

	childBuilder _BACnetValueSourceChildBuilder

	err *utils.MultiError
}

var _ (BACnetValueSourceBuilder) = (*_BACnetValueSourceBuilder)(nil)

func (b *_BACnetValueSourceBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetValueSourceBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetValueSourceBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetValueSourceBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetValueSourceBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetValueSourceBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetValueSourceBuilder) PartialBuild() (BACnetValueSourceContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetValueSource.deepCopy(), nil
}

func (b *_BACnetValueSourceBuilder) PartialMustBuild() BACnetValueSourceContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetValueSourceBuilder) AsBACnetValueSourceNone() interface {
	BACnetValueSourceNoneBuilder
	Done() BACnetValueSourceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetValueSourceNoneBuilder
		Done() BACnetValueSourceBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetValueSourceNoneBuilder().(*_BACnetValueSourceNoneBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetValueSourceBuilder) AsBACnetValueSourceObject() interface {
	BACnetValueSourceObjectBuilder
	Done() BACnetValueSourceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetValueSourceObjectBuilder
		Done() BACnetValueSourceBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetValueSourceObjectBuilder().(*_BACnetValueSourceObjectBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetValueSourceBuilder) AsBACnetValueSourceAddress() interface {
	BACnetValueSourceAddressBuilder
	Done() BACnetValueSourceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetValueSourceAddressBuilder
		Done() BACnetValueSourceBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetValueSourceAddressBuilder().(*_BACnetValueSourceAddressBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetValueSourceBuilder) Build() (BACnetValueSource, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetValueSource()
}

func (b *_BACnetValueSourceBuilder) MustBuild() BACnetValueSource {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetValueSourceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetValueSourceBuilder().(*_BACnetValueSourceBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetValueSourceChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetValueSourceBuilder creates a BACnetValueSourceBuilder
func (b *_BACnetValueSource) CreateBACnetValueSourceBuilder() BACnetValueSourceBuilder {
	if b == nil {
		return NewBACnetValueSourceBuilder()
	}
	return &_BACnetValueSourceBuilder{_BACnetValueSource: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSource) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetValueSource) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetValueSource(structType any) BACnetValueSource {
	if casted, ok := structType.(BACnetValueSource); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSource); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSource) GetTypeName() string {
	return "BACnetValueSource"
}

func (m *_BACnetValueSource) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetValueSource) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetValueSource) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceParse[T BACnetValueSource](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetValueSourceParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceParseWithBufferProducer[T BACnetValueSource]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetValueSourceParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetValueSourceParseWithBuffer[T BACnetValueSource](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetValueSource{}).parse(ctx, readBuffer)
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

func (m *_BACnetValueSource) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetValueSource BACnetValueSource, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSource")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetValueSource
	switch {
	case peekedTagNumber == uint8(0): // BACnetValueSourceNone
		if _child, err = new(_BACnetValueSourceNone).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceNone for type-switch of BACnetValueSource")
		}
	case peekedTagNumber == uint8(1): // BACnetValueSourceObject
		if _child, err = new(_BACnetValueSourceObject).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceObject for type-switch of BACnetValueSource")
		}
	case peekedTagNumber == uint8(2): // BACnetValueSourceAddress
		if _child, err = new(_BACnetValueSourceAddress).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceAddress for type-switch of BACnetValueSource")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSource")
	}

	return _child, nil
}

func (pm *_BACnetValueSource) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetValueSource, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetValueSource"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetValueSource")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetValueSource"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetValueSource")
	}
	return nil
}

func (m *_BACnetValueSource) IsBACnetValueSource() {}

func (m *_BACnetValueSource) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetValueSource) deepCopy() *_BACnetValueSource {
	if m == nil {
		return nil
	}
	_BACnetValueSourceCopy := &_BACnetValueSource{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetValueSourceCopy
}
