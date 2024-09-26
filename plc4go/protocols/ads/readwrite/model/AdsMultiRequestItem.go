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

// AdsMultiRequestItem is the corresponding interface of AdsMultiRequestItem
type AdsMultiRequestItem interface {
	AdsMultiRequestItemContract
	AdsMultiRequestItemRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsAdsMultiRequestItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsMultiRequestItem()
}

// AdsMultiRequestItemContract provides a set of functions which can be overwritten by a sub struct
type AdsMultiRequestItemContract interface {
	utils.Copyable
	// IsAdsMultiRequestItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsMultiRequestItem()
}

// AdsMultiRequestItemRequirements provides a set of functions which need to be implemented by a sub struct
type AdsMultiRequestItemRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIndexGroup returns IndexGroup (discriminator field)
	GetIndexGroup() uint32
}

// _AdsMultiRequestItem is the data-structure of this message
type _AdsMultiRequestItem struct {
	_SubType AdsMultiRequestItem
}

var _ AdsMultiRequestItemContract = (*_AdsMultiRequestItem)(nil)

// NewAdsMultiRequestItem factory function for _AdsMultiRequestItem
func NewAdsMultiRequestItem() *_AdsMultiRequestItem {
	return &_AdsMultiRequestItem{}
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItem(structType any) AdsMultiRequestItem {
	if casted, ok := structType.(AdsMultiRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItem); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *_AdsMultiRequestItem) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_AdsMultiRequestItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func AdsMultiRequestItemParse[T AdsMultiRequestItem](ctx context.Context, theBytes []byte, indexGroup uint32) (T, error) {
	return AdsMultiRequestItemParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), indexGroup)
}

func AdsMultiRequestItemParseWithBufferProducer[T AdsMultiRequestItem](indexGroup uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := AdsMultiRequestItemParseWithBuffer[T](ctx, readBuffer, indexGroup)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func AdsMultiRequestItemParseWithBuffer[T AdsMultiRequestItem](ctx context.Context, readBuffer utils.ReadBuffer, indexGroup uint32) (T, error) {
	v, err := (&_AdsMultiRequestItem{}).parse(ctx, readBuffer, indexGroup)
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

func (m *_AdsMultiRequestItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, indexGroup uint32) (__adsMultiRequestItem AdsMultiRequestItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child AdsMultiRequestItem
	switch {
	case indexGroup == uint32(61568): // AdsMultiRequestItemRead
		if _child, err = new(_AdsMultiRequestItemRead).parse(ctx, readBuffer, m, indexGroup); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsMultiRequestItemRead for type-switch of AdsMultiRequestItem")
		}
	case indexGroup == uint32(61569): // AdsMultiRequestItemWrite
		if _child, err = new(_AdsMultiRequestItemWrite).parse(ctx, readBuffer, m, indexGroup); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsMultiRequestItemWrite for type-switch of AdsMultiRequestItem")
		}
	case indexGroup == uint32(61570): // AdsMultiRequestItemReadWrite
		if _child, err = new(_AdsMultiRequestItemReadWrite).parse(ctx, readBuffer, m, indexGroup); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsMultiRequestItemReadWrite for type-switch of AdsMultiRequestItem")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [indexGroup=%v]", indexGroup)
	}

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItem")
	}

	return _child, nil
}

func (pm *_AdsMultiRequestItem) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AdsMultiRequestItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItem")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsMultiRequestItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsMultiRequestItem")
	}
	return nil
}

func (m *_AdsMultiRequestItem) IsAdsMultiRequestItem() {}

func (m *_AdsMultiRequestItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsMultiRequestItem) deepCopy() *_AdsMultiRequestItem {
	if m == nil {
		return nil
	}
	_AdsMultiRequestItemCopy := &_AdsMultiRequestItem{
		nil, // will be set by child
	}
	return _AdsMultiRequestItemCopy
}
