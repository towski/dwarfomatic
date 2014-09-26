drop table IF EXISTS `User`;
CREATE TABLE User (
    Name VARCHAR(20), 
    Id BIGINT AUTO_INCREMENT,
    PRIMARY KEY (id)
);

drop table IF EXISTS `Dwarf`;
CREATE TABLE Dwarf (
    Name VARCHAR(50), 
    Job VARCHAR(60), 
    Mood VARCHAR(60), 
    Id BIGINT,
    PRIMARY KEY (Id)
);
