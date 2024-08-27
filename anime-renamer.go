// anime-renamer.go
// this is a simple program that renames anime videos and subtitle files so mpv can fuzzy find the subtitles and auto load them
// it assumes the videos and subtitles are in the same folder
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	folderPath := ""
	fmt.Println("Enter the path to the folder containing the videos and subtitles:")
	fmt.Scanln(&folderPath)

	if folderPath == "" {
		fmt.Println("Error: Folder path is empty")
		fmt.Println("Press enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	// won't be used to find files, but will be used to rename files
	animeName := ""
	fmt.Println("Enter the name of the anime:")
	fmt.Scanln(&animeName)

	seasonEpisodeNamingConvention := ""
	fmt.Println("Enter the season and episode naming convention")
	fmt.Println("e.g., S#E#, S# E#, E#, S# - E#, etc.")
	fmt.Scanln(&seasonEpisodeNamingConvention)

	if seasonEpisodeNamingConvention == "" {
		fmt.Println("Error: Season and episode naming convention is empty")
		fmt.Println("Press enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	seasonEpisodeSubtitleNamingConvention := ""
	fmt.Println("Enter the season and episode naming convention for subtitles")
	fmt.Println("e.g., S#E#, S# E#, E#, S# - E#, etc.")
	fmt.Scanln(&seasonEpisodeSubtitleNamingConvention)

	if seasonEpisodeSubtitleNamingConvention == "" {
		fmt.Println("Error: Season and episode naming convention for subtitles is empty")
		fmt.Println("Press enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	videoExtensions := []string{".mkv", ".mp4", ".avi"}
	var videoFiles []string

	for _, ext := range videoExtensions {
		files, err := filepath.Glob(
			filepath.Join(folderPath, "*"+seasonEpisodeNamingConvention+"*"+ext),
		)

		if err != nil {
			fmt.Println("Error: Failed to find video files in folder")
			fmt.Println("Press enter to exit...")
			fmt.Scanln()
			os.Exit(1)
		}

		videoFiles = append(videoFiles, files...)
	}

	subtitleExtensions := []string{".srt", ".ass"}
	var subtitleFiles []string

	for _, ext := range subtitleExtensions {
		files, err := filepath.Glob(
			filepath.Join(folderPath, "*"+seasonEpisodeSubtitleNamingConvention+"*"+ext),
		)

		if err != nil {
			fmt.Println("Error: Failed to find subtitle files in folder")
			fmt.Println("Press enter to exit...")
			fmt.Scanln()
			os.Exit(1)
		}

		subtitleFiles = append(subtitleFiles, files...)
	}

	// warn user if the counts are different
	if len(videoFiles) != len(subtitleFiles) {
		fmt.Println("Number of video files does not match number of subtitle files")
		fmt.Println("Press enter to continue or ctrl+c to exit...")
		fmt.Scanln()
	}

	// the idea is to create pairs of video and subtitle files and
	// warn the user if there's an empty component in the pair
	// (i.e., there's a video or a subtitle without a matching counterpart)

	// also, if either the video or subtitle naming convention has a Season number,
	// and the other doesn't, we need to break it down into episode count
	// (e.g., there's 24 total episodes, and the subtitle file has "E24" in it, but the corresponding video file is "S02E12" or something)

	// TODO

	// pause and wait for user input before exiting
	fmt.Println("All done :) ありがとうございます！")
	fmt.Scanln()
}
