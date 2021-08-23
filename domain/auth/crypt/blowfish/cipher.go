// This is a slightly altered version of the vanilla blowfish package.

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package blowfish implements Bruce Schneier's Blowfish encryption algorithm.
package blowfish

// The code is a port of Bruce Schneier's C implementation.
// See http://www.schneier.com/blowfish.html.

import "strconv"

// The Blowfish block size in bytes.
const BlockSize = 8

// A Cipher is an instance of Blowfish encryption using a particular key.
type Cipher struct {
	p              [18]uint32
	s0, s1, s2, s3 [256]uint32
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/blowfish: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates and returns a Cipher.
// The key argument should be the Blowfish key, from 1 to 56 bytes.
func NewCipher(key []byte) (*Cipher, error) {
	var result Cipher
	if k := len(key); k < 1 || k > 56 {
		return nil, KeySizeError(k)
	}
	initCipher(&result)
	ExpandKey(key, &result)
	return &result, nil
}

// NewSaltedCipher creates a returns a Cipher that folds a salt into its key
// schedule. For most purposes, NewCipher, instead of NewSaltedCipher, is
// sufficient and desirable. For bcrypt compatiblity, the key can be over 56
// bytes.
func NewSaltedCipher(key, salt []byte) (*Cipher, error) {
	if len(salt) == 0 {
		return NewCipher(key)
	}
	var result Cipher
	if k := len(key); k < 1 {
		return nil, KeySizeError(k)
	}
	initCipher(&result)
	expandKeyWithSalt(key, salt, &result)
	return &result, nil
}

// BlockSize returns the Blowfish block size, 8 bytes.
// It is necessary to satisfy the Block interface in the
// package "crypto/cipher".
func (c *Cipher) BlockSize() int { return BlockSize }

// Encrypt encrypts the 8-byte buffer src using the key k
// and stores the result in dst.
// Note that for amounts of data larger than a block,
// it is not safe to just call Encrypt on successive blocks;
// instead, use an encryption mode like CBC (see crypto/cipher/cbc.go).
func (c *Cipher) Encrypt(dst, src []byte) {
	// modified bit conversion
	l := uint32(src[3])<<24 | uint32(src[2])<<16 | uint32(src[1])<<8 | uint32(src[0])
	r := uint32(src[7])<<24 | uint32(src[6])<<16 | uint32(src[5])<<8 | uint32(src[4])

	l, r = encryptBlock(l, r, c)

	// modified bit conversion
	dst[3], dst[2], dst[1], dst[0] = byte(l>>24), byte(l>>16), byte(l>>8), byte(l)
	dst[7], dst[6], dst[5], dst[4] = byte(r>>24), byte(r>>16), byte(r>>8), byte(r)
}

// Decrypt decrypts the 8-byte buffer src using the key k
// and stores the result in dst.
func (c *Cipher) Decrypt(dst, src []byte) {
	// modified bit conversion
	l := uint32(src[3])<<24 | uint32(src[2])<<16 | uint32(src[1])<<8 | uint32(src[0])
	r := uint32(src[7])<<24 | uint32(src[6])<<16 | uint32(src[5])<<8 | uint32(src[4])

	l, r = decryptBlock(l, r, c)

	// modified bit conversion
	dst[3], dst[2], dst[1], dst[0] = byte(l>>24), byte(l>>16), byte(l>>8), byte(l)
	dst[7], dst[6], dst[5], dst[4] = byte(r>>24), byte(r>>16), byte(r>>8), byte(r)
}

func initCipher(c *Cipher) {
	copy(c.p[0:], p[0:])
	copy(c.s0[0:], s0[0:])
	copy(c.s1[0:], s1[0:])
	copy(c.s2[0:], s2[0:])
	copy(c.s3[0:], s3[0:])
}
