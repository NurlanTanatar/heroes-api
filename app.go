package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type Images struct {
	XS string
	SM string
	MD string
	LD string
}

type Hero struct {
	Id     int
	Name   string
	Images Images
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetRandomImageUrl() string {
	Id := rand.Intn(731) + 1
	url := fmt.Sprintf("%s%d%s", "https://cdn.jsdelivr.net/gh/akabab/superhero-api@0.3.0/api/id/", Id, ".json")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Hero
	json.Unmarshal(responseData, &data)

	return data.Images.SM
}

func (a *App) GetHeroList() []string {
	var heroes []string

	response, err := http.Get("https://akabab.github.io/superhero-api/api/all.json")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Hero
	json.Unmarshal(responseData, &data)

	for _, k := range data {
		heroes = append(heroes, fmt.Sprintf("%d: %s", k.Id, k.Name))
	}

	sort.Strings(heroes)

	return heroes
}

func (a *App) GetImageUrlsById(Id string) []string {
	url := fmt.Sprintf("%s%s%s", "https://cdn.jsdelivr.net/gh/akabab/superhero-api@0.3.0/api/id/", strings.Split(Id, ":")[0], ".json")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Hero
	json.Unmarshal(responseData, &data)

	v := reflect.ValueOf(data.Images)
	values := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).String()
	}
	fmt.Println(values)

	return values
}
