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

// OpenChannelMessage is the corresponding interface of OpenChannelMessage
type OpenChannelMessage interface {
	OpenChannelMessageContract
	OpenChannelMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsOpenChannelMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpenChannelMessage()
}

// OpenChannelMessageContract provides a set of functions which can be overwritten by a sub struct
type OpenChannelMessageContract interface {
	utils.Copyable
	// IsOpenChannelMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpenChannelMessage()
}

// OpenChannelMessageRequirements provides a set of functions which need to be implemented by a sub struct
type OpenChannelMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// _OpenChannelMessage is the data-structure of this message
type _OpenChannelMessage struct {
	_SubType OpenChannelMessage
}

var _ OpenChannelMessageContract = (*_OpenChannelMessage)(nil)

// NewOpenChannelMessage factory function for _OpenChannelMessage
func NewOpenChannelMessage() *_OpenChannelMessage {
	return &_OpenChannelMessage{}
}

// Deprecated: use the interface for direct cast
func CastOpenChannelMessage(structType any) OpenChannelMessage {
	if casted, ok := structType.(OpenChannelMessage); ok {
		return casted
	}
	if casted, ok := structType.(*OpenChannelMessage); ok {
		return *casted
	}
	return nil
}

func (m *_OpenChannelMessage) GetTypeName() string {
	return "OpenChannelMessage"
}

func (m *_OpenChannelMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_OpenChannelMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func OpenChannelMessageParse[T OpenChannelMessage](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return OpenChannelMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpenChannelMessageParseWithBufferProducer[T OpenChannelMessage](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := OpenChannelMessageParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func OpenChannelMessageParseWithBuffer[T OpenChannelMessage](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_OpenChannelMessage{}).parse(ctx, readBuffer, response)
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

func (m *_OpenChannelMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__openChannelMessage OpenChannelMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpenChannelMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenChannelMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child OpenChannelMessage
	switch {
	case response == bool(false): // OpenChannelMessageRequest
		if _child, err = new(_OpenChannelMessageRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpenChannelMessageRequest for type-switch of OpenChannelMessage")
		}
	case response == bool(true): // OpenChannelMessageResponse
		if _child, err = new(_OpenChannelMessageResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpenChannelMessageResponse for type-switch of OpenChannelMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [response=%v]", response)
	}

	if closeErr := readBuffer.CloseContext("OpenChannelMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenChannelMessage")
	}

	return _child, nil
}

func (pm *_OpenChannelMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child OpenChannelMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("OpenChannelMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for OpenChannelMessage")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("OpenChannelMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for OpenChannelMessage")
	}
	return nil
}

func (m *_OpenChannelMessage) IsOpenChannelMessage() {}

func (m *_OpenChannelMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpenChannelMessage) deepCopy() *_OpenChannelMessage {
	if m == nil {
		return nil
	}
	_OpenChannelMessageCopy := &_OpenChannelMessage{
		nil, // will be set by child
	}
	return _OpenChannelMessageCopy
}
