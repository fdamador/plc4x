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

public class HistoryUpdateRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "700";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfHistoryUpdateDetails;
  protected final List<ExtensionObject> historyUpdateDetails;

  public HistoryUpdateRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfHistoryUpdateDetails,
      List<ExtensionObject> historyUpdateDetails) {
    super();
    this.requestHeader = requestHeader;
    this.noOfHistoryUpdateDetails = noOfHistoryUpdateDetails;
    this.historyUpdateDetails = historyUpdateDetails;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfHistoryUpdateDetails() {
    return noOfHistoryUpdateDetails;
  }

  public List<ExtensionObject> getHistoryUpdateDetails() {
    return historyUpdateDetails;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("HistoryUpdateRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfHistoryUpdateDetails)
    writeSimpleField(
        "noOfHistoryUpdateDetails", noOfHistoryUpdateDetails, writeSignedInt(writeBuffer, 32));

    // Array Field (historyUpdateDetails)
    writeComplexTypeArrayField("historyUpdateDetails", historyUpdateDetails, writeBuffer);

    writeBuffer.popContext("HistoryUpdateRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    HistoryUpdateRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfHistoryUpdateDetails)
    lengthInBits += 32;

    // Array field
    if (historyUpdateDetails != null) {
      int i = 0;
      for (ExtensionObject element : historyUpdateDetails) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= historyUpdateDetails.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("HistoryUpdateRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfHistoryUpdateDetails =
        readSimpleField("noOfHistoryUpdateDetails", readSignedInt(readBuffer, 32));

    List<ExtensionObject> historyUpdateDetails =
        readCountArrayField(
            "historyUpdateDetails",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer),
            noOfHistoryUpdateDetails);

    readBuffer.closeContext("HistoryUpdateRequest");
    // Create the instance
    return new HistoryUpdateRequestBuilderImpl(
        requestHeader, noOfHistoryUpdateDetails, historyUpdateDetails);
  }

  public static class HistoryUpdateRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfHistoryUpdateDetails;
    private final List<ExtensionObject> historyUpdateDetails;

    public HistoryUpdateRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        int noOfHistoryUpdateDetails,
        List<ExtensionObject> historyUpdateDetails) {
      this.requestHeader = requestHeader;
      this.noOfHistoryUpdateDetails = noOfHistoryUpdateDetails;
      this.historyUpdateDetails = historyUpdateDetails;
    }

    public HistoryUpdateRequest build() {
      HistoryUpdateRequest historyUpdateRequest =
          new HistoryUpdateRequest(requestHeader, noOfHistoryUpdateDetails, historyUpdateDetails);
      return historyUpdateRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HistoryUpdateRequest)) {
      return false;
    }
    HistoryUpdateRequest that = (HistoryUpdateRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfHistoryUpdateDetails() == that.getNoOfHistoryUpdateDetails())
        && (getHistoryUpdateDetails() == that.getHistoryUpdateDetails())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getNoOfHistoryUpdateDetails(),
        getHistoryUpdateDetails());
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
