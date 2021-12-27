package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var cache = []*reddit.Post{}
var cacheUsed = []int{}

func ConnectToDiscord(token string) {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!r" {

		if len(cache) == 0 || len(cacheUsed) == len(cache) {
			newCache, post, err := getRareInsult()
			if err == nil {
				cache = newCache
				cacheUsed = []int{}
				cacheUsed = append(cacheUsed, 0)
				s.ChannelMessageSendEmbed(m.ChannelID, createEmbed(&post))
			}
		} else {
			rand := len(cacheUsed)
			cacheUsed = append(cacheUsed, rand)
			fmt.Println(cacheUsed)
			post := cache[rand]
			s.ChannelMessageSendEmbed(m.ChannelID, createEmbed(post))
		}
	}
}

func createEmbed(post *reddit.Post) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       post.Title,
		Description: post.Body,
		Image:       &discordgo.MessageEmbedImage{URL: post.URL, ProxyURL: post.URL, Width: 200, Height: 200},
		URL:         "https://reddit.com" + post.Permalink,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Author",
				Value:  post.Author,
				Inline: true,
			},
			{
				Name:   "Likes",
				Value:  strconv.Itoa(post.Score),
				Inline: true,
			},
		},
	}
}

func cacheContainsNumber(number int) bool {
	for _, v := range cacheUsed {
		if v == number {
			return true
		}
	}
	return false
}
