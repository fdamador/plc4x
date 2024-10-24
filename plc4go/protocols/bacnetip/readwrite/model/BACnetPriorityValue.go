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

// BACnetPriorityValue is the corresponding interface of BACnetPriorityValue
type BACnetPriorityValue interface {
	BACnetPriorityValueContract
	BACnetPriorityValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetPriorityValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValue()
	// CreateBuilder creates a BACnetPriorityValueBuilder
	CreateBACnetPriorityValueBuilder() BACnetPriorityValueBuilder
}

// BACnetPriorityValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetPriorityValueContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetObjectTypeArgument() returns a parser argument
	GetObjectTypeArgument() BACnetObjectType
	// IsBACnetPriorityValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValue()
	// CreateBuilder creates a BACnetPriorityValueBuilder
	CreateBACnetPriorityValueBuilder() BACnetPriorityValueBuilder
}

// BACnetPriorityValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetPriorityValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedIsContextTag returns PeekedIsContextTag (discriminator field)
	GetPeekedIsContextTag() bool
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetPriorityValue is the data-structure of this message
type _BACnetPriorityValue struct {
	_SubType interface {
		BACnetPriorityValueContract
		BACnetPriorityValueRequirements
	}
	PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetPriorityValueContract = (*_BACnetPriorityValue)(nil)

// NewBACnetPriorityValue factory function for _BACnetPriorityValue
func NewBACnetPriorityValue(peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValue {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetPriorityValue must not be nil")
	}
	return &_BACnetPriorityValue{PeekedTagHeader: peekedTagHeader, ObjectTypeArgument: objectTypeArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueBuilder is a builder for BACnetPriorityValue
type BACnetPriorityValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetPriorityValueBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetPriorityValueBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPriorityValueBuilder
	// AsBACnetPriorityValueNull converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueNull() interface {
		BACnetPriorityValueNullBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueReal converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueReal() interface {
		BACnetPriorityValueRealBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueEnumerated converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueEnumerated() interface {
		BACnetPriorityValueEnumeratedBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueUnsigned converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueUnsigned() interface {
		BACnetPriorityValueUnsignedBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueBoolean converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueBoolean() interface {
		BACnetPriorityValueBooleanBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueInteger converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueInteger() interface {
		BACnetPriorityValueIntegerBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueDouble converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueDouble() interface {
		BACnetPriorityValueDoubleBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueTime converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueTime() interface {
		BACnetPriorityValueTimeBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueCharacterString converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueCharacterString() interface {
		BACnetPriorityValueCharacterStringBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueOctetString converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueOctetString() interface {
		BACnetPriorityValueOctetStringBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueBitString converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueBitString() interface {
		BACnetPriorityValueBitStringBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueDate converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueDate() interface {
		BACnetPriorityValueDateBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueObjectidentifier converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueObjectidentifier() interface {
		BACnetPriorityValueObjectidentifierBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueConstructedValue converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueConstructedValue() interface {
		BACnetPriorityValueConstructedValueBuilder
		Done() BACnetPriorityValueBuilder
	}
	// AsBACnetPriorityValueDateTime converts this build to a subType of BACnetPriorityValue. It is always possible to return to current builder using Done()
	AsBACnetPriorityValueDateTime() interface {
		BACnetPriorityValueDateTimeBuilder
		Done() BACnetPriorityValueBuilder
	}
	// Build builds the BACnetPriorityValue or returns an error if something is wrong
	PartialBuild() (BACnetPriorityValueContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetPriorityValueContract
	// Build builds the BACnetPriorityValue or returns an error if something is wrong
	Build() (BACnetPriorityValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValue
}

// NewBACnetPriorityValueBuilder() creates a BACnetPriorityValueBuilder
func NewBACnetPriorityValueBuilder() BACnetPriorityValueBuilder {
	return &_BACnetPriorityValueBuilder{_BACnetPriorityValue: new(_BACnetPriorityValue)}
}

type _BACnetPriorityValueChildBuilder interface {
	utils.Copyable
	setParent(BACnetPriorityValueContract)
	buildForBACnetPriorityValue() (BACnetPriorityValue, error)
}

type _BACnetPriorityValueBuilder struct {
	*_BACnetPriorityValue

	childBuilder _BACnetPriorityValueChildBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueBuilder) = (*_BACnetPriorityValueBuilder)(nil)

func (b *_BACnetPriorityValueBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetPriorityValueBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetPriorityValueBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetPriorityValueBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetPriorityValueBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPriorityValueBuilder {
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

func (b *_BACnetPriorityValueBuilder) PartialBuild() (BACnetPriorityValueContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValue.deepCopy(), nil
}

func (b *_BACnetPriorityValueBuilder) PartialMustBuild() BACnetPriorityValueContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueNull() interface {
	BACnetPriorityValueNullBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueNullBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueNullBuilder().(*_BACnetPriorityValueNullBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueReal() interface {
	BACnetPriorityValueRealBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueRealBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueRealBuilder().(*_BACnetPriorityValueRealBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueEnumerated() interface {
	BACnetPriorityValueEnumeratedBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueEnumeratedBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueEnumeratedBuilder().(*_BACnetPriorityValueEnumeratedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueUnsigned() interface {
	BACnetPriorityValueUnsignedBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueUnsignedBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueUnsignedBuilder().(*_BACnetPriorityValueUnsignedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueBoolean() interface {
	BACnetPriorityValueBooleanBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueBooleanBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueBooleanBuilder().(*_BACnetPriorityValueBooleanBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueInteger() interface {
	BACnetPriorityValueIntegerBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueIntegerBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueIntegerBuilder().(*_BACnetPriorityValueIntegerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueDouble() interface {
	BACnetPriorityValueDoubleBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueDoubleBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueDoubleBuilder().(*_BACnetPriorityValueDoubleBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueTime() interface {
	BACnetPriorityValueTimeBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueTimeBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueTimeBuilder().(*_BACnetPriorityValueTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueCharacterString() interface {
	BACnetPriorityValueCharacterStringBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueCharacterStringBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueCharacterStringBuilder().(*_BACnetPriorityValueCharacterStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueOctetString() interface {
	BACnetPriorityValueOctetStringBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueOctetStringBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueOctetStringBuilder().(*_BACnetPriorityValueOctetStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueBitString() interface {
	BACnetPriorityValueBitStringBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueBitStringBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueBitStringBuilder().(*_BACnetPriorityValueBitStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueDate() interface {
	BACnetPriorityValueDateBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueDateBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueDateBuilder().(*_BACnetPriorityValueDateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueObjectidentifier() interface {
	BACnetPriorityValueObjectidentifierBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueObjectidentifierBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueObjectidentifierBuilder().(*_BACnetPriorityValueObjectidentifierBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueConstructedValue() interface {
	BACnetPriorityValueConstructedValueBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueConstructedValueBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueConstructedValueBuilder().(*_BACnetPriorityValueConstructedValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) AsBACnetPriorityValueDateTime() interface {
	BACnetPriorityValueDateTimeBuilder
	Done() BACnetPriorityValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetPriorityValueDateTimeBuilder
		Done() BACnetPriorityValueBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetPriorityValueDateTimeBuilder().(*_BACnetPriorityValueDateTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetPriorityValueBuilder) Build() (BACnetPriorityValue, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetPriorityValue()
}

func (b *_BACnetPriorityValueBuilder) MustBuild() BACnetPriorityValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetPriorityValueChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueBuilder creates a BACnetPriorityValueBuilder
func (b *_BACnetPriorityValue) CreateBACnetPriorityValueBuilder() BACnetPriorityValueBuilder {
	if b == nil {
		return NewBACnetPriorityValueBuilder()
	}
	return &_BACnetPriorityValueBuilder{_BACnetPriorityValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValue) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetPriorityValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (pm *_BACnetPriorityValue) GetPeekedIsContextTag() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValue(structType any) BACnetPriorityValue {
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValue) GetTypeName() string {
	return "BACnetPriorityValue"
}

func (m *_BACnetPriorityValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPriorityValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueParse[T BACnetPriorityValue](ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (T, error) {
	return BACnetPriorityValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueParseWithBufferProducer[T BACnetPriorityValue](objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetPriorityValueParseWithBuffer[T](ctx, readBuffer, objectTypeArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetPriorityValueParseWithBuffer[T BACnetPriorityValue](ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (T, error) {
	v, err := (&_BACnetPriorityValue{ObjectTypeArgument: objectTypeArgument}).parse(ctx, readBuffer, objectTypeArgument)
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

func (m *_BACnetPriorityValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (__bACnetPriorityValue BACnetPriorityValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValue")
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

	peekedIsContextTag, err := ReadVirtualField[bool](ctx, "peekedIsContextTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedIsContextTag' field"))
	}
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetPriorityValue
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetPriorityValueNull
		if _child, err = new(_BACnetPriorityValueNull).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueNull for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetPriorityValueReal
		if _child, err = new(_BACnetPriorityValueReal).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueReal for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetPriorityValueEnumerated
		if _child, err = new(_BACnetPriorityValueEnumerated).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueEnumerated for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetPriorityValueUnsigned
		if _child, err = new(_BACnetPriorityValueUnsigned).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueUnsigned for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetPriorityValueBoolean
		if _child, err = new(_BACnetPriorityValueBoolean).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueBoolean for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetPriorityValueInteger
		if _child, err = new(_BACnetPriorityValueInteger).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueInteger for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetPriorityValueDouble
		if _child, err = new(_BACnetPriorityValueDouble).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDouble for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetPriorityValueTime
		if _child, err = new(_BACnetPriorityValueTime).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueTime for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetPriorityValueCharacterString
		if _child, err = new(_BACnetPriorityValueCharacterString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueCharacterString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetPriorityValueOctetString
		if _child, err = new(_BACnetPriorityValueOctetString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueOctetString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetPriorityValueBitString
		if _child, err = new(_BACnetPriorityValueBitString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueBitString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetPriorityValueDate
		if _child, err = new(_BACnetPriorityValueDate).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDate for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetPriorityValueObjectidentifier
		if _child, err = new(_BACnetPriorityValueObjectidentifier).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueObjectidentifier for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetPriorityValueConstructedValue
		if _child, err = new(_BACnetPriorityValueConstructedValue).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueConstructedValue for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == uint8(1) && peekedIsContextTag == bool(true): // BACnetPriorityValueDateTime
		if _child, err = new(_BACnetPriorityValueDateTime).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDateTime for type-switch of BACnetPriorityValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValue")
	}

	return _child, nil
}

func (pm *_BACnetPriorityValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPriorityValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPriorityValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValue")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ = peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPriorityValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPriorityValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPriorityValue) IsBACnetPriorityValue() {}

func (m *_BACnetPriorityValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValue) deepCopy() *_BACnetPriorityValue {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueCopy := &_BACnetPriorityValue{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.ObjectTypeArgument,
	}
	return _BACnetPriorityValueCopy
}
