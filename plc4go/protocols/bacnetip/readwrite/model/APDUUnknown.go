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

// APDUUnknown is the corresponding interface of APDUUnknown
type APDUUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetUnknownTypeRest returns UnknownTypeRest (property field)
	GetUnknownTypeRest() uint8
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
	// IsAPDUUnknown is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUUnknown()
}

// _APDUUnknown is the data-structure of this message
type _APDUUnknown struct {
	APDUContract
	UnknownTypeRest uint8
	UnknownBytes    []byte
}

var _ APDUUnknown = (*_APDUUnknown)(nil)
var _ APDURequirements = (*_APDUUnknown)(nil)

// NewAPDUUnknown factory function for _APDUUnknown
func NewAPDUUnknown(unknownTypeRest uint8, unknownBytes []byte, apduLength uint16) *_APDUUnknown {
	_result := &_APDUUnknown{
		APDUContract:    NewAPDU(apduLength),
		UnknownTypeRest: unknownTypeRest,
		UnknownBytes:    unknownBytes,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUUnknown) GetApduType() ApduType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUUnknown) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUUnknown) GetUnknownTypeRest() uint8 {
	return m.UnknownTypeRest
}

func (m *_APDUUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUUnknown(structType any) APDUUnknown {
	if casted, ok := structType.(APDUUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*APDUUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_APDUUnknown) GetTypeName() string {
	return "APDUUnknown"
}

func (m *_APDUUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Simple field (unknownTypeRest)
	lengthInBits += 4

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_APDUUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUUnknown) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUUnknown APDUUnknown, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unknownTypeRest, err := ReadSimpleField(ctx, "unknownTypeRest", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownTypeRest' field"))
	}
	m.UnknownTypeRest = unknownTypeRest

	unknownBytes, err := readBuffer.ReadByteArray("unknownBytes", int(utils.InlineIf((bool((apduLength) > (0))), func() any { return int32(apduLength) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownBytes' field"))
	}
	m.UnknownBytes = unknownBytes

	if closeErr := readBuffer.CloseContext("APDUUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUUnknown")
	}

	return m, nil
}

func (m *_APDUUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUUnknown")
		}

		if err := WriteSimpleField[uint8](ctx, "unknownTypeRest", m.GetUnknownTypeRest(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownTypeRest' field")
		}

		if err := WriteByteArrayField(ctx, "unknownBytes", m.GetUnknownBytes(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUUnknown")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUUnknown) IsAPDUUnknown() {}

func (m *_APDUUnknown) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUUnknown) deepCopy() *_APDUUnknown {
	if m == nil {
		return nil
	}
	_APDUUnknownCopy := &_APDUUnknown{
		m.APDUContract.(*_APDU).deepCopy(),
		m.UnknownTypeRest,
		utils.DeepCopySlice[byte, byte](m.UnknownBytes),
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUUnknownCopy
}

func (m *_APDUUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
