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

// ModbusPDU is the corresponding interface of ModbusPDU
type ModbusPDU interface {
	ModbusPDUContract
	ModbusPDURequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsModbusPDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDU()
}

// ModbusPDUContract provides a set of functions which can be overwritten by a sub struct
type ModbusPDUContract interface {
	utils.Copyable
	// IsModbusPDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDU()
}

// ModbusPDURequirements provides a set of functions which need to be implemented by a sub struct
type ModbusPDURequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetErrorFlag returns ErrorFlag (discriminator field)
	GetErrorFlag() bool
	// GetFunctionFlag returns FunctionFlag (discriminator field)
	GetFunctionFlag() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// _ModbusPDU is the data-structure of this message
type _ModbusPDU struct {
	_SubType ModbusPDU
}

var _ ModbusPDUContract = (*_ModbusPDU)(nil)

// NewModbusPDU factory function for _ModbusPDU
func NewModbusPDU() *_ModbusPDU {
	return &_ModbusPDU{}
}

// Deprecated: use the interface for direct cast
func CastModbusPDU(structType any) ModbusPDU {
	if casted, ok := structType.(ModbusPDU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDU) GetTypeName() string {
	return "ModbusPDU"
}

func (m *_ModbusPDU) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (errorFlag)
	lengthInBits += 1
	// Discriminator Field (functionFlag)
	lengthInBits += 7

	return lengthInBits
}

func (m *_ModbusPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ModbusPDUParse[T ModbusPDU](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return ModbusPDUParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUParseWithBufferProducer[T ModbusPDU](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ModbusPDUParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ModbusPDUParseWithBuffer[T ModbusPDU](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_ModbusPDU{}).parse(ctx, readBuffer, response)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ModbusPDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__modbusPDU ModbusPDU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorFlag, err := ReadDiscriminatorField[bool](ctx, "errorFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorFlag' field"))
	}

	functionFlag, err := ReadDiscriminatorField[uint8](ctx, "functionFlag", ReadUnsignedByte(readBuffer, uint8(7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionFlag' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ModbusPDU
	switch {
	case errorFlag == bool(true): // ModbusPDUError
		if _child, err = new(_ModbusPDUError).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUError for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(false): // ModbusPDUReadDiscreteInputsRequest
		if _child, err = new(_ModbusPDUReadDiscreteInputsRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadDiscreteInputsRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(true): // ModbusPDUReadDiscreteInputsResponse
		if _child, err = new(_ModbusPDUReadDiscreteInputsResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadDiscreteInputsResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(false): // ModbusPDUReadCoilsRequest
		if _child, err = new(_ModbusPDUReadCoilsRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadCoilsRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(true): // ModbusPDUReadCoilsResponse
		if _child, err = new(_ModbusPDUReadCoilsResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadCoilsResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(false): // ModbusPDUWriteSingleCoilRequest
		if _child, err = new(_ModbusPDUWriteSingleCoilRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteSingleCoilRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(true): // ModbusPDUWriteSingleCoilResponse
		if _child, err = new(_ModbusPDUWriteSingleCoilResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteSingleCoilResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(false): // ModbusPDUWriteMultipleCoilsRequest
		if _child, err = new(_ModbusPDUWriteMultipleCoilsRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteMultipleCoilsRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(true): // ModbusPDUWriteMultipleCoilsResponse
		if _child, err = new(_ModbusPDUWriteMultipleCoilsResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteMultipleCoilsResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(false): // ModbusPDUReadInputRegistersRequest
		if _child, err = new(_ModbusPDUReadInputRegistersRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadInputRegistersRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(true): // ModbusPDUReadInputRegistersResponse
		if _child, err = new(_ModbusPDUReadInputRegistersResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadInputRegistersResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(false): // ModbusPDUReadHoldingRegistersRequest
		if _child, err = new(_ModbusPDUReadHoldingRegistersRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadHoldingRegistersRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(true): // ModbusPDUReadHoldingRegistersResponse
		if _child, err = new(_ModbusPDUReadHoldingRegistersResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadHoldingRegistersResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(false): // ModbusPDUWriteSingleRegisterRequest
		if _child, err = new(_ModbusPDUWriteSingleRegisterRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteSingleRegisterRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(true): // ModbusPDUWriteSingleRegisterResponse
		if _child, err = new(_ModbusPDUWriteSingleRegisterResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteSingleRegisterResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(false): // ModbusPDUWriteMultipleHoldingRegistersRequest
		if _child, err = new(_ModbusPDUWriteMultipleHoldingRegistersRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteMultipleHoldingRegistersRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(true): // ModbusPDUWriteMultipleHoldingRegistersResponse
		if _child, err = new(_ModbusPDUWriteMultipleHoldingRegistersResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteMultipleHoldingRegistersResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(false): // ModbusPDUReadWriteMultipleHoldingRegistersRequest
		if _child, err = new(_ModbusPDUReadWriteMultipleHoldingRegistersRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadWriteMultipleHoldingRegistersRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(true): // ModbusPDUReadWriteMultipleHoldingRegistersResponse
		if _child, err = new(_ModbusPDUReadWriteMultipleHoldingRegistersResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadWriteMultipleHoldingRegistersResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(false): // ModbusPDUMaskWriteHoldingRegisterRequest
		if _child, err = new(_ModbusPDUMaskWriteHoldingRegisterRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUMaskWriteHoldingRegisterRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(true): // ModbusPDUMaskWriteHoldingRegisterResponse
		if _child, err = new(_ModbusPDUMaskWriteHoldingRegisterResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUMaskWriteHoldingRegisterResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(false): // ModbusPDUReadFifoQueueRequest
		if _child, err = new(_ModbusPDUReadFifoQueueRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadFifoQueueRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(true): // ModbusPDUReadFifoQueueResponse
		if _child, err = new(_ModbusPDUReadFifoQueueResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadFifoQueueResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(false): // ModbusPDUReadFileRecordRequest
		if _child, err = new(_ModbusPDUReadFileRecordRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadFileRecordRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(true): // ModbusPDUReadFileRecordResponse
		if _child, err = new(_ModbusPDUReadFileRecordResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadFileRecordResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(false): // ModbusPDUWriteFileRecordRequest
		if _child, err = new(_ModbusPDUWriteFileRecordRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteFileRecordRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(true): // ModbusPDUWriteFileRecordResponse
		if _child, err = new(_ModbusPDUWriteFileRecordResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUWriteFileRecordResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(false): // ModbusPDUReadExceptionStatusRequest
		if _child, err = new(_ModbusPDUReadExceptionStatusRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadExceptionStatusRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(true): // ModbusPDUReadExceptionStatusResponse
		if _child, err = new(_ModbusPDUReadExceptionStatusResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadExceptionStatusResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(false): // ModbusPDUDiagnosticRequest
		if _child, err = new(_ModbusPDUDiagnosticRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUDiagnosticRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(true): // ModbusPDUDiagnosticResponse
		if _child, err = new(_ModbusPDUDiagnosticResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUDiagnosticResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(false): // ModbusPDUGetComEventCounterRequest
		if _child, err = new(_ModbusPDUGetComEventCounterRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUGetComEventCounterRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(true): // ModbusPDUGetComEventCounterResponse
		if _child, err = new(_ModbusPDUGetComEventCounterResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUGetComEventCounterResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(false): // ModbusPDUGetComEventLogRequest
		if _child, err = new(_ModbusPDUGetComEventLogRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUGetComEventLogRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(true): // ModbusPDUGetComEventLogResponse
		if _child, err = new(_ModbusPDUGetComEventLogResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUGetComEventLogResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(false): // ModbusPDUReportServerIdRequest
		if _child, err = new(_ModbusPDUReportServerIdRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReportServerIdRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(true): // ModbusPDUReportServerIdResponse
		if _child, err = new(_ModbusPDUReportServerIdResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReportServerIdResponse for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(false): // ModbusPDUReadDeviceIdentificationRequest
		if _child, err = new(_ModbusPDUReadDeviceIdentificationRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadDeviceIdentificationRequest for type-switch of ModbusPDU")
		}
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(true): // ModbusPDUReadDeviceIdentificationResponse
		if _child, err = new(_ModbusPDUReadDeviceIdentificationResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ModbusPDUReadDeviceIdentificationResponse for type-switch of ModbusPDU")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [errorFlag=%v, functionFlag=%v, response=%v]", errorFlag, functionFlag, response)
	}

	if closeErr := readBuffer.CloseContext("ModbusPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDU")
	}

	return _child, nil
}

func (pm *_ModbusPDU) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ModbusPDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusPDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDU")
	}

	if err := WriteDiscriminatorField(ctx, "errorFlag", m.GetErrorFlag(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorFlag' field")
	}

	if err := WriteDiscriminatorField(ctx, "functionFlag", m.GetFunctionFlag(), WriteUnsignedByte(writeBuffer, 7)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionFlag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDU")
	}
	return nil
}

func (m *_ModbusPDU) IsModbusPDU() {}

func (m *_ModbusPDU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDU) deepCopy() *_ModbusPDU {
	if m == nil {
		return nil
	}
	_ModbusPDUCopy := &_ModbusPDU{
		nil, // will be set by child
	}
	return _ModbusPDUCopy
}
