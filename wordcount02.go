package main 

import (
    "fmt"
    "io"
    "log"
    "os"
    "unicode"
    "github.com/pkg/profile"
    "bufio" // endring A fra wordcount01
)

func readbyte(r io.Reader) (rune, error) {
    var buf [1]byte 
    _, err := r.Read(buf[:])
    return rune(buf[0]), err
}

func main() {

    // For å generee profiler; kun en aktivert av gangen
    defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
    //defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
    //defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatalf("Kunne ikke åpne filen %q: %v", os.Args[1], err)
    }
   
    words := 0  
    inword := false  // tilstandsmaskin at vi er inn et ord (obs! unicode) 
    b := bufio.NewReader(f) // endring B fra wordcount01
    for {
        r, err := readbyte(b)
        if err == io.EOF {
            break
        } 
        if unicode.IsSpace(r) && inword {
            words++
            inword = false
        }
        inword = unicode.IsLetter(r)
    } 
    
    fmt.Printf("%q: %d words\n", os.Args[1], words)

}
