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

// BACnetUnconfirmedServiceRequestUnconfirmedEventNotification is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
type BACnetUnconfirmedServiceRequestUnconfirmedEventNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetMessageText returns MessageText (property field)
	GetMessageText() BACnetContextTagCharacterString
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetContextTagBoolean
	// GetFromState returns FromState (property field)
	GetFromState() BACnetEventStateTagged
	// GetToState returns ToState (property field)
	GetToState() BACnetEventStateTagged
	// GetEventValues returns EventValues (property field)
	GetEventValues() BACnetNotificationParameters
	// IsBACnetUnconfirmedServiceRequestUnconfirmedEventNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestUnconfirmedEventNotification()
}

// _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification struct {
	BACnetUnconfirmedServiceRequestContract
	ProcessIdentifier          BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier BACnetContextTagObjectIdentifier
	EventObjectIdentifier      BACnetContextTagObjectIdentifier
	Timestamp                  BACnetTimeStampEnclosed
	NotificationClass          BACnetContextTagUnsignedInteger
	Priority                   BACnetContextTagUnsignedInteger
	EventType                  BACnetEventTypeTagged
	MessageText                BACnetContextTagCharacterString
	NotifyType                 BACnetNotifyTypeTagged
	AckRequired                BACnetContextTagBoolean
	FromState                  BACnetEventStateTagged
	ToState                    BACnetEventStateTagged
	EventValues                BACnetNotificationParameters
}

var _ BACnetUnconfirmedServiceRequestUnconfirmedEventNotification = (*_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification)(nil)

// NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification factory function for _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
func NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(processIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, eventObjectIdentifier BACnetContextTagObjectIdentifier, timestamp BACnetTimeStampEnclosed, notificationClass BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger, eventType BACnetEventTypeTagged, messageText BACnetContextTagCharacterString, notifyType BACnetNotifyTypeTagged, ackRequired BACnetContextTagBoolean, fromState BACnetEventStateTagged, toState BACnetEventStateTagged, eventValues BACnetNotificationParameters, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	if processIdentifier == nil {
		panic("processIdentifier of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if eventObjectIdentifier == nil {
		panic("eventObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if timestamp == nil {
		panic("timestamp of type BACnetTimeStampEnclosed for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if notificationClass == nil {
		panic("notificationClass of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if priority == nil {
		panic("priority of type BACnetContextTagUnsignedInteger for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if eventType == nil {
		panic("eventType of type BACnetEventTypeTagged for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if notifyType == nil {
		panic("notifyType of type BACnetNotifyTypeTagged for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	if toState == nil {
		panic("toState of type BACnetEventStateTagged for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		ProcessIdentifier:                       processIdentifier,
		InitiatingDeviceIdentifier:              initiatingDeviceIdentifier,
		EventObjectIdentifier:                   eventObjectIdentifier,
		Timestamp:                               timestamp,
		NotificationClass:                       notificationClass,
		Priority:                                priority,
		EventType:                               eventType,
		MessageText:                             messageText,
		NotifyType:                              notifyType,
		AckRequired:                             ackRequired,
		FromState:                               fromState,
		ToState:                                 toState,
		EventValues:                             eventValues,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetNotificationClass() BACnetContextTagUnsignedInteger {
	return m.NotificationClass
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetMessageText() BACnetContextTagCharacterString {
	return m.MessageText
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetAckRequired() BACnetContextTagBoolean {
	return m.AckRequired
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetFromState() BACnetEventStateTagged {
	return m.FromState
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetToState() BACnetEventStateTagged {
	return m.ToState
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventValues() BACnetNotificationParameters {
	return m.EventValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(structType any) BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (notificationClass)
	lengthInBits += m.NotificationClass.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits(ctx)

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// Optional Field (messageText)
	if m.MessageText != nil {
		lengthInBits += m.MessageText.GetLengthInBits(ctx)
	}

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits(ctx)

	// Optional Field (ackRequired)
	if m.AckRequired != nil {
		lengthInBits += m.AckRequired.GetLengthInBits(ctx)
	}

	// Optional Field (fromState)
	if m.FromState != nil {
		lengthInBits += m.FromState.GetLengthInBits(ctx)
	}

	// Simple field (toState)
	lengthInBits += m.ToState.GetLengthInBits(ctx)

	// Optional Field (eventValues)
	if m.EventValues != nil {
		lengthInBits += m.EventValues.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUnconfirmedEventNotification BACnetUnconfirmedServiceRequestUnconfirmedEventNotification, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	m.ProcessIdentifier = processIdentifier

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier

	eventObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventObjectIdentifier' field"))
	}
	m.EventObjectIdentifier = eventObjectIdentifier

	timestamp, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	notificationClass, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationClass", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClass' field"))
	}
	m.NotificationClass = notificationClass

	priority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(6)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}
	m.EventType = eventType

	var messageText BACnetContextTagCharacterString
	_messageText, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "messageText", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(7)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageText' field"))
	}
	if _messageText != nil {
		messageText = *_messageText
		m.MessageText = messageText
	}

	notifyType, err := ReadSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", ReadComplex[BACnetNotifyTypeTagged](BACnetNotifyTypeTaggedParseWithBufferProducer((uint8)(uint8(8)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notifyType' field"))
	}
	m.NotifyType = notifyType

	var ackRequired BACnetContextTagBoolean
	_ackRequired, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(9)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackRequired' field"))
	}
	if _ackRequired != nil {
		ackRequired = *_ackRequired
		m.AckRequired = ackRequired
	}

	var fromState BACnetEventStateTagged
	_fromState, err := ReadOptionalField[BACnetEventStateTagged](ctx, "fromState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(10)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fromState' field"))
	}
	if _fromState != nil {
		fromState = *_fromState
		m.FromState = fromState
	}

	toState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "toState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(11)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toState' field"))
	}
	m.ToState = toState

	var eventValues BACnetNotificationParameters
	_eventValues, err := ReadOptionalField[BACnetNotificationParameters](ctx, "eventValues", ReadComplex[BACnetNotificationParameters](BACnetNotificationParametersParseWithBufferProducer[BACnetNotificationParameters]((uint8)(uint8(12)), (BACnetObjectType)(eventObjectIdentifier.GetObjectType())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventValues' field"))
	}
	if _eventValues != nil {
		eventValues = *_eventValues
		m.EventValues = eventValues
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", m.GetEventObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationClass", m.GetNotificationClass(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationClass' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "priority", m.GetPriority(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteSimpleField[BACnetEventTypeTagged](ctx, "eventType", m.GetEventType(), WriteComplex[BACnetEventTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventType' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "messageText", GetRef(m.GetMessageText()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'messageText' field")
		}

		if err := WriteSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", m.GetNotifyType(), WriteComplex[BACnetNotifyTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notifyType' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", GetRef(m.GetAckRequired()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'ackRequired' field")
		}

		if err := WriteOptionalField[BACnetEventStateTagged](ctx, "fromState", GetRef(m.GetFromState()), WriteComplex[BACnetEventStateTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'fromState' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "toState", m.GetToState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'toState' field")
		}

		if err := WriteOptionalField[BACnetNotificationParameters](ctx, "eventValues", GetRef(m.GetEventValues()), WriteComplex[BACnetNotificationParameters](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'eventValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) IsBACnetUnconfirmedServiceRequestUnconfirmedEventNotification() {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) deepCopy() *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationCopy := &_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		m.BACnetUnconfirmedServiceRequestContract.DeepCopy().(BACnetUnconfirmedServiceRequestContract),
		m.ProcessIdentifier.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.InitiatingDeviceIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.EventObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.Timestamp.DeepCopy().(BACnetTimeStampEnclosed),
		m.NotificationClass.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Priority.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.EventType.DeepCopy().(BACnetEventTypeTagged),
		m.MessageText.DeepCopy().(BACnetContextTagCharacterString),
		m.NotifyType.DeepCopy().(BACnetNotifyTypeTagged),
		m.AckRequired.DeepCopy().(BACnetContextTagBoolean),
		m.FromState.DeepCopy().(BACnetEventStateTagged),
		m.ToState.DeepCopy().(BACnetEventStateTagged),
		m.EventValues.DeepCopy().(BACnetNotificationParameters),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationCopy
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
