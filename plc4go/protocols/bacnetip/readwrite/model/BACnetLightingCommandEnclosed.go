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

// BACnetLightingCommandEnclosed is the corresponding interface of BACnetLightingCommandEnclosed
type BACnetLightingCommandEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() BACnetLightingCommand
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetLightingCommandEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLightingCommandEnclosed()
	// CreateBuilder creates a BACnetLightingCommandEnclosedBuilder
	CreateBACnetLightingCommandEnclosedBuilder() BACnetLightingCommandEnclosedBuilder
}

// _BACnetLightingCommandEnclosed is the data-structure of this message
type _BACnetLightingCommandEnclosed struct {
	OpeningTag      BACnetOpeningTag
	LightingCommand BACnetLightingCommand
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetLightingCommandEnclosed = (*_BACnetLightingCommandEnclosed)(nil)

// NewBACnetLightingCommandEnclosed factory function for _BACnetLightingCommandEnclosed
func NewBACnetLightingCommandEnclosed(openingTag BACnetOpeningTag, lightingCommand BACnetLightingCommand, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLightingCommandEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetLightingCommandEnclosed must not be nil")
	}
	if lightingCommand == nil {
		panic("lightingCommand of type BACnetLightingCommand for BACnetLightingCommandEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetLightingCommandEnclosed must not be nil")
	}
	return &_BACnetLightingCommandEnclosed{OpeningTag: openingTag, LightingCommand: lightingCommand, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLightingCommandEnclosedBuilder is a builder for BACnetLightingCommandEnclosed
type BACnetLightingCommandEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, lightingCommand BACnetLightingCommand, closingTag BACnetClosingTag) BACnetLightingCommandEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetLightingCommandEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLightingCommandEnclosedBuilder
	// WithLightingCommand adds LightingCommand (property field)
	WithLightingCommand(BACnetLightingCommand) BACnetLightingCommandEnclosedBuilder
	// WithLightingCommandBuilder adds LightingCommand (property field) which is build by the builder
	WithLightingCommandBuilder(func(BACnetLightingCommandBuilder) BACnetLightingCommandBuilder) BACnetLightingCommandEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetLightingCommandEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLightingCommandEnclosedBuilder
	// Build builds the BACnetLightingCommandEnclosed or returns an error if something is wrong
	Build() (BACnetLightingCommandEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLightingCommandEnclosed
}

// NewBACnetLightingCommandEnclosedBuilder() creates a BACnetLightingCommandEnclosedBuilder
func NewBACnetLightingCommandEnclosedBuilder() BACnetLightingCommandEnclosedBuilder {
	return &_BACnetLightingCommandEnclosedBuilder{_BACnetLightingCommandEnclosed: new(_BACnetLightingCommandEnclosed)}
}

type _BACnetLightingCommandEnclosedBuilder struct {
	*_BACnetLightingCommandEnclosed

	err *utils.MultiError
}

var _ (BACnetLightingCommandEnclosedBuilder) = (*_BACnetLightingCommandEnclosedBuilder)(nil)

func (b *_BACnetLightingCommandEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, lightingCommand BACnetLightingCommand, closingTag BACnetClosingTag) BACnetLightingCommandEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithLightingCommand(lightingCommand).WithClosingTag(closingTag)
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetLightingCommandEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLightingCommandEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithLightingCommand(lightingCommand BACnetLightingCommand) BACnetLightingCommandEnclosedBuilder {
	b.LightingCommand = lightingCommand
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithLightingCommandBuilder(builderSupplier func(BACnetLightingCommandBuilder) BACnetLightingCommandBuilder) BACnetLightingCommandEnclosedBuilder {
	builder := builderSupplier(b.LightingCommand.CreateBACnetLightingCommandBuilder())
	var err error
	b.LightingCommand, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingCommandBuilder failed"))
	}
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetLightingCommandEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLightingCommandEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetLightingCommandEnclosedBuilder) Build() (BACnetLightingCommandEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.LightingCommand == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lightingCommand' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLightingCommandEnclosed.deepCopy(), nil
}

func (b *_BACnetLightingCommandEnclosedBuilder) MustBuild() BACnetLightingCommandEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLightingCommandEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLightingCommandEnclosedBuilder().(*_BACnetLightingCommandEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLightingCommandEnclosedBuilder creates a BACnetLightingCommandEnclosedBuilder
func (b *_BACnetLightingCommandEnclosed) CreateBACnetLightingCommandEnclosedBuilder() BACnetLightingCommandEnclosedBuilder {
	if b == nil {
		return NewBACnetLightingCommandEnclosedBuilder()
	}
	return &_BACnetLightingCommandEnclosedBuilder{_BACnetLightingCommandEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingCommandEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLightingCommandEnclosed) GetLightingCommand() BACnetLightingCommand {
	return m.LightingCommand
}

func (m *_BACnetLightingCommandEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLightingCommandEnclosed(structType any) BACnetLightingCommandEnclosed {
	if casted, ok := structType.(BACnetLightingCommandEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingCommandEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingCommandEnclosed) GetTypeName() string {
	return "BACnetLightingCommandEnclosed"
}

func (m *_BACnetLightingCommandEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLightingCommandEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingCommandEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLightingCommandEnclosed, error) {
	return BACnetLightingCommandEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLightingCommandEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommandEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommandEnclosed, error) {
		return BACnetLightingCommandEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetLightingCommandEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLightingCommandEnclosed, error) {
	v, err := (&_BACnetLightingCommandEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLightingCommandEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetLightingCommandEnclosed BACnetLightingCommandEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingCommandEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingCommandEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	lightingCommand, err := ReadSimpleField[BACnetLightingCommand](ctx, "lightingCommand", ReadComplex[BACnetLightingCommand](BACnetLightingCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingCommand' field"))
	}
	m.LightingCommand = lightingCommand

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetLightingCommandEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingCommandEnclosed")
	}

	return m, nil
}

func (m *_BACnetLightingCommandEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingCommandEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingCommandEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingCommandEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetLightingCommand](ctx, "lightingCommand", m.GetLightingCommand(), WriteComplex[BACnetLightingCommand](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'lightingCommand' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommandEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingCommandEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingCommandEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLightingCommandEnclosed) IsBACnetLightingCommandEnclosed() {}

func (m *_BACnetLightingCommandEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLightingCommandEnclosed) deepCopy() *_BACnetLightingCommandEnclosed {
	if m == nil {
		return nil
	}
	_BACnetLightingCommandEnclosedCopy := &_BACnetLightingCommandEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.LightingCommand.DeepCopy().(BACnetLightingCommand),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetLightingCommandEnclosedCopy
}

func (m *_BACnetLightingCommandEnclosed) String() string {
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
