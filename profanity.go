package main

import (
	"log"
	"regexp"
	"strings"
)

/** RemoveWords: Function removes input words from the dictionary of bad words.
 Input :
		wordsToBeRemoved ([]string) the words to be removed
 Output :
		wordsToBeRemoved ([]string) words that re removed
        err (error) golang error object
**/
func RemoveWords(wordsToBeRemoved []string) ([]string, error) {
	if wordsToBeRemoved != nil {
		for _, value := range wordsToBeRemoved {
			delete(profanityWords, strings.ToLower(value))
		}
		return wordsToBeRemoved, nil
	}
	return wordsToBeRemoved, INPUT_NOT_FOUND
}

/** AddWords: Function adds given words to the dictionary of bad words.
 Input :
		words ([]string) the words to be added to the dictionary of bad words
 Output :
		words ([]string) words that have been successfully added to the dictionary
        err (error) golang error object
**/
func AddWords(words []string) ([]string, error) {
	if words != nil {
		for _, value := range words {
			if word := wordExistsInDictionary(value); word {
				profanityWords[strings.ToLower(strings.TrimSpace(value))] = 1
			}
		}
		return words, nil
	}
	return words, INPUT_NOT_FOUND
}

/** IsStringDirty: Function checks if the given string has a bad word or not.
 Input :
		words (string) the string to be checked
 Output :
		words (bool) boolean result
**/
func IsStringDirty(message string) bool {
	messageList := strings.Fields(message)
	isDirty := false
	for _, value := range messageList {
		if _, exists := profanityWords[strings.ToLower(value)]; exists {
			isDirty = true
			break
		}
	}
	return isDirty
}

/** MaskProfanity: Function masks bad words present in the input string with masking char.
 Input :
		message (string) the string to be masked
		maskWith (string) the masking character. eg *
 Output :
		words (bool) boolean result
**/
func MaskProfanity(message string, maskWith string) string {
	result := message
	for _, value := range strings.Fields(strings.ToLower(message)) {
		cleanedValue := cleanPunctuations(value)
		if exists := wordExistsInDictionary(cleanedValue); exists {
			replacement := ""
			for i := 0; i < len(cleanedValue); i++ {
				replacement = replacement + maskWith
			}
			result = strings.Replace(result, cleanedValue, replacement, -1)
		}
	}
	return result
}

func cleanPunctuations(word string) string {
	reg, err := regexp.Compile("[,:;]")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(word, "")
}

/** wordExistsInDictionary: Function checks if a given word is present in the dictionary
 of bad words.
 Input  :
		word (string) the word to be checked
 Output :
		exists (boolean) boolean response
**/
func wordExistsInDictionary(word string) bool {
	_, exists := profanityWords[word]
	return exists
}
