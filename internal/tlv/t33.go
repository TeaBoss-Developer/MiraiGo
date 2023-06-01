package tlv

import "github.com/TeaBoss-Developer/MiraiGo/binary"

func T33(guid []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x33)
		w.WriteBytesShort(guid)
	})
}
