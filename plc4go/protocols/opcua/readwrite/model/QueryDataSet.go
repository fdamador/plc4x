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

// QueryDataSet is the corresponding interface of QueryDataSet
type QueryDataSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() ExpandedNodeId
	// GetTypeDefinitionNode returns TypeDefinitionNode (property field)
	GetTypeDefinitionNode() ExpandedNodeId
	// GetNoOfValues returns NoOfValues (property field)
	GetNoOfValues() int32
	// GetValues returns Values (property field)
	GetValues() []Variant
	// IsQueryDataSet is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsQueryDataSet()
}

// _QueryDataSet is the data-structure of this message
type _QueryDataSet struct {
	ExtensionObjectDefinitionContract
	NodeId             ExpandedNodeId
	TypeDefinitionNode ExpandedNodeId
	NoOfValues         int32
	Values             []Variant
}

var _ QueryDataSet = (*_QueryDataSet)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_QueryDataSet)(nil)

// NewQueryDataSet factory function for _QueryDataSet
func NewQueryDataSet(nodeId ExpandedNodeId, typeDefinitionNode ExpandedNodeId, noOfValues int32, values []Variant) *_QueryDataSet {
	if nodeId == nil {
		panic("nodeId of type ExpandedNodeId for QueryDataSet must not be nil")
	}
	if typeDefinitionNode == nil {
		panic("typeDefinitionNode of type ExpandedNodeId for QueryDataSet must not be nil")
	}
	_result := &_QueryDataSet{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		TypeDefinitionNode:                typeDefinitionNode,
		NoOfValues:                        noOfValues,
		Values:                            values,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryDataSet) GetIdentifier() string {
	return "579"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryDataSet) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryDataSet) GetNodeId() ExpandedNodeId {
	return m.NodeId
}

func (m *_QueryDataSet) GetTypeDefinitionNode() ExpandedNodeId {
	return m.TypeDefinitionNode
}

func (m *_QueryDataSet) GetNoOfValues() int32 {
	return m.NoOfValues
}

func (m *_QueryDataSet) GetValues() []Variant {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastQueryDataSet(structType any) QueryDataSet {
	if casted, ok := structType.(QueryDataSet); ok {
		return casted
	}
	if casted, ok := structType.(*QueryDataSet); ok {
		return *casted
	}
	return nil
}

func (m *_QueryDataSet) GetTypeName() string {
	return "QueryDataSet"
}

func (m *_QueryDataSet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (typeDefinitionNode)
	lengthInBits += m.TypeDefinitionNode.GetLengthInBits(ctx)

	// Simple field (noOfValues)
	lengthInBits += 32

	// Array field
	if len(m.Values) > 0 {
		for _curItem, element := range m.Values {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Values), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_QueryDataSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_QueryDataSet) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__queryDataSet QueryDataSet, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QueryDataSet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryDataSet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "nodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	typeDefinitionNode, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinitionNode' field"))
	}
	m.TypeDefinitionNode = typeDefinitionNode

	noOfValues, err := ReadSimpleField(ctx, "noOfValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfValues' field"))
	}
	m.NoOfValues = noOfValues

	values, err := ReadCountArrayField[Variant](ctx, "values", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'values' field"))
	}
	m.Values = values

	if closeErr := readBuffer.CloseContext("QueryDataSet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryDataSet")
	}

	return m, nil
}

func (m *_QueryDataSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryDataSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryDataSet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryDataSet")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", m.GetTypeDefinitionNode(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinitionNode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfValues", m.GetNoOfValues(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "values", m.GetValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'values' field")
		}

		if popErr := writeBuffer.PopContext("QueryDataSet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryDataSet")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryDataSet) IsQueryDataSet() {}

func (m *_QueryDataSet) DeepCopy() any {
	return m.deepCopy()
}

func (m *_QueryDataSet) deepCopy() *_QueryDataSet {
	if m == nil {
		return nil
	}
	_QueryDataSetCopy := &_QueryDataSet{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.NodeId.DeepCopy().(ExpandedNodeId),
		m.TypeDefinitionNode.DeepCopy().(ExpandedNodeId),
		m.NoOfValues,
		utils.DeepCopySlice[Variant, Variant](m.Values),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _QueryDataSetCopy
}

func (m *_QueryDataSet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
