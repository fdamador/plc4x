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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithoutTime_PARAMETER_ACTIVATION extends InformationObjectWithoutTime
    implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.PARAMETER_ACTIVATION;
  }

  // Properties.
  protected final QualifierOfParameterActivation qpa;

  public InformationObjectWithoutTime_PARAMETER_ACTIVATION(
      int address, QualifierOfParameterActivation qpa) {
    super(address);
    this.qpa = qpa;
  }

  public QualifierOfParameterActivation getQpa() {
    return qpa;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithoutTime_PARAMETER_ACTIVATION");

    // Simple Field (qpa)
    writeSimpleField(
        "qpa",
        qpa,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithoutTime_PARAMETER_ACTIVATION");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_PARAMETER_ACTIVATION _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (qpa)
    lengthInBits += qpa.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithoutTime_PARAMETER_ACTIVATION");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    QualifierOfParameterActivation qpa =
        readSimpleField(
            "qpa",
            readComplex(() -> QualifierOfParameterActivation.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithoutTime_PARAMETER_ACTIVATION");
    // Create the instance
    return new InformationObjectWithoutTime_PARAMETER_ACTIVATIONBuilderImpl(qpa);
  }

  public static class InformationObjectWithoutTime_PARAMETER_ACTIVATIONBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final QualifierOfParameterActivation qpa;

    public InformationObjectWithoutTime_PARAMETER_ACTIVATIONBuilderImpl(
        QualifierOfParameterActivation qpa) {
      this.qpa = qpa;
    }

    public InformationObjectWithoutTime_PARAMETER_ACTIVATION build(int address) {
      InformationObjectWithoutTime_PARAMETER_ACTIVATION
          informationObjectWithoutTime_PARAMETER_ACTIVATION =
              new InformationObjectWithoutTime_PARAMETER_ACTIVATION(address, qpa);
      return informationObjectWithoutTime_PARAMETER_ACTIVATION;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithoutTime_PARAMETER_ACTIVATION)) {
      return false;
    }
    InformationObjectWithoutTime_PARAMETER_ACTIVATION that =
        (InformationObjectWithoutTime_PARAMETER_ACTIVATION) o;
    return (getQpa() == that.getQpa()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getQpa());
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
