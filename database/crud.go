package database

import (
	"context"
	"projectservice/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateProject(ctx context.Context, project *entity.Project, logger *zap.SugaredLogger) error {
	logger.Infof("createdat %s", project.CreatedAt)
	var id string
	_, err := db.NewInsert().Model(project).ExcludeColumn("created_at").Returning("id", "created_at").Exec(ctx, &id)
	if err != nil {
		logger.Errorw("Failed to insert project", "error", err)
		return err
	}
	logger.Infof("Inserted %s", id)

	if err != nil {
		logger.Errorw("Failed to retrieve ID of inserted project", "error", err)
		return err
	}

	// Update the project object with the generated ID
	project.Id = id
	return nil
}

func ReadAllByAuthorId(ctx context.Context, authorId string, logger *zap.SugaredLogger) ([]entity.Project, error) {
	logger.Info("Retrieving projects by author ID")

	var projects []entity.Project
	err := db.NewSelect().Model(&projects).Where("author_id = ?", authorId).Order("created_at DESC").Scan(ctx)
	if err != nil {
		logger.Warnw("Failed to retrieve projects by author ID", "error", err)
		return nil, err
	}

	return projects, nil
}

func ReadProjectById(ctx context.Context, id uuid.UUID, logger *zap.SugaredLogger) (*entity.Project, error) {
	logger.Infof("Retrieving project with ID: %s", id)
	project := new(entity.Project)
	err := db.NewSelect().Model(project).Where("id = ?", id).Scan(ctx)
	if err != nil {
		logger.Errorw("Failed to retrieve project by ID", "error", err)
		return nil, err
	}
	return project, nil
}

func UpdateProject(ctx context.Context, project *entity.Project, logger *zap.SugaredLogger) error {
	logger.Infof("Updating project with ID: %s", project.Id)
	_, err := db.NewUpdate().Model(project).Where("id = ?", project.Id).Exec(ctx)
	if err != nil {
		logger.Errorw("Failed to update project", "error", err)
		return err
	}
	return nil
}

func DeleteProject(ctx context.Context, id uuid.UUID, logger *zap.SugaredLogger) error {
	logger.Infof("Deleting %s", id)
	_, err := db.NewDelete().Model(&entity.Project{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		logger.Errorw("Failed to delete project", "error", err)
		return err
	}
	return nil
}

func UpdateProjectsAuthorName(ctx context.Context, user entity.User, logger *zap.SugaredLogger) error {
	logger.Infof("Updating author_name for projects with author ID: %s", user.Id)
	_, err := db.NewUpdate().
		Set("author_name = ?", user.Name).
		Where("author_id = ?", user.Id).
		Exec(ctx)
	if err != nil {
		logger.Errorw("Failed to update projects' author_name", "error", err)
		return err
	}
	return nil
}
