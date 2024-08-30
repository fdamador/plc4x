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

public class BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
    extends BACnetNotificationParametersChangeOfDiscreteValueNewValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagBoolean booleanValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagBoolean booleanValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.booleanValue = booleanValue;
    this.tagNumber = tagNumber;
  }

  public BACnetApplicationTagBoolean getBooleanValue() {
    return booleanValue;
  }

  @Override
  protected void serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean");

    // Simple Field (booleanValue)
    writeSimpleField("booleanValue", booleanValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (booleanValue)
    lengthInBits += booleanValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
      staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
          ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagBoolean booleanValue =
        readSimpleField(
            "booleanValue",
            readComplex(
                () -> (BACnetApplicationTagBoolean) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean");
    // Create the instance
    return new BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilderImpl(
        booleanValue, tagNumber);
  }

  public static class BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilderImpl
      implements BACnetNotificationParametersChangeOfDiscreteValueNewValue
          .BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
    private final BACnetApplicationTagBoolean booleanValue;
    private final Short tagNumber;

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilderImpl(
        BACnetApplicationTagBoolean booleanValue, Short tagNumber) {
      this.booleanValue = booleanValue;
      this.tagNumber = tagNumber;
    }

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
          bACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean =
              new BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(
                  openingTag, peekedTagHeader, closingTag, booleanValue, tagNumber);
      return bACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)) {
      return false;
    }
    BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean that =
        (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) o;
    return (getBooleanValue() == that.getBooleanValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBooleanValue());
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
