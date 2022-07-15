CREATE TABLE users (
    id uuid default uuid_generate_v4(),
    "name" VARCHAR default '',
    "username" VARCHAR default '',
    "password" VARCHAR default '',
    "created_at" TIMESTAMP default current_timestamp,
    "updated_at" TIMESTAMP default current_timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE messages(
    id uuid default uuid_generate_v4(),
    sender_id uuid,
    recipient_id uuid,
    "message" text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    PRIMARY KEY (id)
);
