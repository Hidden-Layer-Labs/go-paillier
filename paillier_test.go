package paillier

import (
	"math/big"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		input  []int64
		output int64
	}{
		{[]int64{10, 20, 30}, 60},
		{[]int64{25, 50, 75, 100}, 250},
		{[]int64{1}, 1},
	}

	pk, sk := GenerateKeypair(1024).ToKeys()

	for _, table := range tables {
		sum := Encrypt(pk, 0)
		for _, input := range table.input {
			sum = Add(pk, sum, Encrypt(pk, input))
		}

		out := Decrypt(sk, sum)

		if out != table.output {
			t.Errorf("Sum of %v was incorrect, got: %v, want: %v", table.input, out, table.output)
		}
	}

}

func TestMul(t *testing.T) {
	tables := []struct {
		base       int64
		multiplier int64
		output     int64
	}{
		{0, 2, 0},
		{25, 5, 125},
		{100, 0, 0},
		{6, 75, 450},
	}

	pk, sk := GenerateKeypair(1024).ToKeys()

	for _, table := range tables {
		this_base := Encrypt(pk, table.base)
		this_mul := Mul(pk, this_base, table.multiplier)
		this_out := Decrypt(sk, this_mul)

		if this_out != table.output {
			t.Errorf("Mul of base %v and multiplier %v incorrect, got: %v, want: %v", table.base, table.multiplier, this_out, table.output)
		}
	}
}

func TestBatchSum(t *testing.T) {
	tables := []struct {
		input  []int64
		output int64
	}{
		{[]int64{10, 20, 30}, 60},
		{[]int64{25, 50, 75, 100}, 250},
		{[]int64{1}, 1},
	}

	pk, sk := GenerateKeypair(1024).ToKeys()

	for _, table := range tables {
		this_input := make([]*big.Int, len(table.input))
		for i, input := range table.input {
			this_input[i] = Encrypt(pk, input)
		}

		this_sum := BatchAdd(pk, this_input...)

		out := Decrypt(sk, this_sum)

		if out != table.output {
			t.Errorf("Batch sum of %v was incorrect, got: %v, want: %v", table.input, out, table.output)
		}
	}
}

func TestPublicKeyStrings(t *testing.T) {
	for i := 0; i < 4; i++ {
		pk, _ := GenerateKeypair(1024).ToKeys()

		pk_str := pk.String()
		pk_back := PublicKeyFromString(pk_str)

		if pk.n.Cmp(pk_back.n) != 0 || pk.g.Cmp(pk_back.g) != 0 || pk.n2.Cmp(pk_back.n2) != 0 {
			t.Errorf("public key conversion for string failed!")
		}
	}
}

func TestKeypairString(t *testing.T) {
	for i := 0; i < 4; i++ {
		kp := GenerateKeypair(1024)

		kp_str := kp.String()
		kp_back := KeypairFromString(kp_str)

		if kp.p.Cmp(kp_back.p) != 0 || kp.q.Cmp(kp_back.q) != 0 {
			t.Errorf("public key conversion for string failed!")
		}
	}
}
