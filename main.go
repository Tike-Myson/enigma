package main

import "fmt"

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
	fmt.Println(answer)
	answer = RotorDecrypt(key, answer)
	fmt.Println(answer)
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