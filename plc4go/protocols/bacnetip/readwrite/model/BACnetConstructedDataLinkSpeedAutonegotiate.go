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

// BACnetConstructedDataLinkSpeedAutonegotiate is the corresponding interface of BACnetConstructedDataLinkSpeedAutonegotiate
type BACnetConstructedDataLinkSpeedAutonegotiate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLinkSpeedAutonegotiate returns LinkSpeedAutonegotiate (property field)
	GetLinkSpeedAutonegotiate() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataLinkSpeedAutonegotiate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLinkSpeedAutonegotiate()
}

// _BACnetConstructedDataLinkSpeedAutonegotiate is the data-structure of this message
type _BACnetConstructedDataLinkSpeedAutonegotiate struct {
	BACnetConstructedDataContract
	LinkSpeedAutonegotiate BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataLinkSpeedAutonegotiate = (*_BACnetConstructedDataLinkSpeedAutonegotiate)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLinkSpeedAutonegotiate)(nil)

// NewBACnetConstructedDataLinkSpeedAutonegotiate factory function for _BACnetConstructedDataLinkSpeedAutonegotiate
func NewBACnetConstructedDataLinkSpeedAutonegotiate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, linkSpeedAutonegotiate BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLinkSpeedAutonegotiate {
	if linkSpeedAutonegotiate == nil {
		panic("linkSpeedAutonegotiate of type BACnetApplicationTagBoolean for BACnetConstructedDataLinkSpeedAutonegotiate must not be nil")
	}
	_result := &_BACnetConstructedDataLinkSpeedAutonegotiate{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LinkSpeedAutonegotiate:        linkSpeedAutonegotiate,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEED_AUTONEGOTIATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLinkSpeedAutonegotiate() BACnetApplicationTagBoolean {
	return m.LinkSpeedAutonegotiate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetLinkSpeedAutonegotiate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLinkSpeedAutonegotiate(structType any) BACnetConstructedDataLinkSpeedAutonegotiate {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeedAutonegotiate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeedAutonegotiate"
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (linkSpeedAutonegotiate)
	lengthInBits += m.LinkSpeedAutonegotiate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLinkSpeedAutonegotiate BACnetConstructedDataLinkSpeedAutonegotiate, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeedAutonegotiate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	linkSpeedAutonegotiate, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "linkSpeedAutonegotiate", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkSpeedAutonegotiate' field"))
	}
	m.LinkSpeedAutonegotiate = linkSpeedAutonegotiate

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), linkSpeedAutonegotiate)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeedAutonegotiate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeedAutonegotiate")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeedAutonegotiate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeedAutonegotiate")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "linkSpeedAutonegotiate", m.GetLinkSpeedAutonegotiate(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkSpeedAutonegotiate' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeedAutonegotiate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeedAutonegotiate")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) IsBACnetConstructedDataLinkSpeedAutonegotiate() {
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) deepCopy() *_BACnetConstructedDataLinkSpeedAutonegotiate {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLinkSpeedAutonegotiateCopy := &_BACnetConstructedDataLinkSpeedAutonegotiate{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.LinkSpeedAutonegotiate.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLinkSpeedAutonegotiateCopy
}

func (m *_BACnetConstructedDataLinkSpeedAutonegotiate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
