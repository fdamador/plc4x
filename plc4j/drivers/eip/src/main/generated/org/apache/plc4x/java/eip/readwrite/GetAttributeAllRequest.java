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
package org.apache.plc4x.java.eip.readwrite;

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

public class GetAttributeAllRequest extends CipService implements Message {

  // Accessors for discriminator values.
  public Byte getService() {
    return (byte) 0x01;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final PathSegment classSegment;
  protected final PathSegment instanceSegment;

  public GetAttributeAllRequest(PathSegment classSegment, PathSegment instanceSegment) {
    super();
    this.classSegment = classSegment;
    this.instanceSegment = instanceSegment;
  }

  public PathSegment getClassSegment() {
    return classSegment;
  }

  public PathSegment getInstanceSegment() {
    return instanceSegment;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("GetAttributeAllRequest");

    // Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short requestPathSize =
        (short)
            ((((getClassSegment().getLengthInBytes()) + (getInstanceSegment().getLengthInBytes())))
                / (2));
    writeImplicitField("requestPathSize", requestPathSize, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (classSegment)
    writeSimpleField("classSegment", classSegment, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (instanceSegment)
    writeSimpleField(
        "instanceSegment", instanceSegment, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("GetAttributeAllRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    GetAttributeAllRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (requestPathSize)
    lengthInBits += 8;

    // Simple field (classSegment)
    lengthInBits += classSegment.getLengthInBits();

    // Simple field (instanceSegment)
    lengthInBits += instanceSegment.getLengthInBits();

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("GetAttributeAllRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short requestPathSize = readImplicitField("requestPathSize", readUnsignedShort(readBuffer, 8));

    PathSegment classSegment =
        readSimpleField(
            "classSegment", readComplex(() -> PathSegment.staticParse(readBuffer), readBuffer));

    PathSegment instanceSegment =
        readSimpleField(
            "instanceSegment", readComplex(() -> PathSegment.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("GetAttributeAllRequest");
    // Create the instance
    return new GetAttributeAllRequestBuilderImpl(classSegment, instanceSegment);
  }

  public static class GetAttributeAllRequestBuilderImpl implements CipService.CipServiceBuilder {
    private final PathSegment classSegment;
    private final PathSegment instanceSegment;

    public GetAttributeAllRequestBuilderImpl(
        PathSegment classSegment, PathSegment instanceSegment) {
      this.classSegment = classSegment;
      this.instanceSegment = instanceSegment;
    }

    public GetAttributeAllRequest build() {
      GetAttributeAllRequest getAttributeAllRequest =
          new GetAttributeAllRequest(classSegment, instanceSegment);
      return getAttributeAllRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof GetAttributeAllRequest)) {
      return false;
    }
    GetAttributeAllRequest that = (GetAttributeAllRequest) o;
    return (getClassSegment() == that.getClassSegment())
        && (getInstanceSegment() == that.getInstanceSegment())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getClassSegment(), getInstanceSegment());
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
