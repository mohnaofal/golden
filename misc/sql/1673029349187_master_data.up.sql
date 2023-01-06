INSERT INTO rekening (rek_norek, rek_updated_at)
    SELECT
    	'r001', 
    	(SELECT EXTRACT(EPOCH FROM  now()))
;