package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <name>", cmd.Name)
	}
	userInfo := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}
	user, err := s.Db.CreateUser(context.Background(), userInfo)
	if err != nil {
		return fmt.Errorf("coulnt create user: %w", err)
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Printf("User was created successfully: \n%s", userToString(user))
	return nil
}

func userToString(usr database.User) string {
	return fmt.Sprintf("id: %v,\nname: %s,\ncreatedAt: %v,\nupdatedAt: %v\n", usr.ID, usr.Name, usr.CreatedAt, usr.UpdatedAt)
}
