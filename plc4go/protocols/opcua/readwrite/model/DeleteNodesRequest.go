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

// DeleteNodesRequest is the corresponding interface of DeleteNodesRequest
type DeleteNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToDelete returns NoOfNodesToDelete (property field)
	GetNoOfNodesToDelete() int32
	// GetNodesToDelete returns NodesToDelete (property field)
	GetNodesToDelete() []ExtensionObjectDefinition
	// IsDeleteNodesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteNodesRequest()
}

// _DeleteNodesRequest is the data-structure of this message
type _DeleteNodesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader     ExtensionObjectDefinition
	NoOfNodesToDelete int32
	NodesToDelete     []ExtensionObjectDefinition
}

var _ DeleteNodesRequest = (*_DeleteNodesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteNodesRequest)(nil)

// NewDeleteNodesRequest factory function for _DeleteNodesRequest
func NewDeleteNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToDelete int32, nodesToDelete []ExtensionObjectDefinition) *_DeleteNodesRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for DeleteNodesRequest must not be nil")
	}
	_result := &_DeleteNodesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfNodesToDelete:                 noOfNodesToDelete,
		NodesToDelete:                     nodesToDelete,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesRequest) GetIdentifier() string {
	return "500"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteNodesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteNodesRequest) GetNoOfNodesToDelete() int32 {
	return m.NoOfNodesToDelete
}

func (m *_DeleteNodesRequest) GetNodesToDelete() []ExtensionObjectDefinition {
	return m.NodesToDelete
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteNodesRequest(structType any) DeleteNodesRequest {
	if casted, ok := structType.(DeleteNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteNodesRequest) GetTypeName() string {
	return "DeleteNodesRequest"
}

func (m *_DeleteNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToDelete)
	lengthInBits += 32

	// Array field
	if len(m.NodesToDelete) > 0 {
		for _curItem, element := range m.NodesToDelete {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToDelete), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteNodesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteNodesRequest DeleteNodesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToDelete, err := ReadSimpleField(ctx, "noOfNodesToDelete", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToDelete' field"))
	}
	m.NoOfNodesToDelete = noOfNodesToDelete

	nodesToDelete, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToDelete", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("384")), readBuffer), uint64(noOfNodesToDelete))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToDelete' field"))
	}
	m.NodesToDelete = nodesToDelete

	if closeErr := readBuffer.CloseContext("DeleteNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteNodesRequest")
	}

	return m, nil
}

func (m *_DeleteNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteNodesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToDelete", m.GetNoOfNodesToDelete(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToDelete' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToDelete", m.GetNodesToDelete(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToDelete' field")
		}

		if popErr := writeBuffer.PopContext("DeleteNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteNodesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteNodesRequest) IsDeleteNodesRequest() {}

func (m *_DeleteNodesRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteNodesRequest) deepCopy() *_DeleteNodesRequest {
	if m == nil {
		return nil
	}
	_DeleteNodesRequestCopy := &_DeleteNodesRequest{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfNodesToDelete,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.NodesToDelete),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteNodesRequestCopy
}

func (m *_DeleteNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
