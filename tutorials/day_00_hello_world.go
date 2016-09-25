package tutorials

import "fmt"
import "bufio"
import "os"

func SayHello() {
	var inputReader *bufio.Reader
	var input string
	inputReader = bufio.NewReader(os.Stdin)
	input, _ = inputReader.ReadString('\n')

	fmt.Printf("Hello, World.\n%s", input)
	return
}
