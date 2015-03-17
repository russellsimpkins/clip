package clip

import (
	"fmt"
	"testing"
)


func TestToken(t *testing.T) {
	token := GenerateToken()
	fmt.Printf("Token.IntValue    = %d\n", token.IntValue)
	fmt.Printf("Token.StringValue = %s\n", token.StringValue)
}
