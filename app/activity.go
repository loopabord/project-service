package app

import (
	"context"
	"log"
	"projectservice/database"
	"projectservice/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func withDBAndLogger(ctx context.Context, fn func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error)) (interface{}, error) {
	// Establish a session with the database cluster
	err := database.Connect()
	if err != nil {
		return nil, err // Return an error if connection fails
	}
	defer database.Close() // Defer closing the session to ensure it's closed after this function returns
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer logger.Sync()

	return fn(ctx, sugar)
}

func CreateProject(ctx context.Context, project entity.Project) (entity.Project, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		err := database.CreateProject(ctx, &project, logger)
		if err != nil {
			return entity.Project{}, err // Return an error if insertion fails
		}
		return project, nil
	})
	if err != nil {
		return entity.Project{}, err
	}
	var value entity.Project
	value = result.(entity.Project)
	log.Println(value.CreatedAt)
	return value, nil
}

func UpdateProject(ctx context.Context, project entity.Project) (entity.Project, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		err := database.UpdateProject(ctx, &project, logger)
		if err != nil {
			return entity.Project{}, err // Return an error if update fails
		}
		return project, nil
	})
	if err != nil {
		return entity.Project{}, err
	}
	return result.(entity.Project), nil
}

func ReadProject(ctx context.Context, id uuid.UUID) (entity.Project, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		project, err := database.ReadProjectById(ctx, id, logger)
		if err != nil {
			return entity.Project{}, err // Return an error if retrieval fails
		}
		return project, nil
	})
	if err != nil {
		return entity.Project{}, err
	}
	return *result.(*entity.Project), nil
}

func ReadAllProjects(ctx context.Context, authorId string) ([]entity.Project, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		projects, err := database.ReadAllByAuthorId(ctx, authorId, logger)
		if err != nil {
			return nil, err // Return an error if reading fails
		}
		return projects, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]entity.Project), nil
}

func DeleteProject(ctx context.Context, id uuid.UUID) (string, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		err := database.DeleteProject(ctx, id, logger)
		if err != nil {
			return "no good (bad)", err // Return an error if deletion fails
		}
		return "good yes", nil
	})
	if err != nil {
		return "no good (bad)", err
	}
	return result.(string), nil
}

func UpdateUser(ctx context.Context, user entity.User) (string, error) {
	result, err := withDBAndLogger(ctx, func(ctx context.Context, logger *zap.SugaredLogger) (interface{}, error) {
		err := database.UpdateProjectsAuthorName(ctx, &user, logger)
		if err != nil {
			return "no good (bad)", err // Return an error if deletion fails
		}
		return "good yes", nil
	})
	if err != nil {
		return "no good (bad)", err
	}
	return result.(string), nil
}
