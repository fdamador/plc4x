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

// AnonymousIdentityToken is the corresponding interface of AnonymousIdentityToken
type AnonymousIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	UserIdentityTokenDefinition
	// IsAnonymousIdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAnonymousIdentityToken()
}

// _AnonymousIdentityToken is the data-structure of this message
type _AnonymousIdentityToken struct {
	UserIdentityTokenDefinitionContract
}

var _ AnonymousIdentityToken = (*_AnonymousIdentityToken)(nil)
var _ UserIdentityTokenDefinitionRequirements = (*_AnonymousIdentityToken)(nil)

// NewAnonymousIdentityToken factory function for _AnonymousIdentityToken
func NewAnonymousIdentityToken() *_AnonymousIdentityToken {
	_result := &_AnonymousIdentityToken{
		UserIdentityTokenDefinitionContract: NewUserIdentityTokenDefinition(),
	}
	_result.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AnonymousIdentityToken) GetIdentifier() string {
	return "anonymous"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AnonymousIdentityToken) GetParent() UserIdentityTokenDefinitionContract {
	return m.UserIdentityTokenDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastAnonymousIdentityToken(structType any) AnonymousIdentityToken {
	if casted, ok := structType.(AnonymousIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*AnonymousIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_AnonymousIdentityToken) GetTypeName() string {
	return "AnonymousIdentityToken"
}

func (m *_AnonymousIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AnonymousIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AnonymousIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_UserIdentityTokenDefinition, identifier string) (__anonymousIdentityToken AnonymousIdentityToken, err error) {
	m.UserIdentityTokenDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AnonymousIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AnonymousIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AnonymousIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AnonymousIdentityToken")
	}

	return m, nil
}

func (m *_AnonymousIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AnonymousIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AnonymousIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AnonymousIdentityToken")
		}

		if popErr := writeBuffer.PopContext("AnonymousIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AnonymousIdentityToken")
		}
		return nil
	}
	return m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AnonymousIdentityToken) IsAnonymousIdentityToken() {}

func (m *_AnonymousIdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AnonymousIdentityToken) deepCopy() *_AnonymousIdentityToken {
	if m == nil {
		return nil
	}
	_AnonymousIdentityTokenCopy := &_AnonymousIdentityToken{
		m.UserIdentityTokenDefinitionContract.DeepCopy().(UserIdentityTokenDefinitionContract),
	}
	m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition)._SubType = m
	return _AnonymousIdentityTokenCopy
}

func (m *_AnonymousIdentityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
