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
package org.apache.plc4x.java.cbus.readwrite;

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

public class TriggerControlDataLabel extends TriggerControlData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final TriggerControlLabelOptions triggerControlOptions;
  protected final byte actionSelector;
  protected final Language language;
  protected final byte[] data;

  public TriggerControlDataLabel(
      TriggerControlCommandTypeContainer commandTypeContainer,
      byte triggerGroup,
      TriggerControlLabelOptions triggerControlOptions,
      byte actionSelector,
      Language language,
      byte[] data) {
    super(commandTypeContainer, triggerGroup);
    this.triggerControlOptions = triggerControlOptions;
    this.actionSelector = actionSelector;
    this.language = language;
    this.data = data;
  }

  public TriggerControlLabelOptions getTriggerControlOptions() {
    return triggerControlOptions;
  }

  public byte getActionSelector() {
    return actionSelector;
  }

  public Language getLanguage() {
    return language;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeTriggerControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TriggerControlDataLabel");

    // Simple Field (triggerControlOptions)
    writeSimpleField(
        "triggerControlOptions",
        triggerControlOptions,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (actionSelector)
    writeSimpleField("actionSelector", actionSelector, writeByte(writeBuffer, 8));

    // Optional Field (language) (Can be skipped, if the value is null)
    writeOptionalEnumField(
        "language",
        "Language",
        language,
        new DataWriterEnumDefault<>(
            Language::getValue, Language::name, writeUnsignedShort(writeBuffer, 8)),
        (getTriggerControlOptions().getLabelType()) != (TriggerControlLabelType.LOAD_DYNAMIC_ICON));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("TriggerControlDataLabel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TriggerControlDataLabel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (triggerControlOptions)
    lengthInBits += triggerControlOptions.getLengthInBits();

    // Simple field (actionSelector)
    lengthInBits += 8;

    // Optional Field (language)
    if (language != null) {
      lengthInBits += 8;
    }

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static TriggerControlDataBuilder staticParseTriggerControlDataBuilder(
      ReadBuffer readBuffer, TriggerControlCommandTypeContainer commandTypeContainer)
      throws ParseException {
    readBuffer.pullContext("TriggerControlDataLabel");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TriggerControlLabelOptions triggerControlOptions =
        readSimpleField(
            "triggerControlOptions",
            readComplex(() -> TriggerControlLabelOptions.staticParse(readBuffer), readBuffer));

    byte actionSelector = readSimpleField("actionSelector", readByte(readBuffer, 8));

    Language language =
        readOptionalField(
            "language",
            new DataReaderEnumDefault<>(Language::enumForValue, readUnsignedShort(readBuffer, 8)),
            (triggerControlOptions.getLabelType()) != (TriggerControlLabelType.LOAD_DYNAMIC_ICON));

    byte[] data =
        readBuffer.readByteArray(
            "data",
            Math.toIntExact(
                ((commandTypeContainer.getNumBytes())
                    - ((((((triggerControlOptions.getLabelType())
                            != (TriggerControlLabelType.LOAD_DYNAMIC_ICON)))
                        ? (4)
                        : (3)))))));

    readBuffer.closeContext("TriggerControlDataLabel");
    // Create the instance
    return new TriggerControlDataLabelBuilderImpl(
        triggerControlOptions, actionSelector, language, data);
  }

  public static class TriggerControlDataLabelBuilderImpl
      implements TriggerControlData.TriggerControlDataBuilder {
    private final TriggerControlLabelOptions triggerControlOptions;
    private final byte actionSelector;
    private final Language language;
    private final byte[] data;

    public TriggerControlDataLabelBuilderImpl(
        TriggerControlLabelOptions triggerControlOptions,
        byte actionSelector,
        Language language,
        byte[] data) {
      this.triggerControlOptions = triggerControlOptions;
      this.actionSelector = actionSelector;
      this.language = language;
      this.data = data;
    }

    public TriggerControlDataLabel build(
        TriggerControlCommandTypeContainer commandTypeContainer, byte triggerGroup) {
      TriggerControlDataLabel triggerControlDataLabel =
          new TriggerControlDataLabel(
              commandTypeContainer,
              triggerGroup,
              triggerControlOptions,
              actionSelector,
              language,
              data);
      return triggerControlDataLabel;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TriggerControlDataLabel)) {
      return false;
    }
    TriggerControlDataLabel that = (TriggerControlDataLabel) o;
    return (getTriggerControlOptions() == that.getTriggerControlOptions())
        && (getActionSelector() == that.getActionSelector())
        && (getLanguage() == that.getLanguage())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTriggerControlOptions(),
        getActionSelector(),
        getLanguage(),
        getData());
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
