package server 

import (
    "errors"
)

//First, define pass in and return parameters
type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

//Second, define serve object
type Arith int

//Third, implement methods of this type
func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("Divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}


