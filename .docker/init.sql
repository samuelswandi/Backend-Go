USE `mahasiswa`;

CREATE TABLE IF NOT EXISTS `mahasiswa` (
  `id` int(11) UNIQUE NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `usia` int(11) NOT NULL,
  `gender` int(11) NOT NULL,
  `tanggal_registrasi` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `jurusan` (
  `id` int(11) UNIQUE NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `hobi` (
  `id` int(11) UNIQUE NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `mahasiswa_hobi` (
  `id_mahasiswa` int(11) NOT NULL,
  `id_hobi` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `mahasiswa_jurusan` (
  `id_mahasiswa` int(11) NOT NULL,
  `id_jurusan` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
