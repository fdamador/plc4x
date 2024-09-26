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

// FirmataCommandProtocolVersion is the corresponding interface of FirmataCommandProtocolVersion
type FirmataCommandProtocolVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataCommand
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// IsFirmataCommandProtocolVersion is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataCommandProtocolVersion()
}

// _FirmataCommandProtocolVersion is the data-structure of this message
type _FirmataCommandProtocolVersion struct {
	FirmataCommandContract
	MajorVersion uint8
	MinorVersion uint8
}

var _ FirmataCommandProtocolVersion = (*_FirmataCommandProtocolVersion)(nil)
var _ FirmataCommandRequirements = (*_FirmataCommandProtocolVersion)(nil)

// NewFirmataCommandProtocolVersion factory function for _FirmataCommandProtocolVersion
func NewFirmataCommandProtocolVersion(majorVersion uint8, minorVersion uint8, response bool) *_FirmataCommandProtocolVersion {
	_result := &_FirmataCommandProtocolVersion{
		FirmataCommandContract: NewFirmataCommand(response),
		MajorVersion:           majorVersion,
		MinorVersion:           minorVersion,
	}
	_result.FirmataCommandContract.(*_FirmataCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandProtocolVersion) GetCommandCode() uint8 {
	return 0x9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandProtocolVersion) GetParent() FirmataCommandContract {
	return m.FirmataCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandProtocolVersion) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_FirmataCommandProtocolVersion) GetMinorVersion() uint8 {
	return m.MinorVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataCommandProtocolVersion(structType any) FirmataCommandProtocolVersion {
	if casted, ok := structType.(FirmataCommandProtocolVersion); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandProtocolVersion); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandProtocolVersion) GetTypeName() string {
	return "FirmataCommandProtocolVersion"
}

func (m *_FirmataCommandProtocolVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataCommandContract.(*_FirmataCommand).getLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	return lengthInBits
}

func (m *_FirmataCommandProtocolVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataCommandProtocolVersion) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataCommand, response bool) (__firmataCommandProtocolVersion FirmataCommandProtocolVersion, err error) {
	m.FirmataCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandProtocolVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandProtocolVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}
	m.MajorVersion = majorVersion

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}
	m.MinorVersion = minorVersion

	if closeErr := readBuffer.CloseContext("FirmataCommandProtocolVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandProtocolVersion")
	}

	return m, nil
}

func (m *_FirmataCommandProtocolVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandProtocolVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandProtocolVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandProtocolVersion")
		}

		if err := WriteSimpleField[uint8](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint8](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandProtocolVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandProtocolVersion")
		}
		return nil
	}
	return m.FirmataCommandContract.(*_FirmataCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataCommandProtocolVersion) IsFirmataCommandProtocolVersion() {}

func (m *_FirmataCommandProtocolVersion) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataCommandProtocolVersion) deepCopy() *_FirmataCommandProtocolVersion {
	if m == nil {
		return nil
	}
	_FirmataCommandProtocolVersionCopy := &_FirmataCommandProtocolVersion{
		m.FirmataCommandContract.DeepCopy().(FirmataCommandContract),
		m.MajorVersion,
		m.MinorVersion,
	}
	m.FirmataCommandContract.(*_FirmataCommand)._SubType = m
	return _FirmataCommandProtocolVersionCopy
}

func (m *_FirmataCommandProtocolVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
