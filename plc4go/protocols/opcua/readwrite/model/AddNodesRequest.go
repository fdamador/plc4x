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

// AddNodesRequest is the corresponding interface of AddNodesRequest
type AddNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToAdd returns NoOfNodesToAdd (property field)
	GetNoOfNodesToAdd() int32
	// GetNodesToAdd returns NodesToAdd (property field)
	GetNodesToAdd() []ExtensionObjectDefinition
	// IsAddNodesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAddNodesRequest()
}

// _AddNodesRequest is the data-structure of this message
type _AddNodesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader  ExtensionObjectDefinition
	NoOfNodesToAdd int32
	NodesToAdd     []ExtensionObjectDefinition
}

var _ AddNodesRequest = (*_AddNodesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AddNodesRequest)(nil)

// NewAddNodesRequest factory function for _AddNodesRequest
func NewAddNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToAdd int32, nodesToAdd []ExtensionObjectDefinition) *_AddNodesRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for AddNodesRequest must not be nil")
	}
	_result := &_AddNodesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfNodesToAdd:                    noOfNodesToAdd,
		NodesToAdd:                        nodesToAdd,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddNodesRequest) GetIdentifier() string {
	return "488"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddNodesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_AddNodesRequest) GetNoOfNodesToAdd() int32 {
	return m.NoOfNodesToAdd
}

func (m *_AddNodesRequest) GetNodesToAdd() []ExtensionObjectDefinition {
	return m.NodesToAdd
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAddNodesRequest(structType any) AddNodesRequest {
	if casted, ok := structType.(AddNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AddNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AddNodesRequest) GetTypeName() string {
	return "AddNodesRequest"
}

func (m *_AddNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToAdd)
	lengthInBits += 32

	// Array field
	if len(m.NodesToAdd) > 0 {
		for _curItem, element := range m.NodesToAdd {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToAdd), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AddNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AddNodesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__addNodesRequest AddNodesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AddNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToAdd, err := ReadSimpleField(ctx, "noOfNodesToAdd", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToAdd' field"))
	}
	m.NoOfNodesToAdd = noOfNodesToAdd

	nodesToAdd, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToAdd", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("378")), readBuffer), uint64(noOfNodesToAdd))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToAdd' field"))
	}
	m.NodesToAdd = nodesToAdd

	if closeErr := readBuffer.CloseContext("AddNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddNodesRequest")
	}

	return m, nil
}

func (m *_AddNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddNodesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToAdd", m.GetNoOfNodesToAdd(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToAdd' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToAdd", m.GetNodesToAdd(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToAdd' field")
		}

		if popErr := writeBuffer.PopContext("AddNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddNodesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddNodesRequest) IsAddNodesRequest() {}

func (m *_AddNodesRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AddNodesRequest) deepCopy() *_AddNodesRequest {
	if m == nil {
		return nil
	}
	_AddNodesRequestCopy := &_AddNodesRequest{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfNodesToAdd,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.NodesToAdd),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AddNodesRequestCopy
}

func (m *_AddNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
