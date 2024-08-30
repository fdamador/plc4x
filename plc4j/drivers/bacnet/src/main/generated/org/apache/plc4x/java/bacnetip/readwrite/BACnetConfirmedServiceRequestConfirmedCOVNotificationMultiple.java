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

public class BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
    extends BACnetConfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.CONFIRMED_COV_NOTIFICATION_MULTIPLE;
  }

  // Properties.
  protected final BACnetContextTagUnsignedInteger subscriberProcessIdentifier;
  protected final BACnetContextTagObjectIdentifier initiatingDeviceIdentifier;
  protected final BACnetContextTagUnsignedInteger timeRemaining;
  protected final BACnetTimeStampEnclosed timestamp;
  protected final ListOfCovNotificationsList listOfCovNotifications;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(
      BACnetContextTagUnsignedInteger subscriberProcessIdentifier,
      BACnetContextTagObjectIdentifier initiatingDeviceIdentifier,
      BACnetContextTagUnsignedInteger timeRemaining,
      BACnetTimeStampEnclosed timestamp,
      ListOfCovNotificationsList listOfCovNotifications,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.subscriberProcessIdentifier = subscriberProcessIdentifier;
    this.initiatingDeviceIdentifier = initiatingDeviceIdentifier;
    this.timeRemaining = timeRemaining;
    this.timestamp = timestamp;
    this.listOfCovNotifications = listOfCovNotifications;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagUnsignedInteger getSubscriberProcessIdentifier() {
    return subscriberProcessIdentifier;
  }

  public BACnetContextTagObjectIdentifier getInitiatingDeviceIdentifier() {
    return initiatingDeviceIdentifier;
  }

  public BACnetContextTagUnsignedInteger getTimeRemaining() {
    return timeRemaining;
  }

  public BACnetTimeStampEnclosed getTimestamp() {
    return timestamp;
  }

  public ListOfCovNotificationsList getListOfCovNotifications() {
    return listOfCovNotifications;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple");

    // Simple Field (subscriberProcessIdentifier)
    writeSimpleField(
        "subscriberProcessIdentifier",
        subscriberProcessIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (initiatingDeviceIdentifier)
    writeSimpleField(
        "initiatingDeviceIdentifier",
        initiatingDeviceIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeRemaining)
    writeSimpleField("timeRemaining", timeRemaining, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (timestamp) (Can be skipped, if the value is null)
    writeOptionalField("timestamp", timestamp, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (listOfCovNotifications)
    writeSimpleField(
        "listOfCovNotifications",
        listOfCovNotifications,
        new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (subscriberProcessIdentifier)
    lengthInBits += subscriberProcessIdentifier.getLengthInBits();

    // Simple field (initiatingDeviceIdentifier)
    lengthInBits += initiatingDeviceIdentifier.getLengthInBits();

    // Simple field (timeRemaining)
    lengthInBits += timeRemaining.getLengthInBits();

    // Optional Field (timestamp)
    if (timestamp != null) {
      lengthInBits += timestamp.getLengthInBits();
    }

    // Simple field (listOfCovNotifications)
    lengthInBits += listOfCovNotifications.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger subscriberProcessIdentifier =
        readSimpleField(
            "subscriberProcessIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagObjectIdentifier initiatingDeviceIdentifier =
        readSimpleField(
            "initiatingDeviceIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetContextTagUnsignedInteger timeRemaining =
        readSimpleField(
            "timeRemaining",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetTimeStampEnclosed timestamp =
        readOptionalField(
            "timestamp",
            readComplex(
                () -> BACnetTimeStampEnclosed.staticParse(readBuffer, (short) (3)), readBuffer));

    ListOfCovNotificationsList listOfCovNotifications =
        readSimpleField(
            "listOfCovNotifications",
            readComplex(
                () -> ListOfCovNotificationsList.staticParse(readBuffer, (short) (4)), readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple");
    // Create the instance
    return new BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilderImpl(
        subscriberProcessIdentifier,
        initiatingDeviceIdentifier,
        timeRemaining,
        timestamp,
        listOfCovNotifications,
        serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagUnsignedInteger subscriberProcessIdentifier;
    private final BACnetContextTagObjectIdentifier initiatingDeviceIdentifier;
    private final BACnetContextTagUnsignedInteger timeRemaining;
    private final BACnetTimeStampEnclosed timestamp;
    private final ListOfCovNotificationsList listOfCovNotifications;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilderImpl(
        BACnetContextTagUnsignedInteger subscriberProcessIdentifier,
        BACnetContextTagObjectIdentifier initiatingDeviceIdentifier,
        BACnetContextTagUnsignedInteger timeRemaining,
        BACnetTimeStampEnclosed timestamp,
        ListOfCovNotificationsList listOfCovNotifications,
        Long serviceRequestLength) {
      this.subscriberProcessIdentifier = subscriberProcessIdentifier;
      this.initiatingDeviceIdentifier = initiatingDeviceIdentifier;
      this.timeRemaining = timeRemaining;
      this.timestamp = timestamp;
      this.listOfCovNotifications = listOfCovNotifications;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple build(
        Long serviceRequestLength) {
      BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
          bACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple =
              new BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(
                  subscriberProcessIdentifier,
                  initiatingDeviceIdentifier,
                  timeRemaining,
                  timestamp,
                  listOfCovNotifications,
                  serviceRequestLength);
      return bACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)) {
      return false;
    }
    BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple that =
        (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) o;
    return (getSubscriberProcessIdentifier() == that.getSubscriberProcessIdentifier())
        && (getInitiatingDeviceIdentifier() == that.getInitiatingDeviceIdentifier())
        && (getTimeRemaining() == that.getTimeRemaining())
        && (getTimestamp() == that.getTimestamp())
        && (getListOfCovNotifications() == that.getListOfCovNotifications())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSubscriberProcessIdentifier(),
        getInitiatingDeviceIdentifier(),
        getTimeRemaining(),
        getTimestamp(),
        getListOfCovNotifications());
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
