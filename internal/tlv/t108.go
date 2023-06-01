package tlv

import "github.com/TeaBoss-Developer/MiraiGo/binary"

func T108(ksid []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x108)
		w.WriteBytesShort(ksid)
	})
}
