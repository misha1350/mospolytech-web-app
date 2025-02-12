CREATE TABLE `countries` (
  ID int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Name varchar(50) NOT NULL
);

CREATE TABLE `offices` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `CountryID` int(11) NOT NULL,
  `Title` varchar(50) COLLATE utf8_bin NOT NULL,
  `Phone` varchar(50) COLLATE utf8_bin NOT NULL,
  `Contact` varchar(250) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_Office_Country` (`CountryID`),
  CONSTRAINT `FK_Office_Country` FOREIGN KEY (`CountryID`) REFERENCES `countries` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE `roles` (
  `ID` int(11) NOT NULL,
  `Title` varchar(50) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
);

CREATE TABLE `users` (
  `ID` int(11) NOT NULL,
  `RoleID` int(11) NOT NULL,
  `Email` varchar(254) COLLATE utf8_bin NOT NULL,
  `Password` varchar(255) COLLATE utf8_bin NOT NULL,
  `FirstName` varchar(50) COLLATE utf8_bin NOT NULL,
  `LastName` varchar(50) COLLATE utf8_bin NOT NULL,
  `OfficeID` int(11) NOT NULL,
  `Birthdate` date NOT NULL,
  `Active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_Users_Offices` (`OfficeID`),
  KEY `FK_Users_Roles` (`RoleID`),
  CONSTRAINT `FK_Users_Offices` FOREIGN KEY (`OfficeID`) REFERENCES `offices` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_Users_Roles` FOREIGN KEY (`RoleID`) REFERENCES `roles` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE `tracking` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `Date` date NOT NULL,
  `TimeIn` datetime NOT NULL,
  `TimeOut` datetime NOT NULL,
  `Hours` decimal(10,2) NOT NULL,
  `Notes` varchar(250) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_Tracking_Users` (`UserID`),
  CONSTRAINT `FK_Tracking_Users` FOREIGN KEY (`UserID`) REFERENCES `users` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE `authentication_tokens` (
  `TokenID` BIGINT NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `AuthToken` varchar(255) COLLATE utf8_bin NOT NULL,
  `GeneratedAt` datetime NOT NULL,
  `ExpiresAt` datetime NOT NULL,
  PRIMARY KEY (`TokenID`),
  CONSTRAINT `FK_Authentication_Tokens_Users` FOREIGN KEY (`UserID`) REFERENCES `users` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
);