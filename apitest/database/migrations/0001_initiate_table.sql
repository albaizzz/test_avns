drop table user

create table user(
	id int not null AUTO_INCREMENT,
	username varchar(25) not null,
	password varchar(100) not null,
	fullname varchar(100),
	email varchar(50),
	address varchar(150),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	primary key(id),
	UNIQUE KEY unique_user_username (username),
	UNIQUE KEY unique_user_email (email)
)

create UNIQUE INDEX unique_idx_user on user(username, email)

drop table user_session

create table user_session(
	id int not null AUTO_INCREMENT,
	user_id int not null,
	authorization varchar(100),
	created_date datetime,
	expired_date datetime,
	created_at timestamp  DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY key(id)
)



---------------------
drop table user

create table user(
	id int not null AUTO_INCREMENT,
	username varchar(25) not null,
	password varchar(100) not null,
	fullname varchar(100),
	email varchar(50),
	address varchar(150),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	primary key(id),
	UNIQUE KEY unique_user_username (username),
	UNIQUE KEY unique_user_email (email)
)

create UNIQUE INDEX unique_idx_user on user(username, email)

select * from user

select * from user_session

update user_session
set user_id =1


create table user_session(
	id int not null AUTO_INCREMENT,
	user_id int not null,
	session varchar(100),
	created_date timestamp  DEFAULT CURRENT_TIMESTAMP,
	expired_date timestamp,
	created_at timestamp  DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY key(id)
)

select us.user_id, us.session, created_date, expired_date, u.id, u.username, u.password, u.fullname, u.email, u.address  from user_session us 
inner join user u on us.user_id = u.id 
where us.session = 'e4ad4a04-622f-4b6f-9c97-927953d6b4f8'
