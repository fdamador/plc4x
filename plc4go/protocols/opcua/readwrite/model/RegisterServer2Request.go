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

// RegisterServer2Request is the corresponding interface of RegisterServer2Request
type RegisterServer2Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetServer returns Server (property field)
	GetServer() ExtensionObjectDefinition
	// GetNoOfDiscoveryConfiguration returns NoOfDiscoveryConfiguration (property field)
	GetNoOfDiscoveryConfiguration() int32
	// GetDiscoveryConfiguration returns DiscoveryConfiguration (property field)
	GetDiscoveryConfiguration() []ExtensionObject
	// IsRegisterServer2Request is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRegisterServer2Request()
}

// _RegisterServer2Request is the data-structure of this message
type _RegisterServer2Request struct {
	ExtensionObjectDefinitionContract
	RequestHeader              ExtensionObjectDefinition
	Server                     ExtensionObjectDefinition
	NoOfDiscoveryConfiguration int32
	DiscoveryConfiguration     []ExtensionObject
}

var _ RegisterServer2Request = (*_RegisterServer2Request)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterServer2Request)(nil)

// NewRegisterServer2Request factory function for _RegisterServer2Request
func NewRegisterServer2Request(requestHeader ExtensionObjectDefinition, server ExtensionObjectDefinition, noOfDiscoveryConfiguration int32, discoveryConfiguration []ExtensionObject) *_RegisterServer2Request {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for RegisterServer2Request must not be nil")
	}
	if server == nil {
		panic("server of type ExtensionObjectDefinition for RegisterServer2Request must not be nil")
	}
	_result := &_RegisterServer2Request{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		Server:                            server,
		NoOfDiscoveryConfiguration:        noOfDiscoveryConfiguration,
		DiscoveryConfiguration:            discoveryConfiguration,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterServer2Request) GetIdentifier() string {
	return "12195"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterServer2Request) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterServer2Request) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_RegisterServer2Request) GetServer() ExtensionObjectDefinition {
	return m.Server
}

func (m *_RegisterServer2Request) GetNoOfDiscoveryConfiguration() int32 {
	return m.NoOfDiscoveryConfiguration
}

func (m *_RegisterServer2Request) GetDiscoveryConfiguration() []ExtensionObject {
	return m.DiscoveryConfiguration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRegisterServer2Request(structType any) RegisterServer2Request {
	if casted, ok := structType.(RegisterServer2Request); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterServer2Request); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterServer2Request) GetTypeName() string {
	return "RegisterServer2Request"
}

func (m *_RegisterServer2Request) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (server)
	lengthInBits += m.Server.GetLengthInBits(ctx)

	// Simple field (noOfDiscoveryConfiguration)
	lengthInBits += 32

	// Array field
	if len(m.DiscoveryConfiguration) > 0 {
		for _curItem, element := range m.DiscoveryConfiguration {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiscoveryConfiguration), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterServer2Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterServer2Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__registerServer2Request RegisterServer2Request, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterServer2Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServer2Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	server, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "server", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("434")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}
	m.Server = server

	noOfDiscoveryConfiguration, err := ReadSimpleField(ctx, "noOfDiscoveryConfiguration", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiscoveryConfiguration' field"))
	}
	m.NoOfDiscoveryConfiguration = noOfDiscoveryConfiguration

	discoveryConfiguration, err := ReadCountArrayField[ExtensionObject](ctx, "discoveryConfiguration", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer), uint64(noOfDiscoveryConfiguration))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryConfiguration' field"))
	}
	m.DiscoveryConfiguration = discoveryConfiguration

	if closeErr := readBuffer.CloseContext("RegisterServer2Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServer2Request")
	}

	return m, nil
}

func (m *_RegisterServer2Request) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterServer2Request) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterServer2Request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterServer2Request")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "server", m.GetServer(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'server' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiscoveryConfiguration", m.GetNoOfDiscoveryConfiguration(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiscoveryConfiguration' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "discoveryConfiguration", m.GetDiscoveryConfiguration(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'discoveryConfiguration' field")
		}

		if popErr := writeBuffer.PopContext("RegisterServer2Request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServer2Request")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterServer2Request) IsRegisterServer2Request() {}

func (m *_RegisterServer2Request) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RegisterServer2Request) deepCopy() *_RegisterServer2Request {
	if m == nil {
		return nil
	}
	_RegisterServer2RequestCopy := &_RegisterServer2Request{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.Server.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfDiscoveryConfiguration,
		utils.DeepCopySlice[ExtensionObject, ExtensionObject](m.DiscoveryConfiguration),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RegisterServer2RequestCopy
}

func (m *_RegisterServer2Request) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
