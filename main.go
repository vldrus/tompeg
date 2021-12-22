package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	i := flag.String("i", "", "Input file")
	o := flag.String("o", "tompeg.mp4", "Output file")
	vh := flag.String("vh", "480", "Video height")
	vb := flag.String("vb", "600k", "Video bitrate")
	ab := flag.String("ab", "66k", "Audio bitrate")

	flag.Parse()

	if len(*i) == 0 {
		fmt.Println("Specify at least input file!")
		flag.Usage()
		os.Exit(1)
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i", *i,
		"-vf", "scale=-2:'min("+*vh+",ih)':flags=lanczos,setsar=1:1",
		"-map_metadata", "-1",
		"-preset", "slower",
		"-pix_fmt", "yuv420p",
		"-b:v", *vb,
		"-ar", "44100",
		"-b:a", *ab,
		*o,
	)

	fmt.Println(cmd.Args)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Cannot start ffmpeg:", err)
		os.Exit(1)
	}
}
