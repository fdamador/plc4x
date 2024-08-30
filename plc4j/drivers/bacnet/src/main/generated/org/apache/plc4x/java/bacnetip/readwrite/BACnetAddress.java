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

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetAddress implements Message {

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger networkNumber;
  protected final BACnetApplicationTagOctetString macAddress;

  public BACnetAddress(
      BACnetApplicationTagUnsignedInteger networkNumber,
      BACnetApplicationTagOctetString macAddress) {
    super();
    this.networkNumber = networkNumber;
    this.macAddress = macAddress;
  }

  public BACnetApplicationTagUnsignedInteger getNetworkNumber() {
    return networkNumber;
  }

  public BACnetApplicationTagOctetString getMacAddress() {
    return macAddress;
  }

  public BigInteger getZero() {
    Object o = 0L;
    if (o instanceof BigInteger) return (BigInteger) o;
    return BigInteger.valueOf(((Number) o).longValue());
  }

  public boolean getIsLocalNetwork() {
    return (boolean) ((getNetworkNumber().getActualValue()) == (getZero()));
  }

  public boolean getIsBroadcast() {
    return (boolean) ((getMacAddress().getActualLength()) == (0));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetAddress");

    // Simple Field (networkNumber)
    writeSimpleField("networkNumber", networkNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BigInteger zero = getZero();
    writeBuffer.writeVirtual("zero", zero);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isLocalNetwork = getIsLocalNetwork();
    writeBuffer.writeVirtual("isLocalNetwork", isLocalNetwork);

    // Simple Field (macAddress)
    writeSimpleField("macAddress", macAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isBroadcast = getIsBroadcast();
    writeBuffer.writeVirtual("isBroadcast", isBroadcast);

    writeBuffer.popContext("BACnetAddress");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAddress _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (networkNumber)
    lengthInBits += networkNumber.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (macAddress)
    lengthInBits += macAddress.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetAddress staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetAddress");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger networkNumber =
        readSimpleField(
            "networkNumber",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));
    BigInteger zero = readVirtualField("zero", BigInteger.class, 0L);
    boolean isLocalNetwork =
        readVirtualField(
            "isLocalNetwork", boolean.class, (networkNumber.getActualValue()) == (zero));

    BACnetApplicationTagOctetString macAddress =
        readSimpleField(
            "macAddress",
            readComplex(
                () ->
                    (BACnetApplicationTagOctetString) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));
    boolean isBroadcast =
        readVirtualField("isBroadcast", boolean.class, (macAddress.getActualLength()) == (0));

    readBuffer.closeContext("BACnetAddress");
    // Create the instance
    BACnetAddress _bACnetAddress;
    _bACnetAddress = new BACnetAddress(networkNumber, macAddress);
    return _bACnetAddress;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAddress)) {
      return false;
    }
    BACnetAddress that = (BACnetAddress) o;
    return (getNetworkNumber() == that.getNetworkNumber())
        && (getMacAddress() == that.getMacAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNetworkNumber(), getMacAddress());
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
