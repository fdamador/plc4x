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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataReserved is the corresponding interface of SALDataReserved
type SALDataReserved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// IsSALDataReserved is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataReserved()
}

// _SALDataReserved is the data-structure of this message
type _SALDataReserved struct {
	SALDataContract
}

var _ SALDataReserved = (*_SALDataReserved)(nil)
var _ SALDataRequirements = (*_SALDataReserved)(nil)

// NewSALDataReserved factory function for _SALDataReserved
func NewSALDataReserved(salData SALData) *_SALDataReserved {
	_result := &_SALDataReserved{
		SALDataContract: NewSALData(salData),
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataReserved) GetApplicationId() ApplicationId {
	return ApplicationId_RESERVED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataReserved) GetParent() SALDataContract {
	return m.SALDataContract
}

// Deprecated: use the interface for direct cast
func CastSALDataReserved(structType any) SALDataReserved {
	if casted, ok := structType.(SALDataReserved); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataReserved); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataReserved) GetTypeName() string {
	return "SALDataReserved"
}

func (m *_SALDataReserved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SALDataReserved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataReserved) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataReserved SALDataReserved, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "RESERVED Not yet implemented"})
	}

	if closeErr := readBuffer.CloseContext("SALDataReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataReserved")
	}

	return m, nil
}

func (m *_SALDataReserved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataReserved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataReserved")
		}

		if popErr := writeBuffer.PopContext("SALDataReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataReserved")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataReserved) IsSALDataReserved() {}

func (m *_SALDataReserved) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataReserved) deepCopy() *_SALDataReserved {
	if m == nil {
		return nil
	}
	_SALDataReservedCopy := &_SALDataReserved{
		m.SALDataContract.DeepCopy().(SALDataContract),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataReservedCopy
}

func (m *_SALDataReserved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
