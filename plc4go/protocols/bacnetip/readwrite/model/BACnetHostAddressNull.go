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

// BACnetHostAddressNull is the corresponding interface of BACnetHostAddressNull
type BACnetHostAddressNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetHostAddress
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
	// IsBACnetHostAddressNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostAddressNull()
}

// _BACnetHostAddressNull is the data-structure of this message
type _BACnetHostAddressNull struct {
	BACnetHostAddressContract
	None BACnetContextTagNull
}

var _ BACnetHostAddressNull = (*_BACnetHostAddressNull)(nil)
var _ BACnetHostAddressRequirements = (*_BACnetHostAddressNull)(nil)

// NewBACnetHostAddressNull factory function for _BACnetHostAddressNull
func NewBACnetHostAddressNull(peekedTagHeader BACnetTagHeader, none BACnetContextTagNull) *_BACnetHostAddressNull {
	if none == nil {
		panic("none of type BACnetContextTagNull for BACnetHostAddressNull must not be nil")
	}
	_result := &_BACnetHostAddressNull{
		BACnetHostAddressContract: NewBACnetHostAddress(peekedTagHeader),
		None:                      none,
	}
	_result.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = _result
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

func (m *_BACnetHostAddressNull) GetParent() BACnetHostAddressContract {
	return m.BACnetHostAddressContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressNull) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressNull(structType any) BACnetHostAddressNull {
	if casted, ok := structType.(BACnetHostAddressNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressNull) GetTypeName() string {
	return "BACnetHostAddressNull"
}

func (m *_BACnetHostAddressNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetHostAddressContract.(*_BACnetHostAddress).getLengthInBits(ctx))

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetHostAddressNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetHostAddress) (__bACnetHostAddressNull BACnetHostAddressNull, err error) {
	m.BACnetHostAddressContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	none, err := ReadSimpleField[BACnetContextTagNull](ctx, "none", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'none' field"))
	}
	m.None = none

	if closeErr := readBuffer.CloseContext("BACnetHostAddressNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressNull")
	}

	return m, nil
}

func (m *_BACnetHostAddressNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressNull")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "none", m.GetNone(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'none' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressNull")
		}
		return nil
	}
	return m.BACnetHostAddressContract.(*_BACnetHostAddress).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetHostAddressNull) IsBACnetHostAddressNull() {}

func (m *_BACnetHostAddressNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetHostAddressNull) deepCopy() *_BACnetHostAddressNull {
	if m == nil {
		return nil
	}
	_BACnetHostAddressNullCopy := &_BACnetHostAddressNull{
		m.BACnetHostAddressContract.(*_BACnetHostAddress).deepCopy(),
		m.None.DeepCopy().(BACnetContextTagNull),
	}
	m.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = m
	return _BACnetHostAddressNullCopy
}

func (m *_BACnetHostAddressNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
