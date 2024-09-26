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

// ApduDataExtAuthorizeResponse is the corresponding interface of ApduDataExtAuthorizeResponse
type ApduDataExtAuthorizeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// GetLevel returns Level (property field)
	GetLevel() uint8
	// IsApduDataExtAuthorizeResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtAuthorizeResponse()
}

// _ApduDataExtAuthorizeResponse is the data-structure of this message
type _ApduDataExtAuthorizeResponse struct {
	ApduDataExtContract
	Level uint8
}

var _ ApduDataExtAuthorizeResponse = (*_ApduDataExtAuthorizeResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtAuthorizeResponse)(nil)

// NewApduDataExtAuthorizeResponse factory function for _ApduDataExtAuthorizeResponse
func NewApduDataExtAuthorizeResponse(level uint8, length uint8) *_ApduDataExtAuthorizeResponse {
	_result := &_ApduDataExtAuthorizeResponse{
		ApduDataExtContract: NewApduDataExt(length),
		Level:               level,
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtAuthorizeResponse) GetExtApciType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtAuthorizeResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtAuthorizeResponse) GetLevel() uint8 {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataExtAuthorizeResponse(structType any) ApduDataExtAuthorizeResponse {
	if casted, ok := structType.(ApduDataExtAuthorizeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtAuthorizeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtAuthorizeResponse) GetTypeName() string {
	return "ApduDataExtAuthorizeResponse"
}

func (m *_ApduDataExtAuthorizeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ApduDataExtAuthorizeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtAuthorizeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtAuthorizeResponse ApduDataExtAuthorizeResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtAuthorizeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	level, err := ReadSimpleField(ctx, "level", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtAuthorizeResponse")
	}

	return m, nil
}

func (m *_ApduDataExtAuthorizeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtAuthorizeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtAuthorizeResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "level", m.GetLevel(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtAuthorizeResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtAuthorizeResponse) IsApduDataExtAuthorizeResponse() {}

func (m *_ApduDataExtAuthorizeResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtAuthorizeResponse) deepCopy() *_ApduDataExtAuthorizeResponse {
	if m == nil {
		return nil
	}
	_ApduDataExtAuthorizeResponseCopy := &_ApduDataExtAuthorizeResponse{
		m.ApduDataExtContract.DeepCopy().(ApduDataExtContract),
		m.Level,
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtAuthorizeResponseCopy
}

func (m *_ApduDataExtAuthorizeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
