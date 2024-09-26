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

// AdsDiscoveryBlockFingerprint is the corresponding interface of AdsDiscoveryBlockFingerprint
type AdsDiscoveryBlockFingerprint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsDiscoveryBlockFingerprint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockFingerprint()
}

// _AdsDiscoveryBlockFingerprint is the data-structure of this message
type _AdsDiscoveryBlockFingerprint struct {
	AdsDiscoveryBlockContract
	Data []byte
}

var _ AdsDiscoveryBlockFingerprint = (*_AdsDiscoveryBlockFingerprint)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockFingerprint)(nil)

// NewAdsDiscoveryBlockFingerprint factory function for _AdsDiscoveryBlockFingerprint
func NewAdsDiscoveryBlockFingerprint(data []byte) *_AdsDiscoveryBlockFingerprint {
	_result := &_AdsDiscoveryBlockFingerprint{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		Data:                      data,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockFingerprint) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_FINGERPRINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockFingerprint) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockFingerprint) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockFingerprint(structType any) AdsDiscoveryBlockFingerprint {
	if casted, ok := structType.(AdsDiscoveryBlockFingerprint); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockFingerprint); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockFingerprint) GetTypeName() string {
	return "AdsDiscoveryBlockFingerprint"
}

func (m *_AdsDiscoveryBlockFingerprint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Implicit Field (dataLen)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsDiscoveryBlockFingerprint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockFingerprint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockFingerprint AdsDiscoveryBlockFingerprint, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockFingerprint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockFingerprint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataLen, err := ReadImplicitField[uint16](ctx, "dataLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataLen' field"))
	}
	_ = dataLen

	data, err := readBuffer.ReadByteArray("data", int(dataLen))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockFingerprint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockFingerprint")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockFingerprint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockFingerprint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockFingerprint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockFingerprint")
		}
		dataLen := uint16(uint16(len(m.GetData())))
		if err := WriteImplicitField(ctx, "dataLen", dataLen, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataLen' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockFingerprint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockFingerprint")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockFingerprint) IsAdsDiscoveryBlockFingerprint() {}

func (m *_AdsDiscoveryBlockFingerprint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockFingerprint) deepCopy() *_AdsDiscoveryBlockFingerprint {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockFingerprintCopy := &_AdsDiscoveryBlockFingerprint{
		m.AdsDiscoveryBlockContract.DeepCopy().(AdsDiscoveryBlockContract),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockFingerprintCopy
}

func (m *_AdsDiscoveryBlockFingerprint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
