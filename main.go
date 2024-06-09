package main

import (
	"net/http"
	_ "roman-numeral-api/docs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Roman Numeral API
// @version 1.0
// @description API to convert a range of numbers (integers) to Roman numerals.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	r.GET("/convert", convertRange)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}

// @Success 200 {object} []string
// @Failure 400 {object} map[string]string
// @Param range query string true "The range of numbers (integers) to convert, in the format 'from-to'"
// @Router /convert [get]
func convertRange(c *gin.Context) {
	rangeStr := c.Query("range")
	rangeParts := strings.Split(rangeStr, "-")

	if len(rangeParts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid input! Ensure that the input is in the format 'from-to'."})
		return
	}

	from, err1 := strconv.Atoi(rangeParts[0])
	to, err2 := strconv.Atoi(rangeParts[1])

	if err1 != nil || err2 != nil || from < 1 || to > 3999 || from > to {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid input. Ensure that numbers are integers and between 1 to 3999. Also, 'from' is less than or equal to 'to'."})
		return
	}

	// Create a sorted slice of integers from `from` to `to`
	numbers := make([]int, to-from+1)
	for i := range numbers {
		numbers[i] = from + i
	}

	// Convert each integer to a Roman numeral and append it to a slice
	romanNumerals := make([]string, len(numbers))
	for i, num := range numbers {
		romanNumerals[i] = toRoman(num)
	}

	c.JSON(http.StatusOK, romanNumerals)
}

func toRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}
	return roman
}
