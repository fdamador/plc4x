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

// BACnetConstructedDataPulseRate is the corresponding interface of BACnetConstructedDataPulseRate
type BACnetConstructedDataPulseRate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPulseRate returns PulseRate (property field)
	GetPulseRate() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPulseRate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPulseRate()
	// CreateBuilder creates a BACnetConstructedDataPulseRateBuilder
	CreateBACnetConstructedDataPulseRateBuilder() BACnetConstructedDataPulseRateBuilder
}

// _BACnetConstructedDataPulseRate is the data-structure of this message
type _BACnetConstructedDataPulseRate struct {
	BACnetConstructedDataContract
	PulseRate BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPulseRate = (*_BACnetConstructedDataPulseRate)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPulseRate)(nil)

// NewBACnetConstructedDataPulseRate factory function for _BACnetConstructedDataPulseRate
func NewBACnetConstructedDataPulseRate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, pulseRate BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPulseRate {
	if pulseRate == nil {
		panic("pulseRate of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPulseRate must not be nil")
	}
	_result := &_BACnetConstructedDataPulseRate{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PulseRate:                     pulseRate,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPulseRateBuilder is a builder for BACnetConstructedDataPulseRate
type BACnetConstructedDataPulseRateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pulseRate BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPulseRateBuilder
	// WithPulseRate adds PulseRate (property field)
	WithPulseRate(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPulseRateBuilder
	// WithPulseRateBuilder adds PulseRate (property field) which is build by the builder
	WithPulseRateBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPulseRateBuilder
	// Build builds the BACnetConstructedDataPulseRate or returns an error if something is wrong
	Build() (BACnetConstructedDataPulseRate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPulseRate
}

// NewBACnetConstructedDataPulseRateBuilder() creates a BACnetConstructedDataPulseRateBuilder
func NewBACnetConstructedDataPulseRateBuilder() BACnetConstructedDataPulseRateBuilder {
	return &_BACnetConstructedDataPulseRateBuilder{_BACnetConstructedDataPulseRate: new(_BACnetConstructedDataPulseRate)}
}

type _BACnetConstructedDataPulseRateBuilder struct {
	*_BACnetConstructedDataPulseRate

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPulseRateBuilder) = (*_BACnetConstructedDataPulseRateBuilder)(nil)

func (b *_BACnetConstructedDataPulseRateBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataPulseRateBuilder) WithMandatoryFields(pulseRate BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPulseRateBuilder {
	return b.WithPulseRate(pulseRate)
}

func (b *_BACnetConstructedDataPulseRateBuilder) WithPulseRate(pulseRate BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPulseRateBuilder {
	b.PulseRate = pulseRate
	return b
}

func (b *_BACnetConstructedDataPulseRateBuilder) WithPulseRateBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPulseRateBuilder {
	builder := builderSupplier(b.PulseRate.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.PulseRate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPulseRateBuilder) Build() (BACnetConstructedDataPulseRate, error) {
	if b.PulseRate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'pulseRate' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPulseRate.deepCopy(), nil
}

func (b *_BACnetConstructedDataPulseRateBuilder) MustBuild() BACnetConstructedDataPulseRate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataPulseRateBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPulseRateBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPulseRateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPulseRateBuilder().(*_BACnetConstructedDataPulseRateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPulseRateBuilder creates a BACnetConstructedDataPulseRateBuilder
func (b *_BACnetConstructedDataPulseRate) CreateBACnetConstructedDataPulseRateBuilder() BACnetConstructedDataPulseRateBuilder {
	if b == nil {
		return NewBACnetConstructedDataPulseRateBuilder()
	}
	return &_BACnetConstructedDataPulseRateBuilder{_BACnetConstructedDataPulseRate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPulseRate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PULSE_RATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPulseRate) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetPulseRate() BACnetApplicationTagUnsignedInteger {
	return m.PulseRate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetPulseRate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPulseRate(structType any) BACnetConstructedDataPulseRate {
	if casted, ok := structType.(BACnetConstructedDataPulseRate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPulseRate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPulseRate) GetTypeName() string {
	return "BACnetConstructedDataPulseRate"
}

func (m *_BACnetConstructedDataPulseRate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (pulseRate)
	lengthInBits += m.PulseRate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPulseRate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPulseRate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPulseRate BACnetConstructedDataPulseRate, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPulseRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPulseRate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pulseRate, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "pulseRate", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pulseRate' field"))
	}
	m.PulseRate = pulseRate

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), pulseRate)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPulseRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPulseRate")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPulseRate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPulseRate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPulseRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPulseRate")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "pulseRate", m.GetPulseRate(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'pulseRate' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPulseRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPulseRate")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPulseRate) IsBACnetConstructedDataPulseRate() {}

func (m *_BACnetConstructedDataPulseRate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPulseRate) deepCopy() *_BACnetConstructedDataPulseRate {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPulseRateCopy := &_BACnetConstructedDataPulseRate{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PulseRate.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPulseRateCopy
}

func (m *_BACnetConstructedDataPulseRate) String() string {
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
