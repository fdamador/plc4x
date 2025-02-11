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

// HistoryModifiedData is the corresponding interface of HistoryModifiedData
type HistoryModifiedData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDataValues returns DataValues (property field)
	GetDataValues() []DataValue
	// GetModificationInfos returns ModificationInfos (property field)
	GetModificationInfos() []ModificationInfo
	// IsHistoryModifiedData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryModifiedData()
	// CreateBuilder creates a HistoryModifiedDataBuilder
	CreateHistoryModifiedDataBuilder() HistoryModifiedDataBuilder
}

// _HistoryModifiedData is the data-structure of this message
type _HistoryModifiedData struct {
	ExtensionObjectDefinitionContract
	DataValues        []DataValue
	ModificationInfos []ModificationInfo
}

var _ HistoryModifiedData = (*_HistoryModifiedData)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryModifiedData)(nil)

// NewHistoryModifiedData factory function for _HistoryModifiedData
func NewHistoryModifiedData(dataValues []DataValue, modificationInfos []ModificationInfo) *_HistoryModifiedData {
	_result := &_HistoryModifiedData{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataValues:                        dataValues,
		ModificationInfos:                 modificationInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryModifiedDataBuilder is a builder for HistoryModifiedData
type HistoryModifiedDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataValues []DataValue, modificationInfos []ModificationInfo) HistoryModifiedDataBuilder
	// WithDataValues adds DataValues (property field)
	WithDataValues(...DataValue) HistoryModifiedDataBuilder
	// WithModificationInfos adds ModificationInfos (property field)
	WithModificationInfos(...ModificationInfo) HistoryModifiedDataBuilder
	// Build builds the HistoryModifiedData or returns an error if something is wrong
	Build() (HistoryModifiedData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryModifiedData
}

// NewHistoryModifiedDataBuilder() creates a HistoryModifiedDataBuilder
func NewHistoryModifiedDataBuilder() HistoryModifiedDataBuilder {
	return &_HistoryModifiedDataBuilder{_HistoryModifiedData: new(_HistoryModifiedData)}
}

type _HistoryModifiedDataBuilder struct {
	*_HistoryModifiedData

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryModifiedDataBuilder) = (*_HistoryModifiedDataBuilder)(nil)

func (b *_HistoryModifiedDataBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_HistoryModifiedDataBuilder) WithMandatoryFields(dataValues []DataValue, modificationInfos []ModificationInfo) HistoryModifiedDataBuilder {
	return b.WithDataValues(dataValues...).WithModificationInfos(modificationInfos...)
}

func (b *_HistoryModifiedDataBuilder) WithDataValues(dataValues ...DataValue) HistoryModifiedDataBuilder {
	b.DataValues = dataValues
	return b
}

func (b *_HistoryModifiedDataBuilder) WithModificationInfos(modificationInfos ...ModificationInfo) HistoryModifiedDataBuilder {
	b.ModificationInfos = modificationInfos
	return b
}

func (b *_HistoryModifiedDataBuilder) Build() (HistoryModifiedData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryModifiedData.deepCopy(), nil
}

func (b *_HistoryModifiedDataBuilder) MustBuild() HistoryModifiedData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_HistoryModifiedDataBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_HistoryModifiedDataBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryModifiedDataBuilder) DeepCopy() any {
	_copy := b.CreateHistoryModifiedDataBuilder().(*_HistoryModifiedDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryModifiedDataBuilder creates a HistoryModifiedDataBuilder
func (b *_HistoryModifiedData) CreateHistoryModifiedDataBuilder() HistoryModifiedDataBuilder {
	if b == nil {
		return NewHistoryModifiedDataBuilder()
	}
	return &_HistoryModifiedDataBuilder{_HistoryModifiedData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryModifiedData) GetExtensionId() int32 {
	return int32(11219)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryModifiedData) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryModifiedData) GetDataValues() []DataValue {
	return m.DataValues
}

func (m *_HistoryModifiedData) GetModificationInfos() []ModificationInfo {
	return m.ModificationInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryModifiedData(structType any) HistoryModifiedData {
	if casted, ok := structType.(HistoryModifiedData); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryModifiedData); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryModifiedData) GetTypeName() string {
	return "HistoryModifiedData"
}

func (m *_HistoryModifiedData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfDataValues)
	lengthInBits += 32

	// Array field
	if len(m.DataValues) > 0 {
		for _curItem, element := range m.DataValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfModificationInfos)
	lengthInBits += 32

	// Array field
	if len(m.ModificationInfos) > 0 {
		for _curItem, element := range m.ModificationInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ModificationInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryModifiedData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryModifiedData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyModifiedData HistoryModifiedData, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryModifiedData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryModifiedData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfDataValues, err := ReadImplicitField[int32](ctx, "noOfDataValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataValues' field"))
	}
	_ = noOfDataValues

	dataValues, err := ReadCountArrayField[DataValue](ctx, "dataValues", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(noOfDataValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataValues' field"))
	}
	m.DataValues = dataValues

	noOfModificationInfos, err := ReadImplicitField[int32](ctx, "noOfModificationInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfModificationInfos' field"))
	}
	_ = noOfModificationInfos

	modificationInfos, err := ReadCountArrayField[ModificationInfo](ctx, "modificationInfos", ReadComplex[ModificationInfo](ExtensionObjectDefinitionParseWithBufferProducer[ModificationInfo]((int32)(int32(11218))), readBuffer), uint64(noOfModificationInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modificationInfos' field"))
	}
	m.ModificationInfos = modificationInfos

	if closeErr := readBuffer.CloseContext("HistoryModifiedData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryModifiedData")
	}

	return m, nil
}

func (m *_HistoryModifiedData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryModifiedData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryModifiedData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryModifiedData")
		}
		noOfDataValues := int32(utils.InlineIf(bool((m.GetDataValues()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDataValues()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDataValues", noOfDataValues, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataValues", m.GetDataValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataValues' field")
		}
		noOfModificationInfos := int32(utils.InlineIf(bool((m.GetModificationInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetModificationInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfModificationInfos", noOfModificationInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfModificationInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "modificationInfos", m.GetModificationInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'modificationInfos' field")
		}

		if popErr := writeBuffer.PopContext("HistoryModifiedData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryModifiedData")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryModifiedData) IsHistoryModifiedData() {}

func (m *_HistoryModifiedData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryModifiedData) deepCopy() *_HistoryModifiedData {
	if m == nil {
		return nil
	}
	_HistoryModifiedDataCopy := &_HistoryModifiedData{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[DataValue, DataValue](m.DataValues),
		utils.DeepCopySlice[ModificationInfo, ModificationInfo](m.ModificationInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryModifiedDataCopy
}

func (m *_HistoryModifiedData) String() string {
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
