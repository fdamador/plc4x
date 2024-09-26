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

// PubSubGroupDataType is the corresponding interface of PubSubGroupDataType
type PubSubGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetSecurityGroupId returns SecurityGroupId (property field)
	GetSecurityGroupId() PascalString
	// GetNoOfSecurityKeyServices returns NoOfSecurityKeyServices (property field)
	GetNoOfSecurityKeyServices() int32
	// GetSecurityKeyServices returns SecurityKeyServices (property field)
	GetSecurityKeyServices() []ExtensionObjectDefinition
	// GetMaxNetworkMessageSize returns MaxNetworkMessageSize (property field)
	GetMaxNetworkMessageSize() uint32
	// GetNoOfGroupProperties returns NoOfGroupProperties (property field)
	GetNoOfGroupProperties() int32
	// GetGroupProperties returns GroupProperties (property field)
	GetGroupProperties() []ExtensionObjectDefinition
	// IsPubSubGroupDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubGroupDataType()
}

// _PubSubGroupDataType is the data-structure of this message
type _PubSubGroupDataType struct {
	ExtensionObjectDefinitionContract
	Name                    PascalString
	Enabled                 bool
	SecurityMode            MessageSecurityMode
	SecurityGroupId         PascalString
	NoOfSecurityKeyServices int32
	SecurityKeyServices     []ExtensionObjectDefinition
	MaxNetworkMessageSize   uint32
	NoOfGroupProperties     int32
	GroupProperties         []ExtensionObjectDefinition
	// Reserved Fields
	reservedField0 *uint8
}

var _ PubSubGroupDataType = (*_PubSubGroupDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubGroupDataType)(nil)

// NewPubSubGroupDataType factory function for _PubSubGroupDataType
func NewPubSubGroupDataType(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, noOfSecurityKeyServices int32, securityKeyServices []ExtensionObjectDefinition, maxNetworkMessageSize uint32, noOfGroupProperties int32, groupProperties []ExtensionObjectDefinition) *_PubSubGroupDataType {
	if name == nil {
		panic("name of type PascalString for PubSubGroupDataType must not be nil")
	}
	if securityGroupId == nil {
		panic("securityGroupId of type PascalString for PubSubGroupDataType must not be nil")
	}
	_result := &_PubSubGroupDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Enabled:                           enabled,
		SecurityMode:                      securityMode,
		SecurityGroupId:                   securityGroupId,
		NoOfSecurityKeyServices:           noOfSecurityKeyServices,
		SecurityKeyServices:               securityKeyServices,
		MaxNetworkMessageSize:             maxNetworkMessageSize,
		NoOfGroupProperties:               noOfGroupProperties,
		GroupProperties:                   groupProperties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubGroupDataType) GetIdentifier() string {
	return "15611"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubGroupDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubGroupDataType) GetName() PascalString {
	return m.Name
}

func (m *_PubSubGroupDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_PubSubGroupDataType) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_PubSubGroupDataType) GetSecurityGroupId() PascalString {
	return m.SecurityGroupId
}

func (m *_PubSubGroupDataType) GetNoOfSecurityKeyServices() int32 {
	return m.NoOfSecurityKeyServices
}

func (m *_PubSubGroupDataType) GetSecurityKeyServices() []ExtensionObjectDefinition {
	return m.SecurityKeyServices
}

func (m *_PubSubGroupDataType) GetMaxNetworkMessageSize() uint32 {
	return m.MaxNetworkMessageSize
}

func (m *_PubSubGroupDataType) GetNoOfGroupProperties() int32 {
	return m.NoOfGroupProperties
}

func (m *_PubSubGroupDataType) GetGroupProperties() []ExtensionObjectDefinition {
	return m.GroupProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPubSubGroupDataType(structType any) PubSubGroupDataType {
	if casted, ok := structType.(PubSubGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubGroupDataType) GetTypeName() string {
	return "PubSubGroupDataType"
}

func (m *_PubSubGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (securityGroupId)
	lengthInBits += m.SecurityGroupId.GetLengthInBits(ctx)

	// Simple field (noOfSecurityKeyServices)
	lengthInBits += 32

	// Array field
	if len(m.SecurityKeyServices) > 0 {
		for _curItem, element := range m.SecurityKeyServices {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityKeyServices), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (maxNetworkMessageSize)
	lengthInBits += 32

	// Simple field (noOfGroupProperties)
	lengthInBits += 32

	// Array field
	if len(m.GroupProperties) > 0 {
		for _curItem, element := range m.GroupProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GroupProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PubSubGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubGroupDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__pubSubGroupDataType PubSubGroupDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	securityMode, err := ReadEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", ReadEnum(MessageSecurityModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityMode' field"))
	}
	m.SecurityMode = securityMode

	securityGroupId, err := ReadSimpleField[PascalString](ctx, "securityGroupId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroupId' field"))
	}
	m.SecurityGroupId = securityGroupId

	noOfSecurityKeyServices, err := ReadSimpleField(ctx, "noOfSecurityKeyServices", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSecurityKeyServices' field"))
	}
	m.NoOfSecurityKeyServices = noOfSecurityKeyServices

	securityKeyServices, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "securityKeyServices", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("314")), readBuffer), uint64(noOfSecurityKeyServices))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityKeyServices' field"))
	}
	m.SecurityKeyServices = securityKeyServices

	maxNetworkMessageSize, err := ReadSimpleField(ctx, "maxNetworkMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNetworkMessageSize' field"))
	}
	m.MaxNetworkMessageSize = maxNetworkMessageSize

	noOfGroupProperties, err := ReadSimpleField(ctx, "noOfGroupProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfGroupProperties' field"))
	}
	m.NoOfGroupProperties = noOfGroupProperties

	groupProperties, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "groupProperties", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("14535")), readBuffer), uint64(noOfGroupProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupProperties' field"))
	}
	m.GroupProperties = groupProperties

	if closeErr := readBuffer.CloseContext("PubSubGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubGroupDataType")
	}

	return m, nil
}

func (m *_PubSubGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubGroupDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if err := WriteSimpleEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", m.GetSecurityMode(), WriteEnum[MessageSecurityMode, uint32](MessageSecurityMode.GetValue, MessageSecurityMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'securityMode' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityGroupId", m.GetSecurityGroupId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityGroupId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfSecurityKeyServices", m.GetNoOfSecurityKeyServices(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSecurityKeyServices' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "securityKeyServices", m.GetSecurityKeyServices(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'securityKeyServices' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNetworkMessageSize", m.GetMaxNetworkMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNetworkMessageSize' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfGroupProperties", m.GetNoOfGroupProperties(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfGroupProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "groupProperties", m.GetGroupProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'groupProperties' field")
		}

		if popErr := writeBuffer.PopContext("PubSubGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubGroupDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubGroupDataType) IsPubSubGroupDataType() {}

func (m *_PubSubGroupDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PubSubGroupDataType) deepCopy() *_PubSubGroupDataType {
	if m == nil {
		return nil
	}
	_PubSubGroupDataTypeCopy := &_PubSubGroupDataType{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.Name.DeepCopy().(PascalString),
		m.Enabled,
		m.SecurityMode,
		m.SecurityGroupId.DeepCopy().(PascalString),
		m.NoOfSecurityKeyServices,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.SecurityKeyServices),
		m.MaxNetworkMessageSize,
		m.NoOfGroupProperties,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.GroupProperties),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubGroupDataTypeCopy
}

func (m *_PubSubGroupDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
