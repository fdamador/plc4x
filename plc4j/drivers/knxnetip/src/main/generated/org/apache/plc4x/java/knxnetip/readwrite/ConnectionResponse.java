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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ConnectionResponse extends KnxNetIpMessage implements Message {

  // Accessors for discriminator values.
  public Integer getMsgType() {
    return (int) 0x0206;
  }

  // Properties.
  protected final short communicationChannelId;
  protected final Status status;
  protected final HPAIDataEndpoint hpaiDataEndpoint;
  protected final ConnectionResponseDataBlock connectionResponseDataBlock;

  public ConnectionResponse(
      short communicationChannelId,
      Status status,
      HPAIDataEndpoint hpaiDataEndpoint,
      ConnectionResponseDataBlock connectionResponseDataBlock) {
    super();
    this.communicationChannelId = communicationChannelId;
    this.status = status;
    this.hpaiDataEndpoint = hpaiDataEndpoint;
    this.connectionResponseDataBlock = connectionResponseDataBlock;
  }

  public short getCommunicationChannelId() {
    return communicationChannelId;
  }

  public Status getStatus() {
    return status;
  }

  public HPAIDataEndpoint getHpaiDataEndpoint() {
    return hpaiDataEndpoint;
  }

  public ConnectionResponseDataBlock getConnectionResponseDataBlock() {
    return connectionResponseDataBlock;
  }

  @Override
  protected void serializeKnxNetIpMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ConnectionResponse");

    // Simple Field (communicationChannelId)
    writeSimpleField(
        "communicationChannelId",
        communicationChannelId,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (status)
    writeSimpleEnumField(
        "status",
        "Status",
        status,
        new DataWriterEnumDefault<>(
            Status::getValue, Status::name, writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Optional Field (hpaiDataEndpoint) (Can be skipped, if the value is null)
    writeOptionalField(
        "hpaiDataEndpoint",
        hpaiDataEndpoint,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Optional Field (connectionResponseDataBlock) (Can be skipped, if the value is null)
    writeOptionalField(
        "connectionResponseDataBlock",
        connectionResponseDataBlock,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("ConnectionResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConnectionResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (communicationChannelId)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Optional Field (hpaiDataEndpoint)
    if (hpaiDataEndpoint != null) {
      lengthInBits += hpaiDataEndpoint.getLengthInBits();
    }

    // Optional Field (connectionResponseDataBlock)
    if (connectionResponseDataBlock != null) {
      lengthInBits += connectionResponseDataBlock.getLengthInBits();
    }

    return lengthInBits;
  }

  public static KnxNetIpMessageBuilder staticParseKnxNetIpMessageBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("ConnectionResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short communicationChannelId =
        readSimpleField(
            "communicationChannelId",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Status status =
        readEnumField(
            "status",
            "Status",
            new DataReaderEnumDefault<>(Status::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    HPAIDataEndpoint hpaiDataEndpoint =
        readOptionalField(
            "hpaiDataEndpoint",
            readComplex(() -> HPAIDataEndpoint.staticParse(readBuffer), readBuffer),
            (status) == (Status.NO_ERROR),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    ConnectionResponseDataBlock connectionResponseDataBlock =
        readOptionalField(
            "connectionResponseDataBlock",
            readComplex(() -> ConnectionResponseDataBlock.staticParse(readBuffer), readBuffer),
            (status) == (Status.NO_ERROR),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("ConnectionResponse");
    // Create the instance
    return new ConnectionResponseBuilderImpl(
        communicationChannelId, status, hpaiDataEndpoint, connectionResponseDataBlock);
  }

  public static class ConnectionResponseBuilderImpl
      implements KnxNetIpMessage.KnxNetIpMessageBuilder {
    private final short communicationChannelId;
    private final Status status;
    private final HPAIDataEndpoint hpaiDataEndpoint;
    private final ConnectionResponseDataBlock connectionResponseDataBlock;

    public ConnectionResponseBuilderImpl(
        short communicationChannelId,
        Status status,
        HPAIDataEndpoint hpaiDataEndpoint,
        ConnectionResponseDataBlock connectionResponseDataBlock) {
      this.communicationChannelId = communicationChannelId;
      this.status = status;
      this.hpaiDataEndpoint = hpaiDataEndpoint;
      this.connectionResponseDataBlock = connectionResponseDataBlock;
    }

    public ConnectionResponse build() {
      ConnectionResponse connectionResponse =
          new ConnectionResponse(
              communicationChannelId, status, hpaiDataEndpoint, connectionResponseDataBlock);
      return connectionResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConnectionResponse)) {
      return false;
    }
    ConnectionResponse that = (ConnectionResponse) o;
    return (getCommunicationChannelId() == that.getCommunicationChannelId())
        && (getStatus() == that.getStatus())
        && (getHpaiDataEndpoint() == that.getHpaiDataEndpoint())
        && (getConnectionResponseDataBlock() == that.getConnectionResponseDataBlock())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getCommunicationChannelId(),
        getStatus(),
        getHpaiDataEndpoint(),
        getConnectionResponseDataBlock());
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
