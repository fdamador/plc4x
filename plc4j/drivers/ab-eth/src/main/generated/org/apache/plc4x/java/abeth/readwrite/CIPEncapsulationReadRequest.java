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
package org.apache.plc4x.java.abeth.readwrite;

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

public class CIPEncapsulationReadRequest extends CIPEncapsulationPacket implements Message {

  // Accessors for discriminator values.
  public Integer getCommandType() {
    return (int) 0x0107;
  }

  // Properties.
  protected final DF1RequestMessage request;

  public CIPEncapsulationReadRequest(
      long sessionHandle,
      long status,
      List<Short> senderContext,
      long options,
      DF1RequestMessage request) {
    super(sessionHandle, status, senderContext, options);
    this.request = request;
  }

  public DF1RequestMessage getRequest() {
    return request;
  }

  @Override
  protected void serializeCIPEncapsulationPacketChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CIPEncapsulationReadRequest");

    // Simple Field (request)
    writeSimpleField(
        "request",
        request,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("CIPEncapsulationReadRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CIPEncapsulationReadRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (request)
    lengthInBits += request.getLengthInBits();

    return lengthInBits;
  }

  public static CIPEncapsulationPacketBuilder staticParseCIPEncapsulationPacketBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CIPEncapsulationReadRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DF1RequestMessage request =
        readSimpleField(
            "request",
            readComplex(() -> DF1RequestMessage.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("CIPEncapsulationReadRequest");
    // Create the instance
    return new CIPEncapsulationReadRequestBuilderImpl(request);
  }

  public static class CIPEncapsulationReadRequestBuilderImpl
      implements CIPEncapsulationPacket.CIPEncapsulationPacketBuilder {
    private final DF1RequestMessage request;

    public CIPEncapsulationReadRequestBuilderImpl(DF1RequestMessage request) {
      this.request = request;
    }

    public CIPEncapsulationReadRequest build(
        long sessionHandle, long status, List<Short> senderContext, long options) {
      CIPEncapsulationReadRequest cIPEncapsulationReadRequest =
          new CIPEncapsulationReadRequest(sessionHandle, status, senderContext, options, request);
      return cIPEncapsulationReadRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CIPEncapsulationReadRequest)) {
      return false;
    }
    CIPEncapsulationReadRequest that = (CIPEncapsulationReadRequest) o;
    return (getRequest() == that.getRequest()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequest());
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
