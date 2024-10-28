DELETE FROM "profile" WHERE user_id IN (SELECT id FROM "users" WHERE email = 'admin@mail.com');
DELETE FROM "users" WHERE email = 'admin@mail.com';