CREATE TABLE IF NOT EXISTS "categories"(
    "id" SERIAL PRIMARY KEY,
    "title" varchar(255) NOT null,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE if not exists "posts"(
    "id" serial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "image_url" VARCHAR(255),
    "user_id" INTEGER NOT NULL,
    "category_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "views_count" INTEGER not NULL default 0
);

create table if not exists "comments"(
    "id"     serial primary key,
    "post_id" integer not null,
    "user_id" integer not null,
    "description" text not null,
    "created_at"  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP WITH TIME ZONE
);

create table if not exists "likes"(
    "id"    serial primary key,
    "post_id" integer not null, 
    "user_id" integer not null,
    "status" boolean
);