package main
// link für ein ähnliches online Spiel: https://www.abcya.com/games/number_ninja_factors
import (
	"fmt"
	"math/rand"
	"math"
	"time"
	"os"
)

type HighScore struct {
	Name  string
	Score int
	Zeit float64
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n <= 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }

    sqrtN := int(math.Sqrt(float64(n)))

    for i := 5; i <= sqrtN; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
    }
    return true
}


func main () {

fmt.Println("Spiele das")
fmt.Println(`#    ____  ___ ___   ____  _____   ___   ____       ____   ____   ____  ___ ___    ___       ____   ____  ___ ___    ___ `)
fmt.Println(`#   /    ||   |   | /    ||     | /   \ |    \     |    \ |    \ |    ||   |   |  /  _]     /    | /    ||   |   |  /  _]`)
fmt.Println(`#  |  o  || _   _ ||  o  ||__/  ||     ||  _  |    |  o  )|  D  ) |  | | _   _ | /  [_     |   __||  o  || _   _ | /  [_ `)
fmt.Println(`#  |     ||  \_/  ||     ||   __||  O  ||  |  |    |   _/ |    /  |  | |  \_/  ||    _]    |  |  ||     ||  \_/  ||    _]`)
fmt.Println(`#  |  _  ||   |   ||  _  ||  /  ||     ||  |  |    |  |   |    \  |  | |   |   ||   [_     |  |_ ||  _  ||   |   ||   [_ `)
fmt.Println(`#  |  |  ||   |   ||  |  ||     ||     ||  |  |    |  |   |  .  \ |  | |   |   ||     |    |     ||  |  ||   |   ||     |`)
fmt.Println(`#  |__|__||___|___||__|__||_____| \___/ |__|__|    |__|   |__|\_||____||___|___||_____|    |___,_||__|__||___|___||_____|`)
fmt.Println("#                                                                                                                        ")
fmt.Println("\tDu bist eine starke griechische Amazone, die fast alle Zahlen besiegen kann.\n\tNur Primzahlen werden dich besiegen. Schütze dich vor ihnen!")
fmt.Println("\t\t\tDir werden nun zufällig Zahlen gezeigt.")
fmt.Println("\t Tippe auf a (=amazon attacs), falls es sich um eine normale Zahl handelt")
fmt.Println("\t oder auf p (=prime protect), falls es sich um eine Primzahl handelt.") 
fmt.Println() 
var highScores []HighScore
for {
	rand.Seed(time.Now().UnixNano())
	score := 0
	playing := true
	startTime := time.Now()

	fmt.Print("Gib den Namen deiner Amazone ein: ")
	var playerName string
	fmt.Scanln(&playerName)

	for playing {
		num := rand.Intn(100)
		fmt.Printf("\nDu wist von der Zahl %d angegriffen: ", num)

		var input string
		fmt.Scanln(&input)

		if (input == "p" && isPrime(num)) || (input == "a" && !isPrime(num)) {
			score++
			fmt.Printf("Richtig! Dein Punktestand: %d\n\n", score)
		} else {
			fmt.Printf("Falsch!\n Dein Endpunktestand: %d\n", score)
			playing = false
		}
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	time := float64(elapsedTime.Seconds())
	fmt.Printf("Deine Geschwindigkeit (bekämpfte Zahlen pro Sekunde): %.2f\n", float64(score)/float64(elapsedTime.Seconds()))
	highScores = append(highScores, HighScore{Name: playerName, Score: score, Zeit: time})

	fmt.Println("Highscores:")
		for _, hs := range highScores {
			fmt.Printf(" %s hat %d Zahlen in %.2f Sekunden besiegt\n", hs.Name, hs.Score, hs.Zeit)
		}

	fmt.Print("\nNächste Amazone? (ja/nein): ")
	var playAgain string
	fmt.Scanln(&playAgain)
		if playAgain != "ja" {
			os.Exit(0)
		}
	}
}

