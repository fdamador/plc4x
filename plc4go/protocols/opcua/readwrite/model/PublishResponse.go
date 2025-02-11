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

// PublishResponse is the corresponding interface of PublishResponse
type PublishResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetAvailableSequenceNumbers returns AvailableSequenceNumbers (property field)
	GetAvailableSequenceNumbers() []uint32
	// GetMoreNotifications returns MoreNotifications (property field)
	GetMoreNotifications() bool
	// GetNotificationMessage returns NotificationMessage (property field)
	GetNotificationMessage() NotificationMessage
	// GetResults returns Results (property field)
	GetResults() []StatusCode
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsPublishResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPublishResponse()
	// CreateBuilder creates a PublishResponseBuilder
	CreatePublishResponseBuilder() PublishResponseBuilder
}

// _PublishResponse is the data-structure of this message
type _PublishResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader           ResponseHeader
	SubscriptionId           uint32
	AvailableSequenceNumbers []uint32
	MoreNotifications        bool
	NotificationMessage      NotificationMessage
	Results                  []StatusCode
	DiagnosticInfos          []DiagnosticInfo
	// Reserved Fields
	reservedField0 *uint8
}

var _ PublishResponse = (*_PublishResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PublishResponse)(nil)

// NewPublishResponse factory function for _PublishResponse
func NewPublishResponse(responseHeader ResponseHeader, subscriptionId uint32, availableSequenceNumbers []uint32, moreNotifications bool, notificationMessage NotificationMessage, results []StatusCode, diagnosticInfos []DiagnosticInfo) *_PublishResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for PublishResponse must not be nil")
	}
	if notificationMessage == nil {
		panic("notificationMessage of type NotificationMessage for PublishResponse must not be nil")
	}
	_result := &_PublishResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		SubscriptionId:                    subscriptionId,
		AvailableSequenceNumbers:          availableSequenceNumbers,
		MoreNotifications:                 moreNotifications,
		NotificationMessage:               notificationMessage,
		Results:                           results,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PublishResponseBuilder is a builder for PublishResponse
type PublishResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, subscriptionId uint32, availableSequenceNumbers []uint32, moreNotifications bool, notificationMessage NotificationMessage, results []StatusCode, diagnosticInfos []DiagnosticInfo) PublishResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) PublishResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) PublishResponseBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) PublishResponseBuilder
	// WithAvailableSequenceNumbers adds AvailableSequenceNumbers (property field)
	WithAvailableSequenceNumbers(...uint32) PublishResponseBuilder
	// WithMoreNotifications adds MoreNotifications (property field)
	WithMoreNotifications(bool) PublishResponseBuilder
	// WithNotificationMessage adds NotificationMessage (property field)
	WithNotificationMessage(NotificationMessage) PublishResponseBuilder
	// WithNotificationMessageBuilder adds NotificationMessage (property field) which is build by the builder
	WithNotificationMessageBuilder(func(NotificationMessageBuilder) NotificationMessageBuilder) PublishResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...StatusCode) PublishResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) PublishResponseBuilder
	// Build builds the PublishResponse or returns an error if something is wrong
	Build() (PublishResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PublishResponse
}

// NewPublishResponseBuilder() creates a PublishResponseBuilder
func NewPublishResponseBuilder() PublishResponseBuilder {
	return &_PublishResponseBuilder{_PublishResponse: new(_PublishResponse)}
}

type _PublishResponseBuilder struct {
	*_PublishResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PublishResponseBuilder) = (*_PublishResponseBuilder)(nil)

func (b *_PublishResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PublishResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, subscriptionId uint32, availableSequenceNumbers []uint32, moreNotifications bool, notificationMessage NotificationMessage, results []StatusCode, diagnosticInfos []DiagnosticInfo) PublishResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithSubscriptionId(subscriptionId).WithAvailableSequenceNumbers(availableSequenceNumbers...).WithMoreNotifications(moreNotifications).WithNotificationMessage(notificationMessage).WithResults(results...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_PublishResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) PublishResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_PublishResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) PublishResponseBuilder {
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

func (b *_PublishResponseBuilder) WithSubscriptionId(subscriptionId uint32) PublishResponseBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_PublishResponseBuilder) WithAvailableSequenceNumbers(availableSequenceNumbers ...uint32) PublishResponseBuilder {
	b.AvailableSequenceNumbers = availableSequenceNumbers
	return b
}

func (b *_PublishResponseBuilder) WithMoreNotifications(moreNotifications bool) PublishResponseBuilder {
	b.MoreNotifications = moreNotifications
	return b
}

func (b *_PublishResponseBuilder) WithNotificationMessage(notificationMessage NotificationMessage) PublishResponseBuilder {
	b.NotificationMessage = notificationMessage
	return b
}

func (b *_PublishResponseBuilder) WithNotificationMessageBuilder(builderSupplier func(NotificationMessageBuilder) NotificationMessageBuilder) PublishResponseBuilder {
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

func (b *_PublishResponseBuilder) WithResults(results ...StatusCode) PublishResponseBuilder {
	b.Results = results
	return b
}

func (b *_PublishResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) PublishResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_PublishResponseBuilder) Build() (PublishResponse, error) {
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
	return b._PublishResponse.deepCopy(), nil
}

func (b *_PublishResponseBuilder) MustBuild() PublishResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_PublishResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_PublishResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PublishResponseBuilder) DeepCopy() any {
	_copy := b.CreatePublishResponseBuilder().(*_PublishResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePublishResponseBuilder creates a PublishResponseBuilder
func (b *_PublishResponse) CreatePublishResponseBuilder() PublishResponseBuilder {
	if b == nil {
		return NewPublishResponseBuilder()
	}
	return &_PublishResponseBuilder{_PublishResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishResponse) GetExtensionId() int32 {
	return int32(829)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_PublishResponse) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_PublishResponse) GetAvailableSequenceNumbers() []uint32 {
	return m.AvailableSequenceNumbers
}

func (m *_PublishResponse) GetMoreNotifications() bool {
	return m.MoreNotifications
}

func (m *_PublishResponse) GetNotificationMessage() NotificationMessage {
	return m.NotificationMessage
}

func (m *_PublishResponse) GetResults() []StatusCode {
	return m.Results
}

func (m *_PublishResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPublishResponse(structType any) PublishResponse {
	if casted, ok := structType.(PublishResponse); ok {
		return casted
	}
	if casted, ok := structType.(*PublishResponse); ok {
		return *casted
	}
	return nil
}

func (m *_PublishResponse) GetTypeName() string {
	return "PublishResponse"
}

func (m *_PublishResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Implicit Field (noOfAvailableSequenceNumbers)
	lengthInBits += 32

	// Array field
	if len(m.AvailableSequenceNumbers) > 0 {
		lengthInBits += 32 * uint16(len(m.AvailableSequenceNumbers))
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (moreNotifications)
	lengthInBits += 1

	// Simple field (notificationMessage)
	lengthInBits += m.NotificationMessage.GetLengthInBits(ctx)

	// Implicit Field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PublishResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PublishResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__publishResponse PublishResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PublishResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	noOfAvailableSequenceNumbers, err := ReadImplicitField[int32](ctx, "noOfAvailableSequenceNumbers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfAvailableSequenceNumbers' field"))
	}
	_ = noOfAvailableSequenceNumbers

	availableSequenceNumbers, err := ReadCountArrayField[uint32](ctx, "availableSequenceNumbers", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfAvailableSequenceNumbers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'availableSequenceNumbers' field"))
	}
	m.AvailableSequenceNumbers = availableSequenceNumbers

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	moreNotifications, err := ReadSimpleField(ctx, "moreNotifications", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreNotifications' field"))
	}
	m.MoreNotifications = moreNotifications

	notificationMessage, err := ReadSimpleField[NotificationMessage](ctx, "notificationMessage", ReadComplex[NotificationMessage](ExtensionObjectDefinitionParseWithBufferProducer[NotificationMessage]((int32)(int32(805))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationMessage' field"))
	}
	m.NotificationMessage = notificationMessage

	noOfResults, err := ReadImplicitField[int32](ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	_ = noOfResults

	results, err := ReadCountArrayField[StatusCode](ctx, "results", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	_ = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("PublishResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishResponse")
	}

	return m, nil
}

func (m *_PublishResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}
		noOfAvailableSequenceNumbers := int32(utils.InlineIf(bool((m.GetAvailableSequenceNumbers()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetAvailableSequenceNumbers()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfAvailableSequenceNumbers", noOfAvailableSequenceNumbers, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfAvailableSequenceNumbers' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "availableSequenceNumbers", m.GetAvailableSequenceNumbers(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'availableSequenceNumbers' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "moreNotifications", m.GetMoreNotifications(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'moreNotifications' field")
		}

		if err := WriteSimpleField[NotificationMessage](ctx, "notificationMessage", m.GetNotificationMessage(), WriteComplex[NotificationMessage](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationMessage' field")
		}
		noOfResults := int32(utils.InlineIf(bool((m.GetResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfResults", noOfResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("PublishResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishResponse) IsPublishResponse() {}

func (m *_PublishResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PublishResponse) deepCopy() *_PublishResponse {
	if m == nil {
		return nil
	}
	_PublishResponseCopy := &_PublishResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ResponseHeader),
		m.SubscriptionId,
		utils.DeepCopySlice[uint32, uint32](m.AvailableSequenceNumbers),
		m.MoreNotifications,
		m.NotificationMessage.DeepCopy().(NotificationMessage),
		utils.DeepCopySlice[StatusCode, StatusCode](m.Results),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PublishResponseCopy
}

func (m *_PublishResponse) String() string {
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
