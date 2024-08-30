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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ConfirmedPrivateTransferError extends BACnetError implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getErrorChoice() {
    return BACnetConfirmedServiceChoice.CONFIRMED_PRIVATE_TRANSFER;
  }

  // Properties.
  protected final ErrorEnclosed errorType;
  protected final BACnetVendorIdTagged vendorId;
  protected final BACnetContextTagUnsignedInteger serviceNumber;
  protected final BACnetConstructedData errorParameters;

  public ConfirmedPrivateTransferError(
      ErrorEnclosed errorType,
      BACnetVendorIdTagged vendorId,
      BACnetContextTagUnsignedInteger serviceNumber,
      BACnetConstructedData errorParameters) {
    super();
    this.errorType = errorType;
    this.vendorId = vendorId;
    this.serviceNumber = serviceNumber;
    this.errorParameters = errorParameters;
  }

  public ErrorEnclosed getErrorType() {
    return errorType;
  }

  public BACnetVendorIdTagged getVendorId() {
    return vendorId;
  }

  public BACnetContextTagUnsignedInteger getServiceNumber() {
    return serviceNumber;
  }

  public BACnetConstructedData getErrorParameters() {
    return errorParameters;
  }

  @Override
  protected void serializeBACnetErrorChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ConfirmedPrivateTransferError");

    // Simple Field (errorType)
    writeSimpleField("errorType", errorType, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (serviceNumber)
    writeSimpleField("serviceNumber", serviceNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (errorParameters) (Can be skipped, if the value is null)
    writeOptionalField(
        "errorParameters", errorParameters, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ConfirmedPrivateTransferError");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConfirmedPrivateTransferError _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (errorType)
    lengthInBits += errorType.getLengthInBits();

    // Simple field (vendorId)
    lengthInBits += vendorId.getLengthInBits();

    // Simple field (serviceNumber)
    lengthInBits += serviceNumber.getLengthInBits();

    // Optional Field (errorParameters)
    if (errorParameters != null) {
      lengthInBits += errorParameters.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetErrorBuilder staticParseBACnetErrorBuilder(
      ReadBuffer readBuffer, BACnetConfirmedServiceChoice errorChoice) throws ParseException {
    readBuffer.pullContext("ConfirmedPrivateTransferError");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ErrorEnclosed errorType =
        readSimpleField(
            "errorType",
            readComplex(() -> ErrorEnclosed.staticParse(readBuffer, (short) (0)), readBuffer));

    BACnetVendorIdTagged vendorId =
        readSimpleField(
            "vendorId",
            readComplex(
                () ->
                    BACnetVendorIdTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger serviceNumber =
        readSimpleField(
            "serviceNumber",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData errorParameters =
        readOptionalField(
            "errorParameters",
            readComplex(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (3),
                        (BACnetObjectType) (BACnetObjectType.VENDOR_PROPRIETARY_VALUE),
                        (BACnetPropertyIdentifier)
                            (BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE),
                        (BACnetTagPayloadUnsignedInteger) (null)),
                readBuffer));

    readBuffer.closeContext("ConfirmedPrivateTransferError");
    // Create the instance
    return new ConfirmedPrivateTransferErrorBuilderImpl(
        errorType, vendorId, serviceNumber, errorParameters);
  }

  public static class ConfirmedPrivateTransferErrorBuilderImpl
      implements BACnetError.BACnetErrorBuilder {
    private final ErrorEnclosed errorType;
    private final BACnetVendorIdTagged vendorId;
    private final BACnetContextTagUnsignedInteger serviceNumber;
    private final BACnetConstructedData errorParameters;

    public ConfirmedPrivateTransferErrorBuilderImpl(
        ErrorEnclosed errorType,
        BACnetVendorIdTagged vendorId,
        BACnetContextTagUnsignedInteger serviceNumber,
        BACnetConstructedData errorParameters) {
      this.errorType = errorType;
      this.vendorId = vendorId;
      this.serviceNumber = serviceNumber;
      this.errorParameters = errorParameters;
    }

    public ConfirmedPrivateTransferError build() {
      ConfirmedPrivateTransferError confirmedPrivateTransferError =
          new ConfirmedPrivateTransferError(errorType, vendorId, serviceNumber, errorParameters);
      return confirmedPrivateTransferError;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConfirmedPrivateTransferError)) {
      return false;
    }
    ConfirmedPrivateTransferError that = (ConfirmedPrivateTransferError) o;
    return (getErrorType() == that.getErrorType())
        && (getVendorId() == that.getVendorId())
        && (getServiceNumber() == that.getServiceNumber())
        && (getErrorParameters() == that.getErrorParameters())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getErrorType(), getVendorId(), getServiceNumber(), getErrorParameters());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
