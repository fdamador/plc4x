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

// BACnetPropertyStatesFileAccessMethod is the corresponding interface of BACnetPropertyStatesFileAccessMethod
type BACnetPropertyStatesFileAccessMethod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetFileAccessMethod returns FileAccessMethod (property field)
	GetFileAccessMethod() BACnetFileAccessMethodTagged
	// IsBACnetPropertyStatesFileAccessMethod is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesFileAccessMethod()
}

// _BACnetPropertyStatesFileAccessMethod is the data-structure of this message
type _BACnetPropertyStatesFileAccessMethod struct {
	BACnetPropertyStatesContract
	FileAccessMethod BACnetFileAccessMethodTagged
}

var _ BACnetPropertyStatesFileAccessMethod = (*_BACnetPropertyStatesFileAccessMethod)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesFileAccessMethod)(nil)

// NewBACnetPropertyStatesFileAccessMethod factory function for _BACnetPropertyStatesFileAccessMethod
func NewBACnetPropertyStatesFileAccessMethod(peekedTagHeader BACnetTagHeader, fileAccessMethod BACnetFileAccessMethodTagged) *_BACnetPropertyStatesFileAccessMethod {
	if fileAccessMethod == nil {
		panic("fileAccessMethod of type BACnetFileAccessMethodTagged for BACnetPropertyStatesFileAccessMethod must not be nil")
	}
	_result := &_BACnetPropertyStatesFileAccessMethod{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		FileAccessMethod:             fileAccessMethod,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
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

func (m *_BACnetPropertyStatesFileAccessMethod) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesFileAccessMethod) GetFileAccessMethod() BACnetFileAccessMethodTagged {
	return m.FileAccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesFileAccessMethod(structType any) BACnetPropertyStatesFileAccessMethod {
	if casted, ok := structType.(BACnetPropertyStatesFileAccessMethod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesFileAccessMethod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetTypeName() string {
	return "BACnetPropertyStatesFileAccessMethod"
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (fileAccessMethod)
	lengthInBits += m.FileAccessMethod.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesFileAccessMethod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesFileAccessMethod) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesFileAccessMethod BACnetPropertyStatesFileAccessMethod, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesFileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesFileAccessMethod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileAccessMethod, err := ReadSimpleField[BACnetFileAccessMethodTagged](ctx, "fileAccessMethod", ReadComplex[BACnetFileAccessMethodTagged](BACnetFileAccessMethodTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileAccessMethod' field"))
	}
	m.FileAccessMethod = fileAccessMethod

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesFileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesFileAccessMethod")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesFileAccessMethod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesFileAccessMethod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesFileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesFileAccessMethod")
		}

		if err := WriteSimpleField[BACnetFileAccessMethodTagged](ctx, "fileAccessMethod", m.GetFileAccessMethod(), WriteComplex[BACnetFileAccessMethodTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileAccessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesFileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesFileAccessMethod")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesFileAccessMethod) IsBACnetPropertyStatesFileAccessMethod() {}

func (m *_BACnetPropertyStatesFileAccessMethod) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesFileAccessMethod) deepCopy() *_BACnetPropertyStatesFileAccessMethod {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesFileAccessMethodCopy := &_BACnetPropertyStatesFileAccessMethod{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.FileAccessMethod.DeepCopy().(BACnetFileAccessMethodTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesFileAccessMethodCopy
}

func (m *_BACnetPropertyStatesFileAccessMethod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
