CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "nombre" varchar NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "hashed_password" varchar,
  "resume" varchar NOT NULL
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "nombre" varchar NOT NULL
);

CREATE TABLE "posts" (
  "id" SERIAL PRIMARY KEY,
  "category" integer NOT NULL,
  "title" varchar NOT NULL,
  "resume" varchar NOT NULL,
  "content" text NOT NULL,
  "author" integer NOT NULL,
  "created_at" timestamp
);

COMMENT ON COLUMN "posts"."content" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("author") REFERENCES "users" ("id");
