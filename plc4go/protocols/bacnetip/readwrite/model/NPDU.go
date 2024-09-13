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

// NPDU is the corresponding interface of NPDU
type NPDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetProtocolVersionNumber returns ProtocolVersionNumber (property field)
	GetProtocolVersionNumber() uint8
	// GetControl returns Control (property field)
	GetControl() NPDUControl
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() *uint16
	// GetDestinationLength returns DestinationLength (property field)
	GetDestinationLength() *uint8
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() []uint8
	// GetSourceNetworkAddress returns SourceNetworkAddress (property field)
	GetSourceNetworkAddress() *uint16
	// GetSourceLength returns SourceLength (property field)
	GetSourceLength() *uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() []uint8
	// GetHopCount returns HopCount (property field)
	GetHopCount() *uint8
	// GetNlm returns Nlm (property field)
	GetNlm() NLM
	// GetApdu returns Apdu (property field)
	GetApdu() APDU
	// GetDestinationLengthAddon returns DestinationLengthAddon (virtual field)
	GetDestinationLengthAddon() uint16
	// GetSourceLengthAddon returns SourceLengthAddon (virtual field)
	GetSourceLengthAddon() uint16
	// GetPayloadSubtraction returns PayloadSubtraction (virtual field)
	GetPayloadSubtraction() uint16
	// IsNPDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNPDU()
}

// _NPDU is the data-structure of this message
type _NPDU struct {
	ProtocolVersionNumber     uint8
	Control                   NPDUControl
	DestinationNetworkAddress *uint16
	DestinationLength         *uint8
	DestinationAddress        []uint8
	SourceNetworkAddress      *uint16
	SourceLength              *uint8
	SourceAddress             []uint8
	HopCount                  *uint8
	Nlm                       NLM
	Apdu                      APDU

	// Arguments.
	NpduLength uint16
}

var _ NPDU = (*_NPDU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NPDU) GetProtocolVersionNumber() uint8 {
	return m.ProtocolVersionNumber
}

func (m *_NPDU) GetControl() NPDUControl {
	return m.Control
}

func (m *_NPDU) GetDestinationNetworkAddress() *uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NPDU) GetDestinationLength() *uint8 {
	return m.DestinationLength
}

func (m *_NPDU) GetDestinationAddress() []uint8 {
	return m.DestinationAddress
}

func (m *_NPDU) GetSourceNetworkAddress() *uint16 {
	return m.SourceNetworkAddress
}

func (m *_NPDU) GetSourceLength() *uint8 {
	return m.SourceLength
}

func (m *_NPDU) GetSourceAddress() []uint8 {
	return m.SourceAddress
}

func (m *_NPDU) GetHopCount() *uint8 {
	return m.HopCount
}

func (m *_NPDU) GetNlm() NLM {
	return m.Nlm
}

func (m *_NPDU) GetApdu() APDU {
	return m.Apdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NPDU) GetDestinationLengthAddon() uint16 {
	ctx := context.Background()
	_ = ctx
	destinationNetworkAddress := m.GetDestinationNetworkAddress()
	_ = destinationNetworkAddress
	destinationLength := m.GetDestinationLength()
	_ = destinationLength
	sourceNetworkAddress := m.GetSourceNetworkAddress()
	_ = sourceNetworkAddress
	sourceLength := m.GetSourceLength()
	_ = sourceLength
	hopCount := m.GetHopCount()
	_ = hopCount
	nlm := m.GetNlm()
	_ = nlm
	apdu := m.GetApdu()
	_ = apdu
	return uint16(utils.InlineIf(m.GetControl().GetDestinationSpecified(), func() any { return uint16((uint16(uint16(3)) + uint16((*m.GetDestinationLength())))) }, func() any { return uint16(uint16(0)) }).(uint16))
}

func (m *_NPDU) GetSourceLengthAddon() uint16 {
	ctx := context.Background()
	_ = ctx
	destinationNetworkAddress := m.GetDestinationNetworkAddress()
	_ = destinationNetworkAddress
	destinationLength := m.GetDestinationLength()
	_ = destinationLength
	sourceNetworkAddress := m.GetSourceNetworkAddress()
	_ = sourceNetworkAddress
	sourceLength := m.GetSourceLength()
	_ = sourceLength
	hopCount := m.GetHopCount()
	_ = hopCount
	nlm := m.GetNlm()
	_ = nlm
	apdu := m.GetApdu()
	_ = apdu
	return uint16(utils.InlineIf(m.GetControl().GetSourceSpecified(), func() any { return uint16((uint16(uint16(3)) + uint16((*m.GetSourceLength())))) }, func() any { return uint16(uint16(0)) }).(uint16))
}

func (m *_NPDU) GetPayloadSubtraction() uint16 {
	ctx := context.Background()
	_ = ctx
	destinationNetworkAddress := m.GetDestinationNetworkAddress()
	_ = destinationNetworkAddress
	destinationLength := m.GetDestinationLength()
	_ = destinationLength
	sourceNetworkAddress := m.GetSourceNetworkAddress()
	_ = sourceNetworkAddress
	sourceLength := m.GetSourceLength()
	_ = sourceLength
	hopCount := m.GetHopCount()
	_ = hopCount
	nlm := m.GetNlm()
	_ = nlm
	apdu := m.GetApdu()
	_ = apdu
	return uint16(uint16(uint16(2)) + uint16((uint16(uint16(m.GetSourceLengthAddon())+uint16(m.GetDestinationLengthAddon())) + uint16((utils.InlineIf((m.GetControl().GetDestinationSpecified()), func() any { return uint16(uint16(1)) }, func() any { return uint16(uint16(0)) }).(uint16))))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNPDU factory function for _NPDU
func NewNPDU(protocolVersionNumber uint8, control NPDUControl, destinationNetworkAddress *uint16, destinationLength *uint8, destinationAddress []uint8, sourceNetworkAddress *uint16, sourceLength *uint8, sourceAddress []uint8, hopCount *uint8, nlm NLM, apdu APDU, npduLength uint16) *_NPDU {
	if control == nil {
		panic("control of type NPDUControl for NPDU must not be nil")
	}
	return &_NPDU{ProtocolVersionNumber: protocolVersionNumber, Control: control, DestinationNetworkAddress: destinationNetworkAddress, DestinationLength: destinationLength, DestinationAddress: destinationAddress, SourceNetworkAddress: sourceNetworkAddress, SourceLength: sourceLength, SourceAddress: sourceAddress, HopCount: hopCount, Nlm: nlm, Apdu: apdu, NpduLength: npduLength}
}

// Deprecated: use the interface for direct cast
func CastNPDU(structType any) NPDU {
	if casted, ok := structType.(NPDU); ok {
		return casted
	}
	if casted, ok := structType.(*NPDU); ok {
		return *casted
	}
	return nil
}

func (m *_NPDU) GetTypeName() string {
	return "NPDU"
}

func (m *_NPDU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (protocolVersionNumber)
	lengthInBits += 8

	// Simple field (control)
	lengthInBits += m.Control.GetLengthInBits(ctx)

	// Optional Field (destinationNetworkAddress)
	if m.DestinationNetworkAddress != nil {
		lengthInBits += 16
	}

	// Optional Field (destinationLength)
	if m.DestinationLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (sourceNetworkAddress)
	if m.SourceNetworkAddress != nil {
		lengthInBits += 16
	}

	// Optional Field (sourceLength)
	if m.SourceLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.SourceAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.SourceAddress))
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (hopCount)
	if m.HopCount != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (nlm)
	if m.Nlm != nil {
		lengthInBits += m.Nlm.GetLengthInBits(ctx)
	}

	// Optional Field (apdu)
	if m.Apdu != nil {
		lengthInBits += m.Apdu.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_NPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NPDUParse(ctx context.Context, theBytes []byte, npduLength uint16) (NPDU, error) {
	return NPDUParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), npduLength)
}

func NPDUParseWithBufferProducer(npduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (NPDU, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NPDU, error) {
		return NPDUParseWithBuffer(ctx, readBuffer, npduLength)
	}
}

func NPDUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, npduLength uint16) (NPDU, error) {
	v, err := (&_NPDU{NpduLength: npduLength}).parse(ctx, readBuffer, npduLength)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_NPDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, npduLength uint16) (__nPDU NPDU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersionNumber, err := ReadSimpleField(ctx, "protocolVersionNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersionNumber' field"))
	}
	m.ProtocolVersionNumber = protocolVersionNumber

	control, err := ReadSimpleField[NPDUControl](ctx, "control", ReadComplex[NPDUControl](NPDUControlParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'control' field"))
	}
	m.Control = control

	// Validation
	if !(bool((control) != (nil))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "control required for further processing"})
	}

	var destinationNetworkAddress *uint16
	destinationNetworkAddress, err = ReadOptionalField[uint16](ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)), control.GetDestinationSpecified())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}
	m.DestinationNetworkAddress = destinationNetworkAddress

	var destinationLength *uint8
	destinationLength, err = ReadOptionalField[uint8](ctx, "destinationLength", ReadUnsignedByte(readBuffer, uint8(8)), control.GetDestinationSpecified())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationLength' field"))
	}
	m.DestinationLength = destinationLength

	destinationAddress, err := ReadCountArrayField[uint8](ctx, "destinationAddress", ReadUnsignedByte(readBuffer, uint8(8)), uint64(utils.InlineIf(control.GetDestinationSpecified(), func() any { return int32((*destinationLength)) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationAddress' field"))
	}
	m.DestinationAddress = destinationAddress

	destinationLengthAddon, err := ReadVirtualField[uint16](ctx, "destinationLengthAddon", (*uint16)(nil), utils.InlineIf(control.GetDestinationSpecified(), func() any { return uint16((uint16(uint16(3)) + uint16((*destinationLength)))) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationLengthAddon' field"))
	}
	_ = destinationLengthAddon

	var sourceNetworkAddress *uint16
	sourceNetworkAddress, err = ReadOptionalField[uint16](ctx, "sourceNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)), control.GetSourceSpecified())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceNetworkAddress' field"))
	}
	m.SourceNetworkAddress = sourceNetworkAddress

	var sourceLength *uint8
	sourceLength, err = ReadOptionalField[uint8](ctx, "sourceLength", ReadUnsignedByte(readBuffer, uint8(8)), control.GetSourceSpecified())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceLength' field"))
	}
	m.SourceLength = sourceLength

	sourceAddress, err := ReadCountArrayField[uint8](ctx, "sourceAddress", ReadUnsignedByte(readBuffer, uint8(8)), uint64(utils.InlineIf(control.GetSourceSpecified(), func() any { return int32((*sourceLength)) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAddress' field"))
	}
	m.SourceAddress = sourceAddress

	sourceLengthAddon, err := ReadVirtualField[uint16](ctx, "sourceLengthAddon", (*uint16)(nil), utils.InlineIf(control.GetSourceSpecified(), func() any { return uint16((uint16(uint16(3)) + uint16((*sourceLength)))) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceLengthAddon' field"))
	}
	_ = sourceLengthAddon

	var hopCount *uint8
	hopCount, err = ReadOptionalField[uint8](ctx, "hopCount", ReadUnsignedByte(readBuffer, uint8(8)), control.GetDestinationSpecified())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hopCount' field"))
	}
	m.HopCount = hopCount

	payloadSubtraction, err := ReadVirtualField[uint16](ctx, "payloadSubtraction", (*uint16)(nil), uint16(uint16(2))+uint16((uint16(uint16(sourceLengthAddon)+uint16(destinationLengthAddon))+uint16((utils.InlineIf((control.GetDestinationSpecified()), func() any { return uint16(uint16(1)) }, func() any { return uint16(uint16(0)) }).(uint16))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payloadSubtraction' field"))
	}
	_ = payloadSubtraction

	var nlm NLM
	_nlm, err := ReadOptionalField[NLM](ctx, "nlm", ReadComplex[NLM](NLMParseWithBufferProducer[NLM]((uint16)(uint16(npduLength)-uint16(payloadSubtraction))), readBuffer), control.GetMessageTypeFieldPresent())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nlm' field"))
	}
	if _nlm != nil {
		nlm = *_nlm
		m.Nlm = nlm
	}

	var apdu APDU
	_apdu, err := ReadOptionalField[APDU](ctx, "apdu", ReadComplex[APDU](APDUParseWithBufferProducer[APDU]((uint16)(uint16(npduLength)-uint16(payloadSubtraction))), readBuffer), !(control.GetMessageTypeFieldPresent()))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apdu' field"))
	}
	if _apdu != nil {
		apdu = *_apdu
		m.Apdu = apdu
	}

	// Validation
	if !(bool(bool((nlm) != (nil))) || bool(bool((apdu) != (nil)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "something is wrong here... apdu and nlm not set"})
	}

	if closeErr := readBuffer.CloseContext("NPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NPDU")
	}

	return m, nil
}

func (m *_NPDU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NPDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NPDU")
	}

	if err := WriteSimpleField[uint8](ctx, "protocolVersionNumber", m.GetProtocolVersionNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolVersionNumber' field")
	}

	if err := WriteSimpleField[NPDUControl](ctx, "control", m.GetControl(), WriteComplex[NPDUControl](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'control' field")
	}

	if err := WriteOptionalField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
	}

	if err := WriteOptionalField[uint8](ctx, "destinationLength", m.GetDestinationLength(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationLength' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "destinationAddress", m.GetDestinationAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationAddress' field")
	}
	// Virtual field
	destinationLengthAddon := m.GetDestinationLengthAddon()
	_ = destinationLengthAddon
	if _destinationLengthAddonErr := writeBuffer.WriteVirtual(ctx, "destinationLengthAddon", m.GetDestinationLengthAddon()); _destinationLengthAddonErr != nil {
		return errors.Wrap(_destinationLengthAddonErr, "Error serializing 'destinationLengthAddon' field")
	}

	if err := WriteOptionalField[uint16](ctx, "sourceNetworkAddress", m.GetSourceNetworkAddress(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceNetworkAddress' field")
	}

	if err := WriteOptionalField[uint8](ctx, "sourceLength", m.GetSourceLength(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceLength' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "sourceAddress", m.GetSourceAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceAddress' field")
	}
	// Virtual field
	sourceLengthAddon := m.GetSourceLengthAddon()
	_ = sourceLengthAddon
	if _sourceLengthAddonErr := writeBuffer.WriteVirtual(ctx, "sourceLengthAddon", m.GetSourceLengthAddon()); _sourceLengthAddonErr != nil {
		return errors.Wrap(_sourceLengthAddonErr, "Error serializing 'sourceLengthAddon' field")
	}

	if err := WriteOptionalField[uint8](ctx, "hopCount", m.GetHopCount(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'hopCount' field")
	}
	// Virtual field
	payloadSubtraction := m.GetPayloadSubtraction()
	_ = payloadSubtraction
	if _payloadSubtractionErr := writeBuffer.WriteVirtual(ctx, "payloadSubtraction", m.GetPayloadSubtraction()); _payloadSubtractionErr != nil {
		return errors.Wrap(_payloadSubtractionErr, "Error serializing 'payloadSubtraction' field")
	}

	if err := WriteOptionalField[NLM](ctx, "nlm", GetRef(m.GetNlm()), WriteComplex[NLM](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'nlm' field")
	}

	if err := WriteOptionalField[APDU](ctx, "apdu", GetRef(m.GetApdu()), WriteComplex[APDU](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'apdu' field")
	}

	if popErr := writeBuffer.PopContext("NPDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NPDU")
	}
	return nil
}

////
// Arguments Getter

func (m *_NPDU) GetNpduLength() uint16 {
	return m.NpduLength
}

//
////

func (m *_NPDU) IsNPDU() {}

func (m *_NPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
