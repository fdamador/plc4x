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

public class BACnetAccessRule implements Message {

  // Properties.
  protected final BACnetAccessRuleTimeRangeSpecifierTagged timeRangeSpecifier;
  protected final BACnetDeviceObjectPropertyReferenceEnclosed timeRange;
  protected final BACnetAccessRuleLocationSpecifierTagged locationSpecifier;
  protected final BACnetDeviceObjectReferenceEnclosed location;
  protected final BACnetContextTagBoolean enable;

  public BACnetAccessRule(
      BACnetAccessRuleTimeRangeSpecifierTagged timeRangeSpecifier,
      BACnetDeviceObjectPropertyReferenceEnclosed timeRange,
      BACnetAccessRuleLocationSpecifierTagged locationSpecifier,
      BACnetDeviceObjectReferenceEnclosed location,
      BACnetContextTagBoolean enable) {
    super();
    this.timeRangeSpecifier = timeRangeSpecifier;
    this.timeRange = timeRange;
    this.locationSpecifier = locationSpecifier;
    this.location = location;
    this.enable = enable;
  }

  public BACnetAccessRuleTimeRangeSpecifierTagged getTimeRangeSpecifier() {
    return timeRangeSpecifier;
  }

  public BACnetDeviceObjectPropertyReferenceEnclosed getTimeRange() {
    return timeRange;
  }

  public BACnetAccessRuleLocationSpecifierTagged getLocationSpecifier() {
    return locationSpecifier;
  }

  public BACnetDeviceObjectReferenceEnclosed getLocation() {
    return location;
  }

  public BACnetContextTagBoolean getEnable() {
    return enable;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetAccessRule");

    // Simple Field (timeRangeSpecifier)
    writeSimpleField("timeRangeSpecifier", timeRangeSpecifier, writeComplex(writeBuffer));

    // Optional Field (timeRange) (Can be skipped, if the value is null)
    writeOptionalField(
        "timeRange", timeRange, writeComplex(writeBuffer), (getTimeRangeSpecifier()) != (null));

    // Simple Field (locationSpecifier)
    writeSimpleField("locationSpecifier", locationSpecifier, writeComplex(writeBuffer));

    // Optional Field (location) (Can be skipped, if the value is null)
    writeOptionalField(
        "location", location, writeComplex(writeBuffer), (getLocationSpecifier()) != (null));

    // Simple Field (enable)
    writeSimpleField("enable", enable, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetAccessRule");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAccessRule _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (timeRangeSpecifier)
    lengthInBits += timeRangeSpecifier.getLengthInBits();

    // Optional Field (timeRange)
    if (timeRange != null) {
      lengthInBits += timeRange.getLengthInBits();
    }

    // Simple field (locationSpecifier)
    lengthInBits += locationSpecifier.getLengthInBits();

    // Optional Field (location)
    if (location != null) {
      lengthInBits += location.getLengthInBits();
    }

    // Simple field (enable)
    lengthInBits += enable.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetAccessRule staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetAccessRule");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetAccessRuleTimeRangeSpecifierTagged timeRangeSpecifier =
        readSimpleField(
            "timeRangeSpecifier",
            readComplex(
                () ->
                    BACnetAccessRuleTimeRangeSpecifierTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetDeviceObjectPropertyReferenceEnclosed timeRange =
        readOptionalField(
            "timeRange",
            readComplex(
                () ->
                    BACnetDeviceObjectPropertyReferenceEnclosed.staticParse(
                        readBuffer, (short) (1)),
                readBuffer),
            (timeRangeSpecifier) != (null));

    BACnetAccessRuleLocationSpecifierTagged locationSpecifier =
        readSimpleField(
            "locationSpecifier",
            readComplex(
                () ->
                    BACnetAccessRuleLocationSpecifierTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetDeviceObjectReferenceEnclosed location =
        readOptionalField(
            "location",
            readComplex(
                () -> BACnetDeviceObjectReferenceEnclosed.staticParse(readBuffer, (short) (3)),
                readBuffer),
            (locationSpecifier) != (null));

    BACnetContextTagBoolean enable =
        readSimpleField(
            "enable",
            readComplex(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (4), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    readBuffer.closeContext("BACnetAccessRule");
    // Create the instance
    BACnetAccessRule _bACnetAccessRule;
    _bACnetAccessRule =
        new BACnetAccessRule(timeRangeSpecifier, timeRange, locationSpecifier, location, enable);
    return _bACnetAccessRule;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAccessRule)) {
      return false;
    }
    BACnetAccessRule that = (BACnetAccessRule) o;
    return (getTimeRangeSpecifier() == that.getTimeRangeSpecifier())
        && (getTimeRange() == that.getTimeRange())
        && (getLocationSpecifier() == that.getLocationSpecifier())
        && (getLocation() == that.getLocation())
        && (getEnable() == that.getEnable())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getTimeRangeSpecifier(),
        getTimeRange(),
        getLocationSpecifier(),
        getLocation(),
        getEnable());
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
