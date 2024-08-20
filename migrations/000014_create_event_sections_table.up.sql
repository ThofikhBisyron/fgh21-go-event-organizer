create table "event_sections" (
	"id" serial primary key,
	"name" VARCHAR(100),
    "price" int,
    "quantity" int,
    "event_id" int REFERENCES "events"("id")
);