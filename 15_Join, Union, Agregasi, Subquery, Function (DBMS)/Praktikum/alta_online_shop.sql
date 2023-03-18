-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 18, 2023 at 05:13 PM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.0.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `alamat`
--

CREATE TABLE `alamat` (
  `alamat_id` int(11) NOT NULL,
  `alamat` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `alamat`
--

INSERT INTO `alamat` (`alamat_id`, `alamat`) VALUES
(1, 'Jalan Raya Tiga No. 123'),
(2, 'Jalan Raya Dua No. 456'),
(3, 'Jalan Raya Satu No. 789'),
(4, 'Jalan Raya Enam No. 246'),
(5, 'Jalan Raya Lima No. 135');

-- --------------------------------------------------------

--
-- Table structure for table `operator`
--

CREATE TABLE `operator` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `operator`
--

INSERT INTO `operator` (`id`, `nama`, `created_at`, `updated_at`) VALUES
(1, 'Puff', '2023-03-18 03:53:14', '2023-03-18 03:53:14'),
(2, 'Yeti', '2023-03-18 03:53:14', '2023-03-18 03:53:14'),
(3, 'Ray Chargels', '2023-03-18 03:53:14', '2023-03-18 03:53:14'),
(4, 'Honey Bear', '2023-03-18 03:53:14', '2023-03-18 03:53:14'),
(5, 'Hadha Kautsar', '2023-03-18 04:02:14', '2023-03-18 04:02:14');

-- --------------------------------------------------------

--
-- Table structure for table `payment_method`
--

CREATE TABLE `payment_method` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `payment_method_description_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment_method`
--

INSERT INTO `payment_method` (`id`, `nama`, `status`, `created_at`, `updated_at`, `payment_method_description_id`) VALUES
(1, 'Cash', 1, '2023-03-18 09:39:08', '2023-03-18 09:39:08', 1),
(2, 'Bank Transfer', 1, '2023-03-18 09:39:08', '2023-03-18 09:39:08', 2),
(3, 'Credit Card', 1, '2023-03-18 09:39:08', '2023-03-18 09:39:08', 3);

-- --------------------------------------------------------

--
-- Table structure for table `payment_method_description`
--

CREATE TABLE `payment_method_description` (
  `payment_method_description_id` int(11) NOT NULL,
  `payment_method_description` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `product_type_id`, `operator_id`, `code`, `nama`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 1, 'HB1', 'Jus Buah', 1, '2023-03-18 08:09:48', '2023-03-18 08:09:48'),
(4, 2, 1, 'YT1', 'Soda', 1, '2023-03-18 08:09:48', '2023-03-18 08:09:48'),
(5, 2, 1, 'RC2', 'Air Tawar', 1, '2023-03-18 08:09:48', '2023-03-18 08:09:48'),
(6, 3, 4, 'HB2', 'Daging', 1, '2023-03-18 08:09:51', '2023-03-18 08:09:51'),
(7, 3, 4, 'RC3', 'Ikan', 1, '2023-03-18 08:09:51', '2023-03-18 08:09:51'),
(8, 3, 4, 'PF2', 'Ayam', 1, '2023-03-18 08:09:51', '2023-03-18 08:09:51');

-- --------------------------------------------------------

--
-- Table structure for table `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Berbagai macam pensil', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(2, 'Berbagai macam penghapus', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(3, 'Berbagai macam jus buah', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(4, 'Berbagai macam soda', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(5, 'Berbagai macam air tawar', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(6, 'Berbagai macam daging', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(7, 'Berbagai macam ikan', '2023-03-18 08:09:55', '2023-03-18 08:09:55'),
(8, 'Berbagai macam ayam', '2023-03-18 08:09:55', '2023-03-18 08:09:55');

-- --------------------------------------------------------

--
-- Table structure for table `product_type`
--

CREATE TABLE `product_type` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_type`
--

INSERT INTO `product_type` (`id`, `nama`, `created_at`, `updated_at`) VALUES
(1, 'Alat Tulis', '2023-03-18 08:09:29', '2023-03-18 08:09:29'),
(2, 'Minuman', '2023-03-18 08:09:29', '2023-03-18 08:09:29'),
(3, 'Makanan', '2023-03-18 08:09:29', '2023-03-18 08:09:29');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` decimal(25,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(2, 1, 2, 'success', 3, '250000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(3, 1, 1, 'canceled', 1, '75000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(4, 2, 2, 'success', 5, '9000000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(5, 2, 1, 'pending', 1, '100000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(6, 2, 2, 'pending', 3, '200000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(7, 3, 1, 'canceled', 2, '200000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(8, 3, 2, 'success', 4, '2500000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(9, 3, 1, 'pending', 1, '50000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(10, 4, 1, 'canceled', 2, '400000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(11, 4, 2, 'success', 4, '4500000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(12, 4, 1, 'pending', 1, '5000000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(13, 5, 1, 'canceled', 2, '600000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(14, 5, 2, 'success', 4, '5500000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16'),
(15, 5, 1, 'pending', 1, '900000.00', '2023-03-18 07:32:16', '2023-03-18 07:32:16');

--
-- Triggers `transaction`
--
DELIMITER $$
CREATE TRIGGER `delete_transaction_details` AFTER DELETE ON `transaction` FOR EACH ROW BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` decimal(25,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(4, 2, 2, 'success', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(5, 2, 4, 'pending', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(6, 2, 5, 'success', 2, '150000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(7, 3, 3, 'success', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(8, 3, 6, 'pending', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(9, 3, 7, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(10, 4, 1, 'success', 3, '50000.00', '2023-03-18 08:40:41', '2023-03-18 09:54:29'),
(11, 4, 2, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(12, 4, 3, 'pending', 1, '25000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(13, 5, 2, 'success', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(14, 5, 4, 'pending', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(15, 5, 5, 'success', 2, '150000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(16, 6, 3, 'success', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(17, 6, 6, 'pending', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(18, 6, 7, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(19, 7, 1, 'success', 3, '50000.00', '2023-03-18 08:40:41', '2023-03-18 09:54:29'),
(20, 7, 2, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(21, 7, 3, 'pending', 1, '25000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(22, 8, 2, 'success', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(23, 8, 4, 'pending', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(24, 8, 5, 'success', 2, '150000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(25, 9, 3, 'success', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(26, 9, 6, 'pending', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(27, 9, 7, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(28, 10, 1, 'success', 3, '50000.00', '2023-03-18 08:40:41', '2023-03-18 09:54:29'),
(29, 10, 2, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(30, 10, 3, 'pending', 1, '25000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(31, 11, 2, 'success', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(32, 11, 4, 'pending', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(33, 11, 5, 'success', 2, '150000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(34, 12, 3, 'success', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(35, 12, 6, 'pending', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(36, 12, 7, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(37, 13, 1, 'success', 3, '50000.00', '2023-03-18 08:40:41', '2023-03-18 09:54:29'),
(38, 13, 2, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(39, 13, 3, 'pending', 1, '25000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(40, 14, 2, 'success', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(41, 14, 4, 'pending', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(42, 14, 5, 'success', 2, '150000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(43, 15, 3, 'success', 1, '50000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(44, 15, 6, 'pending', 2, '100000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41'),
(45, 15, 7, 'success', 1, '75000.00', '2023-03-18 08:40:41', '2023-03-18 08:40:41');

--
-- Triggers `transaction_details`
--
DELIMITER $$
CREATE TRIGGER `update_transaction_qty` AFTER DELETE ON `transaction_details` FOR EACH ROW BEGIN
    UPDATE transaction t SET t.total_qty = (SELECT SUM(td.qty) FROM transaction_details td 
    WHERE td.transaction_id = t.id) 
   WHERE t.id = OLD.transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat_id` int(11) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `status` smallint(6) NOT NULL,
  `gender` enum('male','female') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `nama`, `alamat_id`, `tanggal_lahir`, `status`, `gender`, `created_at`, `updated_at`) VALUES
(27, 'Maulana Ibrahim', 1, '1990-05-15', 1, 'male', '2023-03-18 07:15:03', '2023-03-18 07:15:03'),
(28, 'Susana Listiyani', 2, '1995-07-20', 1, 'female', '2023-03-18 07:15:03', '2023-03-18 07:15:03'),
(29, 'Eko Prasetyanto', 3, '1985-03-12', 1, 'male', '2023-03-18 07:15:03', '2023-03-18 07:15:03'),
(30, 'Liliana Anggita ', 4, '1992-12-28', 1, 'female', '2023-03-18 07:15:03', '2023-03-18 07:15:03'),
(31, 'Bambang Lukman', 5, '1988-08-08', 1, 'male', '2023-03-18 07:15:03', '2023-03-18 07:15:03');

-- --------------------------------------------------------

--
-- Table structure for table `user_payment_method_detail`
--

CREATE TABLE `user_payment_method_detail` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `alamat`
--
ALTER TABLE `alamat`
  ADD PRIMARY KEY (`alamat_id`);

--
-- Indexes for table `operator`
--
ALTER TABLE `operator`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FK_payment_method_description` (`payment_method_description_id`);

--
-- Indexes for table `payment_method_description`
--
ALTER TABLE `payment_method_description`
  ADD PRIMARY KEY (`payment_method_description_id`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_type`
--
ALTER TABLE `product_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_alamat` (`alamat_id`);

--
-- Indexes for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `operator`
--
ALTER TABLE `operator`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `payment_method`
--
ALTER TABLE `payment_method`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `product_type`
--
ALTER TABLE `product_type`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- AUTO_INCREMENT for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `fk_user_alamat` FOREIGN KEY (`alamat_id`) REFERENCES `alamat` (`alamat_id`);

--
-- Constraints for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  ADD CONSTRAINT `user_payment_method_detail_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  ADD CONSTRAINT `user_payment_method_detail_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
