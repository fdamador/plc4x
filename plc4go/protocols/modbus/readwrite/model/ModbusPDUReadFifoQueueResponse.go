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

// ModbusPDUReadFifoQueueResponse is the corresponding interface of ModbusPDUReadFifoQueueResponse
type ModbusPDUReadFifoQueueResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetFifoValue returns FifoValue (property field)
	GetFifoValue() []uint16
	// IsModbusPDUReadFifoQueueResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadFifoQueueResponse()
}

// _ModbusPDUReadFifoQueueResponse is the data-structure of this message
type _ModbusPDUReadFifoQueueResponse struct {
	ModbusPDUContract
	FifoValue []uint16
}

var _ ModbusPDUReadFifoQueueResponse = (*_ModbusPDUReadFifoQueueResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadFifoQueueResponse)(nil)

// NewModbusPDUReadFifoQueueResponse factory function for _ModbusPDUReadFifoQueueResponse
func NewModbusPDUReadFifoQueueResponse(fifoValue []uint16) *_ModbusPDUReadFifoQueueResponse {
	_result := &_ModbusPDUReadFifoQueueResponse{
		ModbusPDUContract: NewModbusPDU(),
		FifoValue:         fifoValue,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFifoQueueResponse) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *_ModbusPDUReadFifoQueueResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) GetFifoValue() []uint16 {
	return m.FifoValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFifoQueueResponse(structType any) ModbusPDUReadFifoQueueResponse {
	if casted, ok := structType.(ModbusPDUReadFifoQueueResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFifoQueueResponse) GetTypeName() string {
	return "ModbusPDUReadFifoQueueResponse"
}

func (m *_ModbusPDUReadFifoQueueResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 16

	// Implicit Field (fifoCount)
	lengthInBits += 16

	// Array field
	if len(m.FifoValue) > 0 {
		lengthInBits += 16 * uint16(len(m.FifoValue))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFifoQueueResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadFifoQueueResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadFifoQueueResponse ModbusPDUReadFifoQueueResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFifoQueueResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint16](ctx, "byteCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	fifoCount, err := ReadImplicitField[uint16](ctx, "fifoCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fifoCount' field"))
	}
	_ = fifoCount

	fifoValue, err := ReadCountArrayField[uint16](ctx, "fifoValue", ReadUnsignedShort(readBuffer, uint8(16)), uint64(fifoCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fifoValue' field"))
	}
	m.FifoValue = fifoValue

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFifoQueueResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadFifoQueueResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFifoQueueResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFifoQueueResponse")
		}
		byteCount := uint16(uint16((uint16(uint16(len(m.GetFifoValue()))) * uint16(uint16(2)))) + uint16(uint16(2)))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}
		fifoCount := uint16(uint16((uint16(uint16(len(m.GetFifoValue()))) * uint16(uint16(2)))) / uint16(uint16(2)))
		if err := WriteImplicitField(ctx, "fifoCount", fifoCount, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'fifoCount' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "fifoValue", m.GetFifoValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'fifoValue' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFifoQueueResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFifoQueueResponse) IsModbusPDUReadFifoQueueResponse() {}

func (m *_ModbusPDUReadFifoQueueResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadFifoQueueResponse) deepCopy() *_ModbusPDUReadFifoQueueResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUReadFifoQueueResponseCopy := &_ModbusPDUReadFifoQueueResponse{
		m.ModbusPDUContract.DeepCopy().(ModbusPDUContract),
		utils.DeepCopySlice[uint16, uint16](m.FifoValue),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadFifoQueueResponseCopy
}

func (m *_ModbusPDUReadFifoQueueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
