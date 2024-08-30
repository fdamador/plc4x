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

public class BrowsePathTarget extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "548";
  }

  // Properties.
  protected final ExpandedNodeId targetId;
  protected final long remainingPathIndex;

  public BrowsePathTarget(ExpandedNodeId targetId, long remainingPathIndex) {
    super();
    this.targetId = targetId;
    this.remainingPathIndex = remainingPathIndex;
  }

  public ExpandedNodeId getTargetId() {
    return targetId;
  }

  public long getRemainingPathIndex() {
    return remainingPathIndex;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BrowsePathTarget");

    // Simple Field (targetId)
    writeSimpleField("targetId", targetId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (remainingPathIndex)
    writeSimpleField("remainingPathIndex", remainingPathIndex, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("BrowsePathTarget");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BrowsePathTarget _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (targetId)
    lengthInBits += targetId.getLengthInBits();

    // Simple field (remainingPathIndex)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("BrowsePathTarget");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExpandedNodeId targetId =
        readSimpleField(
            "targetId", readComplex(() -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    long remainingPathIndex =
        readSimpleField("remainingPathIndex", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("BrowsePathTarget");
    // Create the instance
    return new BrowsePathTargetBuilderImpl(targetId, remainingPathIndex);
  }

  public static class BrowsePathTargetBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExpandedNodeId targetId;
    private final long remainingPathIndex;

    public BrowsePathTargetBuilderImpl(ExpandedNodeId targetId, long remainingPathIndex) {
      this.targetId = targetId;
      this.remainingPathIndex = remainingPathIndex;
    }

    public BrowsePathTarget build() {
      BrowsePathTarget browsePathTarget = new BrowsePathTarget(targetId, remainingPathIndex);
      return browsePathTarget;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BrowsePathTarget)) {
      return false;
    }
    BrowsePathTarget that = (BrowsePathTarget) o;
    return (getTargetId() == that.getTargetId())
        && (getRemainingPathIndex() == that.getRemainingPathIndex())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTargetId(), getRemainingPathIndex());
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
