package god

import (
	"log"
	"testing"
)

func TestUtils_NewSmap(t *testing.T) {
	sm := NewSmap()

	//sm.SetPad("KEY%20s\n")

	sm. String("Local", "local").
		String("Maal", "maal").
		Int(23, "teis").
		Int(98, "attanwe")

	log.Fatalln("\n", sm.SprintLn())
}
