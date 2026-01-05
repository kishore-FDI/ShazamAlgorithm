package main
import(
	"fmt"
	"ShazamAlgorithm/audioprocessing"
)
func main(){
	fmt.Println("Hi Hello package main")
	audioprocessing.ConvertStereoToMono("nodkrai.mp3","op.mp3")
	sr , err := audioprocessing.GetSamplingRate("nodkrai.mp3")
	if err==nil{
		fmt.Println(sr)

	} else {
		fmt.Println("Error")
	}
}
