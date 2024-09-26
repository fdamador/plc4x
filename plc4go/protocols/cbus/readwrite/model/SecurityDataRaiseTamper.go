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

// SecurityDataRaiseTamper is the corresponding interface of SecurityDataRaiseTamper
type SecurityDataRaiseTamper interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataRaiseTamper is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataRaiseTamper()
}

// _SecurityDataRaiseTamper is the data-structure of this message
type _SecurityDataRaiseTamper struct {
	SecurityDataContract
}

var _ SecurityDataRaiseTamper = (*_SecurityDataRaiseTamper)(nil)
var _ SecurityDataRequirements = (*_SecurityDataRaiseTamper)(nil)

// NewSecurityDataRaiseTamper factory function for _SecurityDataRaiseTamper
func NewSecurityDataRaiseTamper(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataRaiseTamper {
	_result := &_SecurityDataRaiseTamper{
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

func (m *_SecurityDataRaiseTamper) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataRaiseTamper(structType any) SecurityDataRaiseTamper {
	if casted, ok := structType.(SecurityDataRaiseTamper); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataRaiseTamper); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataRaiseTamper) GetTypeName() string {
	return "SecurityDataRaiseTamper"
}

func (m *_SecurityDataRaiseTamper) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataRaiseTamper) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataRaiseTamper) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataRaiseTamper SecurityDataRaiseTamper, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataRaiseTamper"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataRaiseTamper")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataRaiseTamper"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataRaiseTamper")
	}

	return m, nil
}

func (m *_SecurityDataRaiseTamper) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataRaiseTamper) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataRaiseTamper"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataRaiseTamper")
		}

		if popErr := writeBuffer.PopContext("SecurityDataRaiseTamper"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataRaiseTamper")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataRaiseTamper) IsSecurityDataRaiseTamper() {}

func (m *_SecurityDataRaiseTamper) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataRaiseTamper) deepCopy() *_SecurityDataRaiseTamper {
	if m == nil {
		return nil
	}
	_SecurityDataRaiseTamperCopy := &_SecurityDataRaiseTamper{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataRaiseTamperCopy
}

func (m *_SecurityDataRaiseTamper) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
