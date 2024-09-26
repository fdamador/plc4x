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

// CreateMonitoredItemsRequest is the corresponding interface of CreateMonitoredItemsRequest
type CreateMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetNoOfItemsToCreate returns NoOfItemsToCreate (property field)
	GetNoOfItemsToCreate() int32
	// GetItemsToCreate returns ItemsToCreate (property field)
	GetItemsToCreate() []ExtensionObjectDefinition
	// IsCreateMonitoredItemsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateMonitoredItemsRequest()
}

// _CreateMonitoredItemsRequest is the data-structure of this message
type _CreateMonitoredItemsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader      ExtensionObjectDefinition
	SubscriptionId     uint32
	TimestampsToReturn TimestampsToReturn
	NoOfItemsToCreate  int32
	ItemsToCreate      []ExtensionObjectDefinition
}

var _ CreateMonitoredItemsRequest = (*_CreateMonitoredItemsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateMonitoredItemsRequest)(nil)

// NewCreateMonitoredItemsRequest factory function for _CreateMonitoredItemsRequest
func NewCreateMonitoredItemsRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, timestampsToReturn TimestampsToReturn, noOfItemsToCreate int32, itemsToCreate []ExtensionObjectDefinition) *_CreateMonitoredItemsRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for CreateMonitoredItemsRequest must not be nil")
	}
	_result := &_CreateMonitoredItemsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		TimestampsToReturn:                timestampsToReturn,
		NoOfItemsToCreate:                 noOfItemsToCreate,
		ItemsToCreate:                     itemsToCreate,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetIdentifier() string {
	return "751"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateMonitoredItemsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_CreateMonitoredItemsRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_CreateMonitoredItemsRequest) GetNoOfItemsToCreate() int32 {
	return m.NoOfItemsToCreate
}

func (m *_CreateMonitoredItemsRequest) GetItemsToCreate() []ExtensionObjectDefinition {
	return m.ItemsToCreate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateMonitoredItemsRequest(structType any) CreateMonitoredItemsRequest {
	if casted, ok := structType.(CreateMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateMonitoredItemsRequest) GetTypeName() string {
	return "CreateMonitoredItemsRequest"
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Simple field (noOfItemsToCreate)
	lengthInBits += 32

	// Array field
	if len(m.ItemsToCreate) > 0 {
		for _curItem, element := range m.ItemsToCreate {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ItemsToCreate), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateMonitoredItemsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createMonitoredItemsRequest CreateMonitoredItemsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}
	m.TimestampsToReturn = timestampsToReturn

	noOfItemsToCreate, err := ReadSimpleField(ctx, "noOfItemsToCreate", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfItemsToCreate' field"))
	}
	m.NoOfItemsToCreate = noOfItemsToCreate

	itemsToCreate, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "itemsToCreate", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("745")), readBuffer), uint64(noOfItemsToCreate))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsToCreate' field"))
	}
	m.ItemsToCreate = itemsToCreate

	if closeErr := readBuffer.CloseContext("CreateMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateMonitoredItemsRequest")
	}

	return m, nil
}

func (m *_CreateMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateMonitoredItemsRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", m.GetTimestampsToReturn(), WriteEnum[TimestampsToReturn, uint32](TimestampsToReturn.GetValue, TimestampsToReturn.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'timestampsToReturn' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfItemsToCreate", m.GetNoOfItemsToCreate(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfItemsToCreate' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "itemsToCreate", m.GetItemsToCreate(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsToCreate' field")
		}

		if popErr := writeBuffer.PopContext("CreateMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateMonitoredItemsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateMonitoredItemsRequest) IsCreateMonitoredItemsRequest() {}

func (m *_CreateMonitoredItemsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CreateMonitoredItemsRequest) deepCopy() *_CreateMonitoredItemsRequest {
	if m == nil {
		return nil
	}
	_CreateMonitoredItemsRequestCopy := &_CreateMonitoredItemsRequest{
		m.ExtensionObjectDefinitionContract.DeepCopy().(ExtensionObjectDefinitionContract),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.SubscriptionId,
		m.TimestampsToReturn,
		m.NoOfItemsToCreate,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.ItemsToCreate),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CreateMonitoredItemsRequestCopy
}

func (m *_CreateMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
