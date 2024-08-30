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

public abstract class BACnetNotificationParameters implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetTagHeader peekedTagHeader;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParameters(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super();
    this.openingTag = openingTag;
    this.peekedTagHeader = peekedTagHeader;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  protected abstract void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParameters");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Switch field (Serialize the sub-type)
    serializeBACnetNotificationParametersChild(writeBuffer);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParameters");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetNotificationParameters _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParameters staticParse(
      ReadBuffer readBuffer, Short tagNumber, BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParameters");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetNotificationParametersBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0)) {
      builder =
          BACnetNotificationParametersChangeOfBitString
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 1)) {
      builder =
          BACnetNotificationParametersChangeOfState.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 2)) {
      builder =
          BACnetNotificationParametersChangeOfValue.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 3)) {
      builder =
          BACnetNotificationParametersCommandFailure.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 4)) {
      builder =
          BACnetNotificationParametersFloatingLimit.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 5)) {
      builder =
          BACnetNotificationParametersOutOfRange.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 6)) {
      builder =
          BACnetNotificationParametersComplexEventType
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 8)) {
      builder =
          BACnetNotificationParametersChangeOfLifeSafety
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 9)) {
      builder =
          BACnetNotificationParametersExtended.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 10)) {
      builder =
          BACnetNotificationParametersBufferReady.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 11)) {
      builder =
          BACnetNotificationParametersUnsignedRange.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 13)) {
      builder =
          BACnetNotificationParametersAccessEvent.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 14)) {
      builder =
          BACnetNotificationParametersDoubleOutOfRange
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 15)) {
      builder =
          BACnetNotificationParametersSignedOutOfRange
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 16)) {
      builder =
          BACnetNotificationParametersUnsignedOutOfRange
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 17)) {
      builder =
          BACnetNotificationParametersChangeOfCharacterString
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 18)) {
      builder =
          BACnetNotificationParametersChangeOfStatusFlags
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 19)) {
      builder =
          BACnetNotificationParametersChangeOfReliability
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 21)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValue
              .staticParseBACnetNotificationParametersBuilder(
                  readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 22)) {
      builder =
          BACnetNotificationParametersChangeOfTimer.staticParseBACnetNotificationParametersBuilder(
              readBuffer, peekedTagNumber, tagNumber, objectTypeArgument);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + "]");
    }

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetNotificationParameters");
    // Create the instance
    BACnetNotificationParameters _bACnetNotificationParameters =
        builder.build(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    return _bACnetNotificationParameters;
  }

  public interface BACnetNotificationParametersBuilder {
    BACnetNotificationParameters build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParameters)) {
      return false;
    }
    BACnetNotificationParameters that = (BACnetNotificationParameters) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getPeekedTagHeader() == that.getPeekedTagHeader())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getPeekedTagHeader(), getClosingTag());
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
