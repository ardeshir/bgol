package main


import (
//    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Question string
}

type Response struct {
     Question string
     Answer   string
}


func handler(e Event) (Response, error) {
/*  if len(e.Username) == 0 {
      return "", fmt.Errorf("No name given")
  }

  if e.Username[0] == 'F' {
    return "", fmt.Errorf("Don't Like : %s", e.Username)
  }
*/
   return Response{
     Question: e.Question,
     Answer: "I don't know. " + e.Question,

   }, nil
  // return fmt.Sprintf("<h2> Question:  %s  <br/> Answer: %s </h1>", e.Question, e.Answer) , nil
}

func main() {
  lambda.Start(handler)
}

