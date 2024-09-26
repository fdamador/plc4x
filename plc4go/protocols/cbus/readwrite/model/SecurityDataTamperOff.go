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

// SecurityDataTamperOff is the corresponding interface of SecurityDataTamperOff
type SecurityDataTamperOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataTamperOff is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataTamperOff()
}

// _SecurityDataTamperOff is the data-structure of this message
type _SecurityDataTamperOff struct {
	SecurityDataContract
}

var _ SecurityDataTamperOff = (*_SecurityDataTamperOff)(nil)
var _ SecurityDataRequirements = (*_SecurityDataTamperOff)(nil)

// NewSecurityDataTamperOff factory function for _SecurityDataTamperOff
func NewSecurityDataTamperOff(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataTamperOff {
	_result := &_SecurityDataTamperOff{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
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

func (m *_SecurityDataTamperOff) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataTamperOff(structType any) SecurityDataTamperOff {
	if casted, ok := structType.(SecurityDataTamperOff); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataTamperOff); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataTamperOff) GetTypeName() string {
	return "SecurityDataTamperOff"
}

func (m *_SecurityDataTamperOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataTamperOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataTamperOff) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataTamperOff SecurityDataTamperOff, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataTamperOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataTamperOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataTamperOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataTamperOff")
	}

	return m, nil
}

func (m *_SecurityDataTamperOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataTamperOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataTamperOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataTamperOff")
		}

		if popErr := writeBuffer.PopContext("SecurityDataTamperOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataTamperOff")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataTamperOff) IsSecurityDataTamperOff() {}

func (m *_SecurityDataTamperOff) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataTamperOff) deepCopy() *_SecurityDataTamperOff {
	if m == nil {
		return nil
	}
	_SecurityDataTamperOffCopy := &_SecurityDataTamperOff{
		m.SecurityDataContract.DeepCopy().(SecurityDataContract),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataTamperOffCopy
}

func (m *_SecurityDataTamperOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
