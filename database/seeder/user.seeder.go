package seeder

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"ordent-assessment/constant"
	"ordent-assessment/entity"
	"ordent-assessment/repository"
	"time"
)

func UserSeeder(db *mongo.Database) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	repo := repository.NewUserRepository(db)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	dataUsers := []entity.User{
		{
			Name:     "Admin",
			Username: "admin",
			Password: string(passwordHash),
			Permissions: []string{
				constant.CAN_READ_PRODUCT,
				constant.CAN_CREATE_PRODUCT,
				constant.CAN_UPDATE_PRODUCT,
				constant.CAN_DELETE_PRODUCT,
			},
		},
	}

	for _, dataUser := range dataUsers {
		_, err := repo.Create(ctx, dataUser)
		if err != nil {
			return err
		}
	}

	fmt.Println("seeder user completed !")
	return nil
}
