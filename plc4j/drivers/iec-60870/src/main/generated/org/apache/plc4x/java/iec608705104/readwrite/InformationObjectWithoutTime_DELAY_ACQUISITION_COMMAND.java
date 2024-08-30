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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND
    extends InformationObjectWithoutTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.DELAY_ACQUISITION_COMMAND;
  }

  // Properties.
  protected final TwoOctetBinaryTime cp16Time2a;

  public InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND(
      int address, TwoOctetBinaryTime cp16Time2a) {
    super(address);
    this.cp16Time2a = cp16Time2a;
  }

  public TwoOctetBinaryTime getCp16Time2a() {
    return cp16Time2a;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND");

    // Simple Field (cp16Time2a)
    writeSimpleField(
        "cp16Time2a",
        cp16Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (cp16Time2a)
    lengthInBits += cp16Time2a.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TwoOctetBinaryTime cp16Time2a =
        readSimpleField(
            "cp16Time2a",
            readComplex(() -> TwoOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND");
    // Create the instance
    return new InformationObjectWithoutTime_DELAY_ACQUISITION_COMMANDBuilderImpl(cp16Time2a);
  }

  public static class InformationObjectWithoutTime_DELAY_ACQUISITION_COMMANDBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final TwoOctetBinaryTime cp16Time2a;

    public InformationObjectWithoutTime_DELAY_ACQUISITION_COMMANDBuilderImpl(
        TwoOctetBinaryTime cp16Time2a) {
      this.cp16Time2a = cp16Time2a;
    }

    public InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND build(int address) {
      InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND
          informationObjectWithoutTime_DELAY_ACQUISITION_COMMAND =
              new InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND(address, cp16Time2a);
      return informationObjectWithoutTime_DELAY_ACQUISITION_COMMAND;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND)) {
      return false;
    }
    InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND that =
        (InformationObjectWithoutTime_DELAY_ACQUISITION_COMMAND) o;
    return (getCp16Time2a() == that.getCp16Time2a()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCp16Time2a());
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
