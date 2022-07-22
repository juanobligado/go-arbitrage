// Utility functions to work with big numbers
package utils

import (
	"fmt"
	"math/big"
)


type BigInt struct {
    big.Int
}

type BigFloat struct {
    big.Float
}

func (b BigInt) MarshalJSON() ([]byte, error) {
    return []byte(b.String()), nil
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
    if string(p) == "null" {
        return nil
    }
    var z big.Int
    _, ok := z.SetString(string(p), 10)
    if !ok {
        return fmt.Errorf("not a valid big integer: %s", p)
    }
    b.Int = z
    return nil
}


func (b BigFloat) MarshalJSON() ([]byte, error) {
    return []byte(b.String()), nil
}

func (b *BigFloat) UnmarshalJSON(p []byte) error {
    if string(p) == "null" {
        return nil
    }
    var z big.Float
    _, ok := z.SetString(string(p))
    if !ok {
        return fmt.Errorf("not a valid big float: %s", p)
    }
    b.Float = z
    return nil
}