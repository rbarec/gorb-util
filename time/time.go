package time

import (
 "fmt"
 "time"
)

func GetInfo() string{
  fmt.Println("V1V1V1",time.Now() )
  return "gorb-util/time  version 1"
}
