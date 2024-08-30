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

public class BVLCDistributeBroadcastToNetwork extends BVLC implements Message {

  // Accessors for discriminator values.
  public Short getBvlcFunction() {
    return (short) 0x09;
  }

  // Properties.
  protected final NPDU npdu;

  // Arguments.
  protected final Integer bvlcPayloadLength;

  public BVLCDistributeBroadcastToNetwork(NPDU npdu, Integer bvlcPayloadLength) {
    super();
    this.npdu = npdu;
    this.bvlcPayloadLength = bvlcPayloadLength;
  }

  public NPDU getNpdu() {
    return npdu;
  }

  @Override
  protected void serializeBVLCChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BVLCDistributeBroadcastToNetwork");

    // Simple Field (npdu)
    writeSimpleField(
        "npdu",
        npdu,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("BVLCDistributeBroadcastToNetwork");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BVLCDistributeBroadcastToNetwork _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (npdu)
    lengthInBits += npdu.getLengthInBits();

    return lengthInBits;
  }

  public static BVLCBuilder staticParseBVLCBuilder(ReadBuffer readBuffer, Integer bvlcPayloadLength)
      throws ParseException {
    readBuffer.pullContext("BVLCDistributeBroadcastToNetwork");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NPDU npdu =
        readSimpleField(
            "npdu",
            readComplex(() -> NPDU.staticParse(readBuffer, (int) (bvlcPayloadLength)), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("BVLCDistributeBroadcastToNetwork");
    // Create the instance
    return new BVLCDistributeBroadcastToNetworkBuilderImpl(npdu, bvlcPayloadLength);
  }

  public static class BVLCDistributeBroadcastToNetworkBuilderImpl implements BVLC.BVLCBuilder {
    private final NPDU npdu;
    private final Integer bvlcPayloadLength;

    public BVLCDistributeBroadcastToNetworkBuilderImpl(NPDU npdu, Integer bvlcPayloadLength) {
      this.npdu = npdu;
      this.bvlcPayloadLength = bvlcPayloadLength;
    }

    public BVLCDistributeBroadcastToNetwork build() {
      BVLCDistributeBroadcastToNetwork bVLCDistributeBroadcastToNetwork =
          new BVLCDistributeBroadcastToNetwork(npdu, bvlcPayloadLength);
      return bVLCDistributeBroadcastToNetwork;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BVLCDistributeBroadcastToNetwork)) {
      return false;
    }
    BVLCDistributeBroadcastToNetwork that = (BVLCDistributeBroadcastToNetwork) o;
    return (getNpdu() == that.getNpdu()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNpdu());
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
