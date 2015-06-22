package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/glycerine/go-capnproto"
	"github.com/meteorhacks/bddp"
)

func main() {
	c := bddp.NewClient(":5000")

	if err := c.Connect(); err != nil {
		panic(err)
	}

	var i int64
	for i = 0; true; i++ {
		time.Sleep(time.Second)
		if err := call(c, i); err != nil {
			fmt.Println(err)
		}
	}
}

func call(c bddp.Client, n int64) (err error) {
	call, err := c.Method("double")
	if err != nil {
		return err
	}

	buff := make([]byte, 8, 8)
	binary.PutVarint(buff, n)

	seg := call.Segment()
	obj := capn.Object(seg.NewData(buff))

	res, err := call.Call(obj)
	if err != nil {
		return err
	}

	out, _ := binary.Varint(res.ToData())
	if n*2 != out {
		return errors.New("incorrect value")
	}

	fmt.Printf("%d * 2 => %d \n", n, out)

	return nil
}
