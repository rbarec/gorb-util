package time

import (
 "time"
 "fmt"
)

func GetInfoV2() string{
  fmt.Println("XXXXXXX ",time.Now())
  return "gorb-util/time  version 2"
}
