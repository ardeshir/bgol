package main


import (
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Username string
}


func handler(e Event) (string, error) {
  if len(e.Username) == 0 {
      return "", fmt.Errorf("No name given")
  }

  if e.Username[0] == 'F' {
    return "", fmt.Errorf("Don't Like : %s", e.Username)
  }
  
  return fmt.Sprintf("<h1>Golambda Serverless Rules -- %s</h1>", e.Username) , nil
}

func main() {
  lambda.Start(handler)
}

