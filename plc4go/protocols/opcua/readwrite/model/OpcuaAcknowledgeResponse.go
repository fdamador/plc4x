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

// OpcuaAcknowledgeResponse is the corresponding interface of OpcuaAcknowledgeResponse
type OpcuaAcknowledgeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MessagePDU
	// GetVersion returns Version (property field)
	GetVersion() uint32
	// GetLimits returns Limits (property field)
	GetLimits() OpcuaProtocolLimits
	// IsOpcuaAcknowledgeResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaAcknowledgeResponse()
	// CreateBuilder creates a OpcuaAcknowledgeResponseBuilder
	CreateOpcuaAcknowledgeResponseBuilder() OpcuaAcknowledgeResponseBuilder
}

// _OpcuaAcknowledgeResponse is the data-structure of this message
type _OpcuaAcknowledgeResponse struct {
	MessagePDUContract
	Version uint32
	Limits  OpcuaProtocolLimits
}

var _ OpcuaAcknowledgeResponse = (*_OpcuaAcknowledgeResponse)(nil)
var _ MessagePDURequirements = (*_OpcuaAcknowledgeResponse)(nil)

// NewOpcuaAcknowledgeResponse factory function for _OpcuaAcknowledgeResponse
func NewOpcuaAcknowledgeResponse(chunk ChunkType, version uint32, limits OpcuaProtocolLimits, binary bool) *_OpcuaAcknowledgeResponse {
	if limits == nil {
		panic("limits of type OpcuaProtocolLimits for OpcuaAcknowledgeResponse must not be nil")
	}
	_result := &_OpcuaAcknowledgeResponse{
		MessagePDUContract: NewMessagePDU(chunk, binary),
		Version:            version,
		Limits:             limits,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaAcknowledgeResponseBuilder is a builder for OpcuaAcknowledgeResponse
type OpcuaAcknowledgeResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint32, limits OpcuaProtocolLimits) OpcuaAcknowledgeResponseBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint32) OpcuaAcknowledgeResponseBuilder
	// WithLimits adds Limits (property field)
	WithLimits(OpcuaProtocolLimits) OpcuaAcknowledgeResponseBuilder
	// WithLimitsBuilder adds Limits (property field) which is build by the builder
	WithLimitsBuilder(func(OpcuaProtocolLimitsBuilder) OpcuaProtocolLimitsBuilder) OpcuaAcknowledgeResponseBuilder
	// Build builds the OpcuaAcknowledgeResponse or returns an error if something is wrong
	Build() (OpcuaAcknowledgeResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaAcknowledgeResponse
}

// NewOpcuaAcknowledgeResponseBuilder() creates a OpcuaAcknowledgeResponseBuilder
func NewOpcuaAcknowledgeResponseBuilder() OpcuaAcknowledgeResponseBuilder {
	return &_OpcuaAcknowledgeResponseBuilder{_OpcuaAcknowledgeResponse: new(_OpcuaAcknowledgeResponse)}
}

type _OpcuaAcknowledgeResponseBuilder struct {
	*_OpcuaAcknowledgeResponse

	parentBuilder *_MessagePDUBuilder

	err *utils.MultiError
}

var _ (OpcuaAcknowledgeResponseBuilder) = (*_OpcuaAcknowledgeResponseBuilder)(nil)

func (b *_OpcuaAcknowledgeResponseBuilder) setParent(contract MessagePDUContract) {
	b.MessagePDUContract = contract
}

func (b *_OpcuaAcknowledgeResponseBuilder) WithMandatoryFields(version uint32, limits OpcuaProtocolLimits) OpcuaAcknowledgeResponseBuilder {
	return b.WithVersion(version).WithLimits(limits)
}

func (b *_OpcuaAcknowledgeResponseBuilder) WithVersion(version uint32) OpcuaAcknowledgeResponseBuilder {
	b.Version = version
	return b
}

func (b *_OpcuaAcknowledgeResponseBuilder) WithLimits(limits OpcuaProtocolLimits) OpcuaAcknowledgeResponseBuilder {
	b.Limits = limits
	return b
}

func (b *_OpcuaAcknowledgeResponseBuilder) WithLimitsBuilder(builderSupplier func(OpcuaProtocolLimitsBuilder) OpcuaProtocolLimitsBuilder) OpcuaAcknowledgeResponseBuilder {
	builder := builderSupplier(b.Limits.CreateOpcuaProtocolLimitsBuilder())
	var err error
	b.Limits, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OpcuaProtocolLimitsBuilder failed"))
	}
	return b
}

func (b *_OpcuaAcknowledgeResponseBuilder) Build() (OpcuaAcknowledgeResponse, error) {
	if b.Limits == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'limits' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaAcknowledgeResponse.deepCopy(), nil
}

func (b *_OpcuaAcknowledgeResponseBuilder) MustBuild() OpcuaAcknowledgeResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_OpcuaAcknowledgeResponseBuilder) Done() MessagePDUBuilder {
	return b.parentBuilder
}

func (b *_OpcuaAcknowledgeResponseBuilder) buildForMessagePDU() (MessagePDU, error) {
	return b.Build()
}

func (b *_OpcuaAcknowledgeResponseBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaAcknowledgeResponseBuilder().(*_OpcuaAcknowledgeResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaAcknowledgeResponseBuilder creates a OpcuaAcknowledgeResponseBuilder
func (b *_OpcuaAcknowledgeResponse) CreateOpcuaAcknowledgeResponseBuilder() OpcuaAcknowledgeResponseBuilder {
	if b == nil {
		return NewOpcuaAcknowledgeResponseBuilder()
	}
	return &_OpcuaAcknowledgeResponseBuilder{_OpcuaAcknowledgeResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaAcknowledgeResponse) GetMessageType() string {
	return "ACK"
}

func (m *_OpcuaAcknowledgeResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaAcknowledgeResponse) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaAcknowledgeResponse) GetVersion() uint32 {
	return m.Version
}

func (m *_OpcuaAcknowledgeResponse) GetLimits() OpcuaProtocolLimits {
	return m.Limits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaAcknowledgeResponse(structType any) OpcuaAcknowledgeResponse {
	if casted, ok := structType.(OpcuaAcknowledgeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaAcknowledgeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaAcknowledgeResponse) GetTypeName() string {
	return "OpcuaAcknowledgeResponse"
}

func (m *_OpcuaAcknowledgeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 32

	// Simple field (limits)
	lengthInBits += m.Limits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaAcknowledgeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaAcknowledgeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, response bool, binary bool) (__opcuaAcknowledgeResponse OpcuaAcknowledgeResponse, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaAcknowledgeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaAcknowledgeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	limits, err := ReadSimpleField[OpcuaProtocolLimits](ctx, "limits", ReadComplex[OpcuaProtocolLimits](OpcuaProtocolLimitsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limits' field"))
	}
	m.Limits = limits

	if closeErr := readBuffer.CloseContext("OpcuaAcknowledgeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaAcknowledgeResponse")
	}

	return m, nil
}

func (m *_OpcuaAcknowledgeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaAcknowledgeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaAcknowledgeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaAcknowledgeResponse")
		}

		if err := WriteSimpleField[uint32](ctx, "version", m.GetVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if err := WriteSimpleField[OpcuaProtocolLimits](ctx, "limits", m.GetLimits(), WriteComplex[OpcuaProtocolLimits](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limits' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaAcknowledgeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaAcknowledgeResponse")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaAcknowledgeResponse) IsOpcuaAcknowledgeResponse() {}

func (m *_OpcuaAcknowledgeResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaAcknowledgeResponse) deepCopy() *_OpcuaAcknowledgeResponse {
	if m == nil {
		return nil
	}
	_OpcuaAcknowledgeResponseCopy := &_OpcuaAcknowledgeResponse{
		m.MessagePDUContract.(*_MessagePDU).deepCopy(),
		m.Version,
		m.Limits.DeepCopy().(OpcuaProtocolLimits),
	}
	m.MessagePDUContract.(*_MessagePDU)._SubType = m
	return _OpcuaAcknowledgeResponseCopy
}

func (m *_OpcuaAcknowledgeResponse) String() string {
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
