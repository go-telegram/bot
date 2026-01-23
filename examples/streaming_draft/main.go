package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// This example demonstrates how to use sendMessageDraft to stream
// partial messages to users, simulating LLM text generation.
//
// Send any text message to the bot after it has been started.

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	// Pass the thread ID to keep the response in the same thread/topic
	threadID := update.Message.MessageThreadID

	// Simulate an LLM response that streams word by word
	response := "Hello! I am an AI assistant. I can help you with many tasks, " +
		"such as answering questions, writing code, and having conversations. " +
		"How can I assist you today?"

	// Generate a unique draft ID for this streaming session
	// Using a simple numeric string
	draftID := fmt.Sprintf("%d", time.Now().UnixNano())

	err := streamResponse(ctx, b, chatID, threadID, draftID, response)
	if err != nil {
		fmt.Printf("error streaming response: %v\n", err)
	}
}

// streamResponse uses sendMessageDraft to stream text progressively
func streamResponse(ctx context.Context, b *bot.Bot, chatID int64, threadID int, draftID, fullText string) error {
	chunks := splitIntoChunks(fullText, 3) // ~3 words per chunk
	var currentText string

	for i, chunk := range chunks {
		currentText += chunk

		// Add cursor indicator while streaming
		displayText := currentText
		if i < len(chunks)-1 {
			displayText += " ▌"
		}

		_, err := b.SendMessageDraft(ctx, &bot.SendMessageDraftParams{
			ChatID:          chatID,
			MessageThreadID: threadID, // Keep response in same thread/topic
			DraftID:         draftID,
			Text:            displayText,
		})
		if err != nil {
			return fmt.Errorf("failed to send draft: %w", err)
		}

		// Simulate LLM generation delay
		time.Sleep(time.Millisecond * 200)
	}

	// Finalize the message with sendMessage
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		Text:            currentText,
	})
	if err != nil {
		return fmt.Errorf("failed to send final message: %w", err)
	}

	return nil
}

// splitIntoChunks splits text into chunks of approximately n words each
func splitIntoChunks(text string, wordsPerChunk int) []string {
	var chunks []string
	var currentChunk string
	wordCount := 0

	for _, char := range text {
		currentChunk += string(char)

		if char == ' ' {
			wordCount++
			if wordCount >= wordsPerChunk {
				chunks = append(chunks, currentChunk)
				currentChunk = ""
				wordCount = 0
			}
		}
	}

	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
