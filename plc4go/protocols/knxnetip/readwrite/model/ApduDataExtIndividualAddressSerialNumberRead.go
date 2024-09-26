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

// ApduDataExtIndividualAddressSerialNumberRead is the corresponding interface of ApduDataExtIndividualAddressSerialNumberRead
type ApduDataExtIndividualAddressSerialNumberRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtIndividualAddressSerialNumberRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtIndividualAddressSerialNumberRead()
}

// _ApduDataExtIndividualAddressSerialNumberRead is the data-structure of this message
type _ApduDataExtIndividualAddressSerialNumberRead struct {
	ApduDataExtContract
}

var _ ApduDataExtIndividualAddressSerialNumberRead = (*_ApduDataExtIndividualAddressSerialNumberRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtIndividualAddressSerialNumberRead)(nil)

// NewApduDataExtIndividualAddressSerialNumberRead factory function for _ApduDataExtIndividualAddressSerialNumberRead
func NewApduDataExtIndividualAddressSerialNumberRead(length uint8) *_ApduDataExtIndividualAddressSerialNumberRead {
	_result := &_ApduDataExtIndividualAddressSerialNumberRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetExtApciType() uint8 {
	return 0x1C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtIndividualAddressSerialNumberRead(structType any) ApduDataExtIndividualAddressSerialNumberRead {
	if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberRead"
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtIndividualAddressSerialNumberRead ApduDataExtIndividualAddressSerialNumberRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberRead")
	}

	return m, nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) IsApduDataExtIndividualAddressSerialNumberRead() {
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) deepCopy() *_ApduDataExtIndividualAddressSerialNumberRead {
	if m == nil {
		return nil
	}
	_ApduDataExtIndividualAddressSerialNumberReadCopy := &_ApduDataExtIndividualAddressSerialNumberRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtIndividualAddressSerialNumberReadCopy
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
