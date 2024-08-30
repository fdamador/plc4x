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

public class BACnetConfirmedServiceRequestAtomicReadFileRecord
    extends BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagSignedInteger fileStartRecord;
  protected final BACnetApplicationTagUnsignedInteger requestRecordCount;

  public BACnetConfirmedServiceRequestAtomicReadFileRecord(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetClosingTag closingTag,
      BACnetApplicationTagSignedInteger fileStartRecord,
      BACnetApplicationTagUnsignedInteger requestRecordCount) {
    super(peekedTagHeader, openingTag, closingTag);
    this.fileStartRecord = fileStartRecord;
    this.requestRecordCount = requestRecordCount;
  }

  public BACnetApplicationTagSignedInteger getFileStartRecord() {
    return fileStartRecord;
  }

  public BACnetApplicationTagUnsignedInteger getRequestRecordCount() {
    return requestRecordCount;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestAtomicReadFileRecord");

    // Simple Field (fileStartRecord)
    writeSimpleField(
        "fileStartRecord", fileStartRecord, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (requestRecordCount)
    writeSimpleField(
        "requestRecordCount", requestRecordCount, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestAtomicReadFileRecord");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestAtomicReadFileRecord _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (fileStartRecord)
    lengthInBits += fileStartRecord.getLengthInBits();

    // Simple field (requestRecordCount)
    lengthInBits += requestRecordCount.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
      staticParseBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder(
          ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestAtomicReadFileRecord");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagSignedInteger fileStartRecord =
        readSimpleField(
            "fileStartRecord",
            readComplex(
                () ->
                    (BACnetApplicationTagSignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagUnsignedInteger requestRecordCount =
        readSimpleField(
            "requestRecordCount",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestAtomicReadFileRecord");
    // Create the instance
    return new BACnetConfirmedServiceRequestAtomicReadFileRecordBuilderImpl(
        fileStartRecord, requestRecordCount);
  }

  public static class BACnetConfirmedServiceRequestAtomicReadFileRecordBuilderImpl
      implements BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
          .BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
    private final BACnetApplicationTagSignedInteger fileStartRecord;
    private final BACnetApplicationTagUnsignedInteger requestRecordCount;

    public BACnetConfirmedServiceRequestAtomicReadFileRecordBuilderImpl(
        BACnetApplicationTagSignedInteger fileStartRecord,
        BACnetApplicationTagUnsignedInteger requestRecordCount) {
      this.fileStartRecord = fileStartRecord;
      this.requestRecordCount = requestRecordCount;
    }

    public BACnetConfirmedServiceRequestAtomicReadFileRecord build(
        BACnetTagHeader peekedTagHeader, BACnetOpeningTag openingTag, BACnetClosingTag closingTag) {
      BACnetConfirmedServiceRequestAtomicReadFileRecord
          bACnetConfirmedServiceRequestAtomicReadFileRecord =
              new BACnetConfirmedServiceRequestAtomicReadFileRecord(
                  peekedTagHeader, openingTag, closingTag, fileStartRecord, requestRecordCount);
      return bACnetConfirmedServiceRequestAtomicReadFileRecord;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestAtomicReadFileRecord)) {
      return false;
    }
    BACnetConfirmedServiceRequestAtomicReadFileRecord that =
        (BACnetConfirmedServiceRequestAtomicReadFileRecord) o;
    return (getFileStartRecord() == that.getFileStartRecord())
        && (getRequestRecordCount() == that.getRequestRecordCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFileStartRecord(), getRequestRecordCount());
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
