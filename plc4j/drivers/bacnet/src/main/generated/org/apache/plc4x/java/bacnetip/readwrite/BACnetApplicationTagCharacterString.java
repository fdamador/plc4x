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

public class BACnetApplicationTagCharacterString extends BACnetApplicationTag implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetTagPayloadCharacterString payload;

  public BACnetApplicationTagCharacterString(
      BACnetTagHeader header, BACnetTagPayloadCharacterString payload) {
    super(header);
    this.payload = payload;
  }

  public BACnetTagPayloadCharacterString getPayload() {
    return payload;
  }

  public String getValue() {
    return String.valueOf(getPayload().getValue());
  }

  @Override
  protected void serializeBACnetApplicationTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetApplicationTagCharacterString");

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    String value = getValue();
    writeBuffer.writeVirtual("value", value);

    writeBuffer.popContext("BACnetApplicationTagCharacterString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetApplicationTagCharacterString _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetApplicationTagBuilder staticParseBACnetApplicationTagBuilder(
      ReadBuffer readBuffer, BACnetTagHeader header) throws ParseException {
    readBuffer.pullContext("BACnetApplicationTagCharacterString");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagPayloadCharacterString payload =
        readSimpleField(
            "payload",
            readComplex(
                () ->
                    BACnetTagPayloadCharacterString.staticParse(
                        readBuffer, (long) (header.getActualLength())),
                readBuffer));
    String value = readVirtualField("value", String.class, payload.getValue());

    readBuffer.closeContext("BACnetApplicationTagCharacterString");
    // Create the instance
    return new BACnetApplicationTagCharacterStringBuilderImpl(payload);
  }

  public static class BACnetApplicationTagCharacterStringBuilderImpl
      implements BACnetApplicationTag.BACnetApplicationTagBuilder {
    private final BACnetTagPayloadCharacterString payload;

    public BACnetApplicationTagCharacterStringBuilderImpl(BACnetTagPayloadCharacterString payload) {
      this.payload = payload;
    }

    public BACnetApplicationTagCharacterString build(BACnetTagHeader header) {
      BACnetApplicationTagCharacterString bACnetApplicationTagCharacterString =
          new BACnetApplicationTagCharacterString(header, payload);
      return bACnetApplicationTagCharacterString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetApplicationTagCharacterString)) {
      return false;
    }
    BACnetApplicationTagCharacterString that = (BACnetApplicationTagCharacterString) o;
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
