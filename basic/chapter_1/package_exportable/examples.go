package package_exportable

import "goplayground/v1/basic/chapter_1/package_exportable/utils"

func ImportScope() int {
	var mt utils.MyInt = 0
	return utils.GetMyInt(mt)
}
