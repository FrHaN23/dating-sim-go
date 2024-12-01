-- Adminer 4.8.1 PostgreSQL 16.3 (Debian 16.3-1.pgdg120+1) dump

DROP TABLE IF EXISTS "pictures";
DROP SEQUENCE IF EXISTS pictures_id_seq;
CREATE SEQUENCE pictures_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."pictures" (
    "id" bigint DEFAULT nextval('pictures_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "path" text,
    "mime" text,
    "original_name" text,
    "user_id" bigint,
    CONSTRAINT "pictures_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_pictures_deleted_at" ON "public"."pictures" USING btree ("deleted_at");


DROP TABLE IF EXISTS "preferances";
DROP SEQUENCE IF EXISTS preferances_id_seq;
CREATE SEQUENCE preferances_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."preferances" (
    "id" bigint DEFAULT nextval('preferances_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "interested_in" text,
    "age_range" bigint,
    "max_distance_km" bigint,
    CONSTRAINT "preferances_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_preferances_deleted_at" ON "public"."preferances" USING btree ("deleted_at");

INSERT INTO "preferances" ("id", "created_at", "updated_at", "deleted_at", "interested_in", "age_range", "max_distance_km") VALUES
(1,	'2024-11-30 14:33:40.554121+00',	'2024-11-30 14:33:40.554121+00',	NULL,	'Music',	25,	50),
(2,	'2024-11-30 17:38:32.861652+00',	'2024-11-30 17:38:32.861652+00',	NULL,	'Music,Movie',	25,	50),
(3,	'2024-11-30 17:39:25.601241+00',	'2024-11-30 17:39:25.601241+00',	NULL,	'Music,Movie,Game',	25,	50),
(4,	'2024-11-30 17:41:54.593614+00',	'2024-11-30 17:41:54.593614+00',	NULL,	'Hiking',	23,	50),
(5,	'2024-11-30 17:42:32.734706+00',	'2024-11-30 17:42:32.734706+00',	NULL,	'Hiking,Shopping',	23,	50),
(6,	'2024-11-30 17:43:35.444311+00',	'2024-11-30 17:43:35.444311+00',	NULL,	'Hiking,Shopping,Eat',	23,	50),
(7,	'2024-11-30 17:43:48.860145+00',	'2024-11-30 17:43:48.860145+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(8,	'2024-11-30 17:43:55.707561+00',	'2024-11-30 17:43:55.707561+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(9,	'2024-11-30 17:43:59.851455+00',	'2024-11-30 17:43:59.851455+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(10,	'2024-11-30 17:44:05.433362+00',	'2024-11-30 17:44:05.433362+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(11,	'2024-12-01 00:04:08.9018+00',	'2024-12-01 00:04:08.9018+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(12,	'2024-12-01 00:04:15.204073+00',	'2024-12-01 00:04:15.204073+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(13,	'2024-12-01 00:04:20.496638+00',	'2024-12-01 00:04:20.496638+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(14,	'2024-12-01 00:04:26.243443+00',	'2024-12-01 00:04:26.243443+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(15,	'2024-12-01 00:04:31.0335+00',	'2024-12-01 00:04:31.0335+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(16,	'2024-12-01 00:04:35.689725+00',	'2024-12-01 00:04:35.689725+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(17,	'2024-12-01 00:29:23.446442+00',	'2024-12-01 00:29:23.446442+00',	NULL,	'Music',	25,	50),
(18,	'2024-12-01 02:37:33.640138+00',	'2024-12-01 02:37:33.640138+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50),
(19,	'2024-12-01 02:56:23.238939+00',	'2024-12-01 02:56:23.238939+00',	NULL,	'Hiking,Shopping,Eat,Sleep',	23,	50);

DROP TABLE IF EXISTS "swipes";
DROP SEQUENCE IF EXISTS swipes_id_seq;
CREATE SEQUENCE swipes_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."swipes" (
    "id" bigint DEFAULT nextval('swipes_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" bigint NOT NULL,
    "profile_id" bigint NOT NULL,
    "action" boolean NOT NULL,
    CONSTRAINT "swipes_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_swipes_deleted_at" ON "public"."swipes" USING btree ("deleted_at");

INSERT INTO "swipes" ("id", "created_at", "updated_at", "deleted_at", "user_id", "profile_id", "action") VALUES
(1,	'2024-11-30 23:59:36.524656+00',	'2024-11-30 23:59:36.524656+00',	NULL,	1,	3,	't'),
(2,	'2024-12-01 00:00:06.943138+00',	'2024-12-01 00:00:06.943138+00',	NULL,	1,	5,	't'),
(3,	'2024-12-01 00:00:17.417467+00',	'2024-12-01 00:00:17.417467+00',	NULL,	1,	10,	't'),
(4,	'2024-12-01 00:01:23.088456+00',	'2024-12-01 00:01:23.088456+00',	NULL,	1,	7,	'f'),
(5,	'2024-12-01 00:03:00.986274+00',	'2024-12-01 00:03:00.986274+00',	NULL,	1,	2,	'f'),
(6,	'2024-12-01 00:03:11.986789+00',	'2024-12-01 00:03:11.986789+00',	NULL,	1,	4,	'f'),
(7,	'2024-12-01 00:03:23.757001+00',	'2024-12-01 00:03:23.757001+00',	NULL,	1,	6,	'f'),
(8,	'2024-12-01 00:03:33.654628+00',	'2024-12-01 00:03:33.654628+00',	NULL,	1,	8,	'f'),
(9,	'2024-12-01 00:03:38.000312+00',	'2024-12-01 00:03:38.000312+00',	NULL,	1,	9,	'f'),
(10,	'2024-12-01 00:05:02.60685+00',	'2024-12-01 00:05:02.60685+00',	NULL,	1,	11,	't'),
(11,	'2024-12-01 00:25:44.196168+00',	'2024-12-01 00:25:44.196168+00',	NULL,	1,	12,	't');

DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."users" (
    "id" bigint DEFAULT nextval('users_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "username" text NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "gender" text NOT NULL,
    "bio" text,
    "date_of_birth" timestamptz,
    "location" bigint,
    "preferance_id" bigint,
    "profile_picture" text,
    "premium" timestamptz,
    "age" bigint,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_users_deleted_at" ON "public"."users" USING btree ("deleted_at");

INSERT INTO "users" ("id", "created_at", "updated_at", "deleted_at", "username", "email", "password", "gender", "bio", "date_of_birth", "location", "preferance_id", "profile_picture", "premium", "age") VALUES
(3,	'2024-11-30 17:39:25.601529+00',	'2024-11-30 17:39:25.601529+00',	NULL,	'cihuy3',	'cihuy@wow.com',	'$2a$14$kPXXhinKvkYnzth3KtmGd.4uu5hIaC3Zm9hnVNR1rfyteYxpV6Iu6',	'Moster Detonator',	'',	NULL,	0,	3,	'',	NULL,	23),
(2,	'2024-11-30 17:38:32.862351+00',	'2024-11-30 17:38:32.862351+00',	NULL,	'cihuy1',	'cihuy@wow.com',	'$2a$14$KPOAz8xgwJoa2uEvxI9PreR/DjroL9rkCwwd6KtbNLc4VnYptTfx6',	'Moster Detonator',	'',	NULL,	0,	2,	'',	NULL,	42),
(4,	'2024-11-30 17:41:54.59411+00',	'2024-11-30 17:41:54.59411+00',	NULL,	'cihuy5',	'cihuy@wow.com',	'$2a$14$0LzxVJZOloxlq9HUgrXHEeAMwCiunJTI03G5x/9BYLhdlDDDOoqC6',	'Moster Detonator',	'',	NULL,	0,	4,	'',	NULL,	25),
(5,	'2024-11-30 17:42:32.734978+00',	'2024-11-30 17:42:32.734978+00',	NULL,	'cihuy6',	'cihuy@wow.com',	'$2a$14$8a0.MaWqv1Q1jZO4uk3dtuhNefUKolqUcQjD.7Or6iDM05LRfrSHu',	'Moster Detonator',	'',	NULL,	0,	5,	'',	NULL,	25),
(6,	'2024-11-30 17:43:35.444537+00',	'2024-11-30 17:43:35.444537+00',	NULL,	'cihuy7',	'cihuy@wow.com',	'$2a$14$4iLMiSx8N5iS5qG.TVH8oOQcwDIhweiAePlK.5oT0kFkKA/wtSuD.',	'Moster Detonator',	'',	NULL,	0,	6,	'',	NULL,	25),
(7,	'2024-11-30 17:43:48.860395+00',	'2024-11-30 17:43:48.860395+00',	NULL,	'cihuy8',	'cihuy@wow.com',	'$2a$14$oFt/j05Djl4vooxIvLV9YO6HOy0pPVnjotM5lmFhutmA/QDH//kCa',	'Moster Detonator',	'',	NULL,	0,	7,	'',	NULL,	25),
(8,	'2024-11-30 17:43:55.707869+00',	'2024-11-30 17:43:55.707869+00',	NULL,	'cihuy9',	'cihuy@wow.com',	'$2a$14$XF1EmofjPaTCi2yWvAP.x.9ZF/avgSuWRh.Mot7W5uWsTkIbbCfoe',	'Moster Detonator',	'',	NULL,	0,	8,	'',	NULL,	25),
(9,	'2024-11-30 17:43:59.851722+00',	'2024-11-30 17:43:59.851722+00',	NULL,	'cihuy10',	'cihuy@wow.com',	'$2a$14$8/1tK..nENcK82O0oOD.Ke3YrxyO7Y6mu0Sy0BRGmeW.756Q6V28C',	'Moster Detonator',	'',	NULL,	0,	9,	'',	NULL,	25),
(10,	'2024-11-30 17:44:05.433672+00',	'2024-11-30 17:44:05.433672+00',	NULL,	'cihuy11',	'cihuy@wow.com',	'$2a$14$DuhRaYeCJH3Te/j.HDTx0.0FgFqD4/se3azosTo7xbzgVksmZrYUi',	'Moster Detonator',	'',	NULL,	0,	10,	'',	NULL,	25),
(11,	'2024-12-01 00:04:08.902358+00',	'2024-12-01 00:04:08.902358+00',	NULL,	'cihuy11',	'cihuy@wow.com',	'$2a$14$wv04Em81UktFpD3Cx3fNbOh4eXfnqL4LnaDbG0Te1fWZWDqNRmmV.',	'Moster Detonator',	'',	NULL,	0,	11,	'',	NULL,	25),
(12,	'2024-12-01 00:04:15.204333+00',	'2024-12-01 00:04:15.204333+00',	NULL,	'cihuy12',	'cihuy@wow.com',	'$2a$14$kLZldLvNtU3nN.A1v8434uMZjgEepFf856LkI/O5ih9F/wZ5DHVp6',	'Moster Detonator',	'',	NULL,	0,	12,	'',	NULL,	25),
(13,	'2024-12-01 00:04:20.496878+00',	'2024-12-01 00:04:20.496878+00',	NULL,	'cihuy13',	'cihuy@wow.com',	'$2a$14$frvizhSp31p0pDUnsUVRau8nAngwNIEsyrKhGXXmNlW.3ZQjr02fW',	'Moster Detonator',	'',	NULL,	0,	13,	'',	NULL,	25),
(14,	'2024-12-01 00:04:26.243697+00',	'2024-12-01 00:04:26.243697+00',	NULL,	'cihuy14',	'cihuy@wow.com',	'$2a$14$6PfpJZolENKZKAvTPE4CDOIytOyJc/BGikdvp4DasuK0ywrQfqfA.',	'Moster Detonator',	'',	NULL,	0,	14,	'',	NULL,	25),
(15,	'2024-12-01 00:04:31.033887+00',	'2024-12-01 00:04:31.033887+00',	NULL,	'cihuy15',	'cihuy@wow.com',	'$2a$14$UzjGj2hUWUOW0T7q8jMQeOwDLAh1csQlgBBEx1X3aTIUnPH0x/jBu',	'Moster Detonator',	'',	NULL,	0,	15,	'',	NULL,	25),
(16,	'2024-12-01 00:04:35.690032+00',	'2024-12-01 00:04:35.690032+00',	NULL,	'cihuy16',	'cihuy@wow.com',	'$2a$14$klhunQdhwU29/O4LWw5g..hIDmeJCxyhQi11k4CPts3da9Lruudwy',	'Moster Detonator',	'',	NULL,	0,	16,	'',	NULL,	25),
(17,	'2024-12-01 00:29:23.447006+00',	'2024-12-01 00:29:23.447006+00',	NULL,	'cihuy',	'cihuy@wow.com',	'$2a$14$fTwP/MTd5KZB30xAbruYfOupgvDrWaphtLlV7o3U2vkYRnlTPZcpS',	'Apache Helicopter',	'',	NULL,	0,	17,	'',	NULL,	0),
(18,	'2024-12-01 02:37:33.640918+00',	'2024-12-01 02:37:33.640918+00',	NULL,	'cihuy16',	'cihuy@wow.com',	'$2a$14$t52yUGeOz2NTIUiOwtjBOuuRq.IptCaTKEgA9cVtHtxvDLu8ajIFy',	'Moster Detonator',	'',	NULL,	0,	18,	'',	NULL,	25),
(19,	'2024-12-01 02:56:23.23924+00',	'2024-12-01 02:56:23.23924+00',	NULL,	'cihuy16',	'cihuy@wow.com',	'$2a$14$3qg9txO7EJ6dZElg7G1qcuP8IOdB7Lh0ABl13GqRuTidJplj/4h4a',	'Moster Detonator',	'',	NULL,	0,	19,	'',	NULL,	25),
(1,	'2024-11-30 14:33:40.554777+00',	'2024-12-01 03:03:57.297748+00',	NULL,	'cihuy',	'cihuy@wow.com',	'$2a$14$VE12w1kJrh4att.Vq3G/qu9YeEalMyDA/0/i9ITtNpGWEAH9kwoHu',	'Apache Helicopter',	'',	NULL,	0,	1,	'',	'2024-12-01 03:01:41.136619+00',	32);

ALTER TABLE ONLY "public"."pictures" ADD CONSTRAINT "fk_users_pictures" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."users" ADD CONSTRAINT "fk_users_preferance" FOREIGN KEY (preferance_id) REFERENCES preferances(id) ON DELETE CASCADE NOT DEFERRABLE;

-- 2024-12-01 05:35:21.012191+00
