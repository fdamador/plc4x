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

public class BACnetUnconfirmedServiceRequestIAm extends BACnetUnconfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetUnconfirmedServiceChoice getServiceChoice() {
    return BACnetUnconfirmedServiceChoice.I_AM;
  }

  // Properties.
  protected final BACnetApplicationTagObjectIdentifier deviceIdentifier;
  protected final BACnetApplicationTagUnsignedInteger maximumApduLengthAcceptedLength;
  protected final BACnetSegmentationTagged segmentationSupported;
  protected final BACnetVendorIdTagged vendorId;

  // Arguments.
  protected final Integer serviceRequestLength;

  public BACnetUnconfirmedServiceRequestIAm(
      BACnetApplicationTagObjectIdentifier deviceIdentifier,
      BACnetApplicationTagUnsignedInteger maximumApduLengthAcceptedLength,
      BACnetSegmentationTagged segmentationSupported,
      BACnetVendorIdTagged vendorId,
      Integer serviceRequestLength) {
    super(serviceRequestLength);
    this.deviceIdentifier = deviceIdentifier;
    this.maximumApduLengthAcceptedLength = maximumApduLengthAcceptedLength;
    this.segmentationSupported = segmentationSupported;
    this.vendorId = vendorId;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetApplicationTagObjectIdentifier getDeviceIdentifier() {
    return deviceIdentifier;
  }

  public BACnetApplicationTagUnsignedInteger getMaximumApduLengthAcceptedLength() {
    return maximumApduLengthAcceptedLength;
  }

  public BACnetSegmentationTagged getSegmentationSupported() {
    return segmentationSupported;
  }

  public BACnetVendorIdTagged getVendorId() {
    return vendorId;
  }

  @Override
  protected void serializeBACnetUnconfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetUnconfirmedServiceRequestIAm");

    // Simple Field (deviceIdentifier)
    writeSimpleField(
        "deviceIdentifier", deviceIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (maximumApduLengthAcceptedLength)
    writeSimpleField(
        "maximumApduLengthAcceptedLength",
        maximumApduLengthAcceptedLength,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (segmentationSupported)
    writeSimpleField(
        "segmentationSupported",
        segmentationSupported,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetUnconfirmedServiceRequestIAm");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetUnconfirmedServiceRequestIAm _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (deviceIdentifier)
    lengthInBits += deviceIdentifier.getLengthInBits();

    // Simple field (maximumApduLengthAcceptedLength)
    lengthInBits += maximumApduLengthAcceptedLength.getLengthInBits();

    // Simple field (segmentationSupported)
    lengthInBits += segmentationSupported.getLengthInBits();

    // Simple field (vendorId)
    lengthInBits += vendorId.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetUnconfirmedServiceRequestBuilder
      staticParseBACnetUnconfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Integer serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetUnconfirmedServiceRequestIAm");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagObjectIdentifier deviceIdentifier =
        readSimpleField(
            "deviceIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagObjectIdentifier)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagUnsignedInteger maximumApduLengthAcceptedLength =
        readSimpleField(
            "maximumApduLengthAcceptedLength",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetSegmentationTagged segmentationSupported =
        readSimpleField(
            "segmentationSupported",
            readComplex(
                () ->
                    BACnetSegmentationTagged.staticParse(
                        readBuffer, (short) (9), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));

    BACnetVendorIdTagged vendorId =
        readSimpleField(
            "vendorId",
            readComplex(
                () ->
                    BACnetVendorIdTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetUnconfirmedServiceRequestIAm");
    // Create the instance
    return new BACnetUnconfirmedServiceRequestIAmBuilderImpl(
        deviceIdentifier,
        maximumApduLengthAcceptedLength,
        segmentationSupported,
        vendorId,
        serviceRequestLength);
  }

  public static class BACnetUnconfirmedServiceRequestIAmBuilderImpl
      implements BACnetUnconfirmedServiceRequest.BACnetUnconfirmedServiceRequestBuilder {
    private final BACnetApplicationTagObjectIdentifier deviceIdentifier;
    private final BACnetApplicationTagUnsignedInteger maximumApduLengthAcceptedLength;
    private final BACnetSegmentationTagged segmentationSupported;
    private final BACnetVendorIdTagged vendorId;
    private final Integer serviceRequestLength;

    public BACnetUnconfirmedServiceRequestIAmBuilderImpl(
        BACnetApplicationTagObjectIdentifier deviceIdentifier,
        BACnetApplicationTagUnsignedInteger maximumApduLengthAcceptedLength,
        BACnetSegmentationTagged segmentationSupported,
        BACnetVendorIdTagged vendorId,
        Integer serviceRequestLength) {
      this.deviceIdentifier = deviceIdentifier;
      this.maximumApduLengthAcceptedLength = maximumApduLengthAcceptedLength;
      this.segmentationSupported = segmentationSupported;
      this.vendorId = vendorId;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetUnconfirmedServiceRequestIAm build(Integer serviceRequestLength) {

      BACnetUnconfirmedServiceRequestIAm bACnetUnconfirmedServiceRequestIAm =
          new BACnetUnconfirmedServiceRequestIAm(
              deviceIdentifier,
              maximumApduLengthAcceptedLength,
              segmentationSupported,
              vendorId,
              serviceRequestLength);
      return bACnetUnconfirmedServiceRequestIAm;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetUnconfirmedServiceRequestIAm)) {
      return false;
    }
    BACnetUnconfirmedServiceRequestIAm that = (BACnetUnconfirmedServiceRequestIAm) o;
    return (getDeviceIdentifier() == that.getDeviceIdentifier())
        && (getMaximumApduLengthAcceptedLength() == that.getMaximumApduLengthAcceptedLength())
        && (getSegmentationSupported() == that.getSegmentationSupported())
        && (getVendorId() == that.getVendorId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getDeviceIdentifier(),
        getMaximumApduLengthAcceptedLength(),
        getSegmentationSupported(),
        getVendorId());
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
