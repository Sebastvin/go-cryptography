package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputLength = 8
	keyLength   = 10
)

var (
	p10 = []uint8{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	p8  = []uint8{6, 3, 7, 4, 8, 5, 10, 9}
	ip  = []uint8{2, 6, 3, 1, 4, 8, 5, 7}
	ip1 = []uint8{4, 1, 3, 5, 7, 2, 8, 6}
	ep  = []uint8{4, 1, 2, 3, 2, 3, 4, 1}
	p4  = []uint8{2, 4, 3, 1}
	ls1 = []uint8{2, 3, 4, 5, 1}
	ls2 = []uint8{3, 4, 5, 1, 2}

	s0 = [][]uint8{
		{1, 0, 3, 2},
		{3, 2, 1, 0},
		{0, 2, 1, 3},
		{3, 1, 3, 2},
	}
	s1 = [][]uint8{
		{0, 1, 2, 3},
		{2, 0, 1, 3},
		{3, 0, 1, 0},
		{2, 1, 0, 3},
	}
)

type sdes struct {
	input   []uint8
	key     []uint8
	decrypt bool
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go [-decrypt] <input> <key>")
	}

	d, err := parseArgs(os.Args)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	result, err := d.process()
	if err != nil {
		log.Fatalf("Error processing S-DES: %v", err)
	}

	for _, v := range result {
		fmt.Print(v)
	}
	fmt.Println()
}

func parseArgs(args []string) (*sdes, error) {
	if len(args) < 3 {
		return nil, errors.New("insufficient arguments")
	}

	var input, key string
	decrypt := false

	if args[1] == "-decrypt" {
		if len(args) != 4 {
			return nil, errors.New("usage: go run main.go -decrypt <input> <key>")
		}
		decrypt = true
		input = args[2]
		key = args[3]
	} else {
		if len(args) != 3 {
			return nil, errors.New("usage: go run main.go <input> <key>")
		}
		input = args[1]
		key = args[2]
	}

	inputBytes, keyBytes, err := formatArguments(input, key)
	if err != nil {
		return nil, err
	}

	return &sdes{input: inputBytes, key: keyBytes, decrypt: decrypt}, nil
}

func formatArguments(input, key string) ([]uint8, []uint8, error) {
	if len(input) != inputLength {
		return nil, nil, fmt.Errorf("invalid input length: got %d, want %d", len(input), inputLength)
	}
	if len(key) != keyLength {
		return nil, nil, fmt.Errorf("invalid key length: got %d, want %d", len(key), keyLength)
	}

	inputBytes := make([]uint8, inputLength)
	keyBytes := make([]uint8, keyLength)

	for i := 0; i < len(input); i++ {
		val, err := strconv.Atoi(string(input[i]))
		if err != nil || val > 1 {
			return nil, nil, fmt.Errorf("invalid binary digit at input[%d]: %c", i, input[i])
		}
		inputBytes[i] = uint8(val)
	}

	for i := 0; i < len(key); i++ {
		val, err := strconv.Atoi(string(key[i]))
		if err != nil || val > 1 {
			return nil, nil, fmt.Errorf("invalid binary digit at key[%d]: %c", i, key[i])
		}
		keyBytes[i] = uint8(val)
	}

	return inputBytes, keyBytes, nil
}

func (d *sdes) process() ([]uint8, error) {
	k1, k2 := generateKeys(d.key)
	if d.decrypt {
		k1, k2 = k2, k1
	}

	output := permute(d.input, ip)
	left, right := split(output)

	f1 := fFunction(left, right, k1)
	left, right = split(f1)

	left, right = right, left
	output = fFunction(left, right, k2)

	return permute(output, ip1), nil
}

func split(input []uint8) ([]uint8, []uint8) {
	mid := len(input) / 2
	return input[:mid], input[mid:]
}

func generateKeys(key []uint8) ([]uint8, []uint8) {
	permuted := permute(key, p10)
	left, right := split(permuted)

	left = permute(left, ls1)
	right = permute(right, ls1)
	k1 := permute(append(left, right...), p8)

	left = permute(left, ls2)
	right = permute(right, ls2)
	k2 := permute(append(left, right...), p8)

	return k1, k2
}

func fFunction(left, right []uint8, key []uint8) []uint8 {
	expanded := permute(right, ep)
	xored := xor(expanded, key)

	s0out := sBox(xored[:4], s0)
	s1out := sBox(xored[4:], s1)

	combined := append(s0out, s1out...)
	permuted := permute(combined, p4)

	return append(xor(left, permuted), right...)
}

func permute(input []uint8, positions []uint8) []uint8 {
	result := make([]uint8, len(positions))
	for i, pos := range positions {
		if int(pos-1) >= len(input) {
			log.Fatalf("permutation out of bounds: position %d, input length %d", pos, len(input))
		}
		result[i] = input[pos-1]
	}
	return result
}

func xor(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		log.Fatalf("xor: length mismatch: %d != %d", len(a), len(b))
	}
	result := make([]uint8, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func sBox(input []uint8, sMatrix [][]uint8) []uint8 {
	if len(input) != 4 {
		log.Fatalf("sBox: invalid input length: got %d, want 4", len(input))
	}
	row := input[0]<<1 | input[3]
	col := input[1]<<1 | input[2]

	val := sMatrix[row][col]
	return []uint8{val / 2, val % 2}
}
