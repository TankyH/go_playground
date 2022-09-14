package package_exportable

import "goplayground/v1/basic/chapter_1/package_exportable/utils"

func ImportScope2() int {
	// unexported type
	var mt utils.myInt = 0
	// unexported function
	return utils.getMyInt(mt)
}
