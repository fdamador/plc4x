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

public class BACnetOpeningTag implements Message {

  // Properties.
  protected final BACnetTagHeader header;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetOpeningTag(BACnetTagHeader header, Short tagNumberArgument) {
    super();
    this.header = header;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetOpeningTag");

    // Simple Field (header)
    writeSimpleField("header", header, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetOpeningTag");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetOpeningTag _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetOpeningTag staticParse(ReadBuffer readBuffer, Short tagNumberArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetOpeningTag");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header", readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getActualTagNumber()) == (tagNumberArgument))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }
    // Validation
    if (!((header.getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS))) {
      throw new ParseValidationException("should be a context tag");
    }
    // Validation
    if (!((header.getLengthValueType()) == (6))) {
      throw new ParseValidationException("opening tag should have a value of 6");
    }

    readBuffer.closeContext("BACnetOpeningTag");
    // Create the instance
    BACnetOpeningTag _bACnetOpeningTag;
    _bACnetOpeningTag = new BACnetOpeningTag(header, tagNumberArgument);
    return _bACnetOpeningTag;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetOpeningTag)) {
      return false;
    }
    BACnetOpeningTag that = (BACnetOpeningTag) o;
    return (getHeader() == that.getHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader());
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
