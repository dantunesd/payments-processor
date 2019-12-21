CREATE DATABASE IF NOT EXISTS source;

USE source;

CREATE TABLE IF NOT EXISTS sources(
    source_id varchar(100) not null primary key,
    card_number  bigint(16) unsigned not null,
    cvv int(4) unsigned not null 
);

INSERT INTO sources (source_id, card_number, cvv) values ("auth", 0000000000000001, 123);
INSERT INTO sources (source_id, card_number, cvv) values ("nonauth", 0000000000000002, 123);
INSERT INTO sources (source_id, card_number, cvv) values ("exp", 0000000000000003, 123);