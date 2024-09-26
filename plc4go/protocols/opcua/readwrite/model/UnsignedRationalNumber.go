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

// UnsignedRationalNumber is the corresponding interface of UnsignedRationalNumber
type UnsignedRationalNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNumerator returns Numerator (property field)
	GetNumerator() uint32
	// GetDenominator returns Denominator (property field)
	GetDenominator() uint32
	// IsUnsignedRationalNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUnsignedRationalNumber()
}

// _UnsignedRationalNumber is the data-structure of this message
type _UnsignedRationalNumber struct {
	ExtensionObjectDefinitionContract
	Numerator   uint32
	Denominator uint32
}

var _ UnsignedRationalNumber = (*_UnsignedRationalNumber)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UnsignedRationalNumber)(nil)

// NewUnsignedRationalNumber factory function for _UnsignedRationalNumber
func NewUnsignedRationalNumber(numerator uint32, denominator uint32) *_UnsignedRationalNumber {
	_result := &_UnsignedRationalNumber{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Numerator:                         numerator,
		Denominator:                       denominator,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnsignedRationalNumber) GetIdentifier() string {
	return "24109"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnsignedRationalNumber) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnsignedRationalNumber) GetNumerator() uint32 {
	return m.Numerator
}

func (m *_UnsignedRationalNumber) GetDenominator() uint32 {
	return m.Denominator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUnsignedRationalNumber(structType any) UnsignedRationalNumber {
	if casted, ok := structType.(UnsignedRationalNumber); ok {
		return casted
	}
	if casted, ok := structType.(*UnsignedRationalNumber); ok {
		return *casted
	}
	return nil
}

func (m *_UnsignedRationalNumber) GetTypeName() string {
	return "UnsignedRationalNumber"
}

func (m *_UnsignedRationalNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (numerator)
	lengthInBits += 32

	// Simple field (denominator)
	lengthInBits += 32

	return lengthInBits
}

func (m *_UnsignedRationalNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UnsignedRationalNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__unsignedRationalNumber UnsignedRationalNumber, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnsignedRationalNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnsignedRationalNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numerator, err := ReadSimpleField(ctx, "numerator", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numerator' field"))
	}
	m.Numerator = numerator

	denominator, err := ReadSimpleField(ctx, "denominator", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'denominator' field"))
	}
	m.Denominator = denominator

	if closeErr := readBuffer.CloseContext("UnsignedRationalNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnsignedRationalNumber")
	}

	return m, nil
}

func (m *_UnsignedRationalNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnsignedRationalNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnsignedRationalNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnsignedRationalNumber")
		}

		if err := WriteSimpleField[uint32](ctx, "numerator", m.GetNumerator(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'numerator' field")
		}

		if err := WriteSimpleField[uint32](ctx, "denominator", m.GetDenominator(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'denominator' field")
		}

		if popErr := writeBuffer.PopContext("UnsignedRationalNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnsignedRationalNumber")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UnsignedRationalNumber) IsUnsignedRationalNumber() {}

func (m *_UnsignedRationalNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UnsignedRationalNumber) deepCopy() *_UnsignedRationalNumber {
	if m == nil {
		return nil
	}
	_UnsignedRationalNumberCopy := &_UnsignedRationalNumber{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.Numerator,
		m.Denominator,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UnsignedRationalNumberCopy
}

func (m *_UnsignedRationalNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
