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

// FieldTargetDataType is the corresponding interface of FieldTargetDataType
type FieldTargetDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDataSetFieldId returns DataSetFieldId (property field)
	GetDataSetFieldId() GuidValue
	// GetReceiverIndexRange returns ReceiverIndexRange (property field)
	GetReceiverIndexRange() PascalString
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetWriteIndexRange returns WriteIndexRange (property field)
	GetWriteIndexRange() PascalString
	// GetOverrideValueHandling returns OverrideValueHandling (property field)
	GetOverrideValueHandling() OverrideValueHandling
	// GetOverrideValue returns OverrideValue (property field)
	GetOverrideValue() Variant
	// IsFieldTargetDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFieldTargetDataType()
}

// _FieldTargetDataType is the data-structure of this message
type _FieldTargetDataType struct {
	ExtensionObjectDefinitionContract
	DataSetFieldId        GuidValue
	ReceiverIndexRange    PascalString
	TargetNodeId          NodeId
	AttributeId           uint32
	WriteIndexRange       PascalString
	OverrideValueHandling OverrideValueHandling
	OverrideValue         Variant
}

var _ FieldTargetDataType = (*_FieldTargetDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FieldTargetDataType)(nil)

// NewFieldTargetDataType factory function for _FieldTargetDataType
func NewFieldTargetDataType(dataSetFieldId GuidValue, receiverIndexRange PascalString, targetNodeId NodeId, attributeId uint32, writeIndexRange PascalString, overrideValueHandling OverrideValueHandling, overrideValue Variant) *_FieldTargetDataType {
	if dataSetFieldId == nil {
		panic("dataSetFieldId of type GuidValue for FieldTargetDataType must not be nil")
	}
	if receiverIndexRange == nil {
		panic("receiverIndexRange of type PascalString for FieldTargetDataType must not be nil")
	}
	if targetNodeId == nil {
		panic("targetNodeId of type NodeId for FieldTargetDataType must not be nil")
	}
	if writeIndexRange == nil {
		panic("writeIndexRange of type PascalString for FieldTargetDataType must not be nil")
	}
	if overrideValue == nil {
		panic("overrideValue of type Variant for FieldTargetDataType must not be nil")
	}
	_result := &_FieldTargetDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataSetFieldId:                    dataSetFieldId,
		ReceiverIndexRange:                receiverIndexRange,
		TargetNodeId:                      targetNodeId,
		AttributeId:                       attributeId,
		WriteIndexRange:                   writeIndexRange,
		OverrideValueHandling:             overrideValueHandling,
		OverrideValue:                     overrideValue,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FieldTargetDataType) GetIdentifier() string {
	return "14746"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FieldTargetDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FieldTargetDataType) GetDataSetFieldId() GuidValue {
	return m.DataSetFieldId
}

func (m *_FieldTargetDataType) GetReceiverIndexRange() PascalString {
	return m.ReceiverIndexRange
}

func (m *_FieldTargetDataType) GetTargetNodeId() NodeId {
	return m.TargetNodeId
}

func (m *_FieldTargetDataType) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_FieldTargetDataType) GetWriteIndexRange() PascalString {
	return m.WriteIndexRange
}

func (m *_FieldTargetDataType) GetOverrideValueHandling() OverrideValueHandling {
	return m.OverrideValueHandling
}

func (m *_FieldTargetDataType) GetOverrideValue() Variant {
	return m.OverrideValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFieldTargetDataType(structType any) FieldTargetDataType {
	if casted, ok := structType.(FieldTargetDataType); ok {
		return casted
	}
	if casted, ok := structType.(*FieldTargetDataType); ok {
		return *casted
	}
	return nil
}

func (m *_FieldTargetDataType) GetTypeName() string {
	return "FieldTargetDataType"
}

func (m *_FieldTargetDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (dataSetFieldId)
	lengthInBits += m.DataSetFieldId.GetLengthInBits(ctx)

	// Simple field (receiverIndexRange)
	lengthInBits += m.ReceiverIndexRange.GetLengthInBits(ctx)

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (writeIndexRange)
	lengthInBits += m.WriteIndexRange.GetLengthInBits(ctx)

	// Simple field (overrideValueHandling)
	lengthInBits += 32

	// Simple field (overrideValue)
	lengthInBits += m.OverrideValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_FieldTargetDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FieldTargetDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__fieldTargetDataType FieldTargetDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FieldTargetDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FieldTargetDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataSetFieldId, err := ReadSimpleField[GuidValue](ctx, "dataSetFieldId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldId' field"))
	}
	m.DataSetFieldId = dataSetFieldId

	receiverIndexRange, err := ReadSimpleField[PascalString](ctx, "receiverIndexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiverIndexRange' field"))
	}
	m.ReceiverIndexRange = receiverIndexRange

	targetNodeId, err := ReadSimpleField[NodeId](ctx, "targetNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeId' field"))
	}
	m.TargetNodeId = targetNodeId

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	writeIndexRange, err := ReadSimpleField[PascalString](ctx, "writeIndexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeIndexRange' field"))
	}
	m.WriteIndexRange = writeIndexRange

	overrideValueHandling, err := ReadEnumField[OverrideValueHandling](ctx, "overrideValueHandling", "OverrideValueHandling", ReadEnum(OverrideValueHandlingByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'overrideValueHandling' field"))
	}
	m.OverrideValueHandling = overrideValueHandling

	overrideValue, err := ReadSimpleField[Variant](ctx, "overrideValue", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'overrideValue' field"))
	}
	m.OverrideValue = overrideValue

	if closeErr := readBuffer.CloseContext("FieldTargetDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FieldTargetDataType")
	}

	return m, nil
}

func (m *_FieldTargetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FieldTargetDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FieldTargetDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FieldTargetDataType")
		}

		if err := WriteSimpleField[GuidValue](ctx, "dataSetFieldId", m.GetDataSetFieldId(), WriteComplex[GuidValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetFieldId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "receiverIndexRange", m.GetReceiverIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'receiverIndexRange' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "targetNodeId", m.GetTargetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNodeId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "writeIndexRange", m.GetWriteIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeIndexRange' field")
		}

		if err := WriteSimpleEnumField[OverrideValueHandling](ctx, "overrideValueHandling", "OverrideValueHandling", m.GetOverrideValueHandling(), WriteEnum[OverrideValueHandling, uint32](OverrideValueHandling.GetValue, OverrideValueHandling.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'overrideValueHandling' field")
		}

		if err := WriteSimpleField[Variant](ctx, "overrideValue", m.GetOverrideValue(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'overrideValue' field")
		}

		if popErr := writeBuffer.PopContext("FieldTargetDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FieldTargetDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FieldTargetDataType) IsFieldTargetDataType() {}

func (m *_FieldTargetDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FieldTargetDataType) deepCopy() *_FieldTargetDataType {
	if m == nil {
		return nil
	}
	_FieldTargetDataTypeCopy := &_FieldTargetDataType{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.DataSetFieldId.DeepCopy().(GuidValue),
		m.ReceiverIndexRange.DeepCopy().(PascalString),
		m.TargetNodeId.DeepCopy().(NodeId),
		m.AttributeId,
		m.WriteIndexRange.DeepCopy().(PascalString),
		m.OverrideValueHandling,
		m.OverrideValue.DeepCopy().(Variant),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FieldTargetDataTypeCopy
}

func (m *_FieldTargetDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
