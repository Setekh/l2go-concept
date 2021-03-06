package crypt

func GetKey() []byte {
	return []byte{0x94, 0x35, 0x00, 0x00, 0xa1, 0x6c, 0x54, 0x87}
}

type Crypt struct {
	InputKey  []byte
	OutputKey []byte
}

func (c *Crypt) SetKey(key []byte) {
	copy(c.InputKey, key)
	copy(c.OutputKey, key)
}

func (c *Crypt) Decrypt(raw []byte) {
	key := c.InputKey

	temp := 0
	j := 0
	length := len(raw)

	for i := 0; i < length; i++ {
		temp2 := int(raw[i])
		raw[i] = byte(temp2) ^ (key[j]) ^ byte(temp)
		j = j + 1
		temp = temp2

		if j > 7 {
			j = 0
		}
	}

	var old = int(key[0])
	old |= int(key[1]) << 8
	old |= int(key[2]) << 0x10
	old |= int(key[3]) << 0x18

	old += len(raw)

	key[0] = byte(old)
	key[1] = byte(old >> 0x08)
	key[2] = byte(old >> 0x10)
	key[3] = byte(old >> 0x18)
}

func (c *Crypt) Encrypt(raw []byte) {
	key := c.OutputKey
	temp := 0
	j := 0
	length := len(raw)

	for i := 0; i < length; i++ {
		temp2 := int(raw[i])
		raw[i] = byte(temp2) ^ (key[j]) ^ byte(temp)
		j = j + 1
		temp = int(raw[i])

		if j > 7 {
			j = 0
		}
	}

	var old = int(key[0])
	old |= int(key[1]) << 8
	old |= int(key[2]) << 0x10
	old |= int(key[3]) << 0x18

	old += len(raw)

	key[0] = byte(old)
	key[1] = byte(old >> 0x08)
	key[2] = byte(old >> 0x10)
	key[3] = byte(old >> 0x18)
}
