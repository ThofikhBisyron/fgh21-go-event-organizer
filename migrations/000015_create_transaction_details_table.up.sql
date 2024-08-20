create table "transaction_details" (
	"id" serial primary key,
	"transaction_id" int references "transactions"("id"),
    "section_id" int references "event_sections"("id"),
    "ticket_qty" int
);