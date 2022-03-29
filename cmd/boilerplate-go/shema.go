package main

// schema for user table
const userSchema = `
		create table if not exists users (
			id 		   Varchar(36) not null,
			email 	   Varchar(100) not null unique,
			username   Varchar(225),
			password   Varchar(225) not null,
			tokenhash  Varchar(15) not null,
			isverified Boolean default false,
			createdat  Timestamp not null,
			updatedat  Timestamp not null,
			Primary Key (id)
		);
`

const verificationSchema = `
		create table if not exists verifications (
			email 		Varchar(100) not null,
			code  		Varchar(10) not null,
			expiresat 	Timestamp not null,
			type        Varchar(10) not null,
			Primary Key (email),
			Constraint fk_user_email Foreign Key(email) References users(email)
				On Delete Cascade On Update Cascade
		)
`
