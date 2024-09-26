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

// ApduDataExtReadRouterMemoryRequest is the corresponding interface of ApduDataExtReadRouterMemoryRequest
type ApduDataExtReadRouterMemoryRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtReadRouterMemoryRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtReadRouterMemoryRequest()
}

// _ApduDataExtReadRouterMemoryRequest is the data-structure of this message
type _ApduDataExtReadRouterMemoryRequest struct {
	ApduDataExtContract
}

var _ ApduDataExtReadRouterMemoryRequest = (*_ApduDataExtReadRouterMemoryRequest)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtReadRouterMemoryRequest)(nil)

// NewApduDataExtReadRouterMemoryRequest factory function for _ApduDataExtReadRouterMemoryRequest
func NewApduDataExtReadRouterMemoryRequest(length uint8) *_ApduDataExtReadRouterMemoryRequest {
	_result := &_ApduDataExtReadRouterMemoryRequest{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterMemoryRequest) GetExtApciType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterMemoryRequest) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterMemoryRequest(structType any) ApduDataExtReadRouterMemoryRequest {
	if casted, ok := structType.(ApduDataExtReadRouterMemoryRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterMemoryRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetTypeName() string {
	return "ApduDataExtReadRouterMemoryRequest"
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtReadRouterMemoryRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtReadRouterMemoryRequest ApduDataExtReadRouterMemoryRequest, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterMemoryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterMemoryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterMemoryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterMemoryRequest")
	}

	return m, nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterMemoryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterMemoryRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterMemoryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterMemoryRequest")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterMemoryRequest) IsApduDataExtReadRouterMemoryRequest() {}

func (m *_ApduDataExtReadRouterMemoryRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtReadRouterMemoryRequest) deepCopy() *_ApduDataExtReadRouterMemoryRequest {
	if m == nil {
		return nil
	}
	_ApduDataExtReadRouterMemoryRequestCopy := &_ApduDataExtReadRouterMemoryRequest{
		m.ApduDataExtContract.DeepCopy().(ApduDataExtContract),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtReadRouterMemoryRequestCopy
}

func (m *_ApduDataExtReadRouterMemoryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
