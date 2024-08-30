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

public class BACnetNotificationParametersDoubleOutOfRange extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetContextTagDouble exceedingValue;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetContextTagDouble deadband;
  protected final BACnetContextTagDouble exceededLimit;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersDoubleOutOfRange(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetContextTagDouble exceedingValue,
      BACnetStatusFlagsTagged statusFlags,
      BACnetContextTagDouble deadband,
      BACnetContextTagDouble exceededLimit,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.exceedingValue = exceedingValue;
    this.statusFlags = statusFlags;
    this.deadband = deadband;
    this.exceededLimit = exceededLimit;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetContextTagDouble getExceedingValue() {
    return exceedingValue;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetContextTagDouble getDeadband() {
    return deadband;
  }

  public BACnetContextTagDouble getExceededLimit() {
    return exceededLimit;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersDoubleOutOfRange");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (exceedingValue)
    writeSimpleField("exceedingValue", exceedingValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (deadband)
    writeSimpleField("deadband", deadband, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (exceededLimit)
    writeSimpleField("exceededLimit", exceededLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersDoubleOutOfRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersDoubleOutOfRange _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (exceedingValue)
    lengthInBits += exceedingValue.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (deadband)
    lengthInBits += deadband.getLengthInBits();

    // Simple field (exceededLimit)
    lengthInBits += exceededLimit.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersBuilder staticParseBACnetNotificationParametersBuilder(
      ReadBuffer readBuffer,
      Short peekedTagNumber,
      Short tagNumber,
      BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersDoubleOutOfRange");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetContextTagDouble exceedingValue =
        readSimpleField(
            "exceedingValue",
            readComplex(
                () ->
                    (BACnetContextTagDouble)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (0), (BACnetDataType) (BACnetDataType.DOUBLE)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagDouble deadband =
        readSimpleField(
            "deadband",
            readComplex(
                () ->
                    (BACnetContextTagDouble)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.DOUBLE)),
                readBuffer));

    BACnetContextTagDouble exceededLimit =
        readSimpleField(
            "exceededLimit",
            readComplex(
                () ->
                    (BACnetContextTagDouble)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (3), (BACnetDataType) (BACnetDataType.DOUBLE)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersDoubleOutOfRange");
    // Create the instance
    return new BACnetNotificationParametersDoubleOutOfRangeBuilderImpl(
        innerOpeningTag,
        exceedingValue,
        statusFlags,
        deadband,
        exceededLimit,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersDoubleOutOfRangeBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetContextTagDouble exceedingValue;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetContextTagDouble deadband;
    private final BACnetContextTagDouble exceededLimit;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersDoubleOutOfRangeBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetContextTagDouble exceedingValue,
        BACnetStatusFlagsTagged statusFlags,
        BACnetContextTagDouble deadband,
        BACnetContextTagDouble exceededLimit,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.exceedingValue = exceedingValue;
      this.statusFlags = statusFlags;
      this.deadband = deadband;
      this.exceededLimit = exceededLimit;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersDoubleOutOfRange build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersDoubleOutOfRange bACnetNotificationParametersDoubleOutOfRange =
          new BACnetNotificationParametersDoubleOutOfRange(
              openingTag,
              peekedTagHeader,
              closingTag,
              innerOpeningTag,
              exceedingValue,
              statusFlags,
              deadband,
              exceededLimit,
              innerClosingTag,
              tagNumber,
              objectTypeArgument);
      return bACnetNotificationParametersDoubleOutOfRange;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersDoubleOutOfRange)) {
      return false;
    }
    BACnetNotificationParametersDoubleOutOfRange that =
        (BACnetNotificationParametersDoubleOutOfRange) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getExceedingValue() == that.getExceedingValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getDeadband() == that.getDeadband())
        && (getExceededLimit() == that.getExceededLimit())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getExceedingValue(),
        getStatusFlags(),
        getDeadband(),
        getExceededLimit(),
        getInnerClosingTag());
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
