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

public class DeviceConfigurationRequest extends KnxNetIpMessage implements Message {

  // Accessors for discriminator values.
  public Integer getMsgType() {
    return (int) 0x0310;
  }

  // Properties.
  protected final DeviceConfigurationRequestDataBlock deviceConfigurationRequestDataBlock;
  protected final CEMI cemi;

  public DeviceConfigurationRequest(
      DeviceConfigurationRequestDataBlock deviceConfigurationRequestDataBlock, CEMI cemi) {
    super();
    this.deviceConfigurationRequestDataBlock = deviceConfigurationRequestDataBlock;
    this.cemi = cemi;
  }

  public DeviceConfigurationRequestDataBlock getDeviceConfigurationRequestDataBlock() {
    return deviceConfigurationRequestDataBlock;
  }

  public CEMI getCemi() {
    return cemi;
  }

  @Override
  protected void serializeKnxNetIpMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeviceConfigurationRequest");

    // Simple Field (deviceConfigurationRequestDataBlock)
    writeSimpleField(
        "deviceConfigurationRequestDataBlock",
        deviceConfigurationRequestDataBlock,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (cemi)
    writeSimpleField(
        "cemi",
        cemi,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("DeviceConfigurationRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DeviceConfigurationRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (deviceConfigurationRequestDataBlock)
    lengthInBits += deviceConfigurationRequestDataBlock.getLengthInBits();

    // Simple field (cemi)
    lengthInBits += cemi.getLengthInBits();

    return lengthInBits;
  }

  public static KnxNetIpMessageBuilder staticParseKnxNetIpMessageBuilder(
      ReadBuffer readBuffer, Integer totalLength) throws ParseException {
    readBuffer.pullContext("DeviceConfigurationRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DeviceConfigurationRequestDataBlock deviceConfigurationRequestDataBlock =
        readSimpleField(
            "deviceConfigurationRequestDataBlock",
            readComplex(
                () -> DeviceConfigurationRequestDataBlock.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    CEMI cemi =
        readSimpleField(
            "cemi",
            readComplex(
                () ->
                    CEMI.staticParse(
                        readBuffer,
                        (int)
                            ((totalLength)
                                - (((6)
                                    + (deviceConfigurationRequestDataBlock.getLengthInBytes()))))),
                readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("DeviceConfigurationRequest");
    // Create the instance
    return new DeviceConfigurationRequestBuilderImpl(deviceConfigurationRequestDataBlock, cemi);
  }

  public static class DeviceConfigurationRequestBuilderImpl
      implements KnxNetIpMessage.KnxNetIpMessageBuilder {
    private final DeviceConfigurationRequestDataBlock deviceConfigurationRequestDataBlock;
    private final CEMI cemi;

    public DeviceConfigurationRequestBuilderImpl(
        DeviceConfigurationRequestDataBlock deviceConfigurationRequestDataBlock, CEMI cemi) {
      this.deviceConfigurationRequestDataBlock = deviceConfigurationRequestDataBlock;
      this.cemi = cemi;
    }

    public DeviceConfigurationRequest build() {
      DeviceConfigurationRequest deviceConfigurationRequest =
          new DeviceConfigurationRequest(deviceConfigurationRequestDataBlock, cemi);
      return deviceConfigurationRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeviceConfigurationRequest)) {
      return false;
    }
    DeviceConfigurationRequest that = (DeviceConfigurationRequest) o;
    return (getDeviceConfigurationRequestDataBlock()
            == that.getDeviceConfigurationRequestDataBlock())
        && (getCemi() == that.getCemi())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDeviceConfigurationRequestDataBlock(), getCemi());
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
