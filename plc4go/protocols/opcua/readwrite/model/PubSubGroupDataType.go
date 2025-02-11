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
	// GetSecurityKeyServices returns SecurityKeyServices (property field)
	GetSecurityKeyServices() []EndpointDescription
	// GetMaxNetworkMessageSize returns MaxNetworkMessageSize (property field)
	GetMaxNetworkMessageSize() uint32
	// GetGroupProperties returns GroupProperties (property field)
	GetGroupProperties() []KeyValuePair
	// IsPubSubGroupDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubGroupDataType()
	// CreateBuilder creates a PubSubGroupDataTypeBuilder
	CreatePubSubGroupDataTypeBuilder() PubSubGroupDataTypeBuilder
}

// _PubSubGroupDataType is the data-structure of this message
type _PubSubGroupDataType struct {
	ExtensionObjectDefinitionContract
	Name                  PascalString
	Enabled               bool
	SecurityMode          MessageSecurityMode
	SecurityGroupId       PascalString
	SecurityKeyServices   []EndpointDescription
	MaxNetworkMessageSize uint32
	GroupProperties       []KeyValuePair
	// Reserved Fields
	reservedField0 *uint8
}

var _ PubSubGroupDataType = (*_PubSubGroupDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubGroupDataType)(nil)

// NewPubSubGroupDataType factory function for _PubSubGroupDataType
func NewPubSubGroupDataType(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair) *_PubSubGroupDataType {
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
		SecurityKeyServices:               securityKeyServices,
		MaxNetworkMessageSize:             maxNetworkMessageSize,
		GroupProperties:                   groupProperties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PubSubGroupDataTypeBuilder is a builder for PubSubGroupDataType
type PubSubGroupDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair) PubSubGroupDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) PubSubGroupDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) PubSubGroupDataTypeBuilder
	// WithEnabled adds Enabled (property field)
	WithEnabled(bool) PubSubGroupDataTypeBuilder
	// WithSecurityMode adds SecurityMode (property field)
	WithSecurityMode(MessageSecurityMode) PubSubGroupDataTypeBuilder
	// WithSecurityGroupId adds SecurityGroupId (property field)
	WithSecurityGroupId(PascalString) PubSubGroupDataTypeBuilder
	// WithSecurityGroupIdBuilder adds SecurityGroupId (property field) which is build by the builder
	WithSecurityGroupIdBuilder(func(PascalStringBuilder) PascalStringBuilder) PubSubGroupDataTypeBuilder
	// WithSecurityKeyServices adds SecurityKeyServices (property field)
	WithSecurityKeyServices(...EndpointDescription) PubSubGroupDataTypeBuilder
	// WithMaxNetworkMessageSize adds MaxNetworkMessageSize (property field)
	WithMaxNetworkMessageSize(uint32) PubSubGroupDataTypeBuilder
	// WithGroupProperties adds GroupProperties (property field)
	WithGroupProperties(...KeyValuePair) PubSubGroupDataTypeBuilder
	// Build builds the PubSubGroupDataType or returns an error if something is wrong
	Build() (PubSubGroupDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PubSubGroupDataType
}

// NewPubSubGroupDataTypeBuilder() creates a PubSubGroupDataTypeBuilder
func NewPubSubGroupDataTypeBuilder() PubSubGroupDataTypeBuilder {
	return &_PubSubGroupDataTypeBuilder{_PubSubGroupDataType: new(_PubSubGroupDataType)}
}

type _PubSubGroupDataTypeBuilder struct {
	*_PubSubGroupDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PubSubGroupDataTypeBuilder) = (*_PubSubGroupDataTypeBuilder)(nil)

func (b *_PubSubGroupDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PubSubGroupDataTypeBuilder) WithMandatoryFields(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair) PubSubGroupDataTypeBuilder {
	return b.WithName(name).WithEnabled(enabled).WithSecurityMode(securityMode).WithSecurityGroupId(securityGroupId).WithSecurityKeyServices(securityKeyServices...).WithMaxNetworkMessageSize(maxNetworkMessageSize).WithGroupProperties(groupProperties...)
}

func (b *_PubSubGroupDataTypeBuilder) WithName(name PascalString) PubSubGroupDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) PubSubGroupDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithEnabled(enabled bool) PubSubGroupDataTypeBuilder {
	b.Enabled = enabled
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithSecurityMode(securityMode MessageSecurityMode) PubSubGroupDataTypeBuilder {
	b.SecurityMode = securityMode
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithSecurityGroupId(securityGroupId PascalString) PubSubGroupDataTypeBuilder {
	b.SecurityGroupId = securityGroupId
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithSecurityGroupIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) PubSubGroupDataTypeBuilder {
	builder := builderSupplier(b.SecurityGroupId.CreatePascalStringBuilder())
	var err error
	b.SecurityGroupId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithSecurityKeyServices(securityKeyServices ...EndpointDescription) PubSubGroupDataTypeBuilder {
	b.SecurityKeyServices = securityKeyServices
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithMaxNetworkMessageSize(maxNetworkMessageSize uint32) PubSubGroupDataTypeBuilder {
	b.MaxNetworkMessageSize = maxNetworkMessageSize
	return b
}

func (b *_PubSubGroupDataTypeBuilder) WithGroupProperties(groupProperties ...KeyValuePair) PubSubGroupDataTypeBuilder {
	b.GroupProperties = groupProperties
	return b
}

func (b *_PubSubGroupDataTypeBuilder) Build() (PubSubGroupDataType, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.SecurityGroupId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityGroupId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PubSubGroupDataType.deepCopy(), nil
}

func (b *_PubSubGroupDataTypeBuilder) MustBuild() PubSubGroupDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_PubSubGroupDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_PubSubGroupDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PubSubGroupDataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePubSubGroupDataTypeBuilder().(*_PubSubGroupDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePubSubGroupDataTypeBuilder creates a PubSubGroupDataTypeBuilder
func (b *_PubSubGroupDataType) CreatePubSubGroupDataTypeBuilder() PubSubGroupDataTypeBuilder {
	if b == nil {
		return NewPubSubGroupDataTypeBuilder()
	}
	return &_PubSubGroupDataTypeBuilder{_PubSubGroupDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubGroupDataType) GetExtensionId() int32 {
	return int32(15611)
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

func (m *_PubSubGroupDataType) GetSecurityKeyServices() []EndpointDescription {
	return m.SecurityKeyServices
}

func (m *_PubSubGroupDataType) GetMaxNetworkMessageSize() uint32 {
	return m.MaxNetworkMessageSize
}

func (m *_PubSubGroupDataType) GetGroupProperties() []KeyValuePair {
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

	// Implicit Field (noOfSecurityKeyServices)
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

	// Implicit Field (noOfGroupProperties)
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

func (m *_PubSubGroupDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__pubSubGroupDataType PubSubGroupDataType, err error) {
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

	noOfSecurityKeyServices, err := ReadImplicitField[int32](ctx, "noOfSecurityKeyServices", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSecurityKeyServices' field"))
	}
	_ = noOfSecurityKeyServices

	securityKeyServices, err := ReadCountArrayField[EndpointDescription](ctx, "securityKeyServices", ReadComplex[EndpointDescription](ExtensionObjectDefinitionParseWithBufferProducer[EndpointDescription]((int32)(int32(314))), readBuffer), uint64(noOfSecurityKeyServices))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityKeyServices' field"))
	}
	m.SecurityKeyServices = securityKeyServices

	maxNetworkMessageSize, err := ReadSimpleField(ctx, "maxNetworkMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNetworkMessageSize' field"))
	}
	m.MaxNetworkMessageSize = maxNetworkMessageSize

	noOfGroupProperties, err := ReadImplicitField[int32](ctx, "noOfGroupProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfGroupProperties' field"))
	}
	_ = noOfGroupProperties

	groupProperties, err := ReadCountArrayField[KeyValuePair](ctx, "groupProperties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfGroupProperties))
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
		noOfSecurityKeyServices := int32(utils.InlineIf(bool((m.GetSecurityKeyServices()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSecurityKeyServices()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSecurityKeyServices", noOfSecurityKeyServices, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSecurityKeyServices' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "securityKeyServices", m.GetSecurityKeyServices(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'securityKeyServices' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNetworkMessageSize", m.GetMaxNetworkMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNetworkMessageSize' field")
		}
		noOfGroupProperties := int32(utils.InlineIf(bool((m.GetGroupProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetGroupProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfGroupProperties", noOfGroupProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
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
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Name.DeepCopy().(PascalString),
		m.Enabled,
		m.SecurityMode,
		m.SecurityGroupId.DeepCopy().(PascalString),
		utils.DeepCopySlice[EndpointDescription, EndpointDescription](m.SecurityKeyServices),
		m.MaxNetworkMessageSize,
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.GroupProperties),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubGroupDataTypeCopy
}

func (m *_PubSubGroupDataType) String() string {
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
