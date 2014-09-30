drop table IF EXISTS `Item`;
CREATE TABLE Item (
    Type VARCHAR(50), 
    Quality VARCHAR(50), 
    Description VARCHAR(100) NOT NULL DEFAULT '', 
    HistFigureId bigint,
    ItemOwnerId Bigint,
    ItemHolderId Bigint,
    Age int,
    Wear int,
    Id Bigint,
    DeadDwarf tinyint,
    InChest tinyint,
    OnGround tinyint,
    InInventory tinyint,
    Owned tinyint,
    PRIMARY KEY (Id)
);

ALTER TABLE Dwarf ADD HistFigureId int;
