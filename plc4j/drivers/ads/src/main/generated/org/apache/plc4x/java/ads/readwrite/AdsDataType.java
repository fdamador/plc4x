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
package org.apache.plc4x.java.ads.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import org.apache.plc4x.java.api.types.PlcValueType;

// Code generated by code-generation. DO NOT EDIT.

public enum AdsDataType {
  BOOL((byte) 0x01, (int) 1, PlcValueType.BOOL),
  BIT((byte) 0x02, (int) 1, PlcValueType.BOOL),
  BIT8((byte) 0x03, (int) 1, PlcValueType.BYTE),
  BYTE((byte) 0x04, (int) 1, PlcValueType.BYTE),
  BITARR8((byte) 0x05, (int) 1, PlcValueType.BYTE),
  WORD((byte) 0x06, (int) 2, PlcValueType.WORD),
  BITARR16((byte) 0x07, (int) 2, PlcValueType.WORD),
  DWORD((byte) 0x08, (int) 4, PlcValueType.DWORD),
  BITARR32((byte) 0x09, (int) 4, PlcValueType.DWORD),
  SINT((byte) 0x0A, (int) 1, PlcValueType.SINT),
  INT8((byte) 0x0B, (int) 1, PlcValueType.SINT),
  USINT((byte) 0x0C, (int) 1, PlcValueType.USINT),
  UINT8((byte) 0x0D, (int) 1, PlcValueType.USINT),
  INT((byte) 0x0E, (int) 2, PlcValueType.INT),
  INT16((byte) 0x0F, (int) 2, PlcValueType.INT),
  UINT((byte) 0x10, (int) 2, PlcValueType.UINT),
  UINT16((byte) 0x11, (int) 2, PlcValueType.UINT),
  DINT((byte) 0x12, (int) 4, PlcValueType.DINT),
  INT32((byte) 0x13, (int) 4, PlcValueType.DINT),
  UDINT((byte) 0x14, (int) 4, PlcValueType.UDINT),
  UINT32((byte) 0x15, (int) 4, PlcValueType.UDINT),
  LINT((byte) 0x16, (int) 8, PlcValueType.LINT),
  INT64((byte) 0x17, (int) 8, PlcValueType.LINT),
  ULINT((byte) 0x18, (int) 8, PlcValueType.ULINT),
  UINT64((byte) 0x19, (int) 8, PlcValueType.ULINT),
  REAL((byte) 0x1A, (int) 4, PlcValueType.REAL),
  FLOAT((byte) 0x1B, (int) 4, PlcValueType.REAL),
  LREAL((byte) 0x1C, (int) 8, PlcValueType.LREAL),
  DOUBLE((byte) 0x1D, (int) 8, PlcValueType.LREAL),
  CHAR((byte) 0x1E, (int) 1, PlcValueType.CHAR),
  WCHAR((byte) 0x1F, (int) 2, PlcValueType.WCHAR),
  STRING((byte) 0x20, (int) 256, PlcValueType.STRING),
  WSTRING((byte) 0x21, (int) 512, PlcValueType.WSTRING),
  TIME((byte) 0x22, (int) 4, PlcValueType.TIME),
  LTIME((byte) 0x23, (int) 8, PlcValueType.LTIME),
  DATE((byte) 0x24, (int) 4, PlcValueType.DATE),
  TIME_OF_DAY((byte) 0x25, (int) 4, PlcValueType.TIME_OF_DAY),
  TOD((byte) 0x26, (int) 4, PlcValueType.TIME_OF_DAY),
  DATE_AND_TIME((byte) 0x27, (int) 4, PlcValueType.DATE_AND_TIME),
  DT((byte) 0x28, (int) 4, PlcValueType.DATE_AND_TIME);
  private static final Map<Byte, AdsDataType> map;

  static {
    map = new HashMap<>();
    for (AdsDataType value : AdsDataType.values()) {
      map.put((byte) value.getValue(), value);
    }
  }

  private final byte value;
  private final int numBytes;
  private final PlcValueType plcValueType;

  AdsDataType(byte value, int numBytes, PlcValueType plcValueType) {
    this.value = value;
    this.numBytes = numBytes;
    this.plcValueType = plcValueType;
  }

  public byte getValue() {
    return value;
  }

  public int getNumBytes() {
    return numBytes;
  }

  public static AdsDataType firstEnumForFieldNumBytes(int fieldValue) {
    for (AdsDataType _val : AdsDataType.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AdsDataType> enumsForFieldNumBytes(int fieldValue) {
    List<AdsDataType> _values = new ArrayList<>();
    for (AdsDataType _val : AdsDataType.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public PlcValueType getPlcValueType() {
    return plcValueType;
  }

  public static AdsDataType firstEnumForFieldPlcValueType(PlcValueType fieldValue) {
    for (AdsDataType _val : AdsDataType.values()) {
      if (_val.getPlcValueType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AdsDataType> enumsForFieldPlcValueType(PlcValueType fieldValue) {
    List<AdsDataType> _values = new ArrayList<>();
    for (AdsDataType _val : AdsDataType.values()) {
      if (_val.getPlcValueType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static AdsDataType enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}
