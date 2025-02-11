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

// BrowsePath is the corresponding interface of BrowsePath
type BrowsePath interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStartingNode returns StartingNode (property field)
	GetStartingNode() NodeId
	// GetRelativePath returns RelativePath (property field)
	GetRelativePath() RelativePath
	// IsBrowsePath is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowsePath()
	// CreateBuilder creates a BrowsePathBuilder
	CreateBrowsePathBuilder() BrowsePathBuilder
}

// _BrowsePath is the data-structure of this message
type _BrowsePath struct {
	ExtensionObjectDefinitionContract
	StartingNode NodeId
	RelativePath RelativePath
}

var _ BrowsePath = (*_BrowsePath)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowsePath)(nil)

// NewBrowsePath factory function for _BrowsePath
func NewBrowsePath(startingNode NodeId, relativePath RelativePath) *_BrowsePath {
	if startingNode == nil {
		panic("startingNode of type NodeId for BrowsePath must not be nil")
	}
	if relativePath == nil {
		panic("relativePath of type RelativePath for BrowsePath must not be nil")
	}
	_result := &_BrowsePath{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StartingNode:                      startingNode,
		RelativePath:                      relativePath,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowsePathBuilder is a builder for BrowsePath
type BrowsePathBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startingNode NodeId, relativePath RelativePath) BrowsePathBuilder
	// WithStartingNode adds StartingNode (property field)
	WithStartingNode(NodeId) BrowsePathBuilder
	// WithStartingNodeBuilder adds StartingNode (property field) which is build by the builder
	WithStartingNodeBuilder(func(NodeIdBuilder) NodeIdBuilder) BrowsePathBuilder
	// WithRelativePath adds RelativePath (property field)
	WithRelativePath(RelativePath) BrowsePathBuilder
	// WithRelativePathBuilder adds RelativePath (property field) which is build by the builder
	WithRelativePathBuilder(func(RelativePathBuilder) RelativePathBuilder) BrowsePathBuilder
	// Build builds the BrowsePath or returns an error if something is wrong
	Build() (BrowsePath, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowsePath
}

// NewBrowsePathBuilder() creates a BrowsePathBuilder
func NewBrowsePathBuilder() BrowsePathBuilder {
	return &_BrowsePathBuilder{_BrowsePath: new(_BrowsePath)}
}

type _BrowsePathBuilder struct {
	*_BrowsePath

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrowsePathBuilder) = (*_BrowsePathBuilder)(nil)

func (b *_BrowsePathBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_BrowsePathBuilder) WithMandatoryFields(startingNode NodeId, relativePath RelativePath) BrowsePathBuilder {
	return b.WithStartingNode(startingNode).WithRelativePath(relativePath)
}

func (b *_BrowsePathBuilder) WithStartingNode(startingNode NodeId) BrowsePathBuilder {
	b.StartingNode = startingNode
	return b
}

func (b *_BrowsePathBuilder) WithStartingNodeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) BrowsePathBuilder {
	builder := builderSupplier(b.StartingNode.CreateNodeIdBuilder())
	var err error
	b.StartingNode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_BrowsePathBuilder) WithRelativePath(relativePath RelativePath) BrowsePathBuilder {
	b.RelativePath = relativePath
	return b
}

func (b *_BrowsePathBuilder) WithRelativePathBuilder(builderSupplier func(RelativePathBuilder) RelativePathBuilder) BrowsePathBuilder {
	builder := builderSupplier(b.RelativePath.CreateRelativePathBuilder())
	var err error
	b.RelativePath, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RelativePathBuilder failed"))
	}
	return b
}

func (b *_BrowsePathBuilder) Build() (BrowsePath, error) {
	if b.StartingNode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'startingNode' not set"))
	}
	if b.RelativePath == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relativePath' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrowsePath.deepCopy(), nil
}

func (b *_BrowsePathBuilder) MustBuild() BrowsePath {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BrowsePathBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_BrowsePathBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrowsePathBuilder) DeepCopy() any {
	_copy := b.CreateBrowsePathBuilder().(*_BrowsePathBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrowsePathBuilder creates a BrowsePathBuilder
func (b *_BrowsePath) CreateBrowsePathBuilder() BrowsePathBuilder {
	if b == nil {
		return NewBrowsePathBuilder()
	}
	return &_BrowsePathBuilder{_BrowsePath: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePath) GetExtensionId() int32 {
	return int32(545)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePath) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePath) GetStartingNode() NodeId {
	return m.StartingNode
}

func (m *_BrowsePath) GetRelativePath() RelativePath {
	return m.RelativePath
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowsePath(structType any) BrowsePath {
	if casted, ok := structType.(BrowsePath); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePath); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePath) GetTypeName() string {
	return "BrowsePath"
}

func (m *_BrowsePath) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (startingNode)
	lengthInBits += m.StartingNode.GetLengthInBits(ctx)

	// Simple field (relativePath)
	lengthInBits += m.RelativePath.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BrowsePath) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowsePath) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__browsePath BrowsePath, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowsePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePath")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingNode, err := ReadSimpleField[NodeId](ctx, "startingNode", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingNode' field"))
	}
	m.StartingNode = startingNode

	relativePath, err := ReadSimpleField[RelativePath](ctx, "relativePath", ReadComplex[RelativePath](ExtensionObjectDefinitionParseWithBufferProducer[RelativePath]((int32)(int32(542))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relativePath' field"))
	}
	m.RelativePath = relativePath

	if closeErr := readBuffer.CloseContext("BrowsePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePath")
	}

	return m, nil
}

func (m *_BrowsePath) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePath) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePath")
		}

		if err := WriteSimpleField[NodeId](ctx, "startingNode", m.GetStartingNode(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingNode' field")
		}

		if err := WriteSimpleField[RelativePath](ctx, "relativePath", m.GetRelativePath(), WriteComplex[RelativePath](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relativePath' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePath")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePath) IsBrowsePath() {}

func (m *_BrowsePath) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowsePath) deepCopy() *_BrowsePath {
	if m == nil {
		return nil
	}
	_BrowsePathCopy := &_BrowsePath{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StartingNode.DeepCopy().(NodeId),
		m.RelativePath.DeepCopy().(RelativePath),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowsePathCopy
}

func (m *_BrowsePath) String() string {
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
