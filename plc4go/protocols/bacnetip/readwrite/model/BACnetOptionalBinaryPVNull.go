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

// BACnetOptionalBinaryPVNull is the corresponding interface of BACnetOptionalBinaryPVNull
type BACnetOptionalBinaryPVNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetOptionalBinaryPV
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetOptionalBinaryPVNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalBinaryPVNull()
}

// _BACnetOptionalBinaryPVNull is the data-structure of this message
type _BACnetOptionalBinaryPVNull struct {
	BACnetOptionalBinaryPVContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetOptionalBinaryPVNull = (*_BACnetOptionalBinaryPVNull)(nil)
var _ BACnetOptionalBinaryPVRequirements = (*_BACnetOptionalBinaryPVNull)(nil)

// NewBACnetOptionalBinaryPVNull factory function for _BACnetOptionalBinaryPVNull
func NewBACnetOptionalBinaryPVNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull) *_BACnetOptionalBinaryPVNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetOptionalBinaryPVNull must not be nil")
	}
	_result := &_BACnetOptionalBinaryPVNull{
		BACnetOptionalBinaryPVContract: NewBACnetOptionalBinaryPV(peekedTagHeader),
		NullValue:                      nullValue,
	}
	_result.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV)._SubType = _result
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

func (m *_BACnetOptionalBinaryPVNull) GetParent() BACnetOptionalBinaryPVContract {
	return m.BACnetOptionalBinaryPVContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalBinaryPVNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalBinaryPVNull(structType any) BACnetOptionalBinaryPVNull {
	if casted, ok := structType.(BACnetOptionalBinaryPVNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPVNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalBinaryPVNull) GetTypeName() string {
	return "BACnetOptionalBinaryPVNull"
}

func (m *_BACnetOptionalBinaryPVNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalBinaryPVNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalBinaryPVNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalBinaryPV) (__bACnetOptionalBinaryPVNull BACnetOptionalBinaryPVNull, err error) {
	m.BACnetOptionalBinaryPVContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPVNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPVNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPVNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPVNull")
	}

	return m, nil
}

func (m *_BACnetOptionalBinaryPVNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalBinaryPVNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPVNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPVNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPVNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPVNull")
		}
		return nil
	}
	return m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalBinaryPVNull) IsBACnetOptionalBinaryPVNull() {}

func (m *_BACnetOptionalBinaryPVNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalBinaryPVNull) deepCopy() *_BACnetOptionalBinaryPVNull {
	if m == nil {
		return nil
	}
	_BACnetOptionalBinaryPVNullCopy := &_BACnetOptionalBinaryPVNull{
		m.BACnetOptionalBinaryPVContract.DeepCopy().(BACnetOptionalBinaryPVContract),
		m.NullValue.DeepCopy().(BACnetApplicationTagNull),
	}
	m.BACnetOptionalBinaryPVContract.(*_BACnetOptionalBinaryPV)._SubType = m
	return _BACnetOptionalBinaryPVNullCopy
}

func (m *_BACnetOptionalBinaryPVNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
