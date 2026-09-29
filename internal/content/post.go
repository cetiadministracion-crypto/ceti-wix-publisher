package content

type Post struct {
	Title    string `yaml:"title"`
	Slug     string `yaml:"slug"`
	Excerpt  string `yaml:"excerpt"`
	Language string `yaml:"language"`
	Article  string `yaml:"article"`
	Hero     string `yaml:"hero"`
	AltText  string `yaml:"alt_text"`
}
