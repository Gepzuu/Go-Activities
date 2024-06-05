/*
In this code, I create a map called musicGenres that serves as a database for storing information 
about different music genres. The musicGenres map uses string keys (music genres) to map to inner maps. 
Each inner map contains specific details about a music genre, including its origin, instruments commonly used, and tempo.
*/

package main

import "fmt"

func main() {
    // Declare and initialize a map to store music genres and their characteristics
    musicGenres := map[string]map[string]string{
        "rock": {
            "origin":     "United States",
            "instruments": "guitar, bass, drums",
            "tempo":      "varied",
        },
        "pop": {
            "origin":     "United States",
            "instruments": "vocals, synthesizer, drums",
            "tempo":      "moderate",
        },
        "hip-hop": {
            "origin":     "United States",
            "instruments": "turntables, microphone",
            "tempo":      "fast",
        },
        "classical": {
            "origin":     "Europe",
            "instruments": "orchestra",
            "tempo":      "slow",
        },
    }

    // Access and print information about specific music genres
    fmt.Println("Information about rock music:")
    printGenreInfo(musicGenres, "rock")

    fmt.Println("\nInformation about pop music:")
    printGenreInfo(musicGenres, "pop")

    fmt.Println("\nInformation about hip-hop music:")
    printGenreInfo(musicGenres, "hip-hop")

    fmt.Println("\nInformation about classical music:")
    printGenreInfo(musicGenres, "classical")
}

// Function to print information about a specific music genre
func printGenreInfo(genres map[string]map[string]string, genre string) {
    info, ok := genres[genre]
    if ok {
        fmt.Println("Origin:", info["origin"])
        fmt.Println("Instruments:", info["instruments"])
        fmt.Println("Tempo:", info["tempo"])
    } else {
        fmt.Println("Genre not found in the map")
    }
}

/*

# Output

Information about rock music:
Origin: United States
Instruments: guitar, bass, drums
Tempo: varied

Information about pop music:
Origin: United States
Instruments: vocals, synthesizer, drums
Tempo: moderate

Information about hip-hop music:
Origin: United States
Instruments: turntables, microphone
Tempo: fast

Information about classical music:
Origin: Europe
Instruments: orchestra
Tempo: slow


*/