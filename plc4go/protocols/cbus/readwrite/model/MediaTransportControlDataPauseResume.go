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

// MediaTransportControlDataPauseResume is the corresponding interface of MediaTransportControlDataPauseResume
type MediaTransportControlDataPauseResume interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsPause returns IsPause (virtual field)
	GetIsPause() bool
	// GetIsResume returns IsResume (virtual field)
	GetIsResume() bool
	// IsMediaTransportControlDataPauseResume is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataPauseResume()
}

// _MediaTransportControlDataPauseResume is the data-structure of this message
type _MediaTransportControlDataPauseResume struct {
	MediaTransportControlDataContract
	Operation byte
}

var _ MediaTransportControlDataPauseResume = (*_MediaTransportControlDataPauseResume)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataPauseResume)(nil)

// NewMediaTransportControlDataPauseResume factory function for _MediaTransportControlDataPauseResume
func NewMediaTransportControlDataPauseResume(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, operation byte) *_MediaTransportControlDataPauseResume {
	_result := &_MediaTransportControlDataPauseResume{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		Operation:                         operation,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
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

func (m *_MediaTransportControlDataPauseResume) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataPauseResume) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataPauseResume) GetIsPause() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataPauseResume) GetIsResume() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataPauseResume(structType any) MediaTransportControlDataPauseResume {
	if casted, ok := structType.(MediaTransportControlDataPauseResume); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataPauseResume); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataPauseResume) GetTypeName() string {
	return "MediaTransportControlDataPauseResume"
}

func (m *_MediaTransportControlDataPauseResume) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataPauseResume) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataPauseResume) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataPauseResume MediaTransportControlDataPauseResume, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataPauseResume"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataPauseResume")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operation, err := ReadSimpleField(ctx, "operation", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}
	m.Operation = operation

	isPause, err := ReadVirtualField[bool](ctx, "isPause", (*bool)(nil), bool((operation) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPause' field"))
	}
	_ = isPause

	isResume, err := ReadVirtualField[bool](ctx, "isResume", (*bool)(nil), bool((operation) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isResume' field"))
	}
	_ = isResume

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataPauseResume"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataPauseResume")
	}

	return m, nil
}

func (m *_MediaTransportControlDataPauseResume) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataPauseResume) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataPauseResume"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataPauseResume")
		}

		if err := WriteSimpleField[byte](ctx, "operation", m.GetOperation(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'operation' field")
		}
		// Virtual field
		isPause := m.GetIsPause()
		_ = isPause
		if _isPauseErr := writeBuffer.WriteVirtual(ctx, "isPause", m.GetIsPause()); _isPauseErr != nil {
			return errors.Wrap(_isPauseErr, "Error serializing 'isPause' field")
		}
		// Virtual field
		isResume := m.GetIsResume()
		_ = isResume
		if _isResumeErr := writeBuffer.WriteVirtual(ctx, "isResume", m.GetIsResume()); _isResumeErr != nil {
			return errors.Wrap(_isResumeErr, "Error serializing 'isResume' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataPauseResume"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataPauseResume")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataPauseResume) IsMediaTransportControlDataPauseResume() {}

func (m *_MediaTransportControlDataPauseResume) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataPauseResume) deepCopy() *_MediaTransportControlDataPauseResume {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataPauseResumeCopy := &_MediaTransportControlDataPauseResume{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.Operation,
	}
	m.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataPauseResumeCopy
}

func (m *_MediaTransportControlDataPauseResume) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
