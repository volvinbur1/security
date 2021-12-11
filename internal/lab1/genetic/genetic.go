package genetic

import (
	"github.com/volvinbur1/security/internal/lab1/decrypt"
	"math/rand"
	"sort"
	"time"
)

const ngramsCnt = 3

type Chromosome struct {
	genes   []byte
	fitness float64
}

type Algorithm struct {
	population      []Chromosome
	trigrams        map[string]float64
	generationCount int
}

func New(popSize, genCnt int, trigramsPath string) *Algorithm {
	a := &Algorithm{
		generationCount: genCnt,
	}

	var err error
	a.trigrams, err = parseTrigramsFile(trigramsPath)
	if err != nil {
		panic(err)
	}

	a.population = generatePopulation(popSize)
	return a
}

func (a *Algorithm) Decrypt(cipher []byte) (string, string) {
	a.evaluate(cipher)

	return "", ""
}

func (a *Algorithm) evaluate(cipher []byte) {
	for i := 0; i < len(a.population); i++ {
		a.population[i].fitness = a.fitness(cipher, a.population[i])
	}

	sort.Slice(a.population, func(i, j int) bool {
		return a.population[i].fitness < a.population[j].fitness
	})
}

func (a *Algorithm) fitness(cipher []byte, chromosome Chromosome) float64 {
	decryption := decrypt.Substitution(cipher, chromosome.genes)
	decryptionTrigrams := ngrams(decryption, ngramsCnt)

	fitnessValue := 0.0
	for key, val := range decryptionTrigrams {
		engTrigrams, _ := a.trigrams[key]
		fitnessValue += val - engTrigrams
	}

	return fitnessValue
}

func generatePopulation(populationSize int) []Chromosome {
	var possibleValues []byte
	for i := byte('A'); i <= byte('Z'); i++ {
		possibleValues = append(possibleValues, i)
	}

	rand.Seed(time.Now().UnixNano())
	var population []Chromosome
	for i := 0; i < populationSize; i++ {
		rand.Shuffle(len(possibleValues), func(i, j int) {
			possibleValues[i], possibleValues[j] = possibleValues[j], possibleValues[i]
		})

		population = append(population, Chromosome{genes: possibleValues})
	}

	return population
}
