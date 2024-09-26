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

// MonitoredSALShortFormBasicMode is the corresponding interface of MonitoredSALShortFormBasicMode
type MonitoredSALShortFormBasicMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MonitoredSAL
	// GetCounts returns Counts (property field)
	GetCounts() byte
	// GetBridgeCount returns BridgeCount (property field)
	GetBridgeCount() *uint8
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() *uint8
	// GetNoCounts returns NoCounts (property field)
	GetNoCounts() *byte
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// IsMonitoredSALShortFormBasicMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoredSALShortFormBasicMode()
}

// _MonitoredSALShortFormBasicMode is the data-structure of this message
type _MonitoredSALShortFormBasicMode struct {
	MonitoredSALContract
	Counts        byte
	BridgeCount   *uint8
	NetworkNumber *uint8
	NoCounts      *byte
	Application   ApplicationIdContainer
	SalData       SALData
}

var _ MonitoredSALShortFormBasicMode = (*_MonitoredSALShortFormBasicMode)(nil)
var _ MonitoredSALRequirements = (*_MonitoredSALShortFormBasicMode)(nil)

// NewMonitoredSALShortFormBasicMode factory function for _MonitoredSALShortFormBasicMode
func NewMonitoredSALShortFormBasicMode(salType byte, counts byte, bridgeCount *uint8, networkNumber *uint8, noCounts *byte, application ApplicationIdContainer, salData SALData, cBusOptions CBusOptions) *_MonitoredSALShortFormBasicMode {
	_result := &_MonitoredSALShortFormBasicMode{
		MonitoredSALContract: NewMonitoredSAL(salType, cBusOptions),
		Counts:               counts,
		BridgeCount:          bridgeCount,
		NetworkNumber:        networkNumber,
		NoCounts:             noCounts,
		Application:          application,
		SalData:              salData,
	}
	_result.MonitoredSALContract.(*_MonitoredSAL)._SubType = _result
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

func (m *_MonitoredSALShortFormBasicMode) GetParent() MonitoredSALContract {
	return m.MonitoredSALContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALShortFormBasicMode) GetCounts() byte {
	return m.Counts
}

func (m *_MonitoredSALShortFormBasicMode) GetBridgeCount() *uint8 {
	return m.BridgeCount
}

func (m *_MonitoredSALShortFormBasicMode) GetNetworkNumber() *uint8 {
	return m.NetworkNumber
}

func (m *_MonitoredSALShortFormBasicMode) GetNoCounts() *byte {
	return m.NoCounts
}

func (m *_MonitoredSALShortFormBasicMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALShortFormBasicMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoredSALShortFormBasicMode(structType any) MonitoredSALShortFormBasicMode {
	if casted, ok := structType.(MonitoredSALShortFormBasicMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALShortFormBasicMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALShortFormBasicMode) GetTypeName() string {
	return "MonitoredSALShortFormBasicMode"
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MonitoredSALContract.(*_MonitoredSAL).getLengthInBits(ctx))

	// Optional Field (bridgeCount)
	if m.BridgeCount != nil {
		lengthInBits += 8
	}

	// Optional Field (networkNumber)
	if m.NetworkNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (noCounts)
	if m.NoCounts != nil {
		lengthInBits += 8
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredSALShortFormBasicMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MonitoredSAL, cBusOptions CBusOptions) (__monitoredSALShortFormBasicMode MonitoredSALShortFormBasicMode, err error) {
	m.MonitoredSALContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALShortFormBasicMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALShortFormBasicMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	counts, err := ReadPeekField[byte](ctx, "counts", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'counts' field"))
	}
	m.Counts = counts

	var bridgeCount *uint8
	bridgeCount, err = ReadOptionalField[uint8](ctx, "bridgeCount", ReadUnsignedByte(readBuffer, uint8(8)), bool((counts) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeCount' field"))
	}
	m.BridgeCount = bridgeCount

	var networkNumber *uint8
	networkNumber, err = ReadOptionalField[uint8](ctx, "networkNumber", ReadUnsignedByte(readBuffer, uint8(8)), bool((counts) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumber' field"))
	}
	m.NetworkNumber = networkNumber

	var noCounts *byte
	noCounts, err = ReadOptionalField[byte](ctx, "noCounts", ReadByte(readBuffer, 8), bool((counts) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noCounts' field"))
	}
	m.NoCounts = noCounts

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	var salData SALData
	_salData, err := ReadOptionalField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(application.ApplicationId())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	if _salData != nil {
		salData = *_salData
		m.SalData = salData
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALShortFormBasicMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALShortFormBasicMode")
	}

	return m, nil
}

func (m *_MonitoredSALShortFormBasicMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALShortFormBasicMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALShortFormBasicMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALShortFormBasicMode")
		}

		if err := WriteOptionalField[uint8](ctx, "bridgeCount", m.GetBridgeCount(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeCount' field")
		}

		if err := WriteOptionalField[uint8](ctx, "networkNumber", m.GetNetworkNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'networkNumber' field")
		}

		if err := WriteOptionalField[byte](ctx, "noCounts", m.GetNoCounts(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'noCounts' field")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteOptionalField[SALData](ctx, "salData", GetRef(m.GetSalData()), WriteComplex[SALData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'salData' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredSALShortFormBasicMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALShortFormBasicMode")
		}
		return nil
	}
	return m.MonitoredSALContract.(*_MonitoredSAL).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALShortFormBasicMode) IsMonitoredSALShortFormBasicMode() {}

func (m *_MonitoredSALShortFormBasicMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MonitoredSALShortFormBasicMode) deepCopy() *_MonitoredSALShortFormBasicMode {
	if m == nil {
		return nil
	}
	_MonitoredSALShortFormBasicModeCopy := &_MonitoredSALShortFormBasicMode{
		m.MonitoredSALContract.DeepCopy().(MonitoredSALContract),
		m.Counts,
		utils.CopyPtr[uint8](m.BridgeCount),
		utils.CopyPtr[uint8](m.NetworkNumber),
		utils.CopyPtr[byte](m.NoCounts),
		m.Application,
		m.SalData.DeepCopy().(SALData),
	}
	m.MonitoredSALContract.(*_MonitoredSAL)._SubType = m
	return _MonitoredSALShortFormBasicModeCopy
}

func (m *_MonitoredSALShortFormBasicMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
