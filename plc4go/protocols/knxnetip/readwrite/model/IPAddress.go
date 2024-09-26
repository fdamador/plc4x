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

// IPAddress is the corresponding interface of IPAddress
type IPAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAddr returns Addr (property field)
	GetAddr() []byte
	// IsIPAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIPAddress()
}

// _IPAddress is the data-structure of this message
type _IPAddress struct {
	Addr []byte
}

var _ IPAddress = (*_IPAddress)(nil)

// NewIPAddress factory function for _IPAddress
func NewIPAddress(addr []byte) *_IPAddress {
	return &_IPAddress{Addr: addr}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IPAddress) GetAddr() []byte {
	return m.Addr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIPAddress(structType any) IPAddress {
	if casted, ok := structType.(IPAddress); ok {
		return casted
	}
	if casted, ok := structType.(*IPAddress); ok {
		return *casted
	}
	return nil
}

func (m *_IPAddress) GetTypeName() string {
	return "IPAddress"
}

func (m *_IPAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *_IPAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IPAddressParse(ctx context.Context, theBytes []byte) (IPAddress, error) {
	return IPAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IPAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (IPAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (IPAddress, error) {
		return IPAddressParseWithBuffer(ctx, readBuffer)
	}
}

func IPAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IPAddress, error) {
	v, err := (&_IPAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_IPAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__iPAddress IPAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IPAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IPAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	addr, err := readBuffer.ReadByteArray("addr", int(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addr' field"))
	}
	m.Addr = addr

	if closeErr := readBuffer.CloseContext("IPAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IPAddress")
	}

	return m, nil
}

func (m *_IPAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IPAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("IPAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IPAddress")
	}

	if err := WriteByteArrayField(ctx, "addr", m.GetAddr(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'addr' field")
	}

	if popErr := writeBuffer.PopContext("IPAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IPAddress")
	}
	return nil
}

func (m *_IPAddress) IsIPAddress() {}

func (m *_IPAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IPAddress) deepCopy() *_IPAddress {
	if m == nil {
		return nil
	}
	_IPAddressCopy := &_IPAddress{
		utils.DeepCopySlice[byte, byte](m.Addr),
	}
	return _IPAddressCopy
}

func (m *_IPAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
