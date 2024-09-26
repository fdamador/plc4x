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

// Error is the corresponding interface of Error
type Error interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetErrorClass returns ErrorClass (property field)
	GetErrorClass() ErrorClassTagged
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() ErrorCodeTagged
	// IsError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsError()
}

// _Error is the data-structure of this message
type _Error struct {
	ErrorClass ErrorClassTagged
	ErrorCode  ErrorCodeTagged
}

var _ Error = (*_Error)(nil)

// NewError factory function for _Error
func NewError(errorClass ErrorClassTagged, errorCode ErrorCodeTagged) *_Error {
	if errorClass == nil {
		panic("errorClass of type ErrorClassTagged for Error must not be nil")
	}
	if errorCode == nil {
		panic("errorCode of type ErrorCodeTagged for Error must not be nil")
	}
	return &_Error{ErrorClass: errorClass, ErrorCode: errorCode}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Error) GetErrorClass() ErrorClassTagged {
	return m.ErrorClass
}

func (m *_Error) GetErrorCode() ErrorCodeTagged {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastError(structType any) Error {
	if casted, ok := structType.(Error); ok {
		return casted
	}
	if casted, ok := structType.(*Error); ok {
		return *casted
	}
	return nil
}

func (m *_Error) GetTypeName() string {
	return "Error"
}

func (m *_Error) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (errorClass)
	lengthInBits += m.ErrorClass.GetLengthInBits(ctx)

	// Simple field (errorCode)
	lengthInBits += m.ErrorCode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Error) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorParse(ctx context.Context, theBytes []byte) (Error, error) {
	return ErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
		return ErrorParseWithBuffer(ctx, readBuffer)
	}
}

func ErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
	v, err := (&_Error{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Error) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__error Error, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Error")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorClass, err := ReadSimpleField[ErrorClassTagged](ctx, "errorClass", ReadComplex[ErrorClassTagged](ErrorClassTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorClass' field"))
	}
	m.ErrorClass = errorClass

	errorCode, err := ReadSimpleField[ErrorCodeTagged](ctx, "errorCode", ReadComplex[ErrorCodeTagged](ErrorCodeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCode' field"))
	}
	m.ErrorCode = errorCode

	if closeErr := readBuffer.CloseContext("Error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Error")
	}

	return m, nil
}

func (m *_Error) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Error) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Error"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Error")
	}

	if err := WriteSimpleField[ErrorClassTagged](ctx, "errorClass", m.GetErrorClass(), WriteComplex[ErrorClassTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorClass' field")
	}

	if err := WriteSimpleField[ErrorCodeTagged](ctx, "errorCode", m.GetErrorCode(), WriteComplex[ErrorCodeTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorCode' field")
	}

	if popErr := writeBuffer.PopContext("Error"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Error")
	}
	return nil
}

func (m *_Error) IsError() {}

func (m *_Error) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Error) deepCopy() *_Error {
	if m == nil {
		return nil
	}
	_ErrorCopy := &_Error{
		m.ErrorClass.DeepCopy().(ErrorClassTagged),
		m.ErrorCode.DeepCopy().(ErrorCodeTagged),
	}
	return _ErrorCopy
}

func (m *_Error) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
