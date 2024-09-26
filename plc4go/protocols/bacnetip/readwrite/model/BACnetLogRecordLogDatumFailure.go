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

// BACnetLogRecordLogDatumFailure is the corresponding interface of BACnetLogRecordLogDatumFailure
type BACnetLogRecordLogDatumFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetFailure returns Failure (property field)
	GetFailure() ErrorEnclosed
	// IsBACnetLogRecordLogDatumFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumFailure()
}

// _BACnetLogRecordLogDatumFailure is the data-structure of this message
type _BACnetLogRecordLogDatumFailure struct {
	BACnetLogRecordLogDatumContract
	Failure ErrorEnclosed
}

var _ BACnetLogRecordLogDatumFailure = (*_BACnetLogRecordLogDatumFailure)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumFailure)(nil)

// NewBACnetLogRecordLogDatumFailure factory function for _BACnetLogRecordLogDatumFailure
func NewBACnetLogRecordLogDatumFailure(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, failure ErrorEnclosed, tagNumber uint8) *_BACnetLogRecordLogDatumFailure {
	if failure == nil {
		panic("failure of type ErrorEnclosed for BACnetLogRecordLogDatumFailure must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumFailure{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		Failure:                         failure,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumFailure) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumFailure) GetFailure() ErrorEnclosed {
	return m.Failure
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumFailure(structType any) BACnetLogRecordLogDatumFailure {
	if casted, ok := structType.(BACnetLogRecordLogDatumFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumFailure) GetTypeName() string {
	return "BACnetLogRecordLogDatumFailure"
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (failure)
	lengthInBits += m.Failure.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumFailure BACnetLogRecordLogDatumFailure, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	failure, err := ReadSimpleField[ErrorEnclosed](ctx, "failure", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'failure' field"))
	}
	m.Failure = failure

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumFailure")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumFailure")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "failure", m.GetFailure(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'failure' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumFailure")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumFailure) IsBACnetLogRecordLogDatumFailure() {}

func (m *_BACnetLogRecordLogDatumFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumFailure) deepCopy() *_BACnetLogRecordLogDatumFailure {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumFailureCopy := &_BACnetLogRecordLogDatumFailure{
		m.BACnetLogRecordLogDatumContract.DeepCopy().(BACnetLogRecordLogDatumContract),
		m.Failure.DeepCopy().(ErrorEnclosed),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumFailureCopy
}

func (m *_BACnetLogRecordLogDatumFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
