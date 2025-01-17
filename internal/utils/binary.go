package utils

import (
	"encoding/binary"
	"fmt"
	"io"
)


func writeSafeSliceLength(w io.Writer, length int) error {
    return binary.Write(w, binary.LittleEndian, uint32(length))
}

// ReadSafeSliceLength reads a slice length as uint32
func readSafeSliceLength(r io.Reader) (int, error) {
    var length uint32
    err := binary.Read(r, binary.LittleEndian, &length)
    return int(length), err
}

func WriteSlice[T int32 | float32](w io.Writer, slice []T) error {
    // Write the length of the slice
    if err := writeSafeSliceLength(w, len(slice)); err != nil {
        return err
    }

    // Write based on type
    switch any(slice).(type) {
    case []int32:
        // Write int slice using WriteInt32
        for _, value := range slice {
            if err := WriteInt32(w, int32(value)); err != nil {
                return err
            }
        }
    case []float32:
        // Write float32 slice
        for _, value := range slice {
            if err := WriteFloat32(w, float32(value)); err != nil {
                return err
            }
        }
    default:
        return fmt.Errorf("unsupported slice type")
    }

    return nil
}

func ReadSlice[T int32|float32](r io.Reader) ([]T, error) {
    // Read the length of the slice
    length, err := readSafeSliceLength(r)
    if err != nil {
        return nil, err
    }

    // Create slice with the read length
    slice := make([]T, length)
	
	// Read each element based on type
	for i := 0; i < length; i++ {
		var value T
		if err := binary.Read(r, binary.LittleEndian, &value); err != nil {
			return nil, err
		}
		slice[i] = value
	}


    return slice, nil
}

func WriteString(w io.Writer, s string) error {
    // Convert string to bytes
    nameBytes := []byte(s)
    
    // Write length of the string as uint32
    if err := binary.Write(w, binary.LittleEndian, uint32(len(nameBytes))); err != nil {
        return err
    }
    
    // Write the string bytes
    if _, err := w.Write(nameBytes); err != nil {
        return err
    }
    
    return nil
}

func ReadString(r io.Reader) (string, error) {
    // Read the length of the string
    var strLength uint32
    if err := binary.Read(r, binary.LittleEndian, &strLength); err != nil {
        return "", err
    }
    
    // Create byte slice to read the string
    strBytes := make([]byte, strLength)
    
    // Read the string bytes
    if _, err := io.ReadFull(r, strBytes); err != nil {
        return "", err
    }
    
    // Convert bytes back to string
    return string(strBytes), nil
}

// Helper functions for specific types
func WriteInt32(w io.Writer, value int32) error {
    return binary.Write(w, binary.LittleEndian, value)
}

func ReadInt32(r io.Reader) (int32, error) {
    var intValue int32
    err := binary.Read(r, binary.LittleEndian, &intValue)
    return intValue, err
}

func WriteFloat32(w io.Writer, value float32) error {
    return binary.Write(w, binary.LittleEndian, value)
}

func ReadFloat32(r io.Reader) (float32, error) {
    var floatValue float32
    err := binary.Read(r, binary.LittleEndian, &floatValue)
    return floatValue, err
}

