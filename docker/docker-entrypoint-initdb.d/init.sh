set -e
psql -U root -d sample << EOSQL
CREATE TABLE Users(
  account_id        SERIAL PRIMARY KEY,
  account_name      VARCHAR(20),
  email             VARCHAR(100),
  password    CHAR(64)
);

CREATE TABLE UserStatus(
  status            VARCHAR(20) PRIMARY KEY
);

INSERT INTO Users (account_name, email,password) VALUES ( 'Jack Otani', 'syazai.co.jp','baka123');
INSERT INTO Users (account_name, email,password) VALUES ( 'Jack Otani', 'syazai.co.jp','baka123');
INSERT INTO Users (account_name, email,password) VALUES ( 'Jack Otani', 'syazai.co.jp','baka123');
EOSQLdoc