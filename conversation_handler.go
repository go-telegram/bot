package bot

import (
	"context"

	"github.com/go-telegram/bot/models"
)

// stageFunction represents a function type for each stage of conversation.
type stageFunction func(ctx context.Context, b *Bot, update *models.Update)

// ConversationHandler is a structure that manages conversation functions.
type ConversationHandler struct {
	active         bool                  // a flag indicating whether the conversation is active
	currentStageId int                   // the identifier of the active conversation stage
	stages         map[int]stageFunction // a map of conversation stages
}

// NewConversationHandler returns a new instance of ConversationHandler.
func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{
		active:         false,
		currentStageId: 0,
		stages:         make(map[int]stageFunction),
	}
}

// AddStage adds a conversation stage to the ConversationHandler.
func (c *ConversationHandler) AddStage(stageId int, sf stageFunction) {
	c.stages[stageId] = sf
}

// SetActiveStage sets the active conversation stage.
func (c *ConversationHandler) SetActiveStage(stageId int) {
	if !c.active {
		c.active = true
	}
	c.currentStageId = stageId
}

// CallStage calls the function of the active conversation stage.
func (c *ConversationHandler) CallStage(ctx context.Context, b *Bot, update *models.Update) {
	if function, ok := c.stages[c.currentStageId]; ok && c.active {
		function(ctx, b, update)
	}
}

// End ends the conversation.
func (c *ConversationHandler) End() {
	c.active = false
}
