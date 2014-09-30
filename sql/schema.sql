drop table IF EXISTS `User`;
CREATE TABLE User (
    Name VARCHAR(20), 
    Password VARCHAR(100), 
    Ip VARCHAR(30), 
    Id BIGINT AUTO_INCREMENT,
    PRIMARY KEY (id)
);

CREATE unique index user_name on User (Name);
CREATE unique index user_ip on User (Ip);

drop table IF EXISTS `Dwarf`;
CREATE TABLE Dwarf (
    Name VARCHAR(50), 
    Job VARCHAR(60), 
    Mood VARCHAR(60), 
    Thoughts TEXT,
    UserId BIGINT,
    Id BIGINT,
    PRIMARY KEY (Id)
);

CREATE unique index user_id on Dwarf (UserId);
