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

public class BACnetSecurityKeySet implements Message {

  // Properties.
  protected final BACnetContextTagUnsignedInteger keyRevision;
  protected final BACnetDateTimeEnclosed activationTime;
  protected final BACnetDateTimeEnclosed expirationTime;
  protected final BACnetSecurityKeySetKeyIds keyIds;

  public BACnetSecurityKeySet(
      BACnetContextTagUnsignedInteger keyRevision,
      BACnetDateTimeEnclosed activationTime,
      BACnetDateTimeEnclosed expirationTime,
      BACnetSecurityKeySetKeyIds keyIds) {
    super();
    this.keyRevision = keyRevision;
    this.activationTime = activationTime;
    this.expirationTime = expirationTime;
    this.keyIds = keyIds;
  }

  public BACnetContextTagUnsignedInteger getKeyRevision() {
    return keyRevision;
  }

  public BACnetDateTimeEnclosed getActivationTime() {
    return activationTime;
  }

  public BACnetDateTimeEnclosed getExpirationTime() {
    return expirationTime;
  }

  public BACnetSecurityKeySetKeyIds getKeyIds() {
    return keyIds;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetSecurityKeySet");

    // Simple Field (keyRevision)
    writeSimpleField("keyRevision", keyRevision, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (activationTime)
    writeSimpleField("activationTime", activationTime, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (expirationTime)
    writeSimpleField("expirationTime", expirationTime, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (keyIds)
    writeSimpleField("keyIds", keyIds, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetSecurityKeySet");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetSecurityKeySet _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (keyRevision)
    lengthInBits += keyRevision.getLengthInBits();

    // Simple field (activationTime)
    lengthInBits += activationTime.getLengthInBits();

    // Simple field (expirationTime)
    lengthInBits += expirationTime.getLengthInBits();

    // Simple field (keyIds)
    lengthInBits += keyIds.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetSecurityKeySet staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetSecurityKeySet");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger keyRevision =
        readSimpleField(
            "keyRevision",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetDateTimeEnclosed activationTime =
        readSimpleField(
            "activationTime",
            readComplex(
                () -> BACnetDateTimeEnclosed.staticParse(readBuffer, (short) (1)), readBuffer));

    BACnetDateTimeEnclosed expirationTime =
        readSimpleField(
            "expirationTime",
            readComplex(
                () -> BACnetDateTimeEnclosed.staticParse(readBuffer, (short) (2)), readBuffer));

    BACnetSecurityKeySetKeyIds keyIds =
        readSimpleField(
            "keyIds",
            readComplex(
                () -> BACnetSecurityKeySetKeyIds.staticParse(readBuffer, (short) (3)), readBuffer));

    readBuffer.closeContext("BACnetSecurityKeySet");
    // Create the instance
    BACnetSecurityKeySet _bACnetSecurityKeySet;
    _bACnetSecurityKeySet =
        new BACnetSecurityKeySet(keyRevision, activationTime, expirationTime, keyIds);
    return _bACnetSecurityKeySet;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetSecurityKeySet)) {
      return false;
    }
    BACnetSecurityKeySet that = (BACnetSecurityKeySet) o;
    return (getKeyRevision() == that.getKeyRevision())
        && (getActivationTime() == that.getActivationTime())
        && (getExpirationTime() == that.getExpirationTime())
        && (getKeyIds() == that.getKeyIds())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getKeyRevision(), getActivationTime(), getExpirationTime(), getKeyIds());
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
