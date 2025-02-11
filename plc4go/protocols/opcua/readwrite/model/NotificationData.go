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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NotificationData is the corresponding interface of NotificationData
type NotificationData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsNotificationData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNotificationData()
	// CreateBuilder creates a NotificationDataBuilder
	CreateNotificationDataBuilder() NotificationDataBuilder
}

// _NotificationData is the data-structure of this message
type _NotificationData struct {
	ExtensionObjectDefinitionContract
}

var _ NotificationData = (*_NotificationData)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NotificationData)(nil)

// NewNotificationData factory function for _NotificationData
func NewNotificationData() *_NotificationData {
	_result := &_NotificationData{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NotificationDataBuilder is a builder for NotificationData
type NotificationDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NotificationDataBuilder
	// Build builds the NotificationData or returns an error if something is wrong
	Build() (NotificationData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NotificationData
}

// NewNotificationDataBuilder() creates a NotificationDataBuilder
func NewNotificationDataBuilder() NotificationDataBuilder {
	return &_NotificationDataBuilder{_NotificationData: new(_NotificationData)}
}

type _NotificationDataBuilder struct {
	*_NotificationData

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (NotificationDataBuilder) = (*_NotificationDataBuilder)(nil)

func (b *_NotificationDataBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_NotificationDataBuilder) WithMandatoryFields() NotificationDataBuilder {
	return b
}

func (b *_NotificationDataBuilder) Build() (NotificationData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NotificationData.deepCopy(), nil
}

func (b *_NotificationDataBuilder) MustBuild() NotificationData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NotificationDataBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NotificationDataBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_NotificationDataBuilder) DeepCopy() any {
	_copy := b.CreateNotificationDataBuilder().(*_NotificationDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNotificationDataBuilder creates a NotificationDataBuilder
func (b *_NotificationData) CreateNotificationDataBuilder() NotificationDataBuilder {
	if b == nil {
		return NewNotificationDataBuilder()
	}
	return &_NotificationDataBuilder{_NotificationData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NotificationData) GetExtensionId() int32 {
	return int32(947)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NotificationData) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastNotificationData(structType any) NotificationData {
	if casted, ok := structType.(NotificationData); ok {
		return casted
	}
	if casted, ok := structType.(*NotificationData); ok {
		return *casted
	}
	return nil
}

func (m *_NotificationData) GetTypeName() string {
	return "NotificationData"
}

func (m *_NotificationData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NotificationData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NotificationData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__notificationData NotificationData, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NotificationData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NotificationData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotificationData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NotificationData")
	}

	return m, nil
}

func (m *_NotificationData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NotificationData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotificationData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NotificationData")
		}

		if popErr := writeBuffer.PopContext("NotificationData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NotificationData")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NotificationData) IsNotificationData() {}

func (m *_NotificationData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NotificationData) deepCopy() *_NotificationData {
	if m == nil {
		return nil
	}
	_NotificationDataCopy := &_NotificationData{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NotificationDataCopy
}

func (m *_NotificationData) String() string {
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
