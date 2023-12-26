package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
    Email string `valid:"email"`
    Password string `valid:"required"`
}

func main() {
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err = govalidator.ValidateStruct(user)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        sess, err := session.NewSession(&aws.Config{
            Region: aws.String("us-west-2")},
        )

        svc := dynamodb.New(sess)

        result, err := svc.GetItem(&dynamodb.GetItemInput{
            TableName: aws.String("Users"),
            Key: map[string]*dynamodb.AttributeValue{
                "Email": {
                    S: aws.String(user.Email),
                },
            },
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var dbUser User
        err = dynamodbattribute.UnmarshalMap(result.Item, &dbUser)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        if user.Password != dbUser.Password {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Forward the request to the main service...
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}