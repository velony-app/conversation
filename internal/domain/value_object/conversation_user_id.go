package value_object

type ConversationUserID struct {
	conversationID ConversationID
	userID         UserID
}

func NewConversationUserID(
	conversationID ConversationID,
	userID UserID,
) ConversationUserID {
	return ConversationUserID{
		conversationID: conversationID,
		userID:         userID,
	}
}

func (id ConversationUserID) Value() (ConversationID, UserID) {
	return id.conversationID, id.userID
}

func (id ConversationUserID) String() string {
	return id.conversationID.String() + ":" + id.userID.String()
}

func (id ConversationUserID) ConversationID() ConversationID {
	return id.conversationID
}

func (id ConversationUserID) UserID() UserID {
	return id.userID
}
