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

public class OpcuaMessageRequest extends MessagePDU implements Message {

  // Accessors for discriminator values.
  public String getMessageType() {
    return (String) "MSG";
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final SecurityHeader securityHeader;
  protected final Payload message;

  public OpcuaMessageRequest(ChunkType chunk, SecurityHeader securityHeader, Payload message) {
    super(chunk);
    this.securityHeader = securityHeader;
    this.message = message;
  }

  public SecurityHeader getSecurityHeader() {
    return securityHeader;
  }

  public Payload getMessage() {
    return message;
  }

  @Override
  protected void serializeMessagePDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpcuaMessageRequest");

    // Simple Field (securityHeader)
    writeSimpleField("securityHeader", securityHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (message)
    writeSimpleField("message", message, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("OpcuaMessageRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpcuaMessageRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (securityHeader)
    lengthInBits += securityHeader.getLengthInBits();

    // Simple field (message)
    lengthInBits += message.getLengthInBits();

    return lengthInBits;
  }

  public static MessagePDUBuilder staticParseMessagePDUBuilder(
      ReadBuffer readBuffer, Long totalLength, Boolean response) throws ParseException {
    readBuffer.pullContext("OpcuaMessageRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SecurityHeader securityHeader =
        readSimpleField(
            "securityHeader",
            readComplex(() -> SecurityHeader.staticParse(readBuffer), readBuffer));

    Payload message =
        readSimpleField(
            "message",
            readComplex(
                () ->
                    Payload.staticParse(
                        readBuffer,
                        (boolean) (false),
                        (long) (((totalLength) - (securityHeader.getLengthInBytes())) - (16L))),
                readBuffer));

    readBuffer.closeContext("OpcuaMessageRequest");
    // Create the instance
    return new OpcuaMessageRequestBuilderImpl(securityHeader, message);
  }

  public static class OpcuaMessageRequestBuilderImpl implements MessagePDU.MessagePDUBuilder {
    private final SecurityHeader securityHeader;
    private final Payload message;

    public OpcuaMessageRequestBuilderImpl(SecurityHeader securityHeader, Payload message) {
      this.securityHeader = securityHeader;
      this.message = message;
    }

    public OpcuaMessageRequest build(ChunkType chunk) {
      OpcuaMessageRequest opcuaMessageRequest =
          new OpcuaMessageRequest(chunk, securityHeader, message);
      return opcuaMessageRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpcuaMessageRequest)) {
      return false;
    }
    OpcuaMessageRequest that = (OpcuaMessageRequest) o;
    return (getSecurityHeader() == that.getSecurityHeader())
        && (getMessage() == that.getMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSecurityHeader(), getMessage());
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
