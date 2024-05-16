package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var convHandler *bot.ConversationHandler

var firstName string // Global variable to store user's first name

func main() {
	// Create a context to terminate the program when a signal is received
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Bot options
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	// Create the bot
	b, err := bot.New("EXAMPLE_TELEGRAM_BOT_TOKEN", opts...)
	if err != nil {
		panic(err)
	}

	// Register commands for the bot
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/cancel", bot.MatchTypeExact, cancelConversation)

	// Create a conversation handler and add stages
	convHandler = bot.NewConversationHandler()
	convHandler.AddStage(firstNameStage, firstNameHandler) // Add first name stage
	convHandler.AddStage(lastNameStage, lastNameHandler)   // Add last name stage

	//start the bot
	b.Start(ctx)
}

const (
	firstNameStage = iota // Definition of the first name stage = 0
	lastNameStage         // Definition of the last name stage = 1
)

// Default handler calls the CallStage function of conversation handler
func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	convHandler.CallStage(ctx, b, update)
}

// Handle /start command to start getting the user's name
func start(ctx context.Context, b *bot.Bot, update *models.Update) {
	convHandler.SetActiveStage(firstNameStage) //start the first name stage

	// Ask user to enter their name
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "enter your name or use the /cancel command to cancel.",
	})
}

// Handle the first name stage to get the user's first name
func firstNameHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	firstName = update.Message.Text

	convHandler.SetActiveStage(lastNameStage) //change stage to last name stage

	// Ask user to enter their last name
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "enter your last name",
	})
}

// Handle the last name stage to get the user's last name and send a hello message
func lastNameHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	lastName := update.Message.Text

	convHandler.End() // end the conversation

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello, " + firstName + " " + lastName + " :)",
	})
}

// Handle /cancel command to end the conversation
func cancelConversation(ctx context.Context, b *bot.Bot, update *models.Update) {
	convHandler.End() // end the conversation

	// Send a message to indicate the conversation has been cancelled
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "conversation cancelled",
	})
}
