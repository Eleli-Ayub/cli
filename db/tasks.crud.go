package db

import(
	"context"
	"go.mongodb.org/mongo-driver/bson"
)
type Task struct {
	Title 	string `json:"title" omitempty:"true" required:"true"`
	Priority string `json:"priority" omitempty:"true" required:"true"`
	Status   string `json:"status" omitempty:"true" required:"false"`
}

func InsertTask(task Task) error {
	taskColl := DBClient.Database("taskdb").Collection("tasks")
	
	_, err := taskColl.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}

func GetTasks() ([]Task, error) {
	taskColl := DBClient.Database("taskdb").Collection("tasks")
	
	cursor, err := taskColl.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var tasks []Task
	for cursor.Next(context.TODO()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(title string) error {
	taskColl := DBClient.Database("taskdb").Collection("tasks")
	
	_, err := taskColl.DeleteOne(context.TODO(), bson.M{"title": title})
	if err != nil {
		return err
	}
	return nil
}
