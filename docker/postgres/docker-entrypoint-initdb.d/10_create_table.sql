--
-- テーブルの作成
--
CREATE TABLE user
(
    user_id    VARCHAR(50) NOT NULL,
    name       VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id)
);

CREATE TABLE todos
(
    todo_id    UUID        NOT NULL DEFAULT gen_random_ulid(),
    name       VARCHAR(50) NOT NULL,
    user_id    VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES companies (user_id) ON DELETE CASCADE
);