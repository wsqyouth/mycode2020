package main

import (
	"fmt"
	"bufio"
	"log"
	"encoding/gob"
)

func GobCodec struce {
	conn io.ReadWriteCloser 
	buf *bufio.Writer 
	dec *gob.Decoder
	ecn *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn:conn,
		buf:buf,
		dec:gob.NewDecoder(conn),
		enc:gob.NewEncoder(buff),
	}
}

func(c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func(c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func(c *GobCodec) Write(h *Header, body interface{})(err error) {
	defer func(){
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}
	if err = c.enc.Encode(h);err!= nil {
		log.Println("rpc: gob encoding header err:",err)
		return 
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: gob encoding body err:",err)
		return 
	}
	return 
}

func(c *GobCodec) Close() error {
	return c.conn.Close()
}
