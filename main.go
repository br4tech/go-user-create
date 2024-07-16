package main

import (
	"fmt"

	"github.com/br4tech/go-user-create/adapter"
	"github.com/br4tech/go-user-create/domain"
	services "github.com/br4tech/go-user-create/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})

	userRepository := adapter.NewGormRepository[domain.User](db)
	userService := services.NewUserService(userRepository)

	// Criar um novo usuário
	newUser := &domain.User{Name: "Guilherme", Email: "Guilherme@example.com"}
	if err := userService.CreateUser(newUser); err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}
	fmt.Println("Usuário criado com sucesso:", newUser)

	// Buscar um usuário pelo ID
	user, err := userService.GetUserByID(newUser.ID)
	if err != nil {
		fmt.Println("Erro ao buscar usuário:", err)
		return
	}
	fmt.Println("Usuário encontrado:", user)

	// Atualizar o usuário
	user.Name = "Silva"
	if err := userService.UpdateUser(user); err != nil {
		fmt.Println("Erro ao atualizar usuário:", err)
		return
	}
	fmt.Println("Usuário atualizado:", user)

	// Excluir o usuário
	if err := userService.DeleteUser(user.ID); err != nil {
		fmt.Println("Erro ao excluir usuário:", err)
		return
	}
	fmt.Println("Usuário excluído com sucesso")
}
