package err

import "errors"

var ErrInvalidExt = errors.New("File have to be jpg or png")
var ErrNoRowsUpdated = errors.New("No data updated")
