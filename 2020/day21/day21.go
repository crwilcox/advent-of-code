package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func readFileToLines(path string) ([]string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Item struct
type Item struct {
	ingredients []string
	allergens   []string
}

func parseLine(line string) Item {
	//mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
	line = strings.TrimSuffix(line, ")")
	split := strings.Split(line, " (contains ")
	ingredients := strings.Split(split[0], " ")
	allergens := strings.Split(split[1], ", ")
	return Item{ingredients, allergens}
}

// Part 1:
func countNonAllergenAppearances(items []Item) int {
	ingredientToAllergen := identifyAllergens(items)

	// once we have a list of allergens, we need to find the other items
	// and also find the counts in our list.
	countNonAllergens := 0
	for _, item := range items {
		for _, ingredient := range item.ingredients {
			if _, ok := ingredientToAllergen[ingredient]; !ok {
				countNonAllergens++
			}
		}
	}
	return countNonAllergens
}

// Part 2:
// Identifies allergesn
// Returns map of ingredient to allergen
func identifyAllergens(items []Item) map[string]string {
	allergenToPotential := make(map[string][]string)

	for _, item := range items {
		for _, allergen := range item.allergens {
			if _, ok := allergenToPotential[allergen]; ok {
				allergenToPotential[allergen] = append(allergenToPotential[allergen], item.ingredients...)
			} else {
				allergenToPotential[allergen] = item.ingredients
			}
		}
	}

	// we now have a map with ingredients that could be the allergen. the highest
	// count (ties both potential contain it) is the allergen match.
	// from there we can find which items don't contain allergens
	allergens := make(map[string]string)
	for len(allergens) < len(allergenToPotential) {
		for allergen, ingredients := range allergenToPotential {
			tally := make(map[string]int)
			for _, ingredient := range ingredients {
				if _, ok := tally[ingredient]; !ok {
					tally[ingredient] = 0
				}
				tally[ingredient]++
			}

			highestValue := 0
			for _, i := range tally {
				if highestValue < i {
					highestValue = i
				}
			}
			possibleIngredientsWithAllergen := []string{}
			for ingredient, count := range tally {
				_, detectedAlergen := allergens[ingredient]
				if count == highestValue && !detectedAlergen {
					possibleIngredientsWithAllergen = append(possibleIngredientsWithAllergen, ingredient)
				}
			}

			if len(possibleIngredientsWithAllergen) == 1 {
				// found the allergen match.
				allergens[possibleIngredientsWithAllergen[0]] = allergen
			}
		}
	}

	return allergens
}

func formatAllergens(ingredientToAllergen map[string]string) string {
	allergenToIngredient := make(map[string]string)
	allergens := []string{}
	for k, v := range ingredientToAllergen {
		allergenToIngredient[v] = k
		allergens = append(allergens, v)
	}

	ingredients := []string{}
	sort.Strings(allergens)
	for _, allergen := range allergens {
		ingredients = append(ingredients, allergenToIngredient[allergen])
	}
	return strings.Join(ingredients, ",")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := readFileToLines(filePath)
	if err != nil {
		panic(err)
	}
	items := []Item{}
	for _, line := range lines {
		item := parseLine(line)
		items = append(items, item)
	}

	count := countNonAllergenAppearances(items)
	fmt.Println("🎄 Part 1 🎁:", count) // Answer: 2724

	allergens := identifyAllergens(items)
	part2 := formatAllergens(allergens)
	fmt.Println("🎄 Part 2 🎁:", part2) // Answer: xlxknk,cskbmx,cjdmk,bmhn,jrmr,tzxcmr,fmgxh,fxzh
}
