package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// determine the game mode
var Mode = "EASY"

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
   0  1  2
0  o  x  x
1  x  x  o
2  o  x  o
*/

// initial board
var board = [][]string{
	[]string{"-", "-", "-"},
	[]string{"-", "-", "-"},
	[]string{"-", "-", "-"},
}

func printBoard() {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	println("")
}

func checkTheWinner(player string) string {
	switch {
	case board[0][0] == player && board[0][1] == player && board[0][2] == player:
		return player
	case board[1][0] == player && board[1][1] == player && board[1][2] == player:
		return player
	case board[2][0] == player && board[2][1] == player && board[2][2] == player:
		return player
	case board[0][0] == player && board[1][0] == player && board[2][0] == player:
		return player
	case board[0][1] == player && board[1][1] == player && board[2][1] == player:
		return player
	case board[0][2] == player && board[1][2] == player && board[2][2] == player:
		return player
	case board[0][0] == player && board[1][1] == player && board[2][2] == player:
		return player
	case board[0][2] == player && board[1][1] == player && board[2][0] == player:
		return player
	default:
		return "-"
	}
}

func isItValidPos(x, y int) bool {
	if (x >= 0 && x <= 2) && (y >= 0 && y <= 2) {
		return board[x][y] == "-"
	}
	return false
}

func pressEnterToContinue() {
	fmt.Println("Press Enter to continue...")
	var input string
	fmt.Scanln(&input)
}

func makeRandomInitialMove() {
	board[rand.Intn(2)][rand.Intn(2)] = "x"
}

func makeComputerMove(computer string) {
	i := 0
	if Mode == "EASY" {
		for {
			fmt.Println("Move", i)
			randX := rand.Intn(2)
			randY := rand.Intn(2)
			if isItValidPos(randX, randY) {
				board[randX][randY] = computer
				break
			}
			i++
		}
	}
}

func checkWinner(computer, player string) string {
	if checkTheWinner(player) == player {
		return player
	} else if checkTheWinner(computer) == computer {
		return computer
	}
	return ""
}

func isItDraw() bool {
	draw := true
	for i := 0; i < len(board); i++ {
		if !draw {
			break
		}
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "-" {
				draw = false
				break
			}
		}
	}

	return draw
}

func main() {
	winner := ""
	var x, y int
	var player, computer string

	// choose who starts first
	if rand.Intn(2) == 0 {
		player = "x"
		computer = "o"
	} else {
		player = "o"
		computer = "x"
	}

	// if the player is o that means computer start first
	if player == "o" {
		makeRandomInitialMove()
	}

	// run as long as winner selected
	for winner == "" {
		for {
			clearScreen()
			fmt.Println("Current board")
			printBoard()
			fmt.Println("Player is", player)
			fmt.Println("Computer is", computer)
			fmt.Print("Please enter x and y pos: ")
			fmt.Scanf("%d %d\n", &x, &y)
			if !isItValidPos(x, y) {
				fmt.Println("Invalid position")
				pressEnterToContinue()
			} else {
				// player makes the move
				board[x][y] = player
				break
			}
		}
		if player == "o" {
			winner = checkWinner(computer, player)
			makeComputerMove(computer)
		} else {
			makeComputerMove(computer)
			winner = checkWinner(computer, player)
		}

		if isItDraw() && winner == "" {
			winner = "-"
		}
	}

	if winner == player {
		fmt.Println("Winner is player")
	} else if winner == computer {
		fmt.Println("Winner is computer")
	} else {
		fmt.Println("Nobody won its draw")
	}

}
