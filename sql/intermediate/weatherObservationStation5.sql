with CTE AS (
  select 
    ROW_NUMBER() OVER () AS idx, 
    city, 
    length(city) AS longchart 
  from 
    station 
  order by 
    length(city), city
), 
min_grp AS (
  select 
    min(idx) idx, 
    longchart 
  from 
    CTE 
  group by 
    longchart 
  limit 
    1
), 
max_grp as (
  select 
    min(idx) idx, 
    longchart 
  from 
    CTE 
  group by 
    longchart 
  order by 
    longchart desc 
  limit 
    1
) 
select 
  city, 
  longchart 
from 
  CTE 
where 
  idx in (
    select 
      idx 
    from 
      min_grp
  ) 
  or idx in (
    select 
      idx 
    from 
      max_grp
  );
