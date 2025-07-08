package models

// ChecklistTask https://core.telegram.org/bots/api#checklisttask
type ChecklistTask struct {
	ID              int             `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User           `json:"completed_by_user,omitempty"`
	CompletionDate  int             `json:"completion_date,omitempty"`
}

// Checklist https://core.telegram.org/bots/api#checklist
type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done,omitempty"`
}

// InputChecklistTask https://core.telegram.org/bots/api#inputchecklisttask
type InputChecklistTask struct {
	ID           int             `json:"id"`
	Text         string          `json:"text"`
	ParseMode    ParseMode       `json:"parse_mode,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

// InputChecklist https://core.telegram.org/bots/api#inputchecklist
type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                ParseMode            `json:"parse_mode,omitempty"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
}

// ChecklistTasksDone https://core.telegram.org/bots/api#checklisttasksdone
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIDs    []int    `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIDs []int    `json:"marked_as_not_done_task_ids,omitempty"`
}

// ChecklistTasksAdded https://core.telegram.org/bots/api#checklisttasksadded
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}
