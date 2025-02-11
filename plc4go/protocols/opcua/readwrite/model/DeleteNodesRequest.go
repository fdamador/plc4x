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
	GetRequestHeader() RequestHeader
	// GetNodesToDelete returns NodesToDelete (property field)
	GetNodesToDelete() []DeleteNodesItem
	// IsDeleteNodesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteNodesRequest()
	// CreateBuilder creates a DeleteNodesRequestBuilder
	CreateDeleteNodesRequestBuilder() DeleteNodesRequestBuilder
}

// _DeleteNodesRequest is the data-structure of this message
type _DeleteNodesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader RequestHeader
	NodesToDelete []DeleteNodesItem
}

var _ DeleteNodesRequest = (*_DeleteNodesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteNodesRequest)(nil)

// NewDeleteNodesRequest factory function for _DeleteNodesRequest
func NewDeleteNodesRequest(requestHeader RequestHeader, nodesToDelete []DeleteNodesItem) *_DeleteNodesRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for DeleteNodesRequest must not be nil")
	}
	_result := &_DeleteNodesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NodesToDelete:                     nodesToDelete,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteNodesRequestBuilder is a builder for DeleteNodesRequest
type DeleteNodesRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, nodesToDelete []DeleteNodesItem) DeleteNodesRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) DeleteNodesRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteNodesRequestBuilder
	// WithNodesToDelete adds NodesToDelete (property field)
	WithNodesToDelete(...DeleteNodesItem) DeleteNodesRequestBuilder
	// Build builds the DeleteNodesRequest or returns an error if something is wrong
	Build() (DeleteNodesRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteNodesRequest
}

// NewDeleteNodesRequestBuilder() creates a DeleteNodesRequestBuilder
func NewDeleteNodesRequestBuilder() DeleteNodesRequestBuilder {
	return &_DeleteNodesRequestBuilder{_DeleteNodesRequest: new(_DeleteNodesRequest)}
}

type _DeleteNodesRequestBuilder struct {
	*_DeleteNodesRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteNodesRequestBuilder) = (*_DeleteNodesRequestBuilder)(nil)

func (b *_DeleteNodesRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteNodesRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, nodesToDelete []DeleteNodesItem) DeleteNodesRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithNodesToDelete(nodesToDelete...)
}

func (b *_DeleteNodesRequestBuilder) WithRequestHeader(requestHeader RequestHeader) DeleteNodesRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_DeleteNodesRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteNodesRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_DeleteNodesRequestBuilder) WithNodesToDelete(nodesToDelete ...DeleteNodesItem) DeleteNodesRequestBuilder {
	b.NodesToDelete = nodesToDelete
	return b
}

func (b *_DeleteNodesRequestBuilder) Build() (DeleteNodesRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteNodesRequest.deepCopy(), nil
}

func (b *_DeleteNodesRequestBuilder) MustBuild() DeleteNodesRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DeleteNodesRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_DeleteNodesRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteNodesRequestBuilder) DeepCopy() any {
	_copy := b.CreateDeleteNodesRequestBuilder().(*_DeleteNodesRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteNodesRequestBuilder creates a DeleteNodesRequestBuilder
func (b *_DeleteNodesRequest) CreateDeleteNodesRequestBuilder() DeleteNodesRequestBuilder {
	if b == nil {
		return NewDeleteNodesRequestBuilder()
	}
	return &_DeleteNodesRequestBuilder{_DeleteNodesRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesRequest) GetExtensionId() int32 {
	return int32(500)
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

func (m *_DeleteNodesRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_DeleteNodesRequest) GetNodesToDelete() []DeleteNodesItem {
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

	// Implicit Field (noOfNodesToDelete)
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

func (m *_DeleteNodesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteNodesRequest DeleteNodesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToDelete, err := ReadImplicitField[int32](ctx, "noOfNodesToDelete", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToDelete' field"))
	}
	_ = noOfNodesToDelete

	nodesToDelete, err := ReadCountArrayField[DeleteNodesItem](ctx, "nodesToDelete", ReadComplex[DeleteNodesItem](ExtensionObjectDefinitionParseWithBufferProducer[DeleteNodesItem]((int32)(int32(384))), readBuffer), uint64(noOfNodesToDelete))
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

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}
		noOfNodesToDelete := int32(utils.InlineIf(bool((m.GetNodesToDelete()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNodesToDelete()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNodesToDelete", noOfNodesToDelete, WriteSignedInt(writeBuffer, 32)); err != nil {
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
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(RequestHeader),
		utils.DeepCopySlice[DeleteNodesItem, DeleteNodesItem](m.NodesToDelete),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteNodesRequestCopy
}

func (m *_DeleteNodesRequest) String() string {
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
