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

// StructureDefinition is the corresponding interface of StructureDefinition
type StructureDefinition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDefaultEncodingId returns DefaultEncodingId (property field)
	GetDefaultEncodingId() NodeId
	// GetBaseDataType returns BaseDataType (property field)
	GetBaseDataType() NodeId
	// GetStructureType returns StructureType (property field)
	GetStructureType() StructureType
	// GetFields returns Fields (property field)
	GetFields() []StructureField
	// IsStructureDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStructureDefinition()
	// CreateBuilder creates a StructureDefinitionBuilder
	CreateStructureDefinitionBuilder() StructureDefinitionBuilder
}

// _StructureDefinition is the data-structure of this message
type _StructureDefinition struct {
	ExtensionObjectDefinitionContract
	DefaultEncodingId NodeId
	BaseDataType      NodeId
	StructureType     StructureType
	Fields            []StructureField
}

var _ StructureDefinition = (*_StructureDefinition)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_StructureDefinition)(nil)

// NewStructureDefinition factory function for _StructureDefinition
func NewStructureDefinition(defaultEncodingId NodeId, baseDataType NodeId, structureType StructureType, fields []StructureField) *_StructureDefinition {
	if defaultEncodingId == nil {
		panic("defaultEncodingId of type NodeId for StructureDefinition must not be nil")
	}
	if baseDataType == nil {
		panic("baseDataType of type NodeId for StructureDefinition must not be nil")
	}
	_result := &_StructureDefinition{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DefaultEncodingId:                 defaultEncodingId,
		BaseDataType:                      baseDataType,
		StructureType:                     structureType,
		Fields:                            fields,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StructureDefinitionBuilder is a builder for StructureDefinition
type StructureDefinitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(defaultEncodingId NodeId, baseDataType NodeId, structureType StructureType, fields []StructureField) StructureDefinitionBuilder
	// WithDefaultEncodingId adds DefaultEncodingId (property field)
	WithDefaultEncodingId(NodeId) StructureDefinitionBuilder
	// WithDefaultEncodingIdBuilder adds DefaultEncodingId (property field) which is build by the builder
	WithDefaultEncodingIdBuilder(func(NodeIdBuilder) NodeIdBuilder) StructureDefinitionBuilder
	// WithBaseDataType adds BaseDataType (property field)
	WithBaseDataType(NodeId) StructureDefinitionBuilder
	// WithBaseDataTypeBuilder adds BaseDataType (property field) which is build by the builder
	WithBaseDataTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) StructureDefinitionBuilder
	// WithStructureType adds StructureType (property field)
	WithStructureType(StructureType) StructureDefinitionBuilder
	// WithFields adds Fields (property field)
	WithFields(...StructureField) StructureDefinitionBuilder
	// Build builds the StructureDefinition or returns an error if something is wrong
	Build() (StructureDefinition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StructureDefinition
}

// NewStructureDefinitionBuilder() creates a StructureDefinitionBuilder
func NewStructureDefinitionBuilder() StructureDefinitionBuilder {
	return &_StructureDefinitionBuilder{_StructureDefinition: new(_StructureDefinition)}
}

type _StructureDefinitionBuilder struct {
	*_StructureDefinition

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (StructureDefinitionBuilder) = (*_StructureDefinitionBuilder)(nil)

func (b *_StructureDefinitionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_StructureDefinitionBuilder) WithMandatoryFields(defaultEncodingId NodeId, baseDataType NodeId, structureType StructureType, fields []StructureField) StructureDefinitionBuilder {
	return b.WithDefaultEncodingId(defaultEncodingId).WithBaseDataType(baseDataType).WithStructureType(structureType).WithFields(fields...)
}

func (b *_StructureDefinitionBuilder) WithDefaultEncodingId(defaultEncodingId NodeId) StructureDefinitionBuilder {
	b.DefaultEncodingId = defaultEncodingId
	return b
}

func (b *_StructureDefinitionBuilder) WithDefaultEncodingIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) StructureDefinitionBuilder {
	builder := builderSupplier(b.DefaultEncodingId.CreateNodeIdBuilder())
	var err error
	b.DefaultEncodingId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_StructureDefinitionBuilder) WithBaseDataType(baseDataType NodeId) StructureDefinitionBuilder {
	b.BaseDataType = baseDataType
	return b
}

func (b *_StructureDefinitionBuilder) WithBaseDataTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) StructureDefinitionBuilder {
	builder := builderSupplier(b.BaseDataType.CreateNodeIdBuilder())
	var err error
	b.BaseDataType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_StructureDefinitionBuilder) WithStructureType(structureType StructureType) StructureDefinitionBuilder {
	b.StructureType = structureType
	return b
}

func (b *_StructureDefinitionBuilder) WithFields(fields ...StructureField) StructureDefinitionBuilder {
	b.Fields = fields
	return b
}

func (b *_StructureDefinitionBuilder) Build() (StructureDefinition, error) {
	if b.DefaultEncodingId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'defaultEncodingId' not set"))
	}
	if b.BaseDataType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'baseDataType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StructureDefinition.deepCopy(), nil
}

func (b *_StructureDefinitionBuilder) MustBuild() StructureDefinition {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_StructureDefinitionBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_StructureDefinitionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_StructureDefinitionBuilder) DeepCopy() any {
	_copy := b.CreateStructureDefinitionBuilder().(*_StructureDefinitionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStructureDefinitionBuilder creates a StructureDefinitionBuilder
func (b *_StructureDefinition) CreateStructureDefinitionBuilder() StructureDefinitionBuilder {
	if b == nil {
		return NewStructureDefinitionBuilder()
	}
	return &_StructureDefinitionBuilder{_StructureDefinition: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StructureDefinition) GetExtensionId() int32 {
	return int32(101)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StructureDefinition) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StructureDefinition) GetDefaultEncodingId() NodeId {
	return m.DefaultEncodingId
}

func (m *_StructureDefinition) GetBaseDataType() NodeId {
	return m.BaseDataType
}

func (m *_StructureDefinition) GetStructureType() StructureType {
	return m.StructureType
}

func (m *_StructureDefinition) GetFields() []StructureField {
	return m.Fields
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStructureDefinition(structType any) StructureDefinition {
	if casted, ok := structType.(StructureDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*StructureDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_StructureDefinition) GetTypeName() string {
	return "StructureDefinition"
}

func (m *_StructureDefinition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (defaultEncodingId)
	lengthInBits += m.DefaultEncodingId.GetLengthInBits(ctx)

	// Simple field (baseDataType)
	lengthInBits += m.BaseDataType.GetLengthInBits(ctx)

	// Simple field (structureType)
	lengthInBits += 32

	// Implicit Field (noOfFields)
	lengthInBits += 32

	// Array field
	if len(m.Fields) > 0 {
		for _curItem, element := range m.Fields {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Fields), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_StructureDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StructureDefinition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__structureDefinition StructureDefinition, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StructureDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StructureDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	defaultEncodingId, err := ReadSimpleField[NodeId](ctx, "defaultEncodingId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'defaultEncodingId' field"))
	}
	m.DefaultEncodingId = defaultEncodingId

	baseDataType, err := ReadSimpleField[NodeId](ctx, "baseDataType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'baseDataType' field"))
	}
	m.BaseDataType = baseDataType

	structureType, err := ReadEnumField[StructureType](ctx, "structureType", "StructureType", ReadEnum(StructureTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureType' field"))
	}
	m.StructureType = structureType

	noOfFields, err := ReadImplicitField[int32](ctx, "noOfFields", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfFields' field"))
	}
	_ = noOfFields

	fields, err := ReadCountArrayField[StructureField](ctx, "fields", ReadComplex[StructureField](ExtensionObjectDefinitionParseWithBufferProducer[StructureField]((int32)(int32(103))), readBuffer), uint64(noOfFields))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fields' field"))
	}
	m.Fields = fields

	if closeErr := readBuffer.CloseContext("StructureDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StructureDefinition")
	}

	return m, nil
}

func (m *_StructureDefinition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StructureDefinition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StructureDefinition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StructureDefinition")
		}

		if err := WriteSimpleField[NodeId](ctx, "defaultEncodingId", m.GetDefaultEncodingId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'defaultEncodingId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "baseDataType", m.GetBaseDataType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'baseDataType' field")
		}

		if err := WriteSimpleEnumField[StructureType](ctx, "structureType", "StructureType", m.GetStructureType(), WriteEnum[StructureType, uint32](StructureType.GetValue, StructureType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'structureType' field")
		}
		noOfFields := int32(utils.InlineIf(bool((m.GetFields()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetFields()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfFields", noOfFields, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfFields' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "fields", m.GetFields(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'fields' field")
		}

		if popErr := writeBuffer.PopContext("StructureDefinition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StructureDefinition")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StructureDefinition) IsStructureDefinition() {}

func (m *_StructureDefinition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StructureDefinition) deepCopy() *_StructureDefinition {
	if m == nil {
		return nil
	}
	_StructureDefinitionCopy := &_StructureDefinition{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.DefaultEncodingId.DeepCopy().(NodeId),
		m.BaseDataType.DeepCopy().(NodeId),
		m.StructureType,
		utils.DeepCopySlice[StructureField, StructureField](m.Fields),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _StructureDefinitionCopy
}

func (m *_StructureDefinition) String() string {
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
