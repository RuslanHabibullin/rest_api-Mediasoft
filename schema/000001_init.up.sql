CREATE TABLE furnitures(
    id serial not null UNIQUE,
    name VARCHAR(255) ,
    manufacturer VARCHAR(255),
    height SMALLINT,
    lenght SMALLINT,
    width SMALLINT
)