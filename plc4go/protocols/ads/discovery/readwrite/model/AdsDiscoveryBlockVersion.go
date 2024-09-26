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

// AdsDiscoveryBlockVersion is the corresponding interface of AdsDiscoveryBlockVersion
type AdsDiscoveryBlockVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetVersionData returns VersionData (property field)
	GetVersionData() []byte
	// IsAdsDiscoveryBlockVersion is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockVersion()
}

// _AdsDiscoveryBlockVersion is the data-structure of this message
type _AdsDiscoveryBlockVersion struct {
	AdsDiscoveryBlockContract
	VersionData []byte
}

var _ AdsDiscoveryBlockVersion = (*_AdsDiscoveryBlockVersion)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockVersion)(nil)

// NewAdsDiscoveryBlockVersion factory function for _AdsDiscoveryBlockVersion
func NewAdsDiscoveryBlockVersion(versionData []byte) *_AdsDiscoveryBlockVersion {
	_result := &_AdsDiscoveryBlockVersion{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		VersionData:               versionData,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockVersion) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockVersion) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockVersion) GetVersionData() []byte {
	return m.VersionData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockVersion(structType any) AdsDiscoveryBlockVersion {
	if casted, ok := structType.(AdsDiscoveryBlockVersion); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockVersion); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockVersion) GetTypeName() string {
	return "AdsDiscoveryBlockVersion"
}

func (m *_AdsDiscoveryBlockVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Implicit Field (versionDataLen)
	lengthInBits += 16

	// Array field
	if len(m.VersionData) > 0 {
		lengthInBits += 8 * uint16(len(m.VersionData))
	}

	return lengthInBits
}

func (m *_AdsDiscoveryBlockVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockVersion) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockVersion AdsDiscoveryBlockVersion, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	versionDataLen, err := ReadImplicitField[uint16](ctx, "versionDataLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'versionDataLen' field"))
	}
	_ = versionDataLen

	versionData, err := readBuffer.ReadByteArray("versionData", int(versionDataLen))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'versionData' field"))
	}
	m.VersionData = versionData

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockVersion")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockVersion")
		}
		versionDataLen := uint16(uint16(len(m.GetVersionData())))
		if err := WriteImplicitField(ctx, "versionDataLen", versionDataLen, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'versionDataLen' field")
		}

		if err := WriteByteArrayField(ctx, "versionData", m.GetVersionData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'versionData' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockVersion")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockVersion) IsAdsDiscoveryBlockVersion() {}

func (m *_AdsDiscoveryBlockVersion) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockVersion) deepCopy() *_AdsDiscoveryBlockVersion {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockVersionCopy := &_AdsDiscoveryBlockVersion{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.VersionData),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockVersionCopy
}

func (m *_AdsDiscoveryBlockVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
