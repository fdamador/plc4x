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

public class ResponseHeader extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "394";
  }

  // Properties.
  protected final long timestamp;
  protected final long requestHandle;
  protected final StatusCode serviceResult;
  protected final DiagnosticInfo serviceDiagnostics;
  protected final int noOfStringTable;
  protected final List<PascalString> stringTable;
  protected final ExtensionObject additionalHeader;

  public ResponseHeader(
      long timestamp,
      long requestHandle,
      StatusCode serviceResult,
      DiagnosticInfo serviceDiagnostics,
      int noOfStringTable,
      List<PascalString> stringTable,
      ExtensionObject additionalHeader) {
    super();
    this.timestamp = timestamp;
    this.requestHandle = requestHandle;
    this.serviceResult = serviceResult;
    this.serviceDiagnostics = serviceDiagnostics;
    this.noOfStringTable = noOfStringTable;
    this.stringTable = stringTable;
    this.additionalHeader = additionalHeader;
  }

  public long getTimestamp() {
    return timestamp;
  }

  public long getRequestHandle() {
    return requestHandle;
  }

  public StatusCode getServiceResult() {
    return serviceResult;
  }

  public DiagnosticInfo getServiceDiagnostics() {
    return serviceDiagnostics;
  }

  public int getNoOfStringTable() {
    return noOfStringTable;
  }

  public List<PascalString> getStringTable() {
    return stringTable;
  }

  public ExtensionObject getAdditionalHeader() {
    return additionalHeader;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ResponseHeader");

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, writeSignedLong(writeBuffer, 64));

    // Simple Field (requestHandle)
    writeSimpleField("requestHandle", requestHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (serviceResult)
    writeSimpleField("serviceResult", serviceResult, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (serviceDiagnostics)
    writeSimpleField(
        "serviceDiagnostics", serviceDiagnostics, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfStringTable)
    writeSimpleField("noOfStringTable", noOfStringTable, writeSignedInt(writeBuffer, 32));

    // Array Field (stringTable)
    writeComplexTypeArrayField("stringTable", stringTable, writeBuffer);

    // Simple Field (additionalHeader)
    writeSimpleField(
        "additionalHeader", additionalHeader, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ResponseHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ResponseHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timestamp)
    lengthInBits += 64;

    // Simple field (requestHandle)
    lengthInBits += 32;

    // Simple field (serviceResult)
    lengthInBits += serviceResult.getLengthInBits();

    // Simple field (serviceDiagnostics)
    lengthInBits += serviceDiagnostics.getLengthInBits();

    // Simple field (noOfStringTable)
    lengthInBits += 32;

    // Array field
    if (stringTable != null) {
      int i = 0;
      for (PascalString element : stringTable) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= stringTable.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (additionalHeader)
    lengthInBits += additionalHeader.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ResponseHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long timestamp = readSimpleField("timestamp", readSignedLong(readBuffer, 64));

    long requestHandle = readSimpleField("requestHandle", readUnsignedLong(readBuffer, 32));

    StatusCode serviceResult =
        readSimpleField(
            "serviceResult", readComplex(() -> StatusCode.staticParse(readBuffer), readBuffer));

    DiagnosticInfo serviceDiagnostics =
        readSimpleField(
            "serviceDiagnostics",
            readComplex(() -> DiagnosticInfo.staticParse(readBuffer), readBuffer));

    int noOfStringTable = readSimpleField("noOfStringTable", readSignedInt(readBuffer, 32));

    List<PascalString> stringTable =
        readCountArrayField(
            "stringTable",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfStringTable);

    ExtensionObject additionalHeader =
        readSimpleField(
            "additionalHeader",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("ResponseHeader");
    // Create the instance
    return new ResponseHeaderBuilderImpl(
        timestamp,
        requestHandle,
        serviceResult,
        serviceDiagnostics,
        noOfStringTable,
        stringTable,
        additionalHeader);
  }

  public static class ResponseHeaderBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long timestamp;
    private final long requestHandle;
    private final StatusCode serviceResult;
    private final DiagnosticInfo serviceDiagnostics;
    private final int noOfStringTable;
    private final List<PascalString> stringTable;
    private final ExtensionObject additionalHeader;

    public ResponseHeaderBuilderImpl(
        long timestamp,
        long requestHandle,
        StatusCode serviceResult,
        DiagnosticInfo serviceDiagnostics,
        int noOfStringTable,
        List<PascalString> stringTable,
        ExtensionObject additionalHeader) {
      this.timestamp = timestamp;
      this.requestHandle = requestHandle;
      this.serviceResult = serviceResult;
      this.serviceDiagnostics = serviceDiagnostics;
      this.noOfStringTable = noOfStringTable;
      this.stringTable = stringTable;
      this.additionalHeader = additionalHeader;
    }

    public ResponseHeader build() {
      ResponseHeader responseHeader =
          new ResponseHeader(
              timestamp,
              requestHandle,
              serviceResult,
              serviceDiagnostics,
              noOfStringTable,
              stringTable,
              additionalHeader);
      return responseHeader;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ResponseHeader)) {
      return false;
    }
    ResponseHeader that = (ResponseHeader) o;
    return (getTimestamp() == that.getTimestamp())
        && (getRequestHandle() == that.getRequestHandle())
        && (getServiceResult() == that.getServiceResult())
        && (getServiceDiagnostics() == that.getServiceDiagnostics())
        && (getNoOfStringTable() == that.getNoOfStringTable())
        && (getStringTable() == that.getStringTable())
        && (getAdditionalHeader() == that.getAdditionalHeader())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTimestamp(),
        getRequestHandle(),
        getServiceResult(),
        getServiceDiagnostics(),
        getNoOfStringTable(),
        getStringTable(),
        getAdditionalHeader());
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
