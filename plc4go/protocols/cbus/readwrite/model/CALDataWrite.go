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

// CALDataWrite is the corresponding interface of CALDataWrite
type CALDataWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCode returns Code (property field)
	GetCode() byte
	// GetParameterValue returns ParameterValue (property field)
	GetParameterValue() ParameterValue
	// IsCALDataWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALDataWrite()
}

// _CALDataWrite is the data-structure of this message
type _CALDataWrite struct {
	CALDataContract
	ParamNo        Parameter
	Code           byte
	ParameterValue ParameterValue
}

var _ CALDataWrite = (*_CALDataWrite)(nil)
var _ CALDataRequirements = (*_CALDataWrite)(nil)

// NewCALDataWrite factory function for _CALDataWrite
func NewCALDataWrite(commandTypeContainer CALCommandTypeContainer, additionalData CALData, paramNo Parameter, code byte, parameterValue ParameterValue, requestContext RequestContext) *_CALDataWrite {
	if parameterValue == nil {
		panic("parameterValue of type ParameterValue for CALDataWrite must not be nil")
	}
	_result := &_CALDataWrite{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		Code:            code,
		ParameterValue:  parameterValue,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataWrite) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataWrite) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataWrite) GetCode() byte {
	return m.Code
}

func (m *_CALDataWrite) GetParameterValue() ParameterValue {
	return m.ParameterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCALDataWrite(structType any) CALDataWrite {
	if casted, ok := structType.(CALDataWrite); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataWrite); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataWrite) GetTypeName() string {
	return "CALDataWrite"
}

func (m *_CALDataWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += m.ParameterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALDataWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (__cALDataWrite CALDataWrite, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	code, err := ReadSimpleField(ctx, "code", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}
	m.Code = code

	parameterValue, err := ReadSimpleField[ParameterValue](ctx, "parameterValue", ReadComplex[ParameterValue](ParameterValueParseWithBufferProducer[ParameterValue]((ParameterType)(paramNo.ParameterType()), (uint8)(uint8(commandTypeContainer.NumBytes())-uint8(uint8(2)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterValue' field"))
	}
	m.ParameterValue = parameterValue

	if closeErr := readBuffer.CloseContext("CALDataWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataWrite")
	}

	return m, nil
}

func (m *_CALDataWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataWrite")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[byte](ctx, "code", m.GetCode(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'code' field")
		}

		if err := WriteSimpleField[ParameterValue](ctx, "parameterValue", m.GetParameterValue(), WriteComplex[ParameterValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CALDataWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataWrite")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataWrite) IsCALDataWrite() {}

func (m *_CALDataWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALDataWrite) deepCopy() *_CALDataWrite {
	if m == nil {
		return nil
	}
	_CALDataWriteCopy := &_CALDataWrite{
		m.CALDataContract.(*_CALData).deepCopy(),
		m.ParamNo,
		m.Code,
		m.ParameterValue.DeepCopy().(ParameterValue),
	}
	m.CALDataContract.(*_CALData)._SubType = m
	return _CALDataWriteCopy
}

func (m *_CALDataWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
