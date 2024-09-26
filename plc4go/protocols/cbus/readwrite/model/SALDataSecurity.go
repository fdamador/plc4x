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

// SALDataSecurity is the corresponding interface of SALDataSecurity
type SALDataSecurity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetSecurityData returns SecurityData (property field)
	GetSecurityData() SecurityData
	// IsSALDataSecurity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataSecurity()
}

// _SALDataSecurity is the data-structure of this message
type _SALDataSecurity struct {
	SALDataContract
	SecurityData SecurityData
}

var _ SALDataSecurity = (*_SALDataSecurity)(nil)
var _ SALDataRequirements = (*_SALDataSecurity)(nil)

// NewSALDataSecurity factory function for _SALDataSecurity
func NewSALDataSecurity(salData SALData, securityData SecurityData) *_SALDataSecurity {
	if securityData == nil {
		panic("securityData of type SecurityData for SALDataSecurity must not be nil")
	}
	_result := &_SALDataSecurity{
		SALDataContract: NewSALData(salData),
		SecurityData:    securityData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataSecurity) GetApplicationId() ApplicationId {
	return ApplicationId_SECURITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataSecurity) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataSecurity) GetSecurityData() SecurityData {
	return m.SecurityData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataSecurity(structType any) SALDataSecurity {
	if casted, ok := structType.(SALDataSecurity); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataSecurity); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataSecurity) GetTypeName() string {
	return "SALDataSecurity"
}

func (m *_SALDataSecurity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (securityData)
	lengthInBits += m.SecurityData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataSecurity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataSecurity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataSecurity SALDataSecurity, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataSecurity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataSecurity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityData, err := ReadSimpleField[SecurityData](ctx, "securityData", ReadComplex[SecurityData](SecurityDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityData' field"))
	}
	m.SecurityData = securityData

	if closeErr := readBuffer.CloseContext("SALDataSecurity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataSecurity")
	}

	return m, nil
}

func (m *_SALDataSecurity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataSecurity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataSecurity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataSecurity")
		}

		if err := WriteSimpleField[SecurityData](ctx, "securityData", m.GetSecurityData(), WriteComplex[SecurityData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataSecurity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataSecurity")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataSecurity) IsSALDataSecurity() {}

func (m *_SALDataSecurity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataSecurity) deepCopy() *_SALDataSecurity {
	if m == nil {
		return nil
	}
	_SALDataSecurityCopy := &_SALDataSecurity{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.SecurityData.DeepCopy().(SecurityData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataSecurityCopy
}

func (m *_SALDataSecurity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
