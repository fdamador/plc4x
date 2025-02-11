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

// BACnetConfirmedServiceRequestDeleteObject is the corresponding interface of BACnetConfirmedServiceRequestDeleteObject
type BACnetConfirmedServiceRequestDeleteObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// IsBACnetConfirmedServiceRequestDeleteObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestDeleteObject()
	// CreateBuilder creates a BACnetConfirmedServiceRequestDeleteObjectBuilder
	CreateBACnetConfirmedServiceRequestDeleteObjectBuilder() BACnetConfirmedServiceRequestDeleteObjectBuilder
}

// _BACnetConfirmedServiceRequestDeleteObject is the data-structure of this message
type _BACnetConfirmedServiceRequestDeleteObject struct {
	BACnetConfirmedServiceRequestContract
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
}

var _ BACnetConfirmedServiceRequestDeleteObject = (*_BACnetConfirmedServiceRequestDeleteObject)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestDeleteObject)(nil)

// NewBACnetConfirmedServiceRequestDeleteObject factory function for _BACnetConfirmedServiceRequestDeleteObject
func NewBACnetConfirmedServiceRequestDeleteObject(objectIdentifier BACnetApplicationTagObjectIdentifier, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestDeleteObject {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetConfirmedServiceRequestDeleteObject must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestDeleteObject{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectIdentifier:                      objectIdentifier,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestDeleteObjectBuilder is a builder for BACnetConfirmedServiceRequestDeleteObject
type BACnetConfirmedServiceRequestDeleteObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestDeleteObjectBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestDeleteObjectBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestDeleteObjectBuilder
	// Build builds the BACnetConfirmedServiceRequestDeleteObject or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestDeleteObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestDeleteObject
}

// NewBACnetConfirmedServiceRequestDeleteObjectBuilder() creates a BACnetConfirmedServiceRequestDeleteObjectBuilder
func NewBACnetConfirmedServiceRequestDeleteObjectBuilder() BACnetConfirmedServiceRequestDeleteObjectBuilder {
	return &_BACnetConfirmedServiceRequestDeleteObjectBuilder{_BACnetConfirmedServiceRequestDeleteObject: new(_BACnetConfirmedServiceRequestDeleteObject)}
}

type _BACnetConfirmedServiceRequestDeleteObjectBuilder struct {
	*_BACnetConfirmedServiceRequestDeleteObject

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestDeleteObjectBuilder) = (*_BACnetConfirmedServiceRequestDeleteObjectBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) WithMandatoryFields(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestDeleteObjectBuilder {
	return b.WithObjectIdentifier(objectIdentifier)
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) WithObjectIdentifier(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestDeleteObjectBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestDeleteObjectBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) Build() (BACnetConfirmedServiceRequestDeleteObject, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestDeleteObject.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) MustBuild() BACnetConfirmedServiceRequestDeleteObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestDeleteObjectBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestDeleteObjectBuilder().(*_BACnetConfirmedServiceRequestDeleteObjectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestDeleteObjectBuilder creates a BACnetConfirmedServiceRequestDeleteObjectBuilder
func (b *_BACnetConfirmedServiceRequestDeleteObject) CreateBACnetConfirmedServiceRequestDeleteObjectBuilder() BACnetConfirmedServiceRequestDeleteObjectBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestDeleteObjectBuilder()
	}
	return &_BACnetConfirmedServiceRequestDeleteObjectBuilder{_BACnetConfirmedServiceRequestDeleteObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DELETE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeleteObject(structType any) BACnetConfirmedServiceRequestDeleteObject {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeleteObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeleteObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeleteObject"
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestDeleteObject BACnetConfirmedServiceRequestDeleteObject, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeleteObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeleteObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeleteObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeleteObject")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeleteObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeleteObject")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeleteObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeleteObject")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) IsBACnetConfirmedServiceRequestDeleteObject() {}

func (m *_BACnetConfirmedServiceRequestDeleteObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) deepCopy() *_BACnetConfirmedServiceRequestDeleteObject {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestDeleteObjectCopy := &_BACnetConfirmedServiceRequestDeleteObject{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.ObjectIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestDeleteObjectCopy
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) String() string {
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
