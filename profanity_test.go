package profanity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanPunctuations(t *testing.T) {
	assert := assert.New(t)
	inputRegex := "abcd"
	inputString := "suraj is, a human"
	replacement := ""
	cleanedString := cleanPunctuations(inputString, inputRegex, replacement)
	assert.NotEqual("suraj is a human", cleanedString)
}

func TestMaskProfanity(t *testing.T) {
	assert := assert.New(t)

	inputMessage := "Hi John, you are a bitch"
	result := MaskProfanity(inputMessage, "*")
	assert.NotNil(result)
	assert.Equal("Hi John, you are a *****", result, "bad words will be replaced with *")
	assert.NotEqual("Hi John, you are a ****", result)
	assert.NotEqual("#####", result)

	inputMessage = "benchod hai tu, maadarjaat"
	result = MaskProfanity(inputMessage, "#")
	assert.NotNil(result)
	assert.Equal("####### hai tu, ##########", result)

	inputMessage = " benchod, how are you"
	result = MaskProfanity(inputMessage, "*")
	assert.NotNil(result)
	assert.Equal(" *******, how are you", result)
}

func TestIsStringDirty(t *testing.T) {
	assert := assert.New(t)

	// Testing for correct input
	inputMessage := "Hi John, stop being a bitch"
	result := IsStringDirty(inputMessage)
	assert.NotNil(result, "Result should not be nil. It will be either true or false")
	assert.True(result, "Should be true as there is a bad word in the string, bitch")

	inputMessage = "hi John, good morning to you"
	result = IsStringDirty(inputMessage)
	assert.NotNil(result, "Result should not be nil. It will be either true or false")
	assert.False(result, "Should be false as there is no bad wor in the string")
}

func TestWordExistsInDictionary(t *testing.T) {
	assert := assert.New(t)

	// Testing for correct input
	word := "boobs"
	exists := wordExistsInDictionary(word)
	assert.Equal(true, exists)

	// Testing for one missing input
	emptyWord := ""
	exists = wordExistsInDictionary(emptyWord)
	assert.NotNil(exists)

	// Testing for second input missing
	word = "kutte"
	exists = wordExistsInDictionary(word)
	assert.NotNil(exists)
}

func TestAddWords(t *testing.T) {
	assert := assert.New(t)

	// Testing for correct input
	words := []string{"asshole", "fucker"}
	result, err := AddWords(words)
	assert.NotNil(result)
	assert.Nil(err)
	assert.IsType([]string{}, result)
	assert.Len(result, 2)

	// Testing for incorrect input
	words = nil
	result, err = AddWords(words)
	assert.Nil(result)
	assert.NotNil(err)
	assert.IsType([]string{}, result)
	assert.Error(err)
	assert.EqualError(err, "input is missing")
}

func TestRemoveWords(t *testing.T) {
	assert := assert.New(t)

	// Testing for correct input
	wordsToBeRemoved := []string{"boobs", "kutte"}
	result, err := RemoveWords(wordsToBeRemoved)
	assert.NotNil(result)
	assert.Nil(err)

	// Testing for incorrect input
	wordsToBeRemoved = nil
	result, err = RemoveWords(wordsToBeRemoved)
	assert.NotNil(err)
	assert.Nil(result)
	assert.EqualError(err, "input is missing")

}
