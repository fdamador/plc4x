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

// ModbusPDUReportServerIdRequest is the corresponding interface of ModbusPDUReportServerIdRequest
type ModbusPDUReportServerIdRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// IsModbusPDUReportServerIdRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReportServerIdRequest()
}

// _ModbusPDUReportServerIdRequest is the data-structure of this message
type _ModbusPDUReportServerIdRequest struct {
	ModbusPDUContract
}

var _ ModbusPDUReportServerIdRequest = (*_ModbusPDUReportServerIdRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReportServerIdRequest)(nil)

// NewModbusPDUReportServerIdRequest factory function for _ModbusPDUReportServerIdRequest
func NewModbusPDUReportServerIdRequest() *_ModbusPDUReportServerIdRequest {
	_result := &_ModbusPDUReportServerIdRequest{
		ModbusPDUContract: NewModbusPDU(),
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReportServerIdRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReportServerIdRequest) GetFunctionFlag() uint8 {
	return 0x11
}

func (m *_ModbusPDUReportServerIdRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReportServerIdRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReportServerIdRequest(structType any) ModbusPDUReportServerIdRequest {
	if casted, ok := structType.(ModbusPDUReportServerIdRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReportServerIdRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReportServerIdRequest) GetTypeName() string {
	return "ModbusPDUReportServerIdRequest"
}

func (m *_ModbusPDUReportServerIdRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ModbusPDUReportServerIdRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReportServerIdRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReportServerIdRequest ModbusPDUReportServerIdRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReportServerIdRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReportServerIdRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReportServerIdRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReportServerIdRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReportServerIdRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReportServerIdRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReportServerIdRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReportServerIdRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReportServerIdRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReportServerIdRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReportServerIdRequest) IsModbusPDUReportServerIdRequest() {}

func (m *_ModbusPDUReportServerIdRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReportServerIdRequest) deepCopy() *_ModbusPDUReportServerIdRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReportServerIdRequestCopy := &_ModbusPDUReportServerIdRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReportServerIdRequestCopy
}

func (m *_ModbusPDUReportServerIdRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
