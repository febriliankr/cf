-- -------------------------------------------------------------
-- TablePlus 4.7.1(428)
--
-- https://tableplus.com/
--
-- Database: cfstore
-- Generation Time: 2022-07-10 18:32:06.8680
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS canteen_balance_canteen_balance_id_seq;

-- Table Definition
CREATE TABLE "public"."canteen_balance" (
    "canteen_balance_id" int4 NOT NULL DEFAULT nextval('canteen_balance_canteen_balance_id_seq'::regclass),
    "amount" int8,
    "type" varchar,
    "student_id" int8,
    PRIMARY KEY ("canteen_balance_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."product" (
    "product_id" uuid DEFAULT uuid_generate_v4(),
    "name" varchar,
    "description" varchar,
    "product_slug" varchar,
    "price" int8,
    "image_url" varchar,
    "student_id" varchar,
    "created_at" timestamptz DEFAULT now(),
    "hidden" bool NOT NULL DEFAULT false
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS user_purchase_user_purchase_id_seq;

-- Table Definition
CREATE TABLE "public"."user_purchase" (
    "user_purchase_id" int4 NOT NULL DEFAULT nextval('user_purchase_user_purchase_id_seq'::regclass),
    "student_id" int8,
    "product_slug" varchar,
    PRIMARY KEY ("user_purchase_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."users" (
    "student_id" int8,
    "password" varchar,
    "created_at" timestamptz DEFAULT now()
);

INSERT INTO "public"."canteen_balance" ("canteen_balance_id", "amount", "type", "student_id") VALUES
(3, 1100000, 'deposit', 12306),
(5, 5000, 'deposit', 12306),
(6, 5000, 'deposit', 12306),
(7, 5000, 'deposit', 12306),
(8, 5000, 'deposit', 12306),
(9, 5000, 'deposit', 12306),
(10, 5000, 'deposit', 12306),
(11, 5000, 'deposit', 12306),
(12, 5000, 'deposit', 12306),
(13, 5000, 'deposit', 12306),
(14, 5000, 'deposit', 12306),
(15, 5000, 'deposit', 12306),
(16, 5000, 'deposit', 12306),
(17, 5000, 'deposit', 12306),
(18, 5000, 'deposit', 12306),
(19, 5000, 'deposit', 12306),
(20, 5000, 'deposit', 12306),
(21, 5000, 'withdraw', 12306),
(22, 5000, 'withdraw', 12306),
(23, 5000, 'deposit', 12306),
(24, 5000, 'deposit', 12306),
(25, 5000, 'withdraw', 12306),
(26, 5000, 'withdraw', 12306),
(27, 5000, 'withdraw', 12306),
(28, 5000, 'withdraw', 12306),
(29, 5000, 'withdraw', 12306),
(30, 5000, 'withdraw', 12306),
(31, 5000, 'withdraw', 12306),
(32, 5000, 'withdraw', 12306),
(33, 5000, 'withdraw', 12306),
(34, 5000, 'withdraw', 12306),
(35, 5000, 'withdraw', 12306),
(36, 5000, 'withdraw', 12306),
(37, 5000, 'withdraw', 12306),
(38, 5000, 'withdraw', 12306),
(39, 5000, 'withdraw', 12306),
(40, 5000, 'withdraw', 12306),
(41, 5000, 'withdraw', 12306),
(42, 5000, 'withdraw', 12306),
(43, 5000, 'withdraw', 12306),
(44, 5000, 'withdraw', 12306),
(45, 5000, 'withdraw', 12306),
(46, 5000, 'withdraw', 12306),
(47, 5000, 'withdraw', 12306),
(48, 5000, 'withdraw', 12306),
(49, 5000, 'withdraw', 12306),
(50, 5000, 'withdraw', 12306),
(51, 5000, 'withdraw', 12306),
(52, 5000, 'withdraw', 12306),
(53, 5000, 'withdraw', 12306),
(54, 5000, 'withdraw', 12306),
(55, 5000, 'withdraw', 12306),
(56, 5000, 'withdraw', 12306),
(57, 5000, 'withdraw', 12306),
(58, 5000, 'withdraw', 12306),
(59, 5000, 'withdraw', 12306),
(60, 5000, 'withdraw', 12306),
(61, 5000, 'withdraw', 12306),
(62, 5000, 'withdraw', 12306),
(63, 5000, 'withdraw', 12306),
(64, 5000, 'withdraw', 12306),
(65, 5000, 'deposit', 12306),
(66, 5000, 'withdraw', 12306),
(67, 5000, 'withdraw', 12306),
(68, 5000, 'withdraw', 12306),
(69, 5000, 'deposit', 12306),
(70, 5000, 'deposit', 12306),
(71, 5000, 'withdraw', 12306),
(72, 5000, 'withdraw', 12306),
(73, 5000, 'deposit', 12306),
(74, 5000, 'deposit', 12306),
(75, 5000, 'withdraw', 12306),
(76, 5000, 'deposit', 12306),
(77, 5000, 'withdraw', 12306),
(78, 5000, 'deposit', 12306),
(79, 5000, 'deposit', 12306),
(80, 5000, 'withdraw', 12306);

INSERT INTO "public"."product" ("product_id", "name", "description", "product_slug", "price", "image_url", "student_id", "created_at", "hidden") VALUES
('7a5044ab-ac86-4096-8926-9203f759982a', 'Toast', 'Delitoast', 'toast_a332b729-7b0f-43f5-a10e-fe6ad9b88f19', 199900, 'https://skala-files.s3.ap-southeast-1.amazonaws.com/compfest_user/91d0d0b58cee4625815aa45b4e2ec975_Toast-3.jpeg.jpeg', '12407', '2022-07-10 16:36:20.885405+07', 't'),
('b7faa690-dc75-4eea-82b1-d65751dcad5b', 'kambing', 'kambing product', 'kambing_64924bea-293f-4ca0-a0a0-9bcd835bf589', 10000000, 'https://skala-files.s3.ap-southeast-1.amazonaws.com/compfest_user/5dac93e5d0514bbcb14b77e9f8664fc4_Risetku Source 7 (1).png.png', '12407', '2022-07-10 14:57:27.356424+07', 't');

INSERT INTO "public"."user_purchase" ("user_purchase_id", "student_id", "product_slug") VALUES
(1, 12407, 'kambing_64924bea-293f-4ca0-a0a0-9bcd835bf589'),
(2, 12407, 'kambing_64924bea-293f-4ca0-a0a0-9bcd835bf589'),
(3, 12407, 'toast_a332b729-7b0f-43f5-a10e-fe6ad9b88f19'),
(4, 12306, 'kambing_64924bea-293f-4ca0-a0a0-9bcd835bf589');

INSERT INTO "public"."users" ("student_id", "password", "created_at") VALUES
(12407, '$2a$10$sP8tkmdyTHhI7nqs8svYH.RLjvoJ.BQDRqyIoFukvEW0uVBrSZOYi', '2022-07-09 20:51:33.982635+07'),
(11103, '$2a$10$hre/S1vcZT7cKpXUAvYrUOpLojNKI5hhEykcQf8OryMrjq/kbgHau', '2022-07-10 16:37:58.912291+07'),
(12306, '$2a$10$/aofIlA630Jh9yXGEJdrLufRkuRVVmChI0TUo2ZWypDMTvZCUSBIe', '2022-07-10 16:55:19.67782+07');

