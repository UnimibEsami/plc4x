#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from plc4py.api.exceptions.exceptions import PlcRuntimeException
from plc4py.api.exceptions.exceptions import SerializationException
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.umas.readwrite.DriverType import DriverType
from plc4py.protocols.umas.readwrite.UmasADU import UmasADU
from plc4py.protocols.umas.readwrite.UmasADU import UmasADUBuilder
from plc4py.protocols.umas.readwrite.UmasPDU import UmasPDU
from plc4py.spi.generation.ReadBuffer import ReadBuffer
from plc4py.spi.generation.WriteBuffer import WriteBuffer
from plc4py.utils.GenericTypes import ByteOrder
import math


@dataclass
class UmasTcpADU(UmasADU):
    transaction_identifier: int
    unit_identifier: int
    pdu: UmasPDU
    # Arguments.
    response: bool
    PROTOCOL_IDENTIFIER: int = 0x0000
    # Accessors for discriminator values.
    driver_type: DriverType = DriverType.Umas_TCP

    def serialize_umas_adu_child(self, write_buffer: WriteBuffer):
        write_buffer.push_context("UmasTcpADU")

        # Simple Field (transactionIdentifier)
        write_buffer.write_unsigned_short(
            self.transaction_identifier, logical_name="transactionIdentifier"
        )

        # Const Field (protocolIdentifier)
        write_buffer.write_unsigned_short(
            self.PROTOCOL_IDENTIFIER, logical_name="protocolIdentifier"
        )

        # Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        length: int = self.pdu.length_in_bytes() + int(1)
        write_buffer.write_unsigned_short(length, logical_name="length")

        # Simple Field (unitIdentifier)
        write_buffer.write_unsigned_byte(
            self.unit_identifier, logical_name="unitIdentifier"
        )

        # Simple Field (pdu)
        write_buffer.write_serializable(self.pdu, logical_name="pdu")

        write_buffer.pop_context("UmasTcpADU")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.length_in_bits() / 8.0)))

    def length_in_bits(self) -> int:
        length_in_bits: int = super().length_in_bits()
        _value: UmasTcpADU = self

        # Simple field (transactionIdentifier)
        length_in_bits += 16

        # Const Field (protocolIdentifier)
        length_in_bits += 16

        # Implicit Field (length)
        length_in_bits += 16

        # Simple field (unitIdentifier)
        length_in_bits += 8

        # Simple field (pdu)
        length_in_bits += self.pdu.length_in_bits()

        return length_in_bits

    @staticmethod
    def static_parse_builder(
        read_buffer: ReadBuffer, driver_type: DriverType, response: bool
    ):
        read_buffer.push_context("UmasTcpADU")

        transaction_identifier: int = read_buffer.read_unsigned_short(
            logical_name="transactionIdentifier",
            bit_length=16,
            byte_order=ByteOrder.BIG_ENDIAN,
            driver_type=driver_type,
            response=response,
        )

        PROTOCOL_IDENTIFIER: int = read_buffer.read_unsigned_short(
            logical_name="protocolIdentifier",
            byte_order=ByteOrder.BIG_ENDIAN,
            driver_type=driver_type,
            response=response,
        )

        length: int = read_buffer.read_unsigned_short(
            logical_name="length",
            byte_order=ByteOrder.BIG_ENDIAN,
            driver_type=driver_type,
            response=response,
        )

        unit_identifier: int = read_buffer.read_unsigned_byte(
            logical_name="unitIdentifier",
            bit_length=8,
            byte_order=ByteOrder.BIG_ENDIAN,
            driver_type=driver_type,
            response=response,
        )

        pdu: UmasPDU = read_buffer.read_complex(
            read_function=UmasPDU.static_parse,
            logical_name="pdu",
            byte_order=ByteOrder.BIG_ENDIAN,
            driver_type=driver_type,
            response=response,
        )

        read_buffer.pop_context("UmasTcpADU")
        # Create the instance
        return UmasTcpADUBuilder(transaction_identifier, unit_identifier, pdu)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, UmasTcpADU):
            return False

        that: UmasTcpADU = UmasTcpADU(o)
        return (
            (self.transaction_identifier == that.transaction_identifier)
            and (self.unit_identifier == that.unit_identifier)
            and (self.pdu == that.pdu)
            and super().equals(that)
            and True
        )

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        pass
        # write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        # try:
        #    write_buffer_box_based.writeSerializable(self)
        # except SerializationException as e:
        #    raise PlcRuntimeException(e)

        # return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class UmasTcpADUBuilder(UmasADUBuilder):
    transaction_identifier: int
    unit_identifier: int
    pdu: UmasPDU

    def build(
        self,
        response: bool,
    ) -> UmasTcpADU:
        umas_tcp_adu: UmasTcpADU = UmasTcpADU(
            response, self.transaction_identifier, self.unit_identifier, self.pdu
        )
        return umas_tcp_adu
