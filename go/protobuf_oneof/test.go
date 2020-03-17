package main

import (
	"fmt"
	"test/mes"
)

func main() {

	// https://developers.google.com/protocol-buffers/docs/reference/go-generated#oneof

	r := &mes.ResourceReply{
		Reason: &mes.ResourceReply_BadLotIDReasons_{
			BadLotIDReasons: mes.ResourceReply_IN_INSPECTION,
		},
	}

	switch t := r.GetReason().(type) {
	case *mes.ResourceReply_BadLotIDReasons_:
		switch t.BadLotIDReasons {
		case mes.ResourceReply_UNSPECIFIED_BAD_LOT_ID_REASON :
			fmt.Println(t)
		case mes.ResourceReply_UNAVAILABLE                   :
			fmt.Println(t)
		case mes.ResourceReply_EXPIRED                       :
			fmt.Println(t)
		case mes.ResourceReply_ON_HOLD                       :
			fmt.Println(t)
		case mes.ResourceReply_IN_INSPECTION                 :
			fmt.Println(mes.ResourceReply_IN_INSPECTION)
		case mes.ResourceReply_ALREADY_MOUNTED               :
			fmt.Println(t)
		case mes.ResourceReply_MISS_PRECONDITION             :
			fmt.Println(t)
		}

	case *mes.ResourceReply_BadQuantityReasons_:
		fmt.Println(t.BadQuantityReasons)
	case *mes.ResourceReply_InternalReasons_:
		fmt.Println(t.InternalReasons)
	case nil:
		// the field is not set
	}

	if r.GetReason() != nil  {
		fmt.Println(fmt.Sprintf("%v",r.GetReason()))
	}
}
