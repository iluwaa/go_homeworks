package zoo

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"time"
)

const animalsString = "Dog,Cow,Cat,Horse,Donkey,Tiger,Lion,Panther,Leopard,Cheetah,Bear,Elephant,Bear,Turtle,Tortoise,Crocodile,Rabbit,Porcupine,Hare,Hen,Pigeon,Albatross,Crow,Fish,Dolphin,Frog,Whale,Alligator,Eagle,Squirrel,Ostrich,Fox,Goat,Jackal,Emu,Armadillo,Eel,Goose,Fox,Wolf,Beagle,Gorilla,Chimpanzee,Monkey,Beaver,Orangutan,Antelope,Bat,Badger,Giraffe,Hermit,Crab,Giant,Panda,Hamster,Cobra,Shark,Camel,Hawk,Deer,Chameleon,Hippopotamus,Jaguar,Chihuahua,King,Cobra,Ibex,Lizard,Koala,Kangaroo,Iguana,Llama,Chinchillas,Dodo,Jellyfish,Rhinoceros,Hedgehog,Zebra,Possum,Wombat,Bison,Bull,Buffalo,Sheep,Meerkat,Mouse,Otter,Sloth,Owl,Vulture,Flamingo,Racoon,Mole,Duck,Swan,Lynx,Monitor,lizard,Elk,Boar,Lemur,Mule,Baboon,Mammoth,Rat,Snake,Peacock"

func MakeZoo(size int) *Zoo {
	animalCount := len(strings.Split(animalsString, ","))
	if size > animalCount {
		fmt.Printf("I know only %[1]d pets, so generate a zoo with %[2]d animals is not possible. Please specify size of zoo less than %[2]d", animalCount, size)
		os.Exit(1)
	}
	var Zoo = new(Zoo)
	animalsList := generateAnimalsList(size)

	for index, animal := range animalsList {
		animalsList := getAnimalsInfo(animal)

		Zoo.Cages = append(Zoo.Cages, settleAnimals(index+1, animalsList))
	}

	return Zoo
}

func settleAnimals(number int, animals []Animal) *Cage {
	result := new(Cage)

	result.Number = number
	result.Animals = animals

	return result
}

func generateAnimalsList(count int) []string {
	var slicedAnimalsList []string
	completeAnimalsList := strings.Split(animalsString, ",")

	for i := 0; i < count; i++ {
		slicedAnimalsList = append(slicedAnimalsList, completeAnimalsList[rand.Intn(len(completeAnimalsList))])
	}

	return slicedAnimalsList
}

func getAnimalsInfo(animal string) []Animal {
	var result []Animal

	baseUrl, err := url.Parse("https://api.api-ninjas.com/v1/animals")
	if err != nil {
		fmt.Println(err)

	}
	urlParams := url.Values{}
	urlParams.Add("name", animal)
	baseUrl.RawQuery = urlParams.Encode()

	request, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	// I know that api key it is sensitive info, but do not matter
	request.Header.Add("X-Api-Key", "bvRGcHgwcwOLsonFMgrt0g==0ZNpc9rceuXSsBnw")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func RandomAnimalsList(zoo *Zoo, count int) []AnimalToCageMapping {
	var completeAnimalsList []AnimalToCageMapping
	var slicedAnimalsList []AnimalToCageMapping

	for i, cage := range zoo.Cages {
		for k, animal := range cage.Animals {
			tmp := new(AnimalToCageMapping)
			tmp.Name = animal.Name
			tmp.CageIndex = i
			tmp.AnimalIndex = k

			completeAnimalsList = append(completeAnimalsList, *tmp)
		}
	}

	for len(slicedAnimalsList) < count {
		randomIndex := rand.Intn(len(completeAnimalsList))
		tmpAnimalMapping := completeAnimalsList[randomIndex]
		if !slices.Contains(slicedAnimalsList, tmpAnimalMapping) {
			slicedAnimalsList = append(slicedAnimalsList, tmpAnimalMapping)
		}
	}

	return slicedAnimalsList
}

// Debug
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
