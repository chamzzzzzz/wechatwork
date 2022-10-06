package user

type User struct {
	Userid      string `json:"userid"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Alias       string `json:"alias"`
	Gender      string `json:"gender"`
	Mobile      string `json:"mobile"`
	Department  []int  `json:"department"`
	Position    string `json:"position"`
	OpenUserid  string `json:"open_userid"`
	Status      int    `json:"status"`
}
