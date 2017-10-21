CREATE TABLE order_fact(
  order_id INTEGER,
  uuid CHAR(36),
  type SMALLINT,
  payment_type SMALLINT,
  tracking_uuid CHAR(36),
  offer_id INTEGER DEFAULT 0,
  user_id INTEGER DEFAULT 0,
  client_id INTEGER DEFAULT 0,
  status SMALLINT,
  status_partner SMALLINT,
  payment FLOAT DEFAULT 0,
  sum FLOAT DEFAULT 0,

  redirect_id INTEGER,
  CONSTRAINT order_fact_pk PRIMARY KEY (order_id)
);

CREATE INDEX order_fact_uuid ON order_fact (uuid);
CREATE INDEX order_fact_offer_id ON order_fact (offer_id);
CREATE INDEX order_fact_user_id ON order_fact (user_id);
CREATE INDEX order_fact_client_id ON order_fact (client_id);
CREATE INDEX order_fact_status ON order_fact (status);
CREATE INDEX order_fact_status_partner ON order_fact (status_partner);
CREATE INDEX order_fact_redirect_id ON order_fact (redirect_id);

CREATE INDEX order_fact_user_status ON order_fact (user_id, status_partner);

CREATE TABLE order_time(
  order_id INTEGER,
  time TIMESTAMP,
  day_id SMALLINT,
  month_id SMALLINT,
  quarter_id SMALLINT,
  year_id SMALLINT,
  CONSTRAINT order_time_pk PRIMARY KEY (order_id)
);

CREATE TABLE order_crm(
  order_id SERIAL,
  external_id INTEGER,
  crm_id SMALLINT,
  CONSTRAINT order_crm_pk PRIMARY KEY (order_id)
);

CREATE TABLE client(
  client_id INTEGER,
  phone_geo_code CHAR(2),
  age SMALLINT,
  gender CHAR(1),
  region CHAR(6),
  CONSTRAINT client_pk PRIMARY KEY (client_id)
);

CREATE TABLE redirect_fact(
  id SERIAL,
  uuid CHAR(36),
  unique_id INTEGER,
  landing_id INTEGER,
  prelanding_id INTEGER,
  stream_id INTEGER,
  device_is_mobile BOOLEAN DEFAULT FALSE,
  location_geo_code CHAR(2),
  sub1_id INTEGER,
  sub2_id INTEGER,
  sub3_id INTEGER,
  sub4_id INTEGER,
  sub5_id INTEGER,
  CONSTRAINT redirect_fact_pk PRIMARY KEY (id)
);

CREATE INDEX redirect_fact_uuid ON redirect_fact (uuid);

CREATE TABLE view_fact(
  id SERIAL,
  uuid CHAR(36),
  CONSTRAINT view_fact_pk PRIMARY KEY (id)
);

CREATE INDEX view_fact_uuid ON view_fact (uuid);

CREATE TABLE sub(
  id SERIAL,
  name VARCHAR(255),
  type SMALLINT,
  CONSTRAINT sub_pk PRIMARY KEY (id)
);

CREATE TABLE sub_user(
  sub_id INTEGER,
  user_id INTEGER,
  CONSTRAINT sub_user_pk PRIMARY KEY (sub_id, user_id)
);