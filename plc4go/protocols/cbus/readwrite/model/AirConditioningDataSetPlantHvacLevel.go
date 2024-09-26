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

// AirConditioningDataSetPlantHvacLevel is the corresponding interface of AirConditioningDataSetPlantHvacLevel
type AirConditioningDataSetPlantHvacLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetLevel returns Level (property field)
	GetLevel() HVACTemperature
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
	// IsAirConditioningDataSetPlantHvacLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetPlantHvacLevel()
}

// _AirConditioningDataSetPlantHvacLevel is the data-structure of this message
type _AirConditioningDataSetPlantHvacLevel struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	HvacModeAndFlags HVACModeAndFlags
	HvacType         HVACType
	Level            HVACTemperature
	RawLevel         HVACRawLevels
	AuxLevel         HVACAuxiliaryLevel
}

var _ AirConditioningDataSetPlantHvacLevel = (*_AirConditioningDataSetPlantHvacLevel)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetPlantHvacLevel)(nil)

// NewAirConditioningDataSetPlantHvacLevel factory function for _AirConditioningDataSetPlantHvacLevel
func NewAirConditioningDataSetPlantHvacLevel(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, hvacModeAndFlags HVACModeAndFlags, hvacType HVACType, level HVACTemperature, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel) *_AirConditioningDataSetPlantHvacLevel {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetPlantHvacLevel must not be nil")
	}
	if hvacModeAndFlags == nil {
		panic("hvacModeAndFlags of type HVACModeAndFlags for AirConditioningDataSetPlantHvacLevel must not be nil")
	}
	_result := &_AirConditioningDataSetPlantHvacLevel{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HvacModeAndFlags:            hvacModeAndFlags,
		HvacType:                    hvacType,
		Level:                       level,
		RawLevel:                    rawLevel,
		AuxLevel:                    auxLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
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

func (m *_AirConditioningDataSetPlantHvacLevel) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetPlantHvacLevel) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLevel() HVACTemperature {
	return m.Level
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetPlantHvacLevel(structType any) AirConditioningDataSetPlantHvacLevel {
	if casted, ok := structType.(AirConditioningDataSetPlantHvacLevel); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetPlantHvacLevel); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetTypeName() string {
	return "AirConditioningDataSetPlantHvacLevel"
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	// Simple field (hvacType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetPlantHvacLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetPlantHvacLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetPlantHvacLevel AirConditioningDataSetPlantHvacLevel, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetPlantHvacLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetPlantHvacLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	hvacType, err := ReadEnumField[HVACType](ctx, "hvacType", "HVACType", ReadEnum(HVACTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacType' field"))
	}
	m.HvacType = hvacType

	var level HVACTemperature
	_level, err := ReadOptionalField[HVACTemperature](ctx, "level", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelTemperature())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	var auxLevel HVACAuxiliaryLevel
	_auxLevel, err := ReadOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", ReadComplex[HVACAuxiliaryLevel](HVACAuxiliaryLevelParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsAuxLevelUsed())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxLevel' field"))
	}
	if _auxLevel != nil {
		auxLevel = *_auxLevel
		m.AuxLevel = auxLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetPlantHvacLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetPlantHvacLevel")
	}

	return m, nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetPlantHvacLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetPlantHvacLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetPlantHvacLevel")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if err := WriteSimpleEnumField[HVACType](ctx, "hvacType", "HVACType", m.GetHvacType(), WriteEnum[HVACType, uint8](HVACType.GetValue, HVACType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacType' field")
		}

		if err := WriteOptionalField[HVACTemperature](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACTemperature](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if err := WriteOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", GetRef(m.GetAuxLevel()), WriteComplex[HVACAuxiliaryLevel](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'auxLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetPlantHvacLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetPlantHvacLevel")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetPlantHvacLevel) IsAirConditioningDataSetPlantHvacLevel() {}

func (m *_AirConditioningDataSetPlantHvacLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetPlantHvacLevel) deepCopy() *_AirConditioningDataSetPlantHvacLevel {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetPlantHvacLevelCopy := &_AirConditioningDataSetPlantHvacLevel{
		m.AirConditioningDataContract.DeepCopy().(AirConditioningDataContract),
		m.ZoneGroup,
		m.ZoneList.DeepCopy().(HVACZoneList),
		m.HvacModeAndFlags.DeepCopy().(HVACModeAndFlags),
		m.HvacType,
		m.Level.DeepCopy().(HVACTemperature),
		m.RawLevel.DeepCopy().(HVACRawLevels),
		m.AuxLevel.DeepCopy().(HVACAuxiliaryLevel),
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetPlantHvacLevelCopy
}

func (m *_AirConditioningDataSetPlantHvacLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
