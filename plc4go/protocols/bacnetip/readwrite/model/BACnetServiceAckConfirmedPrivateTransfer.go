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

// BACnetServiceAckConfirmedPrivateTransfer is the corresponding interface of BACnetServiceAckConfirmedPrivateTransfer
type BACnetServiceAckConfirmedPrivateTransfer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetResultBlock returns ResultBlock (property field)
	GetResultBlock() BACnetConstructedData
	// IsBACnetServiceAckConfirmedPrivateTransfer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckConfirmedPrivateTransfer()
}

// _BACnetServiceAckConfirmedPrivateTransfer is the data-structure of this message
type _BACnetServiceAckConfirmedPrivateTransfer struct {
	BACnetServiceAckContract
	VendorId      BACnetVendorIdTagged
	ServiceNumber BACnetContextTagUnsignedInteger
	ResultBlock   BACnetConstructedData
}

var _ BACnetServiceAckConfirmedPrivateTransfer = (*_BACnetServiceAckConfirmedPrivateTransfer)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckConfirmedPrivateTransfer)(nil)

// NewBACnetServiceAckConfirmedPrivateTransfer factory function for _BACnetServiceAckConfirmedPrivateTransfer
func NewBACnetServiceAckConfirmedPrivateTransfer(vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, resultBlock BACnetConstructedData, serviceAckLength uint32) *_BACnetServiceAckConfirmedPrivateTransfer {
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetServiceAckConfirmedPrivateTransfer must not be nil")
	}
	if serviceNumber == nil {
		panic("serviceNumber of type BACnetContextTagUnsignedInteger for BACnetServiceAckConfirmedPrivateTransfer must not be nil")
	}
	_result := &_BACnetServiceAckConfirmedPrivateTransfer{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		VendorId:                 vendorId,
		ServiceNumber:            serviceNumber,
		ResultBlock:              resultBlock,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetResultBlock() BACnetConstructedData {
	return m.ResultBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckConfirmedPrivateTransfer(structType any) BACnetServiceAckConfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetServiceAckConfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckConfirmedPrivateTransfer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetServiceAckConfirmedPrivateTransfer"
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits(ctx)

	// Optional Field (resultBlock)
	if m.ResultBlock != nil {
		lengthInBits += m.ResultBlock.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckConfirmedPrivateTransfer BACnetServiceAckConfirmedPrivateTransfer, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckConfirmedPrivateTransfer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	serviceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNumber' field"))
	}
	m.ServiceNumber = serviceNumber

	var resultBlock BACnetConstructedData
	_resultBlock, err := ReadOptionalField[BACnetConstructedData](ctx, "resultBlock", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(2)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resultBlock' field"))
	}
	if _resultBlock != nil {
		resultBlock = *_resultBlock
		m.ResultBlock = resultBlock
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckConfirmedPrivateTransfer")
	}

	return m, nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckConfirmedPrivateTransfer")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", m.GetServiceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceNumber' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "resultBlock", GetRef(m.GetResultBlock()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'resultBlock' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckConfirmedPrivateTransfer")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) IsBACnetServiceAckConfirmedPrivateTransfer() {}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) deepCopy() *_BACnetServiceAckConfirmedPrivateTransfer {
	if m == nil {
		return nil
	}
	_BACnetServiceAckConfirmedPrivateTransferCopy := &_BACnetServiceAckConfirmedPrivateTransfer{
		m.BACnetServiceAckContract.DeepCopy().(BACnetServiceAckContract),
		m.VendorId.DeepCopy().(BACnetVendorIdTagged),
		m.ServiceNumber.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ResultBlock.DeepCopy().(BACnetConstructedData),
	}
	m.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckConfirmedPrivateTransferCopy
}

func (m *_BACnetServiceAckConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
