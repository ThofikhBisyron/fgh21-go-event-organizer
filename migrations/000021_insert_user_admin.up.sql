INSERT INTO "users" ("username", "email", "password", "role_id") 
VALUES 
('Admin Super User', 'admin@mail.com', '$argon2id$v=19$m=65536,t=3,p=4$Q/l7gyvwtliBKBwlZ0YfXQ$GpEDDUlPhXtg0Lllfj9Hz0DWghyFr/v0+7pzYolU5xw', '2')
RETURNING id;


INSERT INTO "profile" ("user_id", "full_name", "gender") 
VALUES 
('1', 'Admin Super User', '1');