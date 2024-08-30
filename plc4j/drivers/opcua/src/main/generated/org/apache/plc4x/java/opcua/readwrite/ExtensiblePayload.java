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

public class ExtensiblePayload extends Payload implements Message {

  // Accessors for discriminator values.
  public Boolean getExtensible() {
    return (boolean) true;
  }

  // Properties.
  protected final ExtensionObject payload;

  public ExtensiblePayload(SequenceHeader sequenceHeader, ExtensionObject payload) {
    super(sequenceHeader);
    this.payload = payload;
  }

  public ExtensionObject getPayload() {
    return payload;
  }

  @Override
  protected void serializePayloadChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ExtensiblePayload");

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ExtensiblePayload");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ExtensiblePayload _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static PayloadBuilder staticParsePayloadBuilder(
      ReadBuffer readBuffer, Boolean extensible, Long byteCount) throws ParseException {
    readBuffer.pullContext("ExtensiblePayload");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObject payload =
        readSimpleField(
            "payload",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (false)), readBuffer));

    readBuffer.closeContext("ExtensiblePayload");
    // Create the instance
    return new ExtensiblePayloadBuilderImpl(payload);
  }

  public static class ExtensiblePayloadBuilderImpl implements Payload.PayloadBuilder {
    private final ExtensionObject payload;

    public ExtensiblePayloadBuilderImpl(ExtensionObject payload) {
      this.payload = payload;
    }

    public ExtensiblePayload build(SequenceHeader sequenceHeader) {
      ExtensiblePayload extensiblePayload = new ExtensiblePayload(sequenceHeader, payload);
      return extensiblePayload;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ExtensiblePayload)) {
      return false;
    }
    ExtensiblePayload that = (ExtensiblePayload) o;
    return (getPayload() == that.getPayload()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPayload());
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
