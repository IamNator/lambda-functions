package main 


import (
  "context"
  "github.com/aws/lambda"
)


func handler(ctx context.Context, req interface{}) error {
  return nil
}

func main(){
  lambda.Handle(handler)
}
