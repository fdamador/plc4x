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

// NodeTypeDescription is the corresponding interface of NodeTypeDescription
type NodeTypeDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetTypeDefinitionNode returns TypeDefinitionNode (property field)
	GetTypeDefinitionNode() ExpandedNodeId
	// GetIncludeSubTypes returns IncludeSubTypes (property field)
	GetIncludeSubTypes() bool
	// GetNoOfDataToReturn returns NoOfDataToReturn (property field)
	GetNoOfDataToReturn() int32
	// GetDataToReturn returns DataToReturn (property field)
	GetDataToReturn() []ExtensionObjectDefinition
	// IsNodeTypeDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeTypeDescription()
}

// _NodeTypeDescription is the data-structure of this message
type _NodeTypeDescription struct {
	ExtensionObjectDefinitionContract
	TypeDefinitionNode ExpandedNodeId
	IncludeSubTypes    bool
	NoOfDataToReturn   int32
	DataToReturn       []ExtensionObjectDefinition
	// Reserved Fields
	reservedField0 *uint8
}

var _ NodeTypeDescription = (*_NodeTypeDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NodeTypeDescription)(nil)

// NewNodeTypeDescription factory function for _NodeTypeDescription
func NewNodeTypeDescription(typeDefinitionNode ExpandedNodeId, includeSubTypes bool, noOfDataToReturn int32, dataToReturn []ExtensionObjectDefinition) *_NodeTypeDescription {
	if typeDefinitionNode == nil {
		panic("typeDefinitionNode of type ExpandedNodeId for NodeTypeDescription must not be nil")
	}
	_result := &_NodeTypeDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TypeDefinitionNode:                typeDefinitionNode,
		IncludeSubTypes:                   includeSubTypes,
		NoOfDataToReturn:                  noOfDataToReturn,
		DataToReturn:                      dataToReturn,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeTypeDescription) GetIdentifier() string {
	return "575"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeTypeDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeTypeDescription) GetTypeDefinitionNode() ExpandedNodeId {
	return m.TypeDefinitionNode
}

func (m *_NodeTypeDescription) GetIncludeSubTypes() bool {
	return m.IncludeSubTypes
}

func (m *_NodeTypeDescription) GetNoOfDataToReturn() int32 {
	return m.NoOfDataToReturn
}

func (m *_NodeTypeDescription) GetDataToReturn() []ExtensionObjectDefinition {
	return m.DataToReturn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeTypeDescription(structType any) NodeTypeDescription {
	if casted, ok := structType.(NodeTypeDescription); ok {
		return casted
	}
	if casted, ok := structType.(*NodeTypeDescription); ok {
		return *casted
	}
	return nil
}

func (m *_NodeTypeDescription) GetTypeName() string {
	return "NodeTypeDescription"
}

func (m *_NodeTypeDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (typeDefinitionNode)
	lengthInBits += m.TypeDefinitionNode.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (includeSubTypes)
	lengthInBits += 1

	// Simple field (noOfDataToReturn)
	lengthInBits += 32

	// Array field
	if len(m.DataToReturn) > 0 {
		for _curItem, element := range m.DataToReturn {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataToReturn), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NodeTypeDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeTypeDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__nodeTypeDescription NodeTypeDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeTypeDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeTypeDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	typeDefinitionNode, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinitionNode' field"))
	}
	m.TypeDefinitionNode = typeDefinitionNode

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	includeSubTypes, err := ReadSimpleField(ctx, "includeSubTypes", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'includeSubTypes' field"))
	}
	m.IncludeSubTypes = includeSubTypes

	noOfDataToReturn, err := ReadSimpleField(ctx, "noOfDataToReturn", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataToReturn' field"))
	}
	m.NoOfDataToReturn = noOfDataToReturn

	dataToReturn, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "dataToReturn", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("572")), readBuffer), uint64(noOfDataToReturn))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataToReturn' field"))
	}
	m.DataToReturn = dataToReturn

	if closeErr := readBuffer.CloseContext("NodeTypeDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeTypeDescription")
	}

	return m, nil
}

func (m *_NodeTypeDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeTypeDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeTypeDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeTypeDescription")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", m.GetTypeDefinitionNode(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinitionNode' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "includeSubTypes", m.GetIncludeSubTypes(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'includeSubTypes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDataToReturn", m.GetNoOfDataToReturn(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataToReturn' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataToReturn", m.GetDataToReturn(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataToReturn' field")
		}

		if popErr := writeBuffer.PopContext("NodeTypeDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeTypeDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeTypeDescription) IsNodeTypeDescription() {}

func (m *_NodeTypeDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeTypeDescription) deepCopy() *_NodeTypeDescription {
	if m == nil {
		return nil
	}
	_NodeTypeDescriptionCopy := &_NodeTypeDescription{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.TypeDefinitionNode.DeepCopy().(ExpandedNodeId),
		m.IncludeSubTypes,
		m.NoOfDataToReturn,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.DataToReturn),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NodeTypeDescriptionCopy
}

func (m *_NodeTypeDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
