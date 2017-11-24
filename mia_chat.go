package mia

// Imports used
import (
    "fmt"
    "math/rand"
    "regexp"
    "strings"
    
)

// Function to retrieve reply and format input to give illusion of a response
func ReplyTo(input string) string {

    // Passes input into preprocess function
    // Function then sends back input
    input = preprocess(input)

    for pattern, responses := range phrases {
        re := regexp.MustCompile(pattern)
        matches := re.FindStringSubmatch(input)

        if len(matches) > 0 {
            var fragment string
            if len(matches) > 1 {
                fragment = reflect(matches[1])
            }

            // The output becomes a random choice of responses from the selection of responses
            // If user enter "I don't know if you are a bot" the program will search to see if it has responses for "I don't" or any other key word
            // Once a keyword is found it then randomizes a response from inside the keywords selection (Could have only 1 response, could have 10)
            output := randChoice(responses)
            
            // Merges output and input to give illusion of intellegence
            // If %s is found it 
            if strings.Contains(output, "%s") {
                output = fmt.Sprintf(output, fragment)
            }
            // Returns the output to the 'elizaInterface' function
            return output
        }
    }

    // If a response cannot be found a random conversation starter is selected
    return randChoice(conversationFail)
}

// Processes user input to become more readable by eliza
func preprocess(input string) string {
    input = strings.TrimRight(input, "\n.!")
    input = strings.ToLower(input)

    formattedInput := strings.Split(input, " ")
	for i, word := range formattedInput {
		formattedInput[i] = strings.ToLower(strings.Trim(word, ".! \n"))
    }
    
    formattedInput = PostProcess(formattedInput)

    input = strings.Join(formattedInput," ")

    return input
}

// Searches libary of words to reflect to turn around sentences to be about the user
func reflect(fragment string) string {
    words := strings.Split(fragment, " ")
    for i, word := range words {
        if Reflected, ok := Reflected[word]; ok {
            words[i] = Reflected
        }
    }
    return strings.Join(words, " ")
}

// PostProcess function loops threw the similarTo to 'dumb down' words that are similar
// This wwill make the bot look and feel more intellegent, giving it an illusion of understanding
func PostProcess(parsed []string) (input []string) {

	search := make(map[string]string)
	for key, list := range similarTo {
		for _, word := range list {
			search[word] = key
		}
	}

	input = parsed
	for i, word := range parsed {
		if processedWord, ok := search[word]; ok {
			input[i] = processedWord
		}
	}
	return input
}

// Randomizes runmber for lists in elizas scripts
func randChoice(list []string) string {
    randIndex := rand.Intn(len(list))
    return list[randIndex]
}