package rake

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/neurosnap/sentences.v1"
)

type phrase struct {
	words []string
}

type phraseScore struct {
	phrase string
	score  float64
}

var splitter *regexp.Regexp = regexp.MustCompile("\\s")

type Rake struct {
	stopwords         map[string]bool
	sentenceTokenizer *sentences.DefaultSentenceTokenizer
}

func New(stopwords []string) (Rake, error) {
	o := Rake{}
	o.stopwords = make(map[string]bool)
	for i := 0; i < len(stopwords); i++ {
		o.stopwords[stopwords[i]] = true
	}

	stok, err := NewSentenceTokenizer()
	if err != nil {
		return o, err
	}
	o.sentenceTokenizer = stok

	return o, nil
}

func (o Rake) Extract(text string) ([]string, error) {
	t1 := time.Now()
	text = strings.ToLower(text)
	splitted := splitter.Split(text, -1)
	text = strings.Join(splitted, " ")

	sent1 := o.sentenceTokenizer.Tokenize(text)
	var candidates []phrase
	for i := 0; i < len(sent1); i++ {
		candidates = append(candidates, o.splitSentence(sent1[i].Text))
	}

	scores := wordscores(candidates)
	pscores := make([]phraseScore, len(candidates), len(candidates))
	for i := 0; i < len(candidates); i++ {
		phraseVal := strings.Join(candidates[i].words, " ")
		val := phrasescore(candidates[i], scores)
		p := phraseScore{phraseVal, val}
		pscores[i] = p
	}

	sort.SliceStable(pscores, func(i, j int) bool {
		return pscores[i].score > pscores[j].score
	})
	picknum := len(pscores) / 3
	log.Println(picknum)
	result := make([]string, picknum, picknum)
	for i := 0; i < picknum; i++ {
		result[i] = pscores[i].phrase
	}

	elapsed := time.Now().Sub(t1)
	log.Println("Finished ", elapsed)
	return result, nil
}

func (o Rake) splitSentence(sentence string) phrase {
	p := phrase{words: make([]string, 0, 0)}
	replacer := strings.NewReplacer(",", " ", ".", " ", ";", " ", "!", " ", "?", " ", ":", " ", "\\", " ", "/", " ", "=", " ", "´", " ", "'", " ", "\"", " ", "”", " ", "…", " ",
		"§", " ", "[", " ", "(", " ", "]", " ", ")", " ", "€", " ", "_", " ", "<", " ", ">", " ", "&", " ", "*", " ", "-", " ", "#", " ", "@", " ", "%", " ", "„", " ", "“", " ",
		"+", " ")
	sentence = replacer.Replace(sentence)
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if _, ok := o.stopwords[words[i]]; !ok && len(words[i]) > 3 && !isNumeric(words[i]) && !startsWithNumber(words[i]) {
			p.words = append(p.words, words[i])
		}
	}
	return p
}

func wordscores(candidates []phrase) map[string]float64 {
	deg := make(map[string]int)
	freq := make(map[string]int)
	for i := 0; i < len(candidates); i++ {
		phraseLen := len(candidates[i].words)
		for j := 0; j < phraseLen; j++ {
			freq[candidates[i].words[j]]++
			deg[candidates[i].words[j]] += phraseLen
		}
	}
	scores := make(map[string]float64)
	for word, f := range freq {
		scores[word] = float64(deg[word]) / float64(f)
	}
	return scores
}

func phrasescore(p phrase, scores map[string]float64) float64 {
	var score float64
	for i := 0; i < len(p.words); i++ {
		ws, _ := scores[p.words[i]]
		score += ws
	}
	return score

}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func startsWithNumber(s string) bool {
	var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "9"}
	for i := 0; i < len(digits); i++ {
		if strings.HasPrefix(s, digits[i]) {
			return true
		}
	}
	return false
}
