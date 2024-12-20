INSERT INTO "users" ("username", "email", "password", "role_id") 
VALUES 
('Admin Super User', 'admin@mail.com', '$argon2id$v=19$m=65536,t=3,p=4$mJf+Cbj814Y+uvGR33vwYw$5C1OmZwsJgIN+bsqhw2savxVG4fSflW4i7Yonxt4GfA', '2')
RETURNING id;


INSERT INTO "profile" ("user_id", "full_name") 
VALUES 
('1', 'Admin Super User');