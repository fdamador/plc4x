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

// BACnetPriorityValueObjectidentifier is the corresponding interface of BACnetPriorityValueObjectidentifier
type BACnetPriorityValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetPriorityValueObjectidentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueObjectidentifier()
}

// _BACnetPriorityValueObjectidentifier is the data-structure of this message
type _BACnetPriorityValueObjectidentifier struct {
	BACnetPriorityValueContract
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

var _ BACnetPriorityValueObjectidentifier = (*_BACnetPriorityValueObjectidentifier)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueObjectidentifier)(nil)

// NewBACnetPriorityValueObjectidentifier factory function for _BACnetPriorityValueObjectidentifier
func NewBACnetPriorityValueObjectidentifier(peekedTagHeader BACnetTagHeader, objectidentifierValue BACnetApplicationTagObjectIdentifier, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueObjectidentifier {
	if objectidentifierValue == nil {
		panic("objectidentifierValue of type BACnetApplicationTagObjectIdentifier for BACnetPriorityValueObjectidentifier must not be nil")
	}
	_result := &_BACnetPriorityValueObjectidentifier{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		ObjectidentifierValue:       objectidentifierValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
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

func (m *_BACnetPriorityValueObjectidentifier) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueObjectidentifier(structType any) BACnetPriorityValueObjectidentifier {
	if casted, ok := structType.(BACnetPriorityValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueObjectidentifier) GetTypeName() string {
	return "BACnetPriorityValueObjectidentifier"
}

func (m *_BACnetPriorityValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueObjectidentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueObjectidentifier BACnetPriorityValueObjectidentifier, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectidentifierValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectidentifierValue' field"))
	}
	m.ObjectidentifierValue = objectidentifierValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueObjectidentifier")
	}

	return m, nil
}

func (m *_BACnetPriorityValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueObjectidentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", m.GetObjectidentifierValue(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueObjectidentifier")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueObjectidentifier) IsBACnetPriorityValueObjectidentifier() {}

func (m *_BACnetPriorityValueObjectidentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueObjectidentifier) deepCopy() *_BACnetPriorityValueObjectidentifier {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueObjectidentifierCopy := &_BACnetPriorityValueObjectidentifier{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		m.ObjectidentifierValue.DeepCopy().(BACnetApplicationTagObjectIdentifier),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueObjectidentifierCopy
}

func (m *_BACnetPriorityValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
