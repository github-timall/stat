CREATE TABLE order_fact(
  id SERIAL,
  uuid CHAR(36),
  tracking_uuid CHAR(36),
  offer_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  client_id INTEGER NOT NULL,
  status SMALLINT,
  status_partner SMALLINT,
  payment money,
  sum money,
  time_id INTEGER NOT NULL,
  CONSTRAINT order_fact_pk PRIMARY KEY (id)
);

CREATE INDEX order_fact_uuid ON order_fact (uuid);
CREATE INDEX order_fact_offer_id ON order_fact (offer_id);
CREATE INDEX order_fact_user_id ON order_fact (user_id);
CREATE INDEX order_fact_client_id ON order_fact (client_id);
CREATE INDEX order_fact_status ON order_fact (status);
CREATE INDEX order_fact_status_partner ON order_fact (status_partner);
CREATE INDEX order_fact_time_id ON order_fact (time_id);

CREATE INDEX order_fact_user_status ON order_fact (user_id, status_partner);

CREATE TABLE order_time(
  id SERIAL,
  time TIMESTAMP,
  day_id SMALLINT,
  month_id SMALLINT,
  quarter_id SMALLINT,
  year_id SMALLINT,
  CONSTRAINT order_time_pk PRIMARY KEY (id)
);

CREATE TABLE client(
  client_id INTEGER,
  phone_geo_code CHAR(2),
  age SMALLINT,
  gender CHAR(1),
  region CHAR(6),
  CONSTRAINT client_pk PRIMARY KEY (client_id)
);

CREATE TABLE tracking(
  uuid CHAR(36),
  unique_id INTEGER,
  landing_id INTEGER,
  prelanding_id INTEGER,
  stream_id INTEGER,
  is_mobile BOOLEAN,
  geo_code CHAR(2),
  sub_1_id INTEGER,
  sub_2_id INTEGER,
  sub_3_id INTEGER,
  sub_4_id INTEGER,
  sub_5_id INTEGER,
  CONSTRAINT tracking_pk PRIMARY KEY (uuid)
);