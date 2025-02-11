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

// GenericAttributes is the corresponding interface of GenericAttributes
type GenericAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSpecifiedAttributes returns SpecifiedAttributes (property field)
	GetSpecifiedAttributes() uint32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
	// GetAttributeValues returns AttributeValues (property field)
	GetAttributeValues() []GenericAttributeValue
	// IsGenericAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGenericAttributes()
	// CreateBuilder creates a GenericAttributesBuilder
	CreateGenericAttributesBuilder() GenericAttributesBuilder
}

// _GenericAttributes is the data-structure of this message
type _GenericAttributes struct {
	ExtensionObjectDefinitionContract
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
	AttributeValues     []GenericAttributeValue
}

var _ GenericAttributes = (*_GenericAttributes)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_GenericAttributes)(nil)

// NewGenericAttributes factory function for _GenericAttributes
func NewGenericAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, attributeValues []GenericAttributeValue) *_GenericAttributes {
	if displayName == nil {
		panic("displayName of type LocalizedText for GenericAttributes must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for GenericAttributes must not be nil")
	}
	_result := &_GenericAttributes{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedAttributes:               specifiedAttributes,
		DisplayName:                       displayName,
		Description:                       description,
		WriteMask:                         writeMask,
		UserWriteMask:                     userWriteMask,
		AttributeValues:                   attributeValues,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GenericAttributesBuilder is a builder for GenericAttributes
type GenericAttributesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, attributeValues []GenericAttributeValue) GenericAttributesBuilder
	// WithSpecifiedAttributes adds SpecifiedAttributes (property field)
	WithSpecifiedAttributes(uint32) GenericAttributesBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) GenericAttributesBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) GenericAttributesBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) GenericAttributesBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) GenericAttributesBuilder
	// WithWriteMask adds WriteMask (property field)
	WithWriteMask(uint32) GenericAttributesBuilder
	// WithUserWriteMask adds UserWriteMask (property field)
	WithUserWriteMask(uint32) GenericAttributesBuilder
	// WithAttributeValues adds AttributeValues (property field)
	WithAttributeValues(...GenericAttributeValue) GenericAttributesBuilder
	// Build builds the GenericAttributes or returns an error if something is wrong
	Build() (GenericAttributes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GenericAttributes
}

// NewGenericAttributesBuilder() creates a GenericAttributesBuilder
func NewGenericAttributesBuilder() GenericAttributesBuilder {
	return &_GenericAttributesBuilder{_GenericAttributes: new(_GenericAttributes)}
}

type _GenericAttributesBuilder struct {
	*_GenericAttributes

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (GenericAttributesBuilder) = (*_GenericAttributesBuilder)(nil)

func (b *_GenericAttributesBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_GenericAttributesBuilder) WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, attributeValues []GenericAttributeValue) GenericAttributesBuilder {
	return b.WithSpecifiedAttributes(specifiedAttributes).WithDisplayName(displayName).WithDescription(description).WithWriteMask(writeMask).WithUserWriteMask(userWriteMask).WithAttributeValues(attributeValues...)
}

func (b *_GenericAttributesBuilder) WithSpecifiedAttributes(specifiedAttributes uint32) GenericAttributesBuilder {
	b.SpecifiedAttributes = specifiedAttributes
	return b
}

func (b *_GenericAttributesBuilder) WithDisplayName(displayName LocalizedText) GenericAttributesBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_GenericAttributesBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) GenericAttributesBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_GenericAttributesBuilder) WithDescription(description LocalizedText) GenericAttributesBuilder {
	b.Description = description
	return b
}

func (b *_GenericAttributesBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) GenericAttributesBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_GenericAttributesBuilder) WithWriteMask(writeMask uint32) GenericAttributesBuilder {
	b.WriteMask = writeMask
	return b
}

func (b *_GenericAttributesBuilder) WithUserWriteMask(userWriteMask uint32) GenericAttributesBuilder {
	b.UserWriteMask = userWriteMask
	return b
}

func (b *_GenericAttributesBuilder) WithAttributeValues(attributeValues ...GenericAttributeValue) GenericAttributesBuilder {
	b.AttributeValues = attributeValues
	return b
}

func (b *_GenericAttributesBuilder) Build() (GenericAttributes, error) {
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GenericAttributes.deepCopy(), nil
}

func (b *_GenericAttributesBuilder) MustBuild() GenericAttributes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_GenericAttributesBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_GenericAttributesBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_GenericAttributesBuilder) DeepCopy() any {
	_copy := b.CreateGenericAttributesBuilder().(*_GenericAttributesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGenericAttributesBuilder creates a GenericAttributesBuilder
func (b *_GenericAttributes) CreateGenericAttributesBuilder() GenericAttributesBuilder {
	if b == nil {
		return NewGenericAttributesBuilder()
	}
	return &_GenericAttributesBuilder{_GenericAttributes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GenericAttributes) GetExtensionId() int32 {
	return int32(17609)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GenericAttributes) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GenericAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_GenericAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_GenericAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_GenericAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_GenericAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

func (m *_GenericAttributes) GetAttributeValues() []GenericAttributeValue {
	return m.AttributeValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGenericAttributes(structType any) GenericAttributes {
	if casted, ok := structType.(GenericAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*GenericAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_GenericAttributes) GetTypeName() string {
	return "GenericAttributes"
}

func (m *_GenericAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (specifiedAttributes)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	// Implicit Field (noOfAttributeValues)
	lengthInBits += 32

	// Array field
	if len(m.AttributeValues) > 0 {
		for _curItem, element := range m.AttributeValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AttributeValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_GenericAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GenericAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__genericAttributes GenericAttributes, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GenericAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GenericAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}
	m.SpecifiedAttributes = specifiedAttributes

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}
	m.WriteMask = writeMask

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}
	m.UserWriteMask = userWriteMask

	noOfAttributeValues, err := ReadImplicitField[int32](ctx, "noOfAttributeValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfAttributeValues' field"))
	}
	_ = noOfAttributeValues

	attributeValues, err := ReadCountArrayField[GenericAttributeValue](ctx, "attributeValues", ReadComplex[GenericAttributeValue](ExtensionObjectDefinitionParseWithBufferProducer[GenericAttributeValue]((int32)(int32(17608))), readBuffer), uint64(noOfAttributeValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeValues' field"))
	}
	m.AttributeValues = attributeValues

	if closeErr := readBuffer.CloseContext("GenericAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GenericAttributes")
	}

	return m, nil
}

func (m *_GenericAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GenericAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GenericAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GenericAttributes")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedAttributes", m.GetSpecifiedAttributes(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedAttributes' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleField[uint32](ctx, "writeMask", m.GetWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "userWriteMask", m.GetUserWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'userWriteMask' field")
		}
		noOfAttributeValues := int32(utils.InlineIf(bool((m.GetAttributeValues()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetAttributeValues()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfAttributeValues", noOfAttributeValues, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfAttributeValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "attributeValues", m.GetAttributeValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeValues' field")
		}

		if popErr := writeBuffer.PopContext("GenericAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GenericAttributes")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GenericAttributes) IsGenericAttributes() {}

func (m *_GenericAttributes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GenericAttributes) deepCopy() *_GenericAttributes {
	if m == nil {
		return nil
	}
	_GenericAttributesCopy := &_GenericAttributes{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SpecifiedAttributes,
		m.DisplayName.DeepCopy().(LocalizedText),
		m.Description.DeepCopy().(LocalizedText),
		m.WriteMask,
		m.UserWriteMask,
		utils.DeepCopySlice[GenericAttributeValue, GenericAttributeValue](m.AttributeValues),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _GenericAttributesCopy
}

func (m *_GenericAttributes) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
