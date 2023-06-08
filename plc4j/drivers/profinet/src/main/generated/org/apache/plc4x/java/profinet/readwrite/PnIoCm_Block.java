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
package org.apache.plc4x.java.profinet.readwrite;

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

public abstract class PnIoCm_Block implements Message {

  // Abstract accessors for discriminator values.
  public abstract PnIoCm_BlockType getBlockType();

  public PnIoCm_Block() {
    super();
  }

  protected abstract void serializePnIoCm_BlockChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCm_Block");

    // Discriminator Field (blockType) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "blockType",
        "PnIoCm_BlockType",
        getBlockType(),
        new DataWriterEnumDefault<>(
            PnIoCm_BlockType::getValue, PnIoCm_BlockType::name, writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch field (Serialize the sub-type)
    serializePnIoCm_BlockChild(writeBuffer);

    writeBuffer.popContext("PnIoCm_Block");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnIoCm_Block _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (blockType)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static PnIoCm_Block staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnIoCm_Block staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnIoCm_Block");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PnIoCm_BlockType blockType =
        readDiscriminatorField(
            "blockType",
            new DataReaderEnumDefault<>(
                PnIoCm_BlockType::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    PnIoCm_BlockBuilder builder = null;
    if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOD_WRITE_REQ_HEADER)) {
      builder = IODWriteRequestHeader.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOD_WRITE_RES_HEADER)) {
      builder = IODWriteResponseHeader.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOD_READ_REQ_HEADER)) {
      builder = IODReadRequestHeader.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOD_READ_RES_HEADER)) {
      builder = IODReadResponseHeader.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.PD_INTERFACE_ADJUST)) {
      builder = PDInterfaceAdjust.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.PD_PORT_DATA_CHECK)) {
      builder = PDPortDataCheck.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.CHECK_PEERS)) {
      builder = CheckPeers.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_BLOCK_REQ)) {
      builder = PnIoCm_Block_ArReq.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_BLOCK_RES)) {
      builder = PnIoCm_Block_ArRes.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        blockType, PnIoCm_BlockType.IOD_BLOCK_REQ_CONNECTION_APPLICATION_READY)) {
      builder = PnIoCm_Control_Request.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOX_BLOCK_REQ_CONNECTION)) {
      builder = PnIoCM_Block_Request.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOX_BLOCK_RES_CONNECT)) {
      builder = PnIoCM_Block_ResponseConnect.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IOD_CONTROL_RES_CONNECT)) {
      builder = PnIoCm_Control_ResponseConnect.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IO_CR_BLOCK_REQ)) {
      builder = PnIoCm_Block_IoCrReq.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IO_CR_BLOCK_RES)) {
      builder = PnIoCm_Block_IoCrRes.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.ALARM_CR_BLOCK_REQ)) {
      builder = PnIoCm_Block_AlarmCrReq.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.ALARM_CR_BLOCK_RES)) {
      builder = PnIoCm_Block_AlarmCrRes.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.EXPECTED_SUBMODULE_BLOCK_REQ)) {
      builder = PnIoCm_Block_ExpectedSubmoduleReq.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.MODULE_DIFF_BLOCK)) {
      builder = PnIoCm_Block_ModuleDiff.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_SERVER_BLOCK)) {
      builder = PnIoCm_Block_ArServer.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.REAL_IDENTIFICATION_DATA)) {
      builder = PnIoCm_Block_RealIdentificationData.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.I_AND_M_0)) {
      builder = PnIoCm_Block_IAndM0.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.I_AND_M_1)) {
      builder = PnIoCm_Block_IAndM1.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.I_AND_M_2)) {
      builder = PnIoCm_Block_IAndM2.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.I_AND_M_3)) {
      builder = PnIoCm_Block_IAndM3.staticParsePnIoCm_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.I_AND_M_4)) {
      builder = PnIoCm_Block_IAndM4.staticParsePnIoCm_BlockBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "blockType="
              + blockType
              + "]");
    }

    readBuffer.closeContext("PnIoCm_Block");
    // Create the instance
    PnIoCm_Block _pnIoCm_Block = builder.build();
    return _pnIoCm_Block;
  }

  public interface PnIoCm_BlockBuilder {
    PnIoCm_Block build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block)) {
      return false;
    }
    PnIoCm_Block that = (PnIoCm_Block) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
