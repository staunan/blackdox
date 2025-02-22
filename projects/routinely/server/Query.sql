use routinely;
select * from routines;

-- Insert Data
insert into routines (routine_title) values ("Brush my teeth!");
insert into routines (routine_title, routine_description) values ("Have Bath!", "Take a hot shower");

-- Create Schema
create table routines (
	id bigint primary key auto_increment,
    user_id bigint default 1,
    routine_title varchar(200) default "",
    routine_description varchar(5000) default "",
    routine_mode varchar(50) default "Daily", -- Daily, Weekly, Monthly or Yearly
    routine_time time default '00:00:00',
    is_trash tinyint default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp on update current_timestamp
);
