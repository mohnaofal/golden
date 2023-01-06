-- create table buyback
CREATE TABLE buyback (
	buyback_id serial,
	buyback_gram float NOT NULL DEFAULT 0,
	buyback_harga int NOT NULL DEFAULT 0,
	buyback_norek varchar(20) NOT NULL DEFAULT '',
	buyback_date int NOT NULL DEFAULT 0,
	PRIMARY KEY (buyback_id)
);
CREATE INDEX buyback_buyback_date_idx ON buyback (buyback_date);

-- create table harga
CREATE TABLE harga (
	harga_id serial,
	harga_admin_id varchar(20) NOT NULL DEFAULT '',
	harga_topup int NOT NULL DEFAULT 0,
	harga_buyback int NOT NULL DEFAULT 0,
	harga_date int NOT NULL DEFAULT 0,
	PRIMARY KEY (harga_id)
);
CREATE INDEX harga_harga_date_idx ON harga (harga_date);

-- create table topup
CREATE TABLE topup (
	topup_id serial,
	topup_gram float NOT NULL DEFAULT 0,
	topup_harga int NOT NULL DEFAULT 0,
	topup_norek varchar(20) NOT NULL DEFAULT '',
	topup_date int NOT NULL DEFAULT 0,
	topup_pkey PRIMARY KEY (topup_id)
);
CREATE INDEX topup_norek_idx ON topup (topup_norek);
CREATE INDEX topup_topup_date_idx ON topup (topup_date);

-- create table rekening
CREATE TABLE rekening (
	rek_id serial,
	rek_norek varchar(20) NOT NULL DEFAULT '',
	rek_saldo float NOT NULL DEFAULT 0,
	rek_updated_at int NULL,
	PRIMARY KEY (rek_id)
);
CREATE UNIQUE INDEX rekening_norek_unique ON rekening (rek_norek);

-- create table transaksi
CREATE TABLE transaksi (
	trx_id serial,
	trx_date int NOT NULL DEFAULT 0,
	trx_type varchar(20) NOT NULL DEFAULT '',
	trx_gram float NULL DEFAULT 0,
	trx_harga_topup int NOT NULL DEFAULT 0,
	trx_harga_buyback int NOT NULL DEFAULT 0,
	trx_saldo float8 NOT NULL DEFAULT 0,
	trx_norek varchar(20) NOT NULL DEFAULT '',
	PRIMARY KEY (trx_id)
);
CREATE INDEX trx_date_idx ON transaksi (trx_date);
CREATE INDEX trx_norek_idx ON transaksi (trx_norek);