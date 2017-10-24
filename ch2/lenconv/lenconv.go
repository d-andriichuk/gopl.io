package lenconv

import (
	"fmt"
)

type Kilometre float64
type Mile float64

func (km Kilometre) String() string { return fmt.Sprintf("%gkm", km) }
func (mile Mile) String() string    { return fmt.Sprintf("%gmile", mile) }
