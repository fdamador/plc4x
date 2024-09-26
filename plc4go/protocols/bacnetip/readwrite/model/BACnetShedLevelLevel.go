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

// BACnetShedLevelLevel is the corresponding interface of BACnetShedLevelLevel
type BACnetShedLevelLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetShedLevel
	// GetLevel returns Level (property field)
	GetLevel() BACnetContextTagUnsignedInteger
	// IsBACnetShedLevelLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevelLevel()
}

// _BACnetShedLevelLevel is the data-structure of this message
type _BACnetShedLevelLevel struct {
	BACnetShedLevelContract
	Level BACnetContextTagUnsignedInteger
}

var _ BACnetShedLevelLevel = (*_BACnetShedLevelLevel)(nil)
var _ BACnetShedLevelRequirements = (*_BACnetShedLevelLevel)(nil)

// NewBACnetShedLevelLevel factory function for _BACnetShedLevelLevel
func NewBACnetShedLevelLevel(peekedTagHeader BACnetTagHeader, level BACnetContextTagUnsignedInteger) *_BACnetShedLevelLevel {
	if level == nil {
		panic("level of type BACnetContextTagUnsignedInteger for BACnetShedLevelLevel must not be nil")
	}
	_result := &_BACnetShedLevelLevel{
		BACnetShedLevelContract: NewBACnetShedLevel(peekedTagHeader),
		Level:                   level,
	}
	_result.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = _result
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

func (m *_BACnetShedLevelLevel) GetParent() BACnetShedLevelContract {
	return m.BACnetShedLevelContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelLevel) GetLevel() BACnetContextTagUnsignedInteger {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelLevel(structType any) BACnetShedLevelLevel {
	if casted, ok := structType.(BACnetShedLevelLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelLevel) GetTypeName() string {
	return "BACnetShedLevelLevel"
}

func (m *_BACnetShedLevelLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetShedLevelContract.(*_BACnetShedLevel).getLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += m.Level.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetShedLevelLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetShedLevelLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetShedLevel) (__bACnetShedLevelLevel BACnetShedLevelLevel, err error) {
	m.BACnetShedLevelContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	level, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "level", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	if closeErr := readBuffer.CloseContext("BACnetShedLevelLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelLevel")
	}

	return m, nil
}

func (m *_BACnetShedLevelLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelLevel")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "level", m.GetLevel(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelLevel")
		}
		return nil
	}
	return m.BACnetShedLevelContract.(*_BACnetShedLevel).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetShedLevelLevel) IsBACnetShedLevelLevel() {}

func (m *_BACnetShedLevelLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetShedLevelLevel) deepCopy() *_BACnetShedLevelLevel {
	if m == nil {
		return nil
	}
	_BACnetShedLevelLevelCopy := &_BACnetShedLevelLevel{
		m.BACnetShedLevelContract.DeepCopy().(BACnetShedLevelContract),
		m.Level.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = m
	return _BACnetShedLevelLevelCopy
}

func (m *_BACnetShedLevelLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
