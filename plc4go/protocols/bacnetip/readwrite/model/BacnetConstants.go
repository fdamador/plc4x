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

// Constant values.
const BacnetConstants_BACNETUDPDEFAULTPORT uint16 = uint16(47808)

// BacnetConstants is the corresponding interface of BacnetConstants
type BacnetConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBacnetConstants is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBacnetConstants()
	// CreateBuilder creates a BacnetConstantsBuilder
	CreateBacnetConstantsBuilder() BacnetConstantsBuilder
}

// _BacnetConstants is the data-structure of this message
type _BacnetConstants struct {
}

var _ BacnetConstants = (*_BacnetConstants)(nil)

// NewBacnetConstants factory function for _BacnetConstants
func NewBacnetConstants() *_BacnetConstants {
	return &_BacnetConstants{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BacnetConstantsBuilder is a builder for BacnetConstants
type BacnetConstantsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BacnetConstantsBuilder
	// Build builds the BacnetConstants or returns an error if something is wrong
	Build() (BacnetConstants, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BacnetConstants
}

// NewBacnetConstantsBuilder() creates a BacnetConstantsBuilder
func NewBacnetConstantsBuilder() BacnetConstantsBuilder {
	return &_BacnetConstantsBuilder{_BacnetConstants: new(_BacnetConstants)}
}

type _BacnetConstantsBuilder struct {
	*_BacnetConstants

	err *utils.MultiError
}

var _ (BacnetConstantsBuilder) = (*_BacnetConstantsBuilder)(nil)

func (b *_BacnetConstantsBuilder) WithMandatoryFields() BacnetConstantsBuilder {
	return b
}

func (b *_BacnetConstantsBuilder) Build() (BacnetConstants, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BacnetConstants.deepCopy(), nil
}

func (b *_BacnetConstantsBuilder) MustBuild() BacnetConstants {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BacnetConstantsBuilder) DeepCopy() any {
	_copy := b.CreateBacnetConstantsBuilder().(*_BacnetConstantsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBacnetConstantsBuilder creates a BacnetConstantsBuilder
func (b *_BacnetConstants) CreateBacnetConstantsBuilder() BacnetConstantsBuilder {
	if b == nil {
		return NewBacnetConstantsBuilder()
	}
	return &_BacnetConstantsBuilder{_BacnetConstants: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_BacnetConstants) GetBacnetUdpDefaultPort() uint16 {
	return BacnetConstants_BACNETUDPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBacnetConstants(structType any) BacnetConstants {
	if casted, ok := structType.(BacnetConstants); ok {
		return casted
	}
	if casted, ok := structType.(*BacnetConstants); ok {
		return *casted
	}
	return nil
}

func (m *_BacnetConstants) GetTypeName() string {
	return "BacnetConstants"
}

func (m *_BacnetConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetUdpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BacnetConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BacnetConstantsParse(ctx context.Context, theBytes []byte) (BacnetConstants, error) {
	return BacnetConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BacnetConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BacnetConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BacnetConstants, error) {
		return BacnetConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func BacnetConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BacnetConstants, error) {
	v, err := (&_BacnetConstants{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BacnetConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bacnetConstants BacnetConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BacnetConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BacnetConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetUdpDefaultPort, err := ReadConstField[uint16](ctx, "bacnetUdpDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), BacnetConstants_BACNETUDPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetUdpDefaultPort' field"))
	}
	_ = bacnetUdpDefaultPort

	if closeErr := readBuffer.CloseContext("BacnetConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BacnetConstants")
	}

	return m, nil
}

func (m *_BacnetConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BacnetConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BacnetConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BacnetConstants")
	}

	if err := WriteConstField(ctx, "bacnetUdpDefaultPort", BacnetConstants_BACNETUDPDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'bacnetUdpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("BacnetConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BacnetConstants")
	}
	return nil
}

func (m *_BacnetConstants) IsBacnetConstants() {}

func (m *_BacnetConstants) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BacnetConstants) deepCopy() *_BacnetConstants {
	if m == nil {
		return nil
	}
	_BacnetConstantsCopy := &_BacnetConstants{}
	return _BacnetConstantsCopy
}

func (m *_BacnetConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
