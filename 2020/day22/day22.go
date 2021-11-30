package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crwilcox/advent-of-code/utils"
)

func playCombat(player1 []int, player2 []int) int {
	for i := 0; len(player1) > 0 && len(player2) > 0; i++ {
		// fmt.Println("Round", i)
		// fmt.Println("  Player 1:", player1)
		// fmt.Println("  Player 2:", player2)
		p1Card := player1[0]
		p2Card := player2[0]
		if p1Card > p2Card {
			player1 = append(player1, p1Card)
			player1 = append(player1, p2Card)
		} else {
			player2 = append(player2, p2Card)
			player2 = append(player2, p1Card)
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}

	var winningPlayer []int
	if len(player1) > 0 {
		winningPlayer = player1
	} else {
		// len(player2) > 0
		winningPlayer = player2
	}

	sum := 0
	for i := 1; i <= len(winningPlayer); i++ {
		// is  a multiplier
		card := winningPlayer[len(winningPlayer)-i]
		sum += card * i
	}
	return sum
}

func recursiveCombat(player1 []int, player2 []int, game int) int {
	player1PreviousRounds := make(map[string]bool)
	player2PreviousRounds := make(map[string]bool)
	innerGame := game
	winner := -1
outer:
	for i := 1; len(player1) > 0 && len(player2) > 0; i++ {
		// if there was a previous round in this
		// game that had exactly the same cards in the same order in the same players'
		// decks, the game instantly ends in a win for player 1.
		strPlayer1 := fmt.Sprint(player1)
		strPlayer2 := fmt.Sprint(player2)
		//fmt.Println("G", game, "R", i, "p1:", strPlayer1, "p2:", strPlayer2)
		if _, ok := player1PreviousRounds[strPlayer1]; ok {
			if _, ok := player2PreviousRounds[strPlayer2]; ok {
				winner = 1
				break outer
			}
		}

		// add hands to previous rounds
		player1PreviousRounds[strPlayer1] = true
		player2PreviousRounds[strPlayer2] = true

		// Otherwise, this round's cards must be in a new configuration; the players
		// begin the round by each drawing the top card of their deck as normal.
		p1Card := player1[0]
		p2Card := player2[0]

		player1 = player1[1:]
		player2 = player2[1:]

		// If both players have at least as many cards remaining in their deck as the
		// value of the card they just drew, the winner of the round is determined by
		// playing a new game of Recursive Combat (see below).
		if p1Card <= len(player1) && p2Card <= len(player2) {
			// fmt.Println("recursive game")
			innerGame = innerGame + 1
			recurseP1 := make([]int, p1Card)
			recurseP2 := make([]int, p2Card)
			copy(recurseP1, player1[:p1Card])
			copy(recurseP2, player2[:p2Card])
			res := recursiveCombat(recurseP1, recurseP2, innerGame)
			if res < 0 {
				winner = 1
			} else {
				winner = 2
			}
		} else {
			// Otherwise, at least one player must not have enough cards left in their deck
			// to recurse; the winner of the round is the player with the higher-value card.
			if p1Card > p2Card {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			player1 = append(player1, p1Card)
			player1 = append(player1, p2Card)
		} else {
			player2 = append(player2, p2Card)
			player2 = append(player2, p1Card)
		}
	}

	// As in regular Combat, the winner of the round (even if they won the round by
	// winning a sub-game) takes the two cards dealt at the beginning of the round
	// and places them on the bottom of their own deck (again so that the winner's
	// card is above the other card). Note that the winner's card might be the
	// lower-valued of the two cards if they won the round due to winning a
	// sub-game. If collecting cards by winning the round causes a player to have
	// all of the cards, they win, and the game ends.

	var winningPlayer []int
	if winner == 1 {
		winningPlayer = player1
	} else {
		winningPlayer = player2
	}

	sum := 0
	for i := 1; i <= len(winningPlayer); i++ {
		// i is  a multiplier
		card := winningPlayer[len(winningPlayer)-i]
		sum += card * i
	}

	// to allow differentiation of winner on reentrance,
	//player 1 is a negative score
	if len(player1) > 0 {
		return -1 * sum
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileToLines(filePath)
	if err != nil {
		panic(err)
	}

	player1 := []int{}
	for _, v := range lines[1 : len(lines)/2] {
		i, _ := strconv.Atoi(v)
		player1 = append(player1, i)
	}
	player2 := []int{}
	for _, v := range lines[(len(lines)/2)+2:] {
		i, _ := strconv.Atoi(v)
		player2 = append(player2, i)
	}

	// fmt.Println("player 1:", player1)
	// fmt.Println("player 2:", player2)

	count := playCombat(player1, player2)
	fmt.Println("🎄 Part 1 🎁:", count) // Answer: 32162
	count = recursiveCombat(player1, player2, 1)
	fmt.Println("🎄 Part 2 🎁:", count) // Answer: 32534
}
