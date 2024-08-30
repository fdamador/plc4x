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

public class BitFieldDefinition extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "32423";
  }

  // Properties.
  protected final PascalString name;
  protected final LocalizedText description;
  protected final long startingBitPosition;
  protected final long endingBitPosition;

  public BitFieldDefinition(
      PascalString name,
      LocalizedText description,
      long startingBitPosition,
      long endingBitPosition) {
    super();
    this.name = name;
    this.description = description;
    this.startingBitPosition = startingBitPosition;
    this.endingBitPosition = endingBitPosition;
  }

  public PascalString getName() {
    return name;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public long getStartingBitPosition() {
    return startingBitPosition;
  }

  public long getEndingBitPosition() {
    return endingBitPosition;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BitFieldDefinition");

    // Simple Field (name)
    writeSimpleField("name", name, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Reserved Field (reserved)
    writeReservedField("reserved", (boolean) false, writeBoolean(writeBuffer));

    // Simple Field (startingBitPosition)
    writeSimpleField(
        "startingBitPosition", startingBitPosition, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (endingBitPosition)
    writeSimpleField("endingBitPosition", endingBitPosition, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("BitFieldDefinition");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BitFieldDefinition _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (startingBitPosition)
    lengthInBits += 32;

    // Simple field (endingBitPosition)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("BitFieldDefinition");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    Boolean reservedField1 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    long startingBitPosition =
        readSimpleField("startingBitPosition", readUnsignedLong(readBuffer, 32));

    long endingBitPosition = readSimpleField("endingBitPosition", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("BitFieldDefinition");
    // Create the instance
    return new BitFieldDefinitionBuilderImpl(
        name, description, startingBitPosition, endingBitPosition);
  }

  public static class BitFieldDefinitionBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final LocalizedText description;
    private final long startingBitPosition;
    private final long endingBitPosition;

    public BitFieldDefinitionBuilderImpl(
        PascalString name,
        LocalizedText description,
        long startingBitPosition,
        long endingBitPosition) {
      this.name = name;
      this.description = description;
      this.startingBitPosition = startingBitPosition;
      this.endingBitPosition = endingBitPosition;
    }

    public BitFieldDefinition build() {
      BitFieldDefinition bitFieldDefinition =
          new BitFieldDefinition(name, description, startingBitPosition, endingBitPosition);
      return bitFieldDefinition;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BitFieldDefinition)) {
      return false;
    }
    BitFieldDefinition that = (BitFieldDefinition) o;
    return (getName() == that.getName())
        && (getDescription() == that.getDescription())
        && (getStartingBitPosition() == that.getStartingBitPosition())
        && (getEndingBitPosition() == that.getEndingBitPosition())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getDescription(),
        getStartingBitPosition(),
        getEndingBitPosition());
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
