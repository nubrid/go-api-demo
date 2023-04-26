package handlers

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nubrid/go-api-demo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JSON: { _id, createdAt, updatedAt, title }
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at" validate:"required"`
	Title     string             `json:"title" bson:"title" validate:"required,min=12"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func validateProductStruct(p Product) []*ErrorResponse {
	var errors []*ErrorResponse

	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		// _ = index
		for _, currentErr := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			// e.g.
			// {
			// 	"FailedField": "Product.Title",
			// 	"Tag": "min",
			// 	"Value": "12"
			// }
			element.FailedField = currentErr.StructNamespace()
			element.Tag = currentErr.Tag()
			element.Value = currentErr.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}

func CreateProduct(c *fiber.Ctx) error {
	product := Product{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// product = { title: "some title" }
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	errors := validateProductStruct(product)

	if errors != nil {
		return c.JSON(errors)
	}

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	// const database = client.db("products-api")
	// const collection = database.collection("products")
	collection := client.Database(db.Database).Collection(string(db.ProductsCollection))

	// const product = await collection.insertOne(product)
	_, err = collection.InsertOne(context.TODO(), product)

	if err != nil {
		return err
	}

	return c.JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	var products []*Product

	// const database = client.db("products-api")
	// const collection = database.collection("products")
	collection := client.Database(db.Database).Collection(string(db.ProductsCollection))

	// const cur = collection.find({})
	cur, err := collection.Find(context.TODO(), bson.D{
		primitive.E{},
	})

	if err != nil {
		return err
	}

	// await cur.forEach((p) => {
	// 	products.push(p);
	// })
	for cur.Next(context.TODO()) {
		var p Product

		err := cur.Decode(&p)

		if err != nil {
			return err
		}

		products = append(products, &p)
	}

	return c.JSON(products)
}
