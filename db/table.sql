CREATE SCHEMA info;

CREATE SEQUENCE info.user_user_seq_seq;

CREATE SEQUENCE info.business_card_bc_seq_seq;

CREATE SEQUENCE info.owner_bc_owner_seq_seq;


CREATE TABLE auth_token(
		id VARCHAR(256) NOT NULL PRIMARY KEY,
		grant_type VARCHAR(512),
		access_token VARCHAR(256),
		refresh_token VARCHAR(256)
);

/**********************************/
/* Table Name: info.user */
/**********************************/
CREATE TABLE info.user(
		user_seq INT NOT NULL default nextval('info.user_user_seq_seq'::regclass),
		user_id VARCHAR(256) NOT NULL PRIMARY KEY,
		password VARCHAR(512) NOT NULL,
		name VARCHAR(128),
		company VARCHAR(128),
		position VARCHAR(128),
		duty VARCHAR(128),
		phone VARCHAR(16),
		email VARCHAR(256),
  CONSTRAINT IDX_user_1 UNIQUE (user_seq)
);

/**********************************/
/* Table Name: info.business_card */
/**********************************/
CREATE TABLE info.business_card(
		bc_seq INT NOT NULL PRIMARY KEY default nextval('info.business_card_bc_seq_seq'::regclass),
		user_id VARCHAR(256) NOT NULL REFERENCES info.user (user_id),
		name VARCHAR(128),
		company VARCHAR(128),
		depart VARCHAR(256),
		team VARCHAR(128),
		position VARCHAR(128),
		duty VARCHAR(128),
		phone VARCHAR(16),
		tel VARCHAR(16),
		fax VARCHAR(32),
		address VARCHAR(512),
		email VARCHAR(256),
		priority INT DEFAULT 9 NOT NULL,
  CONSTRAINT IDX_business_card_1 UNIQUE (bc_seq)
);

CREATE TABLE info.owner_bc(
		owner_seq INT NOT NULL default nextval('info.owner_bc_owner_seq_seq'::regclass),
		owner_bc_seq INT,
		owner_user_id VARCHAR(256),
		name VARCHAR(128),
		company VARCHAR(128),
		depart VARCHAR(256),
		team VARCHAR(128),
		position VARCHAR(128),
		duty VARCHAR(128),
		phone VARCHAR(16),
		tel VARCHAR(16),
		fax VARCHAR(32),
		address VARCHAR(512),
		email VARCHAR(256),
  FOREIGN KEY (owner_user_id) REFERENCES info.user (user_id),
  FOREIGN KEY (owner_bc_seq) REFERENCES info.business_card (bc_seq),
  CONSTRAINT IDX_owner_bc_1 UNIQUE (owner_seq)
);

CREATE INDEX IDX_user_2 ON info.user (user_id);

CREATE INDEX IDX_business_card_2 ON info.business_card (user_id);

CREATE INDEX IDX_owner_bc_2 ON info.owner_bc (owner_user_id);
CREATE INDEX IDX_owner_bc_3 ON info.owner_bc (owner_bc_seq);
