package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var chars = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽぁぃぅぇぉっゃゅょゎアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポァィゥェォッャュョヮ")

func _main() error {
	args := os.Args[1:]
	if len(args) < 1 {
		return errors.New("word length is not specified")
	}

	length, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("word length is not valid")
	}

	size := big.NewInt(int64(len(chars)))

	word := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, size)
		if err != nil {
			return fmt.Errorf("failed to get random number: %w", err)
		}
		word = append(word, chars[int(n.Int64())])
	}

	fmt.Println(string(word))

	return nil
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
