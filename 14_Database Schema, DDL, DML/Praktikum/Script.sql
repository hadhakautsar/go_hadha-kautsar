-- Membuat database alta_online_shop
create database alta_online_shop

-- Menggunakan database alta_online_shop
USE alta_online_shop;

-- Membuat table user
CREATE TABLE user (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nama VARCHAR(255) NOT NULL,
  alamat_id INT,
  tanggal_lahir DATE NOT NULL,
  status SMALLINT NOT NULL,
  gender ENUM('male', 'female') NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- Melihat semua data pada table user 
select * from `user` u ; 

-- Membuat table product 
CREATE TABLE product (
  id INT PRIMARY KEY AUTO_INCREMENT,
  product_type_id INT(11) NOT NULL,
  operator_id INT(11) NOT NULL,
  code VARCHAR(50) NOT NULL,
  nama VARCHAR(100) NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table product_type
CREATE TABLE product_type (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nama VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table operator
CREATE TABLE operator (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nama VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table product_descriptions
CREATE TABLE product_descriptions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table payment_method
CREATE TABLE payment_method (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nama VARCHAR(255),
  status SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table transaction
CREATE TABLE transaction (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT(11) NOT NULL,  
  payment_method_id INT(11) NOT NULL,
  status VARCHAR(10) NOT NULL,
  total_qty INT(11) NOT NULL, 
  total_price NUMERIC(25,2) NOT NULL, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table transcation_details
CREATE TABLE transaction_details (
  id INT PRIMARY KEY AUTO_INCREMENT,
  transaction_id INT(11) NOT NULL, 
  product_id INT(11) NOT NULL,  
  status VARCHAR(10) NOT NULL,
  qty INT(11) NOT NULL, 
  price NUMERIC(25,2) NOT NULL, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Membuat table kurir
CREATE TABLE kurir (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nama VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Menambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurir 
ADD ongkos_dasar NUMERIC(25,2) not NULL;

-- Rename tabel kurir menjadi shipping.
ALTER TABLE kurir
RENAME TO shipping;

-- Menghapus tabel shipping
 DROP TABLE shipping; 
 
-- 1-to-1: Membuat tabel payment_method_description
CREATE TABLE payment_method_description (
  payment_method_description_id INT PRIMARY KEY,
  payment_method_description VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 1-to-1: menambahkan kolom payment_method_description_id pada payment_method
ALTER TABLE payment_method
ADD payment_method_description_id INT NOT NULL,
ADD CONSTRAINT FK_payment_method_description
  FOREIGN KEY (payment_method_description_id)
  REFERENCES payment_method_description(payment_method_description_id);
  
 -- 1-to-many: membuat tabel alamat
 CREATE TABLE alamat (
  alamat_id INT PRIMARY KEY,
  alamat VARCHAR(255) NOT NULL
);
 
-- 1-to-many: hubungan antara tabel user dan alamat
ALTER TABLE user
ADD CONSTRAINT fk_user_alamat
FOREIGN KEY (alamat_id) 
REFERENCES alamat(alamat_id);

-- many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_detail (
id INT PRIMARY KEY AUTO_INCREMENT,
user_id INT NOT NULL,
payment_method_id INT NOT NULL,
FOREIGN KEY (user_id) REFERENCES user(id),
FOREIGN KEY (payment_method_id) REFERENCES payment_method(id),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
