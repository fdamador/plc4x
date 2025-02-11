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

// CloseSessionResponse is the corresponding interface of CloseSessionResponse
type CloseSessionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// IsCloseSessionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCloseSessionResponse()
	// CreateBuilder creates a CloseSessionResponseBuilder
	CreateCloseSessionResponseBuilder() CloseSessionResponseBuilder
}

// _CloseSessionResponse is the data-structure of this message
type _CloseSessionResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ResponseHeader
}

var _ CloseSessionResponse = (*_CloseSessionResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSessionResponse)(nil)

// NewCloseSessionResponse factory function for _CloseSessionResponse
func NewCloseSessionResponse(responseHeader ResponseHeader) *_CloseSessionResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for CloseSessionResponse must not be nil")
	}
	_result := &_CloseSessionResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CloseSessionResponseBuilder is a builder for CloseSessionResponse
type CloseSessionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader) CloseSessionResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) CloseSessionResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) CloseSessionResponseBuilder
	// Build builds the CloseSessionResponse or returns an error if something is wrong
	Build() (CloseSessionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CloseSessionResponse
}

// NewCloseSessionResponseBuilder() creates a CloseSessionResponseBuilder
func NewCloseSessionResponseBuilder() CloseSessionResponseBuilder {
	return &_CloseSessionResponseBuilder{_CloseSessionResponse: new(_CloseSessionResponse)}
}

type _CloseSessionResponseBuilder struct {
	*_CloseSessionResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CloseSessionResponseBuilder) = (*_CloseSessionResponseBuilder)(nil)

func (b *_CloseSessionResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CloseSessionResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader) CloseSessionResponseBuilder {
	return b.WithResponseHeader(responseHeader)
}

func (b *_CloseSessionResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) CloseSessionResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_CloseSessionResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) CloseSessionResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_CloseSessionResponseBuilder) Build() (CloseSessionResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CloseSessionResponse.deepCopy(), nil
}

func (b *_CloseSessionResponseBuilder) MustBuild() CloseSessionResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CloseSessionResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CloseSessionResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CloseSessionResponseBuilder) DeepCopy() any {
	_copy := b.CreateCloseSessionResponseBuilder().(*_CloseSessionResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCloseSessionResponseBuilder creates a CloseSessionResponseBuilder
func (b *_CloseSessionResponse) CreateCloseSessionResponseBuilder() CloseSessionResponseBuilder {
	if b == nil {
		return NewCloseSessionResponseBuilder()
	}
	return &_CloseSessionResponseBuilder{_CloseSessionResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSessionResponse) GetExtensionId() int32 {
	return int32(476)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSessionResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSessionResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCloseSessionResponse(structType any) CloseSessionResponse {
	if casted, ok := structType.(CloseSessionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSessionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSessionResponse) GetTypeName() string {
	return "CloseSessionResponse"
}

func (m *_CloseSessionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSessionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSessionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__closeSessionResponse CloseSessionResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSessionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSessionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	if closeErr := readBuffer.CloseContext("CloseSessionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSessionResponse")
	}

	return m, nil
}

func (m *_CloseSessionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSessionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSessionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSessionResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSessionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSessionResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSessionResponse) IsCloseSessionResponse() {}

func (m *_CloseSessionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CloseSessionResponse) deepCopy() *_CloseSessionResponse {
	if m == nil {
		return nil
	}
	_CloseSessionResponseCopy := &_CloseSessionResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ResponseHeader),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CloseSessionResponseCopy
}

func (m *_CloseSessionResponse) String() string {
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
