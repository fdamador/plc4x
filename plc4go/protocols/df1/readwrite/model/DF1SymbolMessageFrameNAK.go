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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1SymbolMessageFrameNAK is the corresponding interface of DF1SymbolMessageFrameNAK
type DF1SymbolMessageFrameNAK interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1Symbol
	// IsDF1SymbolMessageFrameNAK is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1SymbolMessageFrameNAK()
}

// _DF1SymbolMessageFrameNAK is the data-structure of this message
type _DF1SymbolMessageFrameNAK struct {
	DF1SymbolContract
}

var _ DF1SymbolMessageFrameNAK = (*_DF1SymbolMessageFrameNAK)(nil)
var _ DF1SymbolRequirements = (*_DF1SymbolMessageFrameNAK)(nil)

// NewDF1SymbolMessageFrameNAK factory function for _DF1SymbolMessageFrameNAK
func NewDF1SymbolMessageFrameNAK() *_DF1SymbolMessageFrameNAK {
	_result := &_DF1SymbolMessageFrameNAK{
		DF1SymbolContract: NewDF1Symbol(),
	}
	_result.DF1SymbolContract.(*_DF1Symbol)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrameNAK) GetSymbolType() uint8 {
	return 0x15
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrameNAK) GetParent() DF1SymbolContract {
	return m.DF1SymbolContract
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrameNAK(structType any) DF1SymbolMessageFrameNAK {
	if casted, ok := structType.(DF1SymbolMessageFrameNAK); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrameNAK); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrameNAK) GetTypeName() string {
	return "DF1SymbolMessageFrameNAK"
}

func (m *_DF1SymbolMessageFrameNAK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1SymbolContract.(*_DF1Symbol).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DF1SymbolMessageFrameNAK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1SymbolMessageFrameNAK) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Symbol) (__dF1SymbolMessageFrameNAK DF1SymbolMessageFrameNAK, err error) {
	m.DF1SymbolContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrameNAK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrameNAK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrameNAK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrameNAK")
	}

	return m, nil
}

func (m *_DF1SymbolMessageFrameNAK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrameNAK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrameNAK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrameNAK")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrameNAK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrameNAK")
		}
		return nil
	}
	return m.DF1SymbolContract.(*_DF1Symbol).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrameNAK) IsDF1SymbolMessageFrameNAK() {}

func (m *_DF1SymbolMessageFrameNAK) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1SymbolMessageFrameNAK) deepCopy() *_DF1SymbolMessageFrameNAK {
	if m == nil {
		return nil
	}
	_DF1SymbolMessageFrameNAKCopy := &_DF1SymbolMessageFrameNAK{
		m.DF1SymbolContract.DeepCopy().(DF1SymbolContract),
	}
	m.DF1SymbolContract.(*_DF1Symbol)._SubType = m
	return _DF1SymbolMessageFrameNAKCopy
}

func (m *_DF1SymbolMessageFrameNAK) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
