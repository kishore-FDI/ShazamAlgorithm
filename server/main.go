package main
import(
	"fmt"
	"audioprocessing"
)
func main(){
	fmt.Println("Hi Hello package main")
	audioprocessing.convertStereoToMono("nodkrai.mp3","op.mp3")
}
