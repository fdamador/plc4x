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

// BACnetServiceAckReadPropertyMultiple is the corresponding interface of BACnetServiceAckReadPropertyMultiple
type BACnetServiceAckReadPropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetData returns Data (property field)
	GetData() []BACnetReadAccessResult
	// IsBACnetServiceAckReadPropertyMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckReadPropertyMultiple()
}

// _BACnetServiceAckReadPropertyMultiple is the data-structure of this message
type _BACnetServiceAckReadPropertyMultiple struct {
	BACnetServiceAckContract
	Data []BACnetReadAccessResult

	// Arguments.
	ServiceAckPayloadLength uint32
}

var _ BACnetServiceAckReadPropertyMultiple = (*_BACnetServiceAckReadPropertyMultiple)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckReadPropertyMultiple)(nil)

// NewBACnetServiceAckReadPropertyMultiple factory function for _BACnetServiceAckReadPropertyMultiple
func NewBACnetServiceAckReadPropertyMultiple(data []BACnetReadAccessResult, serviceAckPayloadLength uint32, serviceAckLength uint32) *_BACnetServiceAckReadPropertyMultiple {
	_result := &_BACnetServiceAckReadPropertyMultiple{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		Data:                     data,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadPropertyMultiple) GetData() []BACnetReadAccessResult {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadPropertyMultiple(structType any) BACnetServiceAckReadPropertyMultiple {
	if casted, ok := structType.(BACnetServiceAckReadPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetTypeName() string {
	return "BACnetServiceAckReadPropertyMultiple"
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadPropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckReadPropertyMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckPayloadLength uint32, serviceAckLength uint32) (__bACnetServiceAckReadPropertyMultiple BACnetServiceAckReadPropertyMultiple, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadLengthArrayField[BACnetReadAccessResult](ctx, "data", ReadComplex[BACnetReadAccessResult](BACnetReadAccessResultParseWithBuffer, readBuffer), int(serviceAckPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadPropertyMultiple")
	}

	return m, nil
}

func (m *_BACnetServiceAckReadPropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckReadPropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadPropertyMultiple")
		}

		if err := WriteComplexTypeArrayField(ctx, "data", m.GetData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadPropertyMultiple")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetServiceAckReadPropertyMultiple) GetServiceAckPayloadLength() uint32 {
	return m.ServiceAckPayloadLength
}

//
////

func (m *_BACnetServiceAckReadPropertyMultiple) IsBACnetServiceAckReadPropertyMultiple() {}

func (m *_BACnetServiceAckReadPropertyMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckReadPropertyMultiple) deepCopy() *_BACnetServiceAckReadPropertyMultiple {
	if m == nil {
		return nil
	}
	_BACnetServiceAckReadPropertyMultipleCopy := &_BACnetServiceAckReadPropertyMultiple{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopySlice[BACnetReadAccessResult, BACnetReadAccessResult](m.Data),
		m.ServiceAckPayloadLength,
	}
	m.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckReadPropertyMultipleCopy
}

func (m *_BACnetServiceAckReadPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
