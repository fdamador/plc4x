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

// HVACStatusFlags is the corresponding interface of HVACStatusFlags
type HVACStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetExpansion returns Expansion (property field)
	GetExpansion() bool
	// GetError returns Error (property field)
	GetError() bool
	// GetBusy returns Busy (property field)
	GetBusy() bool
	// GetDamperState returns DamperState (property field)
	GetDamperState() bool
	// GetFanActive returns FanActive (property field)
	GetFanActive() bool
	// GetHeatingPlant returns HeatingPlant (property field)
	GetHeatingPlant() bool
	// GetCoolingPlant returns CoolingPlant (property field)
	GetCoolingPlant() bool
	// GetIsDamperStateClosed returns IsDamperStateClosed (virtual field)
	GetIsDamperStateClosed() bool
	// GetIsDamperStateOpen returns IsDamperStateOpen (virtual field)
	GetIsDamperStateOpen() bool
	// IsHVACStatusFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACStatusFlags()
}

// _HVACStatusFlags is the data-structure of this message
type _HVACStatusFlags struct {
	Expansion    bool
	Error        bool
	Busy         bool
	DamperState  bool
	FanActive    bool
	HeatingPlant bool
	CoolingPlant bool
	// Reserved Fields
	reservedField0 *bool
}

var _ HVACStatusFlags = (*_HVACStatusFlags)(nil)

// NewHVACStatusFlags factory function for _HVACStatusFlags
func NewHVACStatusFlags(expansion bool, error bool, busy bool, damperState bool, fanActive bool, heatingPlant bool, coolingPlant bool) *_HVACStatusFlags {
	return &_HVACStatusFlags{Expansion: expansion, Error: error, Busy: busy, DamperState: damperState, FanActive: fanActive, HeatingPlant: heatingPlant, CoolingPlant: coolingPlant}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACStatusFlags) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACStatusFlags) GetError() bool {
	return m.Error
}

func (m *_HVACStatusFlags) GetBusy() bool {
	return m.Busy
}

func (m *_HVACStatusFlags) GetDamperState() bool {
	return m.DamperState
}

func (m *_HVACStatusFlags) GetFanActive() bool {
	return m.FanActive
}

func (m *_HVACStatusFlags) GetHeatingPlant() bool {
	return m.HeatingPlant
}

func (m *_HVACStatusFlags) GetCoolingPlant() bool {
	return m.CoolingPlant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACStatusFlags) GetIsDamperStateClosed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetDamperState()))
}

func (m *_HVACStatusFlags) GetIsDamperStateOpen() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetDamperState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACStatusFlags(structType any) HVACStatusFlags {
	if casted, ok := structType.(HVACStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACStatusFlags) GetTypeName() string {
	return "HVACStatusFlags"
}

func (m *_HVACStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1

	// Simple field (error)
	lengthInBits += 1

	// Simple field (busy)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (damperState)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (fanActive)
	lengthInBits += 1

	// Simple field (heatingPlant)
	lengthInBits += 1

	// Simple field (coolingPlant)
	lengthInBits += 1

	return lengthInBits
}

func (m *_HVACStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACStatusFlagsParse(ctx context.Context, theBytes []byte) (HVACStatusFlags, error) {
	return HVACStatusFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACStatusFlagsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStatusFlags, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStatusFlags, error) {
		return HVACStatusFlagsParseWithBuffer(ctx, readBuffer)
	}
}

func HVACStatusFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStatusFlags, error) {
	v, err := (&_HVACStatusFlags{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACStatusFlags HVACStatusFlags, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	expansion, err := ReadSimpleField(ctx, "expansion", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expansion' field"))
	}
	m.Expansion = expansion

	error, err := ReadSimpleField(ctx, "error", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	busy, err := ReadSimpleField(ctx, "busy", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'busy' field"))
	}
	m.Busy = busy

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	damperState, err := ReadSimpleField(ctx, "damperState", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'damperState' field"))
	}
	m.DamperState = damperState

	isDamperStateClosed, err := ReadVirtualField[bool](ctx, "isDamperStateClosed", (*bool)(nil), !(damperState))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDamperStateClosed' field"))
	}
	_ = isDamperStateClosed

	isDamperStateOpen, err := ReadVirtualField[bool](ctx, "isDamperStateOpen", (*bool)(nil), damperState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDamperStateOpen' field"))
	}
	_ = isDamperStateOpen

	fanActive, err := ReadSimpleField(ctx, "fanActive", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fanActive' field"))
	}
	m.FanActive = fanActive

	heatingPlant, err := ReadSimpleField(ctx, "heatingPlant", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'heatingPlant' field"))
	}
	m.HeatingPlant = heatingPlant

	coolingPlant, err := ReadSimpleField(ctx, "coolingPlant", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'coolingPlant' field"))
	}
	m.CoolingPlant = coolingPlant

	if closeErr := readBuffer.CloseContext("HVACStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACStatusFlags")
	}

	return m, nil
}

func (m *_HVACStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACStatusFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACStatusFlags")
	}

	if err := WriteSimpleField[bool](ctx, "expansion", m.GetExpansion(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'expansion' field")
	}

	if err := WriteSimpleField[bool](ctx, "error", m.GetError(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'error' field")
	}

	if err := WriteSimpleField[bool](ctx, "busy", m.GetBusy(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'busy' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "damperState", m.GetDamperState(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'damperState' field")
	}
	// Virtual field
	isDamperStateClosed := m.GetIsDamperStateClosed()
	_ = isDamperStateClosed
	if _isDamperStateClosedErr := writeBuffer.WriteVirtual(ctx, "isDamperStateClosed", m.GetIsDamperStateClosed()); _isDamperStateClosedErr != nil {
		return errors.Wrap(_isDamperStateClosedErr, "Error serializing 'isDamperStateClosed' field")
	}
	// Virtual field
	isDamperStateOpen := m.GetIsDamperStateOpen()
	_ = isDamperStateOpen
	if _isDamperStateOpenErr := writeBuffer.WriteVirtual(ctx, "isDamperStateOpen", m.GetIsDamperStateOpen()); _isDamperStateOpenErr != nil {
		return errors.Wrap(_isDamperStateOpenErr, "Error serializing 'isDamperStateOpen' field")
	}

	if err := WriteSimpleField[bool](ctx, "fanActive", m.GetFanActive(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'fanActive' field")
	}

	if err := WriteSimpleField[bool](ctx, "heatingPlant", m.GetHeatingPlant(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'heatingPlant' field")
	}

	if err := WriteSimpleField[bool](ctx, "coolingPlant", m.GetCoolingPlant(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'coolingPlant' field")
	}

	if popErr := writeBuffer.PopContext("HVACStatusFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACStatusFlags")
	}
	return nil
}

func (m *_HVACStatusFlags) IsHVACStatusFlags() {}

func (m *_HVACStatusFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACStatusFlags) deepCopy() *_HVACStatusFlags {
	if m == nil {
		return nil
	}
	_HVACStatusFlagsCopy := &_HVACStatusFlags{
		m.Expansion,
		m.Error,
		m.Busy,
		m.DamperState,
		m.FanActive,
		m.HeatingPlant,
		m.CoolingPlant,
		m.reservedField0,
	}
	return _HVACStatusFlagsCopy
}

func (m *_HVACStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
