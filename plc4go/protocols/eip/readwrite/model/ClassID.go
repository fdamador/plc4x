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

// ClassID is the corresponding interface of ClassID
type ClassID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LogicalSegmentType
	// GetFormat returns Format (property field)
	GetFormat() uint8
	// GetSegmentClass returns SegmentClass (property field)
	GetSegmentClass() uint8
	// IsClassID is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClassID()
}

// _ClassID is the data-structure of this message
type _ClassID struct {
	LogicalSegmentTypeContract
	Format       uint8
	SegmentClass uint8
}

var _ ClassID = (*_ClassID)(nil)
var _ LogicalSegmentTypeRequirements = (*_ClassID)(nil)

// NewClassID factory function for _ClassID
func NewClassID(format uint8, segmentClass uint8) *_ClassID {
	_result := &_ClassID{
		LogicalSegmentTypeContract: NewLogicalSegmentType(),
		Format:                     format,
		SegmentClass:               segmentClass,
	}
	_result.LogicalSegmentTypeContract.(*_LogicalSegmentType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ClassID) GetLogicalSegmentType() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClassID) GetParent() LogicalSegmentTypeContract {
	return m.LogicalSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClassID) GetFormat() uint8 {
	return m.Format
}

func (m *_ClassID) GetSegmentClass() uint8 {
	return m.SegmentClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastClassID(structType any) ClassID {
	if casted, ok := structType.(ClassID); ok {
		return casted
	}
	if casted, ok := structType.(*ClassID); ok {
		return *casted
	}
	return nil
}

func (m *_ClassID) GetTypeName() string {
	return "ClassID"
}

func (m *_ClassID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LogicalSegmentTypeContract.(*_LogicalSegmentType).getLengthInBits(ctx))

	// Simple field (format)
	lengthInBits += 2

	// Simple field (segmentClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClassID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ClassID) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LogicalSegmentType) (__classID ClassID, err error) {
	m.LogicalSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClassID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClassID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	format, err := ReadSimpleField(ctx, "format", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'format' field"))
	}
	m.Format = format

	segmentClass, err := ReadSimpleField(ctx, "segmentClass", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentClass' field"))
	}
	m.SegmentClass = segmentClass

	if closeErr := readBuffer.CloseContext("ClassID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClassID")
	}

	return m, nil
}

func (m *_ClassID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClassID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClassID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClassID")
		}

		if err := WriteSimpleField[uint8](ctx, "format", m.GetFormat(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
			return errors.Wrap(err, "Error serializing 'format' field")
		}

		if err := WriteSimpleField[uint8](ctx, "segmentClass", m.GetSegmentClass(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentClass' field")
		}

		if popErr := writeBuffer.PopContext("ClassID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClassID")
		}
		return nil
	}
	return m.LogicalSegmentTypeContract.(*_LogicalSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClassID) IsClassID() {}

func (m *_ClassID) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ClassID) deepCopy() *_ClassID {
	if m == nil {
		return nil
	}
	_ClassIDCopy := &_ClassID{
		m.LogicalSegmentTypeContract.DeepCopy().(LogicalSegmentTypeContract),
		m.Format,
		m.SegmentClass,
	}
	m.LogicalSegmentTypeContract.(*_LogicalSegmentType)._SubType = m
	return _ClassIDCopy
}

func (m *_ClassID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
