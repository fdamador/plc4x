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

// ApduDataExtReadRouterStatusRequest is the corresponding interface of ApduDataExtReadRouterStatusRequest
type ApduDataExtReadRouterStatusRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtReadRouterStatusRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtReadRouterStatusRequest()
}

// _ApduDataExtReadRouterStatusRequest is the data-structure of this message
type _ApduDataExtReadRouterStatusRequest struct {
	ApduDataExtContract
}

var _ ApduDataExtReadRouterStatusRequest = (*_ApduDataExtReadRouterStatusRequest)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtReadRouterStatusRequest)(nil)

// NewApduDataExtReadRouterStatusRequest factory function for _ApduDataExtReadRouterStatusRequest
func NewApduDataExtReadRouterStatusRequest(length uint8) *_ApduDataExtReadRouterStatusRequest {
	_result := &_ApduDataExtReadRouterStatusRequest{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterStatusRequest) GetExtApciType() uint8 {
	return 0x0D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterStatusRequest) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterStatusRequest(structType any) ApduDataExtReadRouterStatusRequest {
	if casted, ok := structType.(ApduDataExtReadRouterStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterStatusRequest) GetTypeName() string {
	return "ApduDataExtReadRouterStatusRequest"
}

func (m *_ApduDataExtReadRouterStatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterStatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtReadRouterStatusRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtReadRouterStatusRequest ApduDataExtReadRouterStatusRequest, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterStatusRequest")
	}

	return m, nil
}

func (m *_ApduDataExtReadRouterStatusRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterStatusRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterStatusRequest")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterStatusRequest) IsApduDataExtReadRouterStatusRequest() {}

func (m *_ApduDataExtReadRouterStatusRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtReadRouterStatusRequest) deepCopy() *_ApduDataExtReadRouterStatusRequest {
	if m == nil {
		return nil
	}
	_ApduDataExtReadRouterStatusRequestCopy := &_ApduDataExtReadRouterStatusRequest{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtReadRouterStatusRequestCopy
}

func (m *_ApduDataExtReadRouterStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
