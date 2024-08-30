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
package org.apache.plc4x.java.ads.discovery.readwrite;

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

public class AdsDiscoveryBlockPassword extends AdsDiscoveryBlock implements Message {

  // Accessors for discriminator values.
  public AdsDiscoveryBlockType getBlockType() {
    return AdsDiscoveryBlockType.PASSWORD;
  }

  // Properties.
  protected final AmsString password;

  public AdsDiscoveryBlockPassword(AmsString password) {
    super();
    this.password = password;
  }

  public AmsString getPassword() {
    return password;
  }

  @Override
  protected void serializeAdsDiscoveryBlockChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsDiscoveryBlockPassword");

    // Simple Field (password)
    writeSimpleField("password", password, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("AdsDiscoveryBlockPassword");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsDiscoveryBlockPassword _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (password)
    lengthInBits += password.getLengthInBits();

    return lengthInBits;
  }

  public static AdsDiscoveryBlockBuilder staticParseAdsDiscoveryBlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsDiscoveryBlockPassword");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    AmsString password =
        readSimpleField(
            "password", readComplex(() -> AmsString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("AdsDiscoveryBlockPassword");
    // Create the instance
    return new AdsDiscoveryBlockPasswordBuilderImpl(password);
  }

  public static class AdsDiscoveryBlockPasswordBuilderImpl
      implements AdsDiscoveryBlock.AdsDiscoveryBlockBuilder {
    private final AmsString password;

    public AdsDiscoveryBlockPasswordBuilderImpl(AmsString password) {
      this.password = password;
    }

    public AdsDiscoveryBlockPassword build() {
      AdsDiscoveryBlockPassword adsDiscoveryBlockPassword = new AdsDiscoveryBlockPassword(password);
      return adsDiscoveryBlockPassword;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsDiscoveryBlockPassword)) {
      return false;
    }
    AdsDiscoveryBlockPassword that = (AdsDiscoveryBlockPassword) o;
    return (getPassword() == that.getPassword()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPassword());
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
