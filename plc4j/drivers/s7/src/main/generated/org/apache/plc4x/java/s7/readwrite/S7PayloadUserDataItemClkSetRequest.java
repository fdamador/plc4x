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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadUserDataItemClkSetRequest extends S7PayloadUserDataItem implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x07;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x04;
  }

  public Short getCpuSubfunction() {
    return (short) 0x04;
  }

  // Properties.
  protected final DateAndTime timeStamp;

  public S7PayloadUserDataItemClkSetRequest(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      DateAndTime timeStamp) {
    super(returnCode, transportSize, dataLength);
    this.timeStamp = timeStamp;
  }

  public DateAndTime getTimeStamp() {
    return timeStamp;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7PayloadUserDataItemClkSetRequest");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (timeStamp)
    writeSimpleField("timeStamp", timeStamp, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("S7PayloadUserDataItemClkSetRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemClkSetRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (timeStamp)
    lengthInBits += timeStamp.getLengthInBits();

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemClkSetRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    Short reservedField1 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    DateAndTime timeStamp =
        readSimpleField(
            "timeStamp", readComplex(() -> DateAndTime.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("S7PayloadUserDataItemClkSetRequest");
    // Create the instance
    return new S7PayloadUserDataItemClkSetRequestBuilderImpl(timeStamp);
  }

  public static class S7PayloadUserDataItemClkSetRequestBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final DateAndTime timeStamp;

    public S7PayloadUserDataItemClkSetRequestBuilderImpl(DateAndTime timeStamp) {
      this.timeStamp = timeStamp;
    }

    public S7PayloadUserDataItemClkSetRequest build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemClkSetRequest s7PayloadUserDataItemClkSetRequest =
          new S7PayloadUserDataItemClkSetRequest(returnCode, transportSize, dataLength, timeStamp);
      return s7PayloadUserDataItemClkSetRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemClkSetRequest)) {
      return false;
    }
    S7PayloadUserDataItemClkSetRequest that = (S7PayloadUserDataItemClkSetRequest) o;
    return (getTimeStamp() == that.getTimeStamp()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTimeStamp());
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
