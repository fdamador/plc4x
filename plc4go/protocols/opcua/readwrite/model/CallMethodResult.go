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

// CallMethodResult is the corresponding interface of CallMethodResult
type CallMethodResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfInputArgumentResults returns NoOfInputArgumentResults (property field)
	GetNoOfInputArgumentResults() int32
	// GetInputArgumentResults returns InputArgumentResults (property field)
	GetInputArgumentResults() []StatusCode
	// GetNoOfInputArgumentDiagnosticInfos returns NoOfInputArgumentDiagnosticInfos (property field)
	GetNoOfInputArgumentDiagnosticInfos() int32
	// GetInputArgumentDiagnosticInfos returns InputArgumentDiagnosticInfos (property field)
	GetInputArgumentDiagnosticInfos() []DiagnosticInfo
	// GetNoOfOutputArguments returns NoOfOutputArguments (property field)
	GetNoOfOutputArguments() int32
	// GetOutputArguments returns OutputArguments (property field)
	GetOutputArguments() []Variant
	// IsCallMethodResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCallMethodResult()
}

// _CallMethodResult is the data-structure of this message
type _CallMethodResult struct {
	ExtensionObjectDefinitionContract
	StatusCode                       StatusCode
	NoOfInputArgumentResults         int32
	InputArgumentResults             []StatusCode
	NoOfInputArgumentDiagnosticInfos int32
	InputArgumentDiagnosticInfos     []DiagnosticInfo
	NoOfOutputArguments              int32
	OutputArguments                  []Variant
}

var _ CallMethodResult = (*_CallMethodResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CallMethodResult)(nil)

// NewCallMethodResult factory function for _CallMethodResult
func NewCallMethodResult(statusCode StatusCode, noOfInputArgumentResults int32, inputArgumentResults []StatusCode, noOfInputArgumentDiagnosticInfos int32, inputArgumentDiagnosticInfos []DiagnosticInfo, noOfOutputArguments int32, outputArguments []Variant) *_CallMethodResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for CallMethodResult must not be nil")
	}
	_result := &_CallMethodResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfInputArgumentResults:          noOfInputArgumentResults,
		InputArgumentResults:              inputArgumentResults,
		NoOfInputArgumentDiagnosticInfos:  noOfInputArgumentDiagnosticInfos,
		InputArgumentDiagnosticInfos:      inputArgumentDiagnosticInfos,
		NoOfOutputArguments:               noOfOutputArguments,
		OutputArguments:                   outputArguments,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallMethodResult) GetIdentifier() string {
	return "709"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallMethodResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallMethodResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_CallMethodResult) GetNoOfInputArgumentResults() int32 {
	return m.NoOfInputArgumentResults
}

func (m *_CallMethodResult) GetInputArgumentResults() []StatusCode {
	return m.InputArgumentResults
}

func (m *_CallMethodResult) GetNoOfInputArgumentDiagnosticInfos() int32 {
	return m.NoOfInputArgumentDiagnosticInfos
}

func (m *_CallMethodResult) GetInputArgumentDiagnosticInfos() []DiagnosticInfo {
	return m.InputArgumentDiagnosticInfos
}

func (m *_CallMethodResult) GetNoOfOutputArguments() int32 {
	return m.NoOfOutputArguments
}

func (m *_CallMethodResult) GetOutputArguments() []Variant {
	return m.OutputArguments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCallMethodResult(structType any) CallMethodResult {
	if casted, ok := structType.(CallMethodResult); ok {
		return casted
	}
	if casted, ok := structType.(*CallMethodResult); ok {
		return *casted
	}
	return nil
}

func (m *_CallMethodResult) GetTypeName() string {
	return "CallMethodResult"
}

func (m *_CallMethodResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfInputArgumentResults)
	lengthInBits += 32

	// Array field
	if len(m.InputArgumentResults) > 0 {
		for _curItem, element := range m.InputArgumentResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArgumentResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfInputArgumentDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.InputArgumentDiagnosticInfos) > 0 {
		for _curItem, element := range m.InputArgumentDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArgumentDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfOutputArguments)
	lengthInBits += 32

	// Array field
	if len(m.OutputArguments) > 0 {
		for _curItem, element := range m.OutputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OutputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallMethodResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CallMethodResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__callMethodResult CallMethodResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CallMethodResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallMethodResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfInputArgumentResults, err := ReadSimpleField(ctx, "noOfInputArgumentResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfInputArgumentResults' field"))
	}
	m.NoOfInputArgumentResults = noOfInputArgumentResults

	inputArgumentResults, err := ReadCountArrayField[StatusCode](ctx, "inputArgumentResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfInputArgumentResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inputArgumentResults' field"))
	}
	m.InputArgumentResults = inputArgumentResults

	noOfInputArgumentDiagnosticInfos, err := ReadSimpleField(ctx, "noOfInputArgumentDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfInputArgumentDiagnosticInfos' field"))
	}
	m.NoOfInputArgumentDiagnosticInfos = noOfInputArgumentDiagnosticInfos

	inputArgumentDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "inputArgumentDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfInputArgumentDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inputArgumentDiagnosticInfos' field"))
	}
	m.InputArgumentDiagnosticInfos = inputArgumentDiagnosticInfos

	noOfOutputArguments, err := ReadSimpleField(ctx, "noOfOutputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOutputArguments' field"))
	}
	m.NoOfOutputArguments = noOfOutputArguments

	outputArguments, err := ReadCountArrayField[Variant](ctx, "outputArguments", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfOutputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'outputArguments' field"))
	}
	m.OutputArguments = outputArguments

	if closeErr := readBuffer.CloseContext("CallMethodResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallMethodResult")
	}

	return m, nil
}

func (m *_CallMethodResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallMethodResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallMethodResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallMethodResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfInputArgumentResults", m.GetNoOfInputArgumentResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfInputArgumentResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "inputArgumentResults", m.GetInputArgumentResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'inputArgumentResults' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfInputArgumentDiagnosticInfos", m.GetNoOfInputArgumentDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfInputArgumentDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "inputArgumentDiagnosticInfos", m.GetInputArgumentDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'inputArgumentDiagnosticInfos' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOutputArguments", m.GetNoOfOutputArguments(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOutputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "outputArguments", m.GetOutputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'outputArguments' field")
		}

		if popErr := writeBuffer.PopContext("CallMethodResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallMethodResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallMethodResult) IsCallMethodResult() {}

func (m *_CallMethodResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CallMethodResult) deepCopy() *_CallMethodResult {
	if m == nil {
		return nil
	}
	_CallMethodResultCopy := &_CallMethodResult{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.StatusCode.DeepCopy().(StatusCode),
		m.NoOfInputArgumentResults,
		utils.DeepCopySlice[StatusCode, StatusCode](m.InputArgumentResults),
		m.NoOfInputArgumentDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.InputArgumentDiagnosticInfos),
		m.NoOfOutputArguments,
		utils.DeepCopySlice[Variant, Variant](m.OutputArguments),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CallMethodResultCopy
}

func (m *_CallMethodResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
