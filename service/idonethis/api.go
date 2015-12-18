package idonethis

type Page struct {
	Next     string
	Results  []Done
	Status   string
	Previous string
	Count    int
	Warnings []string
}

type Done struct {
	Id              int
	Created         string
	Updated         string
	Markedup_text   string
	Done_date       string
	Owner           string
	Team_short_name string
	Tags            []string
	Likes           []string
	Comments        []string
	Meta_data       map[string]string
	Is_goal         bool
	Goal_gompleted  bool
	Url             string
	Team            string
	Raw_text        string
	Permalink       string
}
