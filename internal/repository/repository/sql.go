package repository

const (
	InsertBatch = `
	insert into exchanger.course
	`

	GetStatistic = `
	select
		distinct on (exchanger, exchangers_condition_id)
		exchanger,
		exchangers_condition_id,
		course,
		time
	from exchanger.course
	order by time;
	`

	GetStatisticByExchangerByLastHourAVG = `
	select
    	avg(course.course) as course
	from exchanger.course
	where exchanger = $1 and exchangers_condition_id = $2 and time > $3;
	`

	GetStatisticByExchangerByLastHourMedian = `
	select
		median(course.course) as course
	from exchanger.course
	where exchanger = $1 and exchangers_condition_id = $2 and time > $3;
	`

	GetStatisticByAllExchangerAVG = `
	select
		avg(course) as course
	from (
		 select distinct on (exchanger, exchangers_condition_id)
		     course as course
		 from exchanger.course
		 order by time
	 );
	`

	GetStatisticByAllExchangerMedian = `
	select
		median(course) as course
	from (
		 select distinct on (exchanger, exchangers_condition_id)
		     course as course
		 from exchanger.course
		 order by time
	 );
	`
)
