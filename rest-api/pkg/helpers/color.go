package helpers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/crazy3lf/colorconv"
)

func ColorRandPastel() string {
	rand.Seed(time.Now().UnixNano())
	h := rand.Intn(360)
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(95-25+1) + 25
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(97-92+1) + 92
	c, _ := colorconv.HSLToColor(float64(h), float64(s)/100, float64(l)/100)
	return strings.Replace(colorconv.ColorToHex(c), "0x", "#", 1)
}
