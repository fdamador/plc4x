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

// ModbusPDUMaskWriteHoldingRegisterResponse is the corresponding interface of ModbusPDUMaskWriteHoldingRegisterResponse
type ModbusPDUMaskWriteHoldingRegisterResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetReferenceAddress returns ReferenceAddress (property field)
	GetReferenceAddress() uint16
	// GetAndMask returns AndMask (property field)
	GetAndMask() uint16
	// GetOrMask returns OrMask (property field)
	GetOrMask() uint16
	// IsModbusPDUMaskWriteHoldingRegisterResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUMaskWriteHoldingRegisterResponse()
}

// _ModbusPDUMaskWriteHoldingRegisterResponse is the data-structure of this message
type _ModbusPDUMaskWriteHoldingRegisterResponse struct {
	ModbusPDUContract
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
}

var _ ModbusPDUMaskWriteHoldingRegisterResponse = (*_ModbusPDUMaskWriteHoldingRegisterResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUMaskWriteHoldingRegisterResponse)(nil)

// NewModbusPDUMaskWriteHoldingRegisterResponse factory function for _ModbusPDUMaskWriteHoldingRegisterResponse
func NewModbusPDUMaskWriteHoldingRegisterResponse(referenceAddress uint16, andMask uint16, orMask uint16) *_ModbusPDUMaskWriteHoldingRegisterResponse {
	_result := &_ModbusPDUMaskWriteHoldingRegisterResponse{
		ModbusPDUContract: NewModbusPDU(),
		ReferenceAddress:  referenceAddress,
		AndMask:           andMask,
		OrMask:            orMask,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetFunctionFlag() uint8 {
	return 0x16
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetReferenceAddress() uint16 {
	return m.ReferenceAddress
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetAndMask() uint16 {
	return m.AndMask
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetOrMask() uint16 {
	return m.OrMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUMaskWriteHoldingRegisterResponse(structType any) ModbusPDUMaskWriteHoldingRegisterResponse {
	if casted, ok := structType.(ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterResponse"
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUMaskWriteHoldingRegisterResponse ModbusPDUMaskWriteHoldingRegisterResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUMaskWriteHoldingRegisterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceAddress, err := ReadSimpleField(ctx, "referenceAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceAddress' field"))
	}
	m.ReferenceAddress = referenceAddress

	andMask, err := ReadSimpleField(ctx, "andMask", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'andMask' field"))
	}
	m.AndMask = andMask

	orMask, err := ReadSimpleField(ctx, "orMask", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'orMask' field"))
	}
	m.OrMask = orMask

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUMaskWriteHoldingRegisterResponse")
	}

	return m, nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUMaskWriteHoldingRegisterResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "referenceAddress", m.GetReferenceAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "andMask", m.GetAndMask(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'andMask' field")
		}

		if err := WriteSimpleField[uint16](ctx, "orMask", m.GetOrMask(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUMaskWriteHoldingRegisterResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) IsModbusPDUMaskWriteHoldingRegisterResponse() {}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) deepCopy() *_ModbusPDUMaskWriteHoldingRegisterResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUMaskWriteHoldingRegisterResponseCopy := &_ModbusPDUMaskWriteHoldingRegisterResponse{
		m.ModbusPDUContract.DeepCopy().(ModbusPDUContract),
		m.ReferenceAddress,
		m.AndMask,
		m.OrMask,
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUMaskWriteHoldingRegisterResponseCopy
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
