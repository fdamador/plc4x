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

// SimpleAttributeOperand is the corresponding interface of SimpleAttributeOperand
type SimpleAttributeOperand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetTypeDefinitionId returns TypeDefinitionId (property field)
	GetTypeDefinitionId() NodeId
	// GetBrowsePath returns BrowsePath (property field)
	GetBrowsePath() []QualifiedName
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// IsSimpleAttributeOperand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSimpleAttributeOperand()
	// CreateBuilder creates a SimpleAttributeOperandBuilder
	CreateSimpleAttributeOperandBuilder() SimpleAttributeOperandBuilder
}

// _SimpleAttributeOperand is the data-structure of this message
type _SimpleAttributeOperand struct {
	ExtensionObjectDefinitionContract
	TypeDefinitionId NodeId
	BrowsePath       []QualifiedName
	AttributeId      uint32
	IndexRange       PascalString
}

var _ SimpleAttributeOperand = (*_SimpleAttributeOperand)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SimpleAttributeOperand)(nil)

// NewSimpleAttributeOperand factory function for _SimpleAttributeOperand
func NewSimpleAttributeOperand(typeDefinitionId NodeId, browsePath []QualifiedName, attributeId uint32, indexRange PascalString) *_SimpleAttributeOperand {
	if typeDefinitionId == nil {
		panic("typeDefinitionId of type NodeId for SimpleAttributeOperand must not be nil")
	}
	if indexRange == nil {
		panic("indexRange of type PascalString for SimpleAttributeOperand must not be nil")
	}
	_result := &_SimpleAttributeOperand{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TypeDefinitionId:                  typeDefinitionId,
		BrowsePath:                        browsePath,
		AttributeId:                       attributeId,
		IndexRange:                        indexRange,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SimpleAttributeOperandBuilder is a builder for SimpleAttributeOperand
type SimpleAttributeOperandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(typeDefinitionId NodeId, browsePath []QualifiedName, attributeId uint32, indexRange PascalString) SimpleAttributeOperandBuilder
	// WithTypeDefinitionId adds TypeDefinitionId (property field)
	WithTypeDefinitionId(NodeId) SimpleAttributeOperandBuilder
	// WithTypeDefinitionIdBuilder adds TypeDefinitionId (property field) which is build by the builder
	WithTypeDefinitionIdBuilder(func(NodeIdBuilder) NodeIdBuilder) SimpleAttributeOperandBuilder
	// WithBrowsePath adds BrowsePath (property field)
	WithBrowsePath(...QualifiedName) SimpleAttributeOperandBuilder
	// WithAttributeId adds AttributeId (property field)
	WithAttributeId(uint32) SimpleAttributeOperandBuilder
	// WithIndexRange adds IndexRange (property field)
	WithIndexRange(PascalString) SimpleAttributeOperandBuilder
	// WithIndexRangeBuilder adds IndexRange (property field) which is build by the builder
	WithIndexRangeBuilder(func(PascalStringBuilder) PascalStringBuilder) SimpleAttributeOperandBuilder
	// Build builds the SimpleAttributeOperand or returns an error if something is wrong
	Build() (SimpleAttributeOperand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SimpleAttributeOperand
}

// NewSimpleAttributeOperandBuilder() creates a SimpleAttributeOperandBuilder
func NewSimpleAttributeOperandBuilder() SimpleAttributeOperandBuilder {
	return &_SimpleAttributeOperandBuilder{_SimpleAttributeOperand: new(_SimpleAttributeOperand)}
}

type _SimpleAttributeOperandBuilder struct {
	*_SimpleAttributeOperand

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SimpleAttributeOperandBuilder) = (*_SimpleAttributeOperandBuilder)(nil)

func (b *_SimpleAttributeOperandBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_SimpleAttributeOperandBuilder) WithMandatoryFields(typeDefinitionId NodeId, browsePath []QualifiedName, attributeId uint32, indexRange PascalString) SimpleAttributeOperandBuilder {
	return b.WithTypeDefinitionId(typeDefinitionId).WithBrowsePath(browsePath...).WithAttributeId(attributeId).WithIndexRange(indexRange)
}

func (b *_SimpleAttributeOperandBuilder) WithTypeDefinitionId(typeDefinitionId NodeId) SimpleAttributeOperandBuilder {
	b.TypeDefinitionId = typeDefinitionId
	return b
}

func (b *_SimpleAttributeOperandBuilder) WithTypeDefinitionIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) SimpleAttributeOperandBuilder {
	builder := builderSupplier(b.TypeDefinitionId.CreateNodeIdBuilder())
	var err error
	b.TypeDefinitionId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_SimpleAttributeOperandBuilder) WithBrowsePath(browsePath ...QualifiedName) SimpleAttributeOperandBuilder {
	b.BrowsePath = browsePath
	return b
}

func (b *_SimpleAttributeOperandBuilder) WithAttributeId(attributeId uint32) SimpleAttributeOperandBuilder {
	b.AttributeId = attributeId
	return b
}

func (b *_SimpleAttributeOperandBuilder) WithIndexRange(indexRange PascalString) SimpleAttributeOperandBuilder {
	b.IndexRange = indexRange
	return b
}

func (b *_SimpleAttributeOperandBuilder) WithIndexRangeBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) SimpleAttributeOperandBuilder {
	builder := builderSupplier(b.IndexRange.CreatePascalStringBuilder())
	var err error
	b.IndexRange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_SimpleAttributeOperandBuilder) Build() (SimpleAttributeOperand, error) {
	if b.TypeDefinitionId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'typeDefinitionId' not set"))
	}
	if b.IndexRange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'indexRange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SimpleAttributeOperand.deepCopy(), nil
}

func (b *_SimpleAttributeOperandBuilder) MustBuild() SimpleAttributeOperand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SimpleAttributeOperandBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_SimpleAttributeOperandBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SimpleAttributeOperandBuilder) DeepCopy() any {
	_copy := b.CreateSimpleAttributeOperandBuilder().(*_SimpleAttributeOperandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSimpleAttributeOperandBuilder creates a SimpleAttributeOperandBuilder
func (b *_SimpleAttributeOperand) CreateSimpleAttributeOperandBuilder() SimpleAttributeOperandBuilder {
	if b == nil {
		return NewSimpleAttributeOperandBuilder()
	}
	return &_SimpleAttributeOperandBuilder{_SimpleAttributeOperand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SimpleAttributeOperand) GetExtensionId() int32 {
	return int32(603)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SimpleAttributeOperand) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SimpleAttributeOperand) GetTypeDefinitionId() NodeId {
	return m.TypeDefinitionId
}

func (m *_SimpleAttributeOperand) GetBrowsePath() []QualifiedName {
	return m.BrowsePath
}

func (m *_SimpleAttributeOperand) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_SimpleAttributeOperand) GetIndexRange() PascalString {
	return m.IndexRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSimpleAttributeOperand(structType any) SimpleAttributeOperand {
	if casted, ok := structType.(SimpleAttributeOperand); ok {
		return casted
	}
	if casted, ok := structType.(*SimpleAttributeOperand); ok {
		return *casted
	}
	return nil
}

func (m *_SimpleAttributeOperand) GetTypeName() string {
	return "SimpleAttributeOperand"
}

func (m *_SimpleAttributeOperand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (typeDefinitionId)
	lengthInBits += m.TypeDefinitionId.GetLengthInBits(ctx)

	// Implicit Field (noOfBrowsePath)
	lengthInBits += 32

	// Array field
	if len(m.BrowsePath) > 0 {
		for _curItem, element := range m.BrowsePath {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.BrowsePath), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SimpleAttributeOperand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SimpleAttributeOperand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__simpleAttributeOperand SimpleAttributeOperand, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SimpleAttributeOperand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SimpleAttributeOperand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	typeDefinitionId, err := ReadSimpleField[NodeId](ctx, "typeDefinitionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinitionId' field"))
	}
	m.TypeDefinitionId = typeDefinitionId

	noOfBrowsePath, err := ReadImplicitField[int32](ctx, "noOfBrowsePath", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfBrowsePath' field"))
	}
	_ = noOfBrowsePath

	browsePath, err := ReadCountArrayField[QualifiedName](ctx, "browsePath", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer), uint64(noOfBrowsePath))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browsePath' field"))
	}
	m.BrowsePath = browsePath

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	indexRange, err := ReadSimpleField[PascalString](ctx, "indexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexRange' field"))
	}
	m.IndexRange = indexRange

	if closeErr := readBuffer.CloseContext("SimpleAttributeOperand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SimpleAttributeOperand")
	}

	return m, nil
}

func (m *_SimpleAttributeOperand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SimpleAttributeOperand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SimpleAttributeOperand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SimpleAttributeOperand")
		}

		if err := WriteSimpleField[NodeId](ctx, "typeDefinitionId", m.GetTypeDefinitionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinitionId' field")
		}
		noOfBrowsePath := int32(utils.InlineIf(bool((m.GetBrowsePath()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetBrowsePath()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfBrowsePath", noOfBrowsePath, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfBrowsePath' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "browsePath", m.GetBrowsePath(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'browsePath' field")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "indexRange", m.GetIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexRange' field")
		}

		if popErr := writeBuffer.PopContext("SimpleAttributeOperand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SimpleAttributeOperand")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SimpleAttributeOperand) IsSimpleAttributeOperand() {}

func (m *_SimpleAttributeOperand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SimpleAttributeOperand) deepCopy() *_SimpleAttributeOperand {
	if m == nil {
		return nil
	}
	_SimpleAttributeOperandCopy := &_SimpleAttributeOperand{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.TypeDefinitionId.DeepCopy().(NodeId),
		utils.DeepCopySlice[QualifiedName, QualifiedName](m.BrowsePath),
		m.AttributeId,
		m.IndexRange.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SimpleAttributeOperandCopy
}

func (m *_SimpleAttributeOperand) String() string {
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
