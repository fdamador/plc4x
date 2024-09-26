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

// AdsReadWriteRequest is the corresponding interface of AdsReadWriteRequest
type AdsReadWriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetReadLength returns ReadLength (property field)
	GetReadLength() uint32
	// GetItems returns Items (property field)
	GetItems() []AdsMultiRequestItem
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsReadWriteRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadWriteRequest()
}

// _AdsReadWriteRequest is the data-structure of this message
type _AdsReadWriteRequest struct {
	AmsPacketContract
	IndexGroup  uint32
	IndexOffset uint32
	ReadLength  uint32
	Items       []AdsMultiRequestItem
	Data        []byte
}

var _ AdsReadWriteRequest = (*_AdsReadWriteRequest)(nil)
var _ AmsPacketRequirements = (*_AdsReadWriteRequest)(nil)

// NewAdsReadWriteRequest factory function for _AdsReadWriteRequest
func NewAdsReadWriteRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, indexGroup uint32, indexOffset uint32, readLength uint32, items []AdsMultiRequestItem, data []byte) *_AdsReadWriteRequest {
	_result := &_AdsReadWriteRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		IndexGroup:        indexGroup,
		IndexOffset:       indexOffset,
		ReadLength:        readLength,
		Items:             items,
		Data:              data,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadWriteRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *_AdsReadWriteRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadWriteRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadWriteRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsReadWriteRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsReadWriteRequest) GetReadLength() uint32 {
	return m.ReadLength
}

func (m *_AdsReadWriteRequest) GetItems() []AdsMultiRequestItem {
	return m.Items
}

func (m *_AdsReadWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsReadWriteRequest(structType any) AdsReadWriteRequest {
	if casted, ok := structType.(AdsReadWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadWriteRequest) GetTypeName() string {
	return "AdsReadWriteRequest"
}

func (m *_AdsReadWriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (readLength)
	lengthInBits += 32

	// Implicit Field (writeLength)
	lengthInBits += 32

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadWriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadWriteRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadWriteRequest AdsReadWriteRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	indexGroup, err := ReadSimpleField(ctx, "indexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexGroup' field"))
	}
	m.IndexGroup = indexGroup

	indexOffset, err := ReadSimpleField(ctx, "indexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexOffset' field"))
	}
	m.IndexOffset = indexOffset

	readLength, err := ReadSimpleField(ctx, "readLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readLength' field"))
	}
	m.ReadLength = readLength

	writeLength, err := ReadImplicitField[uint32](ctx, "writeLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeLength' field"))
	}
	_ = writeLength

	items, err := ReadCountArrayField[AdsMultiRequestItem](ctx, "items", ReadComplex[AdsMultiRequestItem](AdsMultiRequestItemParseWithBufferProducer[AdsMultiRequestItem]((uint32)(indexGroup)), readBuffer), uint64(utils.InlineIf((bool(bool((bool((indexGroup) == (61568)))) || bool((bool((indexGroup) == (61569))))) || bool((bool((indexGroup) == (61570))))), func() any { return int32(indexOffset) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	data, err := readBuffer.ReadByteArray("data", int(int32(writeLength)-int32((int32(int32(len(items)))*int32(int32(12))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsReadWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadWriteRequest")
	}

	return m, nil
}

func (m *_AdsReadWriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadWriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadWriteRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "indexGroup", m.GetIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "indexOffset", m.GetIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "readLength", m.GetReadLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'readLength' field")
		}
		writeLength := uint32(uint32((uint32(uint32(len(m.GetItems()))) * uint32((utils.InlineIf((bool((m.GetIndexGroup()) == (61570))), func() any { return uint32(uint32(16)) }, func() any { return uint32(uint32(12)) }).(uint32))))) + uint32(uint32(len(m.GetData()))))
		if err := WriteImplicitField(ctx, "writeLength", writeLength, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadWriteRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadWriteRequest) IsAdsReadWriteRequest() {}

func (m *_AdsReadWriteRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsReadWriteRequest) deepCopy() *_AdsReadWriteRequest {
	if m == nil {
		return nil
	}
	_AdsReadWriteRequestCopy := &_AdsReadWriteRequest{
		m.AmsPacketContract.DeepCopy().(AmsPacketContract),
		m.IndexGroup,
		m.IndexOffset,
		m.ReadLength,
		utils.DeepCopySlice[AdsMultiRequestItem, AdsMultiRequestItem](m.Items),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsReadWriteRequestCopy
}

func (m *_AdsReadWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
