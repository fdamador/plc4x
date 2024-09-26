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

// KnxNetIpRouting is the corresponding interface of KnxNetIpRouting
type KnxNetIpRouting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetIpRouting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetIpRouting()
}

// _KnxNetIpRouting is the data-structure of this message
type _KnxNetIpRouting struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetIpRouting = (*_KnxNetIpRouting)(nil)
var _ ServiceIdRequirements = (*_KnxNetIpRouting)(nil)

// NewKnxNetIpRouting factory function for _KnxNetIpRouting
func NewKnxNetIpRouting(version uint8) *_KnxNetIpRouting {
	_result := &_KnxNetIpRouting{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpRouting) GetServiceType() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpRouting) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpRouting) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetIpRouting(structType any) KnxNetIpRouting {
	if casted, ok := structType.(KnxNetIpRouting); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpRouting); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpRouting) GetTypeName() string {
	return "KnxNetIpRouting"
}

func (m *_KnxNetIpRouting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpRouting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetIpRouting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetIpRouting KnxNetIpRouting, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpRouting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpRouting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetIpRouting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpRouting")
	}

	return m, nil
}

func (m *_KnxNetIpRouting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpRouting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpRouting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpRouting")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpRouting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpRouting")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpRouting) IsKnxNetIpRouting() {}

func (m *_KnxNetIpRouting) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetIpRouting) deepCopy() *_KnxNetIpRouting {
	if m == nil {
		return nil
	}
	_KnxNetIpRoutingCopy := &_KnxNetIpRouting{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	m.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetIpRoutingCopy
}

func (m *_KnxNetIpRouting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
