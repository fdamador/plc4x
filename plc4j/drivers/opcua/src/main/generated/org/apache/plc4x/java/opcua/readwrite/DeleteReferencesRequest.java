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

public class DeleteReferencesRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "506";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfReferencesToDelete;
  protected final List<ExtensionObjectDefinition> referencesToDelete;

  public DeleteReferencesRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfReferencesToDelete,
      List<ExtensionObjectDefinition> referencesToDelete) {
    super();
    this.requestHeader = requestHeader;
    this.noOfReferencesToDelete = noOfReferencesToDelete;
    this.referencesToDelete = referencesToDelete;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfReferencesToDelete() {
    return noOfReferencesToDelete;
  }

  public List<ExtensionObjectDefinition> getReferencesToDelete() {
    return referencesToDelete;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeleteReferencesRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfReferencesToDelete)
    writeSimpleField(
        "noOfReferencesToDelete", noOfReferencesToDelete, writeSignedInt(writeBuffer, 32));

    // Array Field (referencesToDelete)
    writeComplexTypeArrayField("referencesToDelete", referencesToDelete, writeBuffer);

    writeBuffer.popContext("DeleteReferencesRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DeleteReferencesRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfReferencesToDelete)
    lengthInBits += 32;

    // Array field
    if (referencesToDelete != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : referencesToDelete) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= referencesToDelete.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("DeleteReferencesRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfReferencesToDelete =
        readSimpleField("noOfReferencesToDelete", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> referencesToDelete =
        readCountArrayField(
            "referencesToDelete",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("387")),
                readBuffer),
            noOfReferencesToDelete);

    readBuffer.closeContext("DeleteReferencesRequest");
    // Create the instance
    return new DeleteReferencesRequestBuilderImpl(
        requestHeader, noOfReferencesToDelete, referencesToDelete);
  }

  public static class DeleteReferencesRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfReferencesToDelete;
    private final List<ExtensionObjectDefinition> referencesToDelete;

    public DeleteReferencesRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        int noOfReferencesToDelete,
        List<ExtensionObjectDefinition> referencesToDelete) {
      this.requestHeader = requestHeader;
      this.noOfReferencesToDelete = noOfReferencesToDelete;
      this.referencesToDelete = referencesToDelete;
    }

    public DeleteReferencesRequest build() {
      DeleteReferencesRequest deleteReferencesRequest =
          new DeleteReferencesRequest(requestHeader, noOfReferencesToDelete, referencesToDelete);
      return deleteReferencesRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeleteReferencesRequest)) {
      return false;
    }
    DeleteReferencesRequest that = (DeleteReferencesRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfReferencesToDelete() == that.getNoOfReferencesToDelete())
        && (getReferencesToDelete() == that.getReferencesToDelete())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestHeader(), getNoOfReferencesToDelete(), getReferencesToDelete());
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
