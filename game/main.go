package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/danicat/simpleansi"
)

type sprite struct {
	row int
	col int
}

//Config load emojis
type Config struct {
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
	UseEmoji bool   `json:"use_emoji"`
}

var cfg Config
var player sprite
var ghosts []*sprite
var score int
var dots int
var lives = 1

var maze []string

//for emojis
func loadConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}

//reading and loading the file to maze slice function
func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	//capturing players position
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col}
			case 'G':
				ghosts = append(ghosts, &sprite{row, col})
			case '.':
				dots++
			}
		}
	}
	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}

	//move player
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print(cfg.Player)

	//move ghosts
	for _, g := range ghosts {
		simpleansi.MoveCursor(g.row, g.col)
		fmt.Print(cfg.Ghost)
	}

	// Move cursor outside of maze drawing area
	simpleansi.MoveCursor(len(maze)+1, 0)
	//print score
	fmt.Println("Score:", score, "\tlives:", lives)

}

//Enablong Cbreak terminal mode
func initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak terminal mode:", err)
	}
}

//Restoring COoked terminal mode
func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to restore cooked terminal mode:", err)
	}
}

//Reading from Stdin
func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

//movement function
func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

//player movement
func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)

	removeDot := func(row, col int) {
		maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
	}

	switch maze[player.row][player.col] {
	case '.':
		dots--
		score++
		removeDot(player.row, player.col)
	case 'X':
		score += 10
		removeDot(player.row, player.col)
	}

}

//ghosts movement
func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.row, g.col = makeMove(g.row, g.col, dir)
	}
}

func main() {
	err := loadMaze("elmaze.txt")
	if err != nil {
		log.Println("Failed to load maze:", err)
		return
	}

	//load config.json
	err = loadConfig("config.json")
	if err != nil {
		log.Println("failed to load configuration:", err)
		return
	}

	initialize()
	defer cleanup()

	stillOn := true

	go func() {
		for {
			input, err := readInput()
			if err != nil {
				log.Print("error reading input:", err)
				break
			}

			//process collision
			for _, g := range ghosts {
				if player == *g {
					lives = 0
				}
			}

			// process movement
			movePlayer(input)
			//permanent break
			if input == "ESC" {
				stillOn = false
				break
			}

			//winning
			if dots == 0 {
				stillOn = false
				println("You Win")
				break
			}

			if lives <= 0 {
				stillOn = false
				println("You Loose")
				break
			}
		}

	}()

	go func() {
		for {
			moveGhosts()

			time.Sleep(time.Millisecond * 200)
		}

	}()
	//process collisions
	for _, g := range ghosts {
		if player == *g {
			lives = 0
		}
	}

	//Game loop
	for stillOn {
		printScreen()

		//process keyboard inputs

		time.Sleep(time.Millisecond * 100)

	}
}
