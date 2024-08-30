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

public class BACnetServiceAckAtomicWriteFile extends BACnetServiceAck implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.ATOMIC_WRITE_FILE;
  }

  // Properties.
  protected final BACnetContextTagSignedInteger fileStartPosition;

  // Arguments.
  protected final Long serviceAckLength;

  public BACnetServiceAckAtomicWriteFile(
      BACnetContextTagSignedInteger fileStartPosition, Long serviceAckLength) {
    super(serviceAckLength);
    this.fileStartPosition = fileStartPosition;
    this.serviceAckLength = serviceAckLength;
  }

  public BACnetContextTagSignedInteger getFileStartPosition() {
    return fileStartPosition;
  }

  @Override
  protected void serializeBACnetServiceAckChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetServiceAckAtomicWriteFile");

    // Simple Field (fileStartPosition)
    writeSimpleField(
        "fileStartPosition", fileStartPosition, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetServiceAckAtomicWriteFile");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetServiceAckAtomicWriteFile _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (fileStartPosition)
    lengthInBits += fileStartPosition.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetServiceAckBuilder staticParseBACnetServiceAckBuilder(
      ReadBuffer readBuffer, Long serviceAckLength) throws ParseException {
    readBuffer.pullContext("BACnetServiceAckAtomicWriteFile");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagSignedInteger fileStartPosition =
        readSimpleField(
            "fileStartPosition",
            readComplex(
                () ->
                    (BACnetContextTagSignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.SIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetServiceAckAtomicWriteFile");
    // Create the instance
    return new BACnetServiceAckAtomicWriteFileBuilderImpl(fileStartPosition, serviceAckLength);
  }

  public static class BACnetServiceAckAtomicWriteFileBuilderImpl
      implements BACnetServiceAck.BACnetServiceAckBuilder {
    private final BACnetContextTagSignedInteger fileStartPosition;
    private final Long serviceAckLength;

    public BACnetServiceAckAtomicWriteFileBuilderImpl(
        BACnetContextTagSignedInteger fileStartPosition, Long serviceAckLength) {
      this.fileStartPosition = fileStartPosition;
      this.serviceAckLength = serviceAckLength;
    }

    public BACnetServiceAckAtomicWriteFile build(Long serviceAckLength) {

      BACnetServiceAckAtomicWriteFile bACnetServiceAckAtomicWriteFile =
          new BACnetServiceAckAtomicWriteFile(fileStartPosition, serviceAckLength);
      return bACnetServiceAckAtomicWriteFile;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetServiceAckAtomicWriteFile)) {
      return false;
    }
    BACnetServiceAckAtomicWriteFile that = (BACnetServiceAckAtomicWriteFile) o;
    return (getFileStartPosition() == that.getFileStartPosition()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFileStartPosition());
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
