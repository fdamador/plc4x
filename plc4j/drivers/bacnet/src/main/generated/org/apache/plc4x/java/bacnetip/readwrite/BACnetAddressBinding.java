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

public class BACnetAddressBinding implements Message {

  // Properties.
  protected final BACnetApplicationTagObjectIdentifier deviceIdentifier;
  protected final BACnetAddress deviceAddress;

  public BACnetAddressBinding(
      BACnetApplicationTagObjectIdentifier deviceIdentifier, BACnetAddress deviceAddress) {
    super();
    this.deviceIdentifier = deviceIdentifier;
    this.deviceAddress = deviceAddress;
  }

  public BACnetApplicationTagObjectIdentifier getDeviceIdentifier() {
    return deviceIdentifier;
  }

  public BACnetAddress getDeviceAddress() {
    return deviceAddress;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetAddressBinding");

    // Simple Field (deviceIdentifier)
    writeSimpleField(
        "deviceIdentifier", deviceIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (deviceAddress)
    writeSimpleField("deviceAddress", deviceAddress, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetAddressBinding");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAddressBinding _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (deviceIdentifier)
    lengthInBits += deviceIdentifier.getLengthInBits();

    // Simple field (deviceAddress)
    lengthInBits += deviceAddress.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetAddressBinding staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetAddressBinding");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagObjectIdentifier deviceIdentifier =
        readSimpleField(
            "deviceIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagObjectIdentifier)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetAddress deviceAddress =
        readSimpleField(
            "deviceAddress", readComplex(() -> BACnetAddress.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("BACnetAddressBinding");
    // Create the instance
    BACnetAddressBinding _bACnetAddressBinding;
    _bACnetAddressBinding = new BACnetAddressBinding(deviceIdentifier, deviceAddress);
    return _bACnetAddressBinding;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAddressBinding)) {
      return false;
    }
    BACnetAddressBinding that = (BACnetAddressBinding) o;
    return (getDeviceIdentifier() == that.getDeviceIdentifier())
        && (getDeviceAddress() == that.getDeviceAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getDeviceIdentifier(), getDeviceAddress());
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
