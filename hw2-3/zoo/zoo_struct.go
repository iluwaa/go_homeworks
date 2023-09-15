package zoo

type Zoo struct {
	Cages []*Cage
}

type AnimalToCageMapping struct {
	Name        string
	CageIndex   int
	AnimalIndex int
}

type Cage struct {
	Number  int
	Animals []Animal
}

type Animal struct {
	Name            string          `json:"name"`
	Taxonomy        taxonomy        `json:"taxonomy"`
	Locations       []string        `json:"locations"`
	Characteristics characteristics `json:"characteristics"`
}

type taxonomy struct {
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phylum"`
	Class          string `json:"class"`
	Order          string `json:"order"`
	Family         string `json:"family"`
	Genus          string `json:"genus"`
	ScientificName string `json:"scientific_name"`
}

type characteristics struct {
	Prey                    string `json:"prey"`
	NameOfYoung             string `json:"name_of_young"`
	GroupBehavior           string `json:"group_behavior"`
	EstimatedPopulationSize string `json:"estimated_population_size"`
	BiggestThreat           string `json:"biggest_threat"`
	MostDistinctiveFeature  string `json:"most_distinctive_feature"`
	GestationPeriod         string `json:"gestation_period"`
	Habitat                 string `json:"habitat"`
	Diet                    string `json:"diet"`
	AverageLitterSize       string `json:"average_litter_size"`
	Lifestyle               string `json:"lifestyle"`
	CommonName              string `json:"common_name"`
	NumberOfSpecies         string `json:"number_of_species"`
	Location                string `json:"location"`
	Slogan                  string `json:"slogan"`
	Group                   string `json:"group"`
	Color                   string `json:"color"`
	SkinType                string `json:"skin_type"`
	TopSpeed                string `json:"top_speed"`
	Lifespan                string `json:"lifespan"`
	Weight                  string `json:"weight"`
	Height                  string `json:"height"`
	AgeOfSexualMaturity     string `json:"age_of_sexual_maturity"`
	AgeOfWeaning            string `json:"age_of_weaning"`
}
