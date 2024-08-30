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

public class BACnetEventParameterChangeOfTimer extends BACnetEventParameter implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger timeDelay;
  protected final BACnetEventParameterChangeOfTimerAlarmValue alarmValues;
  protected final BACnetDeviceObjectPropertyReferenceEnclosed updateTimeReference;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterChangeOfTimer(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger timeDelay,
      BACnetEventParameterChangeOfTimerAlarmValue alarmValues,
      BACnetDeviceObjectPropertyReferenceEnclosed updateTimeReference,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.timeDelay = timeDelay;
    this.alarmValues = alarmValues;
    this.updateTimeReference = updateTimeReference;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getTimeDelay() {
    return timeDelay;
  }

  public BACnetEventParameterChangeOfTimerAlarmValue getAlarmValues() {
    return alarmValues;
  }

  public BACnetDeviceObjectPropertyReferenceEnclosed getUpdateTimeReference() {
    return updateTimeReference;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetEventParameterChangeOfTimer");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (alarmValues)
    writeSimpleField("alarmValues", alarmValues, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (updateTimeReference)
    writeSimpleField(
        "updateTimeReference", updateTimeReference, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterChangeOfTimer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterChangeOfTimer _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (alarmValues)
    lengthInBits += alarmValues.getLengthInBits();

    // Simple field (updateTimeReference)
    lengthInBits += updateTimeReference.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterChangeOfTimer");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            readComplex(() -> BACnetOpeningTag.staticParse(readBuffer, (short) (22)), readBuffer));

    BACnetContextTagUnsignedInteger timeDelay =
        readSimpleField(
            "timeDelay",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetEventParameterChangeOfTimerAlarmValue alarmValues =
        readSimpleField(
            "alarmValues",
            readComplex(
                () ->
                    BACnetEventParameterChangeOfTimerAlarmValue.staticParse(
                        readBuffer, (short) (1)),
                readBuffer));

    BACnetDeviceObjectPropertyReferenceEnclosed updateTimeReference =
        readSimpleField(
            "updateTimeReference",
            readComplex(
                () ->
                    BACnetDeviceObjectPropertyReferenceEnclosed.staticParse(
                        readBuffer, (short) (2)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            readComplex(() -> BACnetClosingTag.staticParse(readBuffer, (short) (22)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterChangeOfTimer");
    // Create the instance
    return new BACnetEventParameterChangeOfTimerBuilderImpl(
        openingTag, timeDelay, alarmValues, updateTimeReference, closingTag);
  }

  public static class BACnetEventParameterChangeOfTimerBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetContextTagUnsignedInteger timeDelay;
    private final BACnetEventParameterChangeOfTimerAlarmValue alarmValues;
    private final BACnetDeviceObjectPropertyReferenceEnclosed updateTimeReference;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterChangeOfTimerBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetContextTagUnsignedInteger timeDelay,
        BACnetEventParameterChangeOfTimerAlarmValue alarmValues,
        BACnetDeviceObjectPropertyReferenceEnclosed updateTimeReference,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.timeDelay = timeDelay;
      this.alarmValues = alarmValues;
      this.updateTimeReference = updateTimeReference;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterChangeOfTimer build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterChangeOfTimer bACnetEventParameterChangeOfTimer =
          new BACnetEventParameterChangeOfTimer(
              peekedTagHeader, openingTag, timeDelay, alarmValues, updateTimeReference, closingTag);
      return bACnetEventParameterChangeOfTimer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterChangeOfTimer)) {
      return false;
    }
    BACnetEventParameterChangeOfTimer that = (BACnetEventParameterChangeOfTimer) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getTimeDelay() == that.getTimeDelay())
        && (getAlarmValues() == that.getAlarmValues())
        && (getUpdateTimeReference() == that.getUpdateTimeReference())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getOpeningTag(),
        getTimeDelay(),
        getAlarmValues(),
        getUpdateTimeReference(),
        getClosingTag());
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
