CREATE TABLE IF NOT EXISTS `order_fact` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uuid` CHAR(36),
  `type` SMALLINT,
  `payment_type` SMALLINT,
  `tracking_uuid` CHAR(36),

  `offer_id` INTEGER DEFAULT 0,
  `user_id` INTEGER DEFAULT 0,
  `client_id` INTEGER DEFAULT 0,
  `redirect_id` INTEGER,
  `time_id` INTEGER,

  `status` SMALLINT,
  `status_partner` SMALLINT,
  `payment` FLOAT DEFAULT 0,
  `sum` FLOAT DEFAULT 0,

  `time` TIMESTAMP,

  PRIMARY KEY (id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `order_time` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `time` TIMESTAMP,
  `day_id` SMALLINT,
  `month_id` SMALLINT,
  `quarter_id` SMALLINT,
  `year_id` SMALLINT,
  PRIMARY KEY (id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `client`(
  `id` INTEGER,
  `uuid` CHAR(36),
  `phone_geo_code` CHAR(2),
  `age` SMALLINT,
  `gender` CHAR(1),
  `region` CHAR(6),
  PRIMARY KEY (id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `redirect_fact` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uuid` CHAR(36),
  `unique_id` INTEGER,
  `landing_id` INTEGER,
  `prelanding_id` INTEGER,
  `stream_id` INTEGER,
  `device_is_mobile` BOOLEAN DEFAULT FALSE,
  `location_geo_code` CHAR(2),
  `sub1_id` INTEGER,
  `sub2_id` INTEGER,
  `sub3_id` INTEGER,
  `sub4_id` INTEGER,
  `sub5_id` INTEGER,
  PRIMARY KEY (id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `type` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, type)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub_user` (
  `sub_id` INTEGER,
  `user_id` INTEGER,
  PRIMARY KEY (sub_id, user_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub1` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `affiliate_id` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, affiliate_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub2` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `affiliate_id` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, affiliate_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub3` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `affiliate_id` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, affiliate_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub4` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `affiliate_id` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, affiliate_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

CREATE TABLE IF NOT EXISTS `sub5` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `affiliate_id` SMALLINT,
  PRIMARY KEY (id),
  UNIQUE (name, affiliate_id)
) ENGINE='InnoDB' COLLATE 'utf8_unicode_ci';

INSERT IGNORE INTO sub1 (`name`, `affiliate_id`) SELECT `name`, `affiliate_id` FROM `tracking_sub1`;
INSERT IGNORE INTO sub2 (`name`, `affiliate_id`) SELECT `name`, `affiliate_id` FROM `tracking_sub2`;
INSERT IGNORE INTO sub3 (`name`, `affiliate_id`) SELECT `name`, `affiliate_id` FROM `tracking_sub3`;
INSERT IGNORE INTO sub4 (`name`, `affiliate_id`) SELECT `name`, `affiliate_id` FROM `tracking_sub4`;
INSERT IGNORE INTO sub5 (`name`, `affiliate_id`) SELECT `name`, `affiliate_id` FROM `tracking_sub5`;

DELETE FROM `sub1` WHERE `affiliate_id` = 1000;
DELETE FROM `sub2` WHERE `affiliate_id` = 1000;
DELETE FROM `sub3` WHERE `affiliate_id` = 1000;
DELETE FROM `sub4` WHERE `affiliate_id` = 1000;
DELETE FROM `sub5` WHERE `affiliate_id` = 1000;

INSERT IGNORE INTO sub (`name`, `type`) SELECT `name`, 1 FROM `sub1`;
INSERT IGNORE INTO sub (`name`, `type`) SELECT `name`, 2 FROM `sub2`;
INSERT IGNORE INTO sub (`name`, `type`) SELECT `name`, 3 FROM `sub3`;
INSERT IGNORE INTO sub (`name`, `type`) SELECT `name`, 4 FROM `sub4`;
INSERT IGNORE INTO sub (`name`, `type`) SELECT `name`, 5 FROM `sub5`;

INSERT IGNORE INTO `sub_user` (`sub_id`, `user_id`) SELECT (SELECT `s`.`id` FROM `sub` `s` WHERE `s`.`name` = `t`.`name` AND `s`.`type` = 1) as `sub_id`, `affiliate_id` FROM `sub1` `t`;
INSERT IGNORE INTO `sub_user` (`sub_id`, `user_id`) SELECT (SELECT `s`.`id` FROM `sub` `s` WHERE `s`.`name` = `t`.`name` AND `s`.`type` = 2) as `sub_id`, `affiliate_id` FROM `sub2` `t`;
INSERT IGNORE INTO `sub_user` (`sub_id`, `user_id`) SELECT (SELECT `s`.`id` FROM `sub` `s` WHERE `s`.`name` = `t`.`name` AND `s`.`type` = 3) as `sub_id`, `affiliate_id` FROM `sub3` `t`;
INSERT IGNORE INTO `sub_user` (`sub_id`, `user_id`) SELECT (SELECT `s`.`id` FROM `sub` `s` WHERE `s`.`name` = `t`.`name` AND `s`.`type` = 4) as `sub_id`, `affiliate_id` FROM `sub4` `t`;
INSERT IGNORE INTO `sub_user` (`sub_id`, `user_id`) SELECT (SELECT `s`.`id` FROM `sub` `s` WHERE `s`.`name` = `t`.`name` AND `s`.`type` = 5) as `sub_id`, `affiliate_id` FROM `sub5` `t`;

ALTER TABLE `sub_user` ADD INDEX `user_id` (`user_id`);

ALTER TABLE `order_fact` ADD `lead_id` INTEGER;
ALTER TABLE `order_fact` ADD INDEX `lead_id` (`lead_id`);
ALTER TABLE `redirect_fact` ADD `lead_id` INTEGER;
ALTER TABLE `redirect_fact` ADD INDEX `lead_id` (`lead_id`);
ALTER TABLE `redirect_fact` ADD `user_id` INTEGER;
ALTER TABLE `redirect_fact` ADD `sub1_name` VARCHAR(255);
ALTER TABLE `redirect_fact` ADD `sub2_name` VARCHAR(255);
ALTER TABLE `redirect_fact` ADD `sub3_name` VARCHAR(255);
ALTER TABLE `redirect_fact` ADD `sub4_name` VARCHAR(255);
ALTER TABLE `redirect_fact` ADD `sub5_name` VARCHAR(255);

INSERT INTO order_fact (`uuid`,`type`,`payment_type`,`tracking_uuid`,`offer_id`,`user_id`,`client_id`,`status`,`status_partner`,`payment`,`sum`,`redirect_id`,`lead_id`,`time`) SELECT
    `n_id_lead`,
    (SELECT
         CASE
           WHEN `method` = 1 THEN 1
           WHEN `method` = 2 THEN 1
           WHEN `method` = 3 THEN 4
           WHEN `method` = 4 THEN 2
           WHEN `method` = 5 THEN 3
           ELSE 0
         END
    ) AS `type`,
    1,
    '',
    `offer_id`,
    `affiliate_id`,
    `n_id_lead`,
    `status`,
    `status_partner`,
    `payment`,
    `sum`,
    0,
    `n_id_lead`,
    `c_time`
   FROM `tracking_leads`;

INSERT INTO client (`id`,`phone_geo_code`,`age`,`gender`,`region`) SELECT `n_id_lead`,`client_phone_geo`,`client_age`,`client_gender`,`client_region` FROM `tracking_leads`;
INSERT INTO redirect_fact (
  `unique_id`,
  `landing_id`,
  `prelanding_id`,
  `stream_id`,
  `device_is_mobile`,
  `location_geo_code`,
  `lead_id`,
  `sub1_name`,
  `sub2_name`,
  `sub3_name`,
  `sub4_name`,
  `sub5_name`,
  `user_id`
  ) SELECT
      `tl`.`unique_id`,
      `tl`.`landing_id`,
      `tl`.`prelanding_id`,
      `tl`.`stream_id`,
      `tl`.`is_mobile`,
      `tl`.`geo_code`,
      `tl`.`n_id_lead`,
      IF(`sub1_id` > 0, (SELECT `s1`.`name` FROM `tracking_sub1` `s1` WHERE `s1`.`id` = `tl`.`sub1_id`), NULL),
      IF(`sub2_id` > 0, (SELECT `s2`.`name` FROM `tracking_sub2` `s2` WHERE `s2`.`id` = `tl`.`sub2_id`), NULL),
      IF(`sub3_id` > 0, (SELECT `s3`.`name` FROM `tracking_sub3` `s3` WHERE `s3`.`id` = `tl`.`sub3_id`), NULL),
      IF(`sub4_id` > 0, (SELECT `s4`.`name` FROM `tracking_sub4` `s4` WHERE `s4`.`id` = `tl`.`sub4_id`), NULL),
      IF(`sub5_id` > 0, (SELECT `s5`.`name` FROM `tracking_sub5` `s5` WHERE `s5`.`id` = `tl`.`sub5_id`), NULL),
      `tl`.`affiliate_id`
    FROM `tracking_leads` `tl`;

UPDATE `order_fact` `o` SET `o`.`redirect_id` = (SELECT `r`.`id` FROM `redirect_fact` `r` WHERE `r`.`lead_id` = `o`.`lead_id`);

UPDATE `redirect_fact` `r` SET `r`.`sub1_id` = (
SELECT `s`.`id` FROM (
  SELECT * FROM `sub` `ss` LEFT JOIN `sub_user` `su` ON `su`.`sub_id` = `ss`.`id` WHERE `ss`.`type` = 1
) as `s` WHERE `s`.`name` = `r`.`sub1_name` AND `s`.`user_id` = `r`.`user_id`);

UPDATE `redirect_fact` `r` SET `r`.`sub12_id` = (
SELECT `s`.`id` FROM (
SELECT * FROM `sub` `ss` LEFT JOIN `sub_user` `su` ON `su`.`sub_id` = `ss`.`id` WHERE `ss`.`type` = 2
) as `s` WHERE `s`.`name` = `r`.`sub2_name` AND `s`.`user_id` = `r`.`user_id`);

UPDATE `redirect_fact` `r` SET `r`.`sub3_id` = (
SELECT `s`.`id` FROM (
SELECT * FROM `sub` `ss` LEFT JOIN `sub_user` `su` ON `su`.`sub_id` = `ss`.`id` WHERE `ss`.`type` = 3
) as `s` WHERE `s`.`name` = `r`.`sub3_name` AND `s`.`user_id` = `r`.`user_id`);

UPDATE `redirect_fact` `r` SET `r`.`sub4_id` = (
SELECT `s`.`id` FROM (
SELECT * FROM `sub` `ss` LEFT JOIN `sub_user` `su` ON `su`.`sub_id` = `ss`.`id` WHERE `ss`.`type` = 4
) as `s` WHERE `s`.`name` = `r`.`sub4_name` AND `s`.`user_id` = `r`.`user_id`);

UPDATE `redirect_fact` `r` SET `r`.`sub5_id` = (
SELECT `s`.`id` FROM (
SELECT * FROM `sub` `ss` LEFT JOIN `sub_user` `su` ON `su`.`sub_id` = `ss`.`id` WHERE `ss`.`type` = 5
) as `s` WHERE `s`.`name` = `r`.`sub5_name` AND `s`.`user_id` = `r`.`user_id`);


ALTER TABLE `order_fact` DROP COLUMN `lead_id`;