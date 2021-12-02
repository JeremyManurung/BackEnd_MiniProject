-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 28 Nov 2021 pada 12.04
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.4.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `backend`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `bantuans`
--

CREATE TABLE `bantuans` (
  `id` int(20) UNSIGNED NOT NULL,
  `jumlah_bar` int(20) NOT NULL,
  `deskripsi` text NOT NULL,
  `deskripsi_singkat` varchar(255) NOT NULL,
  `list_kondisi` text NOT NULL,
  `prm` varchar(255) NOT NULL,
  `jumlah_pendonasi` int(20) NOT NULL,
  `user_id` int(20) NOT NULL,
  `tittle_bantuan` varchar(255) NOT NULL,
  `jumlah_target` int(11) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `bantuans`
--

INSERT INTO `bantuans` (`id`, `jumlah_bar`, `deskripsi`, `deskripsi_singkat`, `list_kondisi`, `prm`, `jumlah_pendonasi`, `user_id`, `tittle_bantuan`, `jumlah_target`, `created`, `updated`) VALUES
(1, 0, 'Bantuan kepada pak rahman yang luka luka akibat tidak punya rumha', 'bantuan untuk pak rahman', 'luka-luka,sakit,tidak punya tempat tinggal', 'bantuan-satu', 0, 1, 'Bantuan 1', 10000000, '2021-11-27 12:40:26', '2021-11-27 12:40:26'),
(2, 0, 'Bantuan kepada pak joko yang luka luka akibat tidak punya rumha', 'bantuan untuk pak rahman', 'luka-luka,sakit,tidak punya tempat tinggal', 'bantuan-dua', 0, 1, 'Bantuan 2', 10000000, '2021-11-27 12:40:26', '2021-11-27 12:40:26'),
(3, 0, 'Bantuan kepada pak uyax yang luka luka akibat tidak punya rumha', 'bantuan untuk pak uyax', 'luka-luka,sakit,tidak punya tempat tinggal', 'bantuan-tiga', 0, 2, 'Bantuan 3', 10000000, '2021-11-27 12:48:22', '2021-11-27 12:48:22'),
(4, 0, 'dana yang akan dirikimkan kepada tim dan uang akan di gunakan untuk membantu ibu tobox', 'dana untuk tobox yang mengalami luka-luka', 'kelaparan,bingung,luka-luka', 'dana-untuk-tobox-1', 0, 1, 'Dana untuk tobox', 100000000, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `bantuan_imgs`
--

CREATE TABLE `bantuan_imgs` (
  `id` int(20) UNSIGNED NOT NULL,
  `bantuan_id` int(20) NOT NULL,
  `tittle_img` varchar(255) NOT NULL,
  `img_utama` tinyint(7) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `bantuan_imgs`
--

INSERT INTO `bantuan_imgs` (`id`, `bantuan_id`, `tittle_img`, `img_utama`, `created`, `updated`) VALUES
(1, 1, 'satu.jpg', 0, '2021-11-27 13:14:35', '2021-11-27 13:14:35'),
(2, 1, 'sapek.jpg', 0, '2021-11-27 13:14:35', '2021-11-27 13:14:35'),
(3, 1, 'tiga.png', 1, '2021-11-27 13:23:14', '2021-11-27 13:23:14');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `id` int(20) NOT NULL,
  `jumlah_uang` int(20) NOT NULL,
  `status_transaksi` varchar(255) NOT NULL,
  `bantuan_id` int(20) NOT NULL,
  `user_id` int(20) NOT NULL,
  `kode_transaksi` varchar(255) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(20) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `user_img` varchar(255) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `role` varchar(255) NOT NULL,
  `pekerjaan` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `nama`, `user_img`, `created`, `updated`, `role`, `pekerjaan`) VALUES
(1, 'jeremy@gmail.com', 'miror123', 'Jeremy', 'images/cekimg.jpg', '2021-11-23 23:01:58', '2021-11-23 23:01:58', 'user', 'programmer'),
(2, 'sweety@gmail.com', 'mikro11334', 'Sweety', 'image/2-Universitas-Katolik-Santo-Thomas-Sumatera-Utara.png', '2021-11-23 23:02:35', '2021-11-23 23:02:35', 'user', 'dokter'),
(6, 'tes@gmail.com', '$2a$04$ChzE/Wmxk8gjP7sdPo2WOeEq5WT8td5xcOzVmYlMJQUROD9p/HLnG', 'Tes', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Tukang Cambu'),
(9, 'makan@gmail.com', '$2a$04$7b4EWaoP30jYFWEf800Rf.CiUbrtXh8gQuTMHlefV7xiHJsYhaT/G', 'anak tuhan', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Tukang Makan'),
(10, 'pusing@gmail.com', '$2a$04$FKru0kSJwy/Wv60ASuIaHOgLn1WAqYbMZOBrAk7EbnnY1KSyfRysG', 'dewa mabuk', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Tukang Kayu'),
(11, 'bersin@gmail.com', '$2a$04$d5Gmbt5LsCoDzsROok5lruvit8HpJVnYnc5QHldK/yHmlK8jyHPWe', 'dewa silat', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Tukang Hantam'),
(19, 'angin@gmail.com', '$2a$04$yZHCea/VZEy9cqU3RNuKoOXKNBO6uH8WZaIxy10EpHhhd.jlZdg8a', 'dewa angin', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Hembus Angin'),
(20, 'kaneki@gmail.com', '$2a$04$bWTQfFcoRAndEdd/MD3dO.kdu2VIrdCSnN/BIFUiYSgT0PB6FOnia', 'kaneki', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'user', 'Tukang Jaggal');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `bantuans`
--
ALTER TABLE `bantuans`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `bantuan_imgs`
--
ALTER TABLE `bantuan_imgs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `bantuans`
--
ALTER TABLE `bantuans`
  MODIFY `id` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `bantuan_imgs`
--
ALTER TABLE `bantuan_imgs`
  MODIFY `id` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
