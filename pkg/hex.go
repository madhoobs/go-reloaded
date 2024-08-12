package hex

import (
	"fmt"
	"regexp"
	"strconv"
)

func Hello() {
  fmt.Println("Hello, world.")
}

func HexToDecimal(text string) string {
  regex := `(\w+) \(hex\)`
  re := regexp.MustCompile(regex)
  result := re.ReplaceAllFunc([]byte(text), func(match []byte) []byte {
    value := string(match[:len(match)-6])
    var decimalValue int64
    var err error
    decimalValue, err = strconv.ParseInt(value, 16, 64)
    if err != nil {
      return match
    }
    return []byte(fmt.Sprintf("%d", decimalValue))
  })
  return string(result)
}

