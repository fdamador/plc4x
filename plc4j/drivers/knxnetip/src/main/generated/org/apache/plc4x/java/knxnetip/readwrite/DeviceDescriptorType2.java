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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class DeviceDescriptorType2 implements Message {

  // Properties.
  protected final int manufacturerId;
  protected final int deviceType;
  protected final short version;
  protected final boolean readSupported;
  protected final boolean writeSupported;
  protected final byte logicalTagBase;
  protected final ChannelInformation channelInfo1;
  protected final ChannelInformation channelInfo2;
  protected final ChannelInformation channelInfo3;
  protected final ChannelInformation channelInfo4;

  public DeviceDescriptorType2(
      int manufacturerId,
      int deviceType,
      short version,
      boolean readSupported,
      boolean writeSupported,
      byte logicalTagBase,
      ChannelInformation channelInfo1,
      ChannelInformation channelInfo2,
      ChannelInformation channelInfo3,
      ChannelInformation channelInfo4) {
    super();
    this.manufacturerId = manufacturerId;
    this.deviceType = deviceType;
    this.version = version;
    this.readSupported = readSupported;
    this.writeSupported = writeSupported;
    this.logicalTagBase = logicalTagBase;
    this.channelInfo1 = channelInfo1;
    this.channelInfo2 = channelInfo2;
    this.channelInfo3 = channelInfo3;
    this.channelInfo4 = channelInfo4;
  }

  public int getManufacturerId() {
    return manufacturerId;
  }

  public int getDeviceType() {
    return deviceType;
  }

  public short getVersion() {
    return version;
  }

  public boolean getReadSupported() {
    return readSupported;
  }

  public boolean getWriteSupported() {
    return writeSupported;
  }

  public byte getLogicalTagBase() {
    return logicalTagBase;
  }

  public ChannelInformation getChannelInfo1() {
    return channelInfo1;
  }

  public ChannelInformation getChannelInfo2() {
    return channelInfo2;
  }

  public ChannelInformation getChannelInfo3() {
    return channelInfo3;
  }

  public ChannelInformation getChannelInfo4() {
    return channelInfo4;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeviceDescriptorType2");

    // Simple Field (manufacturerId)
    writeSimpleField("manufacturerId", manufacturerId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (deviceType)
    writeSimpleField("deviceType", deviceType, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (version)
    writeSimpleField("version", version, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (readSupported)
    writeSimpleField("readSupported", readSupported, writeBoolean(writeBuffer));

    // Simple Field (writeSupported)
    writeSimpleField("writeSupported", writeSupported, writeBoolean(writeBuffer));

    // Simple Field (logicalTagBase)
    writeSimpleField("logicalTagBase", logicalTagBase, writeUnsignedByte(writeBuffer, 6));

    // Simple Field (channelInfo1)
    writeSimpleField("channelInfo1", channelInfo1, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (channelInfo2)
    writeSimpleField("channelInfo2", channelInfo2, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (channelInfo3)
    writeSimpleField("channelInfo3", channelInfo3, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (channelInfo4)
    writeSimpleField("channelInfo4", channelInfo4, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("DeviceDescriptorType2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DeviceDescriptorType2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (manufacturerId)
    lengthInBits += 16;

    // Simple field (deviceType)
    lengthInBits += 16;

    // Simple field (version)
    lengthInBits += 8;

    // Simple field (readSupported)
    lengthInBits += 1;

    // Simple field (writeSupported)
    lengthInBits += 1;

    // Simple field (logicalTagBase)
    lengthInBits += 6;

    // Simple field (channelInfo1)
    lengthInBits += channelInfo1.getLengthInBits();

    // Simple field (channelInfo2)
    lengthInBits += channelInfo2.getLengthInBits();

    // Simple field (channelInfo3)
    lengthInBits += channelInfo3.getLengthInBits();

    // Simple field (channelInfo4)
    lengthInBits += channelInfo4.getLengthInBits();

    return lengthInBits;
  }

  public static DeviceDescriptorType2 staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DeviceDescriptorType2");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int manufacturerId = readSimpleField("manufacturerId", readUnsignedInt(readBuffer, 16));

    int deviceType = readSimpleField("deviceType", readUnsignedInt(readBuffer, 16));

    short version = readSimpleField("version", readUnsignedShort(readBuffer, 8));

    boolean readSupported = readSimpleField("readSupported", readBoolean(readBuffer));

    boolean writeSupported = readSimpleField("writeSupported", readBoolean(readBuffer));

    byte logicalTagBase = readSimpleField("logicalTagBase", readUnsignedByte(readBuffer, 6));

    ChannelInformation channelInfo1 =
        readSimpleField(
            "channelInfo1",
            readComplex(() -> ChannelInformation.staticParse(readBuffer), readBuffer));

    ChannelInformation channelInfo2 =
        readSimpleField(
            "channelInfo2",
            readComplex(() -> ChannelInformation.staticParse(readBuffer), readBuffer));

    ChannelInformation channelInfo3 =
        readSimpleField(
            "channelInfo3",
            readComplex(() -> ChannelInformation.staticParse(readBuffer), readBuffer));

    ChannelInformation channelInfo4 =
        readSimpleField(
            "channelInfo4",
            readComplex(() -> ChannelInformation.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("DeviceDescriptorType2");
    // Create the instance
    DeviceDescriptorType2 _deviceDescriptorType2;
    _deviceDescriptorType2 =
        new DeviceDescriptorType2(
            manufacturerId,
            deviceType,
            version,
            readSupported,
            writeSupported,
            logicalTagBase,
            channelInfo1,
            channelInfo2,
            channelInfo3,
            channelInfo4);
    return _deviceDescriptorType2;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeviceDescriptorType2)) {
      return false;
    }
    DeviceDescriptorType2 that = (DeviceDescriptorType2) o;
    return (getManufacturerId() == that.getManufacturerId())
        && (getDeviceType() == that.getDeviceType())
        && (getVersion() == that.getVersion())
        && (getReadSupported() == that.getReadSupported())
        && (getWriteSupported() == that.getWriteSupported())
        && (getLogicalTagBase() == that.getLogicalTagBase())
        && (getChannelInfo1() == that.getChannelInfo1())
        && (getChannelInfo2() == that.getChannelInfo2())
        && (getChannelInfo3() == that.getChannelInfo3())
        && (getChannelInfo4() == that.getChannelInfo4())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getManufacturerId(),
        getDeviceType(),
        getVersion(),
        getReadSupported(),
        getWriteSupported(),
        getLogicalTagBase(),
        getChannelInfo1(),
        getChannelInfo2(),
        getChannelInfo3(),
        getChannelInfo4());
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
