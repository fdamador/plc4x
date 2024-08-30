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

public class BACnetDeviceObjectReference implements Message {

  // Properties.
  protected final BACnetContextTagObjectIdentifier deviceIdentifier;
  protected final BACnetContextTagObjectIdentifier objectIdentifier;

  public BACnetDeviceObjectReference(
      BACnetContextTagObjectIdentifier deviceIdentifier,
      BACnetContextTagObjectIdentifier objectIdentifier) {
    super();
    this.deviceIdentifier = deviceIdentifier;
    this.objectIdentifier = objectIdentifier;
  }

  public BACnetContextTagObjectIdentifier getDeviceIdentifier() {
    return deviceIdentifier;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetDeviceObjectReference");

    // Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "deviceIdentifier", deviceIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetDeviceObjectReference");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetDeviceObjectReference _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (deviceIdentifier)
    if (deviceIdentifier != null) {
      lengthInBits += deviceIdentifier.getLengthInBits();
    }

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetDeviceObjectReference staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetDeviceObjectReference");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier deviceIdentifier =
        readOptionalField(
            "deviceIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetContextTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    readBuffer.closeContext("BACnetDeviceObjectReference");
    // Create the instance
    BACnetDeviceObjectReference _bACnetDeviceObjectReference;
    _bACnetDeviceObjectReference =
        new BACnetDeviceObjectReference(deviceIdentifier, objectIdentifier);
    return _bACnetDeviceObjectReference;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetDeviceObjectReference)) {
      return false;
    }
    BACnetDeviceObjectReference that = (BACnetDeviceObjectReference) o;
    return (getDeviceIdentifier() == that.getDeviceIdentifier())
        && (getObjectIdentifier() == that.getObjectIdentifier())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getDeviceIdentifier(), getObjectIdentifier());
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
