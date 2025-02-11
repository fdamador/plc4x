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

// MFuncPropCon is the corresponding interface of MFuncPropCon
type MFuncPropCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMFuncPropCon is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMFuncPropCon()
	// CreateBuilder creates a MFuncPropConBuilder
	CreateMFuncPropConBuilder() MFuncPropConBuilder
}

// _MFuncPropCon is the data-structure of this message
type _MFuncPropCon struct {
	CEMIContract
}

var _ MFuncPropCon = (*_MFuncPropCon)(nil)
var _ CEMIRequirements = (*_MFuncPropCon)(nil)

// NewMFuncPropCon factory function for _MFuncPropCon
func NewMFuncPropCon(size uint16) *_MFuncPropCon {
	_result := &_MFuncPropCon{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MFuncPropConBuilder is a builder for MFuncPropCon
type MFuncPropConBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MFuncPropConBuilder
	// Build builds the MFuncPropCon or returns an error if something is wrong
	Build() (MFuncPropCon, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MFuncPropCon
}

// NewMFuncPropConBuilder() creates a MFuncPropConBuilder
func NewMFuncPropConBuilder() MFuncPropConBuilder {
	return &_MFuncPropConBuilder{_MFuncPropCon: new(_MFuncPropCon)}
}

type _MFuncPropConBuilder struct {
	*_MFuncPropCon

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MFuncPropConBuilder) = (*_MFuncPropConBuilder)(nil)

func (b *_MFuncPropConBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_MFuncPropConBuilder) WithMandatoryFields() MFuncPropConBuilder {
	return b
}

func (b *_MFuncPropConBuilder) Build() (MFuncPropCon, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MFuncPropCon.deepCopy(), nil
}

func (b *_MFuncPropConBuilder) MustBuild() MFuncPropCon {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MFuncPropConBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_MFuncPropConBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MFuncPropConBuilder) DeepCopy() any {
	_copy := b.CreateMFuncPropConBuilder().(*_MFuncPropConBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMFuncPropConBuilder creates a MFuncPropConBuilder
func (b *_MFuncPropCon) CreateMFuncPropConBuilder() MFuncPropConBuilder {
	if b == nil {
		return NewMFuncPropConBuilder()
	}
	return &_MFuncPropConBuilder{_MFuncPropCon: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropCon) GetMessageCode() uint8 {
	return 0xFA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropCon) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMFuncPropCon(structType any) MFuncPropCon {
	if casted, ok := structType.(MFuncPropCon); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropCon); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropCon) GetTypeName() string {
	return "MFuncPropCon"
}

func (m *_MFuncPropCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MFuncPropCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mFuncPropCon MFuncPropCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropCon")
	}

	return m, nil
}

func (m *_MFuncPropCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropCon")
		}

		if popErr := writeBuffer.PopContext("MFuncPropCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropCon) IsMFuncPropCon() {}

func (m *_MFuncPropCon) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MFuncPropCon) deepCopy() *_MFuncPropCon {
	if m == nil {
		return nil
	}
	_MFuncPropConCopy := &_MFuncPropCon{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MFuncPropConCopy
}

func (m *_MFuncPropCon) String() string {
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
