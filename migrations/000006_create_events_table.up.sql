create table "events" (
	"id" serial primary key,
	"image" varchar(255),
	"tittle" varchar(100),
	"date" TIMESTAMP,
	"description" text,
	"location" int references "locations"("id"),
    "created_by" int references "users"("id")
);

