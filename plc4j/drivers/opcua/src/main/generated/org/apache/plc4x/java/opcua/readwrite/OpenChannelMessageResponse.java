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

public class OpenChannelMessageResponse extends OpenChannelMessage implements Message {

  // Accessors for discriminator values.
  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final int secureChannelId;
  protected final PascalString securityPolicyUri;
  protected final PascalByteString senderCertificate;
  protected final PascalByteString receiverCertificateThumbprint;

  public OpenChannelMessageResponse(
      int secureChannelId,
      PascalString securityPolicyUri,
      PascalByteString senderCertificate,
      PascalByteString receiverCertificateThumbprint) {
    super();
    this.secureChannelId = secureChannelId;
    this.securityPolicyUri = securityPolicyUri;
    this.senderCertificate = senderCertificate;
    this.receiverCertificateThumbprint = receiverCertificateThumbprint;
  }

  public int getSecureChannelId() {
    return secureChannelId;
  }

  public PascalString getSecurityPolicyUri() {
    return securityPolicyUri;
  }

  public PascalByteString getSenderCertificate() {
    return senderCertificate;
  }

  public PascalByteString getReceiverCertificateThumbprint() {
    return receiverCertificateThumbprint;
  }

  @Override
  protected void serializeOpenChannelMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenChannelMessageResponse");

    // Simple Field (secureChannelId)
    writeSimpleField("secureChannelId", secureChannelId, writeSignedInt(writeBuffer, 32));

    // Simple Field (securityPolicyUri)
    writeSimpleField(
        "securityPolicyUri", securityPolicyUri, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (senderCertificate)
    writeSimpleField(
        "senderCertificate", senderCertificate, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (receiverCertificateThumbprint)
    writeSimpleField(
        "receiverCertificateThumbprint",
        receiverCertificateThumbprint,
        new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("OpenChannelMessageResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenChannelMessageResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (secureChannelId)
    lengthInBits += 32;

    // Simple field (securityPolicyUri)
    lengthInBits += securityPolicyUri.getLengthInBits();

    // Simple field (senderCertificate)
    lengthInBits += senderCertificate.getLengthInBits();

    // Simple field (receiverCertificateThumbprint)
    lengthInBits += receiverCertificateThumbprint.getLengthInBits();

    return lengthInBits;
  }

  public static OpenChannelMessageBuilder staticParseOpenChannelMessageBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("OpenChannelMessageResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int secureChannelId = readSimpleField("secureChannelId", readSignedInt(readBuffer, 32));

    PascalString securityPolicyUri =
        readSimpleField(
            "securityPolicyUri",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalByteString senderCertificate =
        readSimpleField(
            "senderCertificate",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    PascalByteString receiverCertificateThumbprint =
        readSimpleField(
            "receiverCertificateThumbprint",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("OpenChannelMessageResponse");
    // Create the instance
    return new OpenChannelMessageResponseBuilderImpl(
        secureChannelId, securityPolicyUri, senderCertificate, receiverCertificateThumbprint);
  }

  public static class OpenChannelMessageResponseBuilderImpl
      implements OpenChannelMessage.OpenChannelMessageBuilder {
    private final int secureChannelId;
    private final PascalString securityPolicyUri;
    private final PascalByteString senderCertificate;
    private final PascalByteString receiverCertificateThumbprint;

    public OpenChannelMessageResponseBuilderImpl(
        int secureChannelId,
        PascalString securityPolicyUri,
        PascalByteString senderCertificate,
        PascalByteString receiverCertificateThumbprint) {
      this.secureChannelId = secureChannelId;
      this.securityPolicyUri = securityPolicyUri;
      this.senderCertificate = senderCertificate;
      this.receiverCertificateThumbprint = receiverCertificateThumbprint;
    }

    public OpenChannelMessageResponse build() {
      OpenChannelMessageResponse openChannelMessageResponse =
          new OpenChannelMessageResponse(
              secureChannelId, securityPolicyUri, senderCertificate, receiverCertificateThumbprint);
      return openChannelMessageResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenChannelMessageResponse)) {
      return false;
    }
    OpenChannelMessageResponse that = (OpenChannelMessageResponse) o;
    return (getSecureChannelId() == that.getSecureChannelId())
        && (getSecurityPolicyUri() == that.getSecurityPolicyUri())
        && (getSenderCertificate() == that.getSenderCertificate())
        && (getReceiverCertificateThumbprint() == that.getReceiverCertificateThumbprint())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSecureChannelId(),
        getSecurityPolicyUri(),
        getSenderCertificate(),
        getReceiverCertificateThumbprint());
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
