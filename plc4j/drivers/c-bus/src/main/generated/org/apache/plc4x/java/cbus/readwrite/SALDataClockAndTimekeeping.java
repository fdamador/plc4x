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
package org.apache.plc4x.java.cbus.readwrite;

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

public class SALDataClockAndTimekeeping extends SALData implements Message {

  // Accessors for discriminator values.
  public ApplicationId getApplicationId() {
    return ApplicationId.CLOCK_AND_TIMEKEEPING;
  }

  // Properties.
  protected final ClockAndTimekeepingData clockAndTimekeepingData;

  public SALDataClockAndTimekeeping(
      SALData salData, ClockAndTimekeepingData clockAndTimekeepingData) {
    super(salData);
    this.clockAndTimekeepingData = clockAndTimekeepingData;
  }

  public ClockAndTimekeepingData getClockAndTimekeepingData() {
    return clockAndTimekeepingData;
  }

  @Override
  protected void serializeSALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SALDataClockAndTimekeeping");

    // Simple Field (clockAndTimekeepingData)
    writeSimpleField(
        "clockAndTimekeepingData",
        clockAndTimekeepingData,
        new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SALDataClockAndTimekeeping");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SALDataClockAndTimekeeping _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (clockAndTimekeepingData)
    lengthInBits += clockAndTimekeepingData.getLengthInBits();

    return lengthInBits;
  }

  public static SALDataBuilder staticParseSALDataBuilder(
      ReadBuffer readBuffer, ApplicationId applicationId) throws ParseException {
    readBuffer.pullContext("SALDataClockAndTimekeeping");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ClockAndTimekeepingData clockAndTimekeepingData =
        readSimpleField(
            "clockAndTimekeepingData",
            readComplex(() -> ClockAndTimekeepingData.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SALDataClockAndTimekeeping");
    // Create the instance
    return new SALDataClockAndTimekeepingBuilderImpl(clockAndTimekeepingData);
  }

  public static class SALDataClockAndTimekeepingBuilderImpl implements SALData.SALDataBuilder {
    private final ClockAndTimekeepingData clockAndTimekeepingData;

    public SALDataClockAndTimekeepingBuilderImpl(ClockAndTimekeepingData clockAndTimekeepingData) {
      this.clockAndTimekeepingData = clockAndTimekeepingData;
    }

    public SALDataClockAndTimekeeping build(SALData salData) {
      SALDataClockAndTimekeeping sALDataClockAndTimekeeping =
          new SALDataClockAndTimekeeping(salData, clockAndTimekeepingData);
      return sALDataClockAndTimekeeping;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SALDataClockAndTimekeeping)) {
      return false;
    }
    SALDataClockAndTimekeeping that = (SALDataClockAndTimekeeping) o;
    return (getClockAndTimekeepingData() == that.getClockAndTimekeepingData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getClockAndTimekeepingData());
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
