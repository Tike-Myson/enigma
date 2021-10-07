package main

import (
	"fmt"
	"time"
)

type Pair struct {
	First string
	Second string
}

func main() {
	var key string
	fmt.Println("Enter enigma key")
	fmt.Scanf("%s\n", &key)
	var message string
	fmt.Println("Enter your message")
	fmt.Scanf("%s\n", &message)
	message = Sanitize(message)
	fmt.Printf("Total symbols: %d\n", len([]byte(message)))
	fmt.Println("Analysis: ")
	fmt.Printf("Before: %f\n", FrequencyAnalysis(message))
	//fmt.Println("***********************")
	start := time.Now();
	//commutation
	pairs := make(map[string]string)
	pairs["A"] = "B"
	pairs["C"] = "Q"
	pairs["D"] = "Z"
	pairs["E"] = "L"
	pairs["F"] = "X"
	pairs["J"] = "P"
	pairs["H"] = "U"
	pairs["I"] = "Y"
	pairs["G"] = "N"
	pairs["K"] = "T"


	//answer := Switch(pairs, message)
	answer := Rotor(key, message)
	fmt.Printf("Before: %f\n", FrequencyAnalysis(answer))
	end := time.Now()
	timer := end.Sub(start)
	fmt.Println(timer)
	//answer = RotorDecrypt(key, answer)
	//fmt.Println(answer)
	//answer = Switch(pairs, answer)

}

func RotorDecrypt(key, message string) string {
	key = Reverse(key)
	var encryptedMessage string

	var rotorValue []int
	for _, v := range []byte(key) {
		rotorValue = append(rotorValue, int(v) - 64)
	}
	//rotor1
	for _, char := range []byte(message) {
		if rotorValue[2] > 26 {
			rotorValue[1]++
			rotorValue[2] = 1
		}
		n := SwitchAlphabetReverse(char, rotorValue[2])
		encryptedMessage += string(n)
		rotorValue[2]++
	}

	message = encryptedMessage
	encryptedMessage = ""
	//rotor2
	for _, char := range []byte(message) {
		if rotorValue[1] > 26 {
			rotorValue[0]++
			rotorValue[1] = 1
		}
		n := SwitchAlphabetReverse(char, rotorValue[1])
		encryptedMessage += string(n)
		rotorValue[1]++
	}
	message = encryptedMessage
	encryptedMessage = ""
	//rotor3
	for _, char := range []byte(message) {
		if rotorValue[0] > 26 {
			rotorValue[0] = 1
		}
		n := SwitchAlphabetReverse(char, rotorValue[0])
		encryptedMessage += string(n)
		rotorValue[0]++
	}
	return encryptedMessage

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Rotor(key, message string) string {
	var encryptedMessage string

	var rotorValue []int
	for _, v := range []byte(key) {
		rotorValue = append(rotorValue, int(v) - 64)
	}
	//rotor1
	for _, char := range []byte(message) {
		if rotorValue[2] > 26 {
			rotorValue[1]++
			rotorValue[2] = 1
		}
		n := SwitchAlphabet(char, rotorValue[2])
		encryptedMessage += string(n)
		rotorValue[2]++
	}

	message = encryptedMessage
	encryptedMessage = ""
	//rotor2
	for _, char := range []byte(message) {
		if rotorValue[1] > 26 {
			rotorValue[0]++
			rotorValue[1] = 1
		}
		n := SwitchAlphabet(char, rotorValue[1])
		encryptedMessage += string(n)
		rotorValue[1]++
	}
	message = encryptedMessage
	encryptedMessage = ""
	//rotor3
	for _, char := range []byte(message) {
		if rotorValue[0] > 26 {
			rotorValue[0] = 1
		}
		n := SwitchAlphabet(char, rotorValue[0])
		encryptedMessage += string(n)
		rotorValue[0]++
	}
	return encryptedMessage
}

func Switch(pairs map[string]string, message string) string {
	var answer string
	for _, char := range []byte(message) {
		if val, ok := pairs[string(char)]; ok {
			answer += val
			continue
		}
		ok, key := IsValueExist(pairs, string(char))
		if ok {
			answer += key
			continue
		}
		answer += string(char)
	}
	return answer
}

func IsValueExist(pairs map[string]string, char string) (bool, string) {
	for key, value := range pairs {
		if value == char {
			return true, key
		}
	}
	return false, ""
}

func SwitchAlphabet(char byte, position int) byte {
	if int(char) + position > 90 {
		return byte(int(char) + position - 90 + 65)
	}
	return byte(int(char) + position)
}

func SwitchAlphabetReverse(char byte, position int) byte {
	if int(char) - position < 65 {
		return byte(int(char) - position + 25)
	}
	return byte(int(char) - position)
}

func FrequencyAnalysis(msg string) float64 {

	result := 0.0

	FrequencyTable := map[string]float64{
		"E": 12.02,
		"T": 9.10,
		"A": 8.12,
		"O": 7.68,
		"I": 7.31,
		"N": 6.95,
		"S": 6.28,
		"R": 6.02,
		"H": 5.92,
		"D": 4.32,
		"L": 3.98,
		"U": 2.88,
		"C": 2.71,
		"M": 2.61,
		"F": 2.30,
		"Y": 2.11,
		"W": 2.09,
		"G": 2.03,
		"P": 1.82,
		"B": 1.49,
		"V": 1.11,
		"K": 0.69,
		"X": 0.17,
		"Q": 0.11,
		"J": 0.10,
		"Z": 0.07,
	}

	MessageFrequency := make(map[string]float64)

	EnglishMessage := make(map[string]int)
	//EnglishOriginal := make(map[string]int)
	Alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//fill map with 0 values
	for _, v := range Alphabet {
		EnglishMessage[string(v)] = 0
		//EnglishOriginal[string(v)] = 0
	}

	//increment map values
	for _, v := range msg {
		EnglishMessage[string(v)] += 1
	}

	lenMsg := float64(len(msg))

	//fill MessageFrequency
	for k, v := range EnglishMessage {
		MessageFrequency[k] = float64(v) * 100.0/lenMsg
	}

	//fmt.Println(FrequencyTable)
	//fmt.Println(MessageFrequency)

	for k, v := range MessageFrequency {
		if v >= FrequencyTable[k] - 1.1 && v <= FrequencyTable[k] + 1.1 {
			result += 3.846
		}
	}
	return result
}

func Sanitize(in string) string {
	var out []rune
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}
