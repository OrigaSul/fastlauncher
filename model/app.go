package model

type App struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Command     string `json:"command"`
	Keywords    string `json:"keywords"`
}
