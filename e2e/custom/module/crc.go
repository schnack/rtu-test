package module

import (
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	Mod256 = "mod256"
	ModBus = "modBus"
)

type Crc struct {
	Algorithm string   `yaml:"algorithm"`
	Staffing  bool     `yaml:"staffing"`
	ByteOrder string   `yaml:"byteOrder"`
	Read      []string `yaml:"read"`
	Write     []string `yaml:"write"`
	Error     []string `yaml:"error"`
}

// Подсчет контрольной суммы согласно алгоритму
func (c *Crc) Calc(order binary.ByteOrder, data []byte) []byte {
	switch c.Algorithm {
	case Mod256:
		return []byte{c.CrcMod256(data)}
	case ModBus:
		b := make([]byte, 2)
		// Значит используется параметр по умолчанию
		if c.ByteOrder == "" {
			order.PutUint16(b, c.ModBusCRC(data))
		} else if c.ByteOrder == "little" {
			binary.LittleEndian.PutUint16(b, c.ModBusCRC(data))
		} else {
			binary.BigEndian.PutUint16(b, c.ModBusCRC(data))
		}
		return b
	}
	return nil
}

// CheckSum8 Modulo 256.
func (c *Crc) CrcMod256(data []byte) uint8 {
	var sum uint8
	for _, b := range data {
		sum += b
	}
	return sum & 0xff
}

func (c *Crc) ModBusCRC(data []byte) (crc uint16) {
	crc = 0xffff
	for _, v := range data {
		crc = (crc >> 8) ^ initTableCRC()[(crc^uint16(v))&0x00FF]
	}
	return
}

// Len - возвращает длину данных crc
func (c *Crc) Len() int {
	switch c.Algorithm {
	case Mod256:
		return 1
	case ModBus:
		return 2
	default:
		logrus.Fatalf("crc algorithm no support %s", c.Algorithm)
	}
	return 0
}

// Derived from https://github.com/lammertb/libcrc
/*
 * Library: libcrc
 * File:    src/crc16.c
 * Author:  Lammert Bies
 *
 * This file is licensed under the MIT License as stated below
 *
 * Copyright (c) 1999-2016 Lammert Bies
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 * Description
 * -----------
 * The source file src/crc16.c contains routines which calculate the common
 * CRC16 cyclic redundancy check values for an incomming byte string.
 */

var instanceTableCRC []uint16
var syncTableCRC sync.Once

func initTableCRC() []uint16 {
	syncTableCRC.Do(func() {
		crc16IBM := uint16(0xA001)
		instanceTableCRC = make([]uint16, 256)
		for i := uint16(0); i < 256; i++ {

			var crc uint16
			c := i

			for j := uint16(0); j < 8; j++ {
				if ((crc ^ c) & 0x0001) > 0 {
					crc = (crc >> 1) ^ crc16IBM
				} else {
					crc = crc >> 1
				}
				c = c >> 1
			}
			instanceTableCRC[i] = crc
		}
	})
	return instanceTableCRC
}
