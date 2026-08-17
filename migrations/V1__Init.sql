CREATE TABLE conversations (
    id                       CHAR(36) PRIMARY KEY,
    title                    TEXT NOT NULL,
    description              TEXT NOT NULL,
    avatar                   TEXT NOT NULL,
    create_time              TIMESTAMP(6) NOT NULL,
    update_time              TIMESTAMP(6) NOT NULL,
    delete_time              TIMESTAMP(6) NULL,
    last_message_id          CHAR(36) NULL,
    last_message_create_time TIMESTAMP(6) NULL
) ENGINE = InnoDB;


CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
) ENGINE = InnoDB;


CREATE TABLE conversation_users (
    conversation_id CHAR(36) NOT NULL,
    user_id         CHAR(36) NOT NULL,
    role            TEXT NOT NULL,
    join_time       TIMESTAMP(6) NOT NULL,

    PRIMARY KEY (conversation_id, user_id),

    CONSTRAINT fk_user_conversation_userships_conversation
        FOREIGN KEY (conversation_id)
        REFERENCES conversations(id),

    CONSTRAINT fk_user_conversation_userships_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
) ENGINE = InnoDB;


CREATE TABLE messages (
    id              CHAR(36) PRIMARY KEY,
    conversation_id CHAR(36) NOT NULL,
    sender_id       CHAR(36) NULL,
    content         TEXT NOT NULL,
    create_time     TIMESTAMP(6) NOT NULL,
    update_time     TIMESTAMP(6) NOT NULL,

    CONSTRAINT fk_messages_conversation
        FOREIGN KEY (conversation_id)
        REFERENCES conversations(id),

    CONSTRAINT fk_messages_sender
        FOREIGN KEY (sender_id)
        REFERENCES users(id)
) ENGINE = InnoDB;


ALTER TABLE conversations
    ADD CONSTRAINT fk_conversations_last_message
        FOREIGN KEY (last_message_id)
        REFERENCES messages(id);


CREATE INDEX idx_messages_conversation_create_time
    ON messages (conversation_id, create_time DESC);


CREATE INDEX idx_user_conversation_userships_user_id
    ON user_conversation_userships (user_id);