package genetic

import (
	"fmt"
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

func (a *Algorithm) Decrypt(cipher []byte) ([]byte, string) {
	a.evaluate(cipher)

	for i := 0; i < a.generationCount; i++ {
		a.newGeneration()
		a.evaluate(cipher)
		fmt.Printf("%d -> %s\t%f\n", i, a.population[0].genes, a.population[0].fitness)
	}

	return decrypt.Substitution(cipher, a.population[0].genes), string(a.population[0].genes)
}

func (a *Algorithm) evaluate(cipher []byte) {
	for i := 0; i < len(a.population); i++ {
		a.population[i].fitness = a.fitness(cipher, a.population[i])
	}

	sort.Slice(a.population, func(i, j int) bool {
		return a.population[i].fitness < a.population[j].fitness
	})
}

func (a Algorithm) fitness(cipher []byte, chromosome Chromosome) float64 {
	decryption := decrypt.Substitution(cipher, chromosome.genes)
	decryptionTrigrams := ngrams(decryption, ngramsCnt)

	fitnessValue := 0.0
	for key, val := range decryptionTrigrams {
		engTrigrams, _ := a.trigrams[key]
		fitnessValue += val - engTrigrams
	}

	return fitnessValue
}

func (a *Algorithm) newGeneration() {
	newPopulation := make([]Chromosome, 0)
	newPopulation = append(newPopulation, a.population[:len(a.population)/5]...)

	//rand.Seed(time.Now().UnixNano())
	for i := len(a.population) / 5; i < len(a.population); i += 2 {
		firstParent := a.selectChromosome()
		secondParent := a.selectChromosome()

		newPopulation = a.crossover(firstParent, secondParent, newPopulation)
	}

	a.population = newPopulation
}

func (a Algorithm) crossover(firstParent, secondParent Chromosome, newPopulation []Chromosome) []Chromosome {
	genesCnt := int(byte('Z') - byte('A') + 1)
	end := rand.Intn(genesCnt) + 1
	start := rand.Intn(end)

	firstChild := createEmptyChromosome()
	secondChild := createEmptyChromosome()
	for i := start; i < end; i++ {
		firstChild.genes[i] = secondParent.genes[i]
		secondChild.genes[i] = firstParent.genes[i]
	}

	for i := 0; i < genesCnt; i++ {
		if i == start {
			i += end - start
			if i == genesCnt {
				continue
			}
		}

		idx := i
		if contains(firstChild.genes, firstParent.genes[idx]) {
			idx = indexOf(firstChild.genes, firstParent.genes[idx])
		}
		firstChild.genes[i] = firstParent.genes[idx]

		idx = i
		if contains(secondChild.genes, secondParent.genes[idx]) {
			idx = indexOf(secondChild.genes, secondParent.genes[idx])
		}
		secondChild.genes[i] = secondParent.genes[idx]
	}

	return append(newPopulation, firstChild, secondChild)
}

func (a Algorithm) selectChromosome() Chromosome {
	total := 0.0
	for _, chromosome := range a.population {
		total += chromosome.fitness
	}

	limitation := rand.Float64() * total

	total = 0.0
	for _, chromosome := range a.population {
		total += chromosome.fitness
		if total >= limitation {
			return chromosome
		}
	}

	return a.population[0]
}

func contains(arr []byte, value byte) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}

	return false
}

func indexOf(arr []byte, value byte) int {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return i
		}
	}

	return -1
}

func createEmptyChromosome() Chromosome {
	chromosome := Chromosome{
		genes: make([]byte, byte('Z')-byte('A')+1),
	}

	for i := 0; i < len(chromosome.genes); i++ {
		chromosome.genes[i] = '-'
	}

	return chromosome
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

		newChromosome := Chromosome{genes: make([]byte, len(possibleValues))}
		copy(newChromosome.genes, possibleValues)
		population = append(population, newChromosome)
	}

	return population
}
