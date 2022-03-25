CREATE TABLE IF NOT EXISTS "users" (
	"id" varchar NOT NULL,
	"name" varchar(255) NULL,
	"email" varchar(255) NULL,
	"password_hash" varchar(255) NULL,
	CONSTRAINT "users_pk" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS "tasks" (
	"id" uuid NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" varchar(255) NOT NULL,
	"done" bool null default false,
	"created_at" date null,
	"finished_at" date null,
	"owner_id" uuid null,
	constraint "tasks_pk" primary key (id)
);