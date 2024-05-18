package bot

import (
	"context"
	"log"

	"github.com/go-telegram/bot/models"
)

// ConversationHandler is a structure that manages conversation functions.
type ConversationHandler struct {
	active         bool                // a flag indicating whether the conversation is active
	currentStageId int                 // the identifier of the active conversation stage
	stages         map[int]HandlerFunc // a map of conversation stages
}

// NewConversationHandler returns a new instance of ConversationHandler.
func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{
		active:         false,
		currentStageId: 0,
		stages:         make(map[int]HandlerFunc),
	}
}

// AddStage adds a conversation stage to the ConversationHandler.
func (c *ConversationHandler) AddStage(stageId int, hf HandlerFunc) {
	c.stages[stageId] = hf
}

// SetActiveStage sets the active conversation stage.
// Invalid currentStageId is not checked because if the CallStage function encounters an invalid id,
// it will not process it, so the stageId is not checked.
// if stageId <= len(c.stages)
func (c *ConversationHandler) SetActiveStage(stageId int) {
	if !c.active {
		c.active = true
	}

	c.currentStageId = stageId
}

// CallStage calls the function of the active conversation stage.
func (c *ConversationHandler) CallStage(ctx context.Context, b *Bot, update *models.Update) {
	if c.active {
		// hf = HandlerFunction
		if hf, ok := c.stages[c.currentStageId]; ok {
			hf(ctx, b, update)
		} else {
			log.Println("Error: Invalid stage id. No matching function found for the current stage id.")
		}
	}
}

// End ends the conversation.
func (c *ConversationHandler) End() {
	c.active = false
}
