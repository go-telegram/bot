package bot

import (
	"context"
	"testing"

	"github.com/go-telegram/bot/models"
)

func Test_NewConversationHandler(t *testing.T) {
	convHandler := NewConversationHandler()

	// Check if conversation handler is initialized properly
	if convHandler == nil {
		t.Error("Failed to create a new conversation handler")
		return
	}

	// Check if conversation handler's stages map is initialized properly
	if convHandler.stages == nil {
		t.Error("Conversation handler's stages map is not initialized")
	}
}

func Test_ConversationHandler_AddStage(t *testing.T) {
	convHandler := NewConversationHandler()

	// Add a stage to the conversation handler
	convHandler.AddStage(1, func(ctx context.Context, b *Bot, update *models.Update) {})

	// Check if the stage is added correctly
	if _, ok := convHandler.stages[1]; !ok {
		t.Error("Failed to add stage to the conversation handler")
	}
}

func Test_ConversationHandler_SetActiveStage(t *testing.T) {
	convHandler := NewConversationHandler()

	// Set active stage to 1
	convHandler.SetActiveStage(1)

	// Check if active stage is set correctly
	if convHandler.currentStageId != 1 {
		t.Error("Failed to set active stage")
	}
}

func Test_ConversationHandler_CallStage(t *testing.T) {
	convHandler := NewConversationHandler()

	var state int

	handlerFunction := func(ctx context.Context, b *Bot, update *models.Update) {
		state = 1
	}

	// Add a stage to the conversation handler
	convHandler.AddStage(1, handlerFunction)

	// Set active stage to 1
	convHandler.SetActiveStage(1)

	// Create a dummy context, bot, and update
	ctx := context.Background()
	b := &Bot{}
	update := &models.Update{}

	// Call the active stage function
	convHandler.CallStage(ctx, b, update)

	// Check if the function of the active stage is called
	if state != 1 {
		t.Error("State is not updated as expected.")
	}
}

func Test_ConversationHandler_End(t *testing.T) {
	convHandler := NewConversationHandler()

	// End the conversation
	convHandler.End()

	// Check if conversation is ended
	if convHandler.active {
		t.Error("Conversation should be ended, but it's still active")
	}
}
