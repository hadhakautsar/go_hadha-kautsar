-- 1. INSERT
-- a
insert into operator (nama)
values 
	('Puff'),
	('Yeti'),
	('Ray Chargels'),
	('Honey Bear')
	('Hadha Kautsar');
-- b
INSERT INTO product_type (nama)
VALUES 
('Alat Tulis'),
('Minuman'),
('Makanan');
-- c
INSERT INTO product (product_type_id, operator_id, code, nama, status)
VALUES 
(1, 3, 'RC1','Pensil',1),
(1, 3, 'PF1','Penghapus',1);
-- d
INSERT INTO product (product_type_id, operator_id, code, nama, status)
VALUES 
(2, 1, 'HB1','Jus Buah',1),
(2, 1, 'YT1','Soda',1),
(2, 1, 'RC2','Air Tawar',1);
-- e
INSERT INTO product (product_type_id, operator_id, code, nama, status)
VALUES 
(3, 4, 'HB2','Daging',1),
(3, 4, 'RC3','Ikan',1),
(3, 4, 'PF2','Ayam',1);
-- f
INSERT INTO product_descriptions (description)
VALUES
('Berbagai macam pensil'),
('Berbagai macam penghapus'),
('Berbagai macam jus buah'),
('Berbagai macam soda'),
('Berbagai macam air tawar'),
('Berbagai macam daging'),
('Berbagai macam ikan'),
('Berbagai macam ayam');
-- g
INSERT INTO payment_method (nama, status, payment_method_description_id)
VALUES ('Cash', 1, 1),
       ('Bank Transfer', 1, 2),
       ('Credit Card', 1, 3);
-- h
INSERT INTO user (nama, alamat_id, tanggal_lahir, status, gender)
VALUES
('Maulana Ibrahim', 1, '1990-05-15', 1, 'male'),
('Susana Listiyani', 2, '1995-07-20', 1, 'female'),
('Eko Prasetyanto', 3, '1985-03-12', 1, 'male'),
('Liliana Anggita ', 4, '1992-12-28', 1, 'female'),
('Bambang Lukman', 5, '1988-08-08', 1, 'male');
INSERT INTO alamat (alamat_id, alamat)
values
(1, 'Jalan Raya Tiga No. 123'),
(2, 'Jalan Raya Dua No. 456'),
(3, 'Jalan Raya Satu No. 789'),
(4, 'Jalan Raya Enam No. 246'),
(5, 'Jalan Raya Lima No. 135');
-- i
INSERT INTO transaction (user_id, payment_method_id, status, total_qty, total_price)
VALUES 
(1, 1, 'pending', 2, 150000),
(1, 2, 'success', 3, 250000),
(1, 1, 'canceled', 1, 75000),
(2, 2, 'success', 5, 9000000),
(2, 1, 'pending', 1, 100000),
(2, 2, 'pending', 3, 200000),
(3, 1, 'canceled', 2, 200000),
(3, 2, 'success', 4, 2500000),
(3, 1, 'pending', 1, 50000),
(4, 1, 'canceled', 2, 400000),
(4, 2, 'success', 4, 4500000),
(4, 1, 'pending', 1, 5000000),
(5, 1, 'canceled', 2, 600000),
(5, 2, 'success', 4, 5500000),
(5, 1, 'pending', 1, 900000);
-- j
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES 
(1, 1, 'canceled', 1, 50000),
(1, 2, 'success', 1, 75000),
(1, 3, 'pending', 1, 25000),
(2, 2, 'success', 2, 100000),
(2, 4, 'pending', 1, 50000),
(2, 5, 'success', 2, 150000),
(3, 3, 'success', 1, 50000),
(3, 6, 'pending', 2, 100000),
(3, 7, 'success', 1, 75000),
(4, 1, 'success', 1, 50000),
(4, 2, 'success', 1, 75000),
(4, 3, 'pending', 1, 25000),
(5, 2, 'success', 2, 100000),
(5, 4, 'pending', 1, 50000),
(5, 5, 'success', 2, 150000),
(6, 3, 'success', 1, 50000),
(6, 6, 'pending', 2, 100000),
(6, 7, 'success', 1, 75000),
(7, 1, 'success', 1, 50000),
(7, 2, 'success', 1, 75000),
(7, 3, 'pending', 1, 25000),
(8, 2, 'success', 2, 100000),
(8, 4, 'pending', 1, 50000),
(8, 5, 'success', 2, 150000),
(9, 3, 'success', 1, 50000),
(9, 6, 'pending', 2, 100000),
(9, 7, 'success', 1, 75000),
(10, 1, 'success', 1, 50000),
(10, 2, 'success', 1, 75000),
(10, 3, 'pending', 1, 25000),
(11, 2, 'success', 2, 100000),
(11, 4, 'pending', 1, 50000),
(11, 5, 'success', 2, 150000),
(12, 3, 'success', 1, 50000),
(12, 6, 'pending', 2, 100000),
(12, 7, 'success', 1, 75000),
(13, 1, 'success', 1, 50000),
(13, 2, 'success', 1, 75000),
(13, 3, 'pending', 1, 25000),
(14, 2, 'success', 2, 100000),
(14, 4, 'pending', 1, 50000),
(14, 5, 'success', 2, 150000),
(15, 3, 'success', 1, 50000),
(15, 6, 'pending', 2, 100000),
(15, 7, 'success', 1, 75000);

-- 2. SELECT
-- a 
SELECT nama FROM user WHERE gender = 'male';
-- b 
SELECT * FROM product WHERE id = 3;
-- c
SELECT * FROM user WHERE nama LIKE '%a%' 
AND created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW();
-- d
SELECT COUNT(*) as jumlah_user_perempuan FROM user WHERE gender = 'female';
-- e
SELECT * FROM user ORDER BY nama ASC;
-- f
SELECT * FROM product LIMIT 5;

-- 3. UPDATE 
-- a
UPDATE product SET nama = 'product dummy' WHERE id = 1;
-- b
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- 4. DELETE
-- a
DELETE FROM product WHERE id = 1;
-- b
DELETE FROM product WHERE product_type_id = 1;

-- Join, Union, Sub query, Function
-- 1.
SELECT * FROM transaction WHERE user_id = 1
UNION 
SELECT * FROM transaction WHERE user_id = 2;
-- 2 
SELECT SUM(total_price) as jumlah_harga FROM transaction WHERE user_id = 1;
-- 3
SELECT COUNT(*) as total_transaksi FROM transaction t INNER JOIN transaction_details td 
ON t.id = td.transaction_id
INNER JOIN product p ON td.product_id = p.id
WHERE p.product_type_id = 2;
-- 4
SELECT p.*, pt.nama as product_type_nama FROM product p
INNER JOIN product_type pt ON p.product_type_id = pt.id;
-- 5 
SELECT t.*, p.nama as product_nama, u.nama as user_nama FROM transaction t
INNER JOIN transaction_details td ON t.id = td.transaction_id
INNER JOIN product p ON td.product_id = p.id
INNER JOIN user u ON t.user_id = u.id;
-- 6
DELIMITER //
CREATE TRIGGER delete_transaction_details AFTER DELETE ON transaction
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END; //
DELIMITER ;
-- 7
DELIMITER //
CREATE TRIGGER update_transaction_qty AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transaction t SET t.total_qty = (SELECT SUM(td.qty) FROM transaction_details td 
    WHERE td.transaction_id = t.id) 
   WHERE t.id = OLD.transaction_id;
END; //
DELIMITER ;
-- 8
SELECT * FROM product WHERE id NOT IN ( SELECT product_id FROM transaction_details);

