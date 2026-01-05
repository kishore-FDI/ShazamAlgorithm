package audioprocessing

import(
	"os/exec"
	"fmt"
//	"log"
)

func GetSamplingRate(inputFile string)(int,error){
	out, err := exec.Command(
		"ffprobe",
		"-v", "error",
		"-select_streams", "a:0",
		"-show_entries", "stream=sample_rate",
		"-of", "default=nw=1:nk=1",
		inputFile,
	).Output()
	if err!=nil{
		return 0,fmt.Errorf("ffmpeg failed to find the sampling rate , maybe check whether ffmpeg is installed")}
	var sr int
	_, err = fmt.Sscanf(string(out), "%d", &sr)
	return sr, err
}

func ConvertStereoToMono(inputFile , outputFile string) error {
	
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		inputFile,
		"-ac",
		"1",
		outputFile,
	)

	if err := cmd.Run(); err!=nil{
		return fmt.Errorf("ffmpeg conversion failed: %w", err)
	}
	fmt.Printf("Successfully converted %s to mono %s\n", inputFile, outputFile)
	return nil
}
