package server

import (
	"time"

	"github.com/pkg/errors"
)

const (
	// epoch is the number of milliseconds since Unix epoch
	epoch int64 = 1609459200000 // 2021-01-01 00:00:00 UTC
	// machineBits is the number of bits to represent machine ID
	machineBits uint8 = 10
	// maxMachineID is the maximum machine ID that can be used
	maxMachineID int64 = -1 ^ (-1 << machineBits)
	// sequenceBits is the number of bits to represent sequence number
	sequenceBits uint8 = 12
	// machineShift is the number of bits to shift for machine ID
	machineShift = sequenceBits
	// timestampShift is the number of bits to shift for timestamp
	timestampShift = sequenceBits + machineBits
	// sequenceMask is the mask for sequence number
	sequenceMask int64 = -1 ^ (-1 << sequenceBits)
)

// New creates a new Snowflake instance with the given machine ID.
func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("invalid machine ID")
	}
	return &Snowflake{
		machineID:     machineID,
		sequence:      0,
		lastTimestamp: -1,
	}, nil
}

var SnowflakeSrv *Snowflake

// Snowflake represents a distributed unique ID generator.
type Snowflake struct {
	machineID     int64
	sequence      int64
	lastTimestamp int64
}

// NextID generates the next unique ID.
func (s *Snowflake) NextID() int64 {
	timestamp := time.Now().UnixNano() / 1000000 // convert to milliseconds
	if timestamp < s.lastTimestamp {
		timestamp = s.lastTimestamp
	}
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			timestamp = s.waitNextMillisecond(timestamp)
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timestamp
	return ((timestamp - epoch) << timestampShift) | (s.machineID << machineShift) | s.sequence
}

// waitNextMillisecond waits until the next millisecond.
func (s *Snowflake) waitNextMillisecond(timestamp int64) int64 {
	for timestamp <= s.lastTimestamp {
		time.Sleep(time.Millisecond)
		timestamp = time.Now().UnixNano() / 1000000
	}
	return timestamp
}
