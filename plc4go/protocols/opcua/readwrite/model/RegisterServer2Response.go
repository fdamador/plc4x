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

// RegisterServer2Response is the corresponding interface of RegisterServer2Response
type RegisterServer2Response interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetConfigurationResults returns ConfigurationResults (property field)
	GetConfigurationResults() []StatusCode
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsRegisterServer2Response is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRegisterServer2Response()
	// CreateBuilder creates a RegisterServer2ResponseBuilder
	CreateRegisterServer2ResponseBuilder() RegisterServer2ResponseBuilder
}

// _RegisterServer2Response is the data-structure of this message
type _RegisterServer2Response struct {
	ExtensionObjectDefinitionContract
	ResponseHeader       ResponseHeader
	ConfigurationResults []StatusCode
	DiagnosticInfos      []DiagnosticInfo
}

var _ RegisterServer2Response = (*_RegisterServer2Response)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterServer2Response)(nil)

// NewRegisterServer2Response factory function for _RegisterServer2Response
func NewRegisterServer2Response(responseHeader ResponseHeader, configurationResults []StatusCode, diagnosticInfos []DiagnosticInfo) *_RegisterServer2Response {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for RegisterServer2Response must not be nil")
	}
	_result := &_RegisterServer2Response{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		ConfigurationResults:              configurationResults,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RegisterServer2ResponseBuilder is a builder for RegisterServer2Response
type RegisterServer2ResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, configurationResults []StatusCode, diagnosticInfos []DiagnosticInfo) RegisterServer2ResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) RegisterServer2ResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) RegisterServer2ResponseBuilder
	// WithConfigurationResults adds ConfigurationResults (property field)
	WithConfigurationResults(...StatusCode) RegisterServer2ResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) RegisterServer2ResponseBuilder
	// Build builds the RegisterServer2Response or returns an error if something is wrong
	Build() (RegisterServer2Response, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RegisterServer2Response
}

// NewRegisterServer2ResponseBuilder() creates a RegisterServer2ResponseBuilder
func NewRegisterServer2ResponseBuilder() RegisterServer2ResponseBuilder {
	return &_RegisterServer2ResponseBuilder{_RegisterServer2Response: new(_RegisterServer2Response)}
}

type _RegisterServer2ResponseBuilder struct {
	*_RegisterServer2Response

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RegisterServer2ResponseBuilder) = (*_RegisterServer2ResponseBuilder)(nil)

func (b *_RegisterServer2ResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RegisterServer2ResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, configurationResults []StatusCode, diagnosticInfos []DiagnosticInfo) RegisterServer2ResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithConfigurationResults(configurationResults...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_RegisterServer2ResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) RegisterServer2ResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_RegisterServer2ResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) RegisterServer2ResponseBuilder {
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

func (b *_RegisterServer2ResponseBuilder) WithConfigurationResults(configurationResults ...StatusCode) RegisterServer2ResponseBuilder {
	b.ConfigurationResults = configurationResults
	return b
}

func (b *_RegisterServer2ResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) RegisterServer2ResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_RegisterServer2ResponseBuilder) Build() (RegisterServer2Response, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RegisterServer2Response.deepCopy(), nil
}

func (b *_RegisterServer2ResponseBuilder) MustBuild() RegisterServer2Response {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RegisterServer2ResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RegisterServer2ResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RegisterServer2ResponseBuilder) DeepCopy() any {
	_copy := b.CreateRegisterServer2ResponseBuilder().(*_RegisterServer2ResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRegisterServer2ResponseBuilder creates a RegisterServer2ResponseBuilder
func (b *_RegisterServer2Response) CreateRegisterServer2ResponseBuilder() RegisterServer2ResponseBuilder {
	if b == nil {
		return NewRegisterServer2ResponseBuilder()
	}
	return &_RegisterServer2ResponseBuilder{_RegisterServer2Response: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterServer2Response) GetExtensionId() int32 {
	return int32(12196)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterServer2Response) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterServer2Response) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_RegisterServer2Response) GetConfigurationResults() []StatusCode {
	return m.ConfigurationResults
}

func (m *_RegisterServer2Response) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRegisterServer2Response(structType any) RegisterServer2Response {
	if casted, ok := structType.(RegisterServer2Response); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterServer2Response); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterServer2Response) GetTypeName() string {
	return "RegisterServer2Response"
}

func (m *_RegisterServer2Response) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfConfigurationResults)
	lengthInBits += 32

	// Array field
	if len(m.ConfigurationResults) > 0 {
		for _curItem, element := range m.ConfigurationResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ConfigurationResults), _curItem)
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

func (m *_RegisterServer2Response) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterServer2Response) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__registerServer2Response RegisterServer2Response, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterServer2Response"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServer2Response")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfConfigurationResults, err := ReadImplicitField[int32](ctx, "noOfConfigurationResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConfigurationResults' field"))
	}
	_ = noOfConfigurationResults

	configurationResults, err := ReadCountArrayField[StatusCode](ctx, "configurationResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfConfigurationResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationResults' field"))
	}
	m.ConfigurationResults = configurationResults

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

	if closeErr := readBuffer.CloseContext("RegisterServer2Response"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServer2Response")
	}

	return m, nil
}

func (m *_RegisterServer2Response) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterServer2Response) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterServer2Response"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterServer2Response")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfConfigurationResults := int32(utils.InlineIf(bool((m.GetConfigurationResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetConfigurationResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfConfigurationResults", noOfConfigurationResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConfigurationResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "configurationResults", m.GetConfigurationResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationResults' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("RegisterServer2Response"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServer2Response")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterServer2Response) IsRegisterServer2Response() {}

func (m *_RegisterServer2Response) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RegisterServer2Response) deepCopy() *_RegisterServer2Response {
	if m == nil {
		return nil
	}
	_RegisterServer2ResponseCopy := &_RegisterServer2Response{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ResponseHeader),
		utils.DeepCopySlice[StatusCode, StatusCode](m.ConfigurationResults),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RegisterServer2ResponseCopy
}

func (m *_RegisterServer2Response) String() string {
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
