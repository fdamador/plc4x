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

// AdsDeviceNotificationRequest is the corresponding interface of AdsDeviceNotificationRequest
type AdsDeviceNotificationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetLength returns Length (property field)
	GetLength() uint32
	// GetStamps returns Stamps (property field)
	GetStamps() uint32
	// GetAdsStampHeaders returns AdsStampHeaders (property field)
	GetAdsStampHeaders() []AdsStampHeader
	// IsAdsDeviceNotificationRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDeviceNotificationRequest()
}

// _AdsDeviceNotificationRequest is the data-structure of this message
type _AdsDeviceNotificationRequest struct {
	AmsPacketContract
	Length          uint32
	Stamps          uint32
	AdsStampHeaders []AdsStampHeader
}

var _ AdsDeviceNotificationRequest = (*_AdsDeviceNotificationRequest)(nil)
var _ AmsPacketRequirements = (*_AdsDeviceNotificationRequest)(nil)

// NewAdsDeviceNotificationRequest factory function for _AdsDeviceNotificationRequest
func NewAdsDeviceNotificationRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, length uint32, stamps uint32, adsStampHeaders []AdsStampHeader) *_AdsDeviceNotificationRequest {
	_result := &_AdsDeviceNotificationRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Length:            length,
		Stamps:            stamps,
		AdsStampHeaders:   adsStampHeaders,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_DEVICE_NOTIFICATION
}

func (m *_AdsDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeviceNotificationRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeviceNotificationRequest) GetLength() uint32 {
	return m.Length
}

func (m *_AdsDeviceNotificationRequest) GetStamps() uint32 {
	return m.Stamps
}

func (m *_AdsDeviceNotificationRequest) GetAdsStampHeaders() []AdsStampHeader {
	return m.AdsStampHeaders
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDeviceNotificationRequest(structType any) AdsDeviceNotificationRequest {
	if casted, ok := structType.(AdsDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeviceNotificationRequest) GetTypeName() string {
	return "AdsDeviceNotificationRequest"
}

func (m *_AdsDeviceNotificationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (length)
	lengthInBits += 32

	// Simple field (stamps)
	lengthInBits += 32

	// Array field
	if len(m.AdsStampHeaders) > 0 {
		for _curItem, element := range m.AdsStampHeaders {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AdsStampHeaders), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdsDeviceNotificationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDeviceNotificationRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsDeviceNotificationRequest AdsDeviceNotificationRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	length, err := ReadSimpleField(ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	m.Length = length

	stamps, err := ReadSimpleField(ctx, "stamps", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stamps' field"))
	}
	m.Stamps = stamps

	adsStampHeaders, err := ReadCountArrayField[AdsStampHeader](ctx, "adsStampHeaders", ReadComplex[AdsStampHeader](AdsStampHeaderParseWithBuffer, readBuffer), uint64(stamps))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsStampHeaders' field"))
	}
	m.AdsStampHeaders = adsStampHeaders

	if closeErr := readBuffer.CloseContext("AdsDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeviceNotificationRequest")
	}

	return m, nil
}

func (m *_AdsDeviceNotificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeviceNotificationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeviceNotificationRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "length", m.GetLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleField[uint32](ctx, "stamps", m.GetStamps(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'stamps' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "adsStampHeaders", m.GetAdsStampHeaders(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'adsStampHeaders' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeviceNotificationRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeviceNotificationRequest) IsAdsDeviceNotificationRequest() {}

func (m *_AdsDeviceNotificationRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDeviceNotificationRequest) deepCopy() *_AdsDeviceNotificationRequest {
	if m == nil {
		return nil
	}
	_AdsDeviceNotificationRequestCopy := &_AdsDeviceNotificationRequest{
		m.AmsPacketContract.DeepCopy().(AmsPacketContract),
		m.Length,
		m.Stamps,
		utils.DeepCopySlice[AdsStampHeader, AdsStampHeader](m.AdsStampHeaders),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsDeviceNotificationRequestCopy
}

func (m *_AdsDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
