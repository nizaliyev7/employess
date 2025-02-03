CREATE TABLE "employees" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(100) NOT NULL,
  "last_name" varchar(100) NOT NULL,
  "middle_name" varchar(100), 
  "phone" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "city_id" bigint NOT NULL
);

CREATE TABLE "cities" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(200) NOT NULL,
  "city_code" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "is_active" boolean NOT NULL DEFAULT true
);


CREATE INDEX ON "employees" ("last_name");
CREATE INDEX ON "employees" ("city_id");
CREATE INDEX ON "employees" ("created_at");
CREATE INDEX ON "employees" ("last_name", "city_id", "created_at");
CREATE INDEX ON "cities" ("name");


ALTER TABLE "employees" 
ADD CONSTRAINT "employees_city_id_fkey"
FOREIGN KEY ("city_id") REFERENCES "cities" ("id") ON DELETE RESTRICT;