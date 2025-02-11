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

// RepublishResponse is the corresponding interface of RepublishResponse
type RepublishResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetNotificationMessage returns NotificationMessage (property field)
	GetNotificationMessage() NotificationMessage
	// IsRepublishResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRepublishResponse()
	// CreateBuilder creates a RepublishResponseBuilder
	CreateRepublishResponseBuilder() RepublishResponseBuilder
}

// _RepublishResponse is the data-structure of this message
type _RepublishResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ResponseHeader
	NotificationMessage NotificationMessage
}

var _ RepublishResponse = (*_RepublishResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RepublishResponse)(nil)

// NewRepublishResponse factory function for _RepublishResponse
func NewRepublishResponse(responseHeader ResponseHeader, notificationMessage NotificationMessage) *_RepublishResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for RepublishResponse must not be nil")
	}
	if notificationMessage == nil {
		panic("notificationMessage of type NotificationMessage for RepublishResponse must not be nil")
	}
	_result := &_RepublishResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NotificationMessage:               notificationMessage,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RepublishResponseBuilder is a builder for RepublishResponse
type RepublishResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, notificationMessage NotificationMessage) RepublishResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) RepublishResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) RepublishResponseBuilder
	// WithNotificationMessage adds NotificationMessage (property field)
	WithNotificationMessage(NotificationMessage) RepublishResponseBuilder
	// WithNotificationMessageBuilder adds NotificationMessage (property field) which is build by the builder
	WithNotificationMessageBuilder(func(NotificationMessageBuilder) NotificationMessageBuilder) RepublishResponseBuilder
	// Build builds the RepublishResponse or returns an error if something is wrong
	Build() (RepublishResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RepublishResponse
}

// NewRepublishResponseBuilder() creates a RepublishResponseBuilder
func NewRepublishResponseBuilder() RepublishResponseBuilder {
	return &_RepublishResponseBuilder{_RepublishResponse: new(_RepublishResponse)}
}

type _RepublishResponseBuilder struct {
	*_RepublishResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RepublishResponseBuilder) = (*_RepublishResponseBuilder)(nil)

func (b *_RepublishResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RepublishResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, notificationMessage NotificationMessage) RepublishResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithNotificationMessage(notificationMessage)
}

func (b *_RepublishResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) RepublishResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_RepublishResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) RepublishResponseBuilder {
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

func (b *_RepublishResponseBuilder) WithNotificationMessage(notificationMessage NotificationMessage) RepublishResponseBuilder {
	b.NotificationMessage = notificationMessage
	return b
}

func (b *_RepublishResponseBuilder) WithNotificationMessageBuilder(builderSupplier func(NotificationMessageBuilder) NotificationMessageBuilder) RepublishResponseBuilder {
	builder := builderSupplier(b.NotificationMessage.CreateNotificationMessageBuilder())
	var err error
	b.NotificationMessage, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NotificationMessageBuilder failed"))
	}
	return b
}

func (b *_RepublishResponseBuilder) Build() (RepublishResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.NotificationMessage == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'notificationMessage' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RepublishResponse.deepCopy(), nil
}

func (b *_RepublishResponseBuilder) MustBuild() RepublishResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RepublishResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RepublishResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RepublishResponseBuilder) DeepCopy() any {
	_copy := b.CreateRepublishResponseBuilder().(*_RepublishResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRepublishResponseBuilder creates a RepublishResponseBuilder
func (b *_RepublishResponse) CreateRepublishResponseBuilder() RepublishResponseBuilder {
	if b == nil {
		return NewRepublishResponseBuilder()
	}
	return &_RepublishResponseBuilder{_RepublishResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RepublishResponse) GetExtensionId() int32 {
	return int32(835)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RepublishResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RepublishResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_RepublishResponse) GetNotificationMessage() NotificationMessage {
	return m.NotificationMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRepublishResponse(structType any) RepublishResponse {
	if casted, ok := structType.(RepublishResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RepublishResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RepublishResponse) GetTypeName() string {
	return "RepublishResponse"
}

func (m *_RepublishResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (notificationMessage)
	lengthInBits += m.NotificationMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RepublishResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RepublishResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__republishResponse RepublishResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RepublishResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RepublishResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	notificationMessage, err := ReadSimpleField[NotificationMessage](ctx, "notificationMessage", ReadComplex[NotificationMessage](ExtensionObjectDefinitionParseWithBufferProducer[NotificationMessage]((int32)(int32(805))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationMessage' field"))
	}
	m.NotificationMessage = notificationMessage

	if closeErr := readBuffer.CloseContext("RepublishResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RepublishResponse")
	}

	return m, nil
}

func (m *_RepublishResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RepublishResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RepublishResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RepublishResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[NotificationMessage](ctx, "notificationMessage", m.GetNotificationMessage(), WriteComplex[NotificationMessage](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationMessage' field")
		}

		if popErr := writeBuffer.PopContext("RepublishResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RepublishResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RepublishResponse) IsRepublishResponse() {}

func (m *_RepublishResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RepublishResponse) deepCopy() *_RepublishResponse {
	if m == nil {
		return nil
	}
	_RepublishResponseCopy := &_RepublishResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ResponseHeader),
		m.NotificationMessage.DeepCopy().(NotificationMessage),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RepublishResponseCopy
}

func (m *_RepublishResponse) String() string {
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
