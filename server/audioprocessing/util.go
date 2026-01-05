package audioprocessing

import(
	"os/exec"
	"log"
)

func convertStereoToMono(inputFile , outputFile string) error {
	
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		inputFile,
		"-ac",
		"1",
		outputFile
	)

	if err := cmd.Run(); err!=nil{
		return fmt.Errorf("ffmpeg conversion failed: %w", err)
	}
	fmt.Printf("Successfully converted %s to mono %s\n", inputFile, outputFile)
	return nil
}
