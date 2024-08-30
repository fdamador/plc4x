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

public class BACnetConfirmedServiceRequestDeleteObject extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.DELETE_OBJECT;
  }

  // Properties.
  protected final BACnetApplicationTagObjectIdentifier objectIdentifier;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestDeleteObject(
      BACnetApplicationTagObjectIdentifier objectIdentifier, Long serviceRequestLength) {
    super(serviceRequestLength);
    this.objectIdentifier = objectIdentifier;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetApplicationTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestDeleteObject");

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestDeleteObject");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestDeleteObject _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestDeleteObject");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagObjectIdentifier)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestDeleteObject");
    // Create the instance
    return new BACnetConfirmedServiceRequestDeleteObjectBuilderImpl(
        objectIdentifier, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestDeleteObjectBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetApplicationTagObjectIdentifier objectIdentifier;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestDeleteObjectBuilderImpl(
        BACnetApplicationTagObjectIdentifier objectIdentifier, Long serviceRequestLength) {
      this.objectIdentifier = objectIdentifier;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestDeleteObject build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestDeleteObject bACnetConfirmedServiceRequestDeleteObject =
          new BACnetConfirmedServiceRequestDeleteObject(objectIdentifier, serviceRequestLength);
      return bACnetConfirmedServiceRequestDeleteObject;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestDeleteObject)) {
      return false;
    }
    BACnetConfirmedServiceRequestDeleteObject that = (BACnetConfirmedServiceRequestDeleteObject) o;
    return (getObjectIdentifier() == that.getObjectIdentifier()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getObjectIdentifier());
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
