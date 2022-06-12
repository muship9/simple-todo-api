----------------------------
-- カンパニー系
----------------------------

INSERT INTO user_profile (user_id, name)
VALUES ('testUser', 'testUser')
;

INSERT INTO todos (todo_id, name, user_id)
VALUES ('first-todo', '歯磨き', 'testUser')
     , ('second-todo', '選択', 'testUser')
     , ('third-todo', '掃除', 'testUser')
;