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

// ModificationInfo is the corresponding interface of ModificationInfo
type ModificationInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetModificationTime returns ModificationTime (property field)
	GetModificationTime() int64
	// GetUpdateType returns UpdateType (property field)
	GetUpdateType() HistoryUpdateType
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// IsModificationInfo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModificationInfo()
}

// _ModificationInfo is the data-structure of this message
type _ModificationInfo struct {
	ExtensionObjectDefinitionContract
	ModificationTime int64
	UpdateType       HistoryUpdateType
	UserName         PascalString
}

var _ ModificationInfo = (*_ModificationInfo)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ModificationInfo)(nil)

// NewModificationInfo factory function for _ModificationInfo
func NewModificationInfo(modificationTime int64, updateType HistoryUpdateType, userName PascalString) *_ModificationInfo {
	if userName == nil {
		panic("userName of type PascalString for ModificationInfo must not be nil")
	}
	_result := &_ModificationInfo{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ModificationTime:                  modificationTime,
		UpdateType:                        updateType,
		UserName:                          userName,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModificationInfo) GetIdentifier() string {
	return "11218"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModificationInfo) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModificationInfo) GetModificationTime() int64 {
	return m.ModificationTime
}

func (m *_ModificationInfo) GetUpdateType() HistoryUpdateType {
	return m.UpdateType
}

func (m *_ModificationInfo) GetUserName() PascalString {
	return m.UserName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModificationInfo(structType any) ModificationInfo {
	if casted, ok := structType.(ModificationInfo); ok {
		return casted
	}
	if casted, ok := structType.(*ModificationInfo); ok {
		return *casted
	}
	return nil
}

func (m *_ModificationInfo) GetTypeName() string {
	return "ModificationInfo"
}

func (m *_ModificationInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (modificationTime)
	lengthInBits += 64

	// Simple field (updateType)
	lengthInBits += 32

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ModificationInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModificationInfo) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__modificationInfo ModificationInfo, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModificationInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModificationInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	modificationTime, err := ReadSimpleField(ctx, "modificationTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modificationTime' field"))
	}
	m.ModificationTime = modificationTime

	updateType, err := ReadEnumField[HistoryUpdateType](ctx, "updateType", "HistoryUpdateType", ReadEnum(HistoryUpdateTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateType' field"))
	}
	m.UpdateType = updateType

	userName, err := ReadSimpleField[PascalString](ctx, "userName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	if closeErr := readBuffer.CloseContext("ModificationInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModificationInfo")
	}

	return m, nil
}

func (m *_ModificationInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModificationInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModificationInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModificationInfo")
		}

		if err := WriteSimpleField[int64](ctx, "modificationTime", m.GetModificationTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'modificationTime' field")
		}

		if err := WriteSimpleEnumField[HistoryUpdateType](ctx, "updateType", "HistoryUpdateType", m.GetUpdateType(), WriteEnum[HistoryUpdateType, uint32](HistoryUpdateType.GetValue, HistoryUpdateType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'updateType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "userName", m.GetUserName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if popErr := writeBuffer.PopContext("ModificationInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModificationInfo")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModificationInfo) IsModificationInfo() {}

func (m *_ModificationInfo) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModificationInfo) deepCopy() *_ModificationInfo {
	if m == nil {
		return nil
	}
	_ModificationInfoCopy := &_ModificationInfo{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.ModificationTime,
		m.UpdateType,
		m.UserName.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ModificationInfoCopy
}

func (m *_ModificationInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
