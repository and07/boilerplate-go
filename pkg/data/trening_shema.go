package data

// schema for trening users params table
const treningUserParamsSchema = `
		create table if not exists trening_users_params (
			uid 		   Varchar(36) not null unique,
			user_id 	   Varchar(36) not null unique,
			username       Varchar(225),
			weight 	       integer ,
			height         integer,
			age            integer,
			gender         integer not null,
			activity       integer,
			diet           integer,
			desired_weight integer,
			eat            integer,
			image          Varchar(225),
			createdat      Timestamp not null,
			updatedat      Timestamp not null,
			Primary Key (uid)
		);
`

const treningExercise = `
	create table if not exists trening_exercise (
		uid 		  			Varchar(36) not null unique,
		user_id 	   			Varchar(36) not null,
		name 	       			Varchar(100) not null ,
		duration			    Timestamp,
		relax              	    Timestamp not null,
		count                   integer,
		number_of_sets          integer,
		number_of_repetitions   integer,
		type 				    integer,
		createdat      			Timestamp not null,
		updatedat     			Timestamp not null,
		Primary Key (uid)
	);
`
const trening = `
	create table if not exists trening (
		uid 		  			Varchar(36) not null unique,
		user_id 	   			Varchar(36) not null,
		name 	       			Varchar(100) not null ,
		exercises			    jsonb,
		interval 				Varchar(36) not null,
		type 				    integer default 0,
		status 				    integer default 0,
		date      		     	Timestamp not null,
		createdat      			Timestamp not null,
		updatedat     			Timestamp not null,
		Primary Key (uid)
	);
`
