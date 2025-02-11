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

// RelativePath is the corresponding interface of RelativePath
type RelativePath interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetElements returns Elements (property field)
	GetElements() []RelativePathElement
	// IsRelativePath is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRelativePath()
	// CreateBuilder creates a RelativePathBuilder
	CreateRelativePathBuilder() RelativePathBuilder
}

// _RelativePath is the data-structure of this message
type _RelativePath struct {
	ExtensionObjectDefinitionContract
	Elements []RelativePathElement
}

var _ RelativePath = (*_RelativePath)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RelativePath)(nil)

// NewRelativePath factory function for _RelativePath
func NewRelativePath(elements []RelativePathElement) *_RelativePath {
	_result := &_RelativePath{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Elements:                          elements,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RelativePathBuilder is a builder for RelativePath
type RelativePathBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(elements []RelativePathElement) RelativePathBuilder
	// WithElements adds Elements (property field)
	WithElements(...RelativePathElement) RelativePathBuilder
	// Build builds the RelativePath or returns an error if something is wrong
	Build() (RelativePath, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RelativePath
}

// NewRelativePathBuilder() creates a RelativePathBuilder
func NewRelativePathBuilder() RelativePathBuilder {
	return &_RelativePathBuilder{_RelativePath: new(_RelativePath)}
}

type _RelativePathBuilder struct {
	*_RelativePath

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RelativePathBuilder) = (*_RelativePathBuilder)(nil)

func (b *_RelativePathBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RelativePathBuilder) WithMandatoryFields(elements []RelativePathElement) RelativePathBuilder {
	return b.WithElements(elements...)
}

func (b *_RelativePathBuilder) WithElements(elements ...RelativePathElement) RelativePathBuilder {
	b.Elements = elements
	return b
}

func (b *_RelativePathBuilder) Build() (RelativePath, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RelativePath.deepCopy(), nil
}

func (b *_RelativePathBuilder) MustBuild() RelativePath {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RelativePathBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RelativePathBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RelativePathBuilder) DeepCopy() any {
	_copy := b.CreateRelativePathBuilder().(*_RelativePathBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRelativePathBuilder creates a RelativePathBuilder
func (b *_RelativePath) CreateRelativePathBuilder() RelativePathBuilder {
	if b == nil {
		return NewRelativePathBuilder()
	}
	return &_RelativePathBuilder{_RelativePath: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RelativePath) GetExtensionId() int32 {
	return int32(542)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RelativePath) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativePath) GetElements() []RelativePathElement {
	return m.Elements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRelativePath(structType any) RelativePath {
	if casted, ok := structType.(RelativePath); ok {
		return casted
	}
	if casted, ok := structType.(*RelativePath); ok {
		return *casted
	}
	return nil
}

func (m *_RelativePath) GetTypeName() string {
	return "RelativePath"
}

func (m *_RelativePath) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfElements)
	lengthInBits += 32

	// Array field
	if len(m.Elements) > 0 {
		for _curItem, element := range m.Elements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Elements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RelativePath) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RelativePath) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__relativePath RelativePath, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RelativePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativePath")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElements, err := ReadImplicitField[int32](ctx, "noOfElements", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElements' field"))
	}
	_ = noOfElements

	elements, err := ReadCountArrayField[RelativePathElement](ctx, "elements", ReadComplex[RelativePathElement](ExtensionObjectDefinitionParseWithBufferProducer[RelativePathElement]((int32)(int32(539))), readBuffer), uint64(noOfElements))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elements' field"))
	}
	m.Elements = elements

	if closeErr := readBuffer.CloseContext("RelativePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativePath")
	}

	return m, nil
}

func (m *_RelativePath) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativePath) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RelativePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RelativePath")
		}
		noOfElements := int32(utils.InlineIf(bool((m.GetElements()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetElements()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfElements", noOfElements, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elements", m.GetElements(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elements' field")
		}

		if popErr := writeBuffer.PopContext("RelativePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RelativePath")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RelativePath) IsRelativePath() {}

func (m *_RelativePath) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RelativePath) deepCopy() *_RelativePath {
	if m == nil {
		return nil
	}
	_RelativePathCopy := &_RelativePath{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[RelativePathElement, RelativePathElement](m.Elements),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RelativePathCopy
}

func (m *_RelativePath) String() string {
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
