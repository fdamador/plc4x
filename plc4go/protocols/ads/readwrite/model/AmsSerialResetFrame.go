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

// AmsSerialResetFrame is the corresponding interface of AmsSerialResetFrame
type AmsSerialResetFrame interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetMagicCookie returns MagicCookie (property field)
	GetMagicCookie() uint16
	// GetTransmitterAddress returns TransmitterAddress (property field)
	GetTransmitterAddress() int8
	// GetReceiverAddress returns ReceiverAddress (property field)
	GetReceiverAddress() int8
	// GetFragmentNumber returns FragmentNumber (property field)
	GetFragmentNumber() int8
	// GetLength returns Length (property field)
	GetLength() int8
	// GetCrc returns Crc (property field)
	GetCrc() uint16
	// IsAmsSerialResetFrame is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAmsSerialResetFrame()
}

// _AmsSerialResetFrame is the data-structure of this message
type _AmsSerialResetFrame struct {
	MagicCookie        uint16
	TransmitterAddress int8
	ReceiverAddress    int8
	FragmentNumber     int8
	Length             int8
	Crc                uint16
}

var _ AmsSerialResetFrame = (*_AmsSerialResetFrame)(nil)

// NewAmsSerialResetFrame factory function for _AmsSerialResetFrame
func NewAmsSerialResetFrame(magicCookie uint16, transmitterAddress int8, receiverAddress int8, fragmentNumber int8, length int8, crc uint16) *_AmsSerialResetFrame {
	return &_AmsSerialResetFrame{MagicCookie: magicCookie, TransmitterAddress: transmitterAddress, ReceiverAddress: receiverAddress, FragmentNumber: fragmentNumber, Length: length, Crc: crc}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsSerialResetFrame) GetMagicCookie() uint16 {
	return m.MagicCookie
}

func (m *_AmsSerialResetFrame) GetTransmitterAddress() int8 {
	return m.TransmitterAddress
}

func (m *_AmsSerialResetFrame) GetReceiverAddress() int8 {
	return m.ReceiverAddress
}

func (m *_AmsSerialResetFrame) GetFragmentNumber() int8 {
	return m.FragmentNumber
}

func (m *_AmsSerialResetFrame) GetLength() int8 {
	return m.Length
}

func (m *_AmsSerialResetFrame) GetCrc() uint16 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAmsSerialResetFrame(structType any) AmsSerialResetFrame {
	if casted, ok := structType.(AmsSerialResetFrame); ok {
		return casted
	}
	if casted, ok := structType.(*AmsSerialResetFrame); ok {
		return *casted
	}
	return nil
}

func (m *_AmsSerialResetFrame) GetTypeName() string {
	return "AmsSerialResetFrame"
}

func (m *_AmsSerialResetFrame) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16

	// Simple field (transmitterAddress)
	lengthInBits += 8

	// Simple field (receiverAddress)
	lengthInBits += 8

	// Simple field (fragmentNumber)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AmsSerialResetFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsSerialResetFrameParse(ctx context.Context, theBytes []byte) (AmsSerialResetFrame, error) {
	return AmsSerialResetFrameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AmsSerialResetFrameParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsSerialResetFrame, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsSerialResetFrame, error) {
		return AmsSerialResetFrameParseWithBuffer(ctx, readBuffer)
	}
}

func AmsSerialResetFrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsSerialResetFrame, error) {
	v, err := (&_AmsSerialResetFrame{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AmsSerialResetFrame) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__amsSerialResetFrame AmsSerialResetFrame, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsSerialResetFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsSerialResetFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	magicCookie, err := ReadSimpleField(ctx, "magicCookie", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'magicCookie' field"))
	}
	m.MagicCookie = magicCookie

	transmitterAddress, err := ReadSimpleField(ctx, "transmitterAddress", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmitterAddress' field"))
	}
	m.TransmitterAddress = transmitterAddress

	receiverAddress, err := ReadSimpleField(ctx, "receiverAddress", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiverAddress' field"))
	}
	m.ReceiverAddress = receiverAddress

	fragmentNumber, err := ReadSimpleField(ctx, "fragmentNumber", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fragmentNumber' field"))
	}
	m.FragmentNumber = fragmentNumber

	length, err := ReadSimpleField(ctx, "length", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	m.Length = length

	crc, err := ReadSimpleField(ctx, "crc", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	m.Crc = crc

	if closeErr := readBuffer.CloseContext("AmsSerialResetFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsSerialResetFrame")
	}

	return m, nil
}

func (m *_AmsSerialResetFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsSerialResetFrame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsSerialResetFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsSerialResetFrame")
	}

	if err := WriteSimpleField[uint16](ctx, "magicCookie", m.GetMagicCookie(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'magicCookie' field")
	}

	if err := WriteSimpleField[int8](ctx, "transmitterAddress", m.GetTransmitterAddress(), WriteSignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'transmitterAddress' field")
	}

	if err := WriteSimpleField[int8](ctx, "receiverAddress", m.GetReceiverAddress(), WriteSignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'receiverAddress' field")
	}

	if err := WriteSimpleField[int8](ctx, "fragmentNumber", m.GetFragmentNumber(), WriteSignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'fragmentNumber' field")
	}

	if err := WriteSimpleField[int8](ctx, "length", m.GetLength(), WriteSignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'length' field")
	}

	if err := WriteSimpleField[uint16](ctx, "crc", m.GetCrc(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'crc' field")
	}

	if popErr := writeBuffer.PopContext("AmsSerialResetFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsSerialResetFrame")
	}
	return nil
}

func (m *_AmsSerialResetFrame) IsAmsSerialResetFrame() {}

func (m *_AmsSerialResetFrame) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AmsSerialResetFrame) deepCopy() *_AmsSerialResetFrame {
	if m == nil {
		return nil
	}
	_AmsSerialResetFrameCopy := &_AmsSerialResetFrame{
		m.MagicCookie,
		m.TransmitterAddress,
		m.ReceiverAddress,
		m.FragmentNumber,
		m.Length,
		m.Crc,
	}
	return _AmsSerialResetFrameCopy
}

func (m *_AmsSerialResetFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
