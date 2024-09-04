-- Create "students" table
CREATE TABLE "public"."students" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  "email" text NULL,
  "identifier" text NOT NULL,
  "password_hash" text NULL,
  "ask_new_password" boolean NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_students_deleted_at" to table: "students"
CREATE INDEX "idx_students_deleted_at" ON "public"."students" ("deleted_at");
-- Create index "idx_students_identifier" to table: "students"
CREATE UNIQUE INDEX "idx_students_identifier" ON "public"."students" ("identifier");
-- Create "sessions" table
CREATE TABLE "public"."sessions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "active" boolean NOT NULL,
  "meta_data" text NULL,
  "student_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_students_sessions" FOREIGN KEY ("student_id") REFERENCES "public"."students" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_sessions_deleted_at" to table: "sessions"
CREATE INDEX "idx_sessions_deleted_at" ON "public"."sessions" ("deleted_at");
-- Create "classes" table
CREATE TABLE "public"."classes" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_classes_deleted_at" to table: "classes"
CREATE INDEX "idx_classes_deleted_at" ON "public"."classes" ("deleted_at");
-- Create index "idx_classes_name" to table: "classes"
CREATE UNIQUE INDEX "idx_classes_name" ON "public"."classes" ("name");
-- Create "user_classes" table
CREATE TABLE "public"."user_classes" (
  "student_id" bigint NOT NULL,
  "class_id" bigint NOT NULL,
  PRIMARY KEY ("student_id", "class_id"),
  CONSTRAINT "fk_user_classes_class" FOREIGN KEY ("class_id") REFERENCES "public"."classes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_classes_student" FOREIGN KEY ("student_id") REFERENCES "public"."students" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
