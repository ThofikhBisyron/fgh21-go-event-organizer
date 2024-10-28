INSERT INTO "users" ("username", "email", "password", "role_id") 
VALUES 
('Admin Super User', 'admin@mail.com', '$argon2id$v=19$m=65536,t=3,p=4$svXPb18xktWQ+TPiI07TvA$L2WaPgRhLUl3HpI1nFkwhjN7mavJ0Uv2ZG6bpRkgnAI', '2')
RETURNING id;


INSERT INTO "profile" ("user_id", "full_name") 
VALUES 
('1', 'Admin Super User');