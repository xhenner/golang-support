package uuid

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
}

func BUUID() []byte {
	var entropy = make([]byte, 16)

	rand.Read(entropy)
	entropy[6] = (entropy[6] & 0x0f) | 0x40
	entropy[8] = (entropy[8] & 0x3f) | 0x80
	return entropy
}

func UUID() string {
	entropy := BUUID()
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", entropy[0:4], entropy[4:6], entropy[6:8], entropy[8:10], entropy[10:16])
}
