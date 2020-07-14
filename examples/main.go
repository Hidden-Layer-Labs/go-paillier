package main

import (
	"github.com/hidden-layer-labs/go-paillier"
	"fmt"
)

func main() {
	pk, sk := paillier.GenerateKeypair().ToKeys()

	str := pk.ToString()

	pk = paillier.PublicKeyFromString(str)

	c1 := paillier.Encrypt(pk, 10)
	c2 := paillier.Encrypt(pk, 20)
	c3 := paillier.Encrypt(pk, 30)
	c4 := paillier.Encrypt(pk, 40)
	c5 := paillier.Encrypt(pk, 50)
	c6 := paillier.Encrypt(pk, 60)

	c7 := paillier.Add(pk,
		paillier.Add(pk, c1, c2),
		paillier.Add(pk, c3, c4),
	)

	c := paillier.BatchAdd(pk, c5, c6, c7)

	d := paillier.Mul(pk, c, 2);

	m := paillier.Decrypt(sk, d)

	fmt.Println(m)
}