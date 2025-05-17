package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

const VERSION = "1.0.0"

var apes = []string{
	"🐒", "🦧", "🦍", "🐵",
}

var foods = []string{
	// Food items
	"🍎", "🍐", "🍊", "🍋", "🍌", "🍉", "🍇", "🍓", "🍈", "🍒",
	"🍑", "🥭", "🍍", "🥥", "🥝", "🍅", "🥑", "🥦", "🥬", "🥒",
	"🌶️", "🌽", "🥕", "🧄", "🧅", "🥔", "🍠", "🥐", "🥯", "🍞",
	"🥖", "🥨", "🧀", "🥚", "🍳", "🧈", "🥞", "🧇", "🍤", "🍗",
	"🍖", "🍔", "🍟", "🍕", "🌭", "🥪", "🌮", "🌯", "🥙", "🍝",
	"🍜", "🍲", "🍛", "🍣", "🍱", "🥟", "🦪", "🍦", "🍧", "🍨",
	"🍰", "🧁", "🥧", "🍫", "🍬", "🍭", "🍮", "🍯",
}

var extras = []string{
	"👾", "🤖", "👻", "🦖", "🦕", "🌈", "⭐", "🌟", "🔥", "💧", "🚀", "🎮", "🎨", "💡", "🎉", "🪐", "✨", "🎷", "💻",
}

var animals = []string{
	"🐶", "🐱", "🐭", "🐹", "🐰", "🦊", "🐻", "🐼", "🐨", "🐯",
	"🦁", "🐮", "🐷", "🐸", "🐔", "🐧", "🐦", "🦆", "🦅",
	"🐺", "🐗", "🐴", "🦄", "🐝", "🐛", "🦋", "🐌", "🐙", "🦑",
	"🦐", "🦞", "🦀", "🐬", "🐳", "🐋", "🦈", "🐊", "🐅", "🦓",
	"🐘", "🦛", "🦏", "🐪", "🐫", "🦒", "🦘", "🦬",
}

func listEmojis() {
	allEmojis := append(apes, foods...)
	allEmojis = append(allEmojis, extras...)
	allEmojis = append(allEmojis, animals...)
	for i, emoji := range allEmojis {
		fmt.Fprint(os.Stdout, emoji)
		if i < len(allEmojis)-1 {
			fmt.Fprint(os.Stdout, " ")
		}
	}
	fmt.Fprintln(os.Stdout)
}

func printRandApeEmoji() {
	randomEmoji := apes[rand.Intn(len(apes))]
	fmt.Fprint(os.Stdout, randomEmoji)
}

func printRandFoodEmoji() {
	randomEmoji := foods[rand.Intn(len(foods))]
	fmt.Fprint(os.Stdout, randomEmoji)
}

func printRandExtraEmoji() {
	randomEmoji := extras[rand.Intn(len(extras))]
	fmt.Fprint(os.Stdout, randomEmoji)
}

func printRandAnimalEmoji() {
	randomEmoji := animals[rand.Intn(len(animals))]
	fmt.Fprint(os.Stdout, randomEmoji)
}

func printRandEmoji() {
	allEmojis := append(apes, foods...)
	allEmojis = append(allEmojis, extras...)
	allEmojis = append(allEmojis, animals...)
	randomEmoji := allEmojis[rand.Intn(len(allEmojis))]
	fmt.Fprint(os.Stdout, randomEmoji)
}

func main() {
	listEmojisFlag := flag.Bool("list", false, "List all available emojis")
	apeEmojiFlag := flag.Bool("ape", false, "Use only ape emojis")
	foodEmojiFlag := flag.Bool("food", false, "Use only food emojis")
	extrasEmojiFlag := flag.Bool("extras", false, "Use only extras emojis")
	animalsEmojiFlag := flag.Bool("animals", false, "Use only animal emojis")
	versionFlag := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *listEmojisFlag {
		listEmojis()
		os.Exit(0)
	}

	if *apeEmojiFlag {
		printRandApeEmoji()
		os.Exit(0)
	}

	if *foodEmojiFlag {
		printRandFoodEmoji()
		os.Exit(0)
	}

	if *extrasEmojiFlag {
		printRandExtraEmoji()
		os.Exit(0)
	}

	if *animalsEmojiFlag {
		printRandAnimalEmoji()
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println("Version:", VERSION)
		os.Exit(0)
	}

	printRandEmoji()
}
