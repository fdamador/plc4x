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

// BACnetConstructedDataListOfGroupMembers is the corresponding interface of BACnetConstructedDataListOfGroupMembers
type BACnetConstructedDataListOfGroupMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetListOfGroupMembers returns ListOfGroupMembers (property field)
	GetListOfGroupMembers() []BACnetReadAccessSpecification
	// IsBACnetConstructedDataListOfGroupMembers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataListOfGroupMembers()
}

// _BACnetConstructedDataListOfGroupMembers is the data-structure of this message
type _BACnetConstructedDataListOfGroupMembers struct {
	BACnetConstructedDataContract
	ListOfGroupMembers []BACnetReadAccessSpecification
}

var _ BACnetConstructedDataListOfGroupMembers = (*_BACnetConstructedDataListOfGroupMembers)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataListOfGroupMembers)(nil)

// NewBACnetConstructedDataListOfGroupMembers factory function for _BACnetConstructedDataListOfGroupMembers
func NewBACnetConstructedDataListOfGroupMembers(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, listOfGroupMembers []BACnetReadAccessSpecification, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataListOfGroupMembers {
	_result := &_BACnetConstructedDataListOfGroupMembers{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ListOfGroupMembers:            listOfGroupMembers,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIST_OF_GROUP_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) GetListOfGroupMembers() []BACnetReadAccessSpecification {
	return m.ListOfGroupMembers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataListOfGroupMembers(structType any) BACnetConstructedDataListOfGroupMembers {
	if casted, ok := structType.(BACnetConstructedDataListOfGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataListOfGroupMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataListOfGroupMembers"
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ListOfGroupMembers) > 0 {
		for _, element := range m.ListOfGroupMembers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataListOfGroupMembers) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataListOfGroupMembers BACnetConstructedDataListOfGroupMembers, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataListOfGroupMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataListOfGroupMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	listOfGroupMembers, err := ReadTerminatedArrayField[BACnetReadAccessSpecification](ctx, "listOfGroupMembers", ReadComplex[BACnetReadAccessSpecification](BACnetReadAccessSpecificationParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfGroupMembers' field"))
	}
	m.ListOfGroupMembers = listOfGroupMembers

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataListOfGroupMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataListOfGroupMembers")
	}

	return m, nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataListOfGroupMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataListOfGroupMembers")
		}

		if err := WriteComplexTypeArrayField(ctx, "listOfGroupMembers", m.GetListOfGroupMembers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfGroupMembers' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataListOfGroupMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataListOfGroupMembers")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataListOfGroupMembers) IsBACnetConstructedDataListOfGroupMembers() {}

func (m *_BACnetConstructedDataListOfGroupMembers) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataListOfGroupMembers) deepCopy() *_BACnetConstructedDataListOfGroupMembers {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataListOfGroupMembersCopy := &_BACnetConstructedDataListOfGroupMembers{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetReadAccessSpecification, BACnetReadAccessSpecification](m.ListOfGroupMembers),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataListOfGroupMembersCopy
}

func (m *_BACnetConstructedDataListOfGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
