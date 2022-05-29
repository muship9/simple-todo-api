----------------------------
-- カンパニー系
----------------------------

INSERT INTO user (user_id, name)
VALUES ('testUser1', 'テストユーザー')
;

INSERT INTO todos (todo_id, name, user_id)
VALUES ('0180d65b-b935-ab01-9266-15fe1b29fbd5', '歯磨き', 'testUser1')
     , ('0180d65b-b935-ab01-9266-15fe1b29fbd5', '選択', 'testUser1')
     , ('0180d65b-b935-ab01-9266-15fe1b29fbd5', 'お手洗い', 'testUser1')
;