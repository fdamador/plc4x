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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ContinuationPoint is the corresponding interface of ContinuationPoint
type ContinuationPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsContinuationPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsContinuationPoint()
}

// _ContinuationPoint is the data-structure of this message
type _ContinuationPoint struct {
}

var _ ContinuationPoint = (*_ContinuationPoint)(nil)

// NewContinuationPoint factory function for _ContinuationPoint
func NewContinuationPoint() *_ContinuationPoint {
	return &_ContinuationPoint{}
}

// Deprecated: use the interface for direct cast
func CastContinuationPoint(structType any) ContinuationPoint {
	if casted, ok := structType.(ContinuationPoint); ok {
		return casted
	}
	if casted, ok := structType.(*ContinuationPoint); ok {
		return *casted
	}
	return nil
}

func (m *_ContinuationPoint) GetTypeName() string {
	return "ContinuationPoint"
}

func (m *_ContinuationPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ContinuationPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ContinuationPointParse(ctx context.Context, theBytes []byte) (ContinuationPoint, error) {
	return ContinuationPointParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ContinuationPointParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ContinuationPoint, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ContinuationPoint, error) {
		return ContinuationPointParseWithBuffer(ctx, readBuffer)
	}
}

func ContinuationPointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ContinuationPoint, error) {
	v, err := (&_ContinuationPoint{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ContinuationPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__continuationPoint ContinuationPoint, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContinuationPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContinuationPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ContinuationPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContinuationPoint")
	}

	return m, nil
}

func (m *_ContinuationPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContinuationPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ContinuationPoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ContinuationPoint")
	}

	if popErr := writeBuffer.PopContext("ContinuationPoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ContinuationPoint")
	}
	return nil
}

func (m *_ContinuationPoint) IsContinuationPoint() {}

func (m *_ContinuationPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ContinuationPoint) deepCopy() *_ContinuationPoint {
	if m == nil {
		return nil
	}
	_ContinuationPointCopy := &_ContinuationPoint{}
	return _ContinuationPointCopy
}

func (m *_ContinuationPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
