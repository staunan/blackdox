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
    daily_basis_days varchar(50) default "", -- Empty Value = Everyday
    weekly_basis_weekday varchar(10) default "", -- Empty Value = Anyday of the Week
    monthly_basis_date int default 0, -- 0 = Anyday of the month, 32 = End of the Month, 33 = The day before end of the Month, 1 - 31 = date value
    yearly_basis_month_date varchar(10) default "", -- Empty String = Anyday of the year, 00-00 = Anyday of the year, Format = Month-Date
    routine_time time default '00:00:00', -- 00:00:00 = Anytime of the day, Format = Hour:Minute:Second
    is_trash tinyint default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp on update current_timestamp
);
