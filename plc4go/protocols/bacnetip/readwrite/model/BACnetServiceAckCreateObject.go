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

// BACnetServiceAckCreateObject is the corresponding interface of BACnetServiceAckCreateObject
type BACnetServiceAckCreateObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// IsBACnetServiceAckCreateObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckCreateObject()
}

// _BACnetServiceAckCreateObject is the data-structure of this message
type _BACnetServiceAckCreateObject struct {
	BACnetServiceAckContract
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
}

var _ BACnetServiceAckCreateObject = (*_BACnetServiceAckCreateObject)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckCreateObject)(nil)

// NewBACnetServiceAckCreateObject factory function for _BACnetServiceAckCreateObject
func NewBACnetServiceAckCreateObject(objectIdentifier BACnetApplicationTagObjectIdentifier, serviceAckLength uint32) *_BACnetServiceAckCreateObject {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetServiceAckCreateObject must not be nil")
	}
	_result := &_BACnetServiceAckCreateObject{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ObjectIdentifier:         objectIdentifier,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckCreateObject) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckCreateObject(structType any) BACnetServiceAckCreateObject {
	if casted, ok := structType.(BACnetServiceAckCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckCreateObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckCreateObject) GetTypeName() string {
	return "BACnetServiceAckCreateObject"
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckCreateObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckCreateObject BACnetServiceAckCreateObject, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckCreateObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckCreateObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	if closeErr := readBuffer.CloseContext("BACnetServiceAckCreateObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckCreateObject")
	}

	return m, nil
}

func (m *_BACnetServiceAckCreateObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckCreateObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckCreateObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckCreateObject")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckCreateObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckCreateObject")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckCreateObject) IsBACnetServiceAckCreateObject() {}

func (m *_BACnetServiceAckCreateObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckCreateObject) deepCopy() *_BACnetServiceAckCreateObject {
	if m == nil {
		return nil
	}
	_BACnetServiceAckCreateObjectCopy := &_BACnetServiceAckCreateObject{
		m.BACnetServiceAckContract.DeepCopy().(BACnetServiceAckContract),
		m.ObjectIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
	}
	m.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckCreateObjectCopy
}

func (m *_BACnetServiceAckCreateObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
