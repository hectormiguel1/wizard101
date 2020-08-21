package parser
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

const SizeOfShort = 2
const SizeOfInt = 4
const SizeOfByte = 1
const SizeOfFloat = 4
const SizeOfGuid = 8
const ZeroInitializer = 0
const BeginningOfArray = 0

/**
Function discards n bytes from byte array,
returned trimmed array.
*/
func ReadN(byteArray []byte, bytesToRead int) []byte {
	b := byteArray[bytesToRead:]
	return b
}

/*
Function returns string from byte array and the remaining bytes in array.
*/
func readString(byteArray []byte) ([]byte, string) {
	till := ZeroInitializer
	const StringNullChar = 0x00
	for byteArray[till] != StringNullChar {
		till++
	}
	outputString := string(byteArray[ZeroInitializer:till])
	return byteArray[till:], outputString
}

func ReadStringWithLength(byteArray []byte, length uint16) ([]byte, []byte) {
	return byteArray[length:], byteArray[:length]
}

/*
Function returns signed short (int16) from byte array and the remaining bytes in array.
*/
func ReadShort(byteArray []byte) ([]byte, int16) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfShort])
	var shortValue int16 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &shortValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_SHORT_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfShort:], shortValue
}

/*
Function returns an unsigned short (uint16) from byte array and the remaining bytes in array.
*/
func ReadUShort(byteArray []byte) ([]byte, uint16) {
	buff := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfShort])
	var sortValue uint16 = ZeroInitializer
	err := binary.Read(buff, binary.LittleEndian, &sortValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_U_SHORT_FUNC: Error Reading from Buffer")
	}
	return byteArray[SizeOfShort:], sortValue
}

/*
Function returns int32 from byte array and the remaining bytes in array.
*/
func ReadInt(byteArray []byte) ([]byte, int32) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfInt])
	var intValue int32 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &intValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_INT_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfInt:], intValue
}

func ReadUInt(byteArray []byte) ([]byte, uint32) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfInt])
	var intValue uint32 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &intValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_U_INT_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfInt:], intValue
}

/*
Function returns signed byte (int8) from byte array and the remaining bytes in array.
*/
func ReadByte(byteArray []byte) ([]byte, int8) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfByte])
	var intValue int8 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &intValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_BYTE_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfByte:], intValue
}

/*
Function returns unsigned byte (uint8) from byte array and the remaining bytes in array.
*/
func ReadUByte(byteArray []byte) ([]byte, uint8) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray:SizeOfByte])
	var intValue uint8 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &intValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_U_BYTE_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfByte:], intValue
}

/*
Function returns float from byte array and the remaining bytes in array.
*/
func ReadFloat(byteArray []byte) ([]byte, float32) {
	floatBytes := byteArray[BeginningOfArray:SizeOfFloat]
	remainArray := byteArray[SizeOfFloat:]
	resultingFloat := math.Float32frombits(binary.LittleEndian.Uint32(floatBytes))
	return remainArray, resultingFloat
}
/*
Function returns guid string from byte array and the remaining bytes in array.
*/
func ReadGuid(byteArray []byte) ([]byte, uint64) {
	numBytes := bytes.NewBuffer(byteArray[BeginningOfArray : SizeOfGuid+1])
	var intValue uint64 = ZeroInitializer
	err := binary.Read(numBytes, binary.LittleEndian, &intValue)
	if err != nil {
		fmt.Println("(PARSER.GO) READ_U_INT_FUNC: Error Reading from Buffer")

	}
	return byteArray[SizeOfGuid:], intValue
}
