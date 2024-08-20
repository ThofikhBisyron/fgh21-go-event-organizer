create table "transactions" (
	"id" serial primary key,
	"event_id" INT REFERENCES "events"("id"),
    "payment_method_id" INT REFERENCES "payment_methods"("id"),
    "user_id" INT REFERENCES "users"("id")
);