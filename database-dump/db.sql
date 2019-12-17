CREATE DATABASE IF NOT EXISTS source;

USE source;

CREATE TABLE IF NOT EXISTS sources(
    source_id varchar(100) not null primary key,
    card_number  bigint(16) unsigned not null,
    cvv int(4) unsigned not null 
);

INSERT INTO sources (source_id, card_number, cvv) values ("token-1", 1111222233334444, 123);
INSERT INTO sources (source_id, card_number, cvv) values ("token-2", 5555666677778888, 456);