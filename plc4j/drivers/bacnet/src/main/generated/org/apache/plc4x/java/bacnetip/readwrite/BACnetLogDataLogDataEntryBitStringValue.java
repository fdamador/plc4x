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

public class BACnetLogDataLogDataEntryBitStringValue extends BACnetLogDataLogDataEntry
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagBitString bitStringValue;

  public BACnetLogDataLogDataEntryBitStringValue(
      BACnetTagHeader peekedTagHeader, BACnetContextTagBitString bitStringValue) {
    super(peekedTagHeader);
    this.bitStringValue = bitStringValue;
  }

  public BACnetContextTagBitString getBitStringValue() {
    return bitStringValue;
  }

  @Override
  protected void serializeBACnetLogDataLogDataEntryChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogDataLogDataEntryBitStringValue");

    // Simple Field (bitStringValue)
    writeSimpleField("bitStringValue", bitStringValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogDataLogDataEntryBitStringValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogDataLogDataEntryBitStringValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (bitStringValue)
    lengthInBits += bitStringValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogDataLogDataEntryBuilder staticParseBACnetLogDataLogDataEntryBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogDataLogDataEntryBitStringValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagBitString bitStringValue =
        readSimpleField(
            "bitStringValue",
            readComplex(
                () ->
                    (BACnetContextTagBitString)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (5), (BACnetDataType) (BACnetDataType.BIT_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetLogDataLogDataEntryBitStringValue");
    // Create the instance
    return new BACnetLogDataLogDataEntryBitStringValueBuilderImpl(bitStringValue);
  }

  public static class BACnetLogDataLogDataEntryBitStringValueBuilderImpl
      implements BACnetLogDataLogDataEntry.BACnetLogDataLogDataEntryBuilder {
    private final BACnetContextTagBitString bitStringValue;

    public BACnetLogDataLogDataEntryBitStringValueBuilderImpl(
        BACnetContextTagBitString bitStringValue) {
      this.bitStringValue = bitStringValue;
    }

    public BACnetLogDataLogDataEntryBitStringValue build(BACnetTagHeader peekedTagHeader) {
      BACnetLogDataLogDataEntryBitStringValue bACnetLogDataLogDataEntryBitStringValue =
          new BACnetLogDataLogDataEntryBitStringValue(peekedTagHeader, bitStringValue);
      return bACnetLogDataLogDataEntryBitStringValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogDataLogDataEntryBitStringValue)) {
      return false;
    }
    BACnetLogDataLogDataEntryBitStringValue that = (BACnetLogDataLogDataEntryBitStringValue) o;
    return (getBitStringValue() == that.getBitStringValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBitStringValue());
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
