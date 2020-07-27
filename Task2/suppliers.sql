/*
SQLyog Community v13.1.6 (64 bit)
MySQL - 10.4.8-MariaDB 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `suppliers` (
	`SupplierID` int (11),
	`CompanyName` varchar (120),
	`ContactName` varchar (90),
	`ContactTitle` varchar (90),
	`Address` varchar (180),
	`City` varchar (45),
	`Region` varchar (45),
	`PostalCode` varchar (30),
	`Country` varchar (45),
	`Phone` varchar (72),
	`Fax` varchar (72),
	`HomePage` text 
); 
