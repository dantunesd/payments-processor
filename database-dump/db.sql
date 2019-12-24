CREATE DATABASE IF NOT EXISTS source;

USE source;

CREATE TABLE IF NOT EXISTS sources(
    source_id varchar(100) not null primary key,
    card_number  varchar(16) not null,
    cvv int(4) unsigned not null 
);

INSERT INTO sources (source_id, card_number, cvv) values ("authorized", "0000000000000001", 123);
INSERT INTO sources (source_id, card_number, cvv) values ("not-authorized", "0000000000000002", 123);
INSERT INTO sources (source_id, card_number, cvv) values ("expired", "0000000000000003", 123);
INSERT INTO sources (source_id, card_number, cvv) values ("invalid", "invalid-card", 123);
INSERT INTO sources (source_id, card_number, cvv) values ("rauthorized", "5448280000000007", 123);