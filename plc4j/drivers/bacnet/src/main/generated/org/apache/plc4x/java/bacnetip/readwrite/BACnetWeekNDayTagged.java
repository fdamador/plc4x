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

public class BACnetWeekNDayTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final short month;
  protected final short weekOfMonth;
  protected final short dayOfWeek;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetWeekNDayTagged(
      BACnetTagHeader header,
      short month,
      short weekOfMonth,
      short dayOfWeek,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.month = month;
    this.weekOfMonth = weekOfMonth;
    this.dayOfWeek = dayOfWeek;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public short getMonth() {
    return month;
  }

  public short getWeekOfMonth() {
    return weekOfMonth;
  }

  public short getDayOfWeek() {
    return dayOfWeek;
  }

  public boolean getOddMonths() {
    return (boolean) ((getMonth()) == (13));
  }

  public boolean getEvenMonths() {
    return (boolean) ((getMonth()) == (14));
  }

  public boolean getAnyMonth() {
    return (boolean) ((getMonth()) == (0xFF));
  }

  public boolean getDays1to7() {
    return (boolean) ((getWeekOfMonth()) == (1));
  }

  public boolean getDays8to14() {
    return (boolean) ((getWeekOfMonth()) == (2));
  }

  public boolean getDays15to21() {
    return (boolean) ((getWeekOfMonth()) == (3));
  }

  public boolean getDays22to28() {
    return (boolean) ((getWeekOfMonth()) == (4));
  }

  public boolean getDays29to31() {
    return (boolean) ((getWeekOfMonth()) == (5));
  }

  public boolean getLast7DaysOfThisMonth() {
    return (boolean) ((getWeekOfMonth()) == (6));
  }

  public boolean getAny7DaysPriorToLast7DaysOfThisMonth() {
    return (boolean) ((getWeekOfMonth()) == (7));
  }

  public boolean getAny7DaysPriorToLast14DaysOfThisMonth() {
    return (boolean) ((getWeekOfMonth()) == (8));
  }

  public boolean getAny7DaysPriorToLast21DaysOfThisMonth() {
    return (boolean) ((getWeekOfMonth()) == (9));
  }

  public boolean getAnyWeekOfthisMonth() {
    return (boolean) ((getWeekOfMonth()) == (0xFF));
  }

  public boolean getAnyDayOfWeek() {
    return (boolean) ((getDayOfWeek()) == (0xFF));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetWeekNDayTagged");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (month)
    writeSimpleField("month", month, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean oddMonths = getOddMonths();
    writeBuffer.writeVirtual("oddMonths", oddMonths);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean evenMonths = getEvenMonths();
    writeBuffer.writeVirtual("evenMonths", evenMonths);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean anyMonth = getAnyMonth();
    writeBuffer.writeVirtual("anyMonth", anyMonth);

    // Simple Field (weekOfMonth)
    writeSimpleField("weekOfMonth", weekOfMonth, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean days1to7 = getDays1to7();
    writeBuffer.writeVirtual("days1to7", days1to7);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean days8to14 = getDays8to14();
    writeBuffer.writeVirtual("days8to14", days8to14);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean days15to21 = getDays15to21();
    writeBuffer.writeVirtual("days15to21", days15to21);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean days22to28 = getDays22to28();
    writeBuffer.writeVirtual("days22to28", days22to28);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean days29to31 = getDays29to31();
    writeBuffer.writeVirtual("days29to31", days29to31);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean last7DaysOfThisMonth = getLast7DaysOfThisMonth();
    writeBuffer.writeVirtual("last7DaysOfThisMonth", last7DaysOfThisMonth);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean any7DaysPriorToLast7DaysOfThisMonth = getAny7DaysPriorToLast7DaysOfThisMonth();
    writeBuffer.writeVirtual(
        "any7DaysPriorToLast7DaysOfThisMonth", any7DaysPriorToLast7DaysOfThisMonth);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean any7DaysPriorToLast14DaysOfThisMonth = getAny7DaysPriorToLast14DaysOfThisMonth();
    writeBuffer.writeVirtual(
        "any7DaysPriorToLast14DaysOfThisMonth", any7DaysPriorToLast14DaysOfThisMonth);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean any7DaysPriorToLast21DaysOfThisMonth = getAny7DaysPriorToLast21DaysOfThisMonth();
    writeBuffer.writeVirtual(
        "any7DaysPriorToLast21DaysOfThisMonth", any7DaysPriorToLast21DaysOfThisMonth);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean anyWeekOfthisMonth = getAnyWeekOfthisMonth();
    writeBuffer.writeVirtual("anyWeekOfthisMonth", anyWeekOfthisMonth);

    // Simple Field (dayOfWeek)
    writeSimpleField("dayOfWeek", dayOfWeek, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean anyDayOfWeek = getAnyDayOfWeek();
    writeBuffer.writeVirtual("anyDayOfWeek", anyDayOfWeek);

    writeBuffer.popContext("BACnetWeekNDayTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetWeekNDayTagged _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Simple field (month)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (weekOfMonth)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (dayOfWeek)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetWeekNDayTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetWeekNDayTagged");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header", readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }
    // Validation
    if (!((header.getActualLength()) == (3))) {
      throw new ParseValidationException("We should have at least 3 octets");
    }

    short month = readSimpleField("month", readUnsignedShort(readBuffer, 8));
    boolean oddMonths = readVirtualField("oddMonths", boolean.class, (month) == (13));
    boolean evenMonths = readVirtualField("evenMonths", boolean.class, (month) == (14));
    boolean anyMonth = readVirtualField("anyMonth", boolean.class, (month) == (0xFF));

    short weekOfMonth = readSimpleField("weekOfMonth", readUnsignedShort(readBuffer, 8));
    boolean days1to7 = readVirtualField("days1to7", boolean.class, (weekOfMonth) == (1));
    boolean days8to14 = readVirtualField("days8to14", boolean.class, (weekOfMonth) == (2));
    boolean days15to21 = readVirtualField("days15to21", boolean.class, (weekOfMonth) == (3));
    boolean days22to28 = readVirtualField("days22to28", boolean.class, (weekOfMonth) == (4));
    boolean days29to31 = readVirtualField("days29to31", boolean.class, (weekOfMonth) == (5));
    boolean last7DaysOfThisMonth =
        readVirtualField("last7DaysOfThisMonth", boolean.class, (weekOfMonth) == (6));
    boolean any7DaysPriorToLast7DaysOfThisMonth =
        readVirtualField(
            "any7DaysPriorToLast7DaysOfThisMonth", boolean.class, (weekOfMonth) == (7));
    boolean any7DaysPriorToLast14DaysOfThisMonth =
        readVirtualField(
            "any7DaysPriorToLast14DaysOfThisMonth", boolean.class, (weekOfMonth) == (8));
    boolean any7DaysPriorToLast21DaysOfThisMonth =
        readVirtualField(
            "any7DaysPriorToLast21DaysOfThisMonth", boolean.class, (weekOfMonth) == (9));
    boolean anyWeekOfthisMonth =
        readVirtualField("anyWeekOfthisMonth", boolean.class, (weekOfMonth) == (0xFF));

    short dayOfWeek = readSimpleField("dayOfWeek", readUnsignedShort(readBuffer, 8));
    boolean anyDayOfWeek = readVirtualField("anyDayOfWeek", boolean.class, (dayOfWeek) == (0xFF));

    readBuffer.closeContext("BACnetWeekNDayTagged");
    // Create the instance
    BACnetWeekNDayTagged _bACnetWeekNDayTagged;
    _bACnetWeekNDayTagged =
        new BACnetWeekNDayTagged(header, month, weekOfMonth, dayOfWeek, tagNumber, tagClass);
    return _bACnetWeekNDayTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetWeekNDayTagged)) {
      return false;
    }
    BACnetWeekNDayTagged that = (BACnetWeekNDayTagged) o;
    return (getHeader() == that.getHeader())
        && (getMonth() == that.getMonth())
        && (getWeekOfMonth() == that.getWeekOfMonth())
        && (getDayOfWeek() == that.getDayOfWeek())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getMonth(), getWeekOfMonth(), getDayOfWeek());
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
