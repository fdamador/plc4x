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

public class RegisterServer2Request extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "12195";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final ExtensionObjectDefinition server;
  protected final int noOfDiscoveryConfiguration;
  protected final List<ExtensionObject> discoveryConfiguration;

  public RegisterServer2Request(
      ExtensionObjectDefinition requestHeader,
      ExtensionObjectDefinition server,
      int noOfDiscoveryConfiguration,
      List<ExtensionObject> discoveryConfiguration) {
    super();
    this.requestHeader = requestHeader;
    this.server = server;
    this.noOfDiscoveryConfiguration = noOfDiscoveryConfiguration;
    this.discoveryConfiguration = discoveryConfiguration;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public ExtensionObjectDefinition getServer() {
    return server;
  }

  public int getNoOfDiscoveryConfiguration() {
    return noOfDiscoveryConfiguration;
  }

  public List<ExtensionObject> getDiscoveryConfiguration() {
    return discoveryConfiguration;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RegisterServer2Request");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (server)
    writeSimpleField("server", server, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfDiscoveryConfiguration)
    writeSimpleField(
        "noOfDiscoveryConfiguration", noOfDiscoveryConfiguration, writeSignedInt(writeBuffer, 32));

    // Array Field (discoveryConfiguration)
    writeComplexTypeArrayField("discoveryConfiguration", discoveryConfiguration, writeBuffer);

    writeBuffer.popContext("RegisterServer2Request");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RegisterServer2Request _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (server)
    lengthInBits += server.getLengthInBits();

    // Simple field (noOfDiscoveryConfiguration)
    lengthInBits += 32;

    // Array field
    if (discoveryConfiguration != null) {
      int i = 0;
      for (ExtensionObject element : discoveryConfiguration) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= discoveryConfiguration.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RegisterServer2Request");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    ExtensionObjectDefinition server =
        readSimpleField(
            "server",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("434")),
                readBuffer));

    int noOfDiscoveryConfiguration =
        readSimpleField("noOfDiscoveryConfiguration", readSignedInt(readBuffer, 32));

    List<ExtensionObject> discoveryConfiguration =
        readCountArrayField(
            "discoveryConfiguration",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer),
            noOfDiscoveryConfiguration);

    readBuffer.closeContext("RegisterServer2Request");
    // Create the instance
    return new RegisterServer2RequestBuilderImpl(
        requestHeader, server, noOfDiscoveryConfiguration, discoveryConfiguration);
  }

  public static class RegisterServer2RequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final ExtensionObjectDefinition server;
    private final int noOfDiscoveryConfiguration;
    private final List<ExtensionObject> discoveryConfiguration;

    public RegisterServer2RequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        ExtensionObjectDefinition server,
        int noOfDiscoveryConfiguration,
        List<ExtensionObject> discoveryConfiguration) {
      this.requestHeader = requestHeader;
      this.server = server;
      this.noOfDiscoveryConfiguration = noOfDiscoveryConfiguration;
      this.discoveryConfiguration = discoveryConfiguration;
    }

    public RegisterServer2Request build() {
      RegisterServer2Request registerServer2Request =
          new RegisterServer2Request(
              requestHeader, server, noOfDiscoveryConfiguration, discoveryConfiguration);
      return registerServer2Request;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RegisterServer2Request)) {
      return false;
    }
    RegisterServer2Request that = (RegisterServer2Request) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getServer() == that.getServer())
        && (getNoOfDiscoveryConfiguration() == that.getNoOfDiscoveryConfiguration())
        && (getDiscoveryConfiguration() == that.getDiscoveryConfiguration())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getServer(),
        getNoOfDiscoveryConfiguration(),
        getDiscoveryConfiguration());
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
