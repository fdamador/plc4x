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

// BACnetPropertyStatesShedState is the corresponding interface of BACnetPropertyStatesShedState
type BACnetPropertyStatesShedState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetShedState returns ShedState (property field)
	GetShedState() BACnetShedStateTagged
	// IsBACnetPropertyStatesShedState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesShedState()
	// CreateBuilder creates a BACnetPropertyStatesShedStateBuilder
	CreateBACnetPropertyStatesShedStateBuilder() BACnetPropertyStatesShedStateBuilder
}

// _BACnetPropertyStatesShedState is the data-structure of this message
type _BACnetPropertyStatesShedState struct {
	BACnetPropertyStatesContract
	ShedState BACnetShedStateTagged
}

var _ BACnetPropertyStatesShedState = (*_BACnetPropertyStatesShedState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesShedState)(nil)

// NewBACnetPropertyStatesShedState factory function for _BACnetPropertyStatesShedState
func NewBACnetPropertyStatesShedState(peekedTagHeader BACnetTagHeader, shedState BACnetShedStateTagged) *_BACnetPropertyStatesShedState {
	if shedState == nil {
		panic("shedState of type BACnetShedStateTagged for BACnetPropertyStatesShedState must not be nil")
	}
	_result := &_BACnetPropertyStatesShedState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ShedState:                    shedState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesShedStateBuilder is a builder for BACnetPropertyStatesShedState
type BACnetPropertyStatesShedStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(shedState BACnetShedStateTagged) BACnetPropertyStatesShedStateBuilder
	// WithShedState adds ShedState (property field)
	WithShedState(BACnetShedStateTagged) BACnetPropertyStatesShedStateBuilder
	// WithShedStateBuilder adds ShedState (property field) which is build by the builder
	WithShedStateBuilder(func(BACnetShedStateTaggedBuilder) BACnetShedStateTaggedBuilder) BACnetPropertyStatesShedStateBuilder
	// Build builds the BACnetPropertyStatesShedState or returns an error if something is wrong
	Build() (BACnetPropertyStatesShedState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesShedState
}

// NewBACnetPropertyStatesShedStateBuilder() creates a BACnetPropertyStatesShedStateBuilder
func NewBACnetPropertyStatesShedStateBuilder() BACnetPropertyStatesShedStateBuilder {
	return &_BACnetPropertyStatesShedStateBuilder{_BACnetPropertyStatesShedState: new(_BACnetPropertyStatesShedState)}
}

type _BACnetPropertyStatesShedStateBuilder struct {
	*_BACnetPropertyStatesShedState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesShedStateBuilder) = (*_BACnetPropertyStatesShedStateBuilder)(nil)

func (b *_BACnetPropertyStatesShedStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesShedStateBuilder) WithMandatoryFields(shedState BACnetShedStateTagged) BACnetPropertyStatesShedStateBuilder {
	return b.WithShedState(shedState)
}

func (b *_BACnetPropertyStatesShedStateBuilder) WithShedState(shedState BACnetShedStateTagged) BACnetPropertyStatesShedStateBuilder {
	b.ShedState = shedState
	return b
}

func (b *_BACnetPropertyStatesShedStateBuilder) WithShedStateBuilder(builderSupplier func(BACnetShedStateTaggedBuilder) BACnetShedStateTaggedBuilder) BACnetPropertyStatesShedStateBuilder {
	builder := builderSupplier(b.ShedState.CreateBACnetShedStateTaggedBuilder())
	var err error
	b.ShedState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetShedStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesShedStateBuilder) Build() (BACnetPropertyStatesShedState, error) {
	if b.ShedState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'shedState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesShedState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesShedStateBuilder) MustBuild() BACnetPropertyStatesShedState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesShedStateBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesShedStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesShedStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesShedStateBuilder().(*_BACnetPropertyStatesShedStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesShedStateBuilder creates a BACnetPropertyStatesShedStateBuilder
func (b *_BACnetPropertyStatesShedState) CreateBACnetPropertyStatesShedStateBuilder() BACnetPropertyStatesShedStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesShedStateBuilder()
	}
	return &_BACnetPropertyStatesShedStateBuilder{_BACnetPropertyStatesShedState: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesShedState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesShedState) GetShedState() BACnetShedStateTagged {
	return m.ShedState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesShedState(structType any) BACnetPropertyStatesShedState {
	if casted, ok := structType.(BACnetPropertyStatesShedState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesShedState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesShedState) GetTypeName() string {
	return "BACnetPropertyStatesShedState"
}

func (m *_BACnetPropertyStatesShedState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (shedState)
	lengthInBits += m.ShedState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesShedState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesShedState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesShedState BACnetPropertyStatesShedState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesShedState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesShedState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	shedState, err := ReadSimpleField[BACnetShedStateTagged](ctx, "shedState", ReadComplex[BACnetShedStateTagged](BACnetShedStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'shedState' field"))
	}
	m.ShedState = shedState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesShedState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesShedState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesShedState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesShedState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesShedState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesShedState")
		}

		if err := WriteSimpleField[BACnetShedStateTagged](ctx, "shedState", m.GetShedState(), WriteComplex[BACnetShedStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'shedState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesShedState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesShedState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesShedState) IsBACnetPropertyStatesShedState() {}

func (m *_BACnetPropertyStatesShedState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesShedState) deepCopy() *_BACnetPropertyStatesShedState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesShedStateCopy := &_BACnetPropertyStatesShedState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.ShedState.DeepCopy().(BACnetShedStateTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesShedStateCopy
}

func (m *_BACnetPropertyStatesShedState) String() string {
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
