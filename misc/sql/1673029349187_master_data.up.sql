-- insert rekening
INSERT INTO rekening (rek_norek, rek_saldo, rek_updated_at)
    SELECT
    	'r001',
		1,
    	(SELECT EXTRACT(EPOCH FROM  now()))
;

-- insert harga
INSERT INTO harga (harga_admin_id, harga_topup, harga_buyback, harga_date)
	SELECT
		'a001',
		100000,
		90000,
		(SELECT EXTRACT(EPOCH FROM  now()))
;

INSERT INTO transaksi (trx_date, trx_type, trx_gram, trx_harga_topup, trx_harga_buyback, trx_saldo, trx_norek)
	SELECT
		(SELECT EXTRACT(EPOCH FROM  now())),
		'topup',
		1,
		100000,
		90000,
		2, 
		'r001'
;

INSERT INTO transaksi (trx_date, trx_type, trx_gram, trx_harga_topup, trx_harga_buyback, trx_saldo, trx_norek)
	SELECT
		(SELECT EXTRACT(EPOCH FROM  now())),
		'buyback',
		1,
		100000,
		90000,
		1, 
		'r001'
;
