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

// ListOfCovNotifications is the corresponding interface of ListOfCovNotifications
type ListOfCovNotifications interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []ListOfCovNotificationsValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsListOfCovNotifications is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsListOfCovNotifications()
}

// _ListOfCovNotifications is the data-structure of this message
type _ListOfCovNotifications struct {
	MonitoredObjectIdentifier BACnetContextTagObjectIdentifier
	OpeningTag                BACnetOpeningTag
	ListOfValues              []ListOfCovNotificationsValue
	ClosingTag                BACnetClosingTag
}

var _ ListOfCovNotifications = (*_ListOfCovNotifications)(nil)

// NewListOfCovNotifications factory function for _ListOfCovNotifications
func NewListOfCovNotifications(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, openingTag BACnetOpeningTag, listOfValues []ListOfCovNotificationsValue, closingTag BACnetClosingTag) *_ListOfCovNotifications {
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for ListOfCovNotifications must not be nil")
	}
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for ListOfCovNotifications must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for ListOfCovNotifications must not be nil")
	}
	return &_ListOfCovNotifications{MonitoredObjectIdentifier: monitoredObjectIdentifier, OpeningTag: openingTag, ListOfValues: listOfValues, ClosingTag: closingTag}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotifications) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_ListOfCovNotifications) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ListOfCovNotifications) GetListOfValues() []ListOfCovNotificationsValue {
	return m.ListOfValues
}

func (m *_ListOfCovNotifications) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastListOfCovNotifications(structType any) ListOfCovNotifications {
	if casted, ok := structType.(ListOfCovNotifications); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotifications); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotifications) GetTypeName() string {
	return "ListOfCovNotifications"
}

func (m *_ListOfCovNotifications) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ListOfCovNotifications) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ListOfCovNotificationsParse(ctx context.Context, theBytes []byte) (ListOfCovNotifications, error) {
	return ListOfCovNotificationsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ListOfCovNotificationsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
		return ListOfCovNotificationsParseWithBuffer(ctx, readBuffer)
	}
}

func ListOfCovNotificationsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
	v, err := (&_ListOfCovNotifications{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ListOfCovNotifications) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__listOfCovNotifications ListOfCovNotifications, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ListOfCovNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotifications")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfValues, err := ReadTerminatedArrayField[ListOfCovNotificationsValue](ctx, "listOfValues", ReadComplex[ListOfCovNotificationsValue](ListOfCovNotificationsValueParseWithBufferProducer((BACnetObjectType)(monitoredObjectIdentifier.GetObjectType())), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	m.ListOfValues = listOfValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("ListOfCovNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotifications")
	}

	return m, nil
}

func (m *_ListOfCovNotifications) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListOfCovNotifications) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ListOfCovNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotifications")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfValues", m.GetListOfValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("ListOfCovNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotifications")
	}
	return nil
}

func (m *_ListOfCovNotifications) IsListOfCovNotifications() {}

func (m *_ListOfCovNotifications) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ListOfCovNotifications) deepCopy() *_ListOfCovNotifications {
	if m == nil {
		return nil
	}
	_ListOfCovNotificationsCopy := &_ListOfCovNotifications{
		m.MonitoredObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[ListOfCovNotificationsValue, ListOfCovNotificationsValue](m.ListOfValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	return _ListOfCovNotificationsCopy
}

func (m *_ListOfCovNotifications) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
