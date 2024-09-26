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

// LRawCon is the corresponding interface of LRawCon
type LRawCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsLRawCon is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLRawCon()
}

// _LRawCon is the data-structure of this message
type _LRawCon struct {
	CEMIContract
}

var _ LRawCon = (*_LRawCon)(nil)
var _ CEMIRequirements = (*_LRawCon)(nil)

// NewLRawCon factory function for _LRawCon
func NewLRawCon(size uint16) *_LRawCon {
	_result := &_LRawCon{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawCon) GetMessageCode() uint8 {
	return 0x2F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawCon) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastLRawCon(structType any) LRawCon {
	if casted, ok := structType.(LRawCon); ok {
		return casted
	}
	if casted, ok := structType.(*LRawCon); ok {
		return *casted
	}
	return nil
}

func (m *_LRawCon) GetTypeName() string {
	return "LRawCon"
}

func (m *_LRawCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LRawCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LRawCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lRawCon LRawCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawCon")
	}

	return m, nil
}

func (m *_LRawCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LRawCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawCon")
		}

		if popErr := writeBuffer.PopContext("LRawCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LRawCon) IsLRawCon() {}

func (m *_LRawCon) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LRawCon) deepCopy() *_LRawCon {
	if m == nil {
		return nil
	}
	_LRawConCopy := &_LRawCon{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _LRawConCopy
}

func (m *_LRawCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
