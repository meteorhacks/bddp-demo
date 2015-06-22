package main

import (
	"encoding/binary"
	"fmt"

	"github.com/glycerine/go-capnproto"
	"github.com/meteorhacks/bddp"
)

func main() {
	s := bddp.NewServer(":5000")
	s.Method("double", double)
	if err := s.Listen(); err != nil {
		panic(err)
	}
}

func double(ctx bddp.MContext) {
	params := *ctx.Params()
	n, _ := binary.Varint(params.ToData())
	buff := make([]byte, 8, 8)
	binary.PutVarint(buff, n*2)
	seg := ctx.Segment()
	obj := capn.Object(seg.NewData(buff))
	fmt.Printf("%d * 2 => %d \n", n, n*2)
	ctx.SendResult(&obj)
	ctx.SendUpdated()
}
