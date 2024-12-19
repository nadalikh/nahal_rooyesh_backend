-- MySQL dump 10.13  Distrib 8.0.40, for Linux (x86_64)
--
-- Host: localhost    Database: green_house
-- ------------------------------------------------------
-- Server version	8.0.40-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `config`
--

DROP TABLE IF EXISTS `config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `config` (
  `id` int NOT NULL AUTO_INCREMENT,
  `element_slug` varchar(200) DEFAULT NULL,
  `congif_json` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `config_element_slug_uindex` (`element_slug`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `config`
--

LOCK TABLES `config` WRITE;
/*!40000 ALTER TABLE `config` DISABLE KEYS */;
INSERT INTO `config` VALUES (1,'pipe','{\"galvanize\":\"fabric\", \"props\" : {\"diagonal_id\":1, \"thickness_id\":3, \"length\":100}}');
/*!40000 ALTER TABLE `config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fabric`
--

DROP TABLE IF EXISTS `fabric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fabric` (
  `id` int NOT NULL AUTO_INCREMENT,
  `element_slug` varchar(255) DEFAULT NULL,
  `price` int DEFAULT NULL,
  `thickness_id` int DEFAULT NULL,
  `diagonal_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fabric_properties_id_fk` (`diagonal_id`),
  KEY `fabric_properties_id_fk_2` (`thickness_id`),
  CONSTRAINT `fabric_properties_id_fk` FOREIGN KEY (`diagonal_id`) REFERENCES `properties` (`id`),
  CONSTRAINT `fabric_properties_id_fk_2` FOREIGN KEY (`thickness_id`) REFERENCES `properties` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fabric`
--

LOCK TABLES `fabric` WRITE;
/*!40000 ALTER TABLE `fabric` DISABLE KEYS */;
INSERT INTO `fabric` VALUES (6,'pipe',1000,4,1),(7,'pipe',1000,3,2),(8,'pipe',56100,3,1);
/*!40000 ALTER TABLE `fabric` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `iron_properties`
--

DROP TABLE IF EXISTS `iron_properties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `iron_properties` (
  `slug` varchar(255) NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `value` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `iron_properties`
--

LOCK TABLES `iron_properties` WRITE;
/*!40000 ALTER TABLE `iron_properties` DISABLE KEYS */;
INSERT INTO `iron_properties` VALUES ('warm_price',1,65000);
/*!40000 ALTER TABLE `iron_properties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `properties`
--

DROP TABLE IF EXISTS `properties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `properties` (
  `id` int NOT NULL AUTO_INCREMENT,
  `slug` varchar(200) DEFAULT NULL,
  `value` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `properties`
--

LOCK TABLES `properties` WRITE;
/*!40000 ALTER TABLE `properties` DISABLE KEYS */;
INSERT INTO `properties` VALUES (1,'diagonal',25),(2,'diagonal',32),(3,'thickness',1.5),(4,'thickness',2);
/*!40000 ALTER TABLE `properties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warm`
--

DROP TABLE IF EXISTS `warm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `warm` (
  `id` int NOT NULL AUTO_INCREMENT,
  `diagonal_id` int DEFAULT NULL,
  `thickness_id` int DEFAULT NULL,
  `price` int DEFAULT NULL,
  `element_slug` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `khorshidi_fabric_khorshidi_properties_null_fk_digonal` (`diagonal_id`),
  KEY `khorshidi_fabric_khorshidi_properties_null_fk_thickness` (`thickness_id`),
  CONSTRAINT `khorshidi_fabric_khorshidi_properties_null_fk_digonal` FOREIGN KEY (`diagonal_id`) REFERENCES `properties` (`id`),
  CONSTRAINT `khorshidi_fabric_khorshidi_properties_null_fk_thickness` FOREIGN KEY (`thickness_id`) REFERENCES `properties` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warm`
--

LOCK TABLES `warm` WRITE;
/*!40000 ALTER TABLE `warm` DISABLE KEYS */;
INSERT INTO `warm` VALUES (14,2,4,60000,'khorshidi'),(16,1,4,2000,'khorshidi'),(17,1,3,62000,'khorshidi');
/*!40000 ALTER TABLE `warm` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-12-19 10:34:18
