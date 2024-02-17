package brot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type ChannelType string

const (
	General ChannelType = "general"
	Logs ChannelType = "logs"
)

type Channel struct {
	Type ChannelType
	Id string
}

// Shout sends a message to a channel
//
// returns error if message fails to send
//
// requires environment variables: CHANNEL_ID_GENERAL, CHANNEL_ID_LOGS, BOT_TOKEN
func Shout(msg string, channelType ChannelType ) error {
	sesh, err := brotInit()
	if err != nil {
		return err
	}

	switch channelType {
		case General:
		chennelId := os.Getenv("CHANNEL_ID_GENERAL")
		channel := Channel{
			Type: channelType,
			Id: chennelId,
		}
		case Logs:
		channelId := os.Getenv("CHANNEL_ID_LOGS")
		channel := Channel{
			Type: channelType,
			Id: channelId,
		}
	}

	_, err = sesh.ChannelMessageSend(channel.Id, msg)
	if err != nil {
		return err
	}
}

// load environment variables and initialize bot
func brotInit() *discordgo.Session, error {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// load environment variables
	botToken := os.Getenv("BOT_TOKEN")

	// initialize bot
	sesh, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return nil, err
	}

	return sesh, nil
}
