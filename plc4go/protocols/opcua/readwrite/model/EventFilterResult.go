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

// EventFilterResult is the corresponding interface of EventFilterResult
type EventFilterResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSelectClauseResults returns SelectClauseResults (property field)
	GetSelectClauseResults() []StatusCode
	// GetSelectClauseDiagnosticInfos returns SelectClauseDiagnosticInfos (property field)
	GetSelectClauseDiagnosticInfos() []DiagnosticInfo
	// GetWhereClauseResult returns WhereClauseResult (property field)
	GetWhereClauseResult() ContentFilterResult
	// IsEventFilterResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEventFilterResult()
	// CreateBuilder creates a EventFilterResultBuilder
	CreateEventFilterResultBuilder() EventFilterResultBuilder
}

// _EventFilterResult is the data-structure of this message
type _EventFilterResult struct {
	ExtensionObjectDefinitionContract
	SelectClauseResults         []StatusCode
	SelectClauseDiagnosticInfos []DiagnosticInfo
	WhereClauseResult           ContentFilterResult
}

var _ EventFilterResult = (*_EventFilterResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EventFilterResult)(nil)

// NewEventFilterResult factory function for _EventFilterResult
func NewEventFilterResult(selectClauseResults []StatusCode, selectClauseDiagnosticInfos []DiagnosticInfo, whereClauseResult ContentFilterResult) *_EventFilterResult {
	if whereClauseResult == nil {
		panic("whereClauseResult of type ContentFilterResult for EventFilterResult must not be nil")
	}
	_result := &_EventFilterResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SelectClauseResults:               selectClauseResults,
		SelectClauseDiagnosticInfos:       selectClauseDiagnosticInfos,
		WhereClauseResult:                 whereClauseResult,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EventFilterResultBuilder is a builder for EventFilterResult
type EventFilterResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(selectClauseResults []StatusCode, selectClauseDiagnosticInfos []DiagnosticInfo, whereClauseResult ContentFilterResult) EventFilterResultBuilder
	// WithSelectClauseResults adds SelectClauseResults (property field)
	WithSelectClauseResults(...StatusCode) EventFilterResultBuilder
	// WithSelectClauseDiagnosticInfos adds SelectClauseDiagnosticInfos (property field)
	WithSelectClauseDiagnosticInfos(...DiagnosticInfo) EventFilterResultBuilder
	// WithWhereClauseResult adds WhereClauseResult (property field)
	WithWhereClauseResult(ContentFilterResult) EventFilterResultBuilder
	// WithWhereClauseResultBuilder adds WhereClauseResult (property field) which is build by the builder
	WithWhereClauseResultBuilder(func(ContentFilterResultBuilder) ContentFilterResultBuilder) EventFilterResultBuilder
	// Build builds the EventFilterResult or returns an error if something is wrong
	Build() (EventFilterResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EventFilterResult
}

// NewEventFilterResultBuilder() creates a EventFilterResultBuilder
func NewEventFilterResultBuilder() EventFilterResultBuilder {
	return &_EventFilterResultBuilder{_EventFilterResult: new(_EventFilterResult)}
}

type _EventFilterResultBuilder struct {
	*_EventFilterResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EventFilterResultBuilder) = (*_EventFilterResultBuilder)(nil)

func (b *_EventFilterResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_EventFilterResultBuilder) WithMandatoryFields(selectClauseResults []StatusCode, selectClauseDiagnosticInfos []DiagnosticInfo, whereClauseResult ContentFilterResult) EventFilterResultBuilder {
	return b.WithSelectClauseResults(selectClauseResults...).WithSelectClauseDiagnosticInfos(selectClauseDiagnosticInfos...).WithWhereClauseResult(whereClauseResult)
}

func (b *_EventFilterResultBuilder) WithSelectClauseResults(selectClauseResults ...StatusCode) EventFilterResultBuilder {
	b.SelectClauseResults = selectClauseResults
	return b
}

func (b *_EventFilterResultBuilder) WithSelectClauseDiagnosticInfos(selectClauseDiagnosticInfos ...DiagnosticInfo) EventFilterResultBuilder {
	b.SelectClauseDiagnosticInfos = selectClauseDiagnosticInfos
	return b
}

func (b *_EventFilterResultBuilder) WithWhereClauseResult(whereClauseResult ContentFilterResult) EventFilterResultBuilder {
	b.WhereClauseResult = whereClauseResult
	return b
}

func (b *_EventFilterResultBuilder) WithWhereClauseResultBuilder(builderSupplier func(ContentFilterResultBuilder) ContentFilterResultBuilder) EventFilterResultBuilder {
	builder := builderSupplier(b.WhereClauseResult.CreateContentFilterResultBuilder())
	var err error
	b.WhereClauseResult, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ContentFilterResultBuilder failed"))
	}
	return b
}

func (b *_EventFilterResultBuilder) Build() (EventFilterResult, error) {
	if b.WhereClauseResult == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'whereClauseResult' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EventFilterResult.deepCopy(), nil
}

func (b *_EventFilterResultBuilder) MustBuild() EventFilterResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_EventFilterResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_EventFilterResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EventFilterResultBuilder) DeepCopy() any {
	_copy := b.CreateEventFilterResultBuilder().(*_EventFilterResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEventFilterResultBuilder creates a EventFilterResultBuilder
func (b *_EventFilterResult) CreateEventFilterResultBuilder() EventFilterResultBuilder {
	if b == nil {
		return NewEventFilterResultBuilder()
	}
	return &_EventFilterResultBuilder{_EventFilterResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EventFilterResult) GetExtensionId() int32 {
	return int32(736)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EventFilterResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EventFilterResult) GetSelectClauseResults() []StatusCode {
	return m.SelectClauseResults
}

func (m *_EventFilterResult) GetSelectClauseDiagnosticInfos() []DiagnosticInfo {
	return m.SelectClauseDiagnosticInfos
}

func (m *_EventFilterResult) GetWhereClauseResult() ContentFilterResult {
	return m.WhereClauseResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEventFilterResult(structType any) EventFilterResult {
	if casted, ok := structType.(EventFilterResult); ok {
		return casted
	}
	if casted, ok := structType.(*EventFilterResult); ok {
		return *casted
	}
	return nil
}

func (m *_EventFilterResult) GetTypeName() string {
	return "EventFilterResult"
}

func (m *_EventFilterResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfSelectClauseResults)
	lengthInBits += 32

	// Array field
	if len(m.SelectClauseResults) > 0 {
		for _curItem, element := range m.SelectClauseResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SelectClauseResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfSelectClauseDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.SelectClauseDiagnosticInfos) > 0 {
		for _curItem, element := range m.SelectClauseDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SelectClauseDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (whereClauseResult)
	lengthInBits += m.WhereClauseResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EventFilterResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EventFilterResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__eventFilterResult EventFilterResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EventFilterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EventFilterResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfSelectClauseResults, err := ReadImplicitField[int32](ctx, "noOfSelectClauseResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSelectClauseResults' field"))
	}
	_ = noOfSelectClauseResults

	selectClauseResults, err := ReadCountArrayField[StatusCode](ctx, "selectClauseResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfSelectClauseResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectClauseResults' field"))
	}
	m.SelectClauseResults = selectClauseResults

	noOfSelectClauseDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfSelectClauseDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSelectClauseDiagnosticInfos' field"))
	}
	_ = noOfSelectClauseDiagnosticInfos

	selectClauseDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "selectClauseDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfSelectClauseDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectClauseDiagnosticInfos' field"))
	}
	m.SelectClauseDiagnosticInfos = selectClauseDiagnosticInfos

	whereClauseResult, err := ReadSimpleField[ContentFilterResult](ctx, "whereClauseResult", ReadComplex[ContentFilterResult](ExtensionObjectDefinitionParseWithBufferProducer[ContentFilterResult]((int32)(int32(609))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'whereClauseResult' field"))
	}
	m.WhereClauseResult = whereClauseResult

	if closeErr := readBuffer.CloseContext("EventFilterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EventFilterResult")
	}

	return m, nil
}

func (m *_EventFilterResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EventFilterResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EventFilterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EventFilterResult")
		}
		noOfSelectClauseResults := int32(utils.InlineIf(bool((m.GetSelectClauseResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSelectClauseResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSelectClauseResults", noOfSelectClauseResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSelectClauseResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "selectClauseResults", m.GetSelectClauseResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'selectClauseResults' field")
		}
		noOfSelectClauseDiagnosticInfos := int32(utils.InlineIf(bool((m.GetSelectClauseDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSelectClauseDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSelectClauseDiagnosticInfos", noOfSelectClauseDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSelectClauseDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "selectClauseDiagnosticInfos", m.GetSelectClauseDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'selectClauseDiagnosticInfos' field")
		}

		if err := WriteSimpleField[ContentFilterResult](ctx, "whereClauseResult", m.GetWhereClauseResult(), WriteComplex[ContentFilterResult](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'whereClauseResult' field")
		}

		if popErr := writeBuffer.PopContext("EventFilterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EventFilterResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EventFilterResult) IsEventFilterResult() {}

func (m *_EventFilterResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EventFilterResult) deepCopy() *_EventFilterResult {
	if m == nil {
		return nil
	}
	_EventFilterResultCopy := &_EventFilterResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[StatusCode, StatusCode](m.SelectClauseResults),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.SelectClauseDiagnosticInfos),
		m.WhereClauseResult.DeepCopy().(ContentFilterResult),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EventFilterResultCopy
}

func (m *_EventFilterResult) String() string {
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
