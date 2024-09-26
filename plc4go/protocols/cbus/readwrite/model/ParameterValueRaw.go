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

// ParameterValueRaw is the corresponding interface of ParameterValueRaw
type ParameterValueRaw interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueRaw is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueRaw()
}

// _ParameterValueRaw is the data-structure of this message
type _ParameterValueRaw struct {
	ParameterValueContract
	Data []byte
}

var _ ParameterValueRaw = (*_ParameterValueRaw)(nil)
var _ ParameterValueRequirements = (*_ParameterValueRaw)(nil)

// NewParameterValueRaw factory function for _ParameterValueRaw
func NewParameterValueRaw(data []byte, numBytes uint8) *_ParameterValueRaw {
	_result := &_ParameterValueRaw{
		ParameterValueContract: NewParameterValue(numBytes),
		Data:                   data,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueRaw) GetParameterType() ParameterType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueRaw) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueRaw) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueRaw(structType any) ParameterValueRaw {
	if casted, ok := structType.(ParameterValueRaw); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueRaw); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueRaw) GetTypeName() string {
	return "ParameterValueRaw"
}

func (m *_ParameterValueRaw) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueRaw) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueRaw) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueRaw ParameterValueRaw, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueRaw"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueRaw")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := readBuffer.ReadByteArray("data", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueRaw"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueRaw")
	}

	return m, nil
}

func (m *_ParameterValueRaw) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueRaw) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueRaw"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueRaw")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueRaw"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueRaw")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueRaw) IsParameterValueRaw() {}

func (m *_ParameterValueRaw) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueRaw) deepCopy() *_ParameterValueRaw {
	if m == nil {
		return nil
	}
	_ParameterValueRawCopy := &_ParameterValueRaw{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueRawCopy
}

func (m *_ParameterValueRaw) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
