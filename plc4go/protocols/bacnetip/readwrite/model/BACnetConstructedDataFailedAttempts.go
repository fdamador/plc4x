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

// BACnetConstructedDataFailedAttempts is the corresponding interface of BACnetConstructedDataFailedAttempts
type BACnetConstructedDataFailedAttempts interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFailedAttempts returns FailedAttempts (property field)
	GetFailedAttempts() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataFailedAttempts is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFailedAttempts()
	// CreateBuilder creates a BACnetConstructedDataFailedAttemptsBuilder
	CreateBACnetConstructedDataFailedAttemptsBuilder() BACnetConstructedDataFailedAttemptsBuilder
}

// _BACnetConstructedDataFailedAttempts is the data-structure of this message
type _BACnetConstructedDataFailedAttempts struct {
	BACnetConstructedDataContract
	FailedAttempts BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataFailedAttempts = (*_BACnetConstructedDataFailedAttempts)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFailedAttempts)(nil)

// NewBACnetConstructedDataFailedAttempts factory function for _BACnetConstructedDataFailedAttempts
func NewBACnetConstructedDataFailedAttempts(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, failedAttempts BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFailedAttempts {
	if failedAttempts == nil {
		panic("failedAttempts of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataFailedAttempts must not be nil")
	}
	_result := &_BACnetConstructedDataFailedAttempts{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FailedAttempts:                failedAttempts,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFailedAttemptsBuilder is a builder for BACnetConstructedDataFailedAttempts
type BACnetConstructedDataFailedAttemptsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(failedAttempts BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFailedAttemptsBuilder
	// WithFailedAttempts adds FailedAttempts (property field)
	WithFailedAttempts(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFailedAttemptsBuilder
	// WithFailedAttemptsBuilder adds FailedAttempts (property field) which is build by the builder
	WithFailedAttemptsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFailedAttemptsBuilder
	// Build builds the BACnetConstructedDataFailedAttempts or returns an error if something is wrong
	Build() (BACnetConstructedDataFailedAttempts, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFailedAttempts
}

// NewBACnetConstructedDataFailedAttemptsBuilder() creates a BACnetConstructedDataFailedAttemptsBuilder
func NewBACnetConstructedDataFailedAttemptsBuilder() BACnetConstructedDataFailedAttemptsBuilder {
	return &_BACnetConstructedDataFailedAttemptsBuilder{_BACnetConstructedDataFailedAttempts: new(_BACnetConstructedDataFailedAttempts)}
}

type _BACnetConstructedDataFailedAttemptsBuilder struct {
	*_BACnetConstructedDataFailedAttempts

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFailedAttemptsBuilder) = (*_BACnetConstructedDataFailedAttemptsBuilder)(nil)

func (b *_BACnetConstructedDataFailedAttemptsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) WithMandatoryFields(failedAttempts BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFailedAttemptsBuilder {
	return b.WithFailedAttempts(failedAttempts)
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) WithFailedAttempts(failedAttempts BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFailedAttemptsBuilder {
	b.FailedAttempts = failedAttempts
	return b
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) WithFailedAttemptsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFailedAttemptsBuilder {
	builder := builderSupplier(b.FailedAttempts.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.FailedAttempts, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) Build() (BACnetConstructedDataFailedAttempts, error) {
	if b.FailedAttempts == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'failedAttempts' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFailedAttempts.deepCopy(), nil
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) MustBuild() BACnetConstructedDataFailedAttempts {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataFailedAttemptsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFailedAttemptsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFailedAttemptsBuilder().(*_BACnetConstructedDataFailedAttemptsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFailedAttemptsBuilder creates a BACnetConstructedDataFailedAttemptsBuilder
func (b *_BACnetConstructedDataFailedAttempts) CreateBACnetConstructedDataFailedAttemptsBuilder() BACnetConstructedDataFailedAttemptsBuilder {
	if b == nil {
		return NewBACnetConstructedDataFailedAttemptsBuilder()
	}
	return &_BACnetConstructedDataFailedAttemptsBuilder{_BACnetConstructedDataFailedAttempts: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFailedAttempts) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAILED_ATTEMPTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetFailedAttempts() BACnetApplicationTagUnsignedInteger {
	return m.FailedAttempts
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFailedAttempts())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFailedAttempts(structType any) BACnetConstructedDataFailedAttempts {
	if casted, ok := structType.(BACnetConstructedDataFailedAttempts); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFailedAttempts); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFailedAttempts) GetTypeName() string {
	return "BACnetConstructedDataFailedAttempts"
}

func (m *_BACnetConstructedDataFailedAttempts) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (failedAttempts)
	lengthInBits += m.FailedAttempts.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFailedAttempts) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFailedAttempts) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFailedAttempts BACnetConstructedDataFailedAttempts, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFailedAttempts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFailedAttempts")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	failedAttempts, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "failedAttempts", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'failedAttempts' field"))
	}
	m.FailedAttempts = failedAttempts

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), failedAttempts)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFailedAttempts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFailedAttempts")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFailedAttempts) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFailedAttempts) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFailedAttempts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFailedAttempts")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "failedAttempts", m.GetFailedAttempts(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'failedAttempts' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFailedAttempts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFailedAttempts")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFailedAttempts) IsBACnetConstructedDataFailedAttempts() {}

func (m *_BACnetConstructedDataFailedAttempts) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFailedAttempts) deepCopy() *_BACnetConstructedDataFailedAttempts {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFailedAttemptsCopy := &_BACnetConstructedDataFailedAttempts{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.FailedAttempts.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFailedAttemptsCopy
}

func (m *_BACnetConstructedDataFailedAttempts) String() string {
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
