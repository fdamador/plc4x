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

// RelativeTimestamp is the corresponding interface of RelativeTimestamp
type RelativeTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() uint16
	// IsRelativeTimestamp is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRelativeTimestamp()
}

// _RelativeTimestamp is the data-structure of this message
type _RelativeTimestamp struct {
	Timestamp uint16
}

var _ RelativeTimestamp = (*_RelativeTimestamp)(nil)

// NewRelativeTimestamp factory function for _RelativeTimestamp
func NewRelativeTimestamp(timestamp uint16) *_RelativeTimestamp {
	return &_RelativeTimestamp{Timestamp: timestamp}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativeTimestamp) GetTimestamp() uint16 {
	return m.Timestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRelativeTimestamp(structType any) RelativeTimestamp {
	if casted, ok := structType.(RelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*RelativeTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_RelativeTimestamp) GetTypeName() string {
	return "RelativeTimestamp"
}

func (m *_RelativeTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 16

	return lengthInBits
}

func (m *_RelativeTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RelativeTimestampParse(ctx context.Context, theBytes []byte) (RelativeTimestamp, error) {
	return RelativeTimestampParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RelativeTimestampParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (RelativeTimestamp, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RelativeTimestamp, error) {
		return RelativeTimestampParseWithBuffer(ctx, readBuffer)
	}
}

func RelativeTimestampParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RelativeTimestamp, error) {
	v, err := (&_RelativeTimestamp{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_RelativeTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__relativeTimestamp RelativeTimestamp, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RelativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativeTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField(ctx, "timestamp", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	if closeErr := readBuffer.CloseContext("RelativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativeTimestamp")
	}

	return m, nil
}

func (m *_RelativeTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativeTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("RelativeTimestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RelativeTimestamp")
	}

	if err := WriteSimpleField[uint16](ctx, "timestamp", m.GetTimestamp(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if popErr := writeBuffer.PopContext("RelativeTimestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RelativeTimestamp")
	}
	return nil
}

func (m *_RelativeTimestamp) IsRelativeTimestamp() {}

func (m *_RelativeTimestamp) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RelativeTimestamp) deepCopy() *_RelativeTimestamp {
	if m == nil {
		return nil
	}
	_RelativeTimestampCopy := &_RelativeTimestamp{
		m.Timestamp,
	}
	return _RelativeTimestampCopy
}

func (m *_RelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
