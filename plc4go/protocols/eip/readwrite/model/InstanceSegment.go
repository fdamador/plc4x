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

// InstanceSegment is the corresponding interface of InstanceSegment
type InstanceSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetPathSegmentType returns PathSegmentType (property field)
	GetPathSegmentType() uint8
	// GetLogicalSegmentType returns LogicalSegmentType (property field)
	GetLogicalSegmentType() uint8
	// GetLogicalSegmentFormat returns LogicalSegmentFormat (property field)
	GetLogicalSegmentFormat() uint8
	// GetInstance returns Instance (property field)
	GetInstance() uint8
	// IsInstanceSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsInstanceSegment()
}

// _InstanceSegment is the data-structure of this message
type _InstanceSegment struct {
	PathSegmentType      uint8
	LogicalSegmentType   uint8
	LogicalSegmentFormat uint8
	Instance             uint8
}

var _ InstanceSegment = (*_InstanceSegment)(nil)

// NewInstanceSegment factory function for _InstanceSegment
func NewInstanceSegment(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, instance uint8) *_InstanceSegment {
	return &_InstanceSegment{PathSegmentType: pathSegmentType, LogicalSegmentType: logicalSegmentType, LogicalSegmentFormat: logicalSegmentFormat, Instance: instance}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InstanceSegment) GetPathSegmentType() uint8 {
	return m.PathSegmentType
}

func (m *_InstanceSegment) GetLogicalSegmentType() uint8 {
	return m.LogicalSegmentType
}

func (m *_InstanceSegment) GetLogicalSegmentFormat() uint8 {
	return m.LogicalSegmentFormat
}

func (m *_InstanceSegment) GetInstance() uint8 {
	return m.Instance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastInstanceSegment(structType any) InstanceSegment {
	if casted, ok := structType.(InstanceSegment); ok {
		return casted
	}
	if casted, ok := structType.(*InstanceSegment); ok {
		return *casted
	}
	return nil
}

func (m *_InstanceSegment) GetTypeName() string {
	return "InstanceSegment"
}

func (m *_InstanceSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (pathSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentFormat)
	lengthInBits += 2

	// Simple field (instance)
	lengthInBits += 8

	return lengthInBits
}

func (m *_InstanceSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InstanceSegmentParse(ctx context.Context, theBytes []byte) (InstanceSegment, error) {
	return InstanceSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InstanceSegmentParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (InstanceSegment, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (InstanceSegment, error) {
		return InstanceSegmentParseWithBuffer(ctx, readBuffer)
	}
}

func InstanceSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InstanceSegment, error) {
	v, err := (&_InstanceSegment{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_InstanceSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__instanceSegment InstanceSegment, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InstanceSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InstanceSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pathSegmentType, err := ReadSimpleField(ctx, "pathSegmentType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pathSegmentType' field"))
	}
	m.PathSegmentType = pathSegmentType

	logicalSegmentType, err := ReadSimpleField(ctx, "logicalSegmentType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalSegmentType' field"))
	}
	m.LogicalSegmentType = logicalSegmentType

	logicalSegmentFormat, err := ReadSimpleField(ctx, "logicalSegmentFormat", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalSegmentFormat' field"))
	}
	m.LogicalSegmentFormat = logicalSegmentFormat

	instance, err := ReadSimpleField(ctx, "instance", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instance' field"))
	}
	m.Instance = instance

	if closeErr := readBuffer.CloseContext("InstanceSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InstanceSegment")
	}

	return m, nil
}

func (m *_InstanceSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InstanceSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("InstanceSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InstanceSegment")
	}

	if err := WriteSimpleField[uint8](ctx, "pathSegmentType", m.GetPathSegmentType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'pathSegmentType' field")
	}

	if err := WriteSimpleField[uint8](ctx, "logicalSegmentType", m.GetLogicalSegmentType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalSegmentType' field")
	}

	if err := WriteSimpleField[uint8](ctx, "logicalSegmentFormat", m.GetLogicalSegmentFormat(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalSegmentFormat' field")
	}

	if err := WriteSimpleField[uint8](ctx, "instance", m.GetInstance(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'instance' field")
	}

	if popErr := writeBuffer.PopContext("InstanceSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InstanceSegment")
	}
	return nil
}

func (m *_InstanceSegment) IsInstanceSegment() {}

func (m *_InstanceSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_InstanceSegment) deepCopy() *_InstanceSegment {
	if m == nil {
		return nil
	}
	_InstanceSegmentCopy := &_InstanceSegment{
		m.PathSegmentType,
		m.LogicalSegmentType,
		m.LogicalSegmentFormat,
		m.Instance,
	}
	return _InstanceSegmentCopy
}

func (m *_InstanceSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
