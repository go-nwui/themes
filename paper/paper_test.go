package paper

import (
	"testing"

	"github.com/go-nwui/nwui"
)

func Test_All(*testing.T) {
	w := nwui.NewWindow("nwui paper theme", 800, 600)
	w.UseTheme(Theme())
}
