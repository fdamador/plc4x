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

// BACnetConstructedDataTotalRecordCount is the corresponding interface of BACnetConstructedDataTotalRecordCount
type BACnetConstructedDataTotalRecordCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTotalRecordCount returns TotalRecordCount (property field)
	GetTotalRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTotalRecordCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTotalRecordCount()
}

// _BACnetConstructedDataTotalRecordCount is the data-structure of this message
type _BACnetConstructedDataTotalRecordCount struct {
	BACnetConstructedDataContract
	TotalRecordCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTotalRecordCount = (*_BACnetConstructedDataTotalRecordCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTotalRecordCount)(nil)

// NewBACnetConstructedDataTotalRecordCount factory function for _BACnetConstructedDataTotalRecordCount
func NewBACnetConstructedDataTotalRecordCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, totalRecordCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTotalRecordCount {
	if totalRecordCount == nil {
		panic("totalRecordCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTotalRecordCount must not be nil")
	}
	_result := &_BACnetConstructedDataTotalRecordCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TotalRecordCount:              totalRecordCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTotalRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TOTAL_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetTotalRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.TotalRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTotalRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTotalRecordCount(structType any) BACnetConstructedDataTotalRecordCount {
	if casted, ok := structType.(BACnetConstructedDataTotalRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTotalRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTotalRecordCount) GetTypeName() string {
	return "BACnetConstructedDataTotalRecordCount"
}

func (m *_BACnetConstructedDataTotalRecordCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (totalRecordCount)
	lengthInBits += m.TotalRecordCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTotalRecordCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTotalRecordCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTotalRecordCount BACnetConstructedDataTotalRecordCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTotalRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTotalRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	totalRecordCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "totalRecordCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalRecordCount' field"))
	}
	m.TotalRecordCount = totalRecordCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), totalRecordCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTotalRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTotalRecordCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTotalRecordCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTotalRecordCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTotalRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTotalRecordCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "totalRecordCount", m.GetTotalRecordCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'totalRecordCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTotalRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTotalRecordCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTotalRecordCount) IsBACnetConstructedDataTotalRecordCount() {}

func (m *_BACnetConstructedDataTotalRecordCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTotalRecordCount) deepCopy() *_BACnetConstructedDataTotalRecordCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTotalRecordCountCopy := &_BACnetConstructedDataTotalRecordCount{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.TotalRecordCount.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTotalRecordCountCopy
}

func (m *_BACnetConstructedDataTotalRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
