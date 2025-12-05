package main

import (
	"fmt"
	"log"

	"streaming/internal/models"
	"streaming/internal/repos"
	"streaming/internal/services"
)

func main() {

	repo := repos.NewMemoryRepo()
	service := services.NewStreamingService(repo, repo)

	u, err := service.RegisterUser("u1", "Andy", "andy@example.com")
	if err != nil {
		log.Fatalf("Error creando usuario: %v", err)
	}
	fmt.Println("Usuario creado:", u.GetID(), u.GetName())

	c, err := service.AddContent("c1", "Mi primer video", models.Video, 300)
	if err != nil {
		log.Fatalf("Error creando contenido: %v", err)
	}
	fmt.Println("Contenido creado:", c.GetID(), c.GetTitle())

	msg, err := service.Play("u1", "c1")
	if err != nil {
		log.Fatalf("Error reproduciendo: %v", err)
	}
	fmt.Println(msg)

	_, err = service.Play("u1", "no_existe")
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	history, _ := service.GetHistory("u1")
	fmt.Println("\nHistorial del usuario:")
	for _, item := range history {
		fmt.Println("-", item.GetTitle())
	}
}
