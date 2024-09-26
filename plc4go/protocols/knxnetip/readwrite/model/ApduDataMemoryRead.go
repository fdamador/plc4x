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

// ApduDataMemoryRead is the corresponding interface of ApduDataMemoryRead
type ApduDataMemoryRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// GetNumBytes returns NumBytes (property field)
	GetNumBytes() uint8
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// IsApduDataMemoryRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataMemoryRead()
}

// _ApduDataMemoryRead is the data-structure of this message
type _ApduDataMemoryRead struct {
	ApduDataContract
	NumBytes uint8
	Address  uint16
}

var _ ApduDataMemoryRead = (*_ApduDataMemoryRead)(nil)
var _ ApduDataRequirements = (*_ApduDataMemoryRead)(nil)

// NewApduDataMemoryRead factory function for _ApduDataMemoryRead
func NewApduDataMemoryRead(numBytes uint8, address uint16, dataLength uint8) *_ApduDataMemoryRead {
	_result := &_ApduDataMemoryRead{
		ApduDataContract: NewApduData(dataLength),
		NumBytes:         numBytes,
		Address:          address,
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryRead) GetApciType() uint8 {
	return 0x8
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryRead) GetParent() ApduDataContract {
	return m.ApduDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataMemoryRead) GetNumBytes() uint8 {
	return m.NumBytes
}

func (m *_ApduDataMemoryRead) GetAddress() uint16 {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataMemoryRead(structType any) ApduDataMemoryRead {
	if casted, ok := structType.(ApduDataMemoryRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryRead) GetTypeName() string {
	return "ApduDataMemoryRead"
}

func (m *_ApduDataMemoryRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	// Simple field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ApduDataMemoryRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataMemoryRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataMemoryRead ApduDataMemoryRead, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numBytes, err := ReadSimpleField(ctx, "numBytes", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numBytes' field"))
	}
	m.NumBytes = numBytes

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("ApduDataMemoryRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryRead")
	}

	return m, nil
}

func (m *_ApduDataMemoryRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataMemoryRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryRead")
		}

		if err := WriteSimpleField[uint8](ctx, "numBytes", m.GetNumBytes(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'numBytes' field")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryRead")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataMemoryRead) IsApduDataMemoryRead() {}

func (m *_ApduDataMemoryRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataMemoryRead) deepCopy() *_ApduDataMemoryRead {
	if m == nil {
		return nil
	}
	_ApduDataMemoryReadCopy := &_ApduDataMemoryRead{
		m.ApduDataContract.(*_ApduData).deepCopy(),
		m.NumBytes,
		m.Address,
	}
	m.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataMemoryReadCopy
}

func (m *_ApduDataMemoryRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
