// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package basics

import (
	"bytes"
	"encoding/base32"
	"fmt"

	"github.com/algorand/go-algorand/crypto"
)

type (
	// Address is a unique identifier corresponding to ownership of money
	Address crypto.Digest
)

// ChecksumAddress is a representation of the short address with a checksum
type ChecksumAddress struct {
	shortAddress Address
	checksum     []byte
}

const (
	checksumLength = 4
)

// GetChecksumAddress returns the short address with its checksum as a string
// Checksum in Algorand are the last 4 bytes of the shortAddress Hash. H(Address)[28:]
func (addr Address) GetChecksumAddress() *ChecksumAddress {
	shortAddressHash := crypto.Hash(addr[:])
	return &ChecksumAddress{addr, shortAddressHash[len(shortAddressHash)-checksumLength:]}
}

// GetUserAddress returns the human-readable, checksummed version of the address
func (addr Address) GetUserAddress() string {
	return addr.GetChecksumAddress().String()
}

// IsValid returns true if the address is valid, false otherwise
func (addr ChecksumAddress) IsValid() bool {
	shortAddressHash := crypto.Hash(addr.shortAddress[:])
	return bytes.Equal(shortAddressHash[len(shortAddressHash)-checksumLength:], addr.checksum)
}

// Address returns the address's Address
func (addr ChecksumAddress) Address() Address {
	return addr.shortAddress
}

// UnmarshalChecksumAddress tries to unmarshal the checksummed address string.
func UnmarshalChecksumAddress(address string) (Address, error) {
	decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(address)
	if err != nil {
		return Address{}, fmt.Errorf("failed to decode address %s to base 32", address)
	}
	var short Address
	if len(decoded) < len(short) {
		return Address{}, fmt.Errorf("decoded bad address: %v", address)
	}
	copy(short[:], decoded[:len(short)])

	checksumAddr := ChecksumAddress{short, decoded[len(decoded)-checksumLength:]}
	if !checksumAddr.IsValid() {
		return Address{}, fmt.Errorf("address %s is malformed, checksum verification failed", address)
	}

	// Validate that we had a canonical string representation
	if checksumAddr.String() != address {
		return Address{}, fmt.Errorf("address %s is non-canonical", address)
	}

	return short, nil
}

// String returns a string representation of ChecksumAddress
func (addr *ChecksumAddress) String() string {
	var addrWithChecksum []byte
	addrWithChecksum = append(addr.shortAddress[:], addr.checksum...)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(addrWithChecksum)
}

func (addr Address) String() string {
	return fmt.Sprintf("%v", crypto.Digest(addr))
}

// MarshalText returns the address string as an array of bytes
func (addr Address) MarshalText() ([]byte, error) {
	return []byte(addr.String()), nil
}

// UnmarshalText initializes the Address from an array of bytes.
func (addr *Address) UnmarshalText(text []byte) error {
	d, err := crypto.DigestFromString(string(text))
	if err == nil {
		*addr = Address(d)
	}
	return err
}
