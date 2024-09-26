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

// CALData is the corresponding interface of CALData
type CALData interface {
	CALDataContract
	CALDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCALData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALData()
}

// CALDataContract provides a set of functions which can be overwritten by a sub struct
type CALDataContract interface {
	utils.Copyable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() CALCommandTypeContainer
	// GetAdditionalData returns AdditionalData (property field)
	GetAdditionalData() CALData
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() CALCommandType
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (virtual field)
	GetSendIdentifyRequestBefore() bool
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// IsCALData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALData()
}

// CALDataRequirements provides a set of functions which need to be implemented by a sub struct
type CALDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() CALCommandType
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (discriminator field)
	GetSendIdentifyRequestBefore() bool
}

// _CALData is the data-structure of this message
type _CALData struct {
	_SubType             CALData
	CommandTypeContainer CALCommandTypeContainer
	AdditionalData       CALData

	// Arguments.
	RequestContext RequestContext
}

var _ CALDataContract = (*_CALData)(nil)

// NewCALData factory function for _CALData
func NewCALData(commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALData {
	return &_CALData{CommandTypeContainer: commandTypeContainer, AdditionalData: additionalData, RequestContext: requestContext}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALData) GetCommandTypeContainer() CALCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_CALData) GetAdditionalData() CALData {
	return m.AdditionalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_CALData) GetCommandType() CALCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	additionalData := m.GetAdditionalData()
	_ = additionalData
	return CastCALCommandType(m.GetCommandTypeContainer().CommandType())
}

func (pm *_CALData) GetSendIdentifyRequestBefore() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	additionalData := m.GetAdditionalData()
	_ = additionalData
	return bool(utils.InlineIf(bool((m.GetRequestContext()) != (nil)), func() any { return bool(m.GetRequestContext().GetSendIdentifyRequestBefore()) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCALData(structType any) CALData {
	if casted, ok := structType.(CALData); ok {
		return casted
	}
	if casted, ok := structType.(*CALData); ok {
		return *casted
	}
	return nil
}

func (m *_CALData) GetTypeName() string {
	return "CALData"
}

func (m *_CALData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (additionalData)
	if m.AdditionalData != nil {
		lengthInBits += m.AdditionalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CALData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CALDataParse[T CALData](ctx context.Context, theBytes []byte, requestContext RequestContext) (T, error) {
	return CALDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), requestContext)
}

func CALDataParseWithBufferProducer[T CALData](requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CALDataParseWithBuffer[T](ctx, readBuffer, requestContext)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CALDataParseWithBuffer[T CALData](ctx context.Context, readBuffer utils.ReadBuffer, requestContext RequestContext) (T, error) {
	v, err := (&_CALData{RequestContext: requestContext}).parse(ctx, readBuffer, requestContext)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_CALData) parse(ctx context.Context, readBuffer utils.ReadBuffer, requestContext RequestContext) (__cALData CALData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsCALCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[CALCommandTypeContainer](ctx, "commandTypeContainer", "CALCommandTypeContainer", ReadEnum(CALCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[CALCommandType](ctx, "commandType", (*CALCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	sendIdentifyRequestBefore, err := ReadVirtualField[bool](ctx, "sendIdentifyRequestBefore", (*bool)(nil), utils.InlineIf(bool((requestContext) != (nil)), func() any { return bool(requestContext.GetSendIdentifyRequestBefore()) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sendIdentifyRequestBefore' field"))
	}
	_ = sendIdentifyRequestBefore

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CALData
	switch {
	case commandType == CALCommandType_RESET: // CALDataReset
		if _child, err = new(_CALDataReset).parse(ctx, readBuffer, m, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataReset for type-switch of CALData")
		}
	case commandType == CALCommandType_RECALL: // CALDataRecall
		if _child, err = new(_CALDataRecall).parse(ctx, readBuffer, m, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataRecall for type-switch of CALData")
		}
	case commandType == CALCommandType_IDENTIFY: // CALDataIdentify
		if _child, err = new(_CALDataIdentify).parse(ctx, readBuffer, m, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataIdentify for type-switch of CALData")
		}
	case commandType == CALCommandType_GET_STATUS: // CALDataGetStatus
		if _child, err = new(_CALDataGetStatus).parse(ctx, readBuffer, m, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataGetStatus for type-switch of CALData")
		}
	case commandType == CALCommandType_WRITE: // CALDataWrite
		if _child, err = new(_CALDataWrite).parse(ctx, readBuffer, m, commandTypeContainer, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataWrite for type-switch of CALData")
		}
	case commandType == CALCommandType_REPLY && sendIdentifyRequestBefore == bool(true): // CALDataIdentifyReply
		if _child, err = new(_CALDataIdentifyReply).parse(ctx, readBuffer, m, commandTypeContainer, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataIdentifyReply for type-switch of CALData")
		}
	case commandType == CALCommandType_REPLY: // CALDataReply
		if _child, err = new(_CALDataReply).parse(ctx, readBuffer, m, commandTypeContainer, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataReply for type-switch of CALData")
		}
	case commandType == CALCommandType_ACKNOWLEDGE: // CALDataAcknowledge
		if _child, err = new(_CALDataAcknowledge).parse(ctx, readBuffer, m, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataAcknowledge for type-switch of CALData")
		}
	case commandType == CALCommandType_STATUS: // CALDataStatus
		if _child, err = new(_CALDataStatus).parse(ctx, readBuffer, m, commandTypeContainer, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataStatus for type-switch of CALData")
		}
	case commandType == CALCommandType_STATUS_EXTENDED: // CALDataStatusExtended
		if _child, err = new(_CALDataStatusExtended).parse(ctx, readBuffer, m, commandTypeContainer, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALDataStatusExtended for type-switch of CALData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, sendIdentifyRequestBefore=%v]", commandType, sendIdentifyRequestBefore)
	}

	var additionalData CALData
	_additionalData, err := ReadOptionalField[CALData](ctx, "additionalData", ReadComplex[CALData](CALDataParseWithBufferProducer[CALData]((RequestContext)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalData' field"))
	}
	if _additionalData != nil {
		additionalData = *_additionalData
		m.AdditionalData = additionalData
	}

	if closeErr := readBuffer.CloseContext("CALData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALData")
	}

	return _child, nil
}

func (pm *_CALData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CALData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CALData")
	}

	if err := WriteSimpleEnumField[CALCommandTypeContainer](ctx, "commandTypeContainer", "CALCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[CALCommandTypeContainer, uint8](CALCommandTypeContainer.GetValue, CALCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}
	// Virtual field
	sendIdentifyRequestBefore := m.GetSendIdentifyRequestBefore()
	_ = sendIdentifyRequestBefore
	if _sendIdentifyRequestBeforeErr := writeBuffer.WriteVirtual(ctx, "sendIdentifyRequestBefore", m.GetSendIdentifyRequestBefore()); _sendIdentifyRequestBeforeErr != nil {
		return errors.Wrap(_sendIdentifyRequestBeforeErr, "Error serializing 'sendIdentifyRequestBefore' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteOptionalField[CALData](ctx, "additionalData", GetRef(m.GetAdditionalData()), WriteComplex[CALData](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'additionalData' field")
	}

	if popErr := writeBuffer.PopContext("CALData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CALData")
	}
	return nil
}

////
// Arguments Getter

func (m *_CALData) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_CALData) IsCALData() {}

func (m *_CALData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALData) deepCopy() *_CALData {
	if m == nil {
		return nil
	}
	_CALDataCopy := &_CALData{
		nil, // will be set by child
		m.CommandTypeContainer,
		m.AdditionalData.DeepCopy().(CALData),
		m.RequestContext,
	}
	return _CALDataCopy
}
