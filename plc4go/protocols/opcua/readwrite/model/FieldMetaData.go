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

// FieldMetaData is the corresponding interface of FieldMetaData
type FieldMetaData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetFieldFlags returns FieldFlags (property field)
	GetFieldFlags() DataSetFieldFlags
	// GetBuiltInType returns BuiltInType (property field)
	GetBuiltInType() uint8
	// GetDataType returns DataType (property field)
	GetDataType() NodeId
	// GetValueRank returns ValueRank (property field)
	GetValueRank() int32
	// GetNoOfArrayDimensions returns NoOfArrayDimensions (property field)
	GetNoOfArrayDimensions() int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []uint32
	// GetMaxStringLength returns MaxStringLength (property field)
	GetMaxStringLength() uint32
	// GetDataSetFieldId returns DataSetFieldId (property field)
	GetDataSetFieldId() GuidValue
	// GetNoOfProperties returns NoOfProperties (property field)
	GetNoOfProperties() int32
	// GetProperties returns Properties (property field)
	GetProperties() []ExtensionObjectDefinition
	// IsFieldMetaData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFieldMetaData()
}

// _FieldMetaData is the data-structure of this message
type _FieldMetaData struct {
	ExtensionObjectDefinitionContract
	Name                PascalString
	Description         LocalizedText
	FieldFlags          DataSetFieldFlags
	BuiltInType         uint8
	DataType            NodeId
	ValueRank           int32
	NoOfArrayDimensions int32
	ArrayDimensions     []uint32
	MaxStringLength     uint32
	DataSetFieldId      GuidValue
	NoOfProperties      int32
	Properties          []ExtensionObjectDefinition
}

var _ FieldMetaData = (*_FieldMetaData)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FieldMetaData)(nil)

// NewFieldMetaData factory function for _FieldMetaData
func NewFieldMetaData(name PascalString, description LocalizedText, fieldFlags DataSetFieldFlags, builtInType uint8, dataType NodeId, valueRank int32, noOfArrayDimensions int32, arrayDimensions []uint32, maxStringLength uint32, dataSetFieldId GuidValue, noOfProperties int32, properties []ExtensionObjectDefinition) *_FieldMetaData {
	if name == nil {
		panic("name of type PascalString for FieldMetaData must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for FieldMetaData must not be nil")
	}
	if dataType == nil {
		panic("dataType of type NodeId for FieldMetaData must not be nil")
	}
	if dataSetFieldId == nil {
		panic("dataSetFieldId of type GuidValue for FieldMetaData must not be nil")
	}
	_result := &_FieldMetaData{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Description:                       description,
		FieldFlags:                        fieldFlags,
		BuiltInType:                       builtInType,
		DataType:                          dataType,
		ValueRank:                         valueRank,
		NoOfArrayDimensions:               noOfArrayDimensions,
		ArrayDimensions:                   arrayDimensions,
		MaxStringLength:                   maxStringLength,
		DataSetFieldId:                    dataSetFieldId,
		NoOfProperties:                    noOfProperties,
		Properties:                        properties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FieldMetaData) GetIdentifier() string {
	return "14526"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FieldMetaData) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FieldMetaData) GetName() PascalString {
	return m.Name
}

func (m *_FieldMetaData) GetDescription() LocalizedText {
	return m.Description
}

func (m *_FieldMetaData) GetFieldFlags() DataSetFieldFlags {
	return m.FieldFlags
}

func (m *_FieldMetaData) GetBuiltInType() uint8 {
	return m.BuiltInType
}

func (m *_FieldMetaData) GetDataType() NodeId {
	return m.DataType
}

func (m *_FieldMetaData) GetValueRank() int32 {
	return m.ValueRank
}

func (m *_FieldMetaData) GetNoOfArrayDimensions() int32 {
	return m.NoOfArrayDimensions
}

func (m *_FieldMetaData) GetArrayDimensions() []uint32 {
	return m.ArrayDimensions
}

func (m *_FieldMetaData) GetMaxStringLength() uint32 {
	return m.MaxStringLength
}

func (m *_FieldMetaData) GetDataSetFieldId() GuidValue {
	return m.DataSetFieldId
}

func (m *_FieldMetaData) GetNoOfProperties() int32 {
	return m.NoOfProperties
}

func (m *_FieldMetaData) GetProperties() []ExtensionObjectDefinition {
	return m.Properties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFieldMetaData(structType any) FieldMetaData {
	if casted, ok := structType.(FieldMetaData); ok {
		return casted
	}
	if casted, ok := structType.(*FieldMetaData); ok {
		return *casted
	}
	return nil
}

func (m *_FieldMetaData) GetTypeName() string {
	return "FieldMetaData"
}

func (m *_FieldMetaData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (fieldFlags)
	lengthInBits += 16

	// Simple field (builtInType)
	lengthInBits += 8

	// Simple field (dataType)
	lengthInBits += m.DataType.GetLengthInBits(ctx)

	// Simple field (valueRank)
	lengthInBits += 32

	// Simple field (noOfArrayDimensions)
	lengthInBits += 32

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 32 * uint16(len(m.ArrayDimensions))
	}

	// Simple field (maxStringLength)
	lengthInBits += 32

	// Simple field (dataSetFieldId)
	lengthInBits += m.DataSetFieldId.GetLengthInBits(ctx)

	// Simple field (noOfProperties)
	lengthInBits += 32

	// Array field
	if len(m.Properties) > 0 {
		for _curItem, element := range m.Properties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Properties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FieldMetaData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FieldMetaData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__fieldMetaData FieldMetaData, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FieldMetaData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FieldMetaData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	fieldFlags, err := ReadEnumField[DataSetFieldFlags](ctx, "fieldFlags", "DataSetFieldFlags", ReadEnum(DataSetFieldFlagsByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fieldFlags' field"))
	}
	m.FieldFlags = fieldFlags

	builtInType, err := ReadSimpleField(ctx, "builtInType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'builtInType' field"))
	}
	m.BuiltInType = builtInType

	dataType, err := ReadSimpleField[NodeId](ctx, "dataType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	valueRank, err := ReadSimpleField(ctx, "valueRank", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueRank' field"))
	}
	m.ValueRank = valueRank

	noOfArrayDimensions, err := ReadSimpleField(ctx, "noOfArrayDimensions", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfArrayDimensions' field"))
	}
	m.NoOfArrayDimensions = noOfArrayDimensions

	arrayDimensions, err := ReadCountArrayField[uint32](ctx, "arrayDimensions", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfArrayDimensions))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	maxStringLength, err := ReadSimpleField(ctx, "maxStringLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxStringLength' field"))
	}
	m.MaxStringLength = maxStringLength

	dataSetFieldId, err := ReadSimpleField[GuidValue](ctx, "dataSetFieldId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldId' field"))
	}
	m.DataSetFieldId = dataSetFieldId

	noOfProperties, err := ReadSimpleField(ctx, "noOfProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfProperties' field"))
	}
	m.NoOfProperties = noOfProperties

	properties, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "properties", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("14535")), readBuffer), uint64(noOfProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'properties' field"))
	}
	m.Properties = properties

	if closeErr := readBuffer.CloseContext("FieldMetaData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FieldMetaData")
	}

	return m, nil
}

func (m *_FieldMetaData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FieldMetaData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FieldMetaData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FieldMetaData")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleEnumField[DataSetFieldFlags](ctx, "fieldFlags", "DataSetFieldFlags", m.GetFieldFlags(), WriteEnum[DataSetFieldFlags, uint16](DataSetFieldFlags.GetValue, DataSetFieldFlags.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
			return errors.Wrap(err, "Error serializing 'fieldFlags' field")
		}

		if err := WriteSimpleField[uint8](ctx, "builtInType", m.GetBuiltInType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'builtInType' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataType", m.GetDataType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataType' field")
		}

		if err := WriteSimpleField[int32](ctx, "valueRank", m.GetValueRank(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'valueRank' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfArrayDimensions", m.GetNoOfArrayDimensions(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfArrayDimensions' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "arrayDimensions", m.GetArrayDimensions(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxStringLength", m.GetMaxStringLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxStringLength' field")
		}

		if err := WriteSimpleField[GuidValue](ctx, "dataSetFieldId", m.GetDataSetFieldId(), WriteComplex[GuidValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetFieldId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfProperties", m.GetNoOfProperties(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "properties", m.GetProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'properties' field")
		}

		if popErr := writeBuffer.PopContext("FieldMetaData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FieldMetaData")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FieldMetaData) IsFieldMetaData() {}

func (m *_FieldMetaData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FieldMetaData) deepCopy() *_FieldMetaData {
	if m == nil {
		return nil
	}
	_FieldMetaDataCopy := &_FieldMetaData{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.Name.DeepCopy().(PascalString),
		m.Description.DeepCopy().(LocalizedText),
		m.FieldFlags,
		m.BuiltInType,
		m.DataType.DeepCopy().(NodeId),
		m.ValueRank,
		m.NoOfArrayDimensions,
		utils.DeepCopySlice[uint32, uint32](m.ArrayDimensions),
		m.MaxStringLength,
		m.DataSetFieldId.DeepCopy().(GuidValue),
		m.NoOfProperties,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Properties),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FieldMetaDataCopy
}

func (m *_FieldMetaData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
