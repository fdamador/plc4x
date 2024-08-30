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

public class BACnetLogMultipleRecord implements Message {

  // Properties.
  protected final BACnetDateTimeEnclosed timestamp;
  protected final BACnetLogData logData;

  public BACnetLogMultipleRecord(BACnetDateTimeEnclosed timestamp, BACnetLogData logData) {
    super();
    this.timestamp = timestamp;
    this.logData = logData;
  }

  public BACnetDateTimeEnclosed getTimestamp() {
    return timestamp;
  }

  public BACnetLogData getLogData() {
    return logData;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogMultipleRecord");

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (logData)
    writeSimpleField("logData", logData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogMultipleRecord");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetLogMultipleRecord _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timestamp)
    lengthInBits += timestamp.getLengthInBits();

    // Simple field (logData)
    lengthInBits += logData.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogMultipleRecord staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogMultipleRecord");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetDateTimeEnclosed timestamp =
        readSimpleField(
            "timestamp",
            readComplex(
                () -> BACnetDateTimeEnclosed.staticParse(readBuffer, (short) (0)), readBuffer));

    BACnetLogData logData =
        readSimpleField(
            "logData",
            readComplex(() -> BACnetLogData.staticParse(readBuffer, (short) (1)), readBuffer));

    readBuffer.closeContext("BACnetLogMultipleRecord");
    // Create the instance
    BACnetLogMultipleRecord _bACnetLogMultipleRecord;
    _bACnetLogMultipleRecord = new BACnetLogMultipleRecord(timestamp, logData);
    return _bACnetLogMultipleRecord;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogMultipleRecord)) {
      return false;
    }
    BACnetLogMultipleRecord that = (BACnetLogMultipleRecord) o;
    return (getTimestamp() == that.getTimestamp()) && (getLogData() == that.getLogData()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTimestamp(), getLogData());
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
