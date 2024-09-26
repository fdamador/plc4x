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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TunnelingResponse is the corresponding interface of TunnelingResponse
type TunnelingResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetTunnelingResponseDataBlock returns TunnelingResponseDataBlock (property field)
	GetTunnelingResponseDataBlock() TunnelingResponseDataBlock
	// IsTunnelingResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTunnelingResponse()
}

// _TunnelingResponse is the data-structure of this message
type _TunnelingResponse struct {
	KnxNetIpMessageContract
	TunnelingResponseDataBlock TunnelingResponseDataBlock
}

var _ TunnelingResponse = (*_TunnelingResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_TunnelingResponse)(nil)

// NewTunnelingResponse factory function for _TunnelingResponse
func NewTunnelingResponse(tunnelingResponseDataBlock TunnelingResponseDataBlock) *_TunnelingResponse {
	if tunnelingResponseDataBlock == nil {
		panic("tunnelingResponseDataBlock of type TunnelingResponseDataBlock for TunnelingResponse must not be nil")
	}
	_result := &_TunnelingResponse{
		KnxNetIpMessageContract:    NewKnxNetIpMessage(),
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TunnelingResponse) GetMsgType() uint16 {
	return 0x0421
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TunnelingResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingResponse) GetTunnelingResponseDataBlock() TunnelingResponseDataBlock {
	return m.TunnelingResponseDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTunnelingResponse(structType any) TunnelingResponse {
	if casted, ok := structType.(TunnelingResponse); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingResponse); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingResponse) GetTypeName() string {
	return "TunnelingResponse"
}

func (m *_TunnelingResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (tunnelingResponseDataBlock)
	lengthInBits += m.TunnelingResponseDataBlock.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TunnelingResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TunnelingResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__tunnelingResponse TunnelingResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tunnelingResponseDataBlock, err := ReadSimpleField[TunnelingResponseDataBlock](ctx, "tunnelingResponseDataBlock", ReadComplex[TunnelingResponseDataBlock](TunnelingResponseDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tunnelingResponseDataBlock' field"))
	}
	m.TunnelingResponseDataBlock = tunnelingResponseDataBlock

	if closeErr := readBuffer.CloseContext("TunnelingResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingResponse")
	}

	return m, nil
}

func (m *_TunnelingResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TunnelingResponse")
		}

		if err := WriteSimpleField[TunnelingResponseDataBlock](ctx, "tunnelingResponseDataBlock", m.GetTunnelingResponseDataBlock(), WriteComplex[TunnelingResponseDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'tunnelingResponseDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TunnelingResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TunnelingResponse) IsTunnelingResponse() {}

func (m *_TunnelingResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TunnelingResponse) deepCopy() *_TunnelingResponse {
	if m == nil {
		return nil
	}
	_TunnelingResponseCopy := &_TunnelingResponse{
		m.KnxNetIpMessageContract.DeepCopy().(KnxNetIpMessageContract),
		m.TunnelingResponseDataBlock.DeepCopy().(TunnelingResponseDataBlock),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _TunnelingResponseCopy
}

func (m *_TunnelingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
