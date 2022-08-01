package creational

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

const dbFileName = "capitals.txt"

// Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
// https://refactoring.guru/design-patterns/singleton

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// Feature: thread safety. Two options: sync package Once method ensures one only call. Or init() package method could also ensure it.
// Feature: Lazy loading. Create the object into memory only when it is needed, not before.
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		capitals, e := readData(dbFileName)
		db := singletonDatabase{}
		if e == nil {
			db.capitals = capitals
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

// Problems with the Singleton design pattern
// - May violate the dependency inversion principle
func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		// the line bellow makes this function depend upon a concrete database
		// this violates the Dependency Inversion Principle (DIP)
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

// The implementation bellow adheres to DIP by injecting the Database
type Database interface {
	GetPopulation(name string) int
}

func GetTotalPopulationDIP(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type FakeDatabase struct {
	fakeData map[string]int
}

var cityNames = []string{"alpha", "beta", "gama"}

func (d *FakeDatabase) GetPopulation(name string) int {
	if len(d.fakeData) == 0 {
		d.fakeData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.fakeData[name]
}
