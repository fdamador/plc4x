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

// SysexCommandAnalogMappingQueryRequest is the corresponding interface of SysexCommandAnalogMappingQueryRequest
type SysexCommandAnalogMappingQueryRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// IsSysexCommandAnalogMappingQueryRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandAnalogMappingQueryRequest()
}

// _SysexCommandAnalogMappingQueryRequest is the data-structure of this message
type _SysexCommandAnalogMappingQueryRequest struct {
	SysexCommandContract
}

var _ SysexCommandAnalogMappingQueryRequest = (*_SysexCommandAnalogMappingQueryRequest)(nil)
var _ SysexCommandRequirements = (*_SysexCommandAnalogMappingQueryRequest)(nil)

// NewSysexCommandAnalogMappingQueryRequest factory function for _SysexCommandAnalogMappingQueryRequest
func NewSysexCommandAnalogMappingQueryRequest() *_SysexCommandAnalogMappingQueryRequest {
	_result := &_SysexCommandAnalogMappingQueryRequest{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryRequest) GetCommandType() uint8 {
	return 0x69
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandAnalogMappingQueryRequest) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandAnalogMappingQueryRequest(structType any) SysexCommandAnalogMappingQueryRequest {
	if casted, ok := structType.(SysexCommandAnalogMappingQueryRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingQueryRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryRequest"
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandAnalogMappingQueryRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandAnalogMappingQueryRequest SysexCommandAnalogMappingQueryRequest, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandAnalogMappingQueryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandAnalogMappingQueryRequest")
	}

	return m, nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandAnalogMappingQueryRequest")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandAnalogMappingQueryRequest")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandAnalogMappingQueryRequest) IsSysexCommandAnalogMappingQueryRequest() {}

func (m *_SysexCommandAnalogMappingQueryRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandAnalogMappingQueryRequest) deepCopy() *_SysexCommandAnalogMappingQueryRequest {
	if m == nil {
		return nil
	}
	_SysexCommandAnalogMappingQueryRequestCopy := &_SysexCommandAnalogMappingQueryRequest{
		m.SysexCommandContract.DeepCopy().(SysexCommandContract),
	}
	m.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandAnalogMappingQueryRequestCopy
}

func (m *_SysexCommandAnalogMappingQueryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
