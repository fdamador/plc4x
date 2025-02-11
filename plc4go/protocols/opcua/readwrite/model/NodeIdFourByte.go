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

// NodeIdFourByte is the corresponding interface of NodeIdFourByte
type NodeIdFourByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint8
	// GetId returns Id (property field)
	GetId() uint16
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
	// GetNamespace returns Namespace (virtual field)
	GetNamespace() int16
	// IsNodeIdFourByte is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdFourByte()
	// CreateBuilder creates a NodeIdFourByteBuilder
	CreateNodeIdFourByteBuilder() NodeIdFourByteBuilder
}

// _NodeIdFourByte is the data-structure of this message
type _NodeIdFourByte struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint8
	Id             uint16
}

var _ NodeIdFourByte = (*_NodeIdFourByte)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdFourByte)(nil)

// NewNodeIdFourByte factory function for _NodeIdFourByte
func NewNodeIdFourByte(namespaceIndex uint8, id uint16) *_NodeIdFourByte {
	_result := &_NodeIdFourByte{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		NamespaceIndex:               namespaceIndex,
		Id:                           id,
	}
	_result.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdFourByteBuilder is a builder for NodeIdFourByte
type NodeIdFourByteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceIndex uint8, id uint16) NodeIdFourByteBuilder
	// WithNamespaceIndex adds NamespaceIndex (property field)
	WithNamespaceIndex(uint8) NodeIdFourByteBuilder
	// WithId adds Id (property field)
	WithId(uint16) NodeIdFourByteBuilder
	// Build builds the NodeIdFourByte or returns an error if something is wrong
	Build() (NodeIdFourByte, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdFourByte
}

// NewNodeIdFourByteBuilder() creates a NodeIdFourByteBuilder
func NewNodeIdFourByteBuilder() NodeIdFourByteBuilder {
	return &_NodeIdFourByteBuilder{_NodeIdFourByte: new(_NodeIdFourByte)}
}

type _NodeIdFourByteBuilder struct {
	*_NodeIdFourByte

	parentBuilder *_NodeIdTypeDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeIdFourByteBuilder) = (*_NodeIdFourByteBuilder)(nil)

func (b *_NodeIdFourByteBuilder) setParent(contract NodeIdTypeDefinitionContract) {
	b.NodeIdTypeDefinitionContract = contract
}

func (b *_NodeIdFourByteBuilder) WithMandatoryFields(namespaceIndex uint8, id uint16) NodeIdFourByteBuilder {
	return b.WithNamespaceIndex(namespaceIndex).WithId(id)
}

func (b *_NodeIdFourByteBuilder) WithNamespaceIndex(namespaceIndex uint8) NodeIdFourByteBuilder {
	b.NamespaceIndex = namespaceIndex
	return b
}

func (b *_NodeIdFourByteBuilder) WithId(id uint16) NodeIdFourByteBuilder {
	b.Id = id
	return b
}

func (b *_NodeIdFourByteBuilder) Build() (NodeIdFourByte, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdFourByte.deepCopy(), nil
}

func (b *_NodeIdFourByteBuilder) MustBuild() NodeIdFourByte {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NodeIdFourByteBuilder) Done() NodeIdTypeDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NodeIdFourByteBuilder) buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error) {
	return b.Build()
}

func (b *_NodeIdFourByteBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdFourByteBuilder().(*_NodeIdFourByteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdFourByteBuilder creates a NodeIdFourByteBuilder
func (b *_NodeIdFourByte) CreateNodeIdFourByteBuilder() NodeIdFourByteBuilder {
	if b == nil {
		return NewNodeIdFourByteBuilder()
	}
	return &_NodeIdFourByteBuilder{_NodeIdFourByte: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdFourByte) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeFourByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdFourByte) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdFourByte) GetNamespaceIndex() uint8 {
	return m.NamespaceIndex
}

func (m *_NodeIdFourByte) GetId() uint16 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdFourByte) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

func (m *_NodeIdFourByte) GetNamespace() int16 {
	ctx := context.Background()
	_ = ctx
	return int16(m.GetNamespaceIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdFourByte(structType any) NodeIdFourByte {
	if casted, ok := structType.(NodeIdFourByte); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdFourByte); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdFourByte) GetTypeName() string {
	return "NodeIdFourByte"
}

func (m *_NodeIdFourByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 8

	// Simple field (id)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdFourByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdFourByte) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdFourByte NodeIdFourByte, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdFourByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdFourByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	id, err := ReadSimpleField(ctx, "id", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	namespace, err := ReadVirtualField[int16](ctx, "namespace", (*int16)(nil), namespaceIndex)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespace' field"))
	}
	_ = namespace

	if closeErr := readBuffer.CloseContext("NodeIdFourByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdFourByte")
	}

	return m, nil
}

func (m *_NodeIdFourByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdFourByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdFourByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdFourByte")
		}

		if err := WriteSimpleField[uint8](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteSimpleField[uint16](ctx, "id", m.GetId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}
		// Virtual field
		namespace := m.GetNamespace()
		_ = namespace
		if _namespaceErr := writeBuffer.WriteVirtual(ctx, "namespace", m.GetNamespace()); _namespaceErr != nil {
			return errors.Wrap(_namespaceErr, "Error serializing 'namespace' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdFourByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdFourByte")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdFourByte) IsNodeIdFourByte() {}

func (m *_NodeIdFourByte) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdFourByte) deepCopy() *_NodeIdFourByte {
	if m == nil {
		return nil
	}
	_NodeIdFourByteCopy := &_NodeIdFourByte{
		m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).deepCopy(),
		m.NamespaceIndex,
		m.Id,
	}
	m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = m
	return _NodeIdFourByteCopy
}

func (m *_NodeIdFourByte) String() string {
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
