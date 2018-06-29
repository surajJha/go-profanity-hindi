## BMS-PROFANITY

### Author - Suraj Kumar Jha

#### Version - 1.0.0

This is a small Golang module for dealing with profanity in english and hindi. This Module has functionalities for detecting bad words and cleaning them. Technologies used : Golang

### Installation

```go get -u github.com/surajJha/go-profanity-hindi```

### Import module 

```$xslt
import profanity "github.com/surajJha/go-profanity-hindi"
```

### Dependencies 

```No dependencies```

### Code Examples

#### Mask bad words
```
        inputMessage := "Hi John, you are a bitch"  
	result := MaskProfanity(inputMessage, "*")
	log.Print(result) // Hi John, you are a *****
	
	inputMessage := "Hi John, you are a bitch"
        result := MaskProfanity(inputMessage, "#")
        log.Print(result) // Hi John, you are a #####
```  

#### Check if a string has bad words
```$xslt
    inputMessage := "Hi John, stop being a bitch"
	result := IsStringDirty(inputMessage)
	log.Print(result) // true
	
	inputMessage := "Hi John, how are you"
    	result := IsStringDirty(inputMessage)
    	log.Print(result) // false

``` 

#### Add words to dictionary
   ```$xslt
       words := []string{"asshole", "fucker"}
   	result, err := AddWords(words)
   	log.Print(result) // words that are added
   
   ```
   
#### Remove words from dictionary
```$xslt
    words := []string{"asshole", "fucker"}
	result, err := RemoveWords(words)
	log.Print(result) // words that are removed

```   

### Run Unit Tests

```$xslt
Run command - go test -v
```

### License

``MIT License``