-- 28 November 2024 12:40:41 PM
CREATE TYPE "Currency" AS ENUM (
  'USD',
  'VND',
  'EUR'
);
CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "from_account_id" integer,
  "to_account_id" integer,
  "amount" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "gmail" varchar UNIQUE,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar,
  "amount" bigint NOT NULL,
  "descripton" text,
  "account_id" bigserial,
  "created_at" timestamptz DEFAULT (now())
);
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
CREATE INDEX ON "accounts" ("owner");
CREATE INDEX ON "entries" ("account_id");
COMMENT ON COLUMN "entries"."descripton" IS 'Content of the entry';
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

