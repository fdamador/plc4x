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

public class BACnetChannelValueEnumerated extends BACnetChannelValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagEnumerated enumeratedValue;

  public BACnetChannelValueEnumerated(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagEnumerated enumeratedValue) {
    super(peekedTagHeader);
    this.enumeratedValue = enumeratedValue;
  }

  public BACnetApplicationTagEnumerated getEnumeratedValue() {
    return enumeratedValue;
  }

  @Override
  protected void serializeBACnetChannelValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetChannelValueEnumerated");

    // Simple Field (enumeratedValue)
    writeSimpleField(
        "enumeratedValue", enumeratedValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetChannelValueEnumerated");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetChannelValueEnumerated _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (enumeratedValue)
    lengthInBits += enumeratedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetChannelValueBuilder staticParseBACnetChannelValueBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetChannelValueEnumerated");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagEnumerated enumeratedValue =
        readSimpleField(
            "enumeratedValue",
            readComplex(
                () -> (BACnetApplicationTagEnumerated) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetChannelValueEnumerated");
    // Create the instance
    return new BACnetChannelValueEnumeratedBuilderImpl(enumeratedValue);
  }

  public static class BACnetChannelValueEnumeratedBuilderImpl
      implements BACnetChannelValue.BACnetChannelValueBuilder {
    private final BACnetApplicationTagEnumerated enumeratedValue;

    public BACnetChannelValueEnumeratedBuilderImpl(BACnetApplicationTagEnumerated enumeratedValue) {
      this.enumeratedValue = enumeratedValue;
    }

    public BACnetChannelValueEnumerated build(BACnetTagHeader peekedTagHeader) {
      BACnetChannelValueEnumerated bACnetChannelValueEnumerated =
          new BACnetChannelValueEnumerated(peekedTagHeader, enumeratedValue);
      return bACnetChannelValueEnumerated;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetChannelValueEnumerated)) {
      return false;
    }
    BACnetChannelValueEnumerated that = (BACnetChannelValueEnumerated) o;
    return (getEnumeratedValue() == that.getEnumeratedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEnumeratedValue());
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
