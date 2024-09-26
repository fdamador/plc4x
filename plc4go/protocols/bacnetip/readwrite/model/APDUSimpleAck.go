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

// APDUSimpleAck is the corresponding interface of APDUSimpleAck
type APDUSimpleAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetServiceChoice returns ServiceChoice (property field)
	GetServiceChoice() BACnetConfirmedServiceChoice
	// IsAPDUSimpleAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUSimpleAck()
}

// _APDUSimpleAck is the data-structure of this message
type _APDUSimpleAck struct {
	APDUContract
	OriginalInvokeId uint8
	ServiceChoice    BACnetConfirmedServiceChoice
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUSimpleAck = (*_APDUSimpleAck)(nil)
var _ APDURequirements = (*_APDUSimpleAck)(nil)

// NewAPDUSimpleAck factory function for _APDUSimpleAck
func NewAPDUSimpleAck(originalInvokeId uint8, serviceChoice BACnetConfirmedServiceChoice, apduLength uint16) *_APDUSimpleAck {
	_result := &_APDUSimpleAck{
		APDUContract:     NewAPDU(apduLength),
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSimpleAck) GetApduType() ApduType {
	return ApduType_SIMPLE_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSimpleAck) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSimpleAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSimpleAck) GetServiceChoice() BACnetConfirmedServiceChoice {
	return m.ServiceChoice
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUSimpleAck(structType any) APDUSimpleAck {
	if casted, ok := structType.(APDUSimpleAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSimpleAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSimpleAck) GetTypeName() string {
	return "APDUSimpleAck"
}

func (m *_APDUSimpleAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *_APDUSimpleAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUSimpleAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUSimpleAck APDUSimpleAck, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUSimpleAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSimpleAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	serviceChoice, err := ReadEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}
	m.ServiceChoice = serviceChoice

	if closeErr := readBuffer.CloseContext("APDUSimpleAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSimpleAck")
	}

	return m, nil
}

func (m *_APDUSimpleAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUSimpleAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSimpleAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSimpleAck")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", m.GetServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceChoice' field")
		}

		if popErr := writeBuffer.PopContext("APDUSimpleAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSimpleAck")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUSimpleAck) IsAPDUSimpleAck() {}

func (m *_APDUSimpleAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUSimpleAck) deepCopy() *_APDUSimpleAck {
	if m == nil {
		return nil
	}
	_APDUSimpleAckCopy := &_APDUSimpleAck{
		m.APDUContract.DeepCopy().(APDUContract),
		m.OriginalInvokeId,
		m.ServiceChoice,
		m.reservedField0,
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUSimpleAckCopy
}

func (m *_APDUSimpleAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
