package data

// schema for trening users params table
const treningUserParamsSchema = `
		create table if not exists trening_users_params (
			uid 		   Varchar(36) not null unique,
			user_id 	   Varchar(36) not null unique,
			weight 	       integer ,
			height         integer,
			age            integer,
			gender         integer not null,
			activity       integer,
			diet           integer,
			desired_weight integer,
			eat            integer,
			createdat      Timestamp not null,
			updatedat      Timestamp not null,
			Primary Key (uid)
		);
`
