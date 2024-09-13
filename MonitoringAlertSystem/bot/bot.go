package bot

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
)

const botToken = "7214843261:AAHtji56Baa9CWb1clJykbLv-U4K8CGVJvM"

func ConnectToTelegram() (*telego.Bot, error) {
	// Create bot and enable debugging info
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	// (more on configuration in examples/configuration/main.go)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	botUser, err = bot.GetMe()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Bot user: %+v\n", botUser)

	return bot, nil
}

func SendMessageToOneUser(bot *telego.Bot, message string) error {
	// Get the list of all chat IDs

	// Send the message to each 		// Check if the error is due to the bot being blocked by the useruser
	_, err := bot.SendMessage(tu.Message(tu.ID(465184827), message)) // Attempt to send the message
	if err != nil {
		log.Printf("Failed to send message to user %d: %v", 465184827, err)

	}

	fmt.Printf("Message sent to %s user\n", "Ozodjon")
	return nil
}
