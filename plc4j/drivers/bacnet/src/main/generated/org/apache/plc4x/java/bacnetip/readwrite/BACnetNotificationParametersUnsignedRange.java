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

public class BACnetNotificationParametersUnsignedRange extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetContextTagUnsignedInteger sequenceNumber;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetContextTagUnsignedInteger exceededLimit;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersUnsignedRange(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetContextTagUnsignedInteger sequenceNumber,
      BACnetStatusFlagsTagged statusFlags,
      BACnetContextTagUnsignedInteger exceededLimit,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.sequenceNumber = sequenceNumber;
    this.statusFlags = statusFlags;
    this.exceededLimit = exceededLimit;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetContextTagUnsignedInteger getSequenceNumber() {
    return sequenceNumber;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetContextTagUnsignedInteger getExceededLimit() {
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
    writeBuffer.pushContext("BACnetNotificationParametersUnsignedRange");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (sequenceNumber)
    writeSimpleField("sequenceNumber", sequenceNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (exceededLimit)
    writeSimpleField("exceededLimit", exceededLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersUnsignedRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersUnsignedRange _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (sequenceNumber)
    lengthInBits += sequenceNumber.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

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
    readBuffer.pullContext("BACnetNotificationParametersUnsignedRange");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetContextTagUnsignedInteger sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger exceededLimit =
        readSimpleField(
            "exceededLimit",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersUnsignedRange");
    // Create the instance
    return new BACnetNotificationParametersUnsignedRangeBuilderImpl(
        innerOpeningTag,
        sequenceNumber,
        statusFlags,
        exceededLimit,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersUnsignedRangeBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetContextTagUnsignedInteger sequenceNumber;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetContextTagUnsignedInteger exceededLimit;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersUnsignedRangeBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetContextTagUnsignedInteger sequenceNumber,
        BACnetStatusFlagsTagged statusFlags,
        BACnetContextTagUnsignedInteger exceededLimit,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.sequenceNumber = sequenceNumber;
      this.statusFlags = statusFlags;
      this.exceededLimit = exceededLimit;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersUnsignedRange build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersUnsignedRange bACnetNotificationParametersUnsignedRange =
          new BACnetNotificationParametersUnsignedRange(
              openingTag,
              peekedTagHeader,
              closingTag,
              innerOpeningTag,
              sequenceNumber,
              statusFlags,
              exceededLimit,
              innerClosingTag,
              tagNumber,
              objectTypeArgument);
      return bACnetNotificationParametersUnsignedRange;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersUnsignedRange)) {
      return false;
    }
    BACnetNotificationParametersUnsignedRange that = (BACnetNotificationParametersUnsignedRange) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getStatusFlags() == that.getStatusFlags())
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
        getSequenceNumber(),
        getStatusFlags(),
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
