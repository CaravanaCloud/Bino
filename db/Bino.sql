CREATE TABLE "mentors" (
  "id" SERIAL PRIMARY KEY,
  "firstname" varchar(30),
  "lastname" varchar(60),
  "email" varchar(84),
  "password" varchar(255),
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "bookings" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100),
  "email" varchar(84),
  "reserved_date" date,
  "reserved_time" date,
  "status" int,
  "mentor_id" int
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("mentor_id") REFERENCES "mentors" ("id");
