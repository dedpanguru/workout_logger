CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "password" varchar,
  "created_at" timestamp
);

CREATE TABLE "tokens" (
  "user_id" int PRIMARY KEY,
  "token" varchar
);

CREATE TABLE "workout" (
  "user_id" int PRIMARY KEY,
  "id" SERIAL,
  "created_at" timestamp,
  "day" date
);

CREATE TABLE "exercise" (
  "workout_id" int PRIMARY KEY,
  "type" int,
  "name" varchar,
  "resistance" varchar,
  "sets" int,
  "reps" int,
  "mins" int,
  "distance" varchar
);

ALTER TABLE "workout" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "exercise" ADD FOREIGN KEY ("workout_id") REFERENCES "workout" ("id");

ALTER TABLE "tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
