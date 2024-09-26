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

// ConfirmedPrivateTransferError is the corresponding interface of ConfirmedPrivateTransferError
type ConfirmedPrivateTransferError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetErrorParameters returns ErrorParameters (property field)
	GetErrorParameters() BACnetConstructedData
	// IsConfirmedPrivateTransferError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConfirmedPrivateTransferError()
}

// _ConfirmedPrivateTransferError is the data-structure of this message
type _ConfirmedPrivateTransferError struct {
	BACnetErrorContract
	ErrorType       ErrorEnclosed
	VendorId        BACnetVendorIdTagged
	ServiceNumber   BACnetContextTagUnsignedInteger
	ErrorParameters BACnetConstructedData
}

var _ ConfirmedPrivateTransferError = (*_ConfirmedPrivateTransferError)(nil)
var _ BACnetErrorRequirements = (*_ConfirmedPrivateTransferError)(nil)

// NewConfirmedPrivateTransferError factory function for _ConfirmedPrivateTransferError
func NewConfirmedPrivateTransferError(errorType ErrorEnclosed, vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, errorParameters BACnetConstructedData) *_ConfirmedPrivateTransferError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for ConfirmedPrivateTransferError must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for ConfirmedPrivateTransferError must not be nil")
	}
	if serviceNumber == nil {
		panic("serviceNumber of type BACnetContextTagUnsignedInteger for ConfirmedPrivateTransferError must not be nil")
	}
	_result := &_ConfirmedPrivateTransferError{
		BACnetErrorContract: NewBACnetError(),
		ErrorType:           errorType,
		VendorId:            vendorId,
		ServiceNumber:       serviceNumber,
		ErrorParameters:     errorParameters,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConfirmedPrivateTransferError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConfirmedPrivateTransferError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConfirmedPrivateTransferError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ConfirmedPrivateTransferError) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_ConfirmedPrivateTransferError) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_ConfirmedPrivateTransferError) GetErrorParameters() BACnetConstructedData {
	return m.ErrorParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConfirmedPrivateTransferError(structType any) ConfirmedPrivateTransferError {
	if casted, ok := structType.(ConfirmedPrivateTransferError); ok {
		return casted
	}
	if casted, ok := structType.(*ConfirmedPrivateTransferError); ok {
		return *casted
	}
	return nil
}

func (m *_ConfirmedPrivateTransferError) GetTypeName() string {
	return "ConfirmedPrivateTransferError"
}

func (m *_ConfirmedPrivateTransferError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits(ctx)

	// Optional Field (errorParameters)
	if m.ErrorParameters != nil {
		lengthInBits += m.ErrorParameters.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_ConfirmedPrivateTransferError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConfirmedPrivateTransferError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__confirmedPrivateTransferError ConfirmedPrivateTransferError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConfirmedPrivateTransferError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConfirmedPrivateTransferError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	serviceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNumber' field"))
	}
	m.ServiceNumber = serviceNumber

	var errorParameters BACnetConstructedData
	_errorParameters, err := ReadOptionalField[BACnetConstructedData](ctx, "errorParameters", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(3)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorParameters' field"))
	}
	if _errorParameters != nil {
		errorParameters = *_errorParameters
		m.ErrorParameters = errorParameters
	}

	if closeErr := readBuffer.CloseContext("ConfirmedPrivateTransferError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConfirmedPrivateTransferError")
	}

	return m, nil
}

func (m *_ConfirmedPrivateTransferError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConfirmedPrivateTransferError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConfirmedPrivateTransferError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConfirmedPrivateTransferError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", m.GetServiceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceNumber' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "errorParameters", GetRef(m.GetErrorParameters()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'errorParameters' field")
		}

		if popErr := writeBuffer.PopContext("ConfirmedPrivateTransferError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConfirmedPrivateTransferError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConfirmedPrivateTransferError) IsConfirmedPrivateTransferError() {}

func (m *_ConfirmedPrivateTransferError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConfirmedPrivateTransferError) deepCopy() *_ConfirmedPrivateTransferError {
	if m == nil {
		return nil
	}
	_ConfirmedPrivateTransferErrorCopy := &_ConfirmedPrivateTransferError{
		m.BACnetErrorContract.DeepCopy().(BACnetErrorContract),
		m.ErrorType.DeepCopy().(ErrorEnclosed),
		m.VendorId.DeepCopy().(BACnetVendorIdTagged),
		m.ServiceNumber.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ErrorParameters.DeepCopy().(BACnetConstructedData),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _ConfirmedPrivateTransferErrorCopy
}

func (m *_ConfirmedPrivateTransferError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
