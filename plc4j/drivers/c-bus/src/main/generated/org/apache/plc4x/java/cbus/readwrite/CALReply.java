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

public abstract class CALReply implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final byte calType;
  protected final CALData calData;

  // Arguments.
  protected final CBusOptions cBusOptions;
  protected final RequestContext requestContext;

  public CALReply(
      byte calType, CALData calData, CBusOptions cBusOptions, RequestContext requestContext) {
    super();
    this.calType = calType;
    this.calData = calData;
    this.cBusOptions = cBusOptions;
    this.requestContext = requestContext;
  }

  public byte getCalType() {
    return calType;
  }

  public CALData getCalData() {
    return calData;
  }

  protected abstract void serializeCALReplyChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CALReply");

    // Switch field (Serialize the sub-type)
    serializeCALReplyChild(writeBuffer);

    // Simple Field (calData)
    writeSimpleField("calData", calData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CALReply");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CALReply _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    // Simple field (calData)
    lengthInBits += calData.getLengthInBits();

    return lengthInBits;
  }

  public static CALReply staticParse(
      ReadBuffer readBuffer, CBusOptions cBusOptions, RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("CALReply");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte calType = readPeekField("calType", readByte(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CALReplyBuilder builder = null;
    if (EvaluationHelper.equals(calType, (byte) 0x86)) {
      builder = CALReplyLong.staticParseCALReplyBuilder(readBuffer, cBusOptions, requestContext);
    } else {
      builder = CALReplyShort.staticParseCALReplyBuilder(readBuffer, cBusOptions, requestContext);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "calType=" + calType + "]");
    }

    CALData calData =
        readSimpleField(
            "calData",
            readComplex(
                () -> CALData.staticParse(readBuffer, (RequestContext) (requestContext)),
                readBuffer));

    readBuffer.closeContext("CALReply");
    // Create the instance
    CALReply _cALReply = builder.build(calType, calData, cBusOptions, requestContext);
    return _cALReply;
  }

  public interface CALReplyBuilder {
    CALReply build(
        byte calType, CALData calData, CBusOptions cBusOptions, RequestContext requestContext);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CALReply)) {
      return false;
    }
    CALReply that = (CALReply) o;
    return (getCalType() == that.getCalType()) && (getCalData() == that.getCalData()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCalType(), getCalData());
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
