package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var MutationRate = 0.005
var PopSize = 500
var Target = []byte("Recursion or Iteration? That's a question!")

// 个体
type Organism struct {
	DNA     []byte
	Fitness float64
}

// 生成个体
func createOrganism(target []byte) (organism Organism) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(95) + 32)
	}
	organism = Organism{
		DNA:     ba,
		Fitness: 0,
	}
	organism.calcFitness(target)
	return
}

// 生成族群，根据 PopSize
func createPopulation(target []byte) (population []Organism) {
	population = make([]Organism, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = createOrganism(target)
	}
	return
}

// 计算适应度，通过判断 DNA 中与 Target 相同字母的数量
func (d *Organism) calcFitness(Target []byte) {
	score := 0
	for i := 0; i < len(d.DNA); i++ {
		if d.DNA[i] == Target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.DNA))
	return
}

// 产生新的群体，通过个体的适应度（Fitness / maxFitness）来控制克隆个数，使其被选中的概率大
func createPool(population []Organism, Target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	for i := 0; i < len(population); i++ {
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// 自然选择，随机选择两个个体，然后交配，变异，产生自然选择的下一个族群
func naturalSelection(pool []Organism, population []Organism, Target []byte) []Organism {
	next := make([]Organism, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(Target)

		next[i] = child
	}

	return next
}

// 交叉，mid = rand()，mid 前半段 DNA 复制 d1，后半段 DNA 复制 d2
func crossover(d1 Organism, d2 Organism) Organism {
	child := Organism{
		DNA:     make([]byte, len(d1.DNA)),
		Fitness: 0,
	}

	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}
	}
	return child
}

// 变异，MutationRate（变异概率）
func (d *Organism) mutate() {
	for i := 0; i < len(d.DNA); i++ {
		if rand.Float64() < MutationRate {
			d.DNA[i] = byte(rand.Intn(95) + 32)
		}
	}
}

// 找到族群的最优生物
func getBest(population []Organism) Organism {
	best := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}

	return population[index]
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	population := createPopulation(Target)

	found := false
	generation := 0
	for !found {
		generation++
		bestOrganism := getBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)

		if bytes.Compare(bestOrganism.DNA, Target) == 0 {
			found = true
		} else {
			maxFitness := bestOrganism.Fitness
			pool := createPool(population, Target, maxFitness)
			population = naturalSelection(pool, population, Target)
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime taken: %s\n", elapsed)
}
