package main

import (
	"fmt"
	"go-mongo-tutorial/config"
	"go-mongo-tutorial/src/modules/profile/model"
	"go-mongo-tutorial/src/modules/profile/repository"
	"time"
)

func main() {
	fmt.Println("Go Mongo DB")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	saveProfile(profileRepository)
}

//fungsi Save
func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Ferly"
	p.LastName = "Apiang"
	p.Email = "ferly.apiang9@gmail.com"
	p.Password = "12345"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}
