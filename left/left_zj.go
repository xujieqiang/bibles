package left

import "log"

type Left_zhangjie struct {
}

func Newleftzj() Left_zhangjie {
	return Left_zhangjie{}
}

func (lz *Left_zhangjie) Drawzhangjie() {

	if clickId != 0 {
		log.Println(clickId)
	}
}
