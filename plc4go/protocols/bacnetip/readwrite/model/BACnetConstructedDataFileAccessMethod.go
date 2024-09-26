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

// BACnetConstructedDataFileAccessMethod is the corresponding interface of BACnetConstructedDataFileAccessMethod
type BACnetConstructedDataFileAccessMethod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFileAccessMethod returns FileAccessMethod (property field)
	GetFileAccessMethod() BACnetFileAccessMethodTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetFileAccessMethodTagged
	// IsBACnetConstructedDataFileAccessMethod is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFileAccessMethod()
}

// _BACnetConstructedDataFileAccessMethod is the data-structure of this message
type _BACnetConstructedDataFileAccessMethod struct {
	BACnetConstructedDataContract
	FileAccessMethod BACnetFileAccessMethodTagged
}

var _ BACnetConstructedDataFileAccessMethod = (*_BACnetConstructedDataFileAccessMethod)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFileAccessMethod)(nil)

// NewBACnetConstructedDataFileAccessMethod factory function for _BACnetConstructedDataFileAccessMethod
func NewBACnetConstructedDataFileAccessMethod(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, fileAccessMethod BACnetFileAccessMethodTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFileAccessMethod {
	if fileAccessMethod == nil {
		panic("fileAccessMethod of type BACnetFileAccessMethodTagged for BACnetConstructedDataFileAccessMethod must not be nil")
	}
	_result := &_BACnetConstructedDataFileAccessMethod{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FileAccessMethod:              fileAccessMethod,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFileAccessMethod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFileAccessMethod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FILE_ACCESS_METHOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFileAccessMethod) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFileAccessMethod) GetFileAccessMethod() BACnetFileAccessMethodTagged {
	return m.FileAccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFileAccessMethod) GetActualValue() BACnetFileAccessMethodTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetFileAccessMethodTagged(m.GetFileAccessMethod())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFileAccessMethod(structType any) BACnetConstructedDataFileAccessMethod {
	if casted, ok := structType.(BACnetConstructedDataFileAccessMethod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileAccessMethod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFileAccessMethod) GetTypeName() string {
	return "BACnetConstructedDataFileAccessMethod"
}

func (m *_BACnetConstructedDataFileAccessMethod) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (fileAccessMethod)
	lengthInBits += m.FileAccessMethod.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFileAccessMethod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFileAccessMethod) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFileAccessMethod BACnetConstructedDataFileAccessMethod, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFileAccessMethod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileAccessMethod, err := ReadSimpleField[BACnetFileAccessMethodTagged](ctx, "fileAccessMethod", ReadComplex[BACnetFileAccessMethodTagged](BACnetFileAccessMethodTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileAccessMethod' field"))
	}
	m.FileAccessMethod = fileAccessMethod

	actualValue, err := ReadVirtualField[BACnetFileAccessMethodTagged](ctx, "actualValue", (*BACnetFileAccessMethodTagged)(nil), fileAccessMethod)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFileAccessMethod")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFileAccessMethod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFileAccessMethod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFileAccessMethod")
		}

		if err := WriteSimpleField[BACnetFileAccessMethodTagged](ctx, "fileAccessMethod", m.GetFileAccessMethod(), WriteComplex[BACnetFileAccessMethodTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileAccessMethod' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFileAccessMethod")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFileAccessMethod) IsBACnetConstructedDataFileAccessMethod() {}

func (m *_BACnetConstructedDataFileAccessMethod) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFileAccessMethod) deepCopy() *_BACnetConstructedDataFileAccessMethod {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFileAccessMethodCopy := &_BACnetConstructedDataFileAccessMethod{
		m.BACnetConstructedDataContract.DeepCopy().(BACnetConstructedDataContract),
		m.FileAccessMethod.DeepCopy().(BACnetFileAccessMethodTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFileAccessMethodCopy
}

func (m *_BACnetConstructedDataFileAccessMethod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
