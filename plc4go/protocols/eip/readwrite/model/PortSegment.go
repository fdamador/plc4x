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

// PortSegment is the corresponding interface of PortSegment
type PortSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() PortSegmentType
	// IsPortSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegment()
}

// _PortSegment is the data-structure of this message
type _PortSegment struct {
	PathSegmentContract
	SegmentType PortSegmentType
}

var _ PortSegment = (*_PortSegment)(nil)
var _ PathSegmentRequirements = (*_PortSegment)(nil)

// NewPortSegment factory function for _PortSegment
func NewPortSegment(segmentType PortSegmentType) *_PortSegment {
	if segmentType == nil {
		panic("segmentType of type PortSegmentType for PortSegment must not be nil")
	}
	_result := &_PortSegment{
		PathSegmentContract: NewPathSegment(),
		SegmentType:         segmentType,
	}
	_result.PathSegmentContract.(*_PathSegment)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegment) GetPathSegment() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegment) GetParent() PathSegmentContract {
	return m.PathSegmentContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegment) GetSegmentType() PortSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPortSegment(structType any) PortSegment {
	if casted, ok := structType.(PortSegment); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegment); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegment) GetTypeName() string {
	return "PortSegment"
}

func (m *_PortSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PathSegmentContract.(*_PathSegment).getLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PortSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PathSegment) (__portSegment PortSegment, err error) {
	m.PathSegmentContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentType, err := ReadSimpleField[PortSegmentType](ctx, "segmentType", ReadComplex[PortSegmentType](PortSegmentTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentType' field"))
	}
	m.SegmentType = segmentType

	if closeErr := readBuffer.CloseContext("PortSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegment")
	}

	return m, nil
}

func (m *_PortSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegment")
		}

		if err := WriteSimpleField[PortSegmentType](ctx, "segmentType", m.GetSegmentType(), WriteComplex[PortSegmentType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("PortSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegment")
		}
		return nil
	}
	return m.PathSegmentContract.(*_PathSegment).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegment) IsPortSegment() {}

func (m *_PortSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PortSegment) deepCopy() *_PortSegment {
	if m == nil {
		return nil
	}
	_PortSegmentCopy := &_PortSegment{
		m.PathSegmentContract.DeepCopy().(PathSegmentContract),
		m.SegmentType.DeepCopy().(PortSegmentType),
	}
	m.PathSegmentContract.(*_PathSegment)._SubType = m
	return _PortSegmentCopy
}

func (m *_PortSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
