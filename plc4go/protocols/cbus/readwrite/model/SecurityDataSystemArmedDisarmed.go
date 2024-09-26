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

// SecurityDataSystemArmedDisarmed is the corresponding interface of SecurityDataSystemArmedDisarmed
type SecurityDataSystemArmedDisarmed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetArmCodeType returns ArmCodeType (property field)
	GetArmCodeType() SecurityArmCode
	// IsSecurityDataSystemArmedDisarmed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataSystemArmedDisarmed()
}

// _SecurityDataSystemArmedDisarmed is the data-structure of this message
type _SecurityDataSystemArmedDisarmed struct {
	SecurityDataContract
	ArmCodeType SecurityArmCode
}

var _ SecurityDataSystemArmedDisarmed = (*_SecurityDataSystemArmedDisarmed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataSystemArmedDisarmed)(nil)

// NewSecurityDataSystemArmedDisarmed factory function for _SecurityDataSystemArmedDisarmed
func NewSecurityDataSystemArmedDisarmed(commandTypeContainer SecurityCommandTypeContainer, argument byte, armCodeType SecurityArmCode) *_SecurityDataSystemArmedDisarmed {
	if armCodeType == nil {
		panic("armCodeType of type SecurityArmCode for SecurityDataSystemArmedDisarmed must not be nil")
	}
	_result := &_SecurityDataSystemArmedDisarmed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ArmCodeType:          armCodeType,
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

func (m *_SecurityDataSystemArmedDisarmed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataSystemArmedDisarmed) GetArmCodeType() SecurityArmCode {
	return m.ArmCodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataSystemArmedDisarmed(structType any) SecurityDataSystemArmedDisarmed {
	if casted, ok := structType.(SecurityDataSystemArmedDisarmed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataSystemArmedDisarmed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataSystemArmedDisarmed) GetTypeName() string {
	return "SecurityDataSystemArmedDisarmed"
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (armCodeType)
	lengthInBits += m.ArmCodeType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataSystemArmedDisarmed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataSystemArmedDisarmed SecurityDataSystemArmedDisarmed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataSystemArmedDisarmed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataSystemArmedDisarmed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	armCodeType, err := ReadSimpleField[SecurityArmCode](ctx, "armCodeType", ReadComplex[SecurityArmCode](SecurityArmCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'armCodeType' field"))
	}
	m.ArmCodeType = armCodeType

	if closeErr := readBuffer.CloseContext("SecurityDataSystemArmedDisarmed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataSystemArmedDisarmed")
	}

	return m, nil
}

func (m *_SecurityDataSystemArmedDisarmed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataSystemArmedDisarmed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataSystemArmedDisarmed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataSystemArmedDisarmed")
		}

		if err := WriteSimpleField[SecurityArmCode](ctx, "armCodeType", m.GetArmCodeType(), WriteComplex[SecurityArmCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'armCodeType' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataSystemArmedDisarmed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataSystemArmedDisarmed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataSystemArmedDisarmed) IsSecurityDataSystemArmedDisarmed() {}

func (m *_SecurityDataSystemArmedDisarmed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataSystemArmedDisarmed) deepCopy() *_SecurityDataSystemArmedDisarmed {
	if m == nil {
		return nil
	}
	_SecurityDataSystemArmedDisarmedCopy := &_SecurityDataSystemArmedDisarmed{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.ArmCodeType.DeepCopy().(SecurityArmCode),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataSystemArmedDisarmedCopy
}

func (m *_SecurityDataSystemArmedDisarmed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
