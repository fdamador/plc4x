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

public class RegisteredServer extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "434";
  }

  // Properties.
  protected final PascalString serverUri;
  protected final PascalString productUri;
  protected final int noOfServerNames;
  protected final List<LocalizedText> serverNames;
  protected final ApplicationType serverType;
  protected final PascalString gatewayServerUri;
  protected final int noOfDiscoveryUrls;
  protected final List<PascalString> discoveryUrls;
  protected final PascalString semaphoreFilePath;
  protected final boolean isOnline;

  public RegisteredServer(
      PascalString serverUri,
      PascalString productUri,
      int noOfServerNames,
      List<LocalizedText> serverNames,
      ApplicationType serverType,
      PascalString gatewayServerUri,
      int noOfDiscoveryUrls,
      List<PascalString> discoveryUrls,
      PascalString semaphoreFilePath,
      boolean isOnline) {
    super();
    this.serverUri = serverUri;
    this.productUri = productUri;
    this.noOfServerNames = noOfServerNames;
    this.serverNames = serverNames;
    this.serverType = serverType;
    this.gatewayServerUri = gatewayServerUri;
    this.noOfDiscoveryUrls = noOfDiscoveryUrls;
    this.discoveryUrls = discoveryUrls;
    this.semaphoreFilePath = semaphoreFilePath;
    this.isOnline = isOnline;
  }

  public PascalString getServerUri() {
    return serverUri;
  }

  public PascalString getProductUri() {
    return productUri;
  }

  public int getNoOfServerNames() {
    return noOfServerNames;
  }

  public List<LocalizedText> getServerNames() {
    return serverNames;
  }

  public ApplicationType getServerType() {
    return serverType;
  }

  public PascalString getGatewayServerUri() {
    return gatewayServerUri;
  }

  public int getNoOfDiscoveryUrls() {
    return noOfDiscoveryUrls;
  }

  public List<PascalString> getDiscoveryUrls() {
    return discoveryUrls;
  }

  public PascalString getSemaphoreFilePath() {
    return semaphoreFilePath;
  }

  public boolean getIsOnline() {
    return isOnline;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RegisteredServer");

    // Simple Field (serverUri)
    writeSimpleField("serverUri", serverUri, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (productUri)
    writeSimpleField("productUri", productUri, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfServerNames)
    writeSimpleField("noOfServerNames", noOfServerNames, writeSignedInt(writeBuffer, 32));

    // Array Field (serverNames)
    writeComplexTypeArrayField("serverNames", serverNames, writeBuffer);

    // Simple Field (serverType)
    writeSimpleEnumField(
        "serverType",
        "ApplicationType",
        serverType,
        new DataWriterEnumDefault<>(
            ApplicationType::getValue, ApplicationType::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (gatewayServerUri)
    writeSimpleField(
        "gatewayServerUri", gatewayServerUri, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfDiscoveryUrls)
    writeSimpleField("noOfDiscoveryUrls", noOfDiscoveryUrls, writeSignedInt(writeBuffer, 32));

    // Array Field (discoveryUrls)
    writeComplexTypeArrayField("discoveryUrls", discoveryUrls, writeBuffer);

    // Simple Field (semaphoreFilePath)
    writeSimpleField(
        "semaphoreFilePath", semaphoreFilePath, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (isOnline)
    writeSimpleField("isOnline", isOnline, writeBoolean(writeBuffer));

    writeBuffer.popContext("RegisteredServer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RegisteredServer _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (serverUri)
    lengthInBits += serverUri.getLengthInBits();

    // Simple field (productUri)
    lengthInBits += productUri.getLengthInBits();

    // Simple field (noOfServerNames)
    lengthInBits += 32;

    // Array field
    if (serverNames != null) {
      int i = 0;
      for (LocalizedText element : serverNames) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverNames.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (serverType)
    lengthInBits += 32;

    // Simple field (gatewayServerUri)
    lengthInBits += gatewayServerUri.getLengthInBits();

    // Simple field (noOfDiscoveryUrls)
    lengthInBits += 32;

    // Array field
    if (discoveryUrls != null) {
      int i = 0;
      for (PascalString element : discoveryUrls) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= discoveryUrls.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (semaphoreFilePath)
    lengthInBits += semaphoreFilePath.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (isOnline)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RegisteredServer");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString serverUri =
        readSimpleField(
            "serverUri", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString productUri =
        readSimpleField(
            "productUri", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfServerNames = readSimpleField("noOfServerNames", readSignedInt(readBuffer, 32));

    List<LocalizedText> serverNames =
        readCountArrayField(
            "serverNames",
            readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer),
            noOfServerNames);

    ApplicationType serverType =
        readEnumField(
            "serverType",
            "ApplicationType",
            new DataReaderEnumDefault<>(
                ApplicationType::enumForValue, readUnsignedLong(readBuffer, 32)));

    PascalString gatewayServerUri =
        readSimpleField(
            "gatewayServerUri",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfDiscoveryUrls = readSimpleField("noOfDiscoveryUrls", readSignedInt(readBuffer, 32));

    List<PascalString> discoveryUrls =
        readCountArrayField(
            "discoveryUrls",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfDiscoveryUrls);

    PascalString semaphoreFilePath =
        readSimpleField(
            "semaphoreFilePath",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean isOnline = readSimpleField("isOnline", readBoolean(readBuffer));

    readBuffer.closeContext("RegisteredServer");
    // Create the instance
    return new RegisteredServerBuilderImpl(
        serverUri,
        productUri,
        noOfServerNames,
        serverNames,
        serverType,
        gatewayServerUri,
        noOfDiscoveryUrls,
        discoveryUrls,
        semaphoreFilePath,
        isOnline);
  }

  public static class RegisteredServerBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString serverUri;
    private final PascalString productUri;
    private final int noOfServerNames;
    private final List<LocalizedText> serverNames;
    private final ApplicationType serverType;
    private final PascalString gatewayServerUri;
    private final int noOfDiscoveryUrls;
    private final List<PascalString> discoveryUrls;
    private final PascalString semaphoreFilePath;
    private final boolean isOnline;

    public RegisteredServerBuilderImpl(
        PascalString serverUri,
        PascalString productUri,
        int noOfServerNames,
        List<LocalizedText> serverNames,
        ApplicationType serverType,
        PascalString gatewayServerUri,
        int noOfDiscoveryUrls,
        List<PascalString> discoveryUrls,
        PascalString semaphoreFilePath,
        boolean isOnline) {
      this.serverUri = serverUri;
      this.productUri = productUri;
      this.noOfServerNames = noOfServerNames;
      this.serverNames = serverNames;
      this.serverType = serverType;
      this.gatewayServerUri = gatewayServerUri;
      this.noOfDiscoveryUrls = noOfDiscoveryUrls;
      this.discoveryUrls = discoveryUrls;
      this.semaphoreFilePath = semaphoreFilePath;
      this.isOnline = isOnline;
    }

    public RegisteredServer build() {
      RegisteredServer registeredServer =
          new RegisteredServer(
              serverUri,
              productUri,
              noOfServerNames,
              serverNames,
              serverType,
              gatewayServerUri,
              noOfDiscoveryUrls,
              discoveryUrls,
              semaphoreFilePath,
              isOnline);
      return registeredServer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RegisteredServer)) {
      return false;
    }
    RegisteredServer that = (RegisteredServer) o;
    return (getServerUri() == that.getServerUri())
        && (getProductUri() == that.getProductUri())
        && (getNoOfServerNames() == that.getNoOfServerNames())
        && (getServerNames() == that.getServerNames())
        && (getServerType() == that.getServerType())
        && (getGatewayServerUri() == that.getGatewayServerUri())
        && (getNoOfDiscoveryUrls() == that.getNoOfDiscoveryUrls())
        && (getDiscoveryUrls() == that.getDiscoveryUrls())
        && (getSemaphoreFilePath() == that.getSemaphoreFilePath())
        && (getIsOnline() == that.getIsOnline())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getServerUri(),
        getProductUri(),
        getNoOfServerNames(),
        getServerNames(),
        getServerType(),
        getGatewayServerUri(),
        getNoOfDiscoveryUrls(),
        getDiscoveryUrls(),
        getSemaphoreFilePath(),
        getIsOnline());
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
