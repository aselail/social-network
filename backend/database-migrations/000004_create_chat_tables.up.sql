CREATE TABLE IF NOT EXISTS conversation (
    id INTEGER PRIMARY KEY,
    -- 'direct' for 1-on-1 chats, 'group' for chats with 3+ participants.
    type TEXT NOT NULL CHECK(type IN ('direct', 'group')),
    
    -- The public name of a group chat. Can be NULL for direct message.
    name TEXT, 
    
    -- Timestamp of creation and last activity.
    created_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
    last_message_at TEXT DEFAULT (datetime('now', 'utc')) -- Updated by a trigger
);

CREATE TABLE IF NOT EXISTS conversation_participant (
    user INTEGER NOT NULL,
    conversation INTEGER NOT NULL,
    
    -- The composite primary key prevents a user from being added to the same conversation twice.
    PRIMARY KEY (user, conversation),
    
    FOREIGN KEY (user) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (conversation) REFERENCES conversation(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS message (
    id INTEGER PRIMARY KEY,
    conversation_id INTEGER NOT NULL,
    sender_id INTEGER, -- Can be NULL for system message or if user is deleted
    
    content TEXT, -- The main message content
    message_type TEXT NOT NULL DEFAULT 'text' CHECK(message_type IN ('text', 'image', 'file', 'system')),
    
    sent_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
    status TEXT NOT NULL DEFAULT 'sent' CHECK(status IN ('sent', 'delivered', 'read')),
    
    FOREIGN KEY (conversation_id) REFERENCES conversation(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL,
);

CREATE INDEX idx_message_conversation_id ON message(conversation_id);
CREATE INDEX idx_message_sender_id ON message(sender_id);
CREATE INDEX idx_message_sent_at ON message(sent_at);
CREATE INDEX idx_conversation_participant_user ON conversation_participant(user_id);
CREATE INDEX idx_conversation_participant_conversation ON conversation_participant(conversation_id);

CREATE TRIGGER update_conversation_timestamp_on_new_message
AFTER INSERT ON message
BEGIN
    UPDATE conversation
    SET last_message_at = NEW.sent_at
    WHERE id = NEW.conversation_id;
END;