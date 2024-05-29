package app

import (
	"context"
	"encoding/json"
	"log"

	"projectservice/entity"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type ActionFunc func(ctx context.Context, data []byte) (interface{}, error)

func handleMessage(action ActionFunc) nats.MsgHandler {
	return func(msg *nats.Msg) {
		result, err := action(context.Background(), msg.Data)
		if err != nil {
			log.Println("Error:", err)
			msg.Respond([]byte("Error processing request"))
			return
		}

		response, err := json.Marshal(result)
		if err != nil {
			log.Println("Error marshalling response:", err)
			msg.Respond([]byte("Error marshalling response"))
			return
		}

		msg.Respond(response)
	}
}

// CreateProjectHandler handles the CreateProject messages
func CreateProjectHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		var project entity.Project
		if err := json.Unmarshal(data, &project); err != nil {
			return nil, err
		}
		return CreateProject(ctx, project)
	})
}

// UpdateProjectHandler handles the UpdateProject messages
func UpdateProjectHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		var project entity.Project
		if err := json.Unmarshal(data, &project); err != nil {
			return nil, err
		}
		return UpdateProject(ctx, project)
	})
}

// ReadProjectHandler handles the ReadProject messages
func ReadProjectHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		id, err := uuid.Parse(string(data))
		if err != nil {
			return nil, err
		}
		return ReadProject(ctx, id)
	})
}

// ReadAllProjectsHandler handles the ReadAllProjects messages
func ReadAllProjectsHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		id := string(data)
		return ReadAllProjects(ctx, id)
	})
}

// DeleteProjectHandler handles the DeleteProject messages
func DeleteProjectHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		id, err := uuid.Parse(string(data))
		if err != nil {
			return nil, err
		}
		return DeleteProject(ctx, id)
	})
}

// UpdateUserHandler handles the UpdateUser messages
func UpdateUserHandler() nats.MsgHandler {
	return handleMessage(func(ctx context.Context, data []byte) (interface{}, error) {
		id, err := uuid.Parse(string(data))
		if err != nil {
			return nil, err
		}
		return DeleteProject(ctx, id)
	})
}
