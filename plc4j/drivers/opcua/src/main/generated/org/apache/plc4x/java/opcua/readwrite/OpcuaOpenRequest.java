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
package org.apache.plc4x.java.opcua.readwrite;

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

public class OpcuaOpenRequest extends MessagePDU implements Message {

  // Accessors for discriminator values.
  public String getMessageType() {
    return (String) "OPN";
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final OpenChannelMessage openRequest;
  protected final Payload message;

  public OpcuaOpenRequest(ChunkType chunk, OpenChannelMessage openRequest, Payload message) {
    super(chunk);
    this.openRequest = openRequest;
    this.message = message;
  }

  public OpenChannelMessage getOpenRequest() {
    return openRequest;
  }

  public Payload getMessage() {
    return message;
  }

  @Override
  protected void serializeMessagePDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpcuaOpenRequest");

    // Simple Field (openRequest)
    writeSimpleField("openRequest", openRequest, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (message)
    writeSimpleField("message", message, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("OpcuaOpenRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpcuaOpenRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openRequest)
    lengthInBits += openRequest.getLengthInBits();

    // Simple field (message)
    lengthInBits += message.getLengthInBits();

    return lengthInBits;
  }

  public static MessagePDUBuilder staticParseMessagePDUBuilder(
      ReadBuffer readBuffer, Long totalLength, Boolean response) throws ParseException {
    readBuffer.pullContext("OpcuaOpenRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    OpenChannelMessage openRequest =
        readSimpleField(
            "openRequest",
            readComplex(
                () -> OpenChannelMessage.staticParse(readBuffer, (boolean) (response)),
                readBuffer));

    Payload message =
        readSimpleField(
            "message",
            readComplex(
                () ->
                    Payload.staticParse(
                        readBuffer,
                        (boolean) (false),
                        (long) (((totalLength) - (openRequest.getLengthInBytes())) - (16L))),
                readBuffer));

    readBuffer.closeContext("OpcuaOpenRequest");
    // Create the instance
    return new OpcuaOpenRequestBuilderImpl(openRequest, message);
  }

  public static class OpcuaOpenRequestBuilderImpl implements MessagePDU.MessagePDUBuilder {
    private final OpenChannelMessage openRequest;
    private final Payload message;

    public OpcuaOpenRequestBuilderImpl(OpenChannelMessage openRequest, Payload message) {
      this.openRequest = openRequest;
      this.message = message;
    }

    public OpcuaOpenRequest build(ChunkType chunk) {
      OpcuaOpenRequest opcuaOpenRequest = new OpcuaOpenRequest(chunk, openRequest, message);
      return opcuaOpenRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpcuaOpenRequest)) {
      return false;
    }
    OpcuaOpenRequest that = (OpcuaOpenRequest) o;
    return (getOpenRequest() == that.getOpenRequest())
        && (getMessage() == that.getMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOpenRequest(), getMessage());
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
