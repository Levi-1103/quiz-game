package collection

import (
	"backend/internal/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuizCollection struct {
	collection *mongo.Collection
}

func Quiz(collection *mongo.Collection) *QuizCollection {
	return &QuizCollection{
		collection: collection,
	}
}

func (c QuizCollection) InsertQuiz(quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(context.Background(), quiz)
	return err
}

func (c QuizCollection) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	result := c.collection.FindOne(context.Background(), bson.M{"_id": id})
	var quiz entity.Quiz
	err := result.Decode(&quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (c QuizCollection) GetQuizzes() ([]entity.Quiz, error) {
	cursor, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var quizzes []entity.Quiz
	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		return nil, err
	}

	return quizzes, nil

}
