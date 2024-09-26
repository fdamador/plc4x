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

// BACnetConstructedDataLandingCalls is the corresponding interface of BACnetConstructedDataLandingCalls
type BACnetConstructedDataLandingCalls interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLandingCallStatus returns LandingCallStatus (property field)
	GetLandingCallStatus() []BACnetLandingCallStatus
	// IsBACnetConstructedDataLandingCalls is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLandingCalls()
}

// _BACnetConstructedDataLandingCalls is the data-structure of this message
type _BACnetConstructedDataLandingCalls struct {
	BACnetConstructedDataContract
	LandingCallStatus []BACnetLandingCallStatus
}

var _ BACnetConstructedDataLandingCalls = (*_BACnetConstructedDataLandingCalls)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLandingCalls)(nil)

// NewBACnetConstructedDataLandingCalls factory function for _BACnetConstructedDataLandingCalls
func NewBACnetConstructedDataLandingCalls(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, landingCallStatus []BACnetLandingCallStatus, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLandingCalls {
	_result := &_BACnetConstructedDataLandingCalls{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LandingCallStatus:             landingCallStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLandingCalls) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLandingCalls) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LANDING_CALLS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLandingCalls) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCalls) GetLandingCallStatus() []BACnetLandingCallStatus {
	return m.LandingCallStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLandingCalls(structType any) BACnetConstructedDataLandingCalls {
	if casted, ok := structType.(BACnetConstructedDataLandingCalls); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLandingCalls); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLandingCalls) GetTypeName() string {
	return "BACnetConstructedDataLandingCalls"
}

func (m *_BACnetConstructedDataLandingCalls) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.LandingCallStatus) > 0 {
		for _, element := range m.LandingCallStatus {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLandingCalls) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLandingCalls) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLandingCalls BACnetConstructedDataLandingCalls, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLandingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLandingCalls")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	landingCallStatus, err := ReadTerminatedArrayField[BACnetLandingCallStatus](ctx, "landingCallStatus", ReadComplex[BACnetLandingCallStatus](BACnetLandingCallStatusParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingCallStatus' field"))
	}
	m.LandingCallStatus = landingCallStatus

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLandingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLandingCalls")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLandingCalls) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLandingCalls) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLandingCalls"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLandingCalls")
		}

		if err := WriteComplexTypeArrayField(ctx, "landingCallStatus", m.GetLandingCallStatus(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'landingCallStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLandingCalls"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLandingCalls")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLandingCalls) IsBACnetConstructedDataLandingCalls() {}

func (m *_BACnetConstructedDataLandingCalls) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLandingCalls) deepCopy() *_BACnetConstructedDataLandingCalls {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLandingCallsCopy := &_BACnetConstructedDataLandingCalls{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLandingCallStatus, BACnetLandingCallStatus](m.LandingCallStatus),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLandingCallsCopy
}

func (m *_BACnetConstructedDataLandingCalls) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
