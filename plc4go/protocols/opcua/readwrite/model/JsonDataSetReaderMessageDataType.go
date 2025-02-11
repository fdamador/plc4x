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

// JsonDataSetReaderMessageDataType is the corresponding interface of JsonDataSetReaderMessageDataType
type JsonDataSetReaderMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNetworkMessageContentMask returns NetworkMessageContentMask (property field)
	GetNetworkMessageContentMask() JsonNetworkMessageContentMask
	// GetDataSetMessageContentMask returns DataSetMessageContentMask (property field)
	GetDataSetMessageContentMask() JsonDataSetMessageContentMask
	// IsJsonDataSetReaderMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsJsonDataSetReaderMessageDataType()
	// CreateBuilder creates a JsonDataSetReaderMessageDataTypeBuilder
	CreateJsonDataSetReaderMessageDataTypeBuilder() JsonDataSetReaderMessageDataTypeBuilder
}

// _JsonDataSetReaderMessageDataType is the data-structure of this message
type _JsonDataSetReaderMessageDataType struct {
	ExtensionObjectDefinitionContract
	NetworkMessageContentMask JsonNetworkMessageContentMask
	DataSetMessageContentMask JsonDataSetMessageContentMask
}

var _ JsonDataSetReaderMessageDataType = (*_JsonDataSetReaderMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_JsonDataSetReaderMessageDataType)(nil)

// NewJsonDataSetReaderMessageDataType factory function for _JsonDataSetReaderMessageDataType
func NewJsonDataSetReaderMessageDataType(networkMessageContentMask JsonNetworkMessageContentMask, dataSetMessageContentMask JsonDataSetMessageContentMask) *_JsonDataSetReaderMessageDataType {
	_result := &_JsonDataSetReaderMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NetworkMessageContentMask:         networkMessageContentMask,
		DataSetMessageContentMask:         dataSetMessageContentMask,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// JsonDataSetReaderMessageDataTypeBuilder is a builder for JsonDataSetReaderMessageDataType
type JsonDataSetReaderMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkMessageContentMask JsonNetworkMessageContentMask, dataSetMessageContentMask JsonDataSetMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder
	// WithNetworkMessageContentMask adds NetworkMessageContentMask (property field)
	WithNetworkMessageContentMask(JsonNetworkMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder
	// WithDataSetMessageContentMask adds DataSetMessageContentMask (property field)
	WithDataSetMessageContentMask(JsonDataSetMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder
	// Build builds the JsonDataSetReaderMessageDataType or returns an error if something is wrong
	Build() (JsonDataSetReaderMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() JsonDataSetReaderMessageDataType
}

// NewJsonDataSetReaderMessageDataTypeBuilder() creates a JsonDataSetReaderMessageDataTypeBuilder
func NewJsonDataSetReaderMessageDataTypeBuilder() JsonDataSetReaderMessageDataTypeBuilder {
	return &_JsonDataSetReaderMessageDataTypeBuilder{_JsonDataSetReaderMessageDataType: new(_JsonDataSetReaderMessageDataType)}
}

type _JsonDataSetReaderMessageDataTypeBuilder struct {
	*_JsonDataSetReaderMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (JsonDataSetReaderMessageDataTypeBuilder) = (*_JsonDataSetReaderMessageDataTypeBuilder)(nil)

func (b *_JsonDataSetReaderMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) WithMandatoryFields(networkMessageContentMask JsonNetworkMessageContentMask, dataSetMessageContentMask JsonDataSetMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder {
	return b.WithNetworkMessageContentMask(networkMessageContentMask).WithDataSetMessageContentMask(dataSetMessageContentMask)
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) WithNetworkMessageContentMask(networkMessageContentMask JsonNetworkMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder {
	b.NetworkMessageContentMask = networkMessageContentMask
	return b
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) WithDataSetMessageContentMask(dataSetMessageContentMask JsonDataSetMessageContentMask) JsonDataSetReaderMessageDataTypeBuilder {
	b.DataSetMessageContentMask = dataSetMessageContentMask
	return b
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) Build() (JsonDataSetReaderMessageDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._JsonDataSetReaderMessageDataType.deepCopy(), nil
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) MustBuild() JsonDataSetReaderMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_JsonDataSetReaderMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_JsonDataSetReaderMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateJsonDataSetReaderMessageDataTypeBuilder().(*_JsonDataSetReaderMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateJsonDataSetReaderMessageDataTypeBuilder creates a JsonDataSetReaderMessageDataTypeBuilder
func (b *_JsonDataSetReaderMessageDataType) CreateJsonDataSetReaderMessageDataTypeBuilder() JsonDataSetReaderMessageDataTypeBuilder {
	if b == nil {
		return NewJsonDataSetReaderMessageDataTypeBuilder()
	}
	return &_JsonDataSetReaderMessageDataTypeBuilder{_JsonDataSetReaderMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_JsonDataSetReaderMessageDataType) GetExtensionId() int32 {
	return int32(15667)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_JsonDataSetReaderMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_JsonDataSetReaderMessageDataType) GetNetworkMessageContentMask() JsonNetworkMessageContentMask {
	return m.NetworkMessageContentMask
}

func (m *_JsonDataSetReaderMessageDataType) GetDataSetMessageContentMask() JsonDataSetMessageContentMask {
	return m.DataSetMessageContentMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastJsonDataSetReaderMessageDataType(structType any) JsonDataSetReaderMessageDataType {
	if casted, ok := structType.(JsonDataSetReaderMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*JsonDataSetReaderMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_JsonDataSetReaderMessageDataType) GetTypeName() string {
	return "JsonDataSetReaderMessageDataType"
}

func (m *_JsonDataSetReaderMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (networkMessageContentMask)
	lengthInBits += 32

	// Simple field (dataSetMessageContentMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_JsonDataSetReaderMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_JsonDataSetReaderMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__jsonDataSetReaderMessageDataType JsonDataSetReaderMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("JsonDataSetReaderMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for JsonDataSetReaderMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkMessageContentMask, err := ReadEnumField[JsonNetworkMessageContentMask](ctx, "networkMessageContentMask", "JsonNetworkMessageContentMask", ReadEnum(JsonNetworkMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkMessageContentMask' field"))
	}
	m.NetworkMessageContentMask = networkMessageContentMask

	dataSetMessageContentMask, err := ReadEnumField[JsonDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "JsonDataSetMessageContentMask", ReadEnum(JsonDataSetMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetMessageContentMask' field"))
	}
	m.DataSetMessageContentMask = dataSetMessageContentMask

	if closeErr := readBuffer.CloseContext("JsonDataSetReaderMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for JsonDataSetReaderMessageDataType")
	}

	return m, nil
}

func (m *_JsonDataSetReaderMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_JsonDataSetReaderMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("JsonDataSetReaderMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for JsonDataSetReaderMessageDataType")
		}

		if err := WriteSimpleEnumField[JsonNetworkMessageContentMask](ctx, "networkMessageContentMask", "JsonNetworkMessageContentMask", m.GetNetworkMessageContentMask(), WriteEnum[JsonNetworkMessageContentMask, uint32](JsonNetworkMessageContentMask.GetValue, JsonNetworkMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'networkMessageContentMask' field")
		}

		if err := WriteSimpleEnumField[JsonDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "JsonDataSetMessageContentMask", m.GetDataSetMessageContentMask(), WriteEnum[JsonDataSetMessageContentMask, uint32](JsonDataSetMessageContentMask.GetValue, JsonDataSetMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetMessageContentMask' field")
		}

		if popErr := writeBuffer.PopContext("JsonDataSetReaderMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for JsonDataSetReaderMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_JsonDataSetReaderMessageDataType) IsJsonDataSetReaderMessageDataType() {}

func (m *_JsonDataSetReaderMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_JsonDataSetReaderMessageDataType) deepCopy() *_JsonDataSetReaderMessageDataType {
	if m == nil {
		return nil
	}
	_JsonDataSetReaderMessageDataTypeCopy := &_JsonDataSetReaderMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NetworkMessageContentMask,
		m.DataSetMessageContentMask,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _JsonDataSetReaderMessageDataTypeCopy
}

func (m *_JsonDataSetReaderMessageDataType) String() string {
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
