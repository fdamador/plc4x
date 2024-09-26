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

// SecurityDataZoneSealed is the corresponding interface of SecurityDataZoneSealed
type SecurityDataZoneSealed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// IsSecurityDataZoneSealed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataZoneSealed()
}

// _SecurityDataZoneSealed is the data-structure of this message
type _SecurityDataZoneSealed struct {
	SecurityDataContract
	ZoneNumber uint8
}

var _ SecurityDataZoneSealed = (*_SecurityDataZoneSealed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneSealed)(nil)

// NewSecurityDataZoneSealed factory function for _SecurityDataZoneSealed
func NewSecurityDataZoneSealed(commandTypeContainer SecurityCommandTypeContainer, argument byte, zoneNumber uint8) *_SecurityDataZoneSealed {
	_result := &_SecurityDataZoneSealed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
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

func (m *_SecurityDataZoneSealed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneSealed) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneSealed(structType any) SecurityDataZoneSealed {
	if casted, ok := structType.(SecurityDataZoneSealed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneSealed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneSealed) GetTypeName() string {
	return "SecurityDataZoneSealed"
}

func (m *_SecurityDataZoneSealed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneSealed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneSealed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneSealed SecurityDataZoneSealed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneSealed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneSealed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneSealed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneSealed")
	}

	return m, nil
}

func (m *_SecurityDataZoneSealed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneSealed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneSealed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneSealed")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneSealed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneSealed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneSealed) IsSecurityDataZoneSealed() {}

func (m *_SecurityDataZoneSealed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataZoneSealed) deepCopy() *_SecurityDataZoneSealed {
	if m == nil {
		return nil
	}
	_SecurityDataZoneSealedCopy := &_SecurityDataZoneSealed{
		m.SecurityDataContract.DeepCopy().(SecurityDataContract),
		m.ZoneNumber,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataZoneSealedCopy
}

func (m *_SecurityDataZoneSealed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
