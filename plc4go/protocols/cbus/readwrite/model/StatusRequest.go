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

// StatusRequest is the corresponding interface of StatusRequest
type StatusRequest interface {
	StatusRequestContract
	StatusRequestRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsStatusRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusRequest()
	// CreateBuilder creates a StatusRequestBuilder
	CreateStatusRequestBuilder() StatusRequestBuilder
}

// StatusRequestContract provides a set of functions which can be overwritten by a sub struct
type StatusRequestContract interface {
	// GetStatusType returns StatusType (property field)
	GetStatusType() byte
	// IsStatusRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusRequest()
	// CreateBuilder creates a StatusRequestBuilder
	CreateStatusRequestBuilder() StatusRequestBuilder
}

// StatusRequestRequirements provides a set of functions which need to be implemented by a sub struct
type StatusRequestRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetStatusType returns StatusType (discriminator field)
	GetStatusType() byte
}

// _StatusRequest is the data-structure of this message
type _StatusRequest struct {
	_SubType interface {
		StatusRequestContract
		StatusRequestRequirements
	}
	StatusType byte
}

var _ StatusRequestContract = (*_StatusRequest)(nil)

// NewStatusRequest factory function for _StatusRequest
func NewStatusRequest(statusType byte) *_StatusRequest {
	return &_StatusRequest{StatusType: statusType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StatusRequestBuilder is a builder for StatusRequest
type StatusRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusType byte) StatusRequestBuilder
	// WithStatusType adds StatusType (property field)
	WithStatusType(byte) StatusRequestBuilder
	// AsStatusRequestBinaryState converts this build to a subType of StatusRequest. It is always possible to return to current builder using Done()
	AsStatusRequestBinaryState() interface {
		StatusRequestBinaryStateBuilder
		Done() StatusRequestBuilder
	}
	// AsStatusRequestBinaryStateDeprecated converts this build to a subType of StatusRequest. It is always possible to return to current builder using Done()
	AsStatusRequestBinaryStateDeprecated() interface {
		StatusRequestBinaryStateDeprecatedBuilder
		Done() StatusRequestBuilder
	}
	// AsStatusRequestLevel converts this build to a subType of StatusRequest. It is always possible to return to current builder using Done()
	AsStatusRequestLevel() interface {
		StatusRequestLevelBuilder
		Done() StatusRequestBuilder
	}
	// Build builds the StatusRequest or returns an error if something is wrong
	PartialBuild() (StatusRequestContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() StatusRequestContract
	// Build builds the StatusRequest or returns an error if something is wrong
	Build() (StatusRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StatusRequest
}

// NewStatusRequestBuilder() creates a StatusRequestBuilder
func NewStatusRequestBuilder() StatusRequestBuilder {
	return &_StatusRequestBuilder{_StatusRequest: new(_StatusRequest)}
}

type _StatusRequestChildBuilder interface {
	utils.Copyable
	setParent(StatusRequestContract)
	buildForStatusRequest() (StatusRequest, error)
}

type _StatusRequestBuilder struct {
	*_StatusRequest

	childBuilder _StatusRequestChildBuilder

	err *utils.MultiError
}

var _ (StatusRequestBuilder) = (*_StatusRequestBuilder)(nil)

func (b *_StatusRequestBuilder) WithMandatoryFields(statusType byte) StatusRequestBuilder {
	return b.WithStatusType(statusType)
}

func (b *_StatusRequestBuilder) WithStatusType(statusType byte) StatusRequestBuilder {
	b.StatusType = statusType
	return b
}

func (b *_StatusRequestBuilder) PartialBuild() (StatusRequestContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StatusRequest.deepCopy(), nil
}

func (b *_StatusRequestBuilder) PartialMustBuild() StatusRequestContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StatusRequestBuilder) AsStatusRequestBinaryState() interface {
	StatusRequestBinaryStateBuilder
	Done() StatusRequestBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		StatusRequestBinaryStateBuilder
		Done() StatusRequestBuilder
	}); ok {
		return cb
	}
	cb := NewStatusRequestBinaryStateBuilder().(*_StatusRequestBinaryStateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_StatusRequestBuilder) AsStatusRequestBinaryStateDeprecated() interface {
	StatusRequestBinaryStateDeprecatedBuilder
	Done() StatusRequestBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		StatusRequestBinaryStateDeprecatedBuilder
		Done() StatusRequestBuilder
	}); ok {
		return cb
	}
	cb := NewStatusRequestBinaryStateDeprecatedBuilder().(*_StatusRequestBinaryStateDeprecatedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_StatusRequestBuilder) AsStatusRequestLevel() interface {
	StatusRequestLevelBuilder
	Done() StatusRequestBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		StatusRequestLevelBuilder
		Done() StatusRequestBuilder
	}); ok {
		return cb
	}
	cb := NewStatusRequestLevelBuilder().(*_StatusRequestLevelBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_StatusRequestBuilder) Build() (StatusRequest, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForStatusRequest()
}

func (b *_StatusRequestBuilder) MustBuild() StatusRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StatusRequestBuilder) DeepCopy() any {
	_copy := b.CreateStatusRequestBuilder().(*_StatusRequestBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_StatusRequestChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStatusRequestBuilder creates a StatusRequestBuilder
func (b *_StatusRequest) CreateStatusRequestBuilder() StatusRequestBuilder {
	if b == nil {
		return NewStatusRequestBuilder()
	}
	return &_StatusRequestBuilder{_StatusRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequest) GetStatusType() byte {
	return m.StatusType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStatusRequest(structType any) StatusRequest {
	if casted, ok := structType.(StatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequest) GetTypeName() string {
	return "StatusRequest"
}

func (m *_StatusRequest) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_StatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_StatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func StatusRequestParse[T StatusRequest](ctx context.Context, theBytes []byte) (T, error) {
	return StatusRequestParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusRequestParseWithBufferProducer[T StatusRequest]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := StatusRequestParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func StatusRequestParseWithBuffer[T StatusRequest](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_StatusRequest{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_StatusRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__statusRequest StatusRequest, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusType, err := ReadPeekField[byte](ctx, "statusType", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusType' field"))
	}
	m.StatusType = statusType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child StatusRequest
	switch {
	case statusType == 0x7A: // StatusRequestBinaryState
		if _child, err = new(_StatusRequestBinaryState).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type StatusRequestBinaryState for type-switch of StatusRequest")
		}
	case statusType == 0xFA: // StatusRequestBinaryStateDeprecated
		if _child, err = new(_StatusRequestBinaryStateDeprecated).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type StatusRequestBinaryStateDeprecated for type-switch of StatusRequest")
		}
	case statusType == 0x73: // StatusRequestLevel
		if _child, err = new(_StatusRequestLevel).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type StatusRequestLevel for type-switch of StatusRequest")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [statusType=%v]", statusType)
	}

	if closeErr := readBuffer.CloseContext("StatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequest")
	}

	return _child, nil
}

func (pm *_StatusRequest) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child StatusRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("StatusRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StatusRequest")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("StatusRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StatusRequest")
	}
	return nil
}

func (m *_StatusRequest) IsStatusRequest() {}

func (m *_StatusRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StatusRequest) deepCopy() *_StatusRequest {
	if m == nil {
		return nil
	}
	_StatusRequestCopy := &_StatusRequest{
		nil, // will be set by child
		m.StatusType,
	}
	return _StatusRequestCopy
}
