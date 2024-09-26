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

// ApduDataExtLinkRead is the corresponding interface of ApduDataExtLinkRead
type ApduDataExtLinkRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtLinkRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtLinkRead()
}

// _ApduDataExtLinkRead is the data-structure of this message
type _ApduDataExtLinkRead struct {
	ApduDataExtContract
}

var _ ApduDataExtLinkRead = (*_ApduDataExtLinkRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtLinkRead)(nil)

// NewApduDataExtLinkRead factory function for _ApduDataExtLinkRead
func NewApduDataExtLinkRead(length uint8) *_ApduDataExtLinkRead {
	_result := &_ApduDataExtLinkRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtLinkRead) GetExtApciType() uint8 {
	return 0x25
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtLinkRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtLinkRead(structType any) ApduDataExtLinkRead {
	if casted, ok := structType.(ApduDataExtLinkRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtLinkRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtLinkRead) GetTypeName() string {
	return "ApduDataExtLinkRead"
}

func (m *_ApduDataExtLinkRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtLinkRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtLinkRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtLinkRead ApduDataExtLinkRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtLinkRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkRead")
	}

	return m, nil
}

func (m *_ApduDataExtLinkRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtLinkRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtLinkRead) IsApduDataExtLinkRead() {}

func (m *_ApduDataExtLinkRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtLinkRead) deepCopy() *_ApduDataExtLinkRead {
	if m == nil {
		return nil
	}
	_ApduDataExtLinkReadCopy := &_ApduDataExtLinkRead{
		m.ApduDataExtContract.DeepCopy().(ApduDataExtContract),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtLinkReadCopy
}

func (m *_ApduDataExtLinkRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
