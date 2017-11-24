package mia

// list of phrases
var phrases = map[string][]string{

    // Words
    `hello`: {
        "Hey how are you?",
        "Any craic?",
    },

    `how are you?`: {
        "I'm fine how are you?",
    },

    `what`: {
        "You know what!",
    },

    `how are you feeling`: {
        "I am feeling mighty",
    },
    

    // Responses
    `no`:{
        "Don't be so negative",
    },

    `yes`: {
        "Brilliant, glad to hear!",
    },

    `you (.*)`: {
        "Less about me, more about you, please!",
    },

    `sorry`: {
        "Don't be sorry",
    },
    
    `it is (.*)`: {
        "It sure is %s !",
    },
    
    

    // general

    `I need a (.*)`: {
    "Why do you need a %s?",
    "Would it really help you to get a %s?",
    "Are you sure you need a %s?",
    },

    `any (.*)`: {
        "I'm afraid not, sorry",
    },
    `can you (.*)`: {
        "I cannot %s right now, but maybe I could If you were to teach me",
    },
    `i am (.*)`: {
        "Congrats on being %s ",
    },
    `I think (.*)`: {
        "%s that's very intresting, hmm..",
    },
    `i feel (.*)`: {
        "I sometimes feel like that too",
    },
    `i have (.*)`: {
        "%s is always handy to have",
    },
    `i would (.*)`: {
        "Why would you %s?",
    },
    `i don't (.*)`: {
        "Why don't you %s?",
    },

    // questions
    `why don't you (.*)`: {
        "If I %s will you never ask me again?",
        "Do you really think I don't %s?",
        "Perhaps eventually I will %s.",
        "Do you really want me to %s?",
    },

     `what colour is the sky?` : {
         "Blue!",
     },

    `what is the time?`:{
        "Check your phone buddy",
     
    },
    `are you (.*)`: {
        "I could be %s, I never checked",
    },
    `how do I (.*)`: {
        "I did know I was going to be quized on %s today!",
    },
    `is it (.*)`: {
        "It sure is %s?",
    },
    
    `can I (.*)`: {
        "Can you %s ? .. sure why not!",
        
    },
    `is there (.*)`: {
        "I sure hope there is!!",
    },
    `why (.*)`: {
        "You tell me why?!",
    },
    `what (.*)`: {
        "I find google could answer this question better!",
    },
    `you are (.*)`: {
        "%s .. well thats just great!",
    },
    `(.*)\?`: {
        "hmm",
        "what to try that one again?",
    },
    `how does the thing go?`:{
        "The Ting Goes Skrrrrrrrrrah Pap Pap Kak Kak Kak",
    },
    }

// If inputed phrase/word does not match an output:
var conversationFail = []string{
    "Sorry could you repeat that?",
}

// flip convo back to user if needed
var Reflected = map[string]string{
    "am":       "are",
    "was":      "were",
    "I":        "you",
    "I'd":      "you would",
    "I've":     "you have",
    "I'll":     "you will",
    "my":       "your",
    "are":      "am",
    "you've":   "I have",
    "you'll":   "I will",
    "you're":   "I'm",
    "your":     "my",
    "yours":    "mine",
    "you":      "I",
    "me":       "you",
}

//if mia is asked questions with similar expressions mia should understand
var similarTo = map[string][]string{

    "hello":    []string{"hey","hi","yo","howya","how are you?"},
    "thanks":   []string{"cheers","sound", "thank you"},
    "want":     []string{"need"},
    "you are":  []string{"you're"},
    "sorry":    []string{"apologies","soz"},
    "i am":     []string{"i'm"},

}